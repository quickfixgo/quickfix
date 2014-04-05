package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

func Crack(msg quickfixgo.Message, sessionID quickfixgo.SessionID, router MessageRouter) quickfixgo.MessageReject {

	msgType := new(field.MsgType)
	switch msg.Header.Get(msgType); msgType.Value {
	case "Q":
		return router.OnFIX50SP1DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "R":
		return router.OnFIX50SP1QuoteRequest(QuoteRequest{msg}, sessionID)
	case "BN":
		return router.OnFIX50SP1ExecutionAcknowledgement(ExecutionAcknowledgement{msg}, sessionID)
	case "7":
		return router.OnFIX50SP1Advertisement(Advertisement{msg}, sessionID)
	case "S":
		return router.OnFIX50SP1Quote(Quote{msg}, sessionID)
	case "AK":
		return router.OnFIX50SP1Confirmation(Confirmation{msg}, sessionID)
	case "BR":
		return router.OnFIX50SP1DerivativeSecurityListUpdateReport(DerivativeSecurityListUpdateReport{msg}, sessionID)
	case "AH":
		return router.OnFIX50SP1RFQRequest(RFQRequest{msg}, sessionID)
	case "BG":
		return router.OnFIX50SP1CollateralInquiryAck(CollateralInquiryAck{msg}, sessionID)
	case "C":
		return router.OnFIX50SP1Email(Email{msg}, sessionID)
	case "s":
		return router.OnFIX50SP1NewOrderCross(NewOrderCross{msg}, sessionID)
	case "AY":
		return router.OnFIX50SP1CollateralAssignment(CollateralAssignment{msg}, sessionID)
	case "BA":
		return router.OnFIX50SP1CollateralReport(CollateralReport{msg}, sessionID)
	case "q":
		return router.OnFIX50SP1OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "t":
		return router.OnFIX50SP1CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "B":
		return router.OnFIX50SP1News(News{msg}, sessionID)
	case "j":
		return router.OnFIX50SP1BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "k":
		return router.OnFIX50SP1BidRequest(BidRequest{msg}, sessionID)
	case "BM":
		return router.OnFIX50SP1AllocationInstructionAlert(AllocationInstructionAlert{msg}, sessionID)
	case "BU":
		return router.OnFIX50SP1MarketDefinition(MarketDefinition{msg}, sessionID)
	case "BX":
		return router.OnFIX50SP1ApplicationMessageRequestAck(ApplicationMessageRequestAck{msg}, sessionID)
	case "CA":
		return router.OnFIX50SP1OrderMassActionRequest(OrderMassActionRequest{msg}, sessionID)
	case "AF":
		return router.OnFIX50SP1OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "X":
		return router.OnFIX50SP1MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "v":
		return router.OnFIX50SP1SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "AL":
		return router.OnFIX50SP1PositionMaintenanceRequest(PositionMaintenanceRequest{msg}, sessionID)
	case "h":
		return router.OnFIX50SP1TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "BQ":
		return router.OnFIX50SP1SettlementObligationReport(SettlementObligationReport{msg}, sessionID)
	case "L":
		return router.OnFIX50SP1ListExecute(ListExecute{msg}, sessionID)
	case "AP":
		return router.OnFIX50SP1PositionReport(PositionReport{msg}, sessionID)
	case "BH":
		return router.OnFIX50SP1ConfirmationRequest(ConfirmationRequest{msg}, sessionID)
	case "BS":
		return router.OnFIX50SP1TradingSessionListUpdateReport(TradingSessionListUpdateReport{msg}, sessionID)
	case "Z":
		return router.OnFIX50SP1QuoteCancel(QuoteCancel{msg}, sessionID)
	case "AD":
		return router.OnFIX50SP1TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "AM":
		return router.OnFIX50SP1PositionMaintenanceReport(PositionMaintenanceReport{msg}, sessionID)
	case "f":
		return router.OnFIX50SP1SecurityStatus(SecurityStatus{msg}, sessionID)
	case "AZ":
		return router.OnFIX50SP1CollateralResponse(CollateralResponse{msg}, sessionID)
	case "G":
		return router.OnFIX50SP1OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "V":
		return router.OnFIX50SP1MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "AQ":
		return router.OnFIX50SP1TradeCaptureReportRequestAck(TradeCaptureReportRequestAck{msg}, sessionID)
	case "AX":
		return router.OnFIX50SP1CollateralRequest(CollateralRequest{msg}, sessionID)
	case "BE":
		return router.OnFIX50SP1UserRequest(UserRequest{msg}, sessionID)
	case "BY":
		return router.OnFIX50SP1ApplicationMessageReport(ApplicationMessageReport{msg}, sessionID)
	case "K":
		return router.OnFIX50SP1ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "b":
		return router.OnFIX50SP1MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "p":
		return router.OnFIX50SP1RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "AR":
		return router.OnFIX50SP1TradeCaptureReportAck(TradeCaptureReportAck{msg}, sessionID)
	case "AS":
		return router.OnFIX50SP1AllocationReport(AllocationReport{msg}, sessionID)
	case "AV":
		return router.OnFIX50SP1SettlementInstructionRequest(SettlementInstructionRequest{msg}, sessionID)
	case "E":
		return router.OnFIX50SP1NewOrderList(NewOrderList{msg}, sessionID)
	case "d":
		return router.OnFIX50SP1SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "AT":
		return router.OnFIX50SP1AllocationReportAck(AllocationReportAck{msg}, sessionID)
	case "8":
		return router.OnFIX50SP1ExecutionReport(ExecutionReport{msg}, sessionID)
	case "AC":
		return router.OnFIX50SP1MultilegOrderCancelReplace(MultilegOrderCancelReplace{msg}, sessionID)
	case "BJ":
		return router.OnFIX50SP1TradingSessionList(TradingSessionList{msg}, sessionID)
	case "CB":
		return router.OnFIX50SP1UserNotification(UserNotification{msg}, sessionID)
	case "r":
		return router.OnFIX50SP1OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "AN":
		return router.OnFIX50SP1RequestForPositions(RequestForPositions{msg}, sessionID)
	case "BI":
		return router.OnFIX50SP1TradingSessionListRequest(TradingSessionListRequest{msg}, sessionID)
	case "BV":
		return router.OnFIX50SP1MarketDefinitionUpdateReport(MarketDefinitionUpdateReport{msg}, sessionID)
	case "c":
		return router.OnFIX50SP1SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "y":
		return router.OnFIX50SP1SecurityList(SecurityList{msg}, sessionID)
	case "N":
		return router.OnFIX50SP1ListStatus(ListStatus{msg}, sessionID)
	case "W":
		return router.OnFIX50SP1MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "g":
		return router.OnFIX50SP1TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "o":
		return router.OnFIX50SP1RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "AU":
		return router.OnFIX50SP1ConfirmationAck(ConfirmationAck{msg}, sessionID)
	case "BL":
		return router.OnFIX50SP1AdjustedPositionReport(AdjustedPositionReport{msg}, sessionID)
	case "i":
		return router.OnFIX50SP1MassQuote(MassQuote{msg}, sessionID)
	case "z":
		return router.OnFIX50SP1DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	case "AO":
		return router.OnFIX50SP1RequestForPositionsAck(RequestForPositionsAck{msg}, sessionID)
	case "BK":
		return router.OnFIX50SP1SecurityListUpdateReport(SecurityListUpdateReport{msg}, sessionID)
	case "l":
		return router.OnFIX50SP1BidResponse(BidResponse{msg}, sessionID)
	case "u":
		return router.OnFIX50SP1CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "BB":
		return router.OnFIX50SP1CollateralInquiry(CollateralInquiry{msg}, sessionID)
	case "a":
		return router.OnFIX50SP1QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "D":
		return router.OnFIX50SP1NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "w":
		return router.OnFIX50SP1SecurityTypes(SecurityTypes{msg}, sessionID)
	case "BF":
		return router.OnFIX50SP1UserResponse(UserResponse{msg}, sessionID)
	case "AW":
		return router.OnFIX50SP1AssignmentReport(AssignmentReport{msg}, sessionID)
	case "6":
		return router.OnFIX50SP1IOI(IOI{msg}, sessionID)
	case "Y":
		return router.OnFIX50SP1MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "BC":
		return router.OnFIX50SP1NetworkCounterpartySystemStatusRequest(NetworkCounterpartySystemStatusRequest{msg}, sessionID)
	case "J":
		return router.OnFIX50SP1AllocationInstruction(AllocationInstruction{msg}, sessionID)
	case "e":
		return router.OnFIX50SP1SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "m":
		return router.OnFIX50SP1ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "AA":
		return router.OnFIX50SP1DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "BO":
		return router.OnFIX50SP1ContraryIntentionReport(ContraryIntentionReport{msg}, sessionID)
	case "BP":
		return router.OnFIX50SP1SecurityDefinitionUpdateReport(SecurityDefinitionUpdateReport{msg}, sessionID)
	case "BT":
		return router.OnFIX50SP1MarketDefinitionRequest(MarketDefinitionRequest{msg}, sessionID)
	case "H":
		return router.OnFIX50SP1OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "AB":
		return router.OnFIX50SP1NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "AE":
		return router.OnFIX50SP1TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "AI":
		return router.OnFIX50SP1QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	case "BW":
		return router.OnFIX50SP1ApplicationMessageRequest(ApplicationMessageRequest{msg}, sessionID)
	case "9":
		return router.OnFIX50SP1OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "F":
		return router.OnFIX50SP1OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "M":
		return router.OnFIX50SP1ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "AJ":
		return router.OnFIX50SP1QuoteResponse(QuoteResponse{msg}, sessionID)
	case "P":
		return router.OnFIX50SP1AllocationInstructionAck(AllocationInstructionAck{msg}, sessionID)
	case "T":
		return router.OnFIX50SP1SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "x":
		return router.OnFIX50SP1SecurityListRequest(SecurityListRequest{msg}, sessionID)
	case "AG":
		return router.OnFIX50SP1QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "BD":
		return router.OnFIX50SP1NetworkCounterpartySystemStatusResponse(NetworkCounterpartySystemStatusResponse{msg}, sessionID)
	case "BZ":
		return router.OnFIX50SP1OrderMassActionReport(OrderMassActionReport{msg}, sessionID)
	}
	return quickfixgo.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX50SP1OrderCancelReject(msg OrderCancelReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1OrderCancelRequest(msg OrderCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ListStatusRequest(msg ListStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1QuoteResponse(msg QuoteResponse, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1AllocationInstructionAck(msg AllocationInstructionAck, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SettlementInstructions(msg SettlementInstructions, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SecurityListRequest(msg SecurityListRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1QuoteRequestReject(msg QuoteRequestReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1OrderMassActionReport(msg OrderMassActionReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1DontKnowTrade(msg DontKnowTrade, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1QuoteRequest(msg QuoteRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1Advertisement(msg Advertisement, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1Quote(msg Quote, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1Confirmation(msg Confirmation, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1RFQRequest(msg RFQRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1CollateralInquiryAck(msg CollateralInquiryAck, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1Email(msg Email, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1NewOrderCross(msg NewOrderCross, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1CollateralAssignment(msg CollateralAssignment, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1CollateralReport(msg CollateralReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1News(msg News, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1BusinessMessageReject(msg BusinessMessageReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1BidRequest(msg BidRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1AllocationInstructionAlert(msg AllocationInstructionAlert, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1MarketDefinition(msg MarketDefinition, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1OrderMassActionRequest(msg OrderMassActionRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1TradingSessionStatus(msg TradingSessionStatus, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SettlementObligationReport(msg SettlementObligationReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ListExecute(msg ListExecute, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1PositionReport(msg PositionReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ConfirmationRequest(msg ConfirmationRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1QuoteCancel(msg QuoteCancel, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SecurityStatus(msg SecurityStatus, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1CollateralResponse(msg CollateralResponse, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1MarketDataRequest(msg MarketDataRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1CollateralRequest(msg CollateralRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1UserRequest(msg UserRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ApplicationMessageReport(msg ApplicationMessageReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ListCancelRequest(msg ListCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1AllocationReport(msg AllocationReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1NewOrderList(msg NewOrderList, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SecurityDefinition(msg SecurityDefinition, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1AllocationReportAck(msg AllocationReportAck, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ExecutionReport(msg ExecutionReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1TradingSessionList(msg TradingSessionList, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1UserNotification(msg UserNotification, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1RequestForPositions(msg RequestForPositions, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1TradingSessionListRequest(msg TradingSessionListRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SecurityList(msg SecurityList, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ListStatus(msg ListStatus, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1RegistrationInstructions(msg RegistrationInstructions, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ConfirmationAck(msg ConfirmationAck, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1AdjustedPositionReport(msg AdjustedPositionReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1MassQuote(msg MassQuote, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1RequestForPositionsAck(msg RequestForPositionsAck, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SecurityListUpdateReport(msg SecurityListUpdateReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1BidResponse(msg BidResponse, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1CollateralInquiry(msg CollateralInquiry, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1NewOrderSingle(msg NewOrderSingle, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SecurityTypes(msg SecurityTypes, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1UserResponse(msg UserResponse, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1AssignmentReport(msg AssignmentReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1IOI(msg IOI, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1AllocationInstruction(msg AllocationInstruction, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ListStrikePrice(msg ListStrikePrice, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ContraryIntentionReport(msg ContraryIntentionReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1MarketDefinitionRequest(msg MarketDefinitionRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1OrderStatusRequest(msg OrderStatusRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1NewOrderMultileg(msg NewOrderMultileg, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1TradeCaptureReport(msg TradeCaptureReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1QuoteStatusReport(msg QuoteStatusReport, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
	OnFIX50SP1ApplicationMessageRequest(msg ApplicationMessageRequest, sessionID quickfixgo.SessionID) quickfixgo.MessageReject
}
type FIX50SP1MessageCracker struct{}

func (c *FIX50SP1MessageCracker) OnFIX50SP1ListCancelRequest(msg ListCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReportAck(msg TradeCaptureReportAck, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationReport(msg AllocationReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SettlementInstructionRequest(msg SettlementInstructionRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderList(msg NewOrderList, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityDefinition(msg SecurityDefinition, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationReportAck(msg AllocationReportAck, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ExecutionReport(msg ExecutionReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionList(msg TradingSessionList, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1UserNotification(msg UserNotification, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassCancelReport(msg OrderMassCancelReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RequestForPositions(msg RequestForPositions, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionListRequest(msg TradingSessionListRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityList(msg SecurityList, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListStatus(msg ListStatus, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RegistrationInstructions(msg RegistrationInstructions, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ConfirmationAck(msg ConfirmationAck, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AdjustedPositionReport(msg AdjustedPositionReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MassQuote(msg MassQuote, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RequestForPositionsAck(msg RequestForPositionsAck, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityListUpdateReport(msg SecurityListUpdateReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1BidResponse(msg BidResponse, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralInquiry(msg CollateralInquiry, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteStatusRequest(msg QuoteStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderSingle(msg NewOrderSingle, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityTypes(msg SecurityTypes, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1UserResponse(msg UserResponse, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AssignmentReport(msg AssignmentReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1IOI(msg IOI, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataRequestReject(msg MarketDataRequestReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationInstruction(msg AllocationInstruction, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityStatusRequest(msg SecurityStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListStrikePrice(msg ListStrikePrice, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1DerivativeSecurityList(msg DerivativeSecurityList, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ContraryIntentionReport(msg ContraryIntentionReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDefinitionRequest(msg MarketDefinitionRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderStatusRequest(msg OrderStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderMultileg(msg NewOrderMultileg, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReport(msg TradeCaptureReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteStatusReport(msg QuoteStatusReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ApplicationMessageRequest(msg ApplicationMessageRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderCancelReject(msg OrderCancelReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderCancelRequest(msg OrderCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListStatusRequest(msg ListStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteResponse(msg QuoteResponse, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationInstructionAck(msg AllocationInstructionAck, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SettlementInstructions(msg SettlementInstructions, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityListRequest(msg SecurityListRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteRequestReject(msg QuoteRequestReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassActionReport(msg OrderMassActionReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1DontKnowTrade(msg DontKnowTrade, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteRequest(msg QuoteRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1Advertisement(msg Advertisement, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1Quote(msg Quote, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1Confirmation(msg Confirmation, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RFQRequest(msg RFQRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralInquiryAck(msg CollateralInquiryAck, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1Email(msg Email, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderCross(msg NewOrderCross, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralAssignment(msg CollateralAssignment, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralReport(msg CollateralReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1News(msg News, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1BusinessMessageReject(msg BusinessMessageReject, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1BidRequest(msg BidRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationInstructionAlert(msg AllocationInstructionAlert, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDefinition(msg MarketDefinition, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassActionRequest(msg OrderMassActionRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityTypeRequest(msg SecurityTypeRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionStatus(msg TradingSessionStatus, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SettlementObligationReport(msg SettlementObligationReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListExecute(msg ListExecute, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1PositionReport(msg PositionReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ConfirmationRequest(msg ConfirmationRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteCancel(msg QuoteCancel, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1PositionMaintenanceReport(msg PositionMaintenanceReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityStatus(msg SecurityStatus, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralResponse(msg CollateralResponse, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataRequest(msg MarketDataRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralRequest(msg CollateralRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1UserRequest(msg UserRequest, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ApplicationMessageReport(msg ApplicationMessageReport, sessionId quickfixgo.SessionID) quickfixgo.MessageReject {
	return quickfixgo.NewUnsupportedMessageType(msg.Message)
}
