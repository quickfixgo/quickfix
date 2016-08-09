package quickfix

import "github.com/quickfixgo/quickfix/internal"

type resendState struct{ loggedOn }

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

	if !nextState.IsLoggedOn() || len(session.messageStash) == 0 {
		return
	}

	targetSeqNum := session.store.NextTargetMsgSeqNum()
	if msg, ok := session.messageStash[targetSeqNum]; ok {
		delete(session.messageStash, targetSeqNum)
		nextState = nextState.FixMsgIn(session, msg)
	}

	return
}
