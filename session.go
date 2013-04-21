package quickfixgo

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/log"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session/store"
	"github.com/cbusbey/quickfixgo/settings"
	"github.com/cbusbey/quickfixgo/tag"
	"time"
)

type session struct {
	stateData
	store store.MessageStore

	log log.Log
	SessionID

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
		session.SessionID.BeginString = beginString
	} else {
		return settings.RequiredConfigurationMissing(settings.BeginString)
	}

	if targetCompID, ok := dict.GetString(settings.TargetCompID); ok {
		session.SessionID.TargetCompID = targetCompID
	} else {
		return settings.RequiredConfigurationMissing(settings.TargetCompID)
	}

	if senderCompID, ok := dict.GetString(settings.SenderCompID); ok {
		session.SessionID.SenderCompID = senderCompID
	} else {
		return settings.RequiredConfigurationMissing(settings.SenderCompID)
	}

	if session.SessionID.BeginString == fix.BeginString_FIXT11 {
		if defaultApplVerID, ok := dict.GetString(settings.DefaultApplVerID); ok {
			session.SessionID.DefaultApplVerID = defaultApplVerID
		} else {
			return settings.RequiredConfigurationMissing(settings.DefaultApplVerID)
		}
	}

	session.toSend = make(chan message.Builder)

	session.sessionEvent = make(chan event)
	session.log = logFactory.CreateSessionLog(session.SessionID.String())
	session.callback = callback
	session.stateTimer = eventTimer{Task: func() { session.sessionEvent <- needHeartbeat }}
	session.peerTimer = eventTimer{Task: func() { session.sessionEvent <- peerTimeout }}

	callback.OnCreate(session.SessionID)
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
		builder.SetHeaderField(message.NewUTCTimestampField(tag.SendingTime, sendingTime))
	} else {
		builder.SetHeaderField(message.NewUTCTimestampFieldNoMillis(tag.SendingTime, sendingTime))
	}
}

func (s *session) fillDefaultHeader(builder message.Builder) {
	builder.SetHeaderField(message.NewStringField(tag.BeginString, s.BeginString))
	builder.SetHeaderField(message.NewStringField(tag.SenderCompID, s.SenderCompID))
	builder.SetHeaderField(message.NewStringField(tag.TargetCompID, s.TargetCompID))

	s.insertSendingTime(builder)
}

func (s *session) resend(m *message.Message) {
	m.SetHeaderField(message.NewBooleanField(tag.PossDupFlag, true))

	if origSendingTime, ok := m.Header.Field(tag.SendingTime); ok {
		m.SetHeaderField(message.NewStringField(tag.OrigSendingTime, origSendingTime.Value()))
	}

	s.insertSendingTime(m)

	buffer := m.Build()
	s.sendBuffer(buffer)
}

func (s *session) send(builder message.Builder) {
	s.fillDefaultHeader(builder)
	builder.SetHeaderField(message.NewIntField(tag.MsgSeqNum, s.seqNum))

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
	resend := message.NewMessage()
	resend.Header.SetField(message.NewStringField(tag.MsgType, "2"))
	resend.Body.SetField(message.NewIntField(tag.BeginSeqNo, reject.ExpectedTarget))

	var endSeqNum = 0
	if s.BeginString < fix.BeginString_FIX42 {
		endSeqNum = 999999
	}
	resend.Body.SetField(message.NewIntField(tag.EndSeqNo, endSeqNum))

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
	if msgType, _ := msg.Header.StringValue(tag.MsgType); IsAdminMessageType(msgType) {
		return s.callback.FromAdmin(msg, s.SessionID)
	}

	return s.callback.FromApp(msg, s.SessionID)
}

func (s *session) checkTargetTooLow(msg message.Message) reject.MessageReject {
	switch seqNum, err := msg.Header.IntValue(tag.MsgSeqNum); {
	case err != nil:
		return reject.NewRequiredTagMissing(msg, tag.MsgSeqNum)
	case seqNum < s.expectedSeqNum:
		return reject.NewTargetTooLow(msg, seqNum, s.expectedSeqNum)
	}

	return nil
}

