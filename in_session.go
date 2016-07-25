package quickfix

import (
	"time"

	"github.com/quickfixgo/quickfix/enum"
)

type inSession struct {
}

func (state inSession) String() string   { return "In Session" }
func (state inSession) IsLoggedOn() bool { return true }

func (state inSession) VerifyMsgIn(session *session, msg Message) (err MessageRejectError) {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err == nil {
		switch string(msgType) {
		case enum.MsgType_LOGON:
			return session.verifyLogon(msg)
		case enum.MsgType_LOGOUT:
			return nil
		case enum.MsgType_RESEND_REQUEST:
			return session.verifyIgnoreSeqNumTooHighOrLow(msg)
		case enum.MsgType_SEQUENCE_RESET:
			var gapFillFlag FIXBoolean
			msg.Body.GetField(tagGapFillFlag, &gapFillFlag)
			return session.verifySelect(msg, bool(gapFillFlag), bool(gapFillFlag))
		default:
			return session.verify(msg)
		}
	}
	return nil
}

func (state inSession) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	nextState = state

	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err == nil {
		switch string(msgType) {
		case enum.MsgType_LOGON:
			if err := session.handleLogon(msg); err != nil {
				return session.handleError(err)
			}

			return
		case enum.MsgType_LOGOUT:
			session.log.OnEvent("Received logout request")
			session.log.OnEvent("Sending logout response")
			if err := session.sendLogout(""); err != nil {
				return session.handleError(err)
			}
			nextState = latentState{}
		case enum.MsgType_TEST_REQUEST:
			return state.handleTestRequest(session, msg)
		case enum.MsgType_RESEND_REQUEST:
			return state.handleResendRequest(session, msg)
		case enum.MsgType_SEQUENCE_RESET:
			return state.handleSequenceReset(session, msg)
		}
	}

	if err := session.store.IncrNextTargetMsgSeqNum(); err != nil {
		return session.handleError(err)
	}
	return
}

func (state inSession) FixMsgInRej(session *session, msg Message, rej MessageRejectError) (nextState sessionState) {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err == nil {
		switch string(msgType) {
		case enum.MsgType_LOGON:
			if err := session.initiateLogout(""); err != nil {
				return session.handleError(err)
			}
			return logoutState{}
		case enum.MsgType_LOGOUT:
			return latentState{}
		}
	}
	return state.processReject(session, msg, rej)
}

func (state inSession) Timeout(session *session, event event) (nextState sessionState) {
	switch event {
	case needHeartbeat:
		heartBt := NewMessage()
		heartBt.Header.SetField(tagMsgType, FIXString("0"))
		if err := session.send(heartBt); err != nil {
			return session.handleError(err)
		}
	case peerTimeout:
		testReq := NewMessage()
		testReq.Header.SetField(tagMsgType, FIXString("1"))
		testReq.Body.SetField(tagTestReqID, FIXString("TEST"))
		if err := session.send(testReq); err != nil {
			return session.handleError(err)
		}
		session.log.OnEvent("Sent test request TEST")
		session.peerTimer.Reset(time.Duration(int64(1.2 * float64(session.heartBeatTimeout))))
		return pendingTimeout{}
	}
	return state
}

func (state inSession) handleTestRequest(session *session, msg Message) (nextState sessionState) {
	var testReq FIXString
	if err := msg.Body.GetField(tagTestReqID, &testReq); err != nil {
		session.log.OnEvent("Test Request with no testRequestID")
	} else {
		heartBt := NewMessage()
		heartBt.Header.SetField(tagMsgType, FIXString("0"))
		heartBt.Body.SetField(tagTestReqID, testReq)
		if err := session.send(heartBt); err != nil {
			return session.handleError(err)
		}
	}

	if err := session.store.IncrNextTargetMsgSeqNum(); err != nil {
		return session.handleError(err)
	}
	return state
}

func (state inSession) handleSequenceReset(session *session, msg Message) (nextState sessionState) {
	var newSeqNo FIXInt
	if err := msg.Body.GetField(tagNewSeqNo, &newSeqNo); err == nil {
		expectedSeqNum := FIXInt(session.store.NextTargetMsgSeqNum())
		session.log.OnEventf("Received SequenceReset FROM: %v TO: %v", expectedSeqNum, newSeqNo)

		switch {
		case newSeqNo > expectedSeqNum:
			if err := session.store.SetNextTargetMsgSeqNum(int(newSeqNo)); err != nil {
				return session.handleError(err)
			}
		case newSeqNo < expectedSeqNum:
			//FIXME: to be compliant with legacy tests, do not include tag in reftagid? (11c_NewSeqNoLess)
			if err := session.doReject(msg, valueIsIncorrectNoTag()); err != nil {
				return session.handleError(err)
			}
		}
	}
	return state
}

func (state inSession) handleResendRequest(session *session, msg Message) (nextState sessionState) {
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
	expectedSeqNum := session.store.NextTargetMsgSeqNum()

	if (session.sessionID.BeginString >= enum.BeginStringFIX42 && endSeqNo == 0) ||
		(session.sessionID.BeginString <= enum.BeginStringFIX42 && endSeqNo == 999999) ||
		(endSeqNo >= expectedSeqNum) {
		endSeqNo = expectedSeqNum - 1
	}

	if err := state.resendMessages(session, int(beginSeqNo), endSeqNo); err != nil {
		return session.handleError(err)
	}

	if err := session.store.IncrNextTargetMsgSeqNum(); err != nil {
		return session.handleError(err)
	}
	return state
}

