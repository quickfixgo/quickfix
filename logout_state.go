package quickfix

import "github.com/quickfixgo/quickfix/enum"

type logoutState struct{}

func (state logoutState) String() string   { return "Logout State" }
func (state logoutState) IsLoggedOn() bool { return false }

func (state logoutState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err != nil {
		return session.handleError(err)
	}

	switch string(msgType) {
	case enum.MsgType_LOGOUT:
		if err := session.verifySelect(msg, false, false); err != nil {
			session.logError(err)
			return state
		}

		session.log.OnEvent("Received logout response")
		if err := session.store.IncrNextTargetMsgSeqNum(); err != nil {
			session.logError(err)
		}

		return latentState{}
	default:
		nextState = inSession{}.FixMsgIn(session, msg)
		if nextState, ok := nextState.(latentState); ok {
			return nextState
		}

		return state
	}
}

func (state logoutState) Timeout(session *session, event event) (nextState sessionState) {
	switch event {
	case logoutTimeout:
		session.log.OnEvent("Timed out waiting for logout response")
		return latentState{}
	}

	return state
}
