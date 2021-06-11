package quickfix

import (
	"bytes"
	"time"

	"github.com/quickfixgo/quickfix/internal"
)

type inSession struct{ loggedOn }

func (state inSession) String() string { return "In Session" }

func (state inSession) FixMsgIn(session *session, msg *Message) sessionState {
	msgType, err := msg.Header.GetBytes(tagMsgType)
	if err != nil {
		return handleStateError(session, err)
	}

	switch {
	case bytes.Equal(msgTypeLogon, msgType):
		if err := session.handleLogon(msg); err != nil {
			if err := session.initiateLogoutInReplyTo("", msg); err != nil {
				return handleStateError(session, err)
			}
			return logoutState{}
		}

		return state
	case bytes.Equal(msgTypeLogout, msgType):
		return state.handleLogout(session, msg)
	case bytes.Equal(msgTypeResendRequest, msgType):
		return state.handleResendRequest(session, msg)
	case bytes.Equal(msgTypeSequenceReset, msgType):
		return state.handleSequenceReset(session, msg)
	case bytes.Equal(msgTypeTestRequest, msgType):
		return state.handleTestRequest(session, msg)
	default:
		if err := session.verify(msg); err != nil {
			return state.processReject(session, msg, err)
		}
	}

	if err := session.store.IncrNextTargetMsgSeqNum(); err != nil {
		return handleStateError(session, err)
	}

	return state
}

func (state inSession) Timeout(session *session, event internal.Event) (nextState sessionState) {
	switch event {
	case internal.NeedHeartbeat:
		heartBt := NewMessage()
		heartBt.Header.SetField(tagMsgType, FIXString("0"))
		if err := session.send(heartBt); err != nil {
			return handleStateError(session, err)
		}
	case internal.PeerTimeout:
		testReq := NewMessage()
		testReq.Header.SetField(tagMsgType, FIXString("1"))
		testReq.Body.SetField(tagTestReqID, FIXString("TEST"))
		if err := session.send(testReq); err != nil {
			return handleStateError(session, err)
		}
		session.log.OnEvent("Sent test request TEST")
		session.peerTimer.Reset(time.Duration(float64(1.2) * float64(session.HeartBtInt)))
		return pendingTimeout{state}
	}

	return state
}

func (state inSession) handleLogout(session *session, msg *Message) (nextState sessionState) {
	if err := session.verifySelect(msg, false, false); err != nil {
		return state.processReject(session, msg, err)
	}

	if session.IsLoggedOn() {
		session.log.OnEvent("Received logout request")
		session.log.OnEvent("Sending logout response")

		if err := session.sendLogoutInReplyTo("", msg); err != nil {
			session.logError(err)
		}
	} else {
		session.log.OnEvent("Received logout response")
	}

	if err := session.store.IncrNextTargetMsgSeqNum(); err != nil {
		session.logError(err)
	}

	if session.ResetOnLogout {
		if err := session.dropAndReset(); err != nil {
			session.logError(err)
		}
	}

	return latentState{}
}

func (state inSession) handleTestRequest(session *session, msg *Message) (nextState sessionState) {
	if err := session.verify(msg); err != nil {
		return state.processReject(session, msg, err)
	}
	var testReq FIXString
	if err := msg.Body.GetField(tagTestReqID, &testReq); err != nil {
		session.log.OnEvent("Test Request with no testRequestID")
	} else {
		heartBt := NewMessage()
		heartBt.Header.SetField(tagMsgType, FIXString("0"))
		heartBt.Body.SetField(tagTestReqID, testReq)
		if err := session.sendInReplyTo(heartBt, msg); err != nil {
			return handleStateError(session, err)
		}
	}

	if err := session.store.IncrNextTargetMsgSeqNum(); err != nil {
		return handleStateError(session, err)
	}
	return state
}

