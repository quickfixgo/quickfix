package fix43

import (
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session"
)

func Crack(msg message.Message, sessionID session.ID, router MessageRouter) reject.MessageReject {
	switch msgType, _ := msg.Header().StringValue(fix.MsgType); msgType {
	case "0":
		return router.OnFIX43Heartbeat(Heartbeat{msg}, sessionID)
	case "1":
		return router.OnFIX43TestRequest(TestRequest{msg}, sessionID)
	case "2":
		return router.OnFIX43ResendRequest(ResendRequest{msg}, sessionID)
	case "3":
		return router.OnFIX43Reject(Reject{msg}, sessionID)
	case "4":
		return router.OnFIX43SequenceReset(SequenceReset{msg}, sessionID)
	case "5":
		return router.OnFIX43Logout(Logout{msg}, sessionID)
	case "6":
		return router.OnFIX43IOI(IOI{msg}, sessionID)
	case "7":
		return router.OnFIX43Advertisement(Advertisement{msg}, sessionID)
	case "8":
		return router.OnFIX43ExecutionReport(ExecutionReport{msg}, sessionID)
	case "9":
		return router.OnFIX43OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "A":
		return router.OnFIX43Logon(Logon{msg}, sessionID)
	case "B":
		return router.OnFIX43News(News{msg}, sessionID)
	case "C":
		return router.OnFIX43Email(Email{msg}, sessionID)
	case "D":
		return router.OnFIX43NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "E":
		return router.OnFIX43NewOrderList(NewOrderList{msg}, sessionID)
	case "F":
		return router.OnFIX43OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "G":
		return router.OnFIX43OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "H":
		return router.OnFIX43OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "J":
		return router.OnFIX43Allocation(Allocation{msg}, sessionID)
	case "K":
		return router.OnFIX43ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "L":
		return router.OnFIX43ListExecute(ListExecute{msg}, sessionID)
	case "M":
		return router.OnFIX43ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "N":
		return router.OnFIX43ListStatus(ListStatus{msg}, sessionID)
	case "P":
		return router.OnFIX43AllocationAck(AllocationAck{msg}, sessionID)
	case "Q":
		return router.OnFIX43DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "R":
		return router.OnFIX43QuoteRequest(QuoteRequest{msg}, sessionID)
	case "S":
		return router.OnFIX43Quote(Quote{msg}, sessionID)
	case "T":
		return router.OnFIX43SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "V":
		return router.OnFIX43MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "W":
		return router.OnFIX43MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "X":
		return router.OnFIX43MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "Y":
		return router.OnFIX43MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "Z":
		return router.OnFIX43QuoteCancel(QuoteCancel{msg}, sessionID)
	case "a":
		return router.OnFIX43QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "b":
		return router.OnFIX43MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "c":
		return router.OnFIX43SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "d":
		return router.OnFIX43SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "e":
		return router.OnFIX43SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "f":
		return router.OnFIX43SecurityStatus(SecurityStatus{msg}, sessionID)
	case "g":
		return router.OnFIX43TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "h":
		return router.OnFIX43TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "i":
		return router.OnFIX43MassQuote(MassQuote{msg}, sessionID)
	case "j":
		return router.OnFIX43BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "k":
		return router.OnFIX43BidRequest(BidRequest{msg}, sessionID)
	case "l":
		return router.OnFIX43BidResponse(BidResponse{msg}, sessionID)
	case "m":
		return router.OnFIX43ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "o":
		return router.OnFIX43RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "p":
		return router.OnFIX43RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "q":
		return router.OnFIX43OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "r":
		return router.OnFIX43OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "s":
		return router.OnFIX43NewOrderCross(NewOrderCross{msg}, sessionID)
	case "u":
		return router.OnFIX43CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "t":
		return router.OnFIX43CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "v":
		return router.OnFIX43SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "w":
		return router.OnFIX43SecurityTypes(SecurityTypes{msg}, sessionID)
	case "x":
		return router.OnFIX43SecurityListRequest(SecurityListRequest{msg}, sessionID)
	case "y":
		return router.OnFIX43SecurityList(SecurityList{msg}, sessionID)
	case "z":
		return router.OnFIX43DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	case "AA":
		return router.OnFIX43DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "AB":
		return router.OnFIX43NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "AC":
		return router.OnFIX43MultilegOrderCancelReplaceRequest(MultilegOrderCancelReplaceRequest{msg}, sessionID)
	case "AD":
		return router.OnFIX43TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "AE":
		return router.OnFIX43TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "AF":
		return router.OnFIX43OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "AG":
		return router.OnFIX43QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "AH":
		return router.OnFIX43RFQRequest(RFQRequest{msg}, sessionID)
	case "AI":
		return router.OnFIX43QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	}
	return reject.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX43Heartbeat(msg Heartbeat, sessionID session.ID) reject.MessageReject
	OnFIX43TestRequest(msg TestRequest, sessionID session.ID) reject.MessageReject
	OnFIX43ResendRequest(msg ResendRequest, sessionID session.ID) reject.MessageReject
	OnFIX43Reject(msg Reject, sessionID session.ID) reject.MessageReject
	OnFIX43SequenceReset(msg SequenceReset, sessionID session.ID) reject.MessageReject
	OnFIX43Logout(msg Logout, sessionID session.ID) reject.MessageReject
	OnFIX43IOI(msg IOI, sessionID session.ID) reject.MessageReject
	OnFIX43Advertisement(msg Advertisement, sessionID session.ID) reject.MessageReject
	OnFIX43ExecutionReport(msg ExecutionReport, sessionID session.ID) reject.MessageReject
	OnFIX43OrderCancelReject(msg OrderCancelReject, sessionID session.ID) reject.MessageReject
	OnFIX43Logon(msg Logon, sessionID session.ID) reject.MessageReject
	OnFIX43News(msg News, sessionID session.ID) reject.MessageReject
	OnFIX43Email(msg Email, sessionID session.ID) reject.MessageReject
	OnFIX43NewOrderSingle(msg NewOrderSingle, sessionID session.ID) reject.MessageReject
	OnFIX43NewOrderList(msg NewOrderList, sessionID session.ID) reject.MessageReject
	OnFIX43OrderCancelRequest(msg OrderCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX43OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID session.ID) reject.MessageReject
	OnFIX43OrderStatusRequest(msg OrderStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX43Allocation(msg Allocation, sessionID session.ID) reject.MessageReject
	OnFIX43ListCancelRequest(msg ListCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX43ListExecute(msg ListExecute, sessionID session.ID) reject.MessageReject
	OnFIX43ListStatusRequest(msg ListStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX43ListStatus(msg ListStatus, sessionID session.ID) reject.MessageReject
	OnFIX43AllocationAck(msg AllocationAck, sessionID session.ID) reject.MessageReject
	OnFIX43DontKnowTrade(msg DontKnowTrade, sessionID session.ID) reject.MessageReject
	OnFIX43QuoteRequest(msg QuoteRequest, sessionID session.ID) reject.MessageReject
	OnFIX43Quote(msg Quote, sessionID session.ID) reject.MessageReject
	OnFIX43SettlementInstructions(msg SettlementInstructions, sessionID session.ID) reject.MessageReject
	OnFIX43MarketDataRequest(msg MarketDataRequest, sessionID session.ID) reject.MessageReject
	OnFIX43MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID session.ID) reject.MessageReject
	OnFIX43MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID session.ID) reject.MessageReject
	OnFIX43MarketDataRequestReject(msg MarketDataRequestReject, sessionID session.ID) reject.MessageReject
	OnFIX43QuoteCancel(msg QuoteCancel, sessionID session.ID) reject.MessageReject
	OnFIX43QuoteStatusRequest(msg QuoteStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX43MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID session.ID) reject.MessageReject
	OnFIX43SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID session.ID) reject.MessageReject
	OnFIX43SecurityDefinition(msg SecurityDefinition, sessionID session.ID) reject.MessageReject
	OnFIX43SecurityStatusRequest(msg SecurityStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX43SecurityStatus(msg SecurityStatus, sessionID session.ID) reject.MessageReject
	OnFIX43TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX43TradingSessionStatus(msg TradingSessionStatus, sessionID session.ID) reject.MessageReject
	OnFIX43MassQuote(msg MassQuote, sessionID session.ID) reject.MessageReject
	OnFIX43BusinessMessageReject(msg BusinessMessageReject, sessionID session.ID) reject.MessageReject
	OnFIX43BidRequest(msg BidRequest, sessionID session.ID) reject.MessageReject
	OnFIX43BidResponse(msg BidResponse, sessionID session.ID) reject.MessageReject
	OnFIX43ListStrikePrice(msg ListStrikePrice, sessionID session.ID) reject.MessageReject
	OnFIX43RegistrationInstructions(msg RegistrationInstructions, sessionID session.ID) reject.MessageReject
	OnFIX43RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID session.ID) reject.MessageReject
	OnFIX43OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX43OrderMassCancelReport(msg OrderMassCancelReport, sessionID session.ID) reject.MessageReject
	OnFIX43NewOrderCross(msg NewOrderCross, sessionID session.ID) reject.MessageReject
	OnFIX43CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX43CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID session.ID) reject.MessageReject
	OnFIX43SecurityTypeRequest(msg SecurityTypeRequest, sessionID session.ID) reject.MessageReject
	OnFIX43SecurityTypes(msg SecurityTypes, sessionID session.ID) reject.MessageReject
	OnFIX43SecurityListRequest(msg SecurityListRequest, sessionID session.ID) reject.MessageReject
	OnFIX43SecurityList(msg SecurityList, sessionID session.ID) reject.MessageReject
	OnFIX43DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID session.ID) reject.MessageReject
	OnFIX43DerivativeSecurityList(msg DerivativeSecurityList, sessionID session.ID) reject.MessageReject
	OnFIX43NewOrderMultileg(msg NewOrderMultileg, sessionID session.ID) reject.MessageReject
	OnFIX43MultilegOrderCancelReplaceRequest(msg MultilegOrderCancelReplaceRequest, sessionID session.ID) reject.MessageReject
	OnFIX43TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID session.ID) reject.MessageReject
	OnFIX43TradeCaptureReport(msg TradeCaptureReport, sessionID session.ID) reject.MessageReject
	OnFIX43OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX43QuoteRequestReject(msg QuoteRequestReject, sessionID session.ID) reject.MessageReject
	OnFIX43RFQRequest(msg RFQRequest, sessionID session.ID) reject.MessageReject
	OnFIX43QuoteStatusReport(msg QuoteStatusReport, sessionID session.ID) reject.MessageReject
}
type FIX43MessageCracker struct{}

