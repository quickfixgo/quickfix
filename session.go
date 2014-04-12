package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
	"github.com/quickfixgo/quickfix/log"
	"github.com/quickfixgo/quickfix/message"
	"github.com/quickfixgo/quickfix/settings"
	"time"
)

type session struct {
	stateData
	store messageStore

	log log.Log
	SessionID

	messageOut chan buffer
	toSend     chan *message.MessageBuilder

	sessionEvent            chan event
	application             Application
	currentState            sessionState
	stateTimer              eventTimer
	peerTimer               eventTimer
	messageStash            map[int]message.Message
	dataDictionary          *datadictionary.DataDictionary
	transportDataDictionary *datadictionary.DataDictionary
	appDataDictionary       *datadictionary.DataDictionary
	resetOnLogon            bool
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
		var err error
		if session.dataDictionary, err = datadictionary.Parse(dataDictionaryPath); err != nil {
			return err
		}
	}

	if transportDataDictionaryPath, ok := dict.GetString(settings.TransportDataDictionary); ok {
		var err error
		if session.transportDataDictionary, err = datadictionary.Parse(transportDataDictionaryPath); err != nil {
			return err
		}
	}

	//FIXME: tDictionary w/o appDictionary and vice versa should throw config error
	if appDataDictionaryPath, ok := dict.GetString(settings.AppDataDictionary); ok {
		var err error
		if session.appDataDictionary, err = datadictionary.Parse(appDataDictionaryPath); err != nil {
			return err
		}
	}

	if session.SessionID.BeginString == fix.BeginString_FIXT11 {
		if defaultApplVerID, ok := dict.GetString(settings.DefaultApplVerID); ok {
			session.SessionID.DefaultApplVerID = defaultApplVerID
		} else {
			return settings.RequiredConfigurationMissing(settings.DefaultApplVerID)
		}
	}

	if resetOnLogon, ok := dict.GetBool(settings.ResetOnLogon); ok {
		session.resetOnLogon = resetOnLogon
	} else {
		session.resetOnLogon = false
	}

	session.toSend = make(chan *message.MessageBuilder)

	session.sessionEvent = make(chan event)
	session.log = logFactory.CreateSessionLog(session.SessionID.String())
	session.application = application
	session.stateTimer = eventTimer{Task: func() { session.sessionEvent <- needHeartbeat }}
	session.peerTimer = eventTimer{Task: func() { session.sessionEvent <- peerTimeout }}

	application.OnCreate(session.SessionID)
	sessions.newSession <- session

	return nil
}

func (s *session) accept() (chan buffer, error) {
	s.currentState = logonState{}
	s.messageOut = make(chan buffer)
	s.messageStash = make(map[int]message.Message)
	s.store = newMemoryStore()

	return s.messageOut, nil
}

func (s *session) onDisconnect() {
	s.log.OnEvent("Disconnected")
}

func (s *session) insertSendingTime(builder *message.MessageBuilder) {
	sendingTime := time.Now().UTC()

	if s.BeginString >= fix.BeginString_FIX42 {
		builder.Header.Set(message.NewUTCTimestampField(tag.SendingTime, sendingTime))
	} else {
		builder.Header.Set(message.NewUTCTimestampFieldNoMillis(tag.SendingTime, sendingTime))
	}
}

func (s *session) fillDefaultHeader(builder *message.MessageBuilder) {
	builder.Header.Set(message.NewStringField(tag.BeginString, s.BeginString))
	builder.Header.Set(message.NewStringField(tag.SenderCompID, s.SenderCompID))
	builder.Header.Set(message.NewStringField(tag.TargetCompID, s.TargetCompID))

	s.insertSendingTime(builder)
}

func (s *session) resend(m *message.MessageBuilder) {
	m.Header.Set(message.NewBooleanField(tag.PossDupFlag, true))

	origSendingTime := new(message.StringValue)
	if err := m.Header.GetField(tag.SendingTime, origSendingTime); err == nil {
		m.Header.Set(message.NewStringField(tag.OrigSendingTime, origSendingTime.Value))
	}

	s.insertSendingTime(m)

	if buffer, err := m.Build(); err != nil {
		panic(err)
	} else {
		s.sendBuffer(buffer)
	}
}

func (s *session) send(builder *message.MessageBuilder) {
	s.fillDefaultHeader(builder)
	builder.Header.Set(message.NewIntField(tag.MsgSeqNum, s.seqNum))

	if buffer, err := builder.Build(); err != nil {
		panic(err)
	} else {
		s.store.SaveMessage(s.seqNum, buffer)
		s.sendBuffer(buffer)
		s.seqNum++
	}
}

