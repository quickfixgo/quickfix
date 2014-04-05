package fix41

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

func Crack(msg quickfixgo.Message, sessionID quickfixgo.SessionID, router MessageRouter) quickfixgo.MessageReject {

	msgType := new(field.MsgType)
	switch msg.Header.Get(msgType); msgType.Value {
	case "2":
		return router.OnFIX41ResendRequest(ResendRequest{msg}, sessionID)
	case "4":
		return router.OnFIX41SequenceReset(SequenceReset{msg}, sessionID)
	case "5":
		return router.OnFIX41Logout(Logout{msg}, sessionID)
	case "E":
		return router.OnFIX41NewOrderList(NewOrderList{msg}, sessionID)
	case "T":
		return router.OnFIX41SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "B":
		return router.OnFIX41News(News{msg}, sessionID)
	case "Q":
		return router.OnFIX41DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "8":
		return router.OnFIX41ExecutionReport(ExecutionReport{msg}, sessionID)
	case "N":
		return router.OnFIX41ListStatus(ListStatus{msg}, sessionID)
	case "S":
		return router.OnFIX41Quote(Quote{msg}, sessionID)
	case "0":
		return router.OnFIX41Heartbeat(Heartbeat{msg}, sessionID)
	case "D":
		return router.OnFIX41NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "M":
		return router.OnFIX41ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "P":
		return router.OnFIX41AllocationACK(AllocationACK{msg}, sessionID)
	case "7":
		return router.OnFIX41Advertisement(Advertisement{msg}, sessionID)
	case "9":
		return router.OnFIX41OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "C":
		return router.OnFIX41Email(Email{msg}, sessionID)
	case "J":
		return router.OnFIX41Allocation(Allocation{msg}, sessionID)
	case "L":
		return router.OnFIX41ListExecute(ListExecute{msg}, sessionID)
	case "G":
		return router.OnFIX41OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "H":
		return router.OnFIX41OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "R":
		return router.OnFIX41QuoteRequest(QuoteRequest{msg}, sessionID)
	case "1":
		return router.OnFIX41TestRequest(TestRequest{msg}, sessionID)
	case "3":
		return router.OnFIX41Reject(Reject{msg}, sessionID)
	case "6":
		return router.OnFIX41IndicationofInterest(IndicationofInterest{msg}, sessionID)
	case "A":
		return router.OnFIX41Logon(Logon{msg}, sessionID)
	case "F":
		return router.OnFIX41OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "K":
		return router.OnFIX41ListCancelRequest(ListCancelRequest{msg}, sessionID)
	}
	return quickfixgo.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX41TestRequest(msg TestRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41Reject(msg Reject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41IndicationofInterest(msg IndicationofInterest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41Logon(msg Logon, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41OrderCancelRequest(msg OrderCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41ListCancelRequest(msg ListCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41ResendRequest(msg ResendRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41SequenceReset(msg SequenceReset, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41Logout(msg Logout, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41NewOrderList(msg NewOrderList, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41SettlementInstructions(msg SettlementInstructions, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41News(msg News, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41DontKnowTrade(msg DontKnowTrade, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41ExecutionReport(msg ExecutionReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41ListStatus(msg ListStatus, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41Quote(msg Quote, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41Heartbeat(msg Heartbeat, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41NewOrderSingle(msg NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41ListStatusRequest(msg ListStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41AllocationACK(msg AllocationACK, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41Advertisement(msg Advertisement, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41OrderCancelReject(msg OrderCancelReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41Email(msg Email, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41Allocation(msg Allocation, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41ListExecute(msg ListExecute, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41OrderStatusRequest(msg OrderStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX41QuoteRequest(msg QuoteRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
}
type FIX41MessageCracker struct{}

func (c *FIX41MessageCracker) OnFIX41News(msg News, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41DontKnowTrade(msg DontKnowTrade, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41ExecutionReport(msg ExecutionReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41ListStatus(msg ListStatus, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41Quote(msg Quote, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41Heartbeat(msg Heartbeat, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41NewOrderSingle(msg NewOrderSingle, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41ListStatusRequest(msg ListStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41AllocationACK(msg AllocationACK, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41Advertisement(msg Advertisement, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41OrderCancelReject(msg OrderCancelReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41Email(msg Email, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41Allocation(msg Allocation, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41ListExecute(msg ListExecute, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41OrderStatusRequest(msg OrderStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41QuoteRequest(msg QuoteRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41TestRequest(msg TestRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41Reject(msg Reject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41IndicationofInterest(msg IndicationofInterest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41Logon(msg Logon, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41OrderCancelRequest(msg OrderCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41ListCancelRequest(msg ListCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41ResendRequest(msg ResendRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41SequenceReset(msg SequenceReset, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41Logout(msg Logout, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41NewOrderList(msg NewOrderList, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX41MessageCracker) OnFIX41SettlementInstructions(msg SettlementInstructions, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
