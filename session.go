package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/config"
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/fix/tag"
	"github.com/quickfixgo/quickfix/message"
	"time"
)

//The Session is the primary FIX abstraction for message communication
type Session struct {
	stateData
	store messageStore

	log       Log
	sessionID SessionID

	messageOut chan buffer
	toSend     chan message.MessageBuilder

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
	initiateLogon           bool
	heartBtInt              int

	//required on logon for FIX.T.1 messages
	defaultApplVerID       string
	targetDefaultApplVerID string
}

//TargetDefaultApplicationVersionID returns the default application version ID for messages received by this version.
//Applicable for For FIX.T.1 sessions.
func (s *Session) TargetDefaultApplicationVersionID() string {
	return s.targetDefaultApplVerID
}

//Creates Session, associates with internal session registry
func createSession(sessionID SessionID, settings *SessionSettings, logFactory LogFactory, application Application) error {
	session := &Session{sessionID: sessionID}

	if sessionID.BeginString == fix.BeginString_FIXT11 {
		defaultApplVerID, err := settings.Setting(config.DefaultApplVerID)
		if err != nil {
			return errors.RequiredConfigurationMissing(config.DefaultApplVerID)
		}
		session.defaultApplVerID = defaultApplVerID
	}

	if dataDictionaryPath, err := settings.Setting(config.DataDictionary); err == nil {
		if session.dataDictionary, err = datadictionary.Parse(dataDictionaryPath); err != nil {
			return err
		}
	}

	if transportDataDictionaryPath, err := settings.Setting(config.TransportDataDictionary); err == nil {
		if session.transportDataDictionary, err = datadictionary.Parse(transportDataDictionaryPath); err != nil {
			return err
		}
	}

	//FIXME: tDictionary w/o appDictionary and vice versa should throw config error
	if appDataDictionaryPath, err := settings.Setting(config.AppDataDictionary); err == nil {
		if session.appDataDictionary, err = datadictionary.Parse(appDataDictionaryPath); err != nil {
			return err
		}
	}

	var err error
	if settings.HasSetting(config.ResetOnLogon) {
		if session.resetOnLogon, err = settings.BoolSetting(config.ResetOnLogon); err != nil {
			return err
		}
	}

	if settings.HasSetting(config.HeartBtInt) {
		if session.heartBtInt, err = settings.IntSetting(config.HeartBtInt); err != nil {
			return err
		}
	}

	if session.log, err = logFactory.CreateSessionLog(session.sessionID); err != nil {
		return err
	}

	session.toSend = make(chan message.MessageBuilder)
	session.sessionEvent = make(chan event)
	session.application = application
	session.stateTimer = eventTimer{Task: func() { session.sessionEvent <- needHeartbeat }}
	session.peerTimer = eventTimer{Task: func() { session.sessionEvent <- peerTimeout }}

	application.OnCreate(session.sessionID)
	sessions.newSession <- session

	return nil
}

func (s *Session) initiate() (chan buffer, error) {
	s.currentState = logonState{}
	s.messageOut = make(chan buffer)
	s.messageStash = make(map[int]message.Message)
	s.store = newMemoryStore()
	s.initiateLogon = true

	return s.messageOut, nil
}

func (s *Session) accept() (chan buffer, error) {
	s.currentState = logonState{}
	s.messageOut = make(chan buffer)
	s.messageStash = make(map[int]message.Message)
	s.store = newMemoryStore()

	return s.messageOut, nil
}

func (s *Session) onDisconnect() {
	s.log.OnEvent("Disconnected")
}

func (s *Session) insertSendingTime(builder message.MessageBuilder) {
	sendingTime := time.Now().UTC()

	if s.sessionID.BeginString >= fix.BeginString_FIX42 {
		builder.Header.Set(message.NewUTCTimestampField(tag.SendingTime, sendingTime))
	} else {
		builder.Header.Set(message.NewUTCTimestampFieldNoMillis(tag.SendingTime, sendingTime))
	}
}

func (s *Session) fillDefaultHeader(builder message.MessageBuilder) {
	builder.Header.Set(field.NewBeginString(s.sessionID.BeginString))
	builder.Header.Set(field.NewSenderCompID(s.sessionID.SenderCompID))
	builder.Header.Set(field.NewTargetCompID(s.sessionID.TargetCompID))

	s.insertSendingTime(builder)
}

