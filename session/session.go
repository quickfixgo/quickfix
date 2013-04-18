//fix session specific package 
package session

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/log"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/message/basic"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session/store"
	"github.com/cbusbey/quickfixgo/settings"
	"time"
)

type session struct {
	stateData
	store store.MessageStore

	log log.Log
	ID

	messageOut chan message.Buffer
	toSend     chan message.Builder

	sessionEvent chan event
	callback     Callback
	currentState state
	stateTimer   eventTimer
	peerTimer    eventTimer
	messageStash map[int]message.Message
}

//Creates Session, associates with internal session registry
func Create(dict settings.Dictionary, logFactory log.LogFactory, callback Callback) error {
	session := new(session)

	if beginString, ok := dict.GetString(settings.BeginString); ok {
		session.ID.BeginString = beginString
	} else {
		return settings.RequiredConfigurationMissing(settings.BeginString)
	}

	if targetCompID, ok := dict.GetString(settings.TargetCompID); ok {
		session.ID.TargetCompID = targetCompID
	} else {
		return settings.RequiredConfigurationMissing(settings.TargetCompID)
	}

	if senderCompID, ok := dict.GetString(settings.SenderCompID); ok {
		session.ID.SenderCompID = senderCompID
	} else {
		return settings.RequiredConfigurationMissing(settings.SenderCompID)
	}

	if session.ID.BeginString == fix.BeginString_FIXT11 {
		if defaultApplVerID, ok := dict.GetString(settings.DefaultApplVerID); ok {
			session.ID.DefaultApplVerID = defaultApplVerID
		} else {
			return settings.RequiredConfigurationMissing(settings.DefaultApplVerID)
		}
	}

	session.toSend = make(chan message.Builder)

	session.sessionEvent = make(chan event)
	session.log = logFactory.CreateSessionLog(session.ID.String())
	session.callback = callback
	session.stateTimer = eventTimer{Task: func() { session.sessionEvent <- needHeartbeat }}
	session.peerTimer = eventTimer{Task: func() { session.sessionEvent <- peerTimeout }}

	callback.OnCreate(session.ID)
	sessions.newSession <- session

	return nil
}

func (s *session) accept() (chan message.Buffer, error) {
	s.currentState = logonState{}
	s.messageOut = make(chan message.Buffer)
	s.messageStash = make(map[int]message.Message)
	s.store = store.NewMemoryStore()

	return s.messageOut, nil
}

func (s *session) onDisconnect() {
	s.log.OnEvent("Disconnected")
}

func (s *session) insertSendingTime(builder message.Builder) {
	sendingTime := time.Now().UTC()

	if s.BeginString >= fix.BeginString_FIX42 {
		builder.SetHeaderField(basic.NewUTCTimestampField(fix.SendingTime, sendingTime))
	} else {
		builder.SetHeaderField(basic.NewUTCTimestampFieldNoMillis(fix.SendingTime, sendingTime))
	}
}

func (s *session) fillDefaultHeader(builder message.Builder) {
	builder.SetHeaderField(basic.NewStringField(fix.BeginString, s.BeginString))
	builder.SetHeaderField(basic.NewStringField(fix.SenderCompID, s.SenderCompID))
	builder.SetHeaderField(basic.NewStringField(fix.TargetCompID, s.TargetCompID))

	s.insertSendingTime(builder)
}

func (s *session) resend(message *basic.Message) {
	message.SetHeaderField(basic.NewBooleanField(fix.PossDupFlag, true))

	if origSendingTime, ok := message.MsgHeader.Field(fix.SendingTime); ok {
		message.SetHeaderField(basic.NewStringField(fix.OrigSendingTime, origSendingTime.Value()))
	}

	s.insertSendingTime(message)

	buffer := message.Build()
	s.sendBuffer(buffer)
}

