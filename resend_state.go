package quickfix

type resendState struct{}

func (s resendState) String() string { return "Resend" }

func (s resendState) IsLoggedOn() bool { return true }

func (s resendState) Timeout(session *session, event event) (nextState sessionState) {
	nextState = inSession{}.Timeout(session, event)
	switch nextState.(type) {
	case inSession:
		nextState = s
	case pendingTimeout:
		//wrap pendingTimeout in resend. prevents us falling back to inSession if recovering
		//from pendingTimeout
		nextState = pendingTimeout{s}
	}

	return
}

func (s resendState) VerifyMsgIn(session *session, msg Message) (err MessageRejectError) {
	return inSession{}.VerifyMsgIn(session, msg)
}

func (s resendState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	session.log.OnEventf("Got FIXMsgIn in resend %s", msg.rawMessage)
	return s.handleNextState(session, inSession{}.FixMsgIn(session, msg))
}

func (s resendState) FixMsgInRej(session *session, msg Message, rej MessageRejectError) (nextState sessionState) {
	return s.handleNextState(session, inSession{}.FixMsgInRej(session, msg, rej))
}

func (s resendState) handleNextState(session *session, nextState sessionState) sessionState {
	if !nextState.IsLoggedOn() || len(session.messageStash) == 0 {
		return nextState
	}

	targetSeqNum := session.store.NextTargetMsgSeqNum()
	if msg, ok := session.messageStash[targetSeqNum]; ok {
		delete(session.messageStash, targetSeqNum)
		session.resendIn <- msg
	}

	return s
}
