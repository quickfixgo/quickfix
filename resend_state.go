package quickfix

type resendState struct {
	inSession
}

func (state resendState) String() string { return "Resend" }

func (state resendState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	return state.handleNextState(session, state.inSession.FixMsgIn(session, msg))
}

func (state resendState) FixMsgInRej(session *session, msg Message, rej MessageRejectError) (nextState sessionState) {
	return state.handleNextState(session, state.inSession.FixMsgInRej(session, msg, rej))
}

func (state resendState) handleNextState(session *session, nextState sessionState) sessionState {
	if !nextState.IsLoggedOn() || len(session.messageStash) == 0 {
		return nextState
	}

	targetSeqNum := session.store.NextTargetMsgSeqNum()
	if msg, ok := session.messageStash[targetSeqNum]; ok {
		delete(session.messageStash, targetSeqNum)

		// FIXME add a "resend" channel to the session loop to differentiate between
		// new incoming fix messages, and stashed messages from the resend state
		go func() { session.messageIn <- fixIn{msg.rawMessage, msg.ReceiveTime} }()
	}

	return resendState{}
}
