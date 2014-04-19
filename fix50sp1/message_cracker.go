package fix50sp1

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

func Crack(msg message.Message, sessionID quickfix.SessionID, router MessageRouter) errors.MessageRejectError {

	msgType := new(field.MsgType)
	switch msg.Header.Get(msgType); msgType.Value {
	case "6":
		return router.OnFIX50SP1IOI(IOI{msg}, sessionID)
	case "7":
		return router.OnFIX50SP1Advertisement(Advertisement{msg}, sessionID)
	case "8":
		return router.OnFIX50SP1ExecutionReport(ExecutionReport{msg}, sessionID)
	case "9":
		return router.OnFIX50SP1OrderCancelReject(OrderCancelReject{msg}, sessionID)
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
	case "B":
		return router.OnFIX50SP1News(News{msg}, sessionID)
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
	case "BI":
		return router.OnFIX50SP1TradingSessionListRequest(TradingSessionListRequest{msg}, sessionID)
	case "BJ":
		return router.OnFIX50SP1TradingSessionList(TradingSessionList{msg}, sessionID)
	case "BK":
		return router.OnFIX50SP1SecurityListUpdateReport(SecurityListUpdateReport{msg}, sessionID)
	case "BL":
		return router.OnFIX50SP1AdjustedPositionReport(AdjustedPositionReport{msg}, sessionID)
	case "BM":
		return router.OnFIX50SP1AllocationInstructionAlert(AllocationInstructionAlert{msg}, sessionID)
	case "BN":
		return router.OnFIX50SP1ExecutionAcknowledgement(ExecutionAcknowledgement{msg}, sessionID)
	case "BO":
		return router.OnFIX50SP1ContraryIntentionReport(ContraryIntentionReport{msg}, sessionID)
	case "BP":
		return router.OnFIX50SP1SecurityDefinitionUpdateReport(SecurityDefinitionUpdateReport{msg}, sessionID)
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
	case "C":
		return router.OnFIX50SP1Email(Email{msg}, sessionID)
	case "CA":
		return router.OnFIX50SP1OrderMassActionRequest(OrderMassActionRequest{msg}, sessionID)
	case "CB":
		return router.OnFIX50SP1UserNotification(UserNotification{msg}, sessionID)
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
	}
	return errors.InvalidMessageType()
}

