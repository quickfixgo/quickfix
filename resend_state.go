package quickfix

type resendState struct {
	inSession
}

func (state resendState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	for ok := true; ok; {
		nextState = state.inSession.FixMsgIn(session, msg)

		if !nextState.IsLoggedOn() {
			return
		}

		msg, ok = session.messageStash[session.store.NextTargetMsgSeqNum()]
	}

	if len(session.messageStash) != 0 {
		nextState = resendState{}
	}

	return
}
