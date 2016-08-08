package quickfix

import "github.com/quickfixgo/quickfix/enum"

type logonState struct{}

func (s logonState) String() string   { return "Logon State" }
func (s logonState) IsLoggedOn() bool { return false }
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
		return handleStateError(session, err)
	}
	return inSession{}
}

func (s logonState) Timeout(session *session, e event) (nextState sessionState) {
	if e == logonTimeout {
		session.log.OnEvent("Timed out waiting for logon response")
		return latentState{}
	}

	return s
}