func (state inSession) handleSequenceReset(session *session, msg *Message) (nextState sessionState) {
	var gapFillFlag FIXBoolean
	if msg.Body.Has(tagGapFillFlag) {
		if err := msg.Body.GetField(tagGapFillFlag, &gapFillFlag); err != nil {
			return state.processReject(session, msg, err)
		}
	}

	if err := session.verifySelect(msg, bool(gapFillFlag), bool(gapFillFlag)); err != nil {
		return state.processReject(session, msg, err)
	}

	var newSeqNo FIXInt
	if err := msg.Body.GetField(tagNewSeqNo, &newSeqNo); err == nil {
		expectedSeqNum := FIXInt(session.store.NextTargetMsgSeqNum())
		session.log.OnEventf("Received SequenceReset FROM: %v TO: %v", expectedSeqNum, newSeqNo)

		switch {
		case newSeqNo > expectedSeqNum:
			if err := session.store.SetNextTargetMsgSeqNum(int(newSeqNo)); err != nil {
				return handleStateError(session, err)
			}
		case newSeqNo < expectedSeqNum:
			//FIXME: to be compliant with legacy tests, do not include tag in reftagid? (11c_NewSeqNoLess)
			if err := session.doReject(msg, valueIsIncorrectNoTag()); err != nil {
				return handleStateError(session, err)
			}
		}
	}
	return state
}

func (state inSession) handleResendRequest(session *session, msg *Message) (nextState sessionState) {
	if err := session.verifyIgnoreSeqNumTooHighOrLow(msg); err != nil {
		return state.processReject(session, msg, err)
	}

	var err error
	var beginSeqNoField FIXInt
	if err = msg.Body.GetField(tagBeginSeqNo, &beginSeqNoField); err != nil {
		return state.processReject(session, msg, RequiredTagMissing(tagBeginSeqNo))
	}

	beginSeqNo := beginSeqNoField

	var endSeqNoField FIXInt
	if err = msg.Body.GetField(tagEndSeqNo, &endSeqNoField); err != nil {
		return state.processReject(session, msg, RequiredTagMissing(tagEndSeqNo))
	}

	endSeqNo := int(endSeqNoField)

	session.log.OnEventf("Received ResendRequest FROM: %d TO: %d", beginSeqNo, endSeqNo)
	expectedSeqNum := session.store.NextSenderMsgSeqNum()

	if (session.sessionID.BeginString >= BeginStringFIX42 && endSeqNo == 0) ||
		(session.sessionID.BeginString <= BeginStringFIX42 && endSeqNo == 999999) ||
		(endSeqNo >= expectedSeqNum) {
		endSeqNo = expectedSeqNum - 1
	}

	if err := state.resendMessages(session, int(beginSeqNo), endSeqNo, *msg); err != nil {
		return handleStateError(session, err)
	}

	if err := session.checkTargetTooLow(msg); err != nil {
		return state
	}

	if err := session.checkTargetTooHigh(msg); err != nil {
		return state
	}

	if err := session.store.IncrNextTargetMsgSeqNum(); err != nil {
		return handleStateError(session, err)
	}
	return state
}

func (state inSession) resendMessages(session *session, beginSeqNo, endSeqNo int, inReplyTo Message) (err error) {
	if session.DisableMessagePersist {
		err = state.generateSequenceReset(session, beginSeqNo, endSeqNo+1, inReplyTo)
		return
	}

	msgs, err := session.store.GetMessages(beginSeqNo, endSeqNo)
	if err != nil {
		session.log.OnEventf("error retrieving messages from store: %s", err.Error())
		return
	}

	seqNum := beginSeqNo
	nextSeqNum := seqNum
	msg := NewMessage()
	for _, msgBytes := range msgs {
		_ = ParseMessageWithDataDictionary(msg, bytes.NewBuffer(msgBytes), session.transportDataDictionary, session.appDataDictionary)
		msgType, _ := msg.Header.GetBytes(tagMsgType)
		sentMessageSeqNum, _ := msg.Header.GetInt(tagMsgSeqNum)

		if isAdminMessageType(msgType) {
			nextSeqNum = sentMessageSeqNum + 1
			continue
		}

		if !session.resend(msg) {
			nextSeqNum = sentMessageSeqNum + 1
			continue
		}

		if seqNum != sentMessageSeqNum {
			if err = state.generateSequenceReset(session, seqNum, sentMessageSeqNum, inReplyTo); err != nil {
				return err
			}
		}

		session.log.OnEventf("Resending Message: %v", sentMessageSeqNum)
		msgBytes = msg.build()
		session.EnqueueBytesAndSend(msgBytes)

		seqNum = sentMessageSeqNum + 1
		nextSeqNum = seqNum
	}

	if seqNum != nextSeqNum { // gapfill for catch-up
		if err = state.generateSequenceReset(session, seqNum, nextSeqNum, inReplyTo); err != nil {
			return err
		}
	}

	return
}

