package fix42

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
		return router.OnFIX42Heartbeat(Heartbeat{msg}, sessionID)
	case "1":
		return router.OnFIX42TestRequest(TestRequest{msg}, sessionID)
	case "2":
		return router.OnFIX42ResendRequest(ResendRequest{msg}, sessionID)
	case "3":
		return router.OnFIX42Reject(Reject{msg}, sessionID)
	case "4":
		return router.OnFIX42SequenceReset(SequenceReset{msg}, sessionID)
	case "5":
		return router.OnFIX42Logout(Logout{msg}, sessionID)
	case "6":
		return router.OnFIX42IndicationofInterest(IndicationofInterest{msg}, sessionID)
	case "7":
		return router.OnFIX42Advertisement(Advertisement{msg}, sessionID)
	case "8":
		return router.OnFIX42ExecutionReport(ExecutionReport{msg}, sessionID)
	case "9":
		return router.OnFIX42OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "A":
		return router.OnFIX42Logon(Logon{msg}, sessionID)
	case "B":
		return router.OnFIX42News(News{msg}, sessionID)
	case "C":
		return router.OnFIX42Email(Email{msg}, sessionID)
	case "D":
		return router.OnFIX42NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "E":
		return router.OnFIX42NewOrderList(NewOrderList{msg}, sessionID)
	case "F":
		return router.OnFIX42OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "G":
		return router.OnFIX42OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "H":
		return router.OnFIX42OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "J":
		return router.OnFIX42Allocation(Allocation{msg}, sessionID)
	case "K":
		return router.OnFIX42ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "L":
		return router.OnFIX42ListExecute(ListExecute{msg}, sessionID)
	case "M":
		return router.OnFIX42ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "N":
		return router.OnFIX42ListStatus(ListStatus{msg}, sessionID)
	case "P":
		return router.OnFIX42AllocationACK(AllocationACK{msg}, sessionID)
	case "Q":
		return router.OnFIX42DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "R":
		return router.OnFIX42QuoteRequest(QuoteRequest{msg}, sessionID)
	case "S":
		return router.OnFIX42Quote(Quote{msg}, sessionID)
	case "T":
		return router.OnFIX42SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "V":
		return router.OnFIX42MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "W":
		return router.OnFIX42MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "X":
		return router.OnFIX42MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "Y":
		return router.OnFIX42MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "Z":
		return router.OnFIX42QuoteCancel(QuoteCancel{msg}, sessionID)
	case "a":
		return router.OnFIX42QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "b":
		return router.OnFIX42QuoteAcknowledgement(QuoteAcknowledgement{msg}, sessionID)
	case "c":
		return router.OnFIX42SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "d":
		return router.OnFIX42SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "e":
		return router.OnFIX42SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "f":
		return router.OnFIX42SecurityStatus(SecurityStatus{msg}, sessionID)
	case "g":
		return router.OnFIX42TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "h":
		return router.OnFIX42TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "i":
		return router.OnFIX42MassQuote(MassQuote{msg}, sessionID)
	case "j":
		return router.OnFIX42BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "k":
		return router.OnFIX42BidRequest(BidRequest{msg}, sessionID)
	case "l":
		return router.OnFIX42BidResponse(BidResponse{msg}, sessionID)
	case "m":
		return router.OnFIX42ListStrikePrice(ListStrikePrice{msg}, sessionID)
	}
	return errors.InvalidMessageType()
}

