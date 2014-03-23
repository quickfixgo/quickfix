package quickfixgo

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/field"
	"github.com/cbusbey/quickfixgo/log"
	"github.com/cbusbey/quickfixgo/settings"
	"github.com/cbusbey/quickfixgo/tag"
	"time"
)

type session struct {
	stateData
	store MessageStore

	log log.Log
	SessionID

	messageOut chan Buffer
	toSend     chan MessageBuilder

	sessionEvent chan event
	application  Application
	currentState sessionState
	stateTimer   eventTimer
	peerTimer    eventTimer
	messageStash map[int]Message
	*DataDictionary
}

//Creates Session, associates with internal session registry
func createSession(dict settings.Dictionary, logFactory log.LogFactory, application Application) error {
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

	if dataDictionaryPath, ok := dict.GetString(settings.DataDictionary); ok {
		if dataDictionary, err := NewDataDictionary(dataDictionaryPath); err != nil {
			return err
		} else {
			session.DataDictionary = dataDictionary
		}
	}

	if session.SessionID.BeginString == BeginString_FIXT11 {
		if defaultApplVerID, ok := dict.GetString(settings.DefaultApplVerID); ok {
			session.SessionID.DefaultApplVerID = defaultApplVerID
		} else {
			return settings.RequiredConfigurationMissing(settings.DefaultApplVerID)
		}
	}

	session.toSend = make(chan MessageBuilder)

	session.sessionEvent = make(chan event)
	session.log = logFactory.CreateSessionLog(session.SessionID.String())
	session.application = application
	session.stateTimer = eventTimer{Task: func() { session.sessionEvent <- needHeartbeat }}
	session.peerTimer = eventTimer{Task: func() { session.sessionEvent <- peerTimeout }}

	application.OnCreate(session.SessionID)
	sessions.newSession <- session

	return nil
}

func (s *session) accept() (chan Buffer, error) {
	s.currentState = logonState{}
	s.messageOut = make(chan Buffer)
	s.messageStash = make(map[int]Message)
	s.store = NewMemoryStore()

	return s.messageOut, nil
}

func (s *session) onDisconnect() {
	s.log.OnEvent("Disconnected")
}

func (s *session) insertSendingTime(builder MessageBuilder) {
	sendingTime := time.Now().UTC()

	if s.BeginString >= BeginString_FIX42 {
		builder.SetHeaderField(field.NewUTCTimestampField(tag.SendingTime, sendingTime))
	} else {
		builder.SetHeaderField(field.NewUTCTimestampFieldNoMillis(tag.SendingTime, sendingTime))
	}
}

func (s *session) fillDefaultHeader(builder MessageBuilder) {
	builder.SetHeaderField(field.NewStringField(tag.BeginString, s.BeginString))
	builder.SetHeaderField(field.NewStringField(tag.SenderCompID, s.SenderCompID))
	builder.SetHeaderField(field.NewStringField(tag.TargetCompID, s.TargetCompID))

	s.insertSendingTime(builder)
}

func (s *session) resend(m *Message) {
	m.SetHeaderField(field.NewBooleanField(tag.PossDupFlag, true))

	//FIXME
	origSendingTime := field.NewStringField(tag.SendingTime, "")
	if err := m.Header.Get(origSendingTime); err == nil {
		m.SetHeaderField(field.NewStringField(tag.OrigSendingTime, origSendingTime.Value))
	}

	s.insertSendingTime(m)

	buffer := m.Build()
	s.sendBuffer(buffer)
}

func (s *session) send(builder MessageBuilder) {
	s.fillDefaultHeader(builder)
	builder.SetHeaderField(field.NewIntField(tag.MsgSeqNum, s.seqNum))

	buffer := builder.Build()
	s.store.SaveMessage(s.seqNum, buffer)
	s.sendBuffer(buffer)
	s.seqNum++
}

func (s *session) sendBuffer(buffer Buffer) {
	s.log.OnOutgoing(string(buffer.Bytes()))
	s.messageOut <- buffer
	s.stateTimer.Reset(time.Duration(s.heartBeatTimeout))
}

func (s *session) DoTargetTooHigh(reject TargetTooHigh) {
	resend := NewMessage()
	resend.Header.SetField(field.NewStringField(tag.MsgType, "2"))
	resend.Body.SetField(field.NewIntField(tag.BeginSeqNo, reject.ExpectedTarget))

	var endSeqNum = 0
	if s.BeginString < BeginString_FIX42 {
		endSeqNum = 999999
	}
	resend.Body.SetField(field.NewIntField(tag.EndSeqNo, endSeqNum))

	s.send(resend)
}

func (s *session) verify(msg Message) MessageReject {
	return s.verifySelect(msg, true, true)
}

func (s *session) verifyIgnoreSeqNumTooHigh(msg Message) MessageReject {
	return s.verifySelect(msg, false, true)
}

func (s *session) verifyIgnoreSeqNumTooHighOrLow(msg Message) MessageReject {
	return s.verifySelect(msg, false, false)
}

