package fix41

import (
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session"
)

func Crack(msg message.Message, sessionID session.ID, router MessageRouter) reject.MessageReject {
	switch msgType, _ := msg.Header().StringField(fix.MsgType); msgType.Value() {
	case "0":
		return router.OnFIX41Heartbeat(Heartbeat{msg}, sessionID)
	case "1":
		return router.OnFIX41TestRequest(TestRequest{msg}, sessionID)
	case "2":
		return router.OnFIX41ResendRequest(ResendRequest{msg}, sessionID)
	case "3":
		return router.OnFIX41Reject(Reject{msg}, sessionID)
	case "4":
		return router.OnFIX41SequenceReset(SequenceReset{msg}, sessionID)
	case "5":
		return router.OnFIX41Logout(Logout{msg}, sessionID)
	case "6":
		return router.OnFIX41IndicationofInterest(IndicationofInterest{msg}, sessionID)
	case "7":
		return router.OnFIX41Advertisement(Advertisement{msg}, sessionID)
	case "8":
		return router.OnFIX41ExecutionReport(ExecutionReport{msg}, sessionID)
	case "9":
		return router.OnFIX41OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "A":
		return router.OnFIX41Logon(Logon{msg}, sessionID)
	case "B":
		return router.OnFIX41News(News{msg}, sessionID)
	case "C":
		return router.OnFIX41Email(Email{msg}, sessionID)
	case "D":
		return router.OnFIX41NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "E":
		return router.OnFIX41NewOrderList(NewOrderList{msg}, sessionID)
	case "F":
		return router.OnFIX41OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "G":
		return router.OnFIX41OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "H":
		return router.OnFIX41OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "J":
		return router.OnFIX41Allocation(Allocation{msg}, sessionID)
	case "K":
		return router.OnFIX41ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "L":
		return router.OnFIX41ListExecute(ListExecute{msg}, sessionID)
	case "M":
		return router.OnFIX41ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "N":
		return router.OnFIX41ListStatus(ListStatus{msg}, sessionID)
	case "P":
		return router.OnFIX41AllocationACK(AllocationACK{msg}, sessionID)
	case "Q":
		return router.OnFIX41DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "R":
		return router.OnFIX41QuoteRequest(QuoteRequest{msg}, sessionID)
	case "S":
		return router.OnFIX41Quote(Quote{msg}, sessionID)
	case "T":
		return router.OnFIX41SettlementInstructions(SettlementInstructions{msg}, sessionID)
	}
	return reject.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX41Heartbeat(msg Heartbeat, sessionID session.ID) reject.MessageReject
	OnFIX41TestRequest(msg TestRequest, sessionID session.ID) reject.MessageReject
	OnFIX41ResendRequest(msg ResendRequest, sessionID session.ID) reject.MessageReject
	OnFIX41Reject(msg Reject, sessionID session.ID) reject.MessageReject
	OnFIX41SequenceReset(msg SequenceReset, sessionID session.ID) reject.MessageReject
	OnFIX41Logout(msg Logout, sessionID session.ID) reject.MessageReject
	OnFIX41IndicationofInterest(msg IndicationofInterest, sessionID session.ID) reject.MessageReject
	OnFIX41Advertisement(msg Advertisement, sessionID session.ID) reject.MessageReject
	OnFIX41ExecutionReport(msg ExecutionReport, sessionID session.ID) reject.MessageReject
	OnFIX41OrderCancelReject(msg OrderCancelReject, sessionID session.ID) reject.MessageReject
	OnFIX41Logon(msg Logon, sessionID session.ID) reject.MessageReject
	OnFIX41News(msg News, sessionID session.ID) reject.MessageReject
	OnFIX41Email(msg Email, sessionID session.ID) reject.MessageReject
	OnFIX41NewOrderSingle(msg NewOrderSingle, sessionID session.ID) reject.MessageReject
	OnFIX41NewOrderList(msg NewOrderList, sessionID session.ID) reject.MessageReject
	OnFIX41OrderCancelRequest(msg OrderCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX41OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID session.ID) reject.MessageReject
	OnFIX41OrderStatusRequest(msg OrderStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX41Allocation(msg Allocation, sessionID session.ID) reject.MessageReject
	OnFIX41ListCancelRequest(msg ListCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX41ListExecute(msg ListExecute, sessionID session.ID) reject.MessageReject
	OnFIX41ListStatusRequest(msg ListStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX41ListStatus(msg ListStatus, sessionID session.ID) reject.MessageReject
	OnFIX41AllocationACK(msg AllocationACK, sessionID session.ID) reject.MessageReject
	OnFIX41DontKnowTrade(msg DontKnowTrade, sessionID session.ID) reject.MessageReject
	OnFIX41QuoteRequest(msg QuoteRequest, sessionID session.ID) reject.MessageReject
	OnFIX41Quote(msg Quote, sessionID session.ID) reject.MessageReject
	OnFIX41SettlementInstructions(msg SettlementInstructions, sessionID session.ID) reject.MessageReject
}
type FIX41MessageCracker struct{}

func (c *FIX41MessageCracker) OnFIX41Heartbeat(msg Heartbeat, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41TestRequest(msg TestRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41ResendRequest(msg ResendRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41Reject(msg Reject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41SequenceReset(msg SequenceReset, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41Logout(msg Logout, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41IndicationofInterest(msg IndicationofInterest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41Advertisement(msg Advertisement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41ExecutionReport(msg ExecutionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41OrderCancelReject(msg OrderCancelReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41Logon(msg Logon, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41News(msg News, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41Email(msg Email, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41NewOrderSingle(msg NewOrderSingle, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41NewOrderList(msg NewOrderList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41OrderCancelRequest(msg OrderCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41OrderStatusRequest(msg OrderStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41Allocation(msg Allocation, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41ListCancelRequest(msg ListCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41ListExecute(msg ListExecute, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41ListStatusRequest(msg ListStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41ListStatus(msg ListStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41AllocationACK(msg AllocationACK, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41DontKnowTrade(msg DontKnowTrade, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41QuoteRequest(msg QuoteRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41Quote(msg Quote, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX41MessageCracker) OnFIX41SettlementInstructions(msg SettlementInstructions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
