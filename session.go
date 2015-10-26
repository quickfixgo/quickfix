package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/config"
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/fix/enum"
	"time"
)

//The Session is the primary FIX abstraction for message communication
type Session struct {
	store MessageStore

	log       Log
	sessionID SessionID

	messageOut chan []byte
	toSend     chan Message

	sessionEvent            chan event
	application             Application
	currentState            sessionState
	stateTimer              eventTimer
	peerTimer               eventTimer
	messageStash            map[int]Message
	dataDictionary          *datadictionary.DataDictionary
	transportDataDictionary *datadictionary.DataDictionary
	appDataDictionary       *datadictionary.DataDictionary
	resetOnLogon            bool
	initiateLogon           bool
	heartBtInt              int
	heartBeatTimeout        time.Duration

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
func createSession(sessionID SessionID, storeFactory MessageStoreFactory, settings *SessionSettings, logFactory LogFactory, application Application) error {
	session := &Session{sessionID: sessionID}

	if sessionID.BeginString == enum.BeginStringFIXT11 {
		defaultApplVerID, err := settings.Setting(config.DefaultApplVerID)
		if err != nil {
			return requiredConfigurationMissing(config.DefaultApplVerID)
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

	if session.store, err = storeFactory.Create(session.sessionID); err != nil {
		return err
	}

	session.toSend = make(chan Message)
	session.sessionEvent = make(chan event)
	session.application = application
	session.stateTimer = eventTimer{Task: func() { session.sessionEvent <- needHeartbeat }}
	session.peerTimer = eventTimer{Task: func() { session.sessionEvent <- peerTimeout }}

	application.OnCreate(session.sessionID)
	sessions.newSession <- session

	return nil
}

func (s *Session) initiate() (chan []byte, error) {
	s.currentState = logonState{}
	s.messageOut = make(chan []byte)
	s.messageStash = make(map[int]Message)
	s.initiateLogon = true

	return s.messageOut, nil
}

func (s *Session) accept() (chan []byte, error) {
	s.currentState = logonState{}
	s.messageOut = make(chan []byte)
	s.messageStash = make(map[int]Message)

	return s.messageOut, nil
}

func (s *Session) onDisconnect() {
	s.application.OnLogout(s.sessionID)
	s.log.OnEvent("Disconnected")
}

func (s *Session) insertSendingTime(header FieldMap) {
	sendingTime := time.Now().UTC()

	if s.sessionID.BeginString >= enum.BeginStringFIX42 {
		header.Set(NewUTCTimestampField(tagSendingTime, sendingTime))
	} else {
		header.Set(NewUTCTimestampFieldNoMillis(tagSendingTime, sendingTime))
	}
}

func (s *Session) fillDefaultHeader(msg Message) {
	msg.Header.Set(NewStringField(tagBeginString, s.sessionID.BeginString))
	msg.Header.Set(NewStringField(tagSenderCompID, s.sessionID.SenderCompID))
	msg.Header.Set(NewStringField(tagTargetCompID, s.sessionID.TargetCompID))

	s.insertSendingTime(msg.Header)
}

func (s *Session) resend(msg Message) {
	msg.Header.Set(NewBooleanField(tagPossDupFlag, true))

	origSendingTime := new(StringValue)
	if err := msg.Header.GetField(tagSendingTime, origSendingTime); err == nil {
		msg.Header.Set(NewStringField(tagOrigSendingTime, origSendingTime.Value))
	}

	s.insertSendingTime(msg.Header)

	msg.Build()
	s.sendBytes(msg.rawMessage)
}

func (s *Session) send(msg Message) {
	s.fillDefaultHeader(msg)

	seqNum := s.store.NextSenderMsgSeqNum()
	msg.Header.Set(NewIntField(tagMsgSeqNum, seqNum))

	msgType := new(StringValue)
	if msg.Header.GetField(tagMsgType, msgType); isAdminMessageType(msgType.Value) {
		s.application.ToAdmin(msg, s.sessionID)
	} else {
		s.application.ToApp(msg, s.sessionID)
	}
	if msgBytes, err := msg.Build(); err != nil {
		panic(err)
	} else {
		s.store.SaveMessage(seqNum, msgBytes)
		s.sendBytes(msgBytes)
		s.store.IncrNextSenderMsgSeqNum()
	}
}

func (s *Session) sendBytes(msg []byte) {

	s.log.OnOutgoing(string(msg))
	s.messageOut <- msg
	s.stateTimer.Reset(time.Duration(s.heartBeatTimeout))
}

func (s *Session) doTargetTooHigh(reject targetTooHigh) {
	resend := NewMessage()
	resend.Header.Set(NewStringField(tagMsgType, "2"))
	resend.Body.Set(NewIntField(tagBeginSeqNo, reject.ExpectedTarget))

	var endSeqNum = 0
	if s.sessionID.BeginString < enum.BeginStringFIX42 {
		endSeqNum = 999999
	}
	resend.Body.Set(NewIntField(tagEndSeqNo, endSeqNum))

	s.send(resend)
}

func (s *Session) handleLogon(msg Message) error {
	//Grab default app ver id from fixt.1.1 logon
	if s.sessionID.BeginString == enum.BeginStringFIXT11 {
		targetApplVerID := new(StringValue)

		if err := msg.Body.GetField(tagDefaultApplVerID, targetApplVerID); err != nil {
			return err
		}

		s.targetDefaultApplVerID = targetApplVerID.Value
	}

	if !s.initiateLogon {
		s.log.OnEvent("Received logon request")
		if s.resetOnLogon {
			s.store.Reset()
		}

		resetSeqNumFlag := new(BooleanValue)
		if err := msg.Body.GetField(tagResetSeqNumFlag, resetSeqNumFlag); err == nil {
			if resetSeqNumFlag.Value {
				s.log.OnEvent("Logon contains ResetSeqNumFlag=Y, resetting sequence numbers to 1")
				s.store.Reset()
			}
		}

		if err := s.verifyIgnoreSeqNumTooHigh(msg); err != nil {
			return err
		}

		reply := NewMessage()
		reply.Header.Set(NewStringField(tagMsgType, "A"))
		reply.Header.Set(NewStringField(tagBeginString, s.sessionID.BeginString))
		reply.Header.Set(NewStringField(tagTargetCompID, s.sessionID.TargetCompID))
		reply.Header.Set(NewStringField(tagSenderCompID, s.sessionID.SenderCompID))
		reply.Body.Set(NewStringField(tagEncryptMethod, "0"))

		heartBtInt := new(IntValue)
		if err := msg.Body.GetField(tagHeartBtInt, heartBtInt); err == nil {
			s.heartBeatTimeout = time.Duration(heartBtInt.Value) * time.Second
			reply.Body.Set(NewIntField(tagHeartBtInt, heartBtInt.Value))
		}

		if resetSeqNumFlag.Value {
			reply.Body.Set(NewBooleanField(tagResetSeqNumFlag, resetSeqNumFlag.Value))
		}

		if len(s.defaultApplVerID) > 0 {
			reply.Body.Set(NewStringField(tagDefaultApplVerID, s.defaultApplVerID))
		}

		s.log.OnEvent("Responding to logon request")
		s.send(reply)
	}

	s.application.OnLogon(s.sessionID)

	if err := s.checkTargetTooHigh(msg); err != nil {
		switch TypedError := err.(type) {
		case targetTooHigh:
			s.doTargetTooHigh(TypedError)
		}
	}

	s.store.IncrNextTargetMsgSeqNum()
	return nil
}

func (s *Session) verify(msg Message) MessageRejectError {
	return s.verifySelect(msg, true, true)
}

func (s *Session) verifyIgnoreSeqNumTooHigh(msg Message) MessageRejectError {
	return s.verifySelect(msg, false, true)
}

func (s *Session) verifyIgnoreSeqNumTooHighOrLow(msg Message) MessageRejectError {
	return s.verifySelect(msg, false, false)
}

func (s *Session) verifySelect(msg Message, checkTooHigh bool, checkTooLow bool) MessageRejectError {
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
		if reject := validate(s.dataDictionary, msg); reject != nil {
			return reject
		}
	}

	if s.transportDataDictionary != nil {
		msgType := new(StringField)
		msg.Header.GetField(tagMsgType, msgType)
		if isAdminMessageType(msgType.Value) {
			if reject := validate(s.transportDataDictionary, msg); reject != nil {
				return reject
			}
		} else {
			if reject := validateFIXTApp(s.transportDataDictionary, s.appDataDictionary, msg); reject != nil {
				return reject
			}
		}
	}

	return s.fromCallback(msg)
}

func (s *Session) fromCallback(msg Message) MessageRejectError {
	msgType := new(StringValue)
	if msg.Header.GetField(tagMsgType, msgType); isAdminMessageType(msgType.Value) {
		return s.application.FromAdmin(msg, s.sessionID)
	}

	return s.application.FromApp(msg, s.sessionID)
}

func (s *Session) checkTargetTooLow(msg Message) MessageRejectError {
	seqNum := new(IntField)
	switch err := msg.Header.GetField(tagMsgSeqNum, seqNum); {
	case err != nil:
		return requiredTagMissing(tagMsgSeqNum)
	case seqNum.Value < s.store.NextTargetMsgSeqNum():
		return targetTooLow{ReceivedTarget: seqNum.Value, ExpectedTarget: s.store.NextTargetMsgSeqNum()}
	}

	return nil
}

func (s *Session) checkTargetTooHigh(msg Message) MessageRejectError {
	seqNum := new(IntValue)
	switch err := msg.Header.GetField(tagMsgSeqNum, seqNum); {
	case err != nil:
		return requiredTagMissing(tagMsgSeqNum)
	case seqNum.Value > s.store.NextTargetMsgSeqNum():
		return targetTooHigh{ReceivedTarget: seqNum.Value, ExpectedTarget: s.store.NextTargetMsgSeqNum()}
	}

	return nil
}

func (s *Session) checkCompID(msg Message) MessageRejectError {
	senderCompID := new(StringField)
	targetCompID := new(StringField)

	haveSender := msg.Header.GetField(tagSenderCompID, senderCompID)
	haveTarget := msg.Header.GetField(tagTargetCompID, targetCompID)

	switch {
	case haveSender != nil:
		return requiredTagMissing(tagSenderCompID)
	case haveTarget != nil:
		return requiredTagMissing(tagTargetCompID)
	case len(targetCompID.Value) == 0:
		return tagSpecifiedWithoutAValue(tagTargetCompID)
	case len(senderCompID.Value) == 0:
		return tagSpecifiedWithoutAValue(tagSenderCompID)
	case s.sessionID.SenderCompID != targetCompID.Value || s.sessionID.TargetCompID != senderCompID.Value:
		return compIDProblem()
	}

	return nil
}

func (s *Session) checkSendingTime(msg Message) MessageRejectError {
	if ok := msg.Header.Has(tagSendingTime); !ok {
		return requiredTagMissing(tagSendingTime)
	}

	sendingTime := new(UTCTimestampValue)
	if err := msg.Header.GetField(tagSendingTime, sendingTime); err != nil {
		return err
	}

	if delta := time.Since(sendingTime.Value); delta <= -1*time.Duration(120)*time.Second || delta >= time.Duration(120)*time.Second {
		return sendingTimeAccuracyProblem()
	}

	return nil
}

func (s *Session) checkBeginString(msg Message) MessageRejectError {
	beginString := new(StringValue)
	switch err := msg.Header.GetField(tagBeginString, beginString); {
	case err != nil:
		return requiredTagMissing(tagBeginString)
	case s.sessionID.BeginString != beginString.Value:
		return incorrectBeginString{}
	}

	return nil
}

func (s *Session) doReject(msg Message, rej MessageRejectError) {
	reply := msg.reverseRoute()

	if s.sessionID.BeginString >= enum.BeginStringFIX42 {

		if rej.IsBusinessReject() {
			reply.Header.Set(NewStringField(tagMsgType, "j"))
			reply.Body.Set(NewIntField(tagBusinessRejectReason, int(rej.RejectReason())))
		} else {
			reply.Header.Set(NewStringField(tagMsgType, "3"))
			switch {
			default:
				reply.Body.Set(NewIntField(tagSessionRejectReason, int(rej.RejectReason())))
			case rej.RejectReason() > rejectReasonInvalidMsgType && s.sessionID.BeginString == enum.BeginStringFIX42:
				//fix42 knows up to invalid msg type
			}
		}
		reply.Body.Set(NewStringField(tagText, rej.Error()))

		msgType := new(StringValue)
		if err := msg.Header.GetField(tagMsgType, msgType); err == nil {
			reply.Body.Set(NewStringField(tagRefMsgType, msgType.Value))
		}

		if refTagID := rej.RefTagID(); refTagID != nil {
			reply.Body.Set(NewIntField(tagRefTagID, int(*refTagID)))
		}
	} else {
		reply.Header.Set(NewStringField(tagMsgType, "3"))

		if refTagID := rej.RefTagID(); refTagID != nil {
			reply.Body.Set(NewStringField(tagText, fmt.Sprintf("%s (%d)", rej.Error(), *refTagID)))
		} else {
			reply.Body.Set(NewStringField(tagText, rej.Error()))
		}
	}

	seqNum := new(IntValue)
	if err := msg.Header.GetField(tagMsgSeqNum, seqNum); err == nil {
		reply.Body.Set(NewIntField(tagRefSeqNum, seqNum.Value))
	}

	s.send(reply)
	s.log.OnEventf("Message Rejected: %v", rej.Error())
}

type fixIn struct {
	bytes       []byte
	receiveTime time.Time
}

func (s *Session) run(msgIn chan fixIn, quit chan bool) {
	defer func() {
		close(quit)
		s.messageOut <- nil
		s.onDisconnect()
	}()

	if s.initiateLogon {

		if s.resetOnLogon {
			s.store.Reset()
		}

		logon := NewMessage()
		logon.Header.Set(NewStringField(tagMsgType, "A"))
		logon.Header.Set(NewStringField(tagBeginString, s.sessionID.BeginString))
		logon.Header.Set(NewStringField(tagTargetCompID, s.sessionID.TargetCompID))
		logon.Header.Set(NewStringField(tagSenderCompID, s.sessionID.SenderCompID))
		logon.Body.Set(NewStringField(tagEncryptMethod, "0"))
		logon.Body.Set(NewIntField(tagHeartBtInt, s.heartBtInt))

		s.heartBeatTimeout = time.Duration(s.heartBtInt) * time.Second

		if len(s.defaultApplVerID) > 0 {
			logon.Body.Set(NewStringField(tagDefaultApplVerID, s.defaultApplVerID))
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
		case fixIn, ok := <-msgIn:
			if ok {
				s.log.OnIncoming(string(fixIn.bytes))
				if msg, err := parseMessage(fixIn.bytes); err != nil {
					s.log.OnEventf("Msg Parse Error: %v, %q", err.Error(), fixIn.bytes)
				} else {
					msg.ReceiveTime = fixIn.receiveTime
					s.currentState = s.currentState.FixMsgIn(s, msg)
				}
			} else {
				return
			}
			s.peerTimer.Reset(time.Duration(int64(1.2 * float64(s.heartBeatTimeout))))

		case msg := <-s.toSend:
			s.send(msg)

		case <-quit:
			return

		case evt := <-s.sessionEvent:
			s.currentState = s.currentState.Timeout(s, evt)
		}
	}
}
