package quickfix

import "github.com/quickfixgo/quickfix/internal"

type notSessionTime struct{ latentState }

func (notSessionTime) String() string      { return "Not session time" }
func (notSessionTime) IsSessionTime() bool { return false }

func (state notSessionTime) FixMsgIn(session *session, msg *Message) (nextState sessionState) {
	session.log.OnEventf("Invalid Session State: Unexpected Msg %v while in Latent state", msg)
	return state
}

func (state notSessionTime) Timeout(*session, internal.Event) (nextState sessionState) {
	return state
}

func (state notSessionTime) Stop(*session) (nextState sessionState) {
	return state
}
