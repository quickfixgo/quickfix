package quickfix

import (
	"fmt"
	"sync"
	"time"

	"github.com/quickfixgo/quickfix/config"
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/internal"
)

//The Session is the primary FIX abstraction for message communication
type session struct {
	store MessageStore

	log       Log
	sessionID SessionID

	messageOut chan<- []byte
	messageIn  <-chan fixIn

	//application messages are queued up for send here
	toSend []Message

	//mutex for access to toSend
	sendMutex sync.Mutex

	sessionEvent chan internal.Event
	messageEvent chan bool
	application  Application
	validator
	stateMachine
	stateTimer       internal.EventTimer
	peerTimer        internal.EventTimer
	messageStash     map[int]Message
	resetOnLogon     bool
	resetOnLogout    bool
	initiateLogon    bool
	heartBtInt       int
	heartBeatTimeout time.Duration

	//required on logon for FIX.T.1 messages
	defaultApplVerID       string
	targetDefaultApplVerID string

	sessionTime *internal.TimeRange
	quit        <-chan bool
	admin       chan interface{}
}

func (s *session) logError(err error) {
	s.log.OnEvent(err.Error())
}

//TargetDefaultApplicationVersionID returns the default application version ID for messages received by this version.
//Applicable for For FIX.T.1 sessions.
func (s *session) TargetDefaultApplicationVersionID() string {
	return s.targetDefaultApplVerID
}

func newSession(sessionID SessionID, storeFactory MessageStoreFactory, settings *SessionSettings, logFactory LogFactory, application Application) (*session, error) {
	session := &session{sessionID: sessionID}

	var err error
	var validatorSettings = defaultValidatorSettings
	if settings.HasSetting(config.ValidateFieldsOutOfOrder) {
		if validatorSettings.CheckFieldsOutOfOrder, err = settings.BoolSetting(config.ValidateFieldsOutOfOrder); err != nil {
			return session, err
		}
	}

	if sessionID.IsFIXT() {
		if defaultApplVerID, err := settings.Setting(config.DefaultApplVerID); err == nil {
			session.defaultApplVerID = defaultApplVerID
		} else {
			return session, requiredConfigurationMissing(config.DefaultApplVerID)
		}

		//If the transport or app data dictionary setting is set, the other also needs to be set.
		hasTransportDataDictionary := settings.HasSetting(config.TransportDataDictionary)
		hasAppDataDictionary := settings.HasSetting(config.AppDataDictionary)
		if hasTransportDataDictionary && hasAppDataDictionary {
			transportDataDictionaryPath, _ := settings.Setting(config.TransportDataDictionary)
			transportDataDictionary, err := datadictionary.Parse(transportDataDictionaryPath)
			if err != nil {
				return session, err
			}

			appDataDictionaryPath, _ := settings.Setting(config.AppDataDictionary)
			appDataDictionary, err := datadictionary.Parse(appDataDictionaryPath)
			if err != nil {
				return session, err
			}

			session.validator = &fixtValidator{transportDataDictionary, appDataDictionary, validatorSettings}
		} else if hasTransportDataDictionary {
			return session, requiredConfigurationMissing(config.AppDataDictionary)
		} else if hasAppDataDictionary {
			return session, requiredConfigurationMissing(config.TransportDataDictionary)
		}
	} else {
		var dataDictionary *datadictionary.DataDictionary
		if dataDictionaryPath, err := settings.Setting(config.DataDictionary); err == nil {
			if dataDictionary, err = datadictionary.Parse(dataDictionaryPath); err != nil {
				return session, err
			}

			session.validator = &fixValidator{dataDictionary, validatorSettings}
		}
	}

	if settings.HasSetting(config.ResetOnLogon) {
		if session.resetOnLogon, err = settings.BoolSetting(config.ResetOnLogon); err != nil {
			return session, err
		}
	}

	if settings.HasSetting(config.ResetOnLogout) {
		if session.resetOnLogout, err = settings.BoolSetting(config.ResetOnLogout); err != nil {
			return session, err
		}
	}

	if settings.HasSetting(config.HeartBtInt) {
		if heartBtInt, err := settings.IntSetting(config.HeartBtInt); err != nil {
			return session, err
		} else if heartBtInt <= 0 {
			return session, fmt.Errorf("Heartbeat must be greater than zero")
		} else {
			session.heartBtInt = heartBtInt
		}
	}

	switch {
	case !settings.HasSetting(config.StartTime) && !settings.HasSetting(config.EndTime):
	//no session times
	case settings.HasSetting(config.StartTime) && settings.HasSetting(config.EndTime):
		var startTimeStr, endTimeStr string
		if startTimeStr, err = settings.Setting(config.StartTime); err != nil {
			return session, err
		}

		if endTimeStr, err = settings.Setting(config.EndTime); err != nil {
			return session, err
		}

		session.sessionTime = new(internal.TimeRange)
		if session.sessionTime.StartTime, err = internal.ParseTimeOfDay(startTimeStr); err != nil {
			return session, err
		}
		if session.sessionTime.EndTime, err = internal.ParseTimeOfDay(endTimeStr); err != nil {
			return session, err
		}
	case settings.HasSetting(config.StartTime):
		return session, requiredConfigurationMissing(config.EndTime)
	case settings.HasSetting(config.EndTime):
		return session, requiredConfigurationMissing(config.StartTime)
	}

	if session.log, err = logFactory.CreateSessionLog(session.sessionID); err != nil {
		return session, err
	}

	if session.store, err = storeFactory.Create(session.sessionID); err != nil {
		return session, err
	}

	session.sessionEvent = make(chan internal.Event)
	session.messageEvent = make(chan bool, 1)
	session.admin = make(chan interface{})
	session.application = application
	session.stateTimer = internal.EventTimer{Task: func() { session.sessionEvent <- internal.NeedHeartbeat }}
	session.peerTimer = internal.EventTimer{Task: func() { session.sessionEvent <- internal.PeerTimeout }}
	return session, nil
}

