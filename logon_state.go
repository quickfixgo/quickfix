package quickfix

import (
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/internal"
)

type logonState struct{ connectedNotLoggedOn }

func (s logonState) String() string { return "Logon State" }

func (s logonState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err != nil {
		return handleStateError(session, err)
	}

	if string(msgType) != enum.MsgType_LOGON {
		session.log.OnEventf("Invalid Session State: Received Msg %s while waiting for Logon", msg)
		return latentState{}
	}

	if err := session.handleLogon(msg); err != nil {
		switch err := err.(type) {
		case RejectLogon:
			session.log.OnEvent(err.Text)
			msg := session.buildLogout(err.Text)

			if err := session.dropAndSend(msg, false); err != nil {
				session.logError(err)
			}

			if err := session.store.IncrNextTargetMsgSeqNum(); err != nil {
				session.logError(err)
			}

			return latentState{}

		case targetTooHigh:
			var tooHighErr error
			if nextState, tooHighErr = session.doTargetTooHigh(err); tooHighErr != nil {
				return handleStateError(session, tooHighErr)
			}

			return

		default:
			return handleStateError(session, err)
		}
	}
	return inSession{}
}

func (s logonState) Timeout(session *session, e internal.Event) (nextState sessionState) {
	switch e {
	case internal.LogonTimeout:
		session.log.OnEvent("Timed out waiting for logon response")
		return latentState{}
	}
	return s
}

func (s logonState) Stop(session *session) (nextState sessionState) {
	return latentState{}
}