func (c *FIX43MessageCracker) OnFIX43Heartbeat(msg Heartbeat, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43TestRequest(msg TestRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43ResendRequest(msg ResendRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43Reject(msg Reject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43SequenceReset(msg SequenceReset, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43Logout(msg Logout, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43IOI(msg IOI, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43Advertisement(msg Advertisement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43ExecutionReport(msg ExecutionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43OrderCancelReject(msg OrderCancelReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43Logon(msg Logon, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43News(msg News, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43Email(msg Email, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderSingle(msg NewOrderSingle, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderList(msg NewOrderList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43OrderCancelRequest(msg OrderCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43OrderStatusRequest(msg OrderStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43Allocation(msg Allocation, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43ListCancelRequest(msg ListCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43ListExecute(msg ListExecute, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43ListStatusRequest(msg ListStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43ListStatus(msg ListStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43AllocationAck(msg AllocationAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43DontKnowTrade(msg DontKnowTrade, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43QuoteRequest(msg QuoteRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43Quote(msg Quote, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43SettlementInstructions(msg SettlementInstructions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataRequest(msg MarketDataRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataRequestReject(msg MarketDataRequestReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43QuoteCancel(msg QuoteCancel, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43QuoteStatusRequest(msg QuoteStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43SecurityDefinition(msg SecurityDefinition, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43SecurityStatusRequest(msg SecurityStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43SecurityStatus(msg SecurityStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43TradingSessionStatus(msg TradingSessionStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43MassQuote(msg MassQuote, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43BusinessMessageReject(msg BusinessMessageReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43BidRequest(msg BidRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43BidResponse(msg BidResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43ListStrikePrice(msg ListStrikePrice, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43RegistrationInstructions(msg RegistrationInstructions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43OrderMassCancelReport(msg OrderMassCancelReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderCross(msg NewOrderCross, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43SecurityTypeRequest(msg SecurityTypeRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43SecurityTypes(msg SecurityTypes, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43SecurityListRequest(msg SecurityListRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43SecurityList(msg SecurityList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43DerivativeSecurityList(msg DerivativeSecurityList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderMultileg(msg NewOrderMultileg, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43MultilegOrderCancelReplaceRequest(msg MultilegOrderCancelReplaceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43TradeCaptureReport(msg TradeCaptureReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43QuoteRequestReject(msg QuoteRequestReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43RFQRequest(msg RFQRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX43MessageCracker) OnFIX43QuoteStatusReport(msg QuoteStatusReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