//Creates Session, associates with internal session registry
func createSession(
	sessionID SessionID, storeFactory MessageStoreFactory, settings *SessionSettings,
	logFactory LogFactory, application Application, quit <-chan bool,
) error {
	session, err := newSession(sessionID, storeFactory, settings, logFactory, application)
	if err != nil {
		return err
	}

	session.quit = quit
	application.OnCreate(session.sessionID)
	session.log.OnEvent("Created session")
	sessions.newSession <- session
	go session.run()

	return nil
}

type connect struct {
	messageOut    chan<- []byte
	messageIn     <-chan fixIn
	initiateLogon bool
}

//kicks off session as an initiator
func (s *session) initiate(msgIn <-chan fixIn, msgOut chan<- []byte) {
	s.admin <- connect{
		messageOut:    msgOut,
		messageIn:     msgIn,
		initiateLogon: true,
	}
}

//kicks off session as an acceptor
func (s *session) accept(msgIn chan fixIn, msgOut chan []byte) {
	s.admin <- connect{
		messageOut: msgOut,
		messageIn:  msgIn,
	}
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

func (s *session) sendLogoutAndReset() {
	if err := s.sendLogout(""); err != nil {
		s.logError(err)
	}

	if err := s.dropAndReset(); err != nil {
		s.logError(err)
	}
}

func (s *session) sendLogout(reason string) error {
	logout := NewMessage()
	logout.Header.SetField(tagMsgType, FIXString("5"))
	logout.Header.SetField(tagBeginString, FIXString(s.sessionID.BeginString))
	logout.Header.SetField(tagTargetCompID, FIXString(s.sessionID.TargetCompID))
	logout.Header.SetField(tagSenderCompID, FIXString(s.sessionID.SenderCompID))
	if reason != "" {
		logout.Body.SetField(tagText, FIXString(reason))
	}
	return s.send(logout)
}

func (s *session) resend(msg Message) error {
	msg.Header.SetField(tagPossDupFlag, FIXBoolean(true))

	var origSendingTime FIXString
	if err := msg.Header.GetField(tagSendingTime, &origSendingTime); err == nil {
		msg.Header.SetField(tagOrigSendingTime, origSendingTime)
	}

	s.insertSendingTime(msg.Header)

	if _, err := msg.Build(); err != nil {
		return err
	}
	s.sendBytes(msg.rawMessage)

	return nil
}

//queueForSend will validate, persist, and queue the message for send
func (s *session) queueForSend(msg Message) error {
	s.sendMutex.Lock()
	defer s.sendMutex.Unlock()

	if err := s.prepMessageForSend(&msg); err != nil {
		return err
	}

	s.toSend = append(s.toSend, msg)

	select {
	case s.messageEvent <- true:
	default:
	}

	return nil
}

//send will validate, persist, queue the message. If the session is logged on, send all messages in the queue
func (s *session) send(msg Message) error {
	if !s.IsLoggedOn() {
		return s.queueForSend(msg)
	}

	s.sendMutex.Lock()
	defer s.sendMutex.Unlock()

	if err := s.prepMessageForSend(&msg); err != nil {
		return err
	}

	s.toSend = append(s.toSend, msg)
	s.sendQueued()

	return nil
}

//dropAndReset will drop the send queue and reset the message store
func (s *session) dropAndReset() error {
	s.sendMutex.Lock()
	defer s.sendMutex.Unlock()

	s.dropQueued()
	return s.store.Reset()
}

//dropAndSend will optionally reset the store, validate and persist the message, then drops the send queue and sends the message.
func (s *session) dropAndSend(msg Message, resetStore bool) error {

	s.sendMutex.Lock()
	defer s.sendMutex.Unlock()

	if resetStore {
		if err := s.store.Reset(); err != nil {
			return err
		}
	}

	if err := s.prepMessageForSend(&msg); err != nil {
		return err
	}

	s.dropQueued()
	s.toSend = append(s.toSend, msg)
	s.sendQueued()

	return nil
}

func (s *session) prepMessageForSend(msg *Message) error {
	s.fillDefaultHeader(*msg)
	seqNum := s.store.NextSenderMsgSeqNum()
	msg.Header.SetField(tagMsgSeqNum, FIXInt(seqNum))

	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err != nil {
		return err
	}

	if isAdminMessageType(string(msgType)) {
		s.application.ToAdmin(*msg, s.sessionID)
	} else {
		if err := s.application.ToApp(*msg, s.sessionID); err != nil {
			return err
		}
	}

	msgBytes, err := msg.Build()
	if err == nil {
		err = s.persist(seqNum, msgBytes)
	}

	return err
}

func (s *session) persist(seqNum int, msgBytes []byte) error {
	if err := s.store.SaveMessage(seqNum, msgBytes); err != nil {
		return err
	}

	return s.store.IncrNextSenderMsgSeqNum()
}

func (s *session) sendQueued() {
	for _, msg := range s.toSend {
		s.sendBytes(msg.rawMessage)
	}

	s.dropQueued()
}

func (s *session) dropQueued() {
	s.toSend = s.toSend[:0]
}

func (s *session) sendOrDropAppMessages() {
	s.sendMutex.Lock()
	defer s.sendMutex.Unlock()

	if s.IsLoggedOn() {
		s.sendQueued()
	} else {
		s.dropQueued()
	}
}

func (s *session) sendBytes(msg []byte) {
	s.log.OnOutgoing(string(msg))
	s.messageOut <- msg
	s.stateTimer.Reset(time.Duration(s.heartBeatTimeout))
}

func (s *session) doTargetTooHigh(reject targetTooHigh) error {
	s.log.OnEventf("MsgSeqNum too high, expecting %v but received %v", reject.ExpectedTarget, reject.ReceivedTarget)

	resend := NewMessage()
	resend.Header.SetField(tagMsgType, FIXString("2"))
	resend.Body.SetField(tagBeginSeqNo, FIXInt(reject.ExpectedTarget))

	var endSeqNum = 0
	if s.sessionID.BeginString < enum.BeginStringFIX42 {
		endSeqNum = 999999
	}
	resend.Body.SetField(tagEndSeqNo, FIXInt(endSeqNum))

	if err := s.send(resend); err != nil {
		return err
	}

	s.log.OnEventf("Sent ResendRequest FROM: %v TO: %v", reject.ExpectedTarget, endSeqNum)

	return nil
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

	resetStore := false
	if s.initiateLogon {
		s.log.OnEvent("Received logon response")
	} else {
		s.log.OnEvent("Received logon request")
		resetStore = s.resetOnLogon
	}

	var resetSeqNumFlag FIXBoolean
	if err := msg.Body.GetField(tagResetSeqNumFlag, &resetSeqNumFlag); err == nil {
		if resetSeqNumFlag {
			s.log.OnEvent("Logon contains ResetSeqNumFlag=Y, resetting sequence numbers to 1")
			resetStore = true
		}
	}

	if resetStore {
		if err := s.store.Reset(); err != nil {
			return err
		}
	}

	if err := s.verifyIgnoreSeqNumTooHigh(msg); err != nil {
		return err
	}

	if !s.initiateLogon {
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
		if err := s.dropAndSend(reply, resetStore); err != nil {
			return err
		}
	}

	s.peerTimer.Reset(time.Duration(int64(1.2 * float64(s.heartBeatTimeout))))
	s.application.OnLogon(s.sessionID)

	if err := s.checkTargetTooHigh(msg); err != nil {
		switch TypedError := err.(type) {
		case targetTooHigh:
			return s.doTargetTooHigh(TypedError)
		}
	}

	return s.store.IncrNextTargetMsgSeqNum()
}

func (s *session) initiateLogout(reason string) (err error) {
	if err = s.sendLogout(reason); err != nil {
		s.logError(err)
		return
	}
	s.log.OnEvent("Inititated logout request")
	time.AfterFunc(time.Duration(2)*time.Second, func() { s.sessionEvent <- internal.LogoutTimeout })

	return
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
	if err := msg.Header.GetField(tagMsgType, &msgType); err != nil {
		return err
	}

	if isAdminMessageType(string(msgType)) {
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

func (s *session) doReject(msg Message, rej MessageRejectError) error {
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

	s.log.OnEventf("Message Rejected: %v", rej.Error())
	return s.send(reply)
}

type fixIn struct {
	bytes       []byte
	receiveTime time.Time
}

func (s *session) onDisconnect() {
	if s.messageOut != nil {
		close(s.messageOut)
		s.messageOut = nil
	}

	s.messageIn = nil
}

func (s *session) onIncoming(m fixIn) {
	s.log.OnIncoming(string(m.bytes))
	if msg, err := ParseMessage(m.bytes); err != nil {
		s.log.OnEventf("Msg Parse Error: %v, %q", err.Error(), m.bytes)
	} else {
		msg.ReceiveTime = m.receiveTime
		s.FixMsgIn(s, msg)
	}
	s.peerTimer.Reset(time.Duration(int64(1.2 * float64(s.heartBeatTimeout))))
}

func (s *session) onAdmin(msg interface{}) {
	switch msg := msg.(type) {
	case connect:
		s.messageStash = make(map[int]Message)
		s.messageIn = msg.messageIn
		s.messageOut = msg.messageOut
		s.initiateLogon = msg.initiateLogon

		if s.initiateLogon {
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
			if err := s.dropAndSend(logon, s.resetOnLogon); err != nil {
				s.logError(err)
				return
			}
		}

		s.State = logonState{}
	}
}

func (s *session) run() {
	s.State = latentState{}
	s.CheckSessionTime(s, time.Now())

	ticker := time.NewTicker(time.Second)
	defer func() {
		s.stateTimer.Stop()
		s.peerTimer.Stop()
		ticker.Stop()
	}()

	for {
		select {

		case msg := <-s.admin:
			s.onAdmin(msg)

		case <-s.messageEvent:
			s.sendOrDropAppMessages()

		case fixIn, ok := <-s.messageIn:
			if !ok {
				s.Disconnected(s)
				continue
			}
			s.onIncoming(fixIn)

		case <-s.quit:
			s.quit = nil // prevent infinitly receiving on a closed channel
			if !s.IsLoggedOn() {
				return
			}

			if err := s.initiateLogout(""); err != nil {
				s.logError(err)
				return
			}
			s.State = logoutState{}

		case evt := <-s.sessionEvent:
			s.Timeout(s, evt)

		case now := <-ticker.C:
			s.CheckSessionTime(s, now)
		}
	}
}
