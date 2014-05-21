package quickfix

import (
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/fix/tag"
	"time"
)

type inSession struct {
}

func (state inSession) FixMsgIn(session *Session, msg Message) (nextState sessionState) {
	msgType := field.MsgTypeField{}
	if err := msg.Header.Get(&msgType); err == nil {
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

	session.expectedSeqNum++

	return state
}

func (state inSession) Timeout(session *Session, event event) (nextState sessionState) {
	switch event {
	case needHeartbeat:
		heartBt := NewMessageBuilder()
		heartBt.Header().Set(field.NewMsgType("0"))
		session.send(heartBt)
	case peerTimeout:
		testReq := NewMessageBuilder()
		testReq.Header().Set(field.NewMsgType("1"))
		testReq.Body().Set(field.NewTestReqID("TEST"))
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
	gapFillFlag := new(fix.BooleanField)
	msg.Body.GetField(tag.GapFillFlag, gapFillFlag)

	if err := session.verifySelect(msg, gapFillFlag.Value, gapFillFlag.Value); err != nil {
		return state.processReject(session, msg, err)
	}

	newSeqNo := new(fix.IntField)
	if err := msg.Body.GetField(tag.NewSeqNo, newSeqNo); err == nil {
		session.log.OnEventf("Received SequenceReset FROM: %v TO: %v", session.expectedSeqNum, newSeqNo.Value)

		switch {
		case newSeqNo.Value > session.expectedSeqNum:
			session.expectedSeqNum = newSeqNo.Value
		case newSeqNo.Value < session.expectedSeqNum:
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
	beginSeqNoField := new(fix.IntValue)
	if err = msg.Body.GetField(tag.BeginSeqNo, beginSeqNoField); err != nil {
		return state.processReject(session, msg, requiredTagMissing(tag.BeginSeqNo))
	}

	beginSeqNo := beginSeqNoField.Value

	endSeqNoField := new(fix.IntField)
	if err = msg.Body.GetField(tag.EndSeqNo, endSeqNoField); err != nil {
		return state.processReject(session, msg, requiredTagMissing(tag.EndSeqNo))
	}

	endSeqNo := endSeqNoField.Value

	session.log.OnEventf("Received ResendRequest FROM: %d TO: %d", beginSeqNo, endSeqNo)

	if (session.sessionID.BeginString >= fix.BeginString_FIX42 && endSeqNo == 0) ||
		(session.sessionID.BeginString <= fix.BeginString_FIX42 && endSeqNo == 999999) ||
		(endSeqNo >= session.expectedSeqNum) {
		endSeqNo = session.expectedSeqNum - 1
	}

	state.resendMessages(session, beginSeqNo, endSeqNo)
	session.expectedSeqNum++
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

		msgType := new(fix.StringValue)
		msg.Header.GetField(tag.MsgType, msgType)

		sentMessageSeqNum := new(fix.IntValue)
		msg.Header.GetField(tag.MsgSeqNum, sentMessageSeqNum)

		if fix.IsAdminMessageType(msgType.Value) {
			nextSeqNum = sentMessageSeqNum.Value + 1
		} else {

			if seqNum != sentMessageSeqNum.Value {
				state.generateSequenceReset(session, seqNum, sentMessageSeqNum.Value)
			}

			session.resend(msg.ToBuilder())
			seqNum = sentMessageSeqNum.Value + 1
			nextSeqNum = seqNum
		}
	}
}

func (state inSession) handleTestRequest(session *Session, msg Message) (nextState sessionState) {
	if err := session.verify(msg); err != nil {
		return state.processReject(session, msg, err)
	}

	var testReq field.TestReqIDField
	if err := msg.Body.Get(&testReq); err != nil {
		session.log.OnEvent("Test Request with no testRequestID")
	} else {
		heartBt := NewMessageBuilder()
		heartBt.Header().Set(field.NewMsgType("0"))
		heartBt.Body().Set(field.NewTestReqID(testReq.Value))
		session.send(heartBt)
	}

	session.expectedSeqNum++

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
		session.expectedSeqNum++
		return state
	}
}

func (state inSession) doTargetTooLow(session *Session, msg Message, rej targetTooLow) (nextState sessionState) {
	posDupFlag := new(fix.BooleanValue)
	if err := msg.Header.GetField(tag.PossDupFlag, posDupFlag); err == nil && posDupFlag.Value {

		origSendingTime := new(fix.UTCTimestampValue)
		if err = msg.Header.GetField(tag.OrigSendingTime, origSendingTime); err != nil {
			session.doReject(msg, requiredTagMissing(tag.OrigSendingTime))
			return state
		}

		sendingTime := new(fix.UTCTimestampValue)
		msg.Header.GetField(tag.SendingTime, sendingTime)

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
	sequenceReset := NewMessageBuilder()
	session.fillDefaultHeader(sequenceReset)

	sequenceReset.Header().Set(field.NewMsgType("4"))
	sequenceReset.Header().Set(field.NewMsgSeqNum(beginSeqNo))
	sequenceReset.Header().Set(field.NewPossDupFlag(true))
	sequenceReset.Body().Set(field.NewNewSeqNo(endSeqNo))
	sequenceReset.Body().Set(field.NewGapFillFlag(true))

	origSendingTime := new(fix.StringValue)
	if err := sequenceReset.Header().GetField(tag.SendingTime, origSendingTime); err == nil {
		sequenceReset.Header().Set(fix.NewStringField(tag.OrigSendingTime, origSendingTime.Value))
	}

	//FIXME error check?
	msg, _ := sequenceReset.Build()
	session.sendBytes(msg.Bytes)
}

func (state *inSession) generateLogout(session *Session) {
	state.generateLogoutWithReason(session, "")
}

func (state *inSession) generateLogoutWithReason(session *Session, reason string) {
	reply := NewMessageBuilder()
	reply.Header().Set(field.NewMsgType("5"))
	reply.Header().Set(field.NewBeginString(session.sessionID.BeginString))
	reply.Header().Set(field.NewTargetCompID(session.sessionID.TargetCompID))
	reply.Header().Set(field.NewSenderCompID(session.sessionID.SenderCompID))

	if reason != "" {
		reply.Body().Set(field.NewText(reason))
	}

	session.send(reply)
	session.log.OnEvent("Sending logout response")
}
