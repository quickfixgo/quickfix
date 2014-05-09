package fix50sp1

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

func Crack(msg message.Message, sessionID quickfix.SessionID, router MessageRouter) errors.MessageRejectError {

	msgType := &field.MsgTypeField{}
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

//OnFIX50SP1IOI is a Callback for IOI messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1IOI(msg IOI, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1Advertisement is a Callback for Advertisement messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1Advertisement(msg Advertisement, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ExecutionReport is a Callback for ExecutionReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1OrderCancelReject is a Callback for OrderCancelReject messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1DerivativeSecurityList is a Callback for DerivativeSecurityList messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1NewOrderMultileg is a Callback for NewOrderMultileg messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderMultileg(msg NewOrderMultileg, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1MultilegOrderCancelReplace is a Callback for MultilegOrderCancelReplace messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1TradeCaptureReportRequest is a Callback for TradeCaptureReportRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1TradeCaptureReport is a Callback for TradeCaptureReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReport(msg TradeCaptureReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1OrderMassStatusRequest is a Callback for OrderMassStatusRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1QuoteRequestReject is a Callback for QuoteRequestReject messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteRequestReject(msg QuoteRequestReject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1RFQRequest is a Callback for RFQRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1RFQRequest(msg RFQRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1QuoteStatusReport is a Callback for QuoteStatusReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteStatusReport(msg QuoteStatusReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1QuoteResponse is a Callback for QuoteResponse messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteResponse(msg QuoteResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1Confirmation is a Callback for Confirmation messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1Confirmation(msg Confirmation, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1PositionMaintenanceRequest is a Callback for PositionMaintenanceRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1PositionMaintenanceReport is a Callback for PositionMaintenanceReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1RequestForPositions is a Callback for RequestForPositions messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1RequestForPositions(msg RequestForPositions, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1RequestForPositionsAck is a Callback for RequestForPositionsAck messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1RequestForPositionsAck(msg RequestForPositionsAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1PositionReport is a Callback for PositionReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1PositionReport(msg PositionReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1TradeCaptureReportRequestAck is a Callback for TradeCaptureReportRequestAck messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1TradeCaptureReportAck is a Callback for TradeCaptureReportAck messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1AllocationReport is a Callback for AllocationReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationReport(msg AllocationReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1AllocationReportAck is a Callback for AllocationReportAck messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationReportAck(msg AllocationReportAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ConfirmationAck is a Callback for ConfirmationAck messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ConfirmationAck(msg ConfirmationAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SettlementInstructionRequest is a Callback for SettlementInstructionRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1AssignmentReport is a Callback for AssignmentReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1AssignmentReport(msg AssignmentReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1CollateralRequest is a Callback for CollateralRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralRequest(msg CollateralRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1CollateralAssignment is a Callback for CollateralAssignment messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralAssignment(msg CollateralAssignment, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1CollateralResponse is a Callback for CollateralResponse messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralResponse(msg CollateralResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1News is a Callback for News messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1News(msg News, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1CollateralReport is a Callback for CollateralReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralReport(msg CollateralReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1CollateralInquiry is a Callback for CollateralInquiry messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralInquiry(msg CollateralInquiry, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1NetworkCounterpartySystemStatusRequest is a Callback for NetworkCounterpartySystemStatusRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1NetworkCounterpartySystemStatusResponse is a Callback for NetworkCounterpartySystemStatusResponse messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1UserRequest is a Callback for UserRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1UserRequest(msg UserRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1UserResponse is a Callback for UserResponse messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1UserResponse(msg UserResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1CollateralInquiryAck is a Callback for CollateralInquiryAck messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1CollateralInquiryAck(msg CollateralInquiryAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ConfirmationRequest is a Callback for ConfirmationRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ConfirmationRequest(msg ConfirmationRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1TradingSessionListRequest is a Callback for TradingSessionListRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionListRequest(msg TradingSessionListRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1TradingSessionList is a Callback for TradingSessionList messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionList(msg TradingSessionList, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SecurityListUpdateReport is a Callback for SecurityListUpdateReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityListUpdateReport(msg SecurityListUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1AdjustedPositionReport is a Callback for AdjustedPositionReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1AdjustedPositionReport(msg AdjustedPositionReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1AllocationInstructionAlert is a Callback for AllocationInstructionAlert messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationInstructionAlert(msg AllocationInstructionAlert, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ExecutionAcknowledgement is a Callback for ExecutionAcknowledgement messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ContraryIntentionReport is a Callback for ContraryIntentionReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ContraryIntentionReport(msg ContraryIntentionReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SecurityDefinitionUpdateReport is a Callback for SecurityDefinitionUpdateReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SettlementObligationReport is a Callback for SettlementObligationReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SettlementObligationReport(msg SettlementObligationReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1DerivativeSecurityListUpdateReport is a Callback for DerivativeSecurityListUpdateReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1TradingSessionListUpdateReport is a Callback for TradingSessionListUpdateReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1MarketDefinitionRequest is a Callback for MarketDefinitionRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDefinitionRequest(msg MarketDefinitionRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1MarketDefinition is a Callback for MarketDefinition messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDefinition(msg MarketDefinition, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1MarketDefinitionUpdateReport is a Callback for MarketDefinitionUpdateReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ApplicationMessageRequest is a Callback for ApplicationMessageRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ApplicationMessageRequest(msg ApplicationMessageRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ApplicationMessageRequestAck is a Callback for ApplicationMessageRequestAck messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ApplicationMessageReport is a Callback for ApplicationMessageReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ApplicationMessageReport(msg ApplicationMessageReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1OrderMassActionReport is a Callback for OrderMassActionReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassActionReport(msg OrderMassActionReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1Email is a Callback for Email messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1Email(msg Email, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1OrderMassActionRequest is a Callback for OrderMassActionRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassActionRequest(msg OrderMassActionRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1UserNotification is a Callback for UserNotification messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1UserNotification(msg UserNotification, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1NewOrderSingle is a Callback for NewOrderSingle messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1NewOrderList is a Callback for NewOrderList messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1OrderCancelRequest is a Callback for OrderCancelRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1OrderCancelReplaceRequest is a Callback for OrderCancelReplaceRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1OrderStatusRequest is a Callback for OrderStatusRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1AllocationInstruction is a Callback for AllocationInstruction messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationInstruction(msg AllocationInstruction, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ListCancelRequest is a Callback for ListCancelRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ListExecute is a Callback for ListExecute messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListExecute(msg ListExecute, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ListStatusRequest is a Callback for ListStatusRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ListStatus is a Callback for ListStatus messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListStatus(msg ListStatus, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1AllocationInstructionAck is a Callback for AllocationInstructionAck messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1AllocationInstructionAck(msg AllocationInstructionAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1DontKnowTrade is a Callback for DontKnowTrade messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1QuoteRequest is a Callback for QuoteRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1Quote is a Callback for Quote messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1Quote(msg Quote, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SettlementInstructions is a Callback for SettlementInstructions messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SettlementInstructions(msg SettlementInstructions, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1MarketDataRequest is a Callback for MarketDataRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataRequest(msg MarketDataRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1MarketDataSnapshotFullRefresh is a Callback for MarketDataSnapshotFullRefresh messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1MarketDataIncrementalRefresh is a Callback for MarketDataIncrementalRefresh messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1MarketDataRequestReject is a Callback for MarketDataRequestReject messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1QuoteCancel is a Callback for QuoteCancel messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteCancel(msg QuoteCancel, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1QuoteStatusRequest is a Callback for QuoteStatusRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1MassQuoteAcknowledgement is a Callback for MassQuoteAcknowledgement messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SecurityDefinitionRequest is a Callback for SecurityDefinitionRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SecurityDefinition is a Callback for SecurityDefinition messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityDefinition(msg SecurityDefinition, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SecurityStatusRequest is a Callback for SecurityStatusRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SecurityStatus is a Callback for SecurityStatus messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityStatus(msg SecurityStatus, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1TradingSessionStatusRequest is a Callback for TradingSessionStatusRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1TradingSessionStatus is a Callback for TradingSessionStatus messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1TradingSessionStatus(msg TradingSessionStatus, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1MassQuote is a Callback for MassQuote messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1MassQuote(msg MassQuote, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1BusinessMessageReject is a Callback for BusinessMessageReject messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1BusinessMessageReject(msg BusinessMessageReject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1BidRequest is a Callback for BidRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1BidRequest(msg BidRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1BidResponse is a Callback for BidResponse messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1BidResponse(msg BidResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1ListStrikePrice is a Callback for ListStrikePrice messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1ListStrikePrice(msg ListStrikePrice, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1RegistrationInstructions is a Callback for RegistrationInstructions messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1RegistrationInstructions(msg RegistrationInstructions, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1RegistrationInstructionsResponse is a Callback for RegistrationInstructionsResponse messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1OrderMassCancelRequest is a Callback for OrderMassCancelRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1OrderMassCancelReport is a Callback for OrderMassCancelReport messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1NewOrderCross is a Callback for NewOrderCross messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1NewOrderCross(msg NewOrderCross, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1CrossOrderCancelReplaceRequest is a Callback for CrossOrderCancelReplaceRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1CrossOrderCancelRequest is a Callback for CrossOrderCancelRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SecurityTypeRequest is a Callback for SecurityTypeRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SecurityTypes is a Callback for SecurityTypes messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityTypes(msg SecurityTypes, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SecurityListRequest is a Callback for SecurityListRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityListRequest(msg SecurityListRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1SecurityList is a Callback for SecurityList messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1SecurityList(msg SecurityList, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP1DerivativeSecurityListRequest is a Callback for DerivativeSecurityListRequest messages.
func (c *FIX50SP1MessageCracker) OnFIX50SP1DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
