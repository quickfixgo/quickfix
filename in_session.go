package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
	"time"
)

type inSession struct {
}

func (state inSession) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	msgType := new(StringField)
	if err := msg.Header.GetField(tag.MsgType, msgType); err == nil {
		switch msgType.Value {
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
		heartBt.Header.Set(NewStringField(tag.MsgType, "0"))
		session.send(heartBt)
	case peerTimeout:
		testReq := NewMessageBuilder()
		testReq.Header.Set(NewStringField(tag.MsgType, "1"))
		testReq.Body.Set(NewStringField(tag.TestReqID, "TEST"))
		session.send(testReq)
		session.peerTimer.Reset(time.Duration(int64(1.2 * float64(session.heartBeatTimeout))))
		return pendingTimeout{}
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
	gapFillFlag := new(BooleanField)
	msg.Body.GetField(tag.GapFillFlag, gapFillFlag)

	if err := session.verifySelect(msg, gapFillFlag.Value, gapFillFlag.Value); err != nil {
		return state.processReject(session, err)
	}

	newSeqNo := new(IntField)
	if err := msg.Body.GetField(tag.NewSeqNo, newSeqNo); err == nil {
		session.log.OnEventf("Received SequenceReset FROM: %v TO: %v", session.expectedSeqNum, newSeqNo.Value)

		switch {
		case newSeqNo.Value > session.expectedSeqNum:
			session.expectedSeqNum = newSeqNo.Value
		case newSeqNo.Value < session.expectedSeqNum:
			//FIXME: to be compliant with legacy tests, do not include tag in reftagid? (11c_NewSeqNoLess)
			session.doReject(NewValueIsIncorrect(msg, nil))
		}
	}

	return state
}

func (state inSession) handleResendRequest(session *session, msg Message) (nextState sessionState) {
	if err := session.verifyIgnoreSeqNumTooHighOrLow(msg); err != nil {
		return state.processReject(session, err)
	}

	var beginSeqNo, endSeqNo int
	beginSeqNoField := new(IntValue)
	if err := msg.Body.GetField(tag.BeginSeqNo, beginSeqNoField); err != nil {
		return state.processReject(session, NewRequiredTagMissing(msg, tag.BeginSeqNo))
	} else {
		beginSeqNo = beginSeqNoField.Value
	}

	endSeqNoField := new(IntField)
	if err := msg.Body.GetField(tag.EndSeqNo, endSeqNoField); err != nil {
		return state.processReject(session, NewRequiredTagMissing(msg, tag.EndSeqNo))
	} else {
		endSeqNo = endSeqNoField.Value
	}

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

	for {
		if buf, ok := <-buffers; !ok {
			//gapfill for catch-up
			if seqNum != nextSeqNum {
				state.generateSequenceReset(session, seqNum, nextSeqNum)
			}

			return
		} else {

			message, _ := MessageFromParsedBytes(buf.Bytes())

			msgType := new(StringValue)
			message.Header.GetField(tag.MsgType, msgType)

			sentMessageSeqNum := new(IntValue)
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
}

func (state inSession) handleTestRequest(session *session, msg Message) (nextState sessionState) {
	if err := session.verify(msg); err != nil {
		return state.processReject(session, err)
	}

	testReq := new(StringValue)
	if err := msg.Body.GetField(tag.TestReqID, testReq); err != nil {
		session.log.OnEvent("Test Request with no testRequestID")
	} else {
		heartBt := NewMessageBuilder()
		heartBt.Header.Set(NewStringField(tag.MsgType, "0"))
		heartBt.Body.Set(NewStringField(tag.TestReqID, testReq.Value))
		session.send(heartBt)
	}

	session.expectedSeqNum++

	return state
}

func (state inSession) processReject(session *session, rej MessageReject) (nextState sessionState) {
	switch TypedError := rej.(type) {
	case TargetTooHigh:

		switch session.currentState.(type) {
		default:
			session.DoTargetTooHigh(TypedError)
		case resendState:
			//assumes target too high reject already sent
		}
		session.messageStash[TypedError.ReceivedTarget] = rej.RejectedMessage()
		return resendState{}

	case TargetTooLow:
		return state.doTargetTooLow(session, TypedError)
	case IncorrectBeginString:
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

func (state inSession) doTargetTooLow(session *session, rej TargetTooLow) (nextState sessionState) {
	posDupFlag := new(BooleanValue)
	if err := rej.RejectedMessage().Header.GetField(tag.PossDupFlag, posDupFlag); err == nil && posDupFlag.Value {

		origSendingTime := new(UTCTimestampValue)
		if err := rej.RejectedMessage().Header.GetField(tag.OrigSendingTime, origSendingTime); err != nil {
			session.doReject(NewRequiredTagMissing(rej.RejectedMessage(), tag.OrigSendingTime))
			return state
		} else {
			sendingTime := new(UTCTimestampValue)
			rej.RejectedMessage().Header.GetField(tag.SendingTime, sendingTime)

			if sendingTime.Value.Before(origSendingTime.Value) {
				session.doReject(NewSendingTimeAccuracyProblem(rej.RejectedMessage()))
				return state.initiateLogout(session, "")
			}
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

	sequenceReset.Header.Set(NewStringField(tag.MsgType, "4"))
	sequenceReset.Header.Set(NewIntField(tag.MsgSeqNum, beginSeqNo))
	sequenceReset.Header.Set(NewBooleanField(tag.PossDupFlag, true))
	sequenceReset.Body.Set(NewIntField(tag.NewSeqNo, endSeqNo))
	sequenceReset.Body.Set(NewBooleanField(tag.GapFillFlag, true))

	origSendingTime := new(StringValue)
	if err := sequenceReset.Header.GetField(tag.SendingTime, origSendingTime); err == nil {
		sequenceReset.Header.Set(NewStringField(tag.OrigSendingTime, origSendingTime.Value))
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
	reply.Header.Set(NewStringField(tag.MsgType, "5"))
	reply.Header.Set(NewStringField(tag.BeginString, session.BeginString))
	reply.Header.Set(NewStringField(tag.TargetCompID, session.TargetCompID))
	reply.Header.Set(NewStringField(tag.SenderCompID, session.SenderCompID))

	if reason != "" {
		reply.Body.Set(NewStringField(tag.Text, reason))
	}

	session.send(reply)
	session.log.OnEvent("Sending logout response")
}
