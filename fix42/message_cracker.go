package fix42

import (
	"github.com/cbusbey/quickfixgo"
)

func Crack(msg quickfixgo.Message, sessionID quickfixgo.SessionID, router MessageRouter) quickfixgo.MessageReject {

	msgType := new(quickfixgo.MsgType)
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
	return quickfixgo.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX42Heartbeat(msg Heartbeat, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42TestRequest(msg TestRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42ResendRequest(msg ResendRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42Reject(msg Reject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42SequenceReset(msg SequenceReset, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42Logout(msg Logout, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42IndicationofInterest(msg IndicationofInterest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42Advertisement(msg Advertisement, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42ExecutionReport(msg ExecutionReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42OrderCancelReject(msg OrderCancelReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42Logon(msg Logon, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42News(msg News, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42Email(msg Email, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42NewOrderSingle(msg NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42NewOrderList(msg NewOrderList, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42OrderCancelRequest(msg OrderCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42OrderStatusRequest(msg OrderStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42Allocation(msg Allocation, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42ListCancelRequest(msg ListCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42ListExecute(msg ListExecute, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42ListStatusRequest(msg ListStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42ListStatus(msg ListStatus, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42AllocationACK(msg AllocationACK, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42DontKnowTrade(msg DontKnowTrade, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42QuoteRequest(msg QuoteRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42Quote(msg Quote, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42SettlementInstructions(msg SettlementInstructions, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42MarketDataRequest(msg MarketDataRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42QuoteCancel(msg QuoteCancel, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42QuoteAcknowledgement(msg QuoteAcknowledgement, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42SecurityDefinition(msg SecurityDefinition, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42SecurityStatus(msg SecurityStatus, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42TradingSessionStatus(msg TradingSessionStatus, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42MassQuote(msg MassQuote, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42BusinessMessageReject(msg BusinessMessageReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42BidRequest(msg BidRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42BidResponse(msg BidResponse, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX42ListStrikePrice(msg ListStrikePrice, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
}
type FIX42MessageCracker struct{}

func (c *FIX42MessageCracker) OnFIX42Heartbeat(msg Heartbeat, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42TestRequest(msg TestRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42ResendRequest(msg ResendRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42Reject(msg Reject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42SequenceReset(msg SequenceReset, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42Logout(msg Logout, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42IndicationofInterest(msg IndicationofInterest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42Advertisement(msg Advertisement, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42ExecutionReport(msg ExecutionReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42OrderCancelReject(msg OrderCancelReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42Logon(msg Logon, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42News(msg News, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42Email(msg Email, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42NewOrderSingle(msg NewOrderSingle, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42NewOrderList(msg NewOrderList, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42OrderCancelRequest(msg OrderCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42OrderStatusRequest(msg OrderStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42Allocation(msg Allocation, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42ListCancelRequest(msg ListCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42ListExecute(msg ListExecute, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42ListStatusRequest(msg ListStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42ListStatus(msg ListStatus, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42AllocationACK(msg AllocationACK, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42DontKnowTrade(msg DontKnowTrade, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42QuoteRequest(msg QuoteRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42Quote(msg Quote, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42SettlementInstructions(msg SettlementInstructions, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42MarketDataRequest(msg MarketDataRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42MarketDataRequestReject(msg MarketDataRequestReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42QuoteCancel(msg QuoteCancel, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42QuoteStatusRequest(msg QuoteStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42QuoteAcknowledgement(msg QuoteAcknowledgement, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42SecurityDefinition(msg SecurityDefinition, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42SecurityStatusRequest(msg SecurityStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42SecurityStatus(msg SecurityStatus, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42TradingSessionStatus(msg TradingSessionStatus, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42MassQuote(msg MassQuote, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42BusinessMessageReject(msg BusinessMessageReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42BidRequest(msg BidRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42BidResponse(msg BidResponse, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX42MessageCracker) OnFIX42ListStrikePrice(msg ListStrikePrice, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