func (s *session) checkTargetTooHigh(msg message.Message) reject.MessageReject {
	switch seqNum, err := msg.Header.IntValue(tag.MsgSeqNum); {
	case err != nil:
		return reject.NewRequiredTagMissing(msg, tag.MsgSeqNum)
	case seqNum > s.expectedSeqNum:
		return reject.NewTargetTooHigh(msg, seqNum, s.expectedSeqNum)
	}

	return nil
}

func (s *session) checkCompID(msg message.Message) reject.MessageReject {
	SenderCompID, haveSender := msg.Header.Field(tag.SenderCompID)
	TargetCompID, haveTarget := msg.Header.Field(tag.TargetCompID)

	switch {
	case !haveSender:
		return reject.NewRequiredTagMissing(msg, tag.SenderCompID)
	case !haveTarget:
		return reject.NewRequiredTagMissing(msg, tag.TargetCompID)
	case s.SenderCompID != TargetCompID.Value() || s.TargetCompID != SenderCompID.Value():
		return reject.NewCompIDProblem(msg)
	}

	return nil
}

func (s *session) checkSendingTime(msg message.Message) reject.MessageReject {
	if sendingTime, err := msg.Header.UTCTimestampValue(tag.SendingTime); err != nil {
		return reject.NewRequiredTagMissing(msg, tag.SendingTime)
	} else {
		if delta := time.Since(sendingTime); delta <= -1*time.Duration(120)*time.Second || delta >= time.Duration(120)*time.Second {
			return reject.NewSendingTimeAccuracyProblem(msg)
		}
	}

	return nil
}

func (s *session) checkBeginString(msg message.Message) reject.MessageReject {
	switch BeginString, ok := msg.Header.StringValue(tag.BeginString); {
	case !ok:
		return reject.NewRequiredTagMissing(msg, tag.BeginString)
	case s.BeginString != BeginString:
		return reject.NewIncorrectBeginString(msg)
	}

	return nil
}

func (s *session) doReject(rej reject.MessageReject) {
	reply := message.NewMessage()

	if s.BeginString >= fix.BeginString_FIX42 {

		if rej.IsBusinessReject() {
			reply.Header.SetField(message.NewStringField(tag.MsgType, "j"))
			reply.Body.SetField(message.NewIntField(tag.BusinessRejectReason, int(rej.RejectReason())))
		} else {
			reply.Header.SetField(message.NewStringField(tag.MsgType, "3"))
			reply.Body.SetField(message.NewIntField(tag.SessionRejectReason, int(rej.RejectReason())))
		}
		reply.Body.SetField(message.NewStringField(tag.Text, rej.Error()))

		if MsgType, ok := rej.RejectedMessage().Header.StringValue(tag.MsgType); ok {
			reply.Body.SetField(message.NewStringField(tag.RefMsgType, MsgType))
		}

		switch rej.RejectReason() {
		case reject.RequiredTagMissing:
			reply.Body.SetField(message.NewIntField(tag.RefTagID, int(rej.RefTagID())))
		}
	} else {
		reply.Header.SetField(message.NewStringField(tag.MsgType, "3"))

		switch rej.RejectReason() {
		case reject.RequiredTagMissing:
			reply.Body.SetField(message.NewStringField(tag.Text, fmt.Sprintf("%s (%d)", rej.Error(), rej.RefTagID())))
		default:
			reply.Body.SetField(message.NewStringField(tag.Text, rej.Error()))
		}
	}

	if SeqNum, err := rej.RejectedMessage().Header.IntValue(tag.MsgSeqNum); err == nil {
		reply.Body.SetField(message.NewIntField(tag.RefSeqNum, SeqNum))
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
				if msg, err := message.MessageFromParsedBytes(msgBytes); err != nil {
					s.log.OnEventf("Msg Parse Error: %s, %s", err.Error(), msgBytes)
				} else {
					s.currentState = s.currentState.FixMsgIn(s, *msg)
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