func (s *session) send(builder message.Builder) {
	s.fillDefaultHeader(builder)
	builder.SetHeaderField(basic.NewIntField(fix.MsgSeqNum, s.seqNum))

	buffer := builder.Build()
	s.store.SaveMessage(s.seqNum, buffer)
	s.sendBuffer(buffer)
	s.seqNum++
}

func (s *session) sendBuffer(buffer message.Buffer) {
	s.log.OnOutgoing(string(buffer.Bytes()))
	s.messageOut <- buffer
	s.stateTimer.Reset(time.Duration(s.heartBeatTimeout))
}

func (s *session) DoTargetTooHigh(reject reject.TargetTooHigh) {
	resend := basic.NewMessage()
	resend.MsgHeader.SetField(basic.NewStringField(fix.MsgType, "2"))
	resend.MsgBody.SetField(basic.NewIntField(fix.BeginSeqNo, reject.ExpectedTarget))

	var endSeqNum = 0
	if s.ID.BeginString < fix.BeginString_FIX42 {
		endSeqNum = 999999
	}
	resend.MsgBody.SetField(basic.NewIntField(fix.EndSeqNo, endSeqNum))

	s.send(resend)
}

func (s *session) verify(msg message.Message) reject.MessageReject {
	return s.verifySelect(msg, true, true)
}

func (s *session) verifyIgnoreSeqNumTooHigh(msg message.Message) reject.MessageReject {
	return s.verifySelect(msg, false, true)
}

func (s *session) verifyIgnoreSeqNumTooHighOrLow(msg message.Message) reject.MessageReject {
	return s.verifySelect(msg, false, false)
}

func (s *session) verifySelect(msg message.Message, checkTooHigh bool, checkTooLow bool) reject.MessageReject {
	if err := s.checkBeginString(msg); err != nil {
		return err
	}

	if err := s.checkCompID(msg); err != nil {
		return err
	}

	if err := s.checkSendingTime(msg); err != nil {
		return err
	}

	if checkTooLow {
		if err := s.checkTargetTooLow(msg); err != nil {
			return err
		}
	}

	if checkTooHigh {
		if err := s.checkTargetTooHigh(msg); err != nil {
			return err
		}
	}

	return s.fromCallback(msg)
}

func IsAdminMessageType(msgType string) bool {
	switch msgType {
	case "0", "A", "1", "2", "3", "4", "5":
		return true
	}

	return false
}

func (s *session) fromCallback(msg message.Message) reject.MessageReject {
	if msgType, _ := msg.Header().StringValue(fix.MsgType); IsAdminMessageType(msgType) {
		return s.callback.FromAdmin(msg, s.ID)
	}

	return s.callback.FromApp(msg, s.ID)
}

func (s *session) checkTargetTooLow(msg message.Message) reject.MessageReject {
	switch seqNum, err := msg.Header().IntValue(fix.MsgSeqNum); {
	case err != nil:
		return reject.NewRequiredTagMissing(msg, fix.MsgSeqNum)
	case seqNum < s.expectedSeqNum:
		return reject.NewTargetTooLow(msg, seqNum, s.expectedSeqNum)
	}

	return nil
}

func (s *session) checkTargetTooHigh(msg message.Message) reject.MessageReject {
	switch seqNum, err := msg.Header().IntValue(fix.MsgSeqNum); {
	case err != nil:
		return reject.NewRequiredTagMissing(msg, fix.MsgSeqNum)
	case seqNum > s.expectedSeqNum:
		return reject.NewTargetTooHigh(msg, seqNum, s.expectedSeqNum)
	}

	return nil
}

func (s *session) checkCompID(msg message.Message) reject.MessageReject {
	SenderCompID, haveSender := msg.Header().Field(fix.SenderCompID)
	TargetCompID, haveTarget := msg.Header().Field(fix.TargetCompID)

	switch {
	case !haveSender:
		return reject.NewRequiredTagMissing(msg, fix.SenderCompID)
	case !haveTarget:
		return reject.NewRequiredTagMissing(msg, fix.TargetCompID)
	case s.ID.SenderCompID != TargetCompID.Value() || s.ID.TargetCompID != SenderCompID.Value():
		return reject.NewCompIDProblem(msg)
	}

	return nil
}

