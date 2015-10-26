package quickfix

import (
	"github.com/quickfixgo/quickfix/fix/enum"
	"time"
)

type inSession struct {
}

func (state inSession) FixMsgIn(session *Session, msg Message) (nextState sessionState) {
	msgType := new(StringValue)
	if err := msg.Header.GetField(tagMsgType, msgType); err == nil {
		switch msgType.Value {
		//logon
		case "A":
			return state.handleLogon(session, msg)
		//logout
		case "5":
			return state.handleLogout(session, msg)
		//test request
		case "1":
			return state.handleTestRequest(session, msg)
		//resend request
		case "2":
			return state.handleResendRequest(session, msg)
		//sequence reset
		case "4":
			return state.handleSequenceReset(session, msg)
		default:
			if err := session.verify(msg); err != nil {
				return state.processReject(session, msg, err)
			}
		}
	}

	session.store.IncrNextTargetMsgSeqNum()

	return state
}

func (state inSession) Timeout(session *Session, event event) (nextState sessionState) {
	switch event {
	case needHeartbeat:
		heartBt := NewMessage()
		heartBt.Header.SetField(tagMsgType, &StringValue{Value: "0"})
		session.send(heartBt)
	case peerTimeout:
		testReq := NewMessage()
		testReq.Header.SetField(tagMsgType, &StringValue{Value: "1"})
		testReq.Body.SetField(tagTestReqID, &StringValue{Value: "TEST"})
		session.send(testReq)
		session.peerTimer.Reset(time.Duration(int64(1.2 * float64(session.heartBeatTimeout))))
		return pendingTimeout{}
	}
	return state
}

func (state inSession) handleLogon(session *Session, msg Message) (nextState sessionState) {
	if err := session.handleLogon(msg); err != nil {
		return state.initiateLogout(session, "")
	}

	return state
}

func (state inSession) handleLogout(session *Session, msg Message) (nextState sessionState) {
	session.log.OnEvent("Received logout request")
	state.generateLogout(session)
	session.application.OnLogout(session.sessionID)

	return latentState{}
}

func (state inSession) handleSequenceReset(session *Session, msg Message) (nextState sessionState) {
	gapFillFlag := new(BooleanField)
	msg.Body.GetField(tagGapFillFlag, gapFillFlag)

	if err := session.verifySelect(msg, gapFillFlag.Value, gapFillFlag.Value); err != nil {
		return state.processReject(session, msg, err)
	}

	newSeqNo := new(IntField)
	if err := msg.Body.GetField(tagNewSeqNo, newSeqNo); err == nil {
		expectedSeqNum := session.store.NextTargetMsgSeqNum()
		session.log.OnEventf("Received SequenceReset FROM: %v TO: %v", expectedSeqNum, newSeqNo.Value)

		switch {
		case newSeqNo.Value > expectedSeqNum:
			session.store.SetNextTargetMsgSeqNum(newSeqNo.Value)
		case newSeqNo.Value < expectedSeqNum:
			//FIXME: to be compliant with legacy tests, do not include tag in reftagid? (11c_NewSeqNoLess)
			session.doReject(msg, valueIsIncorrectNoTag())
		}
	}

	return state
}

func (state inSession) handleResendRequest(session *Session, msg Message) (nextState sessionState) {
	if err := session.verifyIgnoreSeqNumTooHighOrLow(msg); err != nil {
		return state.processReject(session, msg, err)
	}

	var err error
	beginSeqNoField := new(IntValue)
	if err = msg.Body.GetField(tagBeginSeqNo, beginSeqNoField); err != nil {
		return state.processReject(session, msg, requiredTagMissing(tagBeginSeqNo))
	}

	beginSeqNo := beginSeqNoField.Value

	endSeqNoField := new(IntField)
	if err = msg.Body.GetField(tagEndSeqNo, endSeqNoField); err != nil {
		return state.processReject(session, msg, requiredTagMissing(tagEndSeqNo))
	}

	endSeqNo := endSeqNoField.Value

	session.log.OnEventf("Received ResendRequest FROM: %d TO: %d", beginSeqNo, endSeqNo)
	expectedSeqNum := session.store.NextTargetMsgSeqNum()

	if (session.sessionID.BeginString >= enum.BeginStringFIX42 && endSeqNo == 0) ||
		(session.sessionID.BeginString <= enum.BeginStringFIX42 && endSeqNo == 999999) ||
		(endSeqNo >= expectedSeqNum) {
		endSeqNo = expectedSeqNum - 1
	}

	state.resendMessages(session, beginSeqNo, endSeqNo)
	session.store.IncrNextTargetMsgSeqNum()
	return state
}

func (state inSession) resendMessages(session *Session, beginSeqNo, endSeqNo int) {
	msgs := session.store.GetMessages(beginSeqNo, endSeqNo)

	seqNum := beginSeqNo
	nextSeqNum := seqNum

	var msgBytes []byte
	var ok bool
	for {
		if msgBytes, ok = <-msgs; !ok {
			//gapfill for catch-up
			if seqNum != nextSeqNum {
				state.generateSequenceReset(session, seqNum, nextSeqNum)
			}

			return
		}

		msg, _ := parseMessage(msgBytes)

		msgType := new(StringValue)
		msg.Header.GetField(tagMsgType, msgType)

		sentMessageSeqNum := new(IntValue)
		msg.Header.GetField(tagMsgSeqNum, sentMessageSeqNum)

		if isAdminMessageType(msgType.Value) {
			nextSeqNum = sentMessageSeqNum.Value + 1
		} else {

			if seqNum != sentMessageSeqNum.Value {
				state.generateSequenceReset(session, seqNum, sentMessageSeqNum.Value)
			}

			session.resend(msg)
			seqNum = sentMessageSeqNum.Value + 1
			nextSeqNum = seqNum
		}
	}
}

