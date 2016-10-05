package quickfix

import "github.com/quickfixgo/quickfix/internal"

type logoutState struct{ connectedNotLoggedOn }

func (state logoutState) String() string { return "Logout State" }

func (state logoutState) FixMsgIn(session *session, msg *Message) (nextState sessionState) {
	nextState = inSession{}.FixMsgIn(session, msg)
	if nextState, ok := nextState.(latentState); ok {
		return nextState
	}

	return state
}

func (state logoutState) Timeout(session *session, event internal.Event) (nextState sessionState) {
	switch event {
	case internal.LogoutTimeout:
		session.log.OnEvent("Timed out waiting for logout response")
		return latentState{}
	}

	return state
}

func (state logoutState) Stop(session *session) (nextstate sessionState) {
	return state
}