func (s *session) sendBuffer(buffer buffer) {

	s.log.OnOutgoing(string(buffer.Bytes()))
	s.messageOut <- buffer
	s.stateTimer.Reset(time.Duration(s.heartBeatTimeout))
}

func (s *session) DoTargetTooHigh(reject message.TargetTooHigh) {
	resend := message.NewMessageBuilder()
	resend.Header.Set(message.NewStringField(tag.MsgType, "2"))
	resend.Body.Set(message.NewIntField(tag.BeginSeqNo, reject.ExpectedTarget))

	var endSeqNum = 0
	if s.BeginString < fix.BeginString_FIX42 {
		endSeqNum = 999999
	}
	resend.Body.Set(message.NewIntField(tag.EndSeqNo, endSeqNum))

	s.send(resend)
}

func (s *session) handleLogon(msg message.Message) error {
	s.log.OnEvent("Received logon request")
	if s.resetOnLogon {
		s.expectedSeqNum = 1
		s.seqNum = 1
	}

	resetSeqNumFlag := new(message.BooleanValue)
	if err := msg.Body.GetField(tag.ResetSeqNumFlag, resetSeqNumFlag); err == nil {
		if resetSeqNumFlag.Value {
			s.log.OnEvent("Logon contains ResetSeqNumFlag=Y, resetting sequence numbers to 1")
			s.seqNum = 1
			s.expectedSeqNum = 1
		}
	}

	if err := s.verifyIgnoreSeqNumTooHigh(msg); err != nil {
		return err
	}

	reply := message.NewMessageBuilder()
	reply.Header.Set(message.NewStringField(tag.MsgType, "A"))
	reply.Header.Set(message.NewStringField(tag.BeginString, s.BeginString))
	reply.Header.Set(message.NewStringField(tag.TargetCompID, s.TargetCompID))
	reply.Header.Set(message.NewStringField(tag.SenderCompID, s.SenderCompID))
	reply.Body.Set(message.NewIntField(tag.EncryptMethod, 0))

	heartBtInt := message.NewIntField(tag.HeartBtInt, 0)
	if err := msg.Body.Get(heartBtInt); err == nil {
		s.heartBeatTimeout = time.Duration(heartBtInt.Value) * time.Second
		reply.Body.Set(heartBtInt)
	}

	if resetSeqNumFlag.Value {
		reply.Body.Set(message.NewBooleanField(tag.ResetSeqNumFlag, resetSeqNumFlag.Value))
	}

	if s.DefaultApplVerID != "" {
		reply.Trailer.Set(message.NewStringField(tag.DefaultApplVerID, s.DefaultApplVerID))
	}

	s.log.OnEvent("Responding to logon request")
	s.send(reply)

	s.application.OnLogon(s.SessionID)

	if err := s.checkTargetTooHigh(msg); err != nil {
		switch TypedError := err.(type) {
		case message.TargetTooHigh:
			s.DoTargetTooHigh(TypedError)
		}
	}

	s.expectedSeqNum++
	return nil
}

func (s *session) verify(msg message.Message) message.MessageReject {
	return s.verifySelect(msg, true, true)
}

func (s *session) verifyIgnoreSeqNumTooHigh(msg message.Message) message.MessageReject {
	return s.verifySelect(msg, false, true)
}

func (s *session) verifyIgnoreSeqNumTooHighOrLow(msg message.Message) message.MessageReject {
	return s.verifySelect(msg, false, false)
}

func (s *session) verifySelect(msg message.Message, checkTooHigh bool, checkTooLow bool) message.MessageReject {
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

	if s.dataDictionary != nil {
		if err := message.Validate(s.dataDictionary, msg); err != nil {
			return err
		}
	}

	if s.transportDataDictionary != nil {
		msgType := new(message.StringValue)
		msg.Header.GetField(tag.MsgType, msgType)
		if fix.IsAdminMessageType(msgType.Value) {
			if err := message.Validate(s.transportDataDictionary, msg); err != nil {
				return err
			}
		} else {
			if err := message.ValidateFIXTApp(s.transportDataDictionary, s.appDataDictionary, msg); err != nil {
				return err
			}
		}
	}

	return s.fromCallback(msg)
}

func (s *session) fromCallback(msg message.Message) message.MessageReject {
	msgType := new(message.StringValue)
	if msg.Header.GetField(tag.MsgType, msgType); fix.IsAdminMessageType(msgType.Value) {
		return s.application.FromAdmin(msg, s.SessionID)
	}

	return s.application.FromApp(msg, s.SessionID)
}

