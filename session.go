package quickfix

import (
	"fmt"
	"sync"
	"time"

	"github.com/quickfixgo/quickfix/config"
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/enum"
)

//The Session is the primary FIX abstraction for message communication
type session struct {
	store MessageStore

	log       Log
	sessionID SessionID

	messageOut chan []byte

	//application messages are queued up to be sent during the run loop.
	toSend []Message
	//mutex for access to toSend
	sendMutex sync.Mutex

	sessionEvent chan event
	messageEvent chan bool
	application  Application
	validator
	sessionState
	stateTimer       eventTimer
	peerTimer        eventTimer
	messageStash     map[int]Message
	resetOnLogon     bool
	initiateLogon    bool
	heartBtInt       int
	heartBeatTimeout time.Duration

	//required on logon for FIX.T.1 messages
	defaultApplVerID       string
	targetDefaultApplVerID string
}

//TargetDefaultApplicationVersionID returns the default application version ID for messages received by this version.
//Applicable for For FIX.T.1 sessions.
func (s *session) TargetDefaultApplicationVersionID() string {
	return s.targetDefaultApplVerID
}

//Creates Session, associates with internal session registry
func createSession(sessionID SessionID, storeFactory MessageStoreFactory, settings *SessionSettings, logFactory LogFactory, application Application) error {
	session := &session{sessionID: sessionID}

	var err error
	var validatorSettings = defaultValidatorSettings
	if settings.HasSetting(config.ValidateFieldsOutOfOrder) {
		if validatorSettings.CheckFieldsOutOfOrder, err = settings.BoolSetting(config.ValidateFieldsOutOfOrder); err != nil {
			return err
		}
	}

	if sessionID.IsFIXT() {
		if defaultApplVerID, err := settings.Setting(config.DefaultApplVerID); err == nil {
			session.defaultApplVerID = defaultApplVerID
		} else {
			return requiredConfigurationMissing(config.DefaultApplVerID)
		}

		//If the transport or app data dictionary setting is set, the other also needs to be set.
		hasTransportDataDictionary := settings.HasSetting(config.TransportDataDictionary)
		hasAppDataDictionary := settings.HasSetting(config.AppDataDictionary)
		if hasTransportDataDictionary && hasAppDataDictionary {
			transportDataDictionaryPath, _ := settings.Setting(config.TransportDataDictionary)
			transportDataDictionary, err := datadictionary.Parse(transportDataDictionaryPath)
			if err != nil {
				return err
			}

			appDataDictionaryPath, _ := settings.Setting(config.AppDataDictionary)
			appDataDictionary, err := datadictionary.Parse(appDataDictionaryPath)
			if err != nil {
				return err
			}

			session.validator = &fixtValidator{transportDataDictionary, appDataDictionary, validatorSettings}
		} else if hasTransportDataDictionary {
			return requiredConfigurationMissing(config.AppDataDictionary)
		} else if hasAppDataDictionary {
			return requiredConfigurationMissing(config.TransportDataDictionary)
		}
	} else {
		var dataDictionary *datadictionary.DataDictionary
		if dataDictionaryPath, err := settings.Setting(config.DataDictionary); err == nil {
			if dataDictionary, err = datadictionary.Parse(dataDictionaryPath); err != nil {
				return err
			}

			session.validator = &fixValidator{dataDictionary, validatorSettings}
		}
	}

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

	session.sessionEvent = make(chan event)
	session.messageEvent = make(chan bool)
	session.application = application
	session.stateTimer = eventTimer{Task: func() { session.sessionEvent <- needHeartbeat }}
	session.peerTimer = eventTimer{Task: func() { session.sessionEvent <- peerTimeout }}
	application.OnCreate(session.sessionID)
	sessions.newSession <- session

	return nil
}

//kicks off session as an initiator
func (s *session) initiate(msgIn chan fixIn, msgOut chan []byte, quit chan bool) {
	s.sessionState = logonState{}
	s.messageStash = make(map[int]Message)
	s.initiateLogon = true

	s.run(msgIn, msgOut, quit)
}

//kicks off session as an acceptor
func (s *session) accept(msgIn chan fixIn, msgOut chan []byte, quit chan bool) {
	s.sessionState = logonState{}
	s.messageStash = make(map[int]Message)

	s.run(msgIn, msgOut, quit)
}

