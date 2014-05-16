package quickfix

type logoutState struct {
}

func (state logoutState) FixMsgIn(session *Session, msg Message) (nextState sessionState) {
	return state
}

func (state logoutState) Timeout(session *Session, event event) (nextState sessionState) {
	switch event {
	case logoutTimeout:
		session.log.OnEvent("Timed out waiting for Logout response")
		return latentState{}
	}

	return state
}