type MessageRouter interface {
	OnFIX42Heartbeat(msg Heartbeat, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42TestRequest(msg TestRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42ResendRequest(msg ResendRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42Reject(msg Reject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42SequenceReset(msg SequenceReset, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42Logout(msg Logout, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42IndicationofInterest(msg IndicationofInterest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42Advertisement(msg Advertisement, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42Logon(msg Logon, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42News(msg News, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42Email(msg Email, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42Allocation(msg Allocation, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42ListExecute(msg ListExecute, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42ListStatus(msg ListStatus, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42AllocationACK(msg AllocationACK, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42Quote(msg Quote, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42SettlementInstructions(msg SettlementInstructions, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42MarketDataRequest(msg MarketDataRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42QuoteCancel(msg QuoteCancel, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42QuoteAcknowledgement(msg QuoteAcknowledgement, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42SecurityDefinition(msg SecurityDefinition, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42SecurityStatus(msg SecurityStatus, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42TradingSessionStatus(msg TradingSessionStatus, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42MassQuote(msg MassQuote, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42BusinessMessageReject(msg BusinessMessageReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42BidRequest(msg BidRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42BidResponse(msg BidResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX42ListStrikePrice(msg ListStrikePrice, sessionID quickfix.SessionID) errors.MessageRejectError
}
type FIX42MessageCracker struct{}

//OnFIX42Heartbeat is a Callback for Heartbeat messages.
func (c *FIX42MessageCracker) OnFIX42Heartbeat(msg Heartbeat, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42TestRequest is a Callback for TestRequest messages.
func (c *FIX42MessageCracker) OnFIX42TestRequest(msg TestRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42ResendRequest is a Callback for ResendRequest messages.
func (c *FIX42MessageCracker) OnFIX42ResendRequest(msg ResendRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42Reject is a Callback for Reject messages.
func (c *FIX42MessageCracker) OnFIX42Reject(msg Reject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42SequenceReset is a Callback for SequenceReset messages.
func (c *FIX42MessageCracker) OnFIX42SequenceReset(msg SequenceReset, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42Logout is a Callback for Logout messages.
func (c *FIX42MessageCracker) OnFIX42Logout(msg Logout, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42IndicationofInterest is a Callback for IndicationofInterest messages.
func (c *FIX42MessageCracker) OnFIX42IndicationofInterest(msg IndicationofInterest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42Advertisement is a Callback for Advertisement messages.
func (c *FIX42MessageCracker) OnFIX42Advertisement(msg Advertisement, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42ExecutionReport is a Callback for ExecutionReport messages.
func (c *FIX42MessageCracker) OnFIX42ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42OrderCancelReject is a Callback for OrderCancelReject messages.
func (c *FIX42MessageCracker) OnFIX42OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42Logon is a Callback for Logon messages.
func (c *FIX42MessageCracker) OnFIX42Logon(msg Logon, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42News is a Callback for News messages.
func (c *FIX42MessageCracker) OnFIX42News(msg News, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42Email is a Callback for Email messages.
func (c *FIX42MessageCracker) OnFIX42Email(msg Email, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42NewOrderSingle is a Callback for NewOrderSingle messages.
func (c *FIX42MessageCracker) OnFIX42NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42NewOrderList is a Callback for NewOrderList messages.
func (c *FIX42MessageCracker) OnFIX42NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42OrderCancelRequest is a Callback for OrderCancelRequest messages.
func (c *FIX42MessageCracker) OnFIX42OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42OrderCancelReplaceRequest is a Callback for OrderCancelReplaceRequest messages.
func (c *FIX42MessageCracker) OnFIX42OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42OrderStatusRequest is a Callback for OrderStatusRequest messages.
func (c *FIX42MessageCracker) OnFIX42OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42Allocation is a Callback for Allocation messages.
func (c *FIX42MessageCracker) OnFIX42Allocation(msg Allocation, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42ListCancelRequest is a Callback for ListCancelRequest messages.
func (c *FIX42MessageCracker) OnFIX42ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42ListExecute is a Callback for ListExecute messages.
func (c *FIX42MessageCracker) OnFIX42ListExecute(msg ListExecute, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42ListStatusRequest is a Callback for ListStatusRequest messages.
func (c *FIX42MessageCracker) OnFIX42ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42ListStatus is a Callback for ListStatus messages.
func (c *FIX42MessageCracker) OnFIX42ListStatus(msg ListStatus, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42AllocationACK is a Callback for AllocationACK messages.
func (c *FIX42MessageCracker) OnFIX42AllocationACK(msg AllocationACK, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42DontKnowTrade is a Callback for DontKnowTrade messages.
func (c *FIX42MessageCracker) OnFIX42DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42QuoteRequest is a Callback for QuoteRequest messages.
func (c *FIX42MessageCracker) OnFIX42QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42Quote is a Callback for Quote messages.
func (c *FIX42MessageCracker) OnFIX42Quote(msg Quote, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42SettlementInstructions is a Callback for SettlementInstructions messages.
func (c *FIX42MessageCracker) OnFIX42SettlementInstructions(msg SettlementInstructions, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42MarketDataRequest is a Callback for MarketDataRequest messages.
func (c *FIX42MessageCracker) OnFIX42MarketDataRequest(msg MarketDataRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42MarketDataSnapshotFullRefresh is a Callback for MarketDataSnapshotFullRefresh messages.
func (c *FIX42MessageCracker) OnFIX42MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42MarketDataIncrementalRefresh is a Callback for MarketDataIncrementalRefresh messages.
func (c *FIX42MessageCracker) OnFIX42MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42MarketDataRequestReject is a Callback for MarketDataRequestReject messages.
func (c *FIX42MessageCracker) OnFIX42MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42QuoteCancel is a Callback for QuoteCancel messages.
func (c *FIX42MessageCracker) OnFIX42QuoteCancel(msg QuoteCancel, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42QuoteStatusRequest is a Callback for QuoteStatusRequest messages.
func (c *FIX42MessageCracker) OnFIX42QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42QuoteAcknowledgement is a Callback for QuoteAcknowledgement messages.
func (c *FIX42MessageCracker) OnFIX42QuoteAcknowledgement(msg QuoteAcknowledgement, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42SecurityDefinitionRequest is a Callback for SecurityDefinitionRequest messages.
func (c *FIX42MessageCracker) OnFIX42SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42SecurityDefinition is a Callback for SecurityDefinition messages.
func (c *FIX42MessageCracker) OnFIX42SecurityDefinition(msg SecurityDefinition, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42SecurityStatusRequest is a Callback for SecurityStatusRequest messages.
func (c *FIX42MessageCracker) OnFIX42SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42SecurityStatus is a Callback for SecurityStatus messages.
func (c *FIX42MessageCracker) OnFIX42SecurityStatus(msg SecurityStatus, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42TradingSessionStatusRequest is a Callback for TradingSessionStatusRequest messages.
func (c *FIX42MessageCracker) OnFIX42TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42TradingSessionStatus is a Callback for TradingSessionStatus messages.
func (c *FIX42MessageCracker) OnFIX42TradingSessionStatus(msg TradingSessionStatus, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42MassQuote is a Callback for MassQuote messages.
func (c *FIX42MessageCracker) OnFIX42MassQuote(msg MassQuote, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42BusinessMessageReject is a Callback for BusinessMessageReject messages.
func (c *FIX42MessageCracker) OnFIX42BusinessMessageReject(msg BusinessMessageReject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42BidRequest is a Callback for BidRequest messages.
func (c *FIX42MessageCracker) OnFIX42BidRequest(msg BidRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42BidResponse is a Callback for BidResponse messages.
func (c *FIX42MessageCracker) OnFIX42BidResponse(msg BidResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX42ListStrikePrice is a Callback for ListStrikePrice messages.
func (c *FIX42MessageCracker) OnFIX42ListStrikePrice(msg ListStrikePrice, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