func (s *session) onDisconnect() {
	s.application.OnLogout(s.sessionID)
	s.log.OnEvent("Disconnected")
}

func (s *session) insertSendingTime(header Header) {
	sendingTime := time.Now().UTC()

	if s.sessionID.BeginString >= enum.BeginStringFIX42 {
		header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime})
	} else {
		header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime, NoMillis: true})
	}
}

func (s *session) fillDefaultHeader(msg Message) {
	msg.Header.SetField(tagBeginString, FIXString(s.sessionID.BeginString))
	msg.Header.SetField(tagSenderCompID, FIXString(s.sessionID.SenderCompID))
	msg.Header.SetField(tagTargetCompID, FIXString(s.sessionID.TargetCompID))

	s.insertSendingTime(msg.Header)
}

func (s *session) resend(msg Message) {
	msg.Header.SetField(tagPossDupFlag, FIXBoolean(true))

	var origSendingTime FIXString
	if err := msg.Header.GetField(tagSendingTime, &origSendingTime); err == nil {
		msg.Header.SetField(tagOrigSendingTime, origSendingTime)
	}

	s.insertSendingTime(msg.Header)

	msg.Build()
	s.sendBytes(msg.rawMessage)
}

//queueForSend will queue up a message to be sent by the session during the next iteration of the run loop
func (s *session) queueForSend(msg Message) {
	s.sendMutex.Lock()
	defer s.sendMutex.Unlock()
	s.toSend = append(s.toSend, msg)
	select {
	case s.messageEvent <- true:
	default:
	}
}

//sends queued messages if session is logged on
func (s *session) sendQueued() {
	if !s.IsLoggedOn() {
		return
	}

	s.sendMutex.Lock()
	defer s.sendMutex.Unlock()

	for _, msg := range s.toSend {
		s.send(msg)
	}

	s.toSend = s.toSend[:0]
}

