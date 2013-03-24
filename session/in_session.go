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
	}
	return state
}

func (state inSession) handleLogout(session *session, msg message.Message) (nextState state) {
	session.log.OnEvent("Received logout request")
	session.generateLogout()
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
		return resendState{}
	case reject.TargetTooLow:
		return state.doTargetTooLow(session, TypedError)
	}

	session.doReject(rej)
	session.initiateLogout("")
	return logoutState{}
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
				session.initiateLogout("")
				return logoutState{}
			}
		}

		if appReject := session.fromCallback(rej.RejectedMessage()); appReject != nil {
			session.doReject(appReject)
			session.initiateLogout("")
			return logoutState{}
		}
	} else {
		session.initiateLogout(rej.Error())
		return logoutState{}
	}

	return state
}
