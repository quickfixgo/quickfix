package quickfix

import "github.com/quickfixgo/quickfix/internal"

type pendingTimeout struct {
	sessionState
}

func (s pendingTimeout) Timeout(session *session, event internal.Event) (nextState sessionState) {
	switch event {

	default:
		return s

	case internal.PeerTimeout:
		session.log.OnEvent("Session Timeout")
	case internal.SessionExpire:
		session.sendLogoutAndReset()
	}

	return latentState{}

}
