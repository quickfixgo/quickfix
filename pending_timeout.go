package quickfix

type pendingTimeout struct {
	sessionState
}

func (s pendingTimeout) Timeout(session *session, event event) (nextState sessionState) {
	switch event {
	case peerTimeout:
		session.log.OnEvent("Session Timeout")
		return latentState{}
	}

	return s
}
