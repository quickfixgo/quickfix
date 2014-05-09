package fix40

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

func Crack(msg message.Message, sessionID quickfix.SessionID, router MessageRouter) errors.MessageRejectError {

	msgType := &field.MsgTypeField{}
	switch msg.Header.Get(msgType); msgType.Value {
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
	return errors.InvalidMessageType()
}

type MessageRouter interface {
	OnFIX40Heartbeat(msg Heartbeat, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40TestRequest(msg TestRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40ResendRequest(msg ResendRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40Reject(msg Reject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40SequenceReset(msg SequenceReset, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40Logout(msg Logout, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40IndicationofInterest(msg IndicationofInterest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40Advertisement(msg Advertisement, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40Logon(msg Logon, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40News(msg News, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40Email(msg Email, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40Allocation(msg Allocation, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40ListExecute(msg ListExecute, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40ListStatus(msg ListStatus, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40AllocationACK(msg AllocationACK, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX40Quote(msg Quote, sessionID quickfix.SessionID) errors.MessageRejectError
}
type FIX40MessageCracker struct{}

//OnFIX40Heartbeat is a Callback for Heartbeat messages.
func (c *FIX40MessageCracker) OnFIX40Heartbeat(msg Heartbeat, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40TestRequest is a Callback for TestRequest messages.
func (c *FIX40MessageCracker) OnFIX40TestRequest(msg TestRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40ResendRequest is a Callback for ResendRequest messages.
func (c *FIX40MessageCracker) OnFIX40ResendRequest(msg ResendRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40Reject is a Callback for Reject messages.
func (c *FIX40MessageCracker) OnFIX40Reject(msg Reject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40SequenceReset is a Callback for SequenceReset messages.
func (c *FIX40MessageCracker) OnFIX40SequenceReset(msg SequenceReset, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40Logout is a Callback for Logout messages.
func (c *FIX40MessageCracker) OnFIX40Logout(msg Logout, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40IndicationofInterest is a Callback for IndicationofInterest messages.
func (c *FIX40MessageCracker) OnFIX40IndicationofInterest(msg IndicationofInterest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40Advertisement is a Callback for Advertisement messages.
func (c *FIX40MessageCracker) OnFIX40Advertisement(msg Advertisement, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40ExecutionReport is a Callback for ExecutionReport messages.
func (c *FIX40MessageCracker) OnFIX40ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40OrderCancelReject is a Callback for OrderCancelReject messages.
func (c *FIX40MessageCracker) OnFIX40OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40Logon is a Callback for Logon messages.
func (c *FIX40MessageCracker) OnFIX40Logon(msg Logon, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40News is a Callback for News messages.
func (c *FIX40MessageCracker) OnFIX40News(msg News, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40Email is a Callback for Email messages.
func (c *FIX40MessageCracker) OnFIX40Email(msg Email, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40NewOrderSingle is a Callback for NewOrderSingle messages.
func (c *FIX40MessageCracker) OnFIX40NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40NewOrderList is a Callback for NewOrderList messages.
func (c *FIX40MessageCracker) OnFIX40NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40OrderCancelRequest is a Callback for OrderCancelRequest messages.
func (c *FIX40MessageCracker) OnFIX40OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40OrderCancelReplaceRequest is a Callback for OrderCancelReplaceRequest messages.
func (c *FIX40MessageCracker) OnFIX40OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40OrderStatusRequest is a Callback for OrderStatusRequest messages.
func (c *FIX40MessageCracker) OnFIX40OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40Allocation is a Callback for Allocation messages.
func (c *FIX40MessageCracker) OnFIX40Allocation(msg Allocation, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40ListCancelRequest is a Callback for ListCancelRequest messages.
func (c *FIX40MessageCracker) OnFIX40ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40ListExecute is a Callback for ListExecute messages.
func (c *FIX40MessageCracker) OnFIX40ListExecute(msg ListExecute, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40ListStatusRequest is a Callback for ListStatusRequest messages.
func (c *FIX40MessageCracker) OnFIX40ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40ListStatus is a Callback for ListStatus messages.
func (c *FIX40MessageCracker) OnFIX40ListStatus(msg ListStatus, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40AllocationACK is a Callback for AllocationACK messages.
func (c *FIX40MessageCracker) OnFIX40AllocationACK(msg AllocationACK, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40DontKnowTrade is a Callback for DontKnowTrade messages.
func (c *FIX40MessageCracker) OnFIX40DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40QuoteRequest is a Callback for QuoteRequest messages.
func (c *FIX40MessageCracker) OnFIX40QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX40Quote is a Callback for Quote messages.
func (c *FIX40MessageCracker) OnFIX40Quote(msg Quote, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
