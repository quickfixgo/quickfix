package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
	"time"
)

type logonState struct{}

func (s logonState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	msgType := new(StringValue)
	if err := msg.Header.GetField(tag.MsgType, msgType); err == nil && msgType.Value == "A" {
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

func (s logonState) Timeout(session *session, e event) (nextState sessionState) {
	if e == logonTimeout {
		session.log.OnEvent("Timed out waiting for logon response")
		return latentState{}
	}

	return s
}

func (s logonState) handleLogon(session *session, msg Message) error {
	//reset on logon
	session.expectedSeqNum = 1
	session.seqNum = 1

	if err := session.verifyIgnoreSeqNumTooHigh(msg); err != nil {
		return err
	}

	reply := NewMessageBuilder()
	reply.Header.SetField(NewStringField(tag.MsgType, "A"))
	reply.Header.SetField(NewStringField(tag.BeginString, session.BeginString))
	reply.Header.SetField(NewStringField(tag.TargetCompID, session.TargetCompID))
	reply.Header.SetField(NewStringField(tag.SenderCompID, session.SenderCompID))
	reply.Body.SetField(NewIntField(tag.EncryptMethod, 0))

	heartBtInt := NewIntField(tag.HeartBtInt, 0)
	if err := msg.Body.Get(heartBtInt); err == nil {
		session.heartBeatTimeout = time.Duration(heartBtInt.Value) * time.Second
		reply.Body.SetField(heartBtInt)
	}

	if session.DefaultApplVerID != "" {
		reply.Trailer.SetField(NewStringField(tag.DefaultApplVerID, session.DefaultApplVerID))
	}

	session.log.OnEvent("Received logon request")
	session.send(reply)
	session.log.OnEvent("Responding to logon request")

	session.application.OnLogon(session.SessionID)

	if err := session.checkTargetTooHigh(msg); err != nil {
		switch TypedError := err.(type) {
		case TargetTooHigh:
			session.DoTargetTooHigh(TypedError)
		}
	}

	session.expectedSeqNum++
	return nil
}
