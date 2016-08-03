package quickfix

type logoutState struct{}

func (state logoutState) String() string   { return "Logout State" }
func (state logoutState) IsLoggedOn() bool { return false }

func (state logoutState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	nextState = inSession{}.FixMsgIn(session, msg)
	if nextState, ok := nextState.(latentState); ok {
		return nextState
	}

	return state
}

func (state logoutState) Timeout(session *session, event event) (nextState sessionState) {
	switch event {
	case logoutTimeout:
		session.log.OnEvent("Timed out waiting for logout response")
		return latentState{}
	}

	return state
}
