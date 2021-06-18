package quickfix

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/quickfixgo/quickfix/datadictionary"
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
	toSend [][]byte

	//mutex for access to toSend
	sendMutex sync.Mutex

	sessionEvent chan internal.Event
	messageEvent chan bool
	application  Application
	validator
	stateMachine
	stateTimer *internal.EventTimer
	peerTimer  *internal.EventTimer
	sentReset  bool

	targetDefaultApplVerID string

	admin chan interface{}
	internal.SessionSettings
	transportDataDictionary *datadictionary.DataDictionary
	appDataDictionary       *datadictionary.DataDictionary

	messagePool
	timestampPrecision TimestampPrecision
}

func (s *session) logError(err error) {
	s.log.OnEvent(err.Error())
}

//TargetDefaultApplicationVersionID returns the default application version ID for messages received by this version.
//Applicable for For FIX.T.1 sessions.
func (s *session) TargetDefaultApplicationVersionID() string {
	return s.targetDefaultApplVerID
}

type connect struct {
	messageOut chan<- []byte
	messageIn  <-chan fixIn
	err        chan<- error
}

func (s *session) connect(msgIn <-chan fixIn, msgOut chan<- []byte) error {
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

func (s *session) insertSendingTime(msg *Message) {
	sendingTime := time.Now().UTC()

	if s.sessionID.BeginString >= BeginStringFIX42 {
		msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime, Precision: s.timestampPrecision})
	} else {
		msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: sendingTime, Precision: Seconds})
	}
}

func optionallySetID(msg *Message, field Tag, value string) {
	if len(value) != 0 {
		msg.Header.SetString(field, value)
	}
}

func (s *session) fillDefaultHeader(msg *Message, inReplyTo *Message) {
	msg.Header.SetString(tagBeginString, s.sessionID.BeginString)
	msg.Header.SetString(tagSenderCompID, s.sessionID.SenderCompID)
	optionallySetID(msg, tagSenderSubID, s.sessionID.SenderSubID)
	optionallySetID(msg, tagSenderLocationID, s.sessionID.SenderLocationID)

	msg.Header.SetString(tagTargetCompID, s.sessionID.TargetCompID)
	optionallySetID(msg, tagTargetSubID, s.sessionID.TargetSubID)
	optionallySetID(msg, tagTargetLocationID, s.sessionID.TargetLocationID)

	s.insertSendingTime(msg)

	if s.EnableLastMsgSeqNumProcessed {
		if inReplyTo != nil {
			if lastSeqNum, err := inReplyTo.Header.GetInt(tagMsgSeqNum); err != nil {
				s.logError(err)
			} else {
				msg.Header.SetInt(tagLastMsgSeqNumProcessed, lastSeqNum)
			}
		} else {
			msg.Header.SetInt(tagLastMsgSeqNumProcessed, s.store.NextTargetMsgSeqNum()-1)
		}
	}
}

func (s *session) shouldSendReset() bool {
	if s.sessionID.BeginString < BeginStringFIX41 {
		return false
	}

	return (s.ResetOnLogon || s.ResetOnDisconnect || s.ResetOnLogout) &&
		s.store.NextTargetMsgSeqNum() == 1 && s.store.NextSenderMsgSeqNum() == 1
}

func (s *session) sendLogon() error {
	return s.sendLogonInReplyTo(s.shouldSendReset(), nil)
}

func (s *session) sendLogonInReplyTo(setResetSeqNum bool, inReplyTo *Message) error {
	logon := NewMessage()
	logon.Header.SetField(tagMsgType, FIXString("A"))
	logon.Header.SetField(tagBeginString, FIXString(s.sessionID.BeginString))
	logon.Header.SetField(tagTargetCompID, FIXString(s.sessionID.TargetCompID))
	logon.Header.SetField(tagSenderCompID, FIXString(s.sessionID.SenderCompID))
	logon.Body.SetField(tagEncryptMethod, FIXString("0"))
	logon.Body.SetField(tagHeartBtInt, FIXInt(s.HeartBtInt.Seconds()))

	if setResetSeqNum {
		logon.Body.SetField(tagResetSeqNumFlag, FIXBoolean(true))
	}

	if len(s.DefaultApplVerID) > 0 {
		logon.Body.SetField(tagDefaultApplVerID, FIXString(s.DefaultApplVerID))
	}

	if err := s.dropAndSendInReplyTo(logon, inReplyTo); err != nil {
		return err
	}

	return nil
}

func (s *session) buildLogout(reason string) *Message {
	logout := NewMessage()
	logout.Header.SetField(tagMsgType, FIXString("5"))
	logout.Header.SetField(tagBeginString, FIXString(s.sessionID.BeginString))
	logout.Header.SetField(tagTargetCompID, FIXString(s.sessionID.TargetCompID))
	logout.Header.SetField(tagSenderCompID, FIXString(s.sessionID.SenderCompID))
	if reason != "" {
		logout.Body.SetField(tagText, FIXString(reason))
	}

	return logout
}

