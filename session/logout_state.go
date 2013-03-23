package session

import (
	"quickfixgo/message"
)

type logoutState struct {
}

func (state logoutState) FixMsgIn(session *session, msg message.Message) (nextState state) {
	return state
}

func (state logoutState) Timeout(session *session, event event) (nextState state) {
	switch event {
	case logoutTimeout:
		session.log.OnEvent("Timed out waiting for Logout response")
		return latentState{}
	}

	return state
}