func (state inSession) handleTestRequest(session *Session, msg Message) (nextState sessionState) {
	if err := session.verify(msg); err != nil {
		return state.processReject(session, msg, err)
	}

	testReq := new(StringValue)
	if err := msg.Body.GetField(tagTestReqID, testReq); err != nil {
		session.log.OnEvent("Test Request with no testRequestID")
	} else {
		heartBt := NewMessage()
		heartBt.Header.SetField(tagMsgType, &StringValue{Value: "0"})
		heartBt.Body.SetField(tagTestReqID, testReq)
		session.send(heartBt)
	}

	session.store.IncrNextTargetMsgSeqNum()

	return state
}

func (state inSession) processReject(session *Session, msg Message, rej MessageRejectError) (nextState sessionState) {
	switch TypedError := rej.(type) {
	case targetTooHigh:

		switch session.currentState.(type) {
		default:
			session.doTargetTooHigh(TypedError)
		case resendState:
			//assumes target too high reject already sent
		}
		session.messageStash[TypedError.ReceivedTarget] = msg
		return resendState{}

	case targetTooLow:
		return state.doTargetTooLow(session, msg, TypedError)
	case incorrectBeginString:
		return state.initiateLogout(session, rej.Error())
	}

	switch rej.RejectReason() {
	case rejectReasonCompIDProblem, rejectReasonSendingTimeAccuracyProblem:
		session.doReject(msg, rej)
		return state.initiateLogout(session, "")
	default:
		session.doReject(msg, rej)
		session.store.IncrNextTargetMsgSeqNum()
		return state
	}
}

func (state inSession) doTargetTooLow(session *Session, msg Message, rej targetTooLow) (nextState sessionState) {
	posDupFlag := new(BooleanValue)
	if err := msg.Header.GetField(tagPossDupFlag, posDupFlag); err == nil && posDupFlag.Value {

		origSendingTime := new(UTCTimestampValue)
		if err = msg.Header.GetField(tagOrigSendingTime, origSendingTime); err != nil {
			session.doReject(msg, requiredTagMissing(tagOrigSendingTime))
			return state
		}

		sendingTime := new(UTCTimestampValue)
		msg.Header.GetField(tagSendingTime, sendingTime)

		if sendingTime.Value.Before(origSendingTime.Value) {
			session.doReject(msg, sendingTimeAccuracyProblem())
			return state.initiateLogout(session, "")
		}

		if appReject := session.fromCallback(msg); appReject != nil {
			session.doReject(msg, appReject)
			return state.initiateLogout(session, "")
		}
	} else {
		return state.initiateLogout(session, rej.Error())
	}

	return state
}

func (state *inSession) initiateLogout(session *Session, reason string) (nextState logoutState) {
	state.generateLogoutWithReason(session, reason)
	time.AfterFunc(time.Duration(2)*time.Second, func() { session.sessionEvent <- logoutTimeout })

	return
}

func (state *inSession) generateSequenceReset(session *Session, beginSeqNo int, endSeqNo int) {
	sequenceReset := NewMessage()
	session.fillDefaultHeader(sequenceReset)

	sequenceReset.Header.Set(NewStringField(tagMsgType, "4"))
	sequenceReset.Header.Set(NewIntField(tagMsgSeqNum, beginSeqNo))
	sequenceReset.Header.Set(NewBooleanField(tagPossDupFlag, true))
	sequenceReset.Body.Set(NewIntField(tagNewSeqNo, endSeqNo))
	sequenceReset.Body.Set(NewBooleanField(tagGapFillFlag, true))

	origSendingTime := new(StringValue)
	if err := sequenceReset.Header.GetField(tagSendingTime, origSendingTime); err == nil {
		sequenceReset.Header.Set(NewStringField(tagOrigSendingTime, origSendingTime.Value))
	}

	//FIXME error check?
	msgBytes, _ := sequenceReset.Build()
	session.sendBytes(msgBytes)
}

func (state *inSession) generateLogout(session *Session) {
	state.generateLogoutWithReason(session, "")
}

func (state *inSession) generateLogoutWithReason(session *Session, reason string) {
	reply := NewMessage()
	reply.Header.Set(NewStringField(tagMsgType, "5"))
	reply.Header.Set(NewStringField(tagBeginString, session.sessionID.BeginString))
	reply.Header.Set(NewStringField(tagTargetCompID, session.sessionID.TargetCompID))
	reply.Header.Set(NewStringField(tagSenderCompID, session.sessionID.SenderCompID))

	if reason != "" {
		reply.Body.Set(NewStringField(tagText, reason))
	}

	session.send(reply)
	session.log.OnEvent("Sending logout response")
}
