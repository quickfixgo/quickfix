package fix42

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

func (c *FIX42MessageCracker) OnFIX42Heartbeat(msg Heartbeat, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42TestRequest(msg TestRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42ResendRequest(msg ResendRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42Reject(msg Reject, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42SequenceReset(msg SequenceReset, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42Logout(msg Logout, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42IndicationofInterest(msg IndicationofInterest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42Advertisement(msg Advertisement, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42ExecutionReport(msg ExecutionReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42OrderCancelReject(msg OrderCancelReject, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42Logon(msg Logon, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42News(msg News, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42Email(msg Email, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42NewOrderSingle(msg NewOrderSingle, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42NewOrderList(msg NewOrderList, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42OrderCancelRequest(msg OrderCancelRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42OrderStatusRequest(msg OrderStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42Allocation(msg Allocation, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42ListCancelRequest(msg ListCancelRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42ListExecute(msg ListExecute, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42ListStatusRequest(msg ListStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42ListStatus(msg ListStatus, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42AllocationACK(msg AllocationACK, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42DontKnowTrade(msg DontKnowTrade, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42QuoteRequest(msg QuoteRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42Quote(msg Quote, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42SettlementInstructions(msg SettlementInstructions, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42MarketDataRequest(msg MarketDataRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42MarketDataRequestReject(msg MarketDataRequestReject, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42QuoteCancel(msg QuoteCancel, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42QuoteStatusRequest(msg QuoteStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42QuoteAcknowledgement(msg QuoteAcknowledgement, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42SecurityDefinition(msg SecurityDefinition, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42SecurityStatusRequest(msg SecurityStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42SecurityStatus(msg SecurityStatus, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42TradingSessionStatus(msg TradingSessionStatus, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42MassQuote(msg MassQuote, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42BusinessMessageReject(msg BusinessMessageReject, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42BidRequest(msg BidRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42BidResponse(msg BidResponse, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX42MessageCracker) OnFIX42ListStrikePrice(msg ListStrikePrice, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
