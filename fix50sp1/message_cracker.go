package fix50sp1

import (
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session"
)

func Crack(msg message.Message, sessionID session.ID, router MessageRouter) reject.MessageReject {
	switch msgType, _ := msg.Header().StringValue(fix.MsgType); msgType {
	case "6":
		return router.OnFIX50SP1IOI(IOI{msg}, sessionID)
	case "7":
		return router.OnFIX50SP1Advertisement(Advertisement{msg}, sessionID)
	case "8":
		return router.OnFIX50SP1ExecutionReport(ExecutionReport{msg}, sessionID)
	case "9":
		return router.OnFIX50SP1OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "B":
		return router.OnFIX50SP1News(News{msg}, sessionID)
	case "C":
		return router.OnFIX50SP1Email(Email{msg}, sessionID)
	case "D":
		return router.OnFIX50SP1NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "E":
		return router.OnFIX50SP1NewOrderList(NewOrderList{msg}, sessionID)
	case "F":
		return router.OnFIX50SP1OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "G":
		return router.OnFIX50SP1OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "H":
		return router.OnFIX50SP1OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "J":
		return router.OnFIX50SP1AllocationInstruction(AllocationInstruction{msg}, sessionID)
	case "K":
		return router.OnFIX50SP1ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "L":
		return router.OnFIX50SP1ListExecute(ListExecute{msg}, sessionID)
	case "M":
		return router.OnFIX50SP1ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "N":
		return router.OnFIX50SP1ListStatus(ListStatus{msg}, sessionID)
	case "P":
		return router.OnFIX50SP1AllocationInstructionAck(AllocationInstructionAck{msg}, sessionID)
	case "Q":
		return router.OnFIX50SP1DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "R":
		return router.OnFIX50SP1QuoteRequest(QuoteRequest{msg}, sessionID)
	case "S":
		return router.OnFIX50SP1Quote(Quote{msg}, sessionID)
	case "T":
		return router.OnFIX50SP1SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "V":
		return router.OnFIX50SP1MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "W":
		return router.OnFIX50SP1MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "X":
		return router.OnFIX50SP1MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "Y":
		return router.OnFIX50SP1MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "Z":
		return router.OnFIX50SP1QuoteCancel(QuoteCancel{msg}, sessionID)
	case "a":
		return router.OnFIX50SP1QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "b":
		return router.OnFIX50SP1MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "c":
		return router.OnFIX50SP1SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "d":
		return router.OnFIX50SP1SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "e":
		return router.OnFIX50SP1SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "f":
		return router.OnFIX50SP1SecurityStatus(SecurityStatus{msg}, sessionID)
	case "g":
		return router.OnFIX50SP1TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "h":
		return router.OnFIX50SP1TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "i":
		return router.OnFIX50SP1MassQuote(MassQuote{msg}, sessionID)
	case "j":
		return router.OnFIX50SP1BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "k":
		return router.OnFIX50SP1BidRequest(BidRequest{msg}, sessionID)
	case "l":
		return router.OnFIX50SP1BidResponse(BidResponse{msg}, sessionID)
	case "m":
		return router.OnFIX50SP1ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "o":
		return router.OnFIX50SP1RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "p":
		return router.OnFIX50SP1RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "q":
		return router.OnFIX50SP1OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "r":
		return router.OnFIX50SP1OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "s":
		return router.OnFIX50SP1NewOrderCross(NewOrderCross{msg}, sessionID)
	case "t":
		return router.OnFIX50SP1CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "u":
		return router.OnFIX50SP1CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "v":
		return router.OnFIX50SP1SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "w":
		return router.OnFIX50SP1SecurityTypes(SecurityTypes{msg}, sessionID)
	case "x":
		return router.OnFIX50SP1SecurityListRequest(SecurityListRequest{msg}, sessionID)
	case "y":
		return router.OnFIX50SP1SecurityList(SecurityList{msg}, sessionID)
	case "z":
		return router.OnFIX50SP1DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	case "AA":
		return router.OnFIX50SP1DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "AB":
		return router.OnFIX50SP1NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "AC":
		return router.OnFIX50SP1MultilegOrderCancelReplace(MultilegOrderCancelReplace{msg}, sessionID)
	case "AD":
		return router.OnFIX50SP1TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "AE":
		return router.OnFIX50SP1TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "AF":
		return router.OnFIX50SP1OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "AG":
		return router.OnFIX50SP1QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "AH":
		return router.OnFIX50SP1RFQRequest(RFQRequest{msg}, sessionID)
	case "AI":
		return router.OnFIX50SP1QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	case "AJ":
		return router.OnFIX50SP1QuoteResponse(QuoteResponse{msg}, sessionID)
	case "AK":
		return router.OnFIX50SP1Confirmation(Confirmation{msg}, sessionID)
	case "AL":
		return router.OnFIX50SP1PositionMaintenanceRequest(PositionMaintenanceRequest{msg}, sessionID)
	case "AM":
		return router.OnFIX50SP1PositionMaintenanceReport(PositionMaintenanceReport{msg}, sessionID)
	case "AN":
		return router.OnFIX50SP1RequestForPositions(RequestForPositions{msg}, sessionID)
	case "AO":
		return router.OnFIX50SP1RequestForPositionsAck(RequestForPositionsAck{msg}, sessionID)
	case "AP":
		return router.OnFIX50SP1PositionReport(PositionReport{msg}, sessionID)
	case "AQ":
		return router.OnFIX50SP1TradeCaptureReportRequestAck(TradeCaptureReportRequestAck{msg}, sessionID)
	case "AR":
		return router.OnFIX50SP1TradeCaptureReportAck(TradeCaptureReportAck{msg}, sessionID)
	case "AS":
		return router.OnFIX50SP1AllocationReport(AllocationReport{msg}, sessionID)
	case "AT":
		return router.OnFIX50SP1AllocationReportAck(AllocationReportAck{msg}, sessionID)
	case "AU":
		return router.OnFIX50SP1ConfirmationAck(ConfirmationAck{msg}, sessionID)
	case "AV":
		return router.OnFIX50SP1SettlementInstructionRequest(SettlementInstructionRequest{msg}, sessionID)
	case "AW":
		return router.OnFIX50SP1AssignmentReport(AssignmentReport{msg}, sessionID)
	case "AX":
		return router.OnFIX50SP1CollateralRequest(CollateralRequest{msg}, sessionID)
	case "AY":
		return router.OnFIX50SP1CollateralAssignment(CollateralAssignment{msg}, sessionID)
	case "AZ":
		return router.OnFIX50SP1CollateralResponse(CollateralResponse{msg}, sessionID)
	case "BA":
		return router.OnFIX50SP1CollateralReport(CollateralReport{msg}, sessionID)
	case "BB":
		return router.OnFIX50SP1CollateralInquiry(CollateralInquiry{msg}, sessionID)
	case "BC":
		return router.OnFIX50SP1NetworkCounterpartySystemStatusRequest(NetworkCounterpartySystemStatusRequest{msg}, sessionID)
	case "BD":
		return router.OnFIX50SP1NetworkCounterpartySystemStatusResponse(NetworkCounterpartySystemStatusResponse{msg}, sessionID)
	case "BE":
		return router.OnFIX50SP1UserRequest(UserRequest{msg}, sessionID)
	case "BF":
		return router.OnFIX50SP1UserResponse(UserResponse{msg}, sessionID)
	case "BG":
		return router.OnFIX50SP1CollateralInquiryAck(CollateralInquiryAck{msg}, sessionID)
	case "BH":
		return router.OnFIX50SP1ConfirmationRequest(ConfirmationRequest{msg}, sessionID)
	case "BO":
		return router.OnFIX50SP1ContraryIntentionReport(ContraryIntentionReport{msg}, sessionID)
	case "BP":
		return router.OnFIX50SP1SecurityDefinitionUpdateReport(SecurityDefinitionUpdateReport{msg}, sessionID)
	case "BK":
		return router.OnFIX50SP1SecurityListUpdateReport(SecurityListUpdateReport{msg}, sessionID)
	case "BL":
		return router.OnFIX50SP1AdjustedPositionReport(AdjustedPositionReport{msg}, sessionID)
	case "BM":
		return router.OnFIX50SP1AllocationInstructionAlert(AllocationInstructionAlert{msg}, sessionID)
	case "BN":
		return router.OnFIX50SP1ExecutionAcknowledgement(ExecutionAcknowledgement{msg}, sessionID)
	case "BJ":
		return router.OnFIX50SP1TradingSessionList(TradingSessionList{msg}, sessionID)
	case "BI":
		return router.OnFIX50SP1TradingSessionListRequest(TradingSessionListRequest{msg}, sessionID)
	case "BQ":
		return router.OnFIX50SP1SettlementObligationReport(SettlementObligationReport{msg}, sessionID)
	case "BR":
		return router.OnFIX50SP1DerivativeSecurityListUpdateReport(DerivativeSecurityListUpdateReport{msg}, sessionID)
	case "BS":
		return router.OnFIX50SP1TradingSessionListUpdateReport(TradingSessionListUpdateReport{msg}, sessionID)
	case "BT":
		return router.OnFIX50SP1MarketDefinitionRequest(MarketDefinitionRequest{msg}, sessionID)
	case "BU":
		return router.OnFIX50SP1MarketDefinition(MarketDefinition{msg}, sessionID)
	case "BV":
		return router.OnFIX50SP1MarketDefinitionUpdateReport(MarketDefinitionUpdateReport{msg}, sessionID)
	case "BW":
		return router.OnFIX50SP1ApplicationMessageRequest(ApplicationMessageRequest{msg}, sessionID)
	case "BX":
		return router.OnFIX50SP1ApplicationMessageRequestAck(ApplicationMessageRequestAck{msg}, sessionID)
	case "BY":
		return router.OnFIX50SP1ApplicationMessageReport(ApplicationMessageReport{msg}, sessionID)
	case "BZ":
		return router.OnFIX50SP1OrderMassActionReport(OrderMassActionReport{msg}, sessionID)
	case "CA":
		return router.OnFIX50SP1OrderMassActionRequest(OrderMassActionRequest{msg}, sessionID)
	case "CB":
		return router.OnFIX50SP1UserNotification(UserNotification{msg}, sessionID)
	}
	return reject.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX50SP1IOI(msg IOI, sessionID session.ID) reject.MessageReject
	OnFIX50SP1Advertisement(msg Advertisement, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ExecutionReport(msg ExecutionReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1OrderCancelReject(msg OrderCancelReject, sessionID session.ID) reject.MessageReject
	OnFIX50SP1News(msg News, sessionID session.ID) reject.MessageReject
	OnFIX50SP1Email(msg Email, sessionID session.ID) reject.MessageReject
	OnFIX50SP1NewOrderSingle(msg NewOrderSingle, sessionID session.ID) reject.MessageReject
	OnFIX50SP1NewOrderList(msg NewOrderList, sessionID session.ID) reject.MessageReject
	OnFIX50SP1OrderCancelRequest(msg OrderCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1OrderStatusRequest(msg OrderStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1AllocationInstruction(msg AllocationInstruction, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ListCancelRequest(msg ListCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ListExecute(msg ListExecute, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ListStatusRequest(msg ListStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ListStatus(msg ListStatus, sessionID session.ID) reject.MessageReject
	OnFIX50SP1AllocationInstructionAck(msg AllocationInstructionAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP1DontKnowTrade(msg DontKnowTrade, sessionID session.ID) reject.MessageReject
	OnFIX50SP1QuoteRequest(msg QuoteRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1Quote(msg Quote, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SettlementInstructions(msg SettlementInstructions, sessionID session.ID) reject.MessageReject
	OnFIX50SP1MarketDataRequest(msg MarketDataRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID session.ID) reject.MessageReject
	OnFIX50SP1MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID session.ID) reject.MessageReject
	OnFIX50SP1MarketDataRequestReject(msg MarketDataRequestReject, sessionID session.ID) reject.MessageReject
	OnFIX50SP1QuoteCancel(msg QuoteCancel, sessionID session.ID) reject.MessageReject
	OnFIX50SP1QuoteStatusRequest(msg QuoteStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SecurityDefinition(msg SecurityDefinition, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SecurityStatusRequest(msg SecurityStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SecurityStatus(msg SecurityStatus, sessionID session.ID) reject.MessageReject
	OnFIX50SP1TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1TradingSessionStatus(msg TradingSessionStatus, sessionID session.ID) reject.MessageReject
	OnFIX50SP1MassQuote(msg MassQuote, sessionID session.ID) reject.MessageReject
	OnFIX50SP1BusinessMessageReject(msg BusinessMessageReject, sessionID session.ID) reject.MessageReject
	OnFIX50SP1BidRequest(msg BidRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1BidResponse(msg BidResponse, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ListStrikePrice(msg ListStrikePrice, sessionID session.ID) reject.MessageReject
	OnFIX50SP1RegistrationInstructions(msg RegistrationInstructions, sessionID session.ID) reject.MessageReject
	OnFIX50SP1RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID session.ID) reject.MessageReject
	OnFIX50SP1OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1OrderMassCancelReport(msg OrderMassCancelReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1NewOrderCross(msg NewOrderCross, sessionID session.ID) reject.MessageReject
	OnFIX50SP1CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SecurityTypeRequest(msg SecurityTypeRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SecurityTypes(msg SecurityTypes, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SecurityListRequest(msg SecurityListRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SecurityList(msg SecurityList, sessionID session.ID) reject.MessageReject
	OnFIX50SP1DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1DerivativeSecurityList(msg DerivativeSecurityList, sessionID session.ID) reject.MessageReject
	OnFIX50SP1NewOrderMultileg(msg NewOrderMultileg, sessionID session.ID) reject.MessageReject
	OnFIX50SP1MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID session.ID) reject.MessageReject
	OnFIX50SP1TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1TradeCaptureReport(msg TradeCaptureReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1QuoteRequestReject(msg QuoteRequestReject, sessionID session.ID) reject.MessageReject
	OnFIX50SP1RFQRequest(msg RFQRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1QuoteStatusReport(msg QuoteStatusReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1QuoteResponse(msg QuoteResponse, sessionID session.ID) reject.MessageReject
	OnFIX50SP1Confirmation(msg Confirmation, sessionID session.ID) reject.MessageReject
	OnFIX50SP1PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1RequestForPositions(msg RequestForPositions, sessionID session.ID) reject.MessageReject
	OnFIX50SP1RequestForPositionsAck(msg RequestForPositionsAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP1PositionReport(msg PositionReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP1TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP1AllocationReport(msg AllocationReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1AllocationReportAck(msg AllocationReportAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ConfirmationAck(msg ConfirmationAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1AssignmentReport(msg AssignmentReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1CollateralRequest(msg CollateralRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1CollateralAssignment(msg CollateralAssignment, sessionID session.ID) reject.MessageReject
	OnFIX50SP1CollateralResponse(msg CollateralResponse, sessionID session.ID) reject.MessageReject
	OnFIX50SP1CollateralReport(msg CollateralReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1CollateralInquiry(msg CollateralInquiry, sessionID session.ID) reject.MessageReject
	OnFIX50SP1NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID session.ID) reject.MessageReject
	OnFIX50SP1UserRequest(msg UserRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1UserResponse(msg UserResponse, sessionID session.ID) reject.MessageReject
	OnFIX50SP1CollateralInquiryAck(msg CollateralInquiryAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ConfirmationRequest(msg ConfirmationRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ContraryIntentionReport(msg ContraryIntentionReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SecurityListUpdateReport(msg SecurityListUpdateReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1AdjustedPositionReport(msg AdjustedPositionReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1AllocationInstructionAlert(msg AllocationInstructionAlert, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionID session.ID) reject.MessageReject
	OnFIX50SP1TradingSessionList(msg TradingSessionList, sessionID session.ID) reject.MessageReject
	OnFIX50SP1TradingSessionListRequest(msg TradingSessionListRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1SettlementObligationReport(msg SettlementObligationReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1MarketDefinitionRequest(msg MarketDefinitionRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1MarketDefinition(msg MarketDefinition, sessionID session.ID) reject.MessageReject
	OnFIX50SP1MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ApplicationMessageRequest(msg ApplicationMessageRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionID session.ID) reject.MessageReject
	OnFIX50SP1ApplicationMessageReport(msg ApplicationMessageReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1OrderMassActionReport(msg OrderMassActionReport, sessionID session.ID) reject.MessageReject
	OnFIX50SP1OrderMassActionRequest(msg OrderMassActionRequest, sessionID session.ID) reject.MessageReject
	OnFIX50SP1UserNotification(msg UserNotification, sessionID session.ID) reject.MessageReject
}
type FIX50SP1MessageCracker struct{}

func (c *FIX50SP1MessageCracker) OnFIX50SP1IOI(msg IOI, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1Advertisement(msg Advertisement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ExecutionReport(msg ExecutionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderCancelReject(msg OrderCancelReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1News(msg News, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1Email(msg Email, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderSingle(msg NewOrderSingle, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderList(msg NewOrderList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderCancelRequest(msg OrderCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderStatusRequest(msg OrderStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationInstruction(msg AllocationInstruction, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListCancelRequest(msg ListCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListExecute(msg ListExecute, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListStatusRequest(msg ListStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListStatus(msg ListStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationInstructionAck(msg AllocationInstructionAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1DontKnowTrade(msg DontKnowTrade, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteRequest(msg QuoteRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1Quote(msg Quote, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SettlementInstructions(msg SettlementInstructions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataRequest(msg MarketDataRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataRequestReject(msg MarketDataRequestReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteCancel(msg QuoteCancel, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteStatusRequest(msg QuoteStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityDefinition(msg SecurityDefinition, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityStatusRequest(msg SecurityStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityStatus(msg SecurityStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionStatus(msg TradingSessionStatus, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MassQuote(msg MassQuote, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1BusinessMessageReject(msg BusinessMessageReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1BidRequest(msg BidRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1BidResponse(msg BidResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListStrikePrice(msg ListStrikePrice, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RegistrationInstructions(msg RegistrationInstructions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassCancelReport(msg OrderMassCancelReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderCross(msg NewOrderCross, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityTypeRequest(msg SecurityTypeRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityTypes(msg SecurityTypes, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityListRequest(msg SecurityListRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityList(msg SecurityList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1DerivativeSecurityList(msg DerivativeSecurityList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderMultileg(msg NewOrderMultileg, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReport(msg TradeCaptureReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteRequestReject(msg QuoteRequestReject, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RFQRequest(msg RFQRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteStatusReport(msg QuoteStatusReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteResponse(msg QuoteResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1Confirmation(msg Confirmation, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1PositionMaintenanceReport(msg PositionMaintenanceReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RequestForPositions(msg RequestForPositions, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RequestForPositionsAck(msg RequestForPositionsAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1PositionReport(msg PositionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReportAck(msg TradeCaptureReportAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationReport(msg AllocationReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationReportAck(msg AllocationReportAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ConfirmationAck(msg ConfirmationAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SettlementInstructionRequest(msg SettlementInstructionRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AssignmentReport(msg AssignmentReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralRequest(msg CollateralRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralAssignment(msg CollateralAssignment, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralResponse(msg CollateralResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralReport(msg CollateralReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralInquiry(msg CollateralInquiry, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1UserRequest(msg UserRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1UserResponse(msg UserResponse, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralInquiryAck(msg CollateralInquiryAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ConfirmationRequest(msg ConfirmationRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ContraryIntentionReport(msg ContraryIntentionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityListUpdateReport(msg SecurityListUpdateReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AdjustedPositionReport(msg AdjustedPositionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationInstructionAlert(msg AllocationInstructionAlert, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionList(msg TradingSessionList, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionListRequest(msg TradingSessionListRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SettlementObligationReport(msg SettlementObligationReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDefinitionRequest(msg MarketDefinitionRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDefinition(msg MarketDefinition, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ApplicationMessageRequest(msg ApplicationMessageRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ApplicationMessageReport(msg ApplicationMessageReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassActionReport(msg OrderMassActionReport, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassActionRequest(msg OrderMassActionRequest, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1UserNotification(msg UserNotification, sessionId session.ID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg)
}
