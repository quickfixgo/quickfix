package quickfix

type latentState struct{}

func (state latentState) String() string   { return "Latent State" }
func (state latentState) IsLoggedOn() bool { return false }

func (state latentState) VerifyMsgIn(session *session, msg Message) MessageRejectError {
	return InvalidMessageType()
}

func (state latentState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	session.log.OnEventf("Invalid Session State: Unexpected Msg %v while in Latent state", msg)
	return state
}

func (state latentState) FixMsgInRej(session *session, msg Message, err MessageRejectError) (nextState sessionState) {
	return state.FixMsgIn(session, msg)
}

func (state latentState) Timeout(*session, event) (nextState sessionState) {
	return state
}
