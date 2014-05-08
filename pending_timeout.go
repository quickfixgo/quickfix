package quickfix

type pendingTimeout struct {
	inSession
}

func (currentState pendingTimeout) Timeout(session *Session, event event) (nextState sessionState) {
	switch event {
	case peerTimeout:
		session.log.OnEvent("Session Timeout")
		return latentState{}
	}

	return currentState
}
