package quickfix

import "github.com/quickfixgo/quickfix/internal"

type pendingTimeout struct {
	sessionState
}

func (s pendingTimeout) Timeout(session *session, event internal.Event) (nextState sessionState) {
	switch event {
	case internal.PeerTimeout:
		session.log.OnEvent("Session Timeout")
		return latentState{}

	case internal.SessionExpire:
		session.sendLogoutAndReset()
		return latentState{}
	}

	return s
}
