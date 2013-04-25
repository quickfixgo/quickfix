package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
	"time"
)

type inSession struct {
}

func (state inSession) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	msgType := new(MsgType)
	if err := msg.Header.Get(msgType); err == nil {
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
		heartBt := NewMessage()
		heartBt.Header.SetField(NewStringField(tag.MsgType, "0"))
		session.send(heartBt)
	case peerTimeout:
		testReq := NewMessage()
		testReq.Header.SetField(NewStringField(tag.MsgType, "1"))
		testReq.Body.SetField(NewStringField(tag.TestReqID, "TEST"))
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
	gapFillFlag := new(GapFillFlag)
	msg.Body.Get(gapFillFlag)

	if err := session.verifySelect(msg, gapFillFlag.Value, gapFillFlag.Value); err != nil {
		return state.processReject(session, err)
	}

	newSeqNo := new(NewSeqNo)
	if err := msg.Body.Get(newSeqNo); err == nil {
		session.log.OnEventf("Received SequenceReset FROM: %v TO: %v", session.expectedSeqNum, newSeqNo.Value)

		switch {
		case newSeqNo.Value > session.expectedSeqNum:
			session.expectedSeqNum = newSeqNo.Value
		case newSeqNo.Value < session.expectedSeqNum:
			session.doReject(NewValueIsIncorrect(msg, tag.NewSeqNo))
		}
	}

	return state
}

func (state inSession) handleResendRequest(session *session, msg Message) (nextState sessionState) {
	if err := session.verifyIgnoreSeqNumTooHighOrLow(msg); err != nil {
		return state.processReject(session, err)
	}

	var beginSeqNo, endSeqNo int
	beginSeqNoField := new(BeginSeqNo)
	if err := msg.Body.Get(beginSeqNoField); err != nil {
		return state.processReject(session, NewRequiredTagMissing(msg, tag.BeginSeqNo))
	} else {
		beginSeqNo = beginSeqNoField.Value
	}

	endSeqNoField := new(EndSeqNo)
	if err := msg.Body.Get(endSeqNoField); err != nil {
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

			msgType := new(MsgType)
			message.Header.Get(msgType)

			sentMessageSeqNum := new(MsgSeqNum)
			message.Header.Get(sentMessageSeqNum)

			if msgType.IsAdminMessageType() {
				nextSeqNum = sentMessageSeqNum.Value + 1
			} else {

				if seqNum != sentMessageSeqNum.Value {
					state.generateSequenceReset(session, seqNum, sentMessageSeqNum.Value)
				}

				session.resend(message)
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

	testReq := new(TestReqID)
	if err := msg.Body.Get(testReq); err != nil {
		session.log.OnEvent("Test Request with no testRequestID")
	} else {
		heartBt := NewMessage()
		heartBt.Header.SetField(NewStringField(tag.MsgType, "0"))
		heartBt.Body.SetField(testReq)
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

	session.doReject(rej)
	return state.initiateLogout(session, "")
}

func (state inSession) doTargetTooLow(session *session, rej TargetTooLow) (nextState sessionState) {
	posDupFlag := new(PossDupFlag)
	if err := rej.RejectedMessage().Header.Get(posDupFlag); err == nil && posDupFlag.Value {

		origSendingTime := new(OrigSendingTime)
		if err := rej.RejectedMessage().Header.Get(origSendingTime); err != nil {
			session.doReject(NewRequiredTagMissing(rej.RejectedMessage(), tag.OrigSendingTime))
			return state
		} else {
			sendingTime := new(SendingTime)
			rej.RejectedMessage().Header.Get(sendingTime)

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
	sequenceReset := NewMessage()
	session.fillDefaultHeader(sequenceReset)

	sequenceReset.Header.SetField(NewStringField(tag.MsgType, "4"))
	sequenceReset.Header.SetField(NewIntField(tag.MsgSeqNum, beginSeqNo))
	sequenceReset.Header.SetField(NewBooleanField(tag.PossDupFlag, true))
	sequenceReset.Body.SetField(NewIntField(tag.NewSeqNo, endSeqNo))
	sequenceReset.Body.SetField(NewBooleanField(tag.GapFillFlag, true))

	//FIXME
	origSendingTime := NewStringField(tag.SendingTime, "")
	if err := sequenceReset.Header.Get(origSendingTime); err == nil {
		sequenceReset.SetHeaderField(NewStringField(tag.OrigSendingTime, origSendingTime.Value))
	}

	buffer := sequenceReset.Build()
	session.sendBuffer(buffer)
}

func (state *inSession) generateLogout(session *session) {
	state.generateLogoutWithReason(session, "")
}

func (state *inSession) generateLogoutWithReason(session *session, reason string) {
	reply := NewMessage()
	reply.Header.SetField(NewStringField(tag.MsgType, "5"))
	reply.Header.SetField(NewStringField(tag.BeginString, session.BeginString))
	reply.Header.SetField(NewStringField(tag.TargetCompID, session.TargetCompID))
	reply.Header.SetField(NewStringField(tag.SenderCompID, session.SenderCompID))

	if reason != "" {
		reply.Body.SetField(NewStringField(tag.Text, reason))
	}

	session.send(reply)
	session.log.OnEvent("Sending logout response")
}