type MessageRouter interface {
	OnFIX50SP1IOI(msg IOI, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1Advertisement(msg Advertisement, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1NewOrderMultileg(msg NewOrderMultileg, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1TradeCaptureReport(msg TradeCaptureReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1QuoteRequestReject(msg QuoteRequestReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1RFQRequest(msg RFQRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1QuoteStatusReport(msg QuoteStatusReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1QuoteResponse(msg QuoteResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1Confirmation(msg Confirmation, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1RequestForPositions(msg RequestForPositions, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1RequestForPositionsAck(msg RequestForPositionsAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1PositionReport(msg PositionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1AllocationReport(msg AllocationReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1AllocationReportAck(msg AllocationReportAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ConfirmationAck(msg ConfirmationAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1AssignmentReport(msg AssignmentReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1CollateralRequest(msg CollateralRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1CollateralAssignment(msg CollateralAssignment, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1CollateralResponse(msg CollateralResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1News(msg News, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1CollateralReport(msg CollateralReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1CollateralInquiry(msg CollateralInquiry, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1UserRequest(msg UserRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1UserResponse(msg UserResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1CollateralInquiryAck(msg CollateralInquiryAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ConfirmationRequest(msg ConfirmationRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1TradingSessionListRequest(msg TradingSessionListRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1TradingSessionList(msg TradingSessionList, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SecurityListUpdateReport(msg SecurityListUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1AdjustedPositionReport(msg AdjustedPositionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1AllocationInstructionAlert(msg AllocationInstructionAlert, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ContraryIntentionReport(msg ContraryIntentionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SettlementObligationReport(msg SettlementObligationReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1MarketDefinitionRequest(msg MarketDefinitionRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1MarketDefinition(msg MarketDefinition, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ApplicationMessageRequest(msg ApplicationMessageRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ApplicationMessageReport(msg ApplicationMessageReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1OrderMassActionReport(msg OrderMassActionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1Email(msg Email, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1OrderMassActionRequest(msg OrderMassActionRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1UserNotification(msg UserNotification, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1AllocationInstruction(msg AllocationInstruction, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ListExecute(msg ListExecute, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ListStatus(msg ListStatus, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1AllocationInstructionAck(msg AllocationInstructionAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1Quote(msg Quote, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SettlementInstructions(msg SettlementInstructions, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1MarketDataRequest(msg MarketDataRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1QuoteCancel(msg QuoteCancel, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SecurityDefinition(msg SecurityDefinition, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SecurityStatus(msg SecurityStatus, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1TradingSessionStatus(msg TradingSessionStatus, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1MassQuote(msg MassQuote, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1BusinessMessageReject(msg BusinessMessageReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1BidRequest(msg BidRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1BidResponse(msg BidResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1ListStrikePrice(msg ListStrikePrice, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1RegistrationInstructions(msg RegistrationInstructions, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1NewOrderCross(msg NewOrderCross, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SecurityTypes(msg SecurityTypes, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SecurityListRequest(msg SecurityListRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1SecurityList(msg SecurityList, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP1DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfix.SessionID) errors.MessageRejectError
}
type FIX50SP1MessageCracker struct{}

func (c *FIX50SP1MessageCracker) OnFIX50SP1IOI(msg IOI, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1Advertisement(msg Advertisement, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ExecutionReport(msg ExecutionReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderCancelReject(msg OrderCancelReject, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1DerivativeSecurityList(msg DerivativeSecurityList, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderMultileg(msg NewOrderMultileg, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReport(msg TradeCaptureReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteRequestReject(msg QuoteRequestReject, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RFQRequest(msg RFQRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteStatusReport(msg QuoteStatusReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteResponse(msg QuoteResponse, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1Confirmation(msg Confirmation, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1PositionMaintenanceReport(msg PositionMaintenanceReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RequestForPositions(msg RequestForPositions, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RequestForPositionsAck(msg RequestForPositionsAck, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1PositionReport(msg PositionReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReportAck(msg TradeCaptureReportAck, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationReport(msg AllocationReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationReportAck(msg AllocationReportAck, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ConfirmationAck(msg ConfirmationAck, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SettlementInstructionRequest(msg SettlementInstructionRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AssignmentReport(msg AssignmentReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralRequest(msg CollateralRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralAssignment(msg CollateralAssignment, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralResponse(msg CollateralResponse, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1News(msg News, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralReport(msg CollateralReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralInquiry(msg CollateralInquiry, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1UserRequest(msg UserRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1UserResponse(msg UserResponse, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralInquiryAck(msg CollateralInquiryAck, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ConfirmationRequest(msg ConfirmationRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionListRequest(msg TradingSessionListRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionList(msg TradingSessionList, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityListUpdateReport(msg SecurityListUpdateReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AdjustedPositionReport(msg AdjustedPositionReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationInstructionAlert(msg AllocationInstructionAlert, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ContraryIntentionReport(msg ContraryIntentionReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SettlementObligationReport(msg SettlementObligationReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDefinitionRequest(msg MarketDefinitionRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDefinition(msg MarketDefinition, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ApplicationMessageRequest(msg ApplicationMessageRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ApplicationMessageReport(msg ApplicationMessageReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassActionReport(msg OrderMassActionReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1Email(msg Email, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassActionRequest(msg OrderMassActionRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1UserNotification(msg UserNotification, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderSingle(msg NewOrderSingle, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderList(msg NewOrderList, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderCancelRequest(msg OrderCancelRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderStatusRequest(msg OrderStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationInstruction(msg AllocationInstruction, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListCancelRequest(msg ListCancelRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListExecute(msg ListExecute, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListStatusRequest(msg ListStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListStatus(msg ListStatus, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationInstructionAck(msg AllocationInstructionAck, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1DontKnowTrade(msg DontKnowTrade, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteRequest(msg QuoteRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1Quote(msg Quote, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SettlementInstructions(msg SettlementInstructions, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataRequest(msg MarketDataRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataRequestReject(msg MarketDataRequestReject, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteCancel(msg QuoteCancel, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteStatusRequest(msg QuoteStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityDefinition(msg SecurityDefinition, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityStatusRequest(msg SecurityStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityStatus(msg SecurityStatus, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionStatus(msg TradingSessionStatus, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1MassQuote(msg MassQuote, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1BusinessMessageReject(msg BusinessMessageReject, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1BidRequest(msg BidRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1BidResponse(msg BidResponse, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListStrikePrice(msg ListStrikePrice, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RegistrationInstructions(msg RegistrationInstructions, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassCancelReport(msg OrderMassCancelReport, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderCross(msg NewOrderCross, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityTypeRequest(msg SecurityTypeRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityTypes(msg SecurityTypes, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityListRequest(msg SecurityListRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityList(msg SecurityList, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
func (c *FIX50SP1MessageCracker) OnFIX50SP1DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
