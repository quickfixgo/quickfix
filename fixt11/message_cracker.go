package fixt11

import (
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session"
)

func Crack(msg message.Message, sessionID session.ID, router MessageRouter) reject.MessageReject {
	switch msgType, _ := msg.Header().StringField(fix.MsgType); msgType.Value() {
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
	OnFIXT11Heartbeat(msg Heartbeat, sessionID session.ID) reject.MessageReject
	OnFIXT11TestRequest(msg TestRequest, sessionID session.ID) reject.MessageReject
	OnFIXT11ResendRequest(msg ResendRequest, sessionID session.ID) reject.MessageReject
	OnFIXT11Reject(msg Reject, sessionID session.ID) reject.MessageReject
	OnFIXT11SequenceReset(msg SequenceReset, sessionID session.ID) reject.MessageReject
	OnFIXT11Logout(msg Logout, sessionID session.ID) reject.MessageReject
	OnFIXT11Logon(msg Logon, sessionID session.ID) reject.MessageReject
}
type FIXT11MessageCracker struct{}

func (c *FIXT11MessageCracker) OnFIXT11Heartbeat(msg Heartbeat, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIXT11MessageCracker) OnFIXT11TestRequest(msg TestRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIXT11MessageCracker) OnFIXT11ResendRequest(msg ResendRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIXT11MessageCracker) OnFIXT11Reject(msg Reject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIXT11MessageCracker) OnFIXT11SequenceReset(msg SequenceReset, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIXT11MessageCracker) OnFIXT11Logout(msg Logout, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIXT11MessageCracker) OnFIXT11Logon(msg Logon, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
