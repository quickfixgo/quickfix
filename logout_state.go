package quickfix

type logoutState struct {
}

func (state logoutState) String() string { return "Logout State" }
func (s logoutState) IsLoggedOn() bool   { return false }

func (state logoutState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	return state
}

func (state logoutState) Timeout(session *session, event event) (nextState sessionState) {
	switch event {
	case logoutTimeout:
		session.log.OnEvent("Timed out waiting for Logout response")
		return latentState{}
	}

	return state
}
