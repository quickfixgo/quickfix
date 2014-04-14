package fix40

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

func Crack(msg message.Message, sessionID quickfix.SessionID, router MessageRouter) message.MessageReject {

	msgType := new(field.MsgType)
	switch msg.Header.Get(msgType); msgType.Value {
	case "H":
		return router.OnFIX40OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "K":
		return router.OnFIX40ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "5":
		return router.OnFIX40Logout(Logout{msg}, sessionID)
	case "6":
		return router.OnFIX40IndicationofInterest(IndicationofInterest{msg}, sessionID)
	case "9":
		return router.OnFIX40OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "G":
		return router.OnFIX40OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "M":
		return router.OnFIX40ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "N":
		return router.OnFIX40ListStatus(ListStatus{msg}, sessionID)
	case "Q":
		return router.OnFIX40DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "0":
		return router.OnFIX40Heartbeat(Heartbeat{msg}, sessionID)
	case "A":
		return router.OnFIX40Logon(Logon{msg}, sessionID)
	case "P":
		return router.OnFIX40AllocationACK(AllocationACK{msg}, sessionID)
	case "S":
		return router.OnFIX40Quote(Quote{msg}, sessionID)
	case "J":
		return router.OnFIX40Allocation(Allocation{msg}, sessionID)
	case "C":
		return router.OnFIX40Email(Email{msg}, sessionID)
	case "D":
		return router.OnFIX40NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "F":
		return router.OnFIX40OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "B":
		return router.OnFIX40News(News{msg}, sessionID)
	case "L":
		return router.OnFIX40ListExecute(ListExecute{msg}, sessionID)
	case "R":
		return router.OnFIX40QuoteRequest(QuoteRequest{msg}, sessionID)
	case "1":
		return router.OnFIX40TestRequest(TestRequest{msg}, sessionID)
	case "3":
		return router.OnFIX40Reject(Reject{msg}, sessionID)
	case "E":
		return router.OnFIX40NewOrderList(NewOrderList{msg}, sessionID)
	case "2":
		return router.OnFIX40ResendRequest(ResendRequest{msg}, sessionID)
	case "4":
		return router.OnFIX40SequenceReset(SequenceReset{msg}, sessionID)
	case "7":
		return router.OnFIX40Advertisement(Advertisement{msg}, sessionID)
	case "8":
		return router.OnFIX40ExecutionReport(ExecutionReport{msg}, sessionID)
	}
	return message.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX40OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40Logout(msg Logout, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40IndicationofInterest(msg IndicationofInterest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40ListStatus(msg ListStatus, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40Heartbeat(msg Heartbeat, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40Logon(msg Logon, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40AllocationACK(msg AllocationACK, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40Quote(msg Quote, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40Allocation(msg Allocation, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40Email(msg Email, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40News(msg News, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40ListExecute(msg ListExecute, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40TestRequest(msg TestRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40Reject(msg Reject, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40ResendRequest(msg ResendRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40SequenceReset(msg SequenceReset, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40Advertisement(msg Advertisement, sessionID quickfix.SessionID) message.MessageReject
	OnFIX40ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) message.MessageReject
}
type FIX40MessageCracker struct{}

func (c *FIX40MessageCracker) OnFIX40OrderStatusRequest(msg OrderStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40ListCancelRequest(msg ListCancelRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Logout(msg Logout, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40IndicationofInterest(msg IndicationofInterest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40OrderCancelReject(msg OrderCancelReject, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40ListStatusRequest(msg ListStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40ListStatus(msg ListStatus, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40DontKnowTrade(msg DontKnowTrade, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Heartbeat(msg Heartbeat, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Logon(msg Logon, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40AllocationACK(msg AllocationACK, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Quote(msg Quote, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Allocation(msg Allocation, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Email(msg Email, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40NewOrderSingle(msg NewOrderSingle, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40OrderCancelRequest(msg OrderCancelRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40News(msg News, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40ListExecute(msg ListExecute, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40QuoteRequest(msg QuoteRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40TestRequest(msg TestRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Reject(msg Reject, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40NewOrderList(msg NewOrderList, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40ResendRequest(msg ResendRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40SequenceReset(msg SequenceReset, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Advertisement(msg Advertisement, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40ExecutionReport(msg ExecutionReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
