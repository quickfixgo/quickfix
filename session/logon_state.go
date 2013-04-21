package session

import (
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/tag"
	"time"
)

type logonState struct{}

func (s logonState) FixMsgIn(session *session, msg message.Message) (nextState state) {
	if msgType, ok := msg.Header.StringValue(tag.MsgType); ok && msgType == "A" {
		if err := s.handleLogon(session, msg); err != nil {
			session.log.OnEvent(err.Error())
			return latentState{}
		} else {
			return inSession{}
		}
	}

	session.log.OnEventf("Invalid Session State: Received Msg %v while waiting for Logon", msg)
	return latentState{}
}

func (s logonState) Timeout(session *session, e event) (nextState state) {
	if e == logonTimeout {
		session.log.OnEvent("Timed out waiting for logon response")
		return latentState{}
	}

	return s
}

func (s logonState) handleLogon(session *session, msg message.Message) error {
	//reset on logon
	session.expectedSeqNum = 1
	session.seqNum = 1

	if err := session.verifyIgnoreSeqNumTooHigh(msg); err != nil {
		return err
	}

	reply := message.NewMessage()
	reply.Header.SetField(message.NewStringField(tag.MsgType, "A"))
	reply.Header.SetField(message.NewStringField(tag.BeginString, session.BeginString))
	reply.Header.SetField(message.NewStringField(tag.TargetCompID, session.TargetCompID))
	reply.Header.SetField(message.NewStringField(tag.SenderCompID, session.SenderCompID))
	reply.Body.SetField(message.NewIntField(tag.EncryptMethod, 0))

	if HeartBtInt, err := msg.Body.IntValue(tag.HeartBtInt); err == nil {
		session.heartBeatTimeout = time.Duration(HeartBtInt) * time.Second
		reply.Body.SetField(message.NewIntField(tag.HeartBtInt, HeartBtInt))
	}

	if session.DefaultApplVerID != "" {
		reply.Trailer.SetField(message.NewStringField(tag.DefaultApplVerID, session.DefaultApplVerID))
	}

	session.log.OnEvent("Received logon request")
	session.send(reply)
	session.log.OnEvent("Responding to logon request")

	session.callback.OnLogon(session.ID)

	if err := session.checkTargetTooHigh(msg); err != nil {
		switch TypedError := err.(type) {
		case reject.TargetTooHigh:
			session.DoTargetTooHigh(TypedError)
		}
	}

	session.expectedSeqNum++
	return nil
}
