package fixt11

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

func Crack(msg quickfix.Message, sessionID quickfix.SessionID, router MessageRouter) quickfix.MessageReject {

	msgType := new(field.MsgType)
	switch msg.Header.Get(msgType); msgType.Value {
	case "0":
		return router.OnFIXT11Heartbeat(Heartbeat{msg}, sessionID)
	case "1":
		return router.OnFIXT11TestRequest(TestRequest{msg}, sessionID)
	case "2":
		return router.OnFIXT11ResendRequest(ResendRequest{msg}, sessionID)
	case "3":
		return router.OnFIXT11Reject(Reject{msg}, sessionID)
	case "4":
		return router.OnFIXT11SequenceReset(SequenceReset{msg}, sessionID)
	case "5":
		return router.OnFIXT11Logout(Logout{msg}, sessionID)
	case "A":
		return router.OnFIXT11Logon(Logon{msg}, sessionID)
	}
	return quickfix.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIXT11Heartbeat(msg Heartbeat, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIXT11TestRequest(msg TestRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIXT11ResendRequest(msg ResendRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIXT11Reject(msg Reject, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIXT11SequenceReset(msg SequenceReset, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIXT11Logout(msg Logout, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIXT11Logon(msg Logon, sessionID quickfix.SessionID) quickfix.MessageReject
}
type FIXT11MessageCracker struct{}

func (c *FIXT11MessageCracker) OnFIXT11Heartbeat(msg Heartbeat, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIXT11MessageCracker) OnFIXT11TestRequest(msg TestRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIXT11MessageCracker) OnFIXT11ResendRequest(msg ResendRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIXT11MessageCracker) OnFIXT11Reject(msg Reject, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIXT11MessageCracker) OnFIXT11SequenceReset(msg SequenceReset, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIXT11MessageCracker) OnFIXT11Logout(msg Logout, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIXT11MessageCracker) OnFIXT11Logon(msg Logon, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
