package quickfixgo

type latentState struct{}

func (state latentState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	session.log.OnEventf("Invalid Session State: Unexpected Msg %v while in Latent state", msg)
	return state
}

func (state latentState) Timeout(*session, event) (nextState sessionState) {
	return state
}
