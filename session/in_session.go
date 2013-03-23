package session

import (
	"quickfixgo/fix"
	"quickfixgo/message"
	"quickfixgo/reject"
)

type inSession struct {
}

func (state inSession) FixMsgIn(session *session, msg message.Message) (nextState state) {
	if msgType, err := msg.Header().StringField(fix.MsgType); err == nil {
		switch msgType.Value() {
		//logout
		case "5":
			return state.handleLogout(session, msg)
		default:
			if err := session.verify(msg); err != nil {
				return state.processReject(session, err)
			}
		}
	}

	session.expectedSeqNum++

	return state
}

func (state inSession) Timeout(*session, event) (nextState state) {
	return state
}

func (state inSession) handleLogout(session *session, msg message.Message) (nextState state) {
	session.log.OnEvent("Received logout request")
	session.generateLogout()
	session.callback.OnLogout(session.ID)

	return latentState{}
}

func (state inSession) processReject(session *session, rej reject.MessageReject) (nextState state) {
	switch TypedError := rej.(type) {
	case reject.TargetTooHigh:
		session.DoTargetTooHigh(TypedError)
		return resendState{}
	case reject.TargetTooLow:
		return state.doTargetTooLow(session, TypedError)
	}

	session.initiateLogout("")
	return logoutState{}
}

func (state inSession) doTargetTooLow(session *session, rej reject.TargetTooLow) (nextState state) {
	session.initiateLogout(rej.Error())
	return logoutState{}
}
