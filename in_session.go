package quickfix

import (
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/tag"
	"time"
)

type inSession struct {
}

func (state inSession) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	msgType := new(field.StringField)
	if err := msg.Header.GetField(tag.MsgType, msgType); err == nil {
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
				return state.processReject(session, err)
			}
		}
	}

	session.expectedSeqNum++

	return state
}

func (state inSession) Timeout(session *session, event event) (nextState sessionState) {
	switch event {
	case needHeartbeat:
		heartBt := NewMessageBuilder()
		heartBt.Header.Set(field.NewStringField(tag.MsgType, "0"))
		session.send(heartBt)
	case peerTimeout:
		testReq := NewMessageBuilder()
		testReq.Header.Set(field.NewStringField(tag.MsgType, "1"))
		testReq.Body.Set(field.NewStringField(tag.TestReqID, "TEST"))
		session.send(testReq)
		session.peerTimer.Reset(time.Duration(int64(1.2 * float64(session.heartBeatTimeout))))
		return pendingTimeout{}
	}
	return state
}

func (state inSession) handleLogon(session *session, msg Message) (nextState sessionState) {
	if err := session.handleLogon(msg); err != nil {
		return state.initiateLogout(session, "")
	}

	return state
}

func (state inSession) handleLogout(session *session, msg Message) (nextState sessionState) {
	session.log.OnEvent("Received logout request")
	state.generateLogout(session)
	session.application.OnLogout(session.SessionID)

	return latentState{}
}

func (state inSession) handleSequenceReset(session *session, msg Message) (nextState sessionState) {
	gapFillFlag := new(field.BooleanField)
	msg.Body.GetField(tag.GapFillFlag, gapFillFlag)

	if err := session.verifySelect(msg, gapFillFlag.Value, gapFillFlag.Value); err != nil {
		return state.processReject(session, err)
	}

	newSeqNo := new(field.IntField)
	if err := msg.Body.GetField(tag.NewSeqNo, newSeqNo); err == nil {
		session.log.OnEventf("Received SequenceReset FROM: %v TO: %v", session.expectedSeqNum, newSeqNo.Value)

		switch {
		case newSeqNo.Value > session.expectedSeqNum:
			session.expectedSeqNum = newSeqNo.Value
		case newSeqNo.Value < session.expectedSeqNum:
			//FIXME: to be compliant with legacy tests, do not include tag in reftagid? (11c_NewSeqNoLess)
			session.doReject(newValueIsIncorrect(msg, nil))
		}
	}

	return state
}

func (state inSession) handleResendRequest(session *session, msg Message) (nextState sessionState) {
	if err := session.verifyIgnoreSeqNumTooHighOrLow(msg); err != nil {
		return state.processReject(session, err)
	}

	var err error
	beginSeqNoField := new(field.IntValue)
	if err = msg.Body.GetField(tag.BeginSeqNo, beginSeqNoField); err != nil {
		return state.processReject(session, newRequiredTagMissing(msg, tag.BeginSeqNo))
	}

	beginSeqNo := beginSeqNoField.Value

	endSeqNoField := new(field.IntField)
	if err = msg.Body.GetField(tag.EndSeqNo, endSeqNoField); err != nil {
		return state.processReject(session, newRequiredTagMissing(msg, tag.EndSeqNo))
	}

	endSeqNo := endSeqNoField.Value

	session.log.OnEventf("Received ResendRequest FROM: %d TO: %d", beginSeqNo, endSeqNo)

	if (session.BeginString >= BeginString_FIX42 && endSeqNo == 0) ||
		(session.BeginString <= BeginString_FIX42 && endSeqNo == 999999) ||
		(endSeqNo >= session.expectedSeqNum) {
		endSeqNo = session.expectedSeqNum - 1
	}

	state.resendMessages(session, beginSeqNo, endSeqNo)
	session.expectedSeqNum++
	return state
}

func (state inSession) resendMessages(session *session, beginSeqNo, endSeqNo int) {
	buffers := session.store.GetMessages(beginSeqNo, endSeqNo)

	seqNum := beginSeqNo
	nextSeqNum := seqNum

	var buf buffer
	var ok bool
	for {
		if buf, ok = <-buffers; !ok {
			//gapfill for catch-up
			if seqNum != nextSeqNum {
				state.generateSequenceReset(session, seqNum, nextSeqNum)
			}

			return
		}

		message, _ := MessageFromParsedBytes(buf.Bytes())

		msgType := new(field.StringValue)
		message.Header.GetField(tag.MsgType, msgType)

		sentMessageSeqNum := new(field.IntValue)
		message.Header.GetField(tag.MsgSeqNum, sentMessageSeqNum)

		if IsAdminMessageType(msgType.Value) {
			nextSeqNum = sentMessageSeqNum.Value + 1
		} else {

			if seqNum != sentMessageSeqNum.Value {
				state.generateSequenceReset(session, seqNum, sentMessageSeqNum.Value)
			}

			session.resend(message.ToBuilder())
			seqNum = sentMessageSeqNum.Value + 1
			nextSeqNum = seqNum
		}
	}
}