func (s *session) sendLogout(reason string) error {
	return s.sendLogoutInReplyTo(reason, nil)
}

func (s *session) sendLogoutInReplyTo(reason string, inReplyTo *Message) error {
	logout := s.buildLogout(reason)
	return s.sendInReplyTo(logout, inReplyTo)
}

func (s *session) resend(msg *Message) bool {
	msg.Header.SetField(tagPossDupFlag, FIXBoolean(true))

	var origSendingTime FIXString
	if err := msg.Header.GetField(tagSendingTime, &origSendingTime); err == nil {
		msg.Header.SetField(tagOrigSendingTime, origSendingTime)
	}

	s.insertSendingTime(msg)

	return s.application.ToApp(msg, s.sessionID) == nil
}

//queueForSend will validate, persist, and queue the message for send
func (s *session) queueForSend(msg *Message) error {
	s.sendMutex.Lock()
	defer s.sendMutex.Unlock()

	msgBytes, err := s.prepMessageForSend(msg, nil)
	if err != nil {
		return err
	}

	s.toSend = append(s.toSend, msgBytes)

	select {
	case s.messageEvent <- true:
	default:
	}

	return nil
}

//send will validate, persist, queue the message. If the session is logged on, send all messages in the queue
func (s *session) send(msg *Message) error {
	return s.sendInReplyTo(msg, nil)
}
func (s *session) sendInReplyTo(msg *Message, inReplyTo *Message) error {
	if !s.IsLoggedOn() {
		return s.queueForSend(msg)
	}

	s.sendMutex.Lock()
	defer s.sendMutex.Unlock()

	msgBytes, err := s.prepMessageForSend(msg, inReplyTo)
	if err != nil {
		return err
	}

	s.toSend = append(s.toSend, msgBytes)
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

//dropAndSend will validate and persist the message, then drops the send queue and sends the message.
func (s *session) dropAndSend(msg *Message) error {
	return s.dropAndSendInReplyTo(msg, nil)
}
func (s *session) dropAndSendInReplyTo(msg *Message, inReplyTo *Message) error {
	s.sendMutex.Lock()
	defer s.sendMutex.Unlock()

	msgBytes, err := s.prepMessageForSend(msg, inReplyTo)
	if err != nil {
		return err
	}

	s.dropQueued()
	s.toSend = append(s.toSend, msgBytes)
	s.sendQueued()

	return nil
}

func (s *session) prepMessageForSend(msg *Message, inReplyTo *Message) (msgBytes []byte, err error) {
	s.fillDefaultHeader(msg, inReplyTo)
	seqNum := s.store.NextSenderMsgSeqNum()
	msg.Header.SetField(tagMsgSeqNum, FIXInt(seqNum))

	msgType, err := msg.Header.GetBytes(tagMsgType)
	if err != nil {
		return
	}

	if isAdminMessageType(msgType) {
		s.application.ToAdmin(msg, s.sessionID)

		if bytes.Equal(msgType, msgTypeLogon) {
			var resetSeqNumFlag FIXBoolean
			if msg.Body.Has(tagResetSeqNumFlag) {
				if err = msg.Body.GetField(tagResetSeqNumFlag, &resetSeqNumFlag); err != nil {
					return
				}
			}

			if resetSeqNumFlag.Bool() {
				if err = s.store.Reset(); err != nil {
					return
				}

				s.sentReset = true
				seqNum = s.store.NextSenderMsgSeqNum()
				msg.Header.SetField(tagMsgSeqNum, FIXInt(seqNum))
			}
		}
	} else {
		if err = s.application.ToApp(msg, s.sessionID); err != nil {
			return
		}
	}

	msgBytes = msg.build()
	err = s.persist(seqNum, msgBytes)

	return
}

func (s *session) persist(seqNum int, msgBytes []byte) error {
	if !s.DisableMessagePersist {
		if err := s.store.SaveMessage(seqNum, msgBytes); err != nil {
			return err
		}
	}

	return s.store.IncrNextSenderMsgSeqNum()
}

func (s *session) sendQueued() {
	for _, msgBytes := range s.toSend {
		s.sendBytes(msgBytes)
	}

	s.dropQueued()
}

func (s *session) dropQueued() {
	s.toSend = s.toSend[:0]
}

func (s *session) EnqueueBytesAndSend(msg []byte) {
	s.sendMutex.Lock()
	defer s.sendMutex.Unlock()

	s.toSend = append(s.toSend, msg)
	s.sendQueued()
}

func (s *session) sendBytes(msg []byte) {
	s.log.OnOutgoing(msg)
	s.messageOut <- msg
	s.stateTimer.Reset(s.HeartBtInt)
}

func (s *session) doTargetTooHigh(reject targetTooHigh) (nextState resendState, err error) {
	s.log.OnEventf("MsgSeqNum too high, expecting %v but received %v", reject.ExpectedTarget, reject.ReceivedTarget)
	return s.sendResendRequest(reject.ExpectedTarget, reject.ReceivedTarget-1)
}

func (s *session) sendResendRequest(beginSeq, endSeq int) (nextState resendState, err error) {
	nextState.resendRangeEnd = endSeq

	resend := NewMessage()
	resend.Header.SetBytes(tagMsgType, msgTypeResendRequest)
	resend.Body.SetField(tagBeginSeqNo, FIXInt(beginSeq))

	var endSeqNo int
	if s.ResendRequestChunkSize != 0 {
		endSeqNo = beginSeq + s.ResendRequestChunkSize - 1
	} else {
		endSeqNo = endSeq
	}

	if endSeqNo < endSeq {
		nextState.currentResendRangeEnd = endSeqNo
	} else {
		if s.sessionID.BeginString < BeginStringFIX42 {
			endSeqNo = 999999
		} else {
			endSeqNo = 0
		}
	}
	resend.Body.SetField(tagEndSeqNo, FIXInt(endSeqNo))

	if err = s.send(resend); err != nil {
		return
	}
	s.log.OnEventf("Sent ResendRequest FROM: %v TO: %v", beginSeq, endSeqNo)

	return
}

func (s *session) handleLogon(msg *Message) error {
	//Grab default app ver id from fixt.1.1 logon
	if s.sessionID.BeginString == BeginStringFIXT11 {
		var targetApplVerID FIXString

		if err := msg.Body.GetField(tagDefaultApplVerID, &targetApplVerID); err != nil {
			return err
		}

		s.targetDefaultApplVerID = string(targetApplVerID)
	}

	resetStore := false
	if s.InitiateLogon {
		s.log.OnEvent("Received logon response")
	} else {
		s.log.OnEvent("Received logon request")
		resetStore = s.ResetOnLogon

		if s.RefreshOnLogon {
			if err := s.store.Refresh(); err != nil {
				return err
			}
		}
	}

	var resetSeqNumFlag FIXBoolean
	if err := msg.Body.GetField(tagResetSeqNumFlag, &resetSeqNumFlag); err == nil {
		if resetSeqNumFlag {
			if !s.sentReset {
				s.log.OnEvent("Logon contains ResetSeqNumFlag=Y, resetting sequence numbers to 1")
				resetStore = true
			}
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

	if !s.InitiateLogon {
		var heartBtInt FIXInt
		if err := msg.Body.GetField(tagHeartBtInt, &heartBtInt); err == nil {
			s.HeartBtInt = time.Duration(heartBtInt) * time.Second
		}

		s.log.OnEvent("Responding to logon request")
		if err := s.sendLogonInReplyTo(resetSeqNumFlag.Bool(), msg); err != nil {
			return err
		}
	}
	s.sentReset = false

	s.peerTimer.Reset(time.Duration(float64(1.2) * float64(s.HeartBtInt)))
	s.application.OnLogon(s.sessionID)

	if err := s.checkTargetTooHigh(msg); err != nil {
		return err
	}

	return s.store.IncrNextTargetMsgSeqNum()
}

func (s *session) initiateLogout(reason string) (err error) {
	return s.initiateLogoutInReplyTo(reason, nil)
}

func (s *session) initiateLogoutInReplyTo(reason string, inReplyTo *Message) (err error) {
	if err = s.sendLogoutInReplyTo(reason, inReplyTo); err != nil {
		s.logError(err)
		return
	}
	s.log.OnEvent("Inititated logout request")
	time.AfterFunc(s.LogoutTimeout, func() { s.sessionEvent <- internal.LogoutTimeout })
	return
}

func (s *session) verify(msg *Message) MessageRejectError {
	return s.verifySelect(msg, true, true)
}

func (s *session) verifyIgnoreSeqNumTooHigh(msg *Message) MessageRejectError {
	return s.verifySelect(msg, false, true)
}

func (s *session) verifyIgnoreSeqNumTooHighOrLow(msg *Message) MessageRejectError {
	return s.verifySelect(msg, false, false)
}

func (s *session) verifySelect(msg *Message, checkTooHigh bool, checkTooLow bool) MessageRejectError {
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

func (s *session) fromCallback(msg *Message) MessageRejectError {
	msgType, err := msg.Header.GetBytes(tagMsgType)
	if err != nil {
		return err
	}

	if isAdminMessageType(msgType) {
		return s.application.FromAdmin(msg, s.sessionID)
	}

	return s.application.FromApp(msg, s.sessionID)
}

func (s *session) checkTargetTooLow(msg *Message) MessageRejectError {
	if !msg.Header.Has(tagMsgSeqNum) {
		return RequiredTagMissing(tagMsgSeqNum)
	}

	seqNum, err := msg.Header.GetInt(tagMsgSeqNum)
	if err != nil {
		return err
	}

	if seqNum < s.store.NextTargetMsgSeqNum() {
		return targetTooLow{ReceivedTarget: seqNum, ExpectedTarget: s.store.NextTargetMsgSeqNum()}
	}

	return nil
}

func (s *session) checkTargetTooHigh(msg *Message) MessageRejectError {
	if !msg.Header.Has(tagMsgSeqNum) {
		return RequiredTagMissing(tagMsgSeqNum)
	}

	seqNum, err := msg.Header.GetInt(tagMsgSeqNum)
	if err != nil {
		return err
	}

	if seqNum > s.store.NextTargetMsgSeqNum() {
		return targetTooHigh{ReceivedTarget: seqNum, ExpectedTarget: s.store.NextTargetMsgSeqNum()}
	}

	return nil
}

func (s *session) checkCompID(msg *Message) MessageRejectError {
	senderCompID, haveSender := msg.Header.GetBytes(tagSenderCompID)
	targetCompID, haveTarget := msg.Header.GetBytes(tagTargetCompID)

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

func (s *session) checkSendingTime(msg *Message) MessageRejectError {
	if s.SkipCheckLatency {
		return nil
	}

	if ok := msg.Header.Has(tagSendingTime); !ok {
		return RequiredTagMissing(tagSendingTime)
	}

	sendingTime, err := msg.Header.GetTime(tagSendingTime)
	if err != nil {
		return err
	}

	if delta := time.Since(sendingTime); delta <= -1*s.MaxLatency || delta >= s.MaxLatency {
		return sendingTimeAccuracyProblem()
	}

	return nil
}

func (s *session) checkBeginString(msg *Message) MessageRejectError {
	switch beginString, err := msg.Header.GetBytes(tagBeginString); {
	case err != nil:
		return RequiredTagMissing(tagBeginString)
	case s.sessionID.BeginString != string(beginString):
		return incorrectBeginString{}
	}

	return nil
}

func (s *session) doReject(msg *Message, rej MessageRejectError) error {
	reply := msg.reverseRoute()

	if s.sessionID.BeginString >= BeginStringFIX42 {

		if rej.IsBusinessReject() {
			reply.Header.SetField(tagMsgType, FIXString("j"))
			reply.Body.SetField(tagBusinessRejectReason, FIXInt(rej.RejectReason()))
			if refID := rej.BusinessRejectRefID(); refID != "" {
				reply.Body.SetField(tagBusinessRejectRefID, FIXString(refID))
			}
		} else {
			reply.Header.SetField(tagMsgType, FIXString("3"))
			switch {
			default:
				reply.Body.SetField(tagSessionRejectReason, FIXInt(rej.RejectReason()))
			case rej.RejectReason() > rejectReasonInvalidMsgType && s.sessionID.BeginString == BeginStringFIX42:
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
	return s.sendInReplyTo(reply, msg)
}

type fixIn struct {
	bytes       *bytes.Buffer
	receiveTime time.Time
}

func (s *session) returnToPool(msg *Message) {
	s.messagePool.Put(msg)
	if msg.rawMessage != nil {
		bufferPool.Put(msg.rawMessage)
		msg.rawMessage = nil
	}
}

func (s *session) onDisconnect() {
	s.log.OnEvent("Disconnected")
	if s.ResetOnDisconnect {
		if err := s.dropAndReset(); err != nil {
			s.logError(err)
		}
	}

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

		if !s.IsSessionTime() {
			s.handleDisconnectState(s)
			if msg.err != nil {
				msg.err <- errors.New("Connection outside of session time")
				close(msg.err)
			}
			return
		}

		if msg.err != nil {
			close(msg.err)
		}

		s.messageIn = msg.messageIn
		s.messageOut = msg.messageOut
		s.sentReset = false

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

	s.stateTimer = internal.NewEventTimer(func() { s.sessionEvent <- internal.NeedHeartbeat })
	s.peerTimer = internal.NewEventTimer(func() { s.sessionEvent <- internal.PeerTimeout })
	ticker := time.NewTicker(time.Second)

	defer func() {
		s.stateTimer.Stop()
		s.peerTimer.Stop()
		ticker.Stop()
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

		case now := <-ticker.C:
			s.CheckSessionTime(s, now)
		}
	}
}
