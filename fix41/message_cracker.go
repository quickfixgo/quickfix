package fix41

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

func Crack(msg message.Message, sessionID quickfix.SessionID, router MessageRouter) errors.MessageRejectError {

	msgType := new(field.MsgType)
	switch msg.Header.Get(msgType); msgType.Value {
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
	return errors.InvalidMessageType()
}

type MessageRouter interface {
	OnFIX41Heartbeat(msg Heartbeat, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41TestRequest(msg TestRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41ResendRequest(msg ResendRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41Reject(msg Reject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41SequenceReset(msg SequenceReset, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41Logout(msg Logout, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41IndicationofInterest(msg IndicationofInterest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41Advertisement(msg Advertisement, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41Logon(msg Logon, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41News(msg News, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41Email(msg Email, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41Allocation(msg Allocation, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41ListExecute(msg ListExecute, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41ListStatus(msg ListStatus, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41AllocationACK(msg AllocationACK, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41Quote(msg Quote, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX41SettlementInstructions(msg SettlementInstructions, sessionID quickfix.SessionID) errors.MessageRejectError
}
type FIX41MessageCracker struct{}

func (c *FIX41MessageCracker) OnFIX41Heartbeat(msg Heartbeat, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41TestRequest(msg TestRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41ResendRequest(msg ResendRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41Reject(msg Reject, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41SequenceReset(msg SequenceReset, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41Logout(msg Logout, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41IndicationofInterest(msg IndicationofInterest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41Advertisement(msg Advertisement, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41ExecutionReport(msg ExecutionReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41OrderCancelReject(msg OrderCancelReject, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41Logon(msg Logon, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41News(msg News, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41Email(msg Email, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41NewOrderSingle(msg NewOrderSingle, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41NewOrderList(msg NewOrderList, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41OrderCancelRequest(msg OrderCancelRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41OrderStatusRequest(msg OrderStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41Allocation(msg Allocation, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41ListCancelRequest(msg ListCancelRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41ListExecute(msg ListExecute, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41ListStatusRequest(msg ListStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41ListStatus(msg ListStatus, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41AllocationACK(msg AllocationACK, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41DontKnowTrade(msg DontKnowTrade, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41QuoteRequest(msg QuoteRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41Quote(msg Quote, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX41MessageCracker) OnFIX41SettlementInstructions(msg SettlementInstructions, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