func (state inSession) processReject(session *session, msg *Message, rej MessageRejectError) sessionState {
	switch TypedError := rej.(type) {
	case targetTooHigh:

		var nextState resendState
		switch currentState := session.State.(type) {
		case resendState:
			//assumes target too high reject already sent
			nextState = currentState
		default:
			var err error
			if nextState, err = session.doTargetTooHigh(TypedError); err != nil {
				return handleStateError(session, err)
			}
		}

		if nextState.messageStash == nil {
			nextState.messageStash = make(map[int]*Message)
		}

		nextState.messageStash[TypedError.ReceivedTarget] = msg
		//do not reclaim stashed message
		msg.keepMessage = true

		return nextState

	case targetTooLow:
		return state.doTargetTooLow(session, msg, TypedError)
	case incorrectBeginString:
		if err := session.initiateLogout(rej.Error()); err != nil {
			return handleStateError(session, err)
		}
		return logoutState{}
	}

	switch rej.RejectReason() {
	case rejectReasonCompIDProblem, rejectReasonSendingTimeAccuracyProblem:
		if err := session.doReject(msg, rej); err != nil {
			return handleStateError(session, err)
		}

		if err := session.initiateLogout(""); err != nil {
			return handleStateError(session, err)
		}
		return logoutState{}
	default:
		if err := session.doReject(msg, rej); err != nil {
			return handleStateError(session, err)
		}

		if err := session.store.IncrNextTargetMsgSeqNum(); err != nil {
			return handleStateError(session, err)
		}
		return state
	}
}

func (state inSession) doTargetTooLow(session *session, msg *Message, rej targetTooLow) (nextState sessionState) {
	var posDupFlag FIXBoolean
	if msg.Header.Has(tagPossDupFlag) {
		if err := msg.Header.GetField(tagPossDupFlag, &posDupFlag); err != nil {
			if rejErr := session.doReject(msg, err); rejErr != nil {
				return handleStateError(session, rejErr)
			}
			return state
		}
	}

	if !posDupFlag.Bool() {
		if err := session.initiateLogout(rej.Error()); err != nil {
			return handleStateError(session, err)
		}
		return logoutState{}
	}

	if !msg.Header.Has(tagOrigSendingTime) {
		if err := session.doReject(msg, RequiredTagMissing(tagOrigSendingTime)); err != nil {
			return handleStateError(session, err)
		}
		return state
	}

	var origSendingTime FIXUTCTimestamp
	if err := msg.Header.GetField(tagOrigSendingTime, &origSendingTime); err != nil {
		if rejErr := session.doReject(msg, err); rejErr != nil {
			return handleStateError(session, rejErr)
		}
		return state
	}

	sendingTime := new(FIXUTCTimestamp)
	if err := msg.Header.GetField(tagSendingTime, sendingTime); err != nil {
		return state.processReject(session, msg, err)
	}

	if sendingTime.Before(origSendingTime.Time) {
		if err := session.doReject(msg, sendingTimeAccuracyProblem()); err != nil {
			return handleStateError(session, err)
		}

		if err := session.initiateLogout(""); err != nil {
			return handleStateError(session, err)
		}
		return logoutState{}
	}

	return state
}

func (state *inSession) generateSequenceReset(session *session, beginSeqNo int, endSeqNo int, inReplyTo Message) (err error) {
	sequenceReset := NewMessage()
	session.fillDefaultHeader(sequenceReset, &inReplyTo)

	sequenceReset.Header.SetField(tagMsgType, FIXString("4"))
	sequenceReset.Header.SetField(tagMsgSeqNum, FIXInt(beginSeqNo))
	sequenceReset.Header.SetField(tagPossDupFlag, FIXBoolean(true))
	sequenceReset.Body.SetField(tagNewSeqNo, FIXInt(endSeqNo))
	sequenceReset.Body.SetField(tagGapFillFlag, FIXBoolean(true))

	var origSendingTime FIXString
	if err := sequenceReset.Header.GetField(tagSendingTime, &origSendingTime); err == nil {
		sequenceReset.Header.SetField(tagOrigSendingTime, origSendingTime)
	}

	session.application.ToAdmin(sequenceReset, session.sessionID)

	msgBytes := sequenceReset.build()

	session.EnqueueBytesAndSend(msgBytes)
	session.log.OnEventf("Sent SequenceReset TO: %v", endSeqNo)

	return
}