func (s *Session) resend(m message.MessageBuilder) {
	m.Header.Set(field.NewPossDupFlag(true))

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

func (s *Session) send(builder message.MessageBuilder) {
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

func (s *Session) sendBuffer(buffer buffer) {

	s.log.OnOutgoing(string(buffer.Bytes()))
	s.messageOut <- buffer
	s.stateTimer.Reset(time.Duration(s.heartBeatTimeout))
}

func (s *Session) doTargetTooHigh(reject errors.TargetTooHigh) {
	resend := message.CreateMessageBuilder()
	resend.Header.Set(field.NewMsgType("2"))
	resend.Body.Set(field.NewBeginSeqNo(reject.ExpectedTarget))

	var endSeqNum = 0
	if s.sessionID.BeginString < fix.BeginString_FIX42 {
		endSeqNum = 999999
	}
	resend.Body.Set(field.NewEndSeqNo(endSeqNum))

	s.send(resend)
}

func (s *Session) handleLogon(msg message.Message) error {
	//Grab default app ver id from fixt.1.1 logon
	if s.sessionID.BeginString == fix.BeginString_FIXT11 {
		targetApplVerID := &field.DefaultApplVerIDField{}

		if err := msg.Body.Get(targetApplVerID); err != nil {
			return err
		}

		s.targetDefaultApplVerID = targetApplVerID.Value
	}

	if !s.initiateLogon {
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

		reply := message.CreateMessageBuilder()
		reply.Header.Set(field.NewMsgType("A"))
		reply.Header.Set(field.NewBeginString(s.sessionID.BeginString))
		reply.Header.Set(field.NewTargetCompID(s.sessionID.TargetCompID))
		reply.Header.Set(field.NewSenderCompID(s.sessionID.SenderCompID))
		reply.Body.Set(field.NewEncryptMethod(0))

		var heartBtInt field.HeartBtIntField
		if err := msg.Body.Get(&heartBtInt); err == nil {
			s.heartBeatTimeout = time.Duration(heartBtInt.Value) * time.Second
			reply.Body.Set(heartBtInt)
		}

		if resetSeqNumFlag.Value {
			reply.Body.Set(field.NewResetSeqNumFlag(resetSeqNumFlag.Value))
		}

		if len(s.defaultApplVerID) > 0 {
			reply.Body.Set(field.NewDefaultApplVerID(s.defaultApplVerID))
		}

		s.log.OnEvent("Responding to logon request")
		s.send(reply)
	}

	s.application.OnLogon(s.sessionID)

	if err := s.checkTargetTooHigh(msg); err != nil {
		switch TypedError := err.(type) {
		case errors.TargetTooHigh:
			s.doTargetTooHigh(TypedError)
		}
	}

	s.expectedSeqNum++
	return nil
}

func (s *Session) verify(msg message.Message) errors.MessageRejectError {
	return s.verifySelect(msg, true, true)
}

func (s *Session) verifyIgnoreSeqNumTooHigh(msg message.Message) errors.MessageRejectError {
	return s.verifySelect(msg, false, true)
}

func (s *Session) verifyIgnoreSeqNumTooHighOrLow(msg message.Message) errors.MessageRejectError {
	return s.verifySelect(msg, false, false)
}

func (s *Session) verifySelect(msg message.Message, checkTooHigh bool, checkTooLow bool) errors.MessageRejectError {
	if reject := s.checkBeginString(msg); reject != nil {
		return reject
	}

	if reject := s.checkCompID(msg); reject != nil {
		return reject
	}

	if reject := s.checkSendingTime(msg); reject != nil {
		return reject
	}

	if checkTooLow {
		if reject := s.checkTargetTooLow(msg); reject != nil {
			return reject
		}
	}

	if checkTooHigh {
		if reject := s.checkTargetTooHigh(msg); reject != nil {
			return reject
		}
	}

	if s.dataDictionary != nil {
		if reject := message.Validate(s.dataDictionary, msg); reject != nil {
			return reject
		}
	}

	if s.transportDataDictionary != nil {
		var msgType field.MsgTypeField
		msg.Header.Get(&msgType)
		if fix.IsAdminMessageType(msgType.Value) {
			if reject := message.Validate(s.transportDataDictionary, msg); reject != nil {
				return reject
			}
		} else {
			if reject := message.ValidateFIXTApp(s.transportDataDictionary, s.appDataDictionary, msg); reject != nil {
				return reject
			}
		}
	}

	return s.fromCallback(msg)
}

func (s *Session) fromCallback(msg message.Message) errors.MessageRejectError {
	msgType := new(message.StringValue)
	if msg.Header.GetField(tag.MsgType, msgType); fix.IsAdminMessageType(msgType.Value) {
		return s.application.FromAdmin(msg, s.sessionID)
	}

	return s.application.FromApp(msg, s.sessionID)
}

func (s *Session) checkTargetTooLow(msg message.Message) errors.MessageRejectError {
	seqNum := new(message.IntField)
	switch err := msg.Header.GetField(tag.MsgSeqNum, seqNum); {
	case err != nil:
		return errors.RequiredTagMissing(tag.MsgSeqNum)
	case seqNum.Value < s.expectedSeqNum:
		return errors.TargetTooLow{ReceivedTarget: seqNum.Value, ExpectedTarget: s.expectedSeqNum}
	}

	return nil
}

func (s *Session) checkTargetTooHigh(msg message.Message) errors.MessageRejectError {
	seqNum := new(message.IntValue)
	switch err := msg.Header.GetField(tag.MsgSeqNum, seqNum); {
	case err != nil:
		return errors.RequiredTagMissing(tag.MsgSeqNum)
	case seqNum.Value > s.expectedSeqNum:
		return errors.TargetTooHigh{ReceivedTarget: seqNum.Value, ExpectedTarget: s.expectedSeqNum}
	}

	return nil
}

func (s *Session) checkCompID(msg message.Message) errors.MessageRejectError {
	senderCompID := &field.SenderCompIDField{}
	targetCompID := &field.TargetCompIDField{}

	haveSender := msg.Header.Get(senderCompID)
	haveTarget := msg.Header.Get(targetCompID)

	switch {
	case haveSender != nil:
		return errors.RequiredTagMissing(tag.SenderCompID)
	case haveTarget != nil:
		return errors.RequiredTagMissing(tag.TargetCompID)
	case len(targetCompID.Value) == 0:
		return errors.TagSpecifiedWithoutAValue(tag.TargetCompID)
	case len(senderCompID.Value) == 0:
		return errors.TagSpecifiedWithoutAValue(tag.SenderCompID)
	case s.sessionID.SenderCompID != targetCompID.Value || s.sessionID.TargetCompID != senderCompID.Value:
		return errors.CompIDProblem()
	}

	return nil
}

func (s *Session) checkSendingTime(msg message.Message) errors.MessageRejectError {
	sendingTime := new(message.UTCTimestampValue)
	if err := msg.Header.GetField(tag.SendingTime, sendingTime); err != nil {
		return err
	}

	if delta := time.Since(sendingTime.Value); delta <= -1*time.Duration(120)*time.Second || delta >= time.Duration(120)*time.Second {
		return errors.SendingTimeAccuracyProblem()
	}

	return nil
}

func (s *Session) checkBeginString(msg message.Message) errors.MessageRejectError {
	beginString := new(message.StringValue)
	switch err := msg.Header.GetField(tag.BeginString, beginString); {
	case err != nil:
		return errors.RequiredTagMissing(tag.BeginString)
	case s.sessionID.BeginString != beginString.Value:
		return errors.IncorrectBeginString{}
	}

	return nil
}

func (s *Session) doReject(msg message.Message, rej errors.MessageRejectError) {
	reply := msg.ReverseRoute()

	if s.sessionID.BeginString >= fix.BeginString_FIX42 {

		if rej.IsBusinessReject() {
			reply.Header.Set(field.NewMsgType("j"))
			reply.Body.Set(field.NewBusinessRejectReason(int(rej.RejectReason())))
		} else {
			reply.Header.Set(field.NewMsgType("3"))
			switch {
			default:
				reply.Body.Set(field.NewSessionRejectReason(int(rej.RejectReason())))
			case rej.RejectReason() > errors.RejectReasonInvalidMsgType && s.sessionID.BeginString == fix.BeginString_FIX42:
				//fix42 knows up to invalid msg type
			}
		}
		reply.Body.Set(field.NewText(rej.Error()))

		var msgType field.MsgTypeField
		if err := msg.Header.Get(&msgType); err == nil {
			reply.Body.Set(field.NewRefMsgType(msgType.Value))
		}

		if refTagID := rej.RefTagID(); refTagID != nil {
			reply.Body.Set(field.NewRefTagID(int(*refTagID)))
		}
	} else {
		reply.Header.Set(field.NewMsgType("3"))

		if refTagID := rej.RefTagID(); refTagID != nil {
			reply.Body.Set(field.NewText(fmt.Sprintf("%s (%d)", rej.Error(), *refTagID)))
		} else {
			reply.Body.Set(field.NewText(rej.Error()))
		}
	}

	var seqNum field.MsgSeqNumField
	if err := msg.Header.Get(&seqNum); err == nil {
		reply.Body.Set(field.NewRefSeqNum(seqNum.Value))
	}

	s.send(reply)
	s.log.OnEventf("Message Rejected: %v", rej.Error())
}

func (s *Session) run(msgIn chan []byte) {
	defer func() {
		s.messageOut <- nil
	}()

	if s.initiateLogon {
		s.expectedSeqNum = 1
		s.seqNum = 1

		logon := message.CreateMessageBuilder()
		logon.Header.Set(field.NewMsgType("A"))
		logon.Header.Set(field.NewBeginString(s.sessionID.BeginString))
		logon.Header.Set(field.NewTargetCompID(s.sessionID.TargetCompID))
		logon.Header.Set(field.NewSenderCompID(s.sessionID.SenderCompID))
		logon.Body.Set(field.NewEncryptMethod(0))
		logon.Body.Set(field.NewHeartBtInt(s.heartBtInt))

		s.heartBeatTimeout = time.Duration(s.heartBtInt) * time.Second

		if len(s.defaultApplVerID) > 0 {
			logon.Body.Set(field.NewDefaultApplVerID(s.defaultApplVerID))
		}

		s.log.OnEvent("Sending logon request")
		s.send(logon)
	}

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
					s.log.OnEventf("Msg Parse Error: %v, %q", err.Error(), msgBytes)
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
