package quickfix

import "github.com/quickfixgo/quickfix/internal"

type latentState struct{}

func (state latentState) String() string    { return "Latent State" }
func (state latentState) IsLoggedOn() bool  { return false }
func (state latentState) IsConnected() bool { return false }

func (state latentState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	session.log.OnEventf("Invalid Session State: Unexpected Msg %v while in Latent state", msg)
	return state
}

func (state latentState) Timeout(*session, internal.Event) (nextState sessionState) {
	return state
}
