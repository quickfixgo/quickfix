package quickfix

import "github.com/quickfixgo/quickfix/internal"

type resendState struct {
	loggedOn
	messageStash          map[int]*Message
	currentResendRangeEnd int
	resendRangeEnd        int
}

func (s resendState) String() string { return "Resend" }

func (s resendState) Timeout(session *session, event internal.Event) (nextState sessionState) {
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

func (s resendState) FixMsgIn(session *session, msg *Message) (nextState sessionState) {
	nextState = inSession{}.FixMsgIn(session, msg)

	if !nextState.IsLoggedOn() {
		return
	}

	if s.currentResendRangeEnd != 0 && s.currentResendRangeEnd < session.store.NextTargetMsgSeqNum() {
		nextResendState, err := session.sendResendRequest(session.store.NextTargetMsgSeqNum(), s.resendRangeEnd)
		if err != nil {
			return handleStateError(session, err)
		}
		nextResendState.messageStash = s.messageStash
		return nextResendState
	}

	if s.resendRangeEnd >= session.store.NextTargetMsgSeqNum() {
		return s
	}

	for len(s.messageStash) > 0 {
		targetSeqNum := session.store.NextTargetMsgSeqNum()
		msg, ok := s.messageStash[targetSeqNum]
		if !ok {
			break
		}

		delete(s.messageStash, targetSeqNum)

		//return stashed message to pool
		session.returnToPool(msg)

		nextState = inSession{}.FixMsgIn(session, msg)
		if !nextState.IsLoggedOn() {
			return
		}
	}

	return
}