//send should NOT be called outside of the run loop
func (s *session) send(msg Message) {
	s.fillDefaultHeader(msg)

	seqNum := s.store.NextSenderMsgSeqNum()
	msg.Header.SetField(tagMsgSeqNum, FIXInt(seqNum))

	var msgType FIXString
	if msg.Header.GetField(tagMsgType, &msgType); isAdminMessageType(string(msgType)) {
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

func (s *session) sendBytes(msg []byte) {
	s.log.OnOutgoing(string(msg))
	s.messageOut <- msg
	s.stateTimer.Reset(time.Duration(s.heartBeatTimeout))
}

func (s *session) doTargetTooHigh(reject targetTooHigh) {
	resend := NewMessage()
	resend.Header.SetField(tagMsgType, FIXString("2"))
	resend.Body.SetField(tagBeginSeqNo, FIXInt(reject.ExpectedTarget))

	var endSeqNum = 0
	if s.sessionID.BeginString < enum.BeginStringFIX42 {
		endSeqNum = 999999
	}
	resend.Body.SetField(tagEndSeqNo, FIXInt(endSeqNum))

	s.send(resend)
}

func (s *session) handleLogon(msg Message) error {
	//Grab default app ver id from fixt.1.1 logon
	if s.sessionID.BeginString == enum.BeginStringFIXT11 {
		var targetApplVerID FIXString

		if err := msg.Body.GetField(tagDefaultApplVerID, &targetApplVerID); err != nil {
			return err
		}

		s.targetDefaultApplVerID = string(targetApplVerID)
	}

	if !s.initiateLogon {
		s.log.OnEvent("Received logon request")
		if s.resetOnLogon {
			s.store.Reset()
		}

		var resetSeqNumFlag FIXBoolean
		if err := msg.Body.GetField(tagResetSeqNumFlag, &resetSeqNumFlag); err == nil {
			if resetSeqNumFlag {
				s.log.OnEvent("Logon contains ResetSeqNumFlag=Y, resetting sequence numbers to 1")
				s.store.Reset()
			}
		}

		if err := s.verifyIgnoreSeqNumTooHigh(msg); err != nil {
			return err
		}

		reply := NewMessage()
		reply.Header.SetField(tagMsgType, FIXString("A"))
		reply.Header.SetField(tagBeginString, FIXString(s.sessionID.BeginString))
		reply.Header.SetField(tagTargetCompID, FIXString(s.sessionID.TargetCompID))
		reply.Header.SetField(tagSenderCompID, FIXString(s.sessionID.SenderCompID))
		reply.Body.SetField(tagEncryptMethod, FIXString("0"))

		var heartBtInt FIXInt
		if err := msg.Body.GetField(tagHeartBtInt, &heartBtInt); err == nil {
			s.heartBeatTimeout = time.Duration(heartBtInt) * time.Second
			reply.Body.SetField(tagHeartBtInt, heartBtInt)
		}

		if resetSeqNumFlag {
			reply.Body.SetField(tagResetSeqNumFlag, resetSeqNumFlag)
		}

		if len(s.defaultApplVerID) > 0 {
			reply.Body.SetField(tagDefaultApplVerID, FIXString(s.defaultApplVerID))
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

func (s *session) verify(msg Message) MessageRejectError {
	return s.verifySelect(msg, true, true)
}

func (s *session) verifyIgnoreSeqNumTooHigh(msg Message) MessageRejectError {
	return s.verifySelect(msg, false, true)
}

func (s *session) verifyIgnoreSeqNumTooHighOrLow(msg Message) MessageRejectError {
	return s.verifySelect(msg, false, false)
}

func (s *session) verifySelect(msg Message, checkTooHigh bool, checkTooLow bool) MessageRejectError {
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

	if s.validator != nil {
		if reject := s.validator.Validate(msg); reject != nil {
			return reject
		}
	}

	return s.fromCallback(msg)
}

func (s *session) fromCallback(msg Message) MessageRejectError {
	var msgType FIXString
	if msg.Header.GetField(tagMsgType, &msgType); isAdminMessageType(string(msgType)) {
		return s.application.FromAdmin(msg, s.sessionID)
	}

	return s.application.FromApp(msg, s.sessionID)
}

func (s *session) checkTargetTooLow(msg Message) MessageRejectError {
	var seqNum FIXInt
	switch err := msg.Header.GetField(tagMsgSeqNum, &seqNum); {
	case err != nil:
		return RequiredTagMissing(tagMsgSeqNum)
	case int(seqNum) < s.store.NextTargetMsgSeqNum():
		return targetTooLow{ReceivedTarget: int(seqNum), ExpectedTarget: s.store.NextTargetMsgSeqNum()}
	}

	return nil
}

func (s *session) checkTargetTooHigh(msg Message) MessageRejectError {
	var seqNum FIXInt
	switch err := msg.Header.GetField(tagMsgSeqNum, &seqNum); {
	case err != nil:
		return RequiredTagMissing(tagMsgSeqNum)
	case int(seqNum) > s.store.NextTargetMsgSeqNum():
		return targetTooHigh{ReceivedTarget: int(seqNum), ExpectedTarget: s.store.NextTargetMsgSeqNum()}
	}

	return nil
}

func (s *session) checkCompID(msg Message) MessageRejectError {
	var senderCompID FIXString
	var targetCompID FIXString

	haveSender := msg.Header.GetField(tagSenderCompID, &senderCompID)
	haveTarget := msg.Header.GetField(tagTargetCompID, &targetCompID)

	switch {
	case haveSender != nil:
		return RequiredTagMissing(tagSenderCompID)
	case haveTarget != nil:
		return RequiredTagMissing(tagTargetCompID)
	case len(targetCompID) == 0:
		return TagSpecifiedWithoutAValue(tagTargetCompID)
	case len(senderCompID) == 0:
		return TagSpecifiedWithoutAValue(tagSenderCompID)
	case s.sessionID.SenderCompID != string(targetCompID) || s.sessionID.TargetCompID != string(senderCompID):
		return compIDProblem()
	}

	return nil
}

func (s *session) checkSendingTime(msg Message) MessageRejectError {
	if ok := msg.Header.Has(tagSendingTime); !ok {
		return RequiredTagMissing(tagSendingTime)
	}

	sendingTime := new(FIXUTCTimestamp)
	if err := msg.Header.GetField(tagSendingTime, sendingTime); err != nil {
		return err
	}

	if delta := time.Since(sendingTime.Time); delta <= -1*time.Duration(120)*time.Second || delta >= time.Duration(120)*time.Second {
		return sendingTimeAccuracyProblem()
	}

	return nil
}

func (s *session) checkBeginString(msg Message) MessageRejectError {
	var beginString FIXString
	switch err := msg.Header.GetField(tagBeginString, &beginString); {
	case err != nil:
		return RequiredTagMissing(tagBeginString)
	case s.sessionID.BeginString != string(beginString):
		return incorrectBeginString{}
	}

	return nil
}

func (s *session) doReject(msg Message, rej MessageRejectError) {
	reply := msg.reverseRoute()

	if s.sessionID.BeginString >= enum.BeginStringFIX42 {

		if rej.IsBusinessReject() {
			reply.Header.SetField(tagMsgType, FIXString("j"))
			reply.Body.SetField(tagBusinessRejectReason, FIXInt(rej.RejectReason()))
		} else {
			reply.Header.SetField(tagMsgType, FIXString("3"))
			switch {
			default:
				reply.Body.SetField(tagSessionRejectReason, FIXInt(rej.RejectReason()))
			case rej.RejectReason() > rejectReasonInvalidMsgType && s.sessionID.BeginString == enum.BeginStringFIX42:
				//fix42 knows up to invalid msg type
			}

			if refTagID := rej.RefTagID(); refTagID != nil {
				reply.Body.SetField(tagRefTagID, FIXInt(*refTagID))
			}
		}
		reply.Body.SetField(tagText, FIXString(rej.Error()))

		var msgType FIXString
		if err := msg.Header.GetField(tagMsgType, &msgType); err == nil {
			reply.Body.SetField(tagRefMsgType, msgType)
		}
	} else {
		reply.Header.SetField(tagMsgType, FIXString("3"))

		if refTagID := rej.RefTagID(); refTagID != nil {
			reply.Body.SetField(tagText, FIXString(fmt.Sprintf("%s (%d)", rej.Error(), *refTagID)))
		} else {
			reply.Body.SetField(tagText, FIXString(rej.Error()))
		}
	}

	seqNum := new(FIXInt)
	if err := msg.Header.GetField(tagMsgSeqNum, seqNum); err == nil {
		reply.Body.SetField(tagRefSeqNum, seqNum)
	}

	s.send(reply)
	s.log.OnEventf("Message Rejected: %v", rej.Error())
}

type fixIn struct {
	bytes       []byte
	receiveTime time.Time
}

func (s *session) run(msgIn chan fixIn, msgOut chan []byte, quit chan bool) {
	s.messageOut = msgOut

	defer func() {
		close(s.messageOut)
		s.onDisconnect()
	}()

	if s.initiateLogon {
		if s.resetOnLogon {
			s.store.Reset()
		}

		logon := NewMessage()
		logon.Header.SetField(tagMsgType, FIXString("A"))
		logon.Header.SetField(tagBeginString, FIXString(s.sessionID.BeginString))
		logon.Header.SetField(tagTargetCompID, FIXString(s.sessionID.TargetCompID))
		logon.Header.SetField(tagSenderCompID, FIXString(s.sessionID.SenderCompID))
		logon.Body.SetField(tagEncryptMethod, FIXString("0"))
		logon.Body.SetField(tagHeartBtInt, FIXInt(s.heartBtInt))

		s.heartBeatTimeout = time.Duration(s.heartBtInt) * time.Second

		if len(s.defaultApplVerID) > 0 {
			logon.Body.SetField(tagDefaultApplVerID, FIXString(s.defaultApplVerID))
		}

		s.log.OnEvent("Sending logon request")
		s.send(logon)
	}

	for {

		switch s.sessionState.(type) {
		case latentState:
			return
		}

		s.sendQueued()

		select {
		case fixIn, ok := <-msgIn:
			if ok {
				s.log.OnIncoming(string(fixIn.bytes))
				if msg, err := ParseMessage(fixIn.bytes); err != nil {
					s.log.OnEventf("Msg Parse Error: %v, %q", err.Error(), fixIn.bytes)
				} else {
					msg.ReceiveTime = fixIn.receiveTime
					s.sessionState = s.FixMsgIn(s, msg)
				}
			} else {
				return
			}
			s.peerTimer.Reset(time.Duration(int64(1.2 * float64(s.heartBeatTimeout))))

		case <-quit:
			return

		case evt := <-s.sessionEvent:
			s.sessionState = s.Timeout(s, evt)
		case <-s.messageEvent:
		}
	}
}
