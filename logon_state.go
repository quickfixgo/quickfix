package quickfix

import "github.com/quickfixgo/quickfix/enum"

type logonState struct{}

func (state logonState) String() string { return "Logon State" }
func (s logonState) IsLoggedOn() bool   { return false }

func (s logonState) VerifyMsgIn(session *session, msg Message) MessageRejectError {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err != nil {
		return RequiredTagMissing(tagMsgType)
	}

	switch string(msgType) {
	case enum.MsgType_LOGON:
		err := session.verifyLogon(msg)
		if err != nil {
			session.log.OnEvent(err.Error())
		}
		return err
	default:
		session.log.OnEventf("Invalid Session State: Received Msg %v while waiting for Logon", msg)
		return InvalidMessageType()
	}
}

func (s logonState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	if err := session.handleLogon(msg); err != nil {
		session.log.OnEvent(err.Error())
		return latentState{}
	}
	return inSession{}
}

func (s logonState) FixMsgInRej(session *session, msg Message, err MessageRejectError) sessionState {
	return latentState{}
}

func (s logonState) Timeout(session *session, e event) (nextState sessionState) {
	if e == logonTimeout {
		session.log.OnEvent("Timed out waiting for logon response")
		return latentState{}
	}

	return s
}