func (state inSession) handleTestRequest(session *session, msg Message) (nextState sessionState) {
	if err := session.verify(msg); err != nil {
		return state.processReject(session, err)
	}

	testReq := new(field.StringValue)
	if err := msg.Body.GetField(tag.TestReqID, testReq); err != nil {
		session.log.OnEvent("Test Request with no testRequestID")
	} else {
		heartBt := NewMessageBuilder()
		heartBt.Header.Set(field.NewStringField(tag.MsgType, "0"))
		heartBt.Body.Set(field.NewStringField(tag.TestReqID, testReq.Value))
		session.send(heartBt)
	}

	session.expectedSeqNum++

	return state
}

func (state inSession) processReject(session *session, rej MessageReject) (nextState sessionState) {
	switch TypedError := rej.(type) {
	case targetTooHigh:

		switch session.currentState.(type) {
		default:
			session.DoTargetTooHigh(TypedError)
		case resendState:
			//assumes target too high reject already sent
		}
		session.messageStash[TypedError.ReceivedTarget] = rej.RejectedMessage()
		return resendState{}

	case targetTooLow:
		return state.doTargetTooLow(session, TypedError)
	case incorrectBeginString:
		return state.initiateLogout(session, rej.Error())
	}

	switch rej.RejectReason() {
	case CompIDProblem, SendingTimeAccuracyProblem:
		session.doReject(rej)
		return state.initiateLogout(session, "")
	default:
		session.doReject(rej)
		session.expectedSeqNum++
		return state
	}
}

func (state inSession) doTargetTooLow(session *session, rej targetTooLow) (nextState sessionState) {
	posDupFlag := new(field.BooleanValue)
	if err := rej.RejectedMessage().Header.GetField(tag.PossDupFlag, posDupFlag); err == nil && posDupFlag.Value {

		origSendingTime := new(field.UTCTimestampValue)
		if err = rej.RejectedMessage().Header.GetField(tag.OrigSendingTime, origSendingTime); err != nil {
			session.doReject(newRequiredTagMissing(rej.RejectedMessage(), tag.OrigSendingTime))
			return state
		}

		sendingTime := new(field.UTCTimestampValue)
		rej.RejectedMessage().Header.GetField(tag.SendingTime, sendingTime)

		if sendingTime.Value.Before(origSendingTime.Value) {
			session.doReject(newSendingTimeAccuracyProblem(rej.RejectedMessage()))
			return state.initiateLogout(session, "")
		}

		if appReject := session.fromCallback(rej.RejectedMessage()); appReject != nil {
			session.doReject(appReject)
			return state.initiateLogout(session, "")
		}
	} else {
		return state.initiateLogout(session, rej.Error())
	}

	return state
}

func (state *inSession) initiateLogout(session *session, reason string) (nextState logoutState) {
	state.generateLogoutWithReason(session, reason)
	time.AfterFunc(time.Duration(2)*time.Second, func() { session.sessionEvent <- logoutTimeout })

	return
}

func (state *inSession) generateSequenceReset(session *session, beginSeqNo int, endSeqNo int) {
	sequenceReset := NewMessageBuilder()
	session.fillDefaultHeader(sequenceReset)

	sequenceReset.Header.Set(field.NewStringField(tag.MsgType, "4"))
	sequenceReset.Header.Set(field.NewIntField(tag.MsgSeqNum, beginSeqNo))
	sequenceReset.Header.Set(field.NewBooleanField(tag.PossDupFlag, true))
	sequenceReset.Body.Set(field.NewIntField(tag.NewSeqNo, endSeqNo))
	sequenceReset.Body.Set(field.NewBooleanField(tag.GapFillFlag, true))

	origSendingTime := new(field.StringValue)
	if err := sequenceReset.Header.GetField(tag.SendingTime, origSendingTime); err == nil {
		sequenceReset.Header.Set(field.NewStringField(tag.OrigSendingTime, origSendingTime.Value))
	}

	//FIXME error check?
	buffer, _ := sequenceReset.Build()
	session.sendBuffer(buffer)
}

func (state *inSession) generateLogout(session *session) {
	state.generateLogoutWithReason(session, "")
}

func (state *inSession) generateLogoutWithReason(session *session, reason string) {
	reply := NewMessageBuilder()
	reply.Header.Set(field.NewStringField(tag.MsgType, "5"))
	reply.Header.Set(field.NewStringField(tag.BeginString, session.BeginString))
	reply.Header.Set(field.NewStringField(tag.TargetCompID, session.TargetCompID))
	reply.Header.Set(field.NewStringField(tag.SenderCompID, session.SenderCompID))

	if reason != "" {
		reply.Body.Set(field.NewStringField(tag.Text, reason))
	}

	session.send(reply)
	session.log.OnEvent("Sending logout response")
}
