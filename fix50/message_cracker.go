package fix50

import (
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session"
)

func Crack(msg message.Message, sessionID session.ID, router MessageRouter) reject.MessageReject {
	switch msgType, _ := msg.Header().StringValue(fix.MsgType); msgType {
	case "6":
		return router.OnFIX50IOI(IOI{msg}, sessionID)
	case "7":
		return router.OnFIX50Advertisement(Advertisement{msg}, sessionID)
	case "8":
		return router.OnFIX50ExecutionReport(ExecutionReport{msg}, sessionID)
	case "9":
		return router.OnFIX50OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "B":
		return router.OnFIX50News(News{msg}, sessionID)
	case "C":
		return router.OnFIX50Email(Email{msg}, sessionID)
	case "D":
		return router.OnFIX50NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "E":
		return router.OnFIX50NewOrderList(NewOrderList{msg}, sessionID)
	case "F":
		return router.OnFIX50OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "G":
		return router.OnFIX50OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "H":
		return router.OnFIX50OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "J":
		return router.OnFIX50AllocationInstruction(AllocationInstruction{msg}, sessionID)
	case "K":
		return router.OnFIX50ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "L":
		return router.OnFIX50ListExecute(ListExecute{msg}, sessionID)
	case "M":
		return router.OnFIX50ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "N":
		return router.OnFIX50ListStatus(ListStatus{msg}, sessionID)
	case "P":
		return router.OnFIX50AllocationInstructionAck(AllocationInstructionAck{msg}, sessionID)
	case "Q":
		return router.OnFIX50DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "R":
		return router.OnFIX50QuoteRequest(QuoteRequest{msg}, sessionID)
	case "S":
		return router.OnFIX50Quote(Quote{msg}, sessionID)
	case "T":
		return router.OnFIX50SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "V":
		return router.OnFIX50MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "W":
		return router.OnFIX50MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "X":
		return router.OnFIX50MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "Y":
		return router.OnFIX50MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "Z":
		return router.OnFIX50QuoteCancel(QuoteCancel{msg}, sessionID)
	case "a":
		return router.OnFIX50QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "b":
		return router.OnFIX50MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "c":
		return router.OnFIX50SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "d":
		return router.OnFIX50SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "e":
		return router.OnFIX50SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "f":
		return router.OnFIX50SecurityStatus(SecurityStatus{msg}, sessionID)
	case "g":
		return router.OnFIX50TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "h":
		return router.OnFIX50TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "i":
		return router.OnFIX50MassQuote(MassQuote{msg}, sessionID)
	case "j":
		return router.OnFIX50BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "k":
		return router.OnFIX50BidRequest(BidRequest{msg}, sessionID)
	case "l":
		return router.OnFIX50BidResponse(BidResponse{msg}, sessionID)
	case "m":
		return router.OnFIX50ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "o":
		return router.OnFIX50RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "p":
		return router.OnFIX50RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "q":
		return router.OnFIX50OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "r":
		return router.OnFIX50OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "s":
		return router.OnFIX50NewOrderCross(NewOrderCross{msg}, sessionID)
	case "t":
		return router.OnFIX50CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "u":
		return router.OnFIX50CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "v":
		return router.OnFIX50SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "w":
		return router.OnFIX50SecurityTypes(SecurityTypes{msg}, sessionID)
	case "x":
		return router.OnFIX50SecurityListRequest(SecurityListRequest{msg}, sessionID)
	case "y":
		return router.OnFIX50SecurityList(SecurityList{msg}, sessionID)
	case "z":
		return router.OnFIX50DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	case "AA":
		return router.OnFIX50DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "AB":
		return router.OnFIX50NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "AC":
		return router.OnFIX50MultilegOrderCancelReplace(MultilegOrderCancelReplace{msg}, sessionID)
	case "AD":
		return router.OnFIX50TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "AE":
		return router.OnFIX50TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "AF":
		return router.OnFIX50OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "AG":
		return router.OnFIX50QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "AH":
		return router.OnFIX50RFQRequest(RFQRequest{msg}, sessionID)
	case "AI":
		return router.OnFIX50QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	case "AJ":
		return router.OnFIX50QuoteResponse(QuoteResponse{msg}, sessionID)
	case "AK":
		return router.OnFIX50Confirmation(Confirmation{msg}, sessionID)
	case "AL":
		return router.OnFIX50PositionMaintenanceRequest(PositionMaintenanceRequest{msg}, sessionID)
	case "AM":
		return router.OnFIX50PositionMaintenanceReport(PositionMaintenanceReport{msg}, sessionID)
	case "AN":
		return router.OnFIX50RequestForPositions(RequestForPositions{msg}, sessionID)
	case "AO":
		return router.OnFIX50RequestForPositionsAck(RequestForPositionsAck{msg}, sessionID)
	case "AP":
		return router.OnFIX50PositionReport(PositionReport{msg}, sessionID)
	case "AQ":
		return router.OnFIX50TradeCaptureReportRequestAck(TradeCaptureReportRequestAck{msg}, sessionID)
	case "AR":
		return router.OnFIX50TradeCaptureReportAck(TradeCaptureReportAck{msg}, sessionID)
	case "AS":
		return router.OnFIX50AllocationReport(AllocationReport{msg}, sessionID)
	case "AT":
		return router.OnFIX50AllocationReportAck(AllocationReportAck{msg}, sessionID)
	case "AU":
		return router.OnFIX50ConfirmationAck(ConfirmationAck{msg}, sessionID)
	case "AV":
		return router.OnFIX50SettlementInstructionRequest(SettlementInstructionRequest{msg}, sessionID)
	case "AW":
		return router.OnFIX50AssignmentReport(AssignmentReport{msg}, sessionID)
	case "AX":
		return router.OnFIX50CollateralRequest(CollateralRequest{msg}, sessionID)
	case "AY":
		return router.OnFIX50CollateralAssignment(CollateralAssignment{msg}, sessionID)
	case "AZ":
		return router.OnFIX50CollateralResponse(CollateralResponse{msg}, sessionID)
	case "BA":
		return router.OnFIX50CollateralReport(CollateralReport{msg}, sessionID)
	case "BB":
		return router.OnFIX50CollateralInquiry(CollateralInquiry{msg}, sessionID)
	case "BC":
		return router.OnFIX50NetworkCounterpartySystemStatusRequest(NetworkCounterpartySystemStatusRequest{msg}, sessionID)
	case "BD":
		return router.OnFIX50NetworkCounterpartySystemStatusResponse(NetworkCounterpartySystemStatusResponse{msg}, sessionID)
	case "BE":
		return router.OnFIX50UserRequest(UserRequest{msg}, sessionID)
	case "BF":
		return router.OnFIX50UserResponse(UserResponse{msg}, sessionID)
	case "BG":
		return router.OnFIX50CollateralInquiryAck(CollateralInquiryAck{msg}, sessionID)
	case "BH":
		return router.OnFIX50ConfirmationRequest(ConfirmationRequest{msg}, sessionID)
	case "BO":
		return router.OnFIX50ContraryIntentionReport(ContraryIntentionReport{msg}, sessionID)
	case "BP":
		return router.OnFIX50SecurityDefinitionUpdateReport(SecurityDefinitionUpdateReport{msg}, sessionID)
	case "BK":
		return router.OnFIX50SecurityListUpdateReport(SecurityListUpdateReport{msg}, sessionID)
	case "BL":
		return router.OnFIX50AdjustedPositionReport(AdjustedPositionReport{msg}, sessionID)
	case "BM":
		return router.OnFIX50AllocationInstructionAlert(AllocationInstructionAlert{msg}, sessionID)
	case "BN":
		return router.OnFIX50ExecutionAcknowledgement(ExecutionAcknowledgement{msg}, sessionID)
	case "BJ":
		return router.OnFIX50TradingSessionList(TradingSessionList{msg}, sessionID)
	case "BI":
		return router.OnFIX50TradingSessionListRequest(TradingSessionListRequest{msg}, sessionID)
	}
	return reject.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX50IOI(msg IOI, sessionID session.ID) reject.MessageReject
	OnFIX50Advertisement(msg Advertisement, sessionID session.ID) reject.MessageReject
	OnFIX50ExecutionReport(msg ExecutionReport, sessionID session.ID) reject.MessageReject
	OnFIX50OrderCancelReject(msg OrderCancelReject, sessionID session.ID) reject.MessageReject
	OnFIX50News(msg News, sessionID session.ID) reject.MessageReject
	OnFIX50Email(msg Email, sessionID session.ID) reject.MessageReject
	OnFIX50NewOrderSingle(msg NewOrderSingle, sessionID session.ID) reject.MessageReject
	OnFIX50NewOrderList(msg NewOrderList, sessionID session.ID) reject.MessageReject
	OnFIX50OrderCancelRequest(msg OrderCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX50OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID session.ID) reject.MessageReject
	OnFIX50OrderStatusRequest(msg OrderStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50AllocationInstruction(msg AllocationInstruction, sessionID session.ID) reject.MessageReject
	OnFIX50ListCancelRequest(msg ListCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX50ListExecute(msg ListExecute, sessionID session.ID) reject.MessageReject
	OnFIX50ListStatusRequest(msg ListStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50ListStatus(msg ListStatus, sessionID session.ID) reject.MessageReject
	OnFIX50AllocationInstructionAck(msg AllocationInstructionAck, sessionID session.ID) reject.MessageReject
	OnFIX50DontKnowTrade(msg DontKnowTrade, sessionID session.ID) reject.MessageReject
	OnFIX50QuoteRequest(msg QuoteRequest, sessionID session.ID) reject.MessageReject
	OnFIX50Quote(msg Quote, sessionID session.ID) reject.MessageReject
	OnFIX50SettlementInstructions(msg SettlementInstructions, sessionID session.ID) reject.MessageReject
	OnFIX50MarketDataRequest(msg MarketDataRequest, sessionID session.ID) reject.MessageReject
	OnFIX50MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID session.ID) reject.MessageReject
	OnFIX50MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID session.ID) reject.MessageReject
	OnFIX50MarketDataRequestReject(msg MarketDataRequestReject, sessionID session.ID) reject.MessageReject
	OnFIX50QuoteCancel(msg QuoteCancel, sessionID session.ID) reject.MessageReject
	OnFIX50QuoteStatusRequest(msg QuoteStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID session.ID) reject.MessageReject
	OnFIX50SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SecurityDefinition(msg SecurityDefinition, sessionID session.ID) reject.MessageReject
	OnFIX50SecurityStatusRequest(msg SecurityStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SecurityStatus(msg SecurityStatus, sessionID session.ID) reject.MessageReject
	OnFIX50TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50TradingSessionStatus(msg TradingSessionStatus, sessionID session.ID) reject.MessageReject
	OnFIX50MassQuote(msg MassQuote, sessionID session.ID) reject.MessageReject
	OnFIX50BusinessMessageReject(msg BusinessMessageReject, sessionID session.ID) reject.MessageReject
	OnFIX50BidRequest(msg BidRequest, sessionID session.ID) reject.MessageReject
	OnFIX50BidResponse(msg BidResponse, sessionID session.ID) reject.MessageReject
	OnFIX50ListStrikePrice(msg ListStrikePrice, sessionID session.ID) reject.MessageReject
	OnFIX50RegistrationInstructions(msg RegistrationInstructions, sessionID session.ID) reject.MessageReject
	OnFIX50RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID session.ID) reject.MessageReject
	OnFIX50OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX50OrderMassCancelReport(msg OrderMassCancelReport, sessionID session.ID) reject.MessageReject
	OnFIX50NewOrderCross(msg NewOrderCross, sessionID session.ID) reject.MessageReject
	OnFIX50CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID session.ID) reject.MessageReject
	OnFIX50CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SecurityTypeRequest(msg SecurityTypeRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SecurityTypes(msg SecurityTypes, sessionID session.ID) reject.MessageReject
	OnFIX50SecurityListRequest(msg SecurityListRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SecurityList(msg SecurityList, sessionID session.ID) reject.MessageReject
	OnFIX50DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID session.ID) reject.MessageReject
	OnFIX50DerivativeSecurityList(msg DerivativeSecurityList, sessionID session.ID) reject.MessageReject
	OnFIX50NewOrderMultileg(msg NewOrderMultileg, sessionID session.ID) reject.MessageReject
	OnFIX50MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID session.ID) reject.MessageReject
	OnFIX50TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID session.ID) reject.MessageReject
	OnFIX50TradeCaptureReport(msg TradeCaptureReport, sessionID session.ID) reject.MessageReject
	OnFIX50OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50QuoteRequestReject(msg QuoteRequestReject, sessionID session.ID) reject.MessageReject
	OnFIX50RFQRequest(msg RFQRequest, sessionID session.ID) reject.MessageReject
	OnFIX50QuoteStatusReport(msg QuoteStatusReport, sessionID session.ID) reject.MessageReject
	OnFIX50QuoteResponse(msg QuoteResponse, sessionID session.ID) reject.MessageReject
	OnFIX50Confirmation(msg Confirmation, sessionID session.ID) reject.MessageReject
	OnFIX50PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID session.ID) reject.MessageReject
	OnFIX50PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID session.ID) reject.MessageReject
	OnFIX50RequestForPositions(msg RequestForPositions, sessionID session.ID) reject.MessageReject
	OnFIX50RequestForPositionsAck(msg RequestForPositionsAck, sessionID session.ID) reject.MessageReject
	OnFIX50PositionReport(msg PositionReport, sessionID session.ID) reject.MessageReject
	OnFIX50TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID session.ID) reject.MessageReject
	OnFIX50TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID session.ID) reject.MessageReject
	OnFIX50AllocationReport(msg AllocationReport, sessionID session.ID) reject.MessageReject
	OnFIX50AllocationReportAck(msg AllocationReportAck, sessionID session.ID) reject.MessageReject
	OnFIX50ConfirmationAck(msg ConfirmationAck, sessionID session.ID) reject.MessageReject
	OnFIX50SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID session.ID) reject.MessageReject
	OnFIX50AssignmentReport(msg AssignmentReport, sessionID session.ID) reject.MessageReject
	OnFIX50CollateralRequest(msg CollateralRequest, sessionID session.ID) reject.MessageReject
	OnFIX50CollateralAssignment(msg CollateralAssignment, sessionID session.ID) reject.MessageReject
	OnFIX50CollateralResponse(msg CollateralResponse, sessionID session.ID) reject.MessageReject
	OnFIX50CollateralReport(msg CollateralReport, sessionID session.ID) reject.MessageReject
	OnFIX50CollateralInquiry(msg CollateralInquiry, sessionID session.ID) reject.MessageReject
	OnFIX50NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID session.ID) reject.MessageReject
	OnFIX50UserRequest(msg UserRequest, sessionID session.ID) reject.MessageReject
	OnFIX50UserResponse(msg UserResponse, sessionID session.ID) reject.MessageReject
	OnFIX50CollateralInquiryAck(msg CollateralInquiryAck, sessionID session.ID) reject.MessageReject
	OnFIX50ConfirmationRequest(msg ConfirmationRequest, sessionID session.ID) reject.MessageReject
	OnFIX50ContraryIntentionReport(msg ContraryIntentionReport, sessionID session.ID) reject.MessageReject
	OnFIX50SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionID session.ID) reject.MessageReject
	OnFIX50SecurityListUpdateReport(msg SecurityListUpdateReport, sessionID session.ID) reject.MessageReject
	OnFIX50AdjustedPositionReport(msg AdjustedPositionReport, sessionID session.ID) reject.MessageReject
	OnFIX50AllocationInstructionAlert(msg AllocationInstructionAlert, sessionID session.ID) reject.MessageReject
	OnFIX50ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionID session.ID) reject.MessageReject
	OnFIX50TradingSessionList(msg TradingSessionList, sessionID session.ID) reject.MessageReject
	OnFIX50TradingSessionListRequest(msg TradingSessionListRequest, sessionID session.ID) reject.MessageReject
}
type FIX50MessageCracker struct{}

