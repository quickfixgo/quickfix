package fix40

import (
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session"
	"github.com/cbusbey/quickfixgo/tag"
)

func Crack(msg message.Message, sessionID session.ID, router MessageRouter) reject.MessageReject {
	switch msgType, _ := msg.Header().StringValue(tag.MsgType); msgType {
	case "0":
		return router.OnFIX40Heartbeat(Heartbeat{msg}, sessionID)
	case "1":
		return router.OnFIX40TestRequest(TestRequest{msg}, sessionID)
	case "2":
		return router.OnFIX40ResendRequest(ResendRequest{msg}, sessionID)
	case "3":
		return router.OnFIX40Reject(Reject{msg}, sessionID)
	case "4":
		return router.OnFIX40SequenceReset(SequenceReset{msg}, sessionID)
	case "5":
		return router.OnFIX40Logout(Logout{msg}, sessionID)
	case "6":
		return router.OnFIX40IndicationofInterest(IndicationofInterest{msg}, sessionID)
	case "7":
		return router.OnFIX40Advertisement(Advertisement{msg}, sessionID)
	case "8":
		return router.OnFIX40ExecutionReport(ExecutionReport{msg}, sessionID)
	case "9":
		return router.OnFIX40OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "A":
		return router.OnFIX40Logon(Logon{msg}, sessionID)
	case "B":
		return router.OnFIX40News(News{msg}, sessionID)
	case "C":
		return router.OnFIX40Email(Email{msg}, sessionID)
	case "D":
		return router.OnFIX40NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "E":
		return router.OnFIX40NewOrderList(NewOrderList{msg}, sessionID)
	case "F":
		return router.OnFIX40OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "G":
		return router.OnFIX40OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "H":
		return router.OnFIX40OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "J":
		return router.OnFIX40Allocation(Allocation{msg}, sessionID)
	case "K":
		return router.OnFIX40ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "L":
		return router.OnFIX40ListExecute(ListExecute{msg}, sessionID)
	case "M":
		return router.OnFIX40ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "N":
		return router.OnFIX40ListStatus(ListStatus{msg}, sessionID)
	case "P":
		return router.OnFIX40AllocationACK(AllocationACK{msg}, sessionID)
	case "Q":
		return router.OnFIX40DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "R":
		return router.OnFIX40QuoteRequest(QuoteRequest{msg}, sessionID)
	case "S":
		return router.OnFIX40Quote(Quote{msg}, sessionID)
	}
	return reject.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX40Heartbeat(msg Heartbeat, sessionID session.ID) reject.MessageReject
	OnFIX40TestRequest(msg TestRequest, sessionID session.ID) reject.MessageReject
	OnFIX40ResendRequest(msg ResendRequest, sessionID session.ID) reject.MessageReject
	OnFIX40Reject(msg Reject, sessionID session.ID) reject.MessageReject
	OnFIX40SequenceReset(msg SequenceReset, sessionID session.ID) reject.MessageReject
	OnFIX40Logout(msg Logout, sessionID session.ID) reject.MessageReject
	OnFIX40IndicationofInterest(msg IndicationofInterest, sessionID session.ID) reject.MessageReject
	OnFIX40Advertisement(msg Advertisement, sessionID session.ID) reject.MessageReject
	OnFIX40ExecutionReport(msg ExecutionReport, sessionID session.ID) reject.MessageReject
	OnFIX40OrderCancelReject(msg OrderCancelReject, sessionID session.ID) reject.MessageReject
	OnFIX40Logon(msg Logon, sessionID session.ID) reject.MessageReject
	OnFIX40News(msg News, sessionID session.ID) reject.MessageReject
	OnFIX40Email(msg Email, sessionID session.ID) reject.MessageReject
	OnFIX40NewOrderSingle(msg NewOrderSingle, sessionID session.ID) reject.MessageReject
	OnFIX40NewOrderList(msg NewOrderList, sessionID session.ID) reject.MessageReject
	OnFIX40OrderCancelRequest(msg OrderCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX40OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID session.ID) reject.MessageReject
	OnFIX40OrderStatusRequest(msg OrderStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX40Allocation(msg Allocation, sessionID session.ID) reject.MessageReject
	OnFIX40ListCancelRequest(msg ListCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX40ListExecute(msg ListExecute, sessionID session.ID) reject.MessageReject
	OnFIX40ListStatusRequest(msg ListStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX40ListStatus(msg ListStatus, sessionID session.ID) reject.MessageReject
	OnFIX40AllocationACK(msg AllocationACK, sessionID session.ID) reject.MessageReject
	OnFIX40DontKnowTrade(msg DontKnowTrade, sessionID session.ID) reject.MessageReject
	OnFIX40QuoteRequest(msg QuoteRequest, sessionID session.ID) reject.MessageReject
	OnFIX40Quote(msg Quote, sessionID session.ID) reject.MessageReject
}
type FIX40MessageCracker struct{}

func (c *FIX40MessageCracker) OnFIX40Heartbeat(msg Heartbeat, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40TestRequest(msg TestRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40ResendRequest(msg ResendRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40Reject(msg Reject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40SequenceReset(msg SequenceReset, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40Logout(msg Logout, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40IndicationofInterest(msg IndicationofInterest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40Advertisement(msg Advertisement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40ExecutionReport(msg ExecutionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40OrderCancelReject(msg OrderCancelReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40Logon(msg Logon, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40News(msg News, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40Email(msg Email, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40NewOrderSingle(msg NewOrderSingle, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40NewOrderList(msg NewOrderList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40OrderCancelRequest(msg OrderCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40OrderStatusRequest(msg OrderStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40Allocation(msg Allocation, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40ListCancelRequest(msg ListCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40ListExecute(msg ListExecute, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40ListStatusRequest(msg ListStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40ListStatus(msg ListStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40AllocationACK(msg AllocationACK, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40DontKnowTrade(msg DontKnowTrade, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40QuoteRequest(msg QuoteRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX40MessageCracker) OnFIX40Quote(msg Quote, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