func (state inSession) resendMessages(session *session, beginSeqNo, endSeqNo int) (err error) {
	msgs, err := session.store.GetMessages(beginSeqNo, endSeqNo)
	if err != nil {
		session.log.OnEventf("error retrieving messages from store: %s", err.Error())
		return
	}

	seqNum := beginSeqNo
	nextSeqNum := seqNum
	for _, msgBytes := range msgs {
		msg, _ := ParseMessage(msgBytes)
		msgType, _ := msg.Header.GetString(tagMsgType)
		sentMessageSeqNum, _ := msg.Header.GetInt(tagMsgSeqNum)

		if isAdminMessageType(msgType) {
			nextSeqNum = sentMessageSeqNum + 1
			continue
		}

		if seqNum != sentMessageSeqNum {
			state.generateSequenceReset(session, seqNum, sentMessageSeqNum)
		}

		session.log.OnEventf("Resending Message: %v", sentMessageSeqNum)
		if err = session.resend(msg); err != nil {
			return
		}

		seqNum = sentMessageSeqNum + 1
		nextSeqNum = seqNum
	}

	if seqNum != nextSeqNum { // gapfill for catch-up
		state.generateSequenceReset(session, seqNum, nextSeqNum)
	}

	return
}

func (state inSession) processReject(session *session, msg Message, rej MessageRejectError) (nextState sessionState) {
	switch TypedError := rej.(type) {
	case targetTooHigh:

		switch session.sessionState.(type) {
		default:
			if err := session.doTargetTooHigh(TypedError); err != nil {
				return session.handleError(err)
			}
		case resendState:
			//assumes target too high reject already sent
		}

		session.messageStash[TypedError.ReceivedTarget] = msg
		return resendState{}

	case targetTooLow:
		return state.doTargetTooLow(session, msg, TypedError)
	case incorrectBeginString:
		if err := session.initiateLogout(rej.Error()); err != nil {
			session.handleError(err)
		}
		return logoutState{}
	}

	switch rej.RejectReason() {
	case rejectReasonCompIDProblem, rejectReasonSendingTimeAccuracyProblem:
		if err := session.doReject(msg, rej); err != nil {
			return session.handleError(err)
		}

		if err := session.initiateLogout(""); err != nil {
			return session.handleError(err)
		}
		return logoutState{}
	default:
		if err := session.doReject(msg, rej); err != nil {
			return session.handleError(err)
		}

		if err := session.store.IncrNextTargetMsgSeqNum(); err != nil {
			return session.handleError(err)
		}
		return state
	}
}

func (state inSession) doTargetTooLow(session *session, msg Message, rej targetTooLow) (nextState sessionState) {
	var posDupFlag FIXBoolean
	if err := msg.Header.GetField(tagPossDupFlag, &posDupFlag); err == nil && posDupFlag {

		origSendingTime := new(FIXUTCTimestamp)
		if err = msg.Header.GetField(tagOrigSendingTime, origSendingTime); err != nil {
			if rejErr := session.doReject(msg, RequiredTagMissing(tagOrigSendingTime)); rejErr != nil {
				return session.handleError(rejErr)
			}
			return state
		}

		sendingTime := new(FIXUTCTimestamp)
		msg.Header.GetField(tagSendingTime, sendingTime)

		if sendingTime.Before(origSendingTime.Time) {
			if err := session.doReject(msg, sendingTimeAccuracyProblem()); err != nil {
				return session.handleError(err)
			}

			if err := session.initiateLogout(""); err != nil {
				return session.handleError(err)
			}
			return logoutState{}
		}

		if appReject := session.fromCallback(msg); appReject != nil {
			if err := session.doReject(msg, appReject); err != nil {
				return session.handleError(err)
			}

			if err := session.initiateLogout(""); err != nil {
				return session.handleError(err)
			}
			return logoutState{}
		}
	} else {
		if err := session.initiateLogout(rej.Error()); err != nil {
			return session.handleError(err)
		}
		return logoutState{}
	}

	return state
}

func (state *inSession) generateSequenceReset(session *session, beginSeqNo int, endSeqNo int) {
	sequenceReset := NewMessage()
	session.fillDefaultHeader(sequenceReset)

	sequenceReset.Header.SetField(tagMsgType, FIXString("4"))
	sequenceReset.Header.SetField(tagMsgSeqNum, FIXInt(beginSeqNo))
	sequenceReset.Header.SetField(tagPossDupFlag, FIXBoolean(true))
	sequenceReset.Body.SetField(tagNewSeqNo, FIXInt(endSeqNo))
	sequenceReset.Body.SetField(tagGapFillFlag, FIXBoolean(true))

	var origSendingTime FIXString
	if err := sequenceReset.Header.GetField(tagSendingTime, &origSendingTime); err == nil {
		sequenceReset.Header.SetField(tagOrigSendingTime, origSendingTime)
	}

	//FIXME error check?
	msgBytes, _ := sequenceReset.Build()
	session.sendBytes(msgBytes)

	session.log.OnEventf("Sent SequenceReset TO: %v", endSeqNo)
}
