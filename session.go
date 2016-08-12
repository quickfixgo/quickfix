package quickfix

import (
	"errors"
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
	stateTimer     internal.EventTimer
	peerTimer      internal.EventTimer
	messageStash   map[int]Message
	resetOnLogon   bool
	refreshOnLogon bool
	resetOnLogout  bool
	initiateLogon  bool
	heartBtInt     time.Duration

	//required on logon for FIX.T.1 messages
	defaultApplVerID       string
	targetDefaultApplVerID string

	sessionTime *internal.TimeRange
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

var dayLookup = map[string]time.Weekday{
	"Sunday":    time.Sunday,
	"Monday":    time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Saturday":  time.Saturday,

	"Sun": time.Sunday,
	"Mon": time.Monday,
	"Tue": time.Tuesday,
	"Wed": time.Wednesday,
	"Thu": time.Thursday,
	"Fri": time.Friday,
	"Sat": time.Saturday,
}

var applVerIDLookup = map[string]string{
	enum.BeginStringFIX40: "2",
	enum.BeginStringFIX41: "3",
	enum.BeginStringFIX42: "4",
	enum.BeginStringFIX43: "5",
	enum.BeginStringFIX44: "6",
	"FIX.5.0":             "7",
	"FIX.5.0SP1":          "8",
	"FIX.5.0SP2":          "9",
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
		if session.defaultApplVerID, err = settings.Setting(config.DefaultApplVerID); err != nil {
			return session, requiredConfigurationMissing(config.DefaultApplVerID)
		}

		if applVerID, ok := applVerIDLookup[session.defaultApplVerID]; ok {
			session.defaultApplVerID = applVerID
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

	if settings.HasSetting(config.RefreshOnLogon) {
		if session.refreshOnLogon, err = settings.BoolSetting(config.RefreshOnLogon); err != nil {
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
			session.heartBtInt = time.Duration(heartBtInt) * time.Second
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

		start, err := internal.ParseTimeOfDay(startTimeStr)
		if err != nil {
			return session, err
		}

		end, err := internal.ParseTimeOfDay(endTimeStr)
		if err != nil {
			return session, err
		}

		loc := time.UTC
		if settings.HasSetting(config.TimeZone) {
			locStr, err := settings.Setting(config.TimeZone)
			if err != nil {
				return session, err
			}

			loc, err = time.LoadLocation(locStr)
			if err != nil {
				return session, err
			}
		}

		switch {
		case !settings.HasSetting(config.StartDay) && !settings.HasSetting(config.EndDay):
			session.sessionTime = internal.NewTimeRangeInLocation(start, end, loc)

		case settings.HasSetting(config.StartDay) && settings.HasSetting(config.EndDay):
			var startDayStr, endDayStr string
			if startDayStr, err = settings.Setting(config.StartDay); err != nil {
				return session, err
			}

			if endDayStr, err = settings.Setting(config.EndDay); err != nil {
				return session, err
			}

			parseDay := func(dayStr string) (day time.Weekday, err error) {
				day, ok := dayLookup[dayStr]
				if !ok {
					err = fmt.Errorf("Cannot parse %v", dayStr)
				}
				return
			}

			startDay, err := parseDay(startDayStr)
			if err != nil {
				return session, err
			}

			endDay, err := parseDay(endDayStr)
			if err != nil {
				return session, err
			}

			session.sessionTime = internal.NewWeekRangeInLocation(start, end, startDay, endDay, loc)
		case settings.HasSetting(config.StartDay):
			return session, requiredConfigurationMissing(config.EndDay)
		case settings.HasSetting(config.EndDay):
			return session, requiredConfigurationMissing(config.StartDay)
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
	logFactory LogFactory, application Application,
) (session *session, err error) {

	if session, err = newSession(sessionID, storeFactory, settings, logFactory, application); err != nil {
		return
	}

	if err = registerSession(session); err != nil {
		return
	}
	application.OnCreate(session.sessionID)
	session.log.OnEvent("Created session")

	return
}

type connect struct {
	messageOut    chan<- []byte
	messageIn     <-chan fixIn
	initiateLogon bool
	err           chan<- error
}

//kicks off session as an initiator
func (s *session) initiate(msgIn <-chan fixIn, msgOut chan<- []byte) error {
	rep := make(chan error)
	s.admin <- connect{
		messageOut:    msgOut,
		messageIn:     msgIn,
		initiateLogon: true,
		err:           rep,
	}

	return <-rep
}

//kicks off session as an acceptor
func (s *session) accept(msgIn chan fixIn, msgOut chan []byte) error {
	rep := make(chan error)
	s.admin <- connect{
		messageOut: msgOut,
		messageIn:  msgIn,
		err:        rep,
	}
	return <-rep
}

type stopReq struct{}

func (s *session) stop() {
	s.admin <- stopReq{}
}

type waitChan <-chan interface{}

type waitForInSessionReq struct{ rep chan<- waitChan }

func (s *session) waitForInSessionTime() {
	rep := make(chan waitChan)
	s.admin <- waitForInSessionReq{rep}
	if wait, ok := <-rep; ok {
		<-wait
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

func (s *session) sendLogon(resetStore, setResetSeqNum bool) error {
	logon := NewMessage()
	logon.Header.SetField(tagMsgType, FIXString("A"))
	logon.Header.SetField(tagBeginString, FIXString(s.sessionID.BeginString))
	logon.Header.SetField(tagTargetCompID, FIXString(s.sessionID.TargetCompID))
	logon.Header.SetField(tagSenderCompID, FIXString(s.sessionID.SenderCompID))
	logon.Body.SetField(tagEncryptMethod, FIXString("0"))
	logon.Body.SetField(tagHeartBtInt, FIXInt(s.heartBtInt.Seconds()))

	if setResetSeqNum {
		logon.Body.SetField(tagResetSeqNumFlag, FIXBoolean(true))
	}

	if len(s.defaultApplVerID) > 0 {
		logon.Body.SetField(tagDefaultApplVerID, FIXString(s.defaultApplVerID))
	}

	if err := s.dropAndSend(logon, resetStore); err != nil {
		return err
	}

	return nil
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

func (s *session) sendBytes(msg []byte) {
	s.log.OnOutgoing(string(msg))
	s.messageOut <- msg
	s.stateTimer.Reset(s.heartBtInt)
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

		if s.refreshOnLogon {
			if err := s.store.Refresh(); err != nil {
				return err
			}
		}
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
		var heartBtInt FIXInt
		if err := msg.Body.GetField(tagHeartBtInt, &heartBtInt); err == nil {
			s.heartBtInt = time.Duration(heartBtInt) * time.Second
		}

		s.log.OnEvent("Responding to logon request")
		if err := s.sendLogon(resetStore, resetSeqNumFlag.Bool()); err != nil {
			return err
		}
	}

	s.peerTimer.Reset(time.Duration(float64(1.2) * float64(s.heartBtInt)))
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
	s.log.OnEvent("Disconnected")
	if s.messageOut != nil {
		close(s.messageOut)
		s.messageOut = nil
	}

	s.messageIn = nil
}

func (s *session) onAdmin(msg interface{}) {
	switch msg := msg.(type) {

	case connect:

		if s.IsConnected() {
			if msg.err != nil {
				msg.err <- errors.New("Already connected")
				close(msg.err)
			}
			return
		}

		if msg.err != nil {
			close(msg.err)
		}

		s.initiateLogon = msg.initiateLogon
		s.messageIn = msg.messageIn
		s.messageOut = msg.messageOut
		s.messageStash = make(map[int]Message)

		s.Connect(s)

	case stopReq:
		s.Stop(s)

	case waitForInSessionReq:
		if !s.IsSessionTime() {
			msg.rep <- s.stateMachine.notifyOnInSessionTime
		}
		close(msg.rep)
	}
}

func (s *session) run() {
	s.Start(s)

	defer func() {
		s.stateTimer.Stop()
		s.peerTimer.Stop()
	}()

	for !s.Stopped() {
		select {

		case msg := <-s.admin:
			s.onAdmin(msg)

		case <-s.messageEvent:
			s.SendAppMessages(s)

		case fixIn, ok := <-s.messageIn:
			if !ok {
				s.Disconnected(s)
			} else {
				s.Incoming(s, fixIn)
			}

		case evt := <-s.sessionEvent:
			s.Timeout(s, evt)

		case now := <-time.After(time.Second):
			s.CheckSessionTime(s, now)
		}
	}
}