func (s *session) checkTargetTooLow(msg message.Message) message.MessageReject {
	seqNum := new(message.IntField)
	switch err := msg.Header.GetField(tag.MsgSeqNum, seqNum); {
	case err != nil:
		return message.NewRequiredTagMissing(msg, tag.MsgSeqNum)
	case seqNum.Value < s.expectedSeqNum:
		return message.NewTargetTooLow(msg, seqNum.Value, s.expectedSeqNum)
	}

	return nil
}

func (s *session) checkTargetTooHigh(msg message.Message) message.MessageReject {
	seqNum := new(message.IntValue)
	switch err := msg.Header.GetField(tag.MsgSeqNum, seqNum); {
	case err != nil:
		return message.NewRequiredTagMissing(msg, tag.MsgSeqNum)
	case seqNum.Value > s.expectedSeqNum:
		return message.NewTargetTooHigh(msg, seqNum.Value, s.expectedSeqNum)
	}

	return nil
}

func (s *session) checkCompID(msg message.Message) message.MessageReject {
	senderCompID := new(message.StringValue)
	targetCompID := new(message.StringValue)
	haveSender := msg.Header.GetField(tag.SenderCompID, senderCompID)
	haveTarget := msg.Header.GetField(tag.TargetCompID, targetCompID)

	switch {
	case haveSender != nil:
		return message.NewRequiredTagMissing(msg, tag.SenderCompID)
	case haveTarget != nil:
		return message.NewRequiredTagMissing(msg, tag.TargetCompID)
	case len(targetCompID.Value) == 0:
		return message.NewTagSpecifiedWithoutAValue(msg, tag.TargetCompID)
	case len(senderCompID.Value) == 0:
		return message.NewTagSpecifiedWithoutAValue(msg, tag.SenderCompID)
	case s.SenderCompID != targetCompID.Value || s.TargetCompID != senderCompID.Value:
		return message.NewCompIDProblem(msg)
	}

	return nil
}

func (s *session) checkSendingTime(msg message.Message) message.MessageReject {
	sendingTime := new(message.UTCTimestampValue)
	if err := msg.Header.GetField(tag.SendingTime, sendingTime); err != nil {
		return message.NewRequiredTagMissing(msg, tag.SendingTime)
	}

	if delta := time.Since(sendingTime.Value); delta <= -1*time.Duration(120)*time.Second || delta >= time.Duration(120)*time.Second {
		return message.NewSendingTimeAccuracyProblem(msg)
	}

	return nil
}

func (s *session) checkBeginString(msg message.Message) message.MessageReject {
	beginString := new(message.StringValue)
	switch err := msg.Header.GetField(tag.BeginString, beginString); {
	case err != nil:
		return message.NewRequiredTagMissing(msg, tag.BeginString)
	case s.BeginString != beginString.Value:
		return message.NewIncorrectBeginString(msg)
	}

	return nil
}

func (s *session) doReject(rej message.MessageReject) {
	reply := rej.RejectedMessage().ReverseRoute()

	if s.BeginString >= fix.BeginString_FIX42 {

		if rej.IsBusinessReject() {
			reply.Header.Set(message.NewStringField(tag.MsgType, "j"))
			reply.Body.Set(message.NewIntField(tag.BusinessRejectReason, int(rej.RejectReason())))
		} else {
			reply.Header.Set(message.NewStringField(tag.MsgType, "3"))
			switch {
			default:
				reply.Body.Set(message.NewIntField(tag.SessionRejectReason, int(rej.RejectReason())))
			case rej.RejectReason() > message.InvalidMsgType && s.BeginString == fix.BeginString_FIX42:
				//fix42 knows up to invalid msg type
			}
		}
		reply.Body.Set(message.NewStringField(tag.Text, rej.Error()))

		msgType := new(message.StringValue)
		if err := rej.RejectedMessage().Header.GetField(tag.MsgType, msgType); err == nil {
			reply.Body.Set(message.NewStringField(tag.RefMsgType, msgType.Value))
		}

		if refTagID := rej.RefTagID(); refTagID != nil {
			reply.Body.Set(message.NewIntField(tag.RefTagID, int(*refTagID)))
		}
	} else {
		reply.Header.Set(message.NewStringField(tag.MsgType, "3"))

		if refTagID := rej.RefTagID(); refTagID != nil {
			reply.Body.Set(message.NewStringField(tag.Text, fmt.Sprintf("%s (%d)", rej.Error(), *refTagID)))
		} else {
			reply.Body.Set(message.NewStringField(tag.Text, rej.Error()))
		}
	}

	seqNum := new(message.IntValue)
	if err := rej.RejectedMessage().Header.GetField(tag.MsgSeqNum, seqNum); err == nil {
		reply.Body.Set(message.NewIntField(tag.RefSeqNum, seqNum.Value))
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
