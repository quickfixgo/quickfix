package quickfix

import (
	"bytes"

	"github.com/quickfixgo/quickfix/internal"
)

type logonState struct{ connectedNotLoggedOn }

func (s logonState) String() string { return "Logon State" }

func (s logonState) FixMsgIn(session *session, msg *Message) (nextState sessionState) {
	msgType, err := msg.Header.GetBytes(tagMsgType)
	if err != nil {
		return handleStateError(session, err)
	}

	if !bytes.Equal(msgType, msgTypeLogon) {
		session.log.OnEventf("Invalid Session State: Received Msg %s while waiting for Logon", msg)
		return latentState{}
	}

	if err := session.handleLogon(msg); err != nil {
		switch err := err.(type) {
		case RejectLogon:
			return shutdownWithReason(session, msg, true, err.Error())

		case targetTooLow:
			return shutdownWithReason(session, msg, false, err.Error())

		case targetTooHigh:
			var tooHighErr error
			if nextState, tooHighErr = session.doTargetTooHigh(err); tooHighErr != nil {
				return shutdownWithReason(session, msg, false, tooHighErr.Error())
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

func shutdownWithReason(session *session, msg *Message, incrNextTargetMsgSeqNum bool, reason string) (nextState sessionState) {
	session.log.OnEvent(reason)
	logout := session.buildLogout(reason)

	if err := session.dropAndSendInReplyTo(logout, msg); err != nil {
		session.logError(err)
	}

	if incrNextTargetMsgSeqNum {
		if err := session.store.IncrNextTargetMsgSeqNum(); err != nil {
			session.logError(err)
		}
	}

	return latentState{}
}
