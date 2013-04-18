package session

import (
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/message/basic"
	"github.com/cbusbey/quickfixgo/reject"
	"time"
)

type inSession struct {
}

func (state inSession) FixMsgIn(session *session, msg message.Message) (nextState state) {
	if msgType, ok := msg.Header().StringValue(fix.MsgType); ok {
		switch msgType {
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

func (state inSession) Timeout(session *session, event event) (nextState state) {
	switch event {
	case needHeartbeat:
		heartBt := basic.NewMessage()
		heartBt.MsgHeader.SetField(basic.NewStringField(fix.MsgType, "0"))
		session.send(heartBt)
	case peerTimeout:
		testReq := basic.NewMessage()
		testReq.MsgHeader.SetField(basic.NewStringField(fix.MsgType, "1"))
		testReq.MsgBody.SetField(basic.NewStringField(fix.TestReqID, "TEST"))
		session.send(testReq)
		session.peerTimer.Reset(time.Duration(int64(1.2 * float64(session.heartBeatTimeout))))
		return pendingTimeout{}
	}
	return state
}

func (state inSession) handleLogout(session *session, msg message.Message) (nextState state) {
	session.log.OnEvent("Received logout request")
	state.generateLogout(session)
	session.callback.OnLogout(session.ID)

	return latentState{}
}

func (state inSession) handleSequenceReset(session *session, msg message.Message) (nextState state) {
	isGapFill := false

	if gapFillFlag, err := msg.Body().BooleanValue(fix.GapFillFlag); err == nil {
		isGapFill = gapFillFlag
	}

	if err := session.verifySelect(msg, isGapFill, isGapFill); err != nil {
		return state.processReject(session, err)
	}

	if newSeqNo, err := msg.Body().IntValue(fix.NewSeqNo); err == nil {
		session.log.OnEventf("Received SequenceReset FROM: %v TO: %v", session.expectedSeqNum, newSeqNo)

		switch {
		case newSeqNo > session.expectedSeqNum:
			session.expectedSeqNum = newSeqNo
		case newSeqNo < session.expectedSeqNum:
			session.doReject(reject.NewValueIsIncorrect(msg, fix.NewSeqNo))
		}
	}

	return state
}

func (state inSession) handleResendRequest(session *session, msg message.Message) (nextState state) {
	if err := session.verifyIgnoreSeqNumTooHighOrLow(msg); err != nil {
		return state.processReject(session, err)
	}

	var beginSeqNo, endSeqNo int
	if beginSeqNoField, err := msg.Body().IntValue(fix.BeginSeqNo); err != nil {
		return state.processReject(session, reject.NewRequiredTagMissing(msg, fix.BeginSeqNo))
	} else {
		beginSeqNo = beginSeqNoField
	}

	if endSeqNoField, err := msg.Body().IntValue(fix.EndSeqNo); err != nil {
		return state.processReject(session, reject.NewRequiredTagMissing(msg, fix.EndSeqNo))
	} else {
		endSeqNo = endSeqNoField
	}

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

	for {
		if buf, ok := <-buffers; !ok {
			//gapfill for catch-up
			if seqNum != nextSeqNum {
				state.generateSequenceReset(session, seqNum, nextSeqNum)
			}

			return
		} else {

			message, _ := basic.MessageFromParsedBytes(buf.Bytes())
			msgType, _ := message.MsgHeader.StringValue(fix.MsgType)
			sentMessageSeqNum, _ := message.MsgHeader.IntValue(fix.MsgSeqNum)

			if IsAdminMessageType(msgType) {
				nextSeqNum = sentMessageSeqNum + 1
			} else {

				if seqNum != sentMessageSeqNum {
					state.generateSequenceReset(session, seqNum, sentMessageSeqNum)
				}

				session.resend(message)
				seqNum = sentMessageSeqNum + 1
				nextSeqNum = seqNum
			}
		}
	}
}

func (state inSession) handleTestRequest(session *session, msg message.Message) (nextState state) {
	if err := session.verify(msg); err != nil {
		return state.processReject(session, err)
	}

	if testReq, ok := msg.Body().Field(fix.TestReqID); !ok {
		session.log.OnEvent("Test Request with no testRequestID")
	} else {
		heartBt := basic.NewMessage()
		heartBt.MsgHeader.SetField(basic.NewStringField(fix.MsgType, "0"))
		heartBt.MsgBody.SetField(testReq)
		session.send(heartBt)
	}

	session.expectedSeqNum++

	return state
}

func (state inSession) processReject(session *session, rej reject.MessageReject) (nextState state) {
	switch TypedError := rej.(type) {
	case reject.TargetTooHigh:

		switch session.currentState.(type) {
		default:
			session.DoTargetTooHigh(TypedError)
		case resendState:
			//assumes target too high reject already sent
		}
		session.messageStash[TypedError.ReceivedTarget] = rej.RejectedMessage()
		return resendState{}

	case reject.TargetTooLow:
		return state.doTargetTooLow(session, TypedError)
	case reject.IncorrectBeginString:
		return state.initiateLogout(session, rej.Error())
	}

	session.doReject(rej)
	return state.initiateLogout(session, "")
}

func (state inSession) doTargetTooLow(session *session, rej reject.TargetTooLow) (nextState state) {
	if posDupFlag, err := rej.RejectedMessage().Header().BooleanValue(fix.PossDupFlag); err == nil && posDupFlag {

		if origSendingTime, err := rej.RejectedMessage().Header().UTCTimestampValue(fix.OrigSendingTime); err != nil {
			session.doReject(reject.NewRequiredTagMissing(rej.RejectedMessage(), fix.OrigSendingTime))
			return state
		} else {
			sendingTime, _ := rej.RejectedMessage().Header().UTCTimestampValue(fix.SendingTime)

			if sendingTime.Before(origSendingTime) {
				session.doReject(reject.NewSendingTimeAccuracyProblem(rej.RejectedMessage()))
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
	sequenceReset := basic.NewMessage()
	session.fillDefaultHeader(sequenceReset)

	sequenceReset.MsgHeader.SetField(basic.NewStringField(fix.MsgType, "4"))
	sequenceReset.MsgHeader.SetField(basic.NewIntField(fix.MsgSeqNum, beginSeqNo))
	sequenceReset.MsgHeader.SetField(basic.NewBooleanField(fix.PossDupFlag, true))
	sequenceReset.MsgBody.SetField(basic.NewIntField(fix.NewSeqNo, endSeqNo))
	sequenceReset.MsgBody.SetField(basic.NewBooleanField(fix.GapFillFlag, true))

	if origSendingTime, ok := sequenceReset.MsgHeader.Field(fix.SendingTime); ok {
		sequenceReset.SetHeaderField(basic.NewStringField(fix.OrigSendingTime, origSendingTime.Value()))
	}

	buffer := sequenceReset.Build()
	session.sendBuffer(buffer)
}

func (state *inSession) generateLogout(session *session) {
	state.generateLogoutWithReason(session, "")
}

func (state *inSession) generateLogoutWithReason(session *session, reason string) {
	reply := basic.NewMessage()
	reply.MsgHeader.SetField(basic.NewStringField(fix.MsgType, "5"))
	reply.MsgHeader.SetField(basic.NewStringField(fix.BeginString, session.BeginString))
	reply.MsgHeader.SetField(basic.NewStringField(fix.TargetCompID, session.TargetCompID))
	reply.MsgHeader.SetField(basic.NewStringField(fix.SenderCompID, session.SenderCompID))

	if reason != "" {
		reply.MsgBody.SetField(basic.NewStringField(fix.Text, reason))
	}

	session.send(reply)
	session.log.OnEvent("Sending logout response")
}