func (s *session) checkSendingTime(msg message.Message) reject.MessageReject {
	if sendingTime, err := msg.Header().UTCTimestampValue(fix.SendingTime); err != nil {
		return reject.NewRequiredTagMissing(msg, fix.SendingTime)
	} else {
		if delta := time.Since(sendingTime); delta <= -1*time.Duration(120)*time.Second || delta >= time.Duration(120)*time.Second {
			return reject.NewSendingTimeAccuracyProblem(msg)
		}
	}

	return nil
}

func (s *session) checkBeginString(msg message.Message) reject.MessageReject {
	switch BeginString, ok := msg.Header().StringValue(fix.BeginString); {
	case !ok:
		return reject.NewRequiredTagMissing(msg, fix.BeginString)
	case s.BeginString != BeginString:
		return reject.NewIncorrectBeginString(msg)
	}

	return nil
}

func (s *session) doReject(rej reject.MessageReject) {
	reply := basic.NewMessage()

	if s.BeginString >= fix.BeginString_FIX42 {

		if rej.IsBusinessReject() {
			reply.MsgHeader.SetField(basic.NewStringField(fix.MsgType, "j"))
			reply.MsgBody.SetField(basic.NewIntField(fix.BusinessRejectReason, int(rej.RejectReason())))
		} else {
			reply.MsgHeader.SetField(basic.NewStringField(fix.MsgType, "3"))
			reply.MsgBody.SetField(basic.NewIntField(fix.SessionRejectReason, int(rej.RejectReason())))
		}
		reply.MsgBody.SetField(basic.NewStringField(fix.Text, rej.Error()))

		if MsgType, ok := rej.RejectedMessage().Header().StringValue(fix.MsgType); ok {
			reply.MsgBody.SetField(basic.NewStringField(fix.RefMsgType, MsgType))
		}

		switch rej.RejectReason() {
		case reject.RequiredTagMissing:
			reply.MsgBody.SetField(basic.NewIntField(fix.RefTagID, int(rej.RefTagID())))
		}
	} else {
		reply.MsgHeader.SetField(basic.NewStringField(fix.MsgType, "3"))

		switch rej.RejectReason() {
		case reject.RequiredTagMissing:
			reply.MsgBody.SetField(basic.NewStringField(fix.Text, fmt.Sprintf("%s (%d)", rej.Error(), rej.RefTagID())))
		default:
			reply.MsgBody.SetField(basic.NewStringField(fix.Text, rej.Error()))
		}
	}

	if SeqNum, err := rej.RejectedMessage().Header().IntValue(fix.MsgSeqNum); err == nil {
		reply.MsgBody.SetField(basic.NewIntField(fix.RefSeqNum, SeqNum))
	}

	s.send(reply)
	s.log.OnEventf("Message Rejected: %s", rej.Error())
}

func (s *session) run(msgIn chan []byte) {
	defer func() {
		close(s.messageOut)
	}()

	for {

		switch s.currentState.(type) {
		case latentState:
			return
		}

		select {
		case msgBytes, ok := <-msgIn:
			if ok {
				s.log.OnIncoming(string(msgBytes))
				if msg, err := basic.MessageFromParsedBytes(msgBytes); err != nil {
					s.log.OnEventf("Msg Parse Error: %s, %s", err.Error(), msgBytes)
				} else {
					s.currentState = s.currentState.FixMsgIn(s, msg)
				}
			} else {
				s.onDisconnect()
				return
			}
			s.peerTimer.Reset(time.Duration(int64(1.2 * float64(s.heartBeatTimeout))))

		case msg := <-s.toSend:
			s.send(msg)

		case evt := <-s.sessionEvent:
			s.currentState = s.currentState.Timeout(s, evt)
		}
	}
}
