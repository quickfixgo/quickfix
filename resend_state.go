package quickfix

import "github.com/quickfixgo/quickfix/internal"

type resendState struct {
	loggedOn
	messageStash map[int]Message
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

func (s resendState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	nextState = inSession{}.FixMsgIn(session, msg)

	if _, stillInResend := nextState.(resendState); stillInResend || !nextState.IsLoggedOn() {
		return
	}

	for len(s.messageStash) > 0 {
		targetSeqNum := session.store.NextTargetMsgSeqNum()
		msg, ok := s.messageStash[targetSeqNum]
		if !ok {
			return s
		}
		delete(s.messageStash, targetSeqNum)

		nextState = inSession{}.FixMsgIn(session, msg)
		if !nextState.IsLoggedOn() {
			return
		}
	}

	return
}