func (c *FIX50MessageCracker) OnFIX50IOI(msg IOI, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50Advertisement(msg Advertisement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50ExecutionReport(msg ExecutionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50OrderCancelReject(msg OrderCancelReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50News(msg News, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50Email(msg Email, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50NewOrderSingle(msg NewOrderSingle, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50NewOrderList(msg NewOrderList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50OrderCancelRequest(msg OrderCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50OrderStatusRequest(msg OrderStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50AllocationInstruction(msg AllocationInstruction, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50ListCancelRequest(msg ListCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50ListExecute(msg ListExecute, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50ListStatusRequest(msg ListStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50ListStatus(msg ListStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50AllocationInstructionAck(msg AllocationInstructionAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50DontKnowTrade(msg DontKnowTrade, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50QuoteRequest(msg QuoteRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50Quote(msg Quote, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50SettlementInstructions(msg SettlementInstructions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50MarketDataRequest(msg MarketDataRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50MarketDataRequestReject(msg MarketDataRequestReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50QuoteCancel(msg QuoteCancel, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50QuoteStatusRequest(msg QuoteStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50SecurityDefinition(msg SecurityDefinition, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50SecurityStatusRequest(msg SecurityStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50SecurityStatus(msg SecurityStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50TradingSessionStatus(msg TradingSessionStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50MassQuote(msg MassQuote, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50BusinessMessageReject(msg BusinessMessageReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50BidRequest(msg BidRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50BidResponse(msg BidResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50ListStrikePrice(msg ListStrikePrice, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50RegistrationInstructions(msg RegistrationInstructions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50OrderMassCancelReport(msg OrderMassCancelReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50NewOrderCross(msg NewOrderCross, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50SecurityTypeRequest(msg SecurityTypeRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50SecurityTypes(msg SecurityTypes, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50SecurityListRequest(msg SecurityListRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50SecurityList(msg SecurityList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50DerivativeSecurityList(msg DerivativeSecurityList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50NewOrderMultileg(msg NewOrderMultileg, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50TradeCaptureReport(msg TradeCaptureReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50QuoteRequestReject(msg QuoteRequestReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50RFQRequest(msg RFQRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50QuoteStatusReport(msg QuoteStatusReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50QuoteResponse(msg QuoteResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50Confirmation(msg Confirmation, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50PositionMaintenanceReport(msg PositionMaintenanceReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50RequestForPositions(msg RequestForPositions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50RequestForPositionsAck(msg RequestForPositionsAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50PositionReport(msg PositionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50TradeCaptureReportAck(msg TradeCaptureReportAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50AllocationReport(msg AllocationReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50AllocationReportAck(msg AllocationReportAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50ConfirmationAck(msg ConfirmationAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50SettlementInstructionRequest(msg SettlementInstructionRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50AssignmentReport(msg AssignmentReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50CollateralRequest(msg CollateralRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50CollateralAssignment(msg CollateralAssignment, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50CollateralResponse(msg CollateralResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50CollateralReport(msg CollateralReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50CollateralInquiry(msg CollateralInquiry, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50UserRequest(msg UserRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50UserResponse(msg UserResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50CollateralInquiryAck(msg CollateralInquiryAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50ConfirmationRequest(msg ConfirmationRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50ContraryIntentionReport(msg ContraryIntentionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50SecurityListUpdateReport(msg SecurityListUpdateReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50AdjustedPositionReport(msg AdjustedPositionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50AllocationInstructionAlert(msg AllocationInstructionAlert, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50TradingSessionList(msg TradingSessionList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50MessageCracker) OnFIX50TradingSessionListRequest(msg TradingSessionListRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
