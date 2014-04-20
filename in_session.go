package quickfix

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/fix/tag"
	"github.com/quickfixgo/quickfix/message"
	"time"
)

type inSession struct {
}

func (state inSession) FixMsgIn(session *session, msg message.Message) (nextState sessionState) {
	var msgType field.MsgType
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

func (state inSession) Timeout(session *session, event event) (nextState sessionState) {
	switch event {
	case needHeartbeat:
		heartBt := message.CreateMessageBuilder()
		heartBt.Header.Set(field.BuildMsgType("0"))
		session.send(heartBt)
	case peerTimeout:
		testReq := message.CreateMessageBuilder()
		testReq.Header.Set(field.BuildMsgType("1"))
		testReq.Body.Set(field.BuildTestReqID("TEST"))
		session.send(testReq)
		session.peerTimer.Reset(time.Duration(int64(1.2 * float64(session.heartBeatTimeout))))
		return pendingTimeout{}
	}
	return state
}

func (state inSession) handleLogon(session *session, msg message.Message) (nextState sessionState) {
	if err := session.handleLogon(msg); err != nil {
		return state.initiateLogout(session, "")
	}

	return state
}

func (state inSession) handleLogout(session *session, msg message.Message) (nextState sessionState) {
	session.log.OnEvent("Received logout request")
	state.generateLogout(session)
	session.application.OnLogout(session.SessionID)

	return latentState{}
}

func (state inSession) handleSequenceReset(session *session, msg message.Message) (nextState sessionState) {
	gapFillFlag := new(message.BooleanField)
	msg.Body.GetField(tag.GapFillFlag, gapFillFlag)

	if err := session.verifySelect(msg, gapFillFlag.Value, gapFillFlag.Value); err != nil {
		return state.processReject(session, msg, err)
	}

	newSeqNo := new(message.IntField)
	if err := msg.Body.GetField(tag.NewSeqNo, newSeqNo); err == nil {
		session.log.OnEventf("Received SequenceReset FROM: %v TO: %v", session.expectedSeqNum, newSeqNo.Value)

		switch {
		case newSeqNo.Value > session.expectedSeqNum:
			session.expectedSeqNum = newSeqNo.Value
		case newSeqNo.Value < session.expectedSeqNum:
			//FIXME: to be compliant with legacy tests, do not include tag in reftagid? (11c_NewSeqNoLess)
			session.doReject(msg, errors.ValueIsIncorrectNoTag())
		}
	}

	return state
}

func (state inSession) handleResendRequest(session *session, msg message.Message) (nextState sessionState) {
	if err := session.verifyIgnoreSeqNumTooHighOrLow(msg); err != nil {
		return state.processReject(session, msg, err)
	}

	var err error
	beginSeqNoField := new(message.IntValue)
	if err = msg.Body.GetField(tag.BeginSeqNo, beginSeqNoField); err != nil {
		return state.processReject(session, msg, errors.RequiredTagMissing(tag.BeginSeqNo))
	}

	beginSeqNo := beginSeqNoField.Value

	endSeqNoField := new(message.IntField)
	if err = msg.Body.GetField(tag.EndSeqNo, endSeqNoField); err != nil {
		return state.processReject(session, msg, errors.RequiredTagMissing(tag.EndSeqNo))
	}

	endSeqNo := endSeqNoField.Value

	session.log.OnEventf("Received ResendRequest FROM: %d TO: %d", beginSeqNo, endSeqNo)

	if (session.BeginString >= fix.BeginString_FIX42 && endSeqNo == 0) ||
		(session.BeginString <= fix.BeginString_FIX42 && endSeqNo == 999999) ||
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

		msg, _ := message.MessageFromParsedBytes(buf.Bytes())

		msgType := new(message.StringValue)
		msg.Header.GetField(tag.MsgType, msgType)

		sentMessageSeqNum := new(message.IntValue)
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

func (state inSession) handleTestRequest(session *session, msg message.Message) (nextState sessionState) {
	if err := session.verify(msg); err != nil {
		return state.processReject(session, msg, err)
	}

	var testReq field.TestReqID
	if err := msg.Body.Get(&testReq); err != nil {
		session.log.OnEvent("Test Request with no testRequestID")
	} else {
		heartBt := message.CreateMessageBuilder()
		heartBt.Header.Set(field.BuildMsgType("0"))
		heartBt.Body.Set(field.BuildTestReqID(testReq.Value))
		session.send(heartBt)
	}

	session.expectedSeqNum++

	return state
}

func (state inSession) processReject(session *session, msg message.Message, rej errors.MessageRejectError) (nextState sessionState) {
	switch TypedError := rej.(type) {
	case errors.TargetTooHigh:

		switch session.currentState.(type) {
		default:
			session.DoTargetTooHigh(TypedError)
		case resendState:
			//assumes target too high reject already sent
		}
		session.messageStash[TypedError.ReceivedTarget] = msg
		return resendState{}

	case errors.TargetTooLow:
		return state.doTargetTooLow(session, msg, TypedError)
	case errors.IncorrectBeginString:
		return state.initiateLogout(session, rej.Error())
	}

	switch rej.RejectReason() {
	case errors.RejectReasonCompIDProblem, errors.RejectReasonSendingTimeAccuracyProblem:
		session.doReject(msg, rej)
		return state.initiateLogout(session, "")
	default:
		session.doReject(msg, rej)
		session.expectedSeqNum++
		return state
	}
}

func (state inSession) doTargetTooLow(session *session, msg message.Message, rej errors.TargetTooLow) (nextState sessionState) {
	posDupFlag := new(message.BooleanValue)
	if err := msg.Header.GetField(tag.PossDupFlag, posDupFlag); err == nil && posDupFlag.Value {

		origSendingTime := new(message.UTCTimestampValue)
		if err = msg.Header.GetField(tag.OrigSendingTime, origSendingTime); err != nil {
			session.doReject(msg, errors.RequiredTagMissing(tag.OrigSendingTime))
			return state
		}

		sendingTime := new(message.UTCTimestampValue)
		msg.Header.GetField(tag.SendingTime, sendingTime)

		if sendingTime.Value.Before(origSendingTime.Value) {
			session.doReject(msg, errors.SendingTimeAccuracyProblem())
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

func (state *inSession) initiateLogout(session *session, reason string) (nextState logoutState) {
	state.generateLogoutWithReason(session, reason)
	time.AfterFunc(time.Duration(2)*time.Second, func() { session.sessionEvent <- logoutTimeout })

	return
}

func (state *inSession) generateSequenceReset(session *session, beginSeqNo int, endSeqNo int) {
	sequenceReset := message.CreateMessageBuilder()
	session.fillDefaultHeader(sequenceReset)

	sequenceReset.Header.Set(field.BuildMsgType("4"))
	sequenceReset.Header.Set(field.BuildMsgSeqNum(beginSeqNo))
	sequenceReset.Header.Set(field.BuildPossDupFlag(true))
	sequenceReset.Body.Set(field.BuildNewSeqNo(endSeqNo))
	sequenceReset.Body.Set(field.BuildGapFillFlag(true))

	origSendingTime := new(message.StringValue)
	if err := sequenceReset.Header.GetField(tag.SendingTime, origSendingTime); err == nil {
		sequenceReset.Header.Set(message.NewStringField(tag.OrigSendingTime, origSendingTime.Value))
	}

	//FIXME error check?
	buffer, _ := sequenceReset.Build()
	session.sendBuffer(buffer)
}

func (state *inSession) generateLogout(session *session) {
	state.generateLogoutWithReason(session, "")
}

func (state *inSession) generateLogoutWithReason(session *session, reason string) {
	reply := message.CreateMessageBuilder()
	reply.Header.Set(field.BuildMsgType("5"))
	reply.Header.Set(field.BuildBeginString(session.BeginString))
	reply.Header.Set(field.BuildTargetCompID(session.TargetCompID))
	reply.Header.Set(field.BuildSenderCompID(session.SenderCompID))

	if reason != "" {
		reply.Body.Set(field.BuildText(reason))
	}

	session.send(reply)
	session.log.OnEvent("Sending logout response")
}
