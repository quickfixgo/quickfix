package session

import (
	"quickfixgo/message"
)

type latentState struct{}

func (state latentState) FixMsgIn(session *session, msg message.Message) (nextState state) {
	session.log.OnEventf("Invalid Session State: Unexpected Msg %v while in Latent state", msg)
	return state
}

func (state latentState) Timeout(*session, event) (nextState state) {
	return state
}