func (s *session) verifySelect(msg Message, checkTooHigh bool, checkTooLow bool) MessageReject {
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

	if s.DataDictionary != nil {
		if err := s.DataDictionary.validate(msg); err != nil {
			switch TypedError := err.(type) {
			case InvalidTagNumberError:
				return NewInvalidTagNumber(msg, TypedError.Tag)
			default:
				s.log.OnEventf("Error validating : %s", err.Error())
			}
		}
	}

	return s.fromCallback(msg)
}

func (s *session) fromCallback(msg Message) MessageReject {
	msgType := new(field.MsgType)
	if msg.Header.Get(msgType); msgType.IsAdminMessageType() {
		return s.application.FromAdmin(msg, s.SessionID)
	}

	return s.application.FromApp(msg, s.SessionID)
}

func (s *session) checkTargetTooLow(msg Message) MessageReject {
	seqNum := new(field.MsgSeqNum)
	switch err := msg.Header.Get(seqNum); {
	case err != nil:
		return NewRequiredTagMissing(msg, tag.MsgSeqNum)
	case seqNum.Value < s.expectedSeqNum:
		return NewTargetTooLow(msg, seqNum.Value, s.expectedSeqNum)
	}

	return nil
}

func (s *session) checkTargetTooHigh(msg Message) MessageReject {
	seqNum := new(field.MsgSeqNum)
	switch err := msg.Header.Get(seqNum); {
	case err != nil:
		return NewRequiredTagMissing(msg, tag.MsgSeqNum)
	case seqNum.Value > s.expectedSeqNum:
		return NewTargetTooHigh(msg, seqNum.Value, s.expectedSeqNum)
	}

	return nil
}

func (s *session) checkCompID(msg Message) MessageReject {
	senderCompID := new(field.SenderCompID)
	targetCompID := new(field.TargetCompID)
	haveSender := msg.Header.Get(senderCompID)
	haveTarget := msg.Header.Get(targetCompID)

	switch {
	case haveSender != nil:
		return NewRequiredTagMissing(msg, tag.SenderCompID)
	case haveTarget != nil:
		return NewRequiredTagMissing(msg, tag.TargetCompID)
	case s.SenderCompID != targetCompID.Value || s.TargetCompID != senderCompID.Value:
		return NewCompIDProblem(msg)
	}

	return nil
}

func (s *session) checkSendingTime(msg Message) MessageReject {
	sendingTime := new(field.SendingTime)
	if err := msg.Header.Get(sendingTime); err != nil {
		return NewRequiredTagMissing(msg, tag.SendingTime)
	} else {
		if delta := time.Since(sendingTime.Value); delta <= -1*time.Duration(120)*time.Second || delta >= time.Duration(120)*time.Second {
			return NewSendingTimeAccuracyProblem(msg)
		}
	}

	return nil
}

func (s *session) checkBeginString(msg Message) MessageReject {
	beginString := new(field.BeginString)
	switch err := msg.Header.Get(beginString); {
	case err != nil:
		return NewRequiredTagMissing(msg, tag.BeginString)
	case s.BeginString != beginString.Value:
		return NewIncorrectBeginString(msg)
	}

	return nil
}

func (s *session) doReject(rej MessageReject) {
	reply := NewMessage()

	if s.BeginString >= BeginString_FIX42 {

		if rej.IsBusinessReject() {
			reply.Header.SetField(field.NewStringField(tag.MsgType, "j"))
			reply.Body.SetField(field.NewIntField(tag.BusinessRejectReason, int(rej.RejectReason())))
		} else {
			reply.Header.SetField(field.NewStringField(tag.MsgType, "3"))
			reply.Body.SetField(field.NewIntField(tag.SessionRejectReason, int(rej.RejectReason())))
		}
		reply.Body.SetField(field.NewStringField(tag.Text, rej.Error()))

		msgType := new(field.MsgType)
		if err := rej.RejectedMessage().Header.Get(msgType); err == nil {
			reply.Body.SetField(field.NewStringField(tag.RefMsgType, msgType.Value))
		}

		if refTagId := rej.RefTagID(); refTagId != nil {
			reply.Body.SetField(field.NewIntField(tag.RefTagID, int(*rej.RefTagID())))
		}
	} else {
		reply.Header.SetField(field.NewStringField(tag.MsgType, "3"))

		if refTagId := rej.RefTagID(); refTagId != nil {
			reply.Body.SetField(field.NewStringField(tag.Text, fmt.Sprintf("%s (%d)", rej.Error(), *refTagId)))
		} else {
			reply.Body.SetField(field.NewStringField(tag.Text, rej.Error()))
		}
	}

	seqNum := new(field.MsgSeqNum)
	if err := rej.RejectedMessage().Header.Get(seqNum); err == nil {
		reply.Body.SetField(field.NewIntField(tag.RefSeqNum, seqNum.Value))
	}

	s.send(reply)
	s.log.OnEventf("Message Rejected: %s", rej.Error())
}

func (s *session) run(msgIn chan []byte) {
	defer func() {
		s.messageOut <- nil
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
				if msg, err := MessageFromParsedBytes(msgBytes); err != nil {
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
