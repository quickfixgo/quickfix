package fixt11

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/tag"
)

func Crack(msg message.Message, sessionID quickfixgo.SessionID, router MessageRouter) reject.MessageReject {
	switch msgType, _ := msg.Header.StringValue(tag.MsgType); msgType {
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
	return reject.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIXT11Heartbeat(msg Heartbeat, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIXT11TestRequest(msg TestRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIXT11ResendRequest(msg ResendRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIXT11Reject(msg Reject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIXT11SequenceReset(msg SequenceReset, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIXT11Logout(msg Logout, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIXT11Logon(msg Logon, sessionID quickfixgo.SessionID) reject.MessageReject
}
type FIXT11MessageCracker struct{}

func (c *FIXT11MessageCracker) OnFIXT11Heartbeat(msg Heartbeat, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIXT11MessageCracker) OnFIXT11TestRequest(msg TestRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIXT11MessageCracker) OnFIXT11ResendRequest(msg ResendRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIXT11MessageCracker) OnFIXT11Reject(msg Reject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIXT11MessageCracker) OnFIXT11SequenceReset(msg SequenceReset, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIXT11MessageCracker) OnFIXT11Logout(msg Logout, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIXT11MessageCracker) OnFIXT11Logon(msg Logon, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
