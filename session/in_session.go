package session

import (
	"quickfixgo/fix"
	"quickfixgo/message"
	"quickfixgo/message/basic"
	"quickfixgo/reject"
	"time"
)

type inSession struct {
}

func (state inSession) FixMsgIn(session *session, msg message.Message) (nextState state) {
	if msgType, err := msg.Header().StringField(fix.MsgType); err == nil {
		switch msgType.Value() {
		//logout
		case "5":
			return state.handleLogout(session, msg)
		//test request
		case "1":
			return state.handleTestRequest(session, msg)
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
		heartBt.MsgHeader.Set(basic.NewStringField(fix.MsgType, "0"))
		session.send(heartBt)
	case peerTimeout:
		testReq := basic.NewMessage()
		testReq.MsgHeader.Set(basic.NewStringField(fix.MsgType, "1"))
		testReq.MsgBody.Set(basic.NewStringField(fix.TestReqID, "TEST"))
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

func (state inSession) handleTestRequest(session *session, msg message.Message) (nextState state) {
	if err := session.verify(msg); err != nil {
		return state.processReject(session, err)
	}

	if testReq, err := msg.Body().StringField(fix.TestReqID); err != nil {
		session.log.OnEvent("Test Request with no testRequestID")
	} else {
		heartBt := basic.NewMessage()
		heartBt.MsgHeader.Set(basic.NewStringField(fix.MsgType, "0"))
		heartBt.MsgBody.Set(testReq)
		session.send(heartBt)
	}

	session.expectedSeqNum++

	return state
}

func (state inSession) processReject(session *session, rej reject.MessageReject) (nextState state) {
	switch TypedError := rej.(type) {
	case reject.TargetTooHigh:
		session.DoTargetTooHigh(TypedError)
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
	if posDupFlag, ok := rej.RejectedMessage().Header().Get(fix.PossDupFlag); ok && posDupFlag.Value() == "Y" {

		if origSendingTime, err := rej.RejectedMessage().Header().UTCTimestampField(fix.OrigSendingTime); err != nil {
			session.doReject(reject.NewRequiredTagMissing(rej.RejectedMessage(), fix.OrigSendingTime))
			return state
		} else {
			sendingTime, _ := rej.RejectedMessage().Header().UTCTimestampField(fix.SendingTime)

			if sendingTime.UTCTimestampValue().Before(origSendingTime.UTCTimestampValue()) {
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

func (state *inSession) generateLogout(session *session) {
	state.generateLogoutWithReason(session, "")
}

func (state *inSession) generateLogoutWithReason(session *session, reason string) {
	reply := basic.NewMessage()
	reply.MsgHeader.Set(basic.NewStringField(fix.MsgType, "5"))
	reply.MsgHeader.Set(basic.NewStringField(fix.BeginString, session.BeginString))
	reply.MsgHeader.Set(basic.NewStringField(fix.TargetCompID, session.TargetCompID))
	reply.MsgHeader.Set(basic.NewStringField(fix.SenderCompID, session.SenderCompID))

	if reason != "" {
		reply.MsgBody.Set(basic.NewStringField(fix.Text, reason))
	}

	session.send(reply)
	session.log.OnEvent("Sending logout response")
}
