package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

func Crack(msg quickfixgo.Message, sessionID quickfixgo.SessionID, router MessageRouter) quickfixgo.MessageReject {

	msgType := new(field.MsgType)
	switch msg.Header.Get(msgType); msgType.Value {
	case "D":
		return router.OnFIX43NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "P":
		return router.OnFIX43AllocationAck(AllocationAck{msg}, sessionID)
	case "d":
		return router.OnFIX43SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "f":
		return router.OnFIX43SecurityStatus(SecurityStatus{msg}, sessionID)
	case "v":
		return router.OnFIX43SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "9":
		return router.OnFIX43OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "R":
		return router.OnFIX43QuoteRequest(QuoteRequest{msg}, sessionID)
	case "S":
		return router.OnFIX43Quote(Quote{msg}, sessionID)
	case "z":
		return router.OnFIX43DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	case "L":
		return router.OnFIX43ListExecute(ListExecute{msg}, sessionID)
	case "6":
		return router.OnFIX43IOI(IOI{msg}, sessionID)
	case "B":
		return router.OnFIX43News(News{msg}, sessionID)
	case "G":
		return router.OnFIX43OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "N":
		return router.OnFIX43ListStatus(ListStatus{msg}, sessionID)
	case "Z":
		return router.OnFIX43QuoteCancel(QuoteCancel{msg}, sessionID)
	case "s":
		return router.OnFIX43NewOrderCross(NewOrderCross{msg}, sessionID)
	case "w":
		return router.OnFIX43SecurityTypes(SecurityTypes{msg}, sessionID)
	case "7":
		return router.OnFIX43Advertisement(Advertisement{msg}, sessionID)
	case "J":
		return router.OnFIX43Allocation(Allocation{msg}, sessionID)
	case "c":
		return router.OnFIX43SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "p":
		return router.OnFIX43RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "r":
		return router.OnFIX43OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "AB":
		return router.OnFIX43NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "AG":
		return router.OnFIX43QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "F":
		return router.OnFIX43OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "M":
		return router.OnFIX43ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "i":
		return router.OnFIX43MassQuote(MassQuote{msg}, sessionID)
	case "l":
		return router.OnFIX43BidResponse(BidResponse{msg}, sessionID)
	case "t":
		return router.OnFIX43CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "AD":
		return router.OnFIX43TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "AI":
		return router.OnFIX43QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	case "1":
		return router.OnFIX43TestRequest(TestRequest{msg}, sessionID)
	case "3":
		return router.OnFIX43Reject(Reject{msg}, sessionID)
	case "C":
		return router.OnFIX43Email(Email{msg}, sessionID)
	case "y":
		return router.OnFIX43SecurityList(SecurityList{msg}, sessionID)
	case "AE":
		return router.OnFIX43TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "H":
		return router.OnFIX43OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "g":
		return router.OnFIX43TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "h":
		return router.OnFIX43TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "AF":
		return router.OnFIX43OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "4":
		return router.OnFIX43SequenceReset(SequenceReset{msg}, sessionID)
	case "j":
		return router.OnFIX43BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "q":
		return router.OnFIX43OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "0":
		return router.OnFIX43Heartbeat(Heartbeat{msg}, sessionID)
	case "A":
		return router.OnFIX43Logon(Logon{msg}, sessionID)
	case "W":
		return router.OnFIX43MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "a":
		return router.OnFIX43QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "m":
		return router.OnFIX43ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "AA":
		return router.OnFIX43DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "T":
		return router.OnFIX43SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "b":
		return router.OnFIX43MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "k":
		return router.OnFIX43BidRequest(BidRequest{msg}, sessionID)
	case "u":
		return router.OnFIX43CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "5":
		return router.OnFIX43Logout(Logout{msg}, sessionID)
	case "8":
		return router.OnFIX43ExecutionReport(ExecutionReport{msg}, sessionID)
	case "Y":
		return router.OnFIX43MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "Q":
		return router.OnFIX43DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "AC":
		return router.OnFIX43MultilegOrderCancelReplaceRequest(MultilegOrderCancelReplaceRequest{msg}, sessionID)
	case "E":
		return router.OnFIX43NewOrderList(NewOrderList{msg}, sessionID)
	case "K":
		return router.OnFIX43ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "o":
		return router.OnFIX43RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "x":
		return router.OnFIX43SecurityListRequest(SecurityListRequest{msg}, sessionID)
	case "e":
		return router.OnFIX43SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "2":
		return router.OnFIX43ResendRequest(ResendRequest{msg}, sessionID)
	case "V":
		return router.OnFIX43MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "X":
		return router.OnFIX43MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "AH":
		return router.OnFIX43RFQRequest(RFQRequest{msg}, sessionID)
	}
	return quickfixgo.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX43SequenceReset(msg SequenceReset, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43BusinessMessageReject(msg BusinessMessageReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43Heartbeat(msg Heartbeat, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43Logon(msg Logon, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43ListStrikePrice(msg ListStrikePrice, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43SettlementInstructions(msg SettlementInstructions, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43BidRequest(msg BidRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43Logout(msg Logout, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43ExecutionReport(msg ExecutionReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43DontKnowTrade(msg DontKnowTrade, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43MultilegOrderCancelReplaceRequest(msg MultilegOrderCancelReplaceRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43NewOrderList(msg NewOrderList, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43ListCancelRequest(msg ListCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43RegistrationInstructions(msg RegistrationInstructions, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43SecurityListRequest(msg SecurityListRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43ResendRequest(msg ResendRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43MarketDataRequest(msg MarketDataRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43RFQRequest(msg RFQRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43NewOrderSingle(msg NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43AllocationAck(msg AllocationAck, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43SecurityDefinition(msg SecurityDefinition, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43SecurityStatus(msg SecurityStatus, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43OrderCancelReject(msg OrderCancelReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43QuoteRequest(msg QuoteRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43Quote(msg Quote, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43ListExecute(msg ListExecute, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43IOI(msg IOI, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43News(msg News, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43ListStatus(msg ListStatus, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43QuoteCancel(msg QuoteCancel, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43NewOrderCross(msg NewOrderCross, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43SecurityTypes(msg SecurityTypes, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43Advertisement(msg Advertisement, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43Allocation(msg Allocation, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43NewOrderMultileg(msg NewOrderMultileg, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43QuoteRequestReject(msg QuoteRequestReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43OrderCancelRequest(msg OrderCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43ListStatusRequest(msg ListStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43MassQuote(msg MassQuote, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43BidResponse(msg BidResponse, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43QuoteStatusReport(msg QuoteStatusReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43TestRequest(msg TestRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43Reject(msg Reject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43Email(msg Email, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43SecurityList(msg SecurityList, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43TradeCaptureReport(msg TradeCaptureReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43OrderStatusRequest(msg OrderStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43TradingSessionStatus(msg TradingSessionStatus, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX43OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
}
type FIX43MessageCracker struct{}

func (c *FIX43MessageCracker) OnFIX43TestRequest(msg TestRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Reject(msg Reject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Email(msg Email, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityList(msg SecurityList, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TradeCaptureReport(msg TradeCaptureReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderStatusRequest(msg OrderStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TradingSessionStatus(msg TradingSessionStatus, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SequenceReset(msg SequenceReset, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43BusinessMessageReject(msg BusinessMessageReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Heartbeat(msg Heartbeat, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Logon(msg Logon, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteStatusRequest(msg QuoteStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListStrikePrice(msg ListStrikePrice, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43DerivativeSecurityList(msg DerivativeSecurityList, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SettlementInstructions(msg SettlementInstructions, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43BidRequest(msg BidRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Logout(msg Logout, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ExecutionReport(msg ExecutionReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataRequestReject(msg MarketDataRequestReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43DontKnowTrade(msg DontKnowTrade, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MultilegOrderCancelReplaceRequest(msg MultilegOrderCancelReplaceRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderList(msg NewOrderList, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListCancelRequest(msg ListCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43RegistrationInstructions(msg RegistrationInstructions, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityListRequest(msg SecurityListRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityStatusRequest(msg SecurityStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ResendRequest(msg ResendRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataRequest(msg MarketDataRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43RFQRequest(msg RFQRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderSingle(msg NewOrderSingle, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43AllocationAck(msg AllocationAck, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityDefinition(msg SecurityDefinition, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityStatus(msg SecurityStatus, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityTypeRequest(msg SecurityTypeRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderCancelReject(msg OrderCancelReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteRequest(msg QuoteRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Quote(msg Quote, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListExecute(msg ListExecute, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43IOI(msg IOI, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43News(msg News, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListStatus(msg ListStatus, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteCancel(msg QuoteCancel, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderCross(msg NewOrderCross, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityTypes(msg SecurityTypes, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Advertisement(msg Advertisement, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Allocation(msg Allocation, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderMassCancelReport(msg OrderMassCancelReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderMultileg(msg NewOrderMultileg, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteRequestReject(msg QuoteRequestReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderCancelRequest(msg OrderCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListStatusRequest(msg ListStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MassQuote(msg MassQuote, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43BidResponse(msg BidResponse, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteStatusReport(msg QuoteStatusReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
