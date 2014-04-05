package fix40

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

func Crack(msg quickfixgo.Message, sessionID quickfixgo.SessionID, router MessageRouter) quickfixgo.MessageReject {

	msgType := new(field.MsgType)
	switch msg.Header.Get(msgType); msgType.Value {
	case "0":
		return router.OnFIX40Heartbeat(Heartbeat{msg}, sessionID)
	case "4":
		return router.OnFIX40SequenceReset(SequenceReset{msg}, sessionID)
	case "8":
		return router.OnFIX40ExecutionReport(ExecutionReport{msg}, sessionID)
	case "D":
		return router.OnFIX40NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "M":
		return router.OnFIX40ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "C":
		return router.OnFIX40Email(Email{msg}, sessionID)
	case "H":
		return router.OnFIX40OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "Q":
		return router.OnFIX40DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "5":
		return router.OnFIX40Logout(Logout{msg}, sessionID)
	case "7":
		return router.OnFIX40Advertisement(Advertisement{msg}, sessionID)
	case "L":
		return router.OnFIX40ListExecute(ListExecute{msg}, sessionID)
	case "3":
		return router.OnFIX40Reject(Reject{msg}, sessionID)
	case "P":
		return router.OnFIX40AllocationACK(AllocationACK{msg}, sessionID)
	case "S":
		return router.OnFIX40Quote(Quote{msg}, sessionID)
	case "A":
		return router.OnFIX40Logon(Logon{msg}, sessionID)
	case "E":
		return router.OnFIX40NewOrderList(NewOrderList{msg}, sessionID)
	case "F":
		return router.OnFIX40OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "1":
		return router.OnFIX40TestRequest(TestRequest{msg}, sessionID)
	case "B":
		return router.OnFIX40News(News{msg}, sessionID)
	case "J":
		return router.OnFIX40Allocation(Allocation{msg}, sessionID)
	case "N":
		return router.OnFIX40ListStatus(ListStatus{msg}, sessionID)
	case "R":
		return router.OnFIX40QuoteRequest(QuoteRequest{msg}, sessionID)
	case "9":
		return router.OnFIX40OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "G":
		return router.OnFIX40OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "K":
		return router.OnFIX40ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "2":
		return router.OnFIX40ResendRequest(ResendRequest{msg}, sessionID)
	case "6":
		return router.OnFIX40IndicationofInterest(IndicationofInterest{msg}, sessionID)
	}
	return quickfixgo.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX40TestRequest(msg TestRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40News(msg News, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40Allocation(msg Allocation, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40ListStatus(msg ListStatus, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40QuoteRequest(msg QuoteRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40OrderCancelReject(msg OrderCancelReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40ListCancelRequest(msg ListCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40ResendRequest(msg ResendRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40IndicationofInterest(msg IndicationofInterest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40Heartbeat(msg Heartbeat, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40SequenceReset(msg SequenceReset, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40ExecutionReport(msg ExecutionReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40NewOrderSingle(msg NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40ListStatusRequest(msg ListStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40Email(msg Email, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40OrderStatusRequest(msg OrderStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40DontKnowTrade(msg DontKnowTrade, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40Logout(msg Logout, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40Advertisement(msg Advertisement, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40ListExecute(msg ListExecute, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40Reject(msg Reject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40AllocationACK(msg AllocationACK, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40Quote(msg Quote, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40Logon(msg Logon, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40NewOrderList(msg NewOrderList, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX40OrderCancelRequest(msg OrderCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
}
type FIX40MessageCracker struct{}

func (c *FIX40MessageCracker) OnFIX40ResendRequest(msg ResendRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40IndicationofInterest(msg IndicationofInterest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Heartbeat(msg Heartbeat, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40SequenceReset(msg SequenceReset, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40ExecutionReport(msg ExecutionReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40NewOrderSingle(msg NewOrderSingle, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40ListStatusRequest(msg ListStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Email(msg Email, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40OrderStatusRequest(msg OrderStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40DontKnowTrade(msg DontKnowTrade, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Logout(msg Logout, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Advertisement(msg Advertisement, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40ListExecute(msg ListExecute, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Reject(msg Reject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40AllocationACK(msg AllocationACK, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Quote(msg Quote, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Logon(msg Logon, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40NewOrderList(msg NewOrderList, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40OrderCancelRequest(msg OrderCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40TestRequest(msg TestRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40News(msg News, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40Allocation(msg Allocation, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40ListStatus(msg ListStatus, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40QuoteRequest(msg QuoteRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40OrderCancelReject(msg OrderCancelReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX40MessageCracker) OnFIX40ListCancelRequest(msg ListCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
