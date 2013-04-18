package fix42

import (
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session"
)

func Crack(msg message.Message, sessionID session.ID, router MessageRouter) reject.MessageReject {
	switch msgType, _ := msg.Header().StringValue(fix.MsgType); msgType {
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
	return reject.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX42Heartbeat(msg Heartbeat, sessionID session.ID) reject.MessageReject
	OnFIX42TestRequest(msg TestRequest, sessionID session.ID) reject.MessageReject
	OnFIX42ResendRequest(msg ResendRequest, sessionID session.ID) reject.MessageReject
	OnFIX42Reject(msg Reject, sessionID session.ID) reject.MessageReject
	OnFIX42SequenceReset(msg SequenceReset, sessionID session.ID) reject.MessageReject
	OnFIX42Logout(msg Logout, sessionID session.ID) reject.MessageReject
	OnFIX42IndicationofInterest(msg IndicationofInterest, sessionID session.ID) reject.MessageReject
	OnFIX42Advertisement(msg Advertisement, sessionID session.ID) reject.MessageReject
	OnFIX42ExecutionReport(msg ExecutionReport, sessionID session.ID) reject.MessageReject
	OnFIX42OrderCancelReject(msg OrderCancelReject, sessionID session.ID) reject.MessageReject
	OnFIX42Logon(msg Logon, sessionID session.ID) reject.MessageReject
	OnFIX42News(msg News, sessionID session.ID) reject.MessageReject
	OnFIX42Email(msg Email, sessionID session.ID) reject.MessageReject
	OnFIX42NewOrderSingle(msg NewOrderSingle, sessionID session.ID) reject.MessageReject
	OnFIX42NewOrderList(msg NewOrderList, sessionID session.ID) reject.MessageReject
	OnFIX42OrderCancelRequest(msg OrderCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX42OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID session.ID) reject.MessageReject
	OnFIX42OrderStatusRequest(msg OrderStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX42Allocation(msg Allocation, sessionID session.ID) reject.MessageReject
	OnFIX42ListCancelRequest(msg ListCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX42ListExecute(msg ListExecute, sessionID session.ID) reject.MessageReject
	OnFIX42ListStatusRequest(msg ListStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX42ListStatus(msg ListStatus, sessionID session.ID) reject.MessageReject
	OnFIX42AllocationACK(msg AllocationACK, sessionID session.ID) reject.MessageReject
	OnFIX42DontKnowTrade(msg DontKnowTrade, sessionID session.ID) reject.MessageReject
	OnFIX42QuoteRequest(msg QuoteRequest, sessionID session.ID) reject.MessageReject
	OnFIX42Quote(msg Quote, sessionID session.ID) reject.MessageReject
	OnFIX42SettlementInstructions(msg SettlementInstructions, sessionID session.ID) reject.MessageReject
	OnFIX42MarketDataRequest(msg MarketDataRequest, sessionID session.ID) reject.MessageReject
	OnFIX42MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID session.ID) reject.MessageReject
	OnFIX42MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID session.ID) reject.MessageReject
	OnFIX42MarketDataRequestReject(msg MarketDataRequestReject, sessionID session.ID) reject.MessageReject
	OnFIX42QuoteCancel(msg QuoteCancel, sessionID session.ID) reject.MessageReject
	OnFIX42QuoteStatusRequest(msg QuoteStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX42QuoteAcknowledgement(msg QuoteAcknowledgement, sessionID session.ID) reject.MessageReject
	OnFIX42SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID session.ID) reject.MessageReject
	OnFIX42SecurityDefinition(msg SecurityDefinition, sessionID session.ID) reject.MessageReject
	OnFIX42SecurityStatusRequest(msg SecurityStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX42SecurityStatus(msg SecurityStatus, sessionID session.ID) reject.MessageReject
	OnFIX42TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX42TradingSessionStatus(msg TradingSessionStatus, sessionID session.ID) reject.MessageReject
	OnFIX42MassQuote(msg MassQuote, sessionID session.ID) reject.MessageReject
	OnFIX42BusinessMessageReject(msg BusinessMessageReject, sessionID session.ID) reject.MessageReject
	OnFIX42BidRequest(msg BidRequest, sessionID session.ID) reject.MessageReject
	OnFIX42BidResponse(msg BidResponse, sessionID session.ID) reject.MessageReject
	OnFIX42ListStrikePrice(msg ListStrikePrice, sessionID session.ID) reject.MessageReject
}
type FIX42MessageCracker struct{}

func (c *FIX42MessageCracker) OnFIX42Heartbeat(msg Heartbeat, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42TestRequest(msg TestRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42ResendRequest(msg ResendRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42Reject(msg Reject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42SequenceReset(msg SequenceReset, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42Logout(msg Logout, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42IndicationofInterest(msg IndicationofInterest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42Advertisement(msg Advertisement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42ExecutionReport(msg ExecutionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42OrderCancelReject(msg OrderCancelReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42Logon(msg Logon, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42News(msg News, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42Email(msg Email, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42NewOrderSingle(msg NewOrderSingle, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42NewOrderList(msg NewOrderList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42OrderCancelRequest(msg OrderCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42OrderStatusRequest(msg OrderStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42Allocation(msg Allocation, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42ListCancelRequest(msg ListCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42ListExecute(msg ListExecute, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42ListStatusRequest(msg ListStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42ListStatus(msg ListStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42AllocationACK(msg AllocationACK, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42DontKnowTrade(msg DontKnowTrade, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42QuoteRequest(msg QuoteRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42Quote(msg Quote, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42SettlementInstructions(msg SettlementInstructions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42MarketDataRequest(msg MarketDataRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42MarketDataRequestReject(msg MarketDataRequestReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42QuoteCancel(msg QuoteCancel, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42QuoteStatusRequest(msg QuoteStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42QuoteAcknowledgement(msg QuoteAcknowledgement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42SecurityDefinition(msg SecurityDefinition, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42SecurityStatusRequest(msg SecurityStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42SecurityStatus(msg SecurityStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42TradingSessionStatus(msg TradingSessionStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42MassQuote(msg MassQuote, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42BusinessMessageReject(msg BusinessMessageReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42BidRequest(msg BidRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42BidResponse(msg BidResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX42MessageCracker) OnFIX42ListStrikePrice(msg ListStrikePrice, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
