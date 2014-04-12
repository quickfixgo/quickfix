package fix43

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

func Crack(msg message.Message, sessionID quickfix.SessionID, router MessageRouter) message.MessageReject {

	msgType := new(field.MsgType)
	switch msg.Header.Get(msgType); msgType.Value {
	case "1":
		return router.OnFIX43TestRequest(TestRequest{msg}, sessionID)
	case "D":
		return router.OnFIX43NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "J":
		return router.OnFIX43Allocation(Allocation{msg}, sessionID)
	case "L":
		return router.OnFIX43ListExecute(ListExecute{msg}, sessionID)
	case "N":
		return router.OnFIX43ListStatus(ListStatus{msg}, sessionID)
	case "V":
		return router.OnFIX43MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "C":
		return router.OnFIX43Email(Email{msg}, sessionID)
	case "Q":
		return router.OnFIX43DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "T":
		return router.OnFIX43SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "X":
		return router.OnFIX43MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "c":
		return router.OnFIX43SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "h":
		return router.OnFIX43TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "6":
		return router.OnFIX43IOI(IOI{msg}, sessionID)
	case "y":
		return router.OnFIX43SecurityList(SecurityList{msg}, sessionID)
	case "AI":
		return router.OnFIX43QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	case "5":
		return router.OnFIX43Logout(Logout{msg}, sessionID)
	case "P":
		return router.OnFIX43AllocationAck(AllocationAck{msg}, sessionID)
	case "W":
		return router.OnFIX43MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "i":
		return router.OnFIX43MassQuote(MassQuote{msg}, sessionID)
	case "0":
		return router.OnFIX43Heartbeat(Heartbeat{msg}, sessionID)
	case "H":
		return router.OnFIX43OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "Y":
		return router.OnFIX43MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "m":
		return router.OnFIX43ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "r":
		return router.OnFIX43OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "u":
		return router.OnFIX43CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "Z":
		return router.OnFIX43QuoteCancel(QuoteCancel{msg}, sessionID)
	case "b":
		return router.OnFIX43MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "d":
		return router.OnFIX43SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "g":
		return router.OnFIX43TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "k":
		return router.OnFIX43BidRequest(BidRequest{msg}, sessionID)
	case "l":
		return router.OnFIX43BidResponse(BidResponse{msg}, sessionID)
	case "q":
		return router.OnFIX43OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "v":
		return router.OnFIX43SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "A":
		return router.OnFIX43Logon(Logon{msg}, sessionID)
	case "R":
		return router.OnFIX43QuoteRequest(QuoteRequest{msg}, sessionID)
	case "AA":
		return router.OnFIX43DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "AC":
		return router.OnFIX43MultilegOrderCancelReplaceRequest(MultilegOrderCancelReplaceRequest{msg}, sessionID)
	case "AD":
		return router.OnFIX43TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "B":
		return router.OnFIX43News(News{msg}, sessionID)
	case "2":
		return router.OnFIX43ResendRequest(ResendRequest{msg}, sessionID)
	case "3":
		return router.OnFIX43Reject(Reject{msg}, sessionID)
	case "p":
		return router.OnFIX43RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "w":
		return router.OnFIX43SecurityTypes(SecurityTypes{msg}, sessionID)
	case "K":
		return router.OnFIX43ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "o":
		return router.OnFIX43RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "x":
		return router.OnFIX43SecurityListRequest(SecurityListRequest{msg}, sessionID)
	case "AB":
		return router.OnFIX43NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "4":
		return router.OnFIX43SequenceReset(SequenceReset{msg}, sessionID)
	case "8":
		return router.OnFIX43ExecutionReport(ExecutionReport{msg}, sessionID)
	case "E":
		return router.OnFIX43NewOrderList(NewOrderList{msg}, sessionID)
	case "G":
		return router.OnFIX43OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "a":
		return router.OnFIX43QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "AE":
		return router.OnFIX43TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "AH":
		return router.OnFIX43RFQRequest(RFQRequest{msg}, sessionID)
	case "e":
		return router.OnFIX43SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "t":
		return router.OnFIX43CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "9":
		return router.OnFIX43OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "M":
		return router.OnFIX43ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "S":
		return router.OnFIX43Quote(Quote{msg}, sessionID)
	case "AF":
		return router.OnFIX43OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "AG":
		return router.OnFIX43QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "7":
		return router.OnFIX43Advertisement(Advertisement{msg}, sessionID)
	case "f":
		return router.OnFIX43SecurityStatus(SecurityStatus{msg}, sessionID)
	case "j":
		return router.OnFIX43BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "F":
		return router.OnFIX43OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "s":
		return router.OnFIX43NewOrderCross(NewOrderCross{msg}, sessionID)
	case "z":
		return router.OnFIX43DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	}
	return message.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX43Heartbeat(msg Heartbeat, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43ListStrikePrice(msg ListStrikePrice, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43QuoteCancel(msg QuoteCancel, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43SecurityDefinition(msg SecurityDefinition, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43BidRequest(msg BidRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43BidResponse(msg BidResponse, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43Logon(msg Logon, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43MultilegOrderCancelReplaceRequest(msg MultilegOrderCancelReplaceRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43News(msg News, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43ResendRequest(msg ResendRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43Reject(msg Reject, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43SecurityTypes(msg SecurityTypes, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43RegistrationInstructions(msg RegistrationInstructions, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43SecurityListRequest(msg SecurityListRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43NewOrderMultileg(msg NewOrderMultileg, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43SequenceReset(msg SequenceReset, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43TradeCaptureReport(msg TradeCaptureReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43RFQRequest(msg RFQRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43Quote(msg Quote, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43QuoteRequestReject(msg QuoteRequestReject, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43Advertisement(msg Advertisement, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43SecurityStatus(msg SecurityStatus, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43BusinessMessageReject(msg BusinessMessageReject, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43NewOrderCross(msg NewOrderCross, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43TestRequest(msg TestRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43Allocation(msg Allocation, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43ListExecute(msg ListExecute, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43ListStatus(msg ListStatus, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43MarketDataRequest(msg MarketDataRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43Email(msg Email, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43SettlementInstructions(msg SettlementInstructions, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43TradingSessionStatus(msg TradingSessionStatus, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43IOI(msg IOI, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43SecurityList(msg SecurityList, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43QuoteStatusReport(msg QuoteStatusReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43Logout(msg Logout, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43AllocationAck(msg AllocationAck, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfix.SessionID) message.MessageReject
	OnFIX43MassQuote(msg MassQuote, sessionID quickfix.SessionID) message.MessageReject
}
type FIX43MessageCracker struct{}

func (c *FIX43MessageCracker) OnFIX43Advertisement(msg Advertisement, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityStatus(msg SecurityStatus, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43BusinessMessageReject(msg BusinessMessageReject, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderCancelRequest(msg OrderCancelRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderCross(msg NewOrderCross, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TestRequest(msg TestRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderSingle(msg NewOrderSingle, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Allocation(msg Allocation, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListExecute(msg ListExecute, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListStatus(msg ListStatus, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataRequest(msg MarketDataRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Email(msg Email, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43DontKnowTrade(msg DontKnowTrade, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SettlementInstructions(msg SettlementInstructions, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TradingSessionStatus(msg TradingSessionStatus, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43IOI(msg IOI, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityList(msg SecurityList, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteStatusReport(msg QuoteStatusReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Logout(msg Logout, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43AllocationAck(msg AllocationAck, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MassQuote(msg MassQuote, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Heartbeat(msg Heartbeat, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderStatusRequest(msg OrderStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataRequestReject(msg MarketDataRequestReject, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListStrikePrice(msg ListStrikePrice, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderMassCancelReport(msg OrderMassCancelReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteCancel(msg QuoteCancel, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityDefinition(msg SecurityDefinition, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43BidRequest(msg BidRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43BidResponse(msg BidResponse, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityTypeRequest(msg SecurityTypeRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Logon(msg Logon, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteRequest(msg QuoteRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43DerivativeSecurityList(msg DerivativeSecurityList, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MultilegOrderCancelReplaceRequest(msg MultilegOrderCancelReplaceRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43News(msg News, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ResendRequest(msg ResendRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Reject(msg Reject, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityTypes(msg SecurityTypes, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListCancelRequest(msg ListCancelRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43RegistrationInstructions(msg RegistrationInstructions, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityListRequest(msg SecurityListRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderMultileg(msg NewOrderMultileg, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SequenceReset(msg SequenceReset, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ExecutionReport(msg ExecutionReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderList(msg NewOrderList, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteStatusRequest(msg QuoteStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TradeCaptureReport(msg TradeCaptureReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43RFQRequest(msg RFQRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityStatusRequest(msg SecurityStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderCancelReject(msg OrderCancelReject, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListStatusRequest(msg ListStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Quote(msg Quote, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteRequestReject(msg QuoteRequestReject, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
