package fix50sp2

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
		return router.OnFIX50SP2IOI(IOI{msg}, sessionID)
	case "7":
		return router.OnFIX50SP2Advertisement(Advertisement{msg}, sessionID)
	case "8":
		return router.OnFIX50SP2ExecutionReport(ExecutionReport{msg}, sessionID)
	case "9":
		return router.OnFIX50SP2OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "AA":
		return router.OnFIX50SP2DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "AB":
		return router.OnFIX50SP2NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "AC":
		return router.OnFIX50SP2MultilegOrderCancelReplace(MultilegOrderCancelReplace{msg}, sessionID)
	case "AD":
		return router.OnFIX50SP2TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "AE":
		return router.OnFIX50SP2TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "AF":
		return router.OnFIX50SP2OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "AG":
		return router.OnFIX50SP2QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "AH":
		return router.OnFIX50SP2RFQRequest(RFQRequest{msg}, sessionID)
	case "AI":
		return router.OnFIX50SP2QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	case "AJ":
		return router.OnFIX50SP2QuoteResponse(QuoteResponse{msg}, sessionID)
	case "AK":
		return router.OnFIX50SP2Confirmation(Confirmation{msg}, sessionID)
	case "AL":
		return router.OnFIX50SP2PositionMaintenanceRequest(PositionMaintenanceRequest{msg}, sessionID)
	case "AM":
		return router.OnFIX50SP2PositionMaintenanceReport(PositionMaintenanceReport{msg}, sessionID)
	case "AN":
		return router.OnFIX50SP2RequestForPositions(RequestForPositions{msg}, sessionID)
	case "AO":
		return router.OnFIX50SP2RequestForPositionsAck(RequestForPositionsAck{msg}, sessionID)
	case "AP":
		return router.OnFIX50SP2PositionReport(PositionReport{msg}, sessionID)
	case "AQ":
		return router.OnFIX50SP2TradeCaptureReportRequestAck(TradeCaptureReportRequestAck{msg}, sessionID)
	case "AR":
		return router.OnFIX50SP2TradeCaptureReportAck(TradeCaptureReportAck{msg}, sessionID)
	case "AS":
		return router.OnFIX50SP2AllocationReport(AllocationReport{msg}, sessionID)
	case "AT":
		return router.OnFIX50SP2AllocationReportAck(AllocationReportAck{msg}, sessionID)
	case "AU":
		return router.OnFIX50SP2ConfirmationAck(ConfirmationAck{msg}, sessionID)
	case "AV":
		return router.OnFIX50SP2SettlementInstructionRequest(SettlementInstructionRequest{msg}, sessionID)
	case "AW":
		return router.OnFIX50SP2AssignmentReport(AssignmentReport{msg}, sessionID)
	case "AX":
		return router.OnFIX50SP2CollateralRequest(CollateralRequest{msg}, sessionID)
	case "AY":
		return router.OnFIX50SP2CollateralAssignment(CollateralAssignment{msg}, sessionID)
	case "AZ":
		return router.OnFIX50SP2CollateralResponse(CollateralResponse{msg}, sessionID)
	case "B":
		return router.OnFIX50SP2News(News{msg}, sessionID)
	case "BA":
		return router.OnFIX50SP2CollateralReport(CollateralReport{msg}, sessionID)
	case "BB":
		return router.OnFIX50SP2CollateralInquiry(CollateralInquiry{msg}, sessionID)
	case "BC":
		return router.OnFIX50SP2NetworkCounterpartySystemStatusRequest(NetworkCounterpartySystemStatusRequest{msg}, sessionID)
	case "BD":
		return router.OnFIX50SP2NetworkCounterpartySystemStatusResponse(NetworkCounterpartySystemStatusResponse{msg}, sessionID)
	case "BE":
		return router.OnFIX50SP2UserRequest(UserRequest{msg}, sessionID)
	case "BF":
		return router.OnFIX50SP2UserResponse(UserResponse{msg}, sessionID)
	case "BG":
		return router.OnFIX50SP2CollateralInquiryAck(CollateralInquiryAck{msg}, sessionID)
	case "BH":
		return router.OnFIX50SP2ConfirmationRequest(ConfirmationRequest{msg}, sessionID)
	case "BI":
		return router.OnFIX50SP2TradingSessionListRequest(TradingSessionListRequest{msg}, sessionID)
	case "BJ":
		return router.OnFIX50SP2TradingSessionList(TradingSessionList{msg}, sessionID)
	case "BK":
		return router.OnFIX50SP2SecurityListUpdateReport(SecurityListUpdateReport{msg}, sessionID)
	case "BL":
		return router.OnFIX50SP2AdjustedPositionReport(AdjustedPositionReport{msg}, sessionID)
	case "BM":
		return router.OnFIX50SP2AllocationInstructionAlert(AllocationInstructionAlert{msg}, sessionID)
	case "BN":
		return router.OnFIX50SP2ExecutionAcknowledgement(ExecutionAcknowledgement{msg}, sessionID)
	case "BO":
		return router.OnFIX50SP2ContraryIntentionReport(ContraryIntentionReport{msg}, sessionID)
	case "BP":
		return router.OnFIX50SP2SecurityDefinitionUpdateReport(SecurityDefinitionUpdateReport{msg}, sessionID)
	case "BQ":
		return router.OnFIX50SP2SettlementObligationReport(SettlementObligationReport{msg}, sessionID)
	case "BR":
		return router.OnFIX50SP2DerivativeSecurityListUpdateReport(DerivativeSecurityListUpdateReport{msg}, sessionID)
	case "BS":
		return router.OnFIX50SP2TradingSessionListUpdateReport(TradingSessionListUpdateReport{msg}, sessionID)
	case "BT":
		return router.OnFIX50SP2MarketDefinitionRequest(MarketDefinitionRequest{msg}, sessionID)
	case "BU":
		return router.OnFIX50SP2MarketDefinition(MarketDefinition{msg}, sessionID)
	case "BV":
		return router.OnFIX50SP2MarketDefinitionUpdateReport(MarketDefinitionUpdateReport{msg}, sessionID)
	case "BW":
		return router.OnFIX50SP2ApplicationMessageRequest(ApplicationMessageRequest{msg}, sessionID)
	case "BX":
		return router.OnFIX50SP2ApplicationMessageRequestAck(ApplicationMessageRequestAck{msg}, sessionID)
	case "BY":
		return router.OnFIX50SP2ApplicationMessageReport(ApplicationMessageReport{msg}, sessionID)
	case "BZ":
		return router.OnFIX50SP2OrderMassActionReport(OrderMassActionReport{msg}, sessionID)
	case "C":
		return router.OnFIX50SP2Email(Email{msg}, sessionID)
	case "CA":
		return router.OnFIX50SP2OrderMassActionRequest(OrderMassActionRequest{msg}, sessionID)
	case "CB":
		return router.OnFIX50SP2UserNotification(UserNotification{msg}, sessionID)
	case "CC":
		return router.OnFIX50SP2StreamAssignmentRequest(StreamAssignmentRequest{msg}, sessionID)
	case "CD":
		return router.OnFIX50SP2StreamAssignmentReport(StreamAssignmentReport{msg}, sessionID)
	case "CE":
		return router.OnFIX50SP2StreamAssignmentReportACK(StreamAssignmentReportACK{msg}, sessionID)
	case "CF":
		return router.OnFIX50SP2PartyDetailsListRequest(PartyDetailsListRequest{msg}, sessionID)
	case "CG":
		return router.OnFIX50SP2PartyDetailsListReport(PartyDetailsListReport{msg}, sessionID)
	case "D":
		return router.OnFIX50SP2NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "E":
		return router.OnFIX50SP2NewOrderList(NewOrderList{msg}, sessionID)
	case "F":
		return router.OnFIX50SP2OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "G":
		return router.OnFIX50SP2OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "H":
		return router.OnFIX50SP2OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "J":
		return router.OnFIX50SP2AllocationInstruction(AllocationInstruction{msg}, sessionID)
	case "K":
		return router.OnFIX50SP2ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "L":
		return router.OnFIX50SP2ListExecute(ListExecute{msg}, sessionID)
	case "M":
		return router.OnFIX50SP2ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "N":
		return router.OnFIX50SP2ListStatus(ListStatus{msg}, sessionID)
	case "P":
		return router.OnFIX50SP2AllocationInstructionAck(AllocationInstructionAck{msg}, sessionID)
	case "Q":
		return router.OnFIX50SP2DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "R":
		return router.OnFIX50SP2QuoteRequest(QuoteRequest{msg}, sessionID)
	case "S":
		return router.OnFIX50SP2Quote(Quote{msg}, sessionID)
	case "T":
		return router.OnFIX50SP2SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "V":
		return router.OnFIX50SP2MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "W":
		return router.OnFIX50SP2MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "X":
		return router.OnFIX50SP2MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "Y":
		return router.OnFIX50SP2MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "Z":
		return router.OnFIX50SP2QuoteCancel(QuoteCancel{msg}, sessionID)
	case "a":
		return router.OnFIX50SP2QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "b":
		return router.OnFIX50SP2MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "c":
		return router.OnFIX50SP2SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "d":
		return router.OnFIX50SP2SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "e":
		return router.OnFIX50SP2SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "f":
		return router.OnFIX50SP2SecurityStatus(SecurityStatus{msg}, sessionID)
	case "g":
		return router.OnFIX50SP2TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "h":
		return router.OnFIX50SP2TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "i":
		return router.OnFIX50SP2MassQuote(MassQuote{msg}, sessionID)
	case "j":
		return router.OnFIX50SP2BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "k":
		return router.OnFIX50SP2BidRequest(BidRequest{msg}, sessionID)
	case "l":
		return router.OnFIX50SP2BidResponse(BidResponse{msg}, sessionID)
	case "m":
		return router.OnFIX50SP2ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "o":
		return router.OnFIX50SP2RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "p":
		return router.OnFIX50SP2RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "q":
		return router.OnFIX50SP2OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "r":
		return router.OnFIX50SP2OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "s":
		return router.OnFIX50SP2NewOrderCross(NewOrderCross{msg}, sessionID)
	case "t":
		return router.OnFIX50SP2CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "u":
		return router.OnFIX50SP2CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "v":
		return router.OnFIX50SP2SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "w":
		return router.OnFIX50SP2SecurityTypes(SecurityTypes{msg}, sessionID)
	case "x":
		return router.OnFIX50SP2SecurityListRequest(SecurityListRequest{msg}, sessionID)
	case "y":
		return router.OnFIX50SP2SecurityList(SecurityList{msg}, sessionID)
	case "z":
		return router.OnFIX50SP2DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	}
	return errors.InvalidMessageType()
}

type MessageRouter interface {
	OnFIX50SP2IOI(msg IOI, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2Advertisement(msg Advertisement, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2NewOrderMultileg(msg NewOrderMultileg, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2TradeCaptureReport(msg TradeCaptureReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2QuoteRequestReject(msg QuoteRequestReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2RFQRequest(msg RFQRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2QuoteStatusReport(msg QuoteStatusReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2QuoteResponse(msg QuoteResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2Confirmation(msg Confirmation, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2RequestForPositions(msg RequestForPositions, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2RequestForPositionsAck(msg RequestForPositionsAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2PositionReport(msg PositionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2AllocationReport(msg AllocationReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2AllocationReportAck(msg AllocationReportAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ConfirmationAck(msg ConfirmationAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2AssignmentReport(msg AssignmentReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2CollateralRequest(msg CollateralRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2CollateralAssignment(msg CollateralAssignment, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2CollateralResponse(msg CollateralResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2News(msg News, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2CollateralReport(msg CollateralReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2CollateralInquiry(msg CollateralInquiry, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2UserRequest(msg UserRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2UserResponse(msg UserResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2CollateralInquiryAck(msg CollateralInquiryAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ConfirmationRequest(msg ConfirmationRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2TradingSessionListRequest(msg TradingSessionListRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2TradingSessionList(msg TradingSessionList, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SecurityListUpdateReport(msg SecurityListUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2AdjustedPositionReport(msg AdjustedPositionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2AllocationInstructionAlert(msg AllocationInstructionAlert, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ContraryIntentionReport(msg ContraryIntentionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SettlementObligationReport(msg SettlementObligationReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2MarketDefinitionRequest(msg MarketDefinitionRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2MarketDefinition(msg MarketDefinition, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ApplicationMessageRequest(msg ApplicationMessageRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ApplicationMessageReport(msg ApplicationMessageReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2OrderMassActionReport(msg OrderMassActionReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2Email(msg Email, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2OrderMassActionRequest(msg OrderMassActionRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2UserNotification(msg UserNotification, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2StreamAssignmentRequest(msg StreamAssignmentRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2StreamAssignmentReport(msg StreamAssignmentReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2StreamAssignmentReportACK(msg StreamAssignmentReportACK, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2PartyDetailsListRequest(msg PartyDetailsListRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2PartyDetailsListReport(msg PartyDetailsListReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2AllocationInstruction(msg AllocationInstruction, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ListExecute(msg ListExecute, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ListStatus(msg ListStatus, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2AllocationInstructionAck(msg AllocationInstructionAck, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2Quote(msg Quote, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SettlementInstructions(msg SettlementInstructions, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2MarketDataRequest(msg MarketDataRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2QuoteCancel(msg QuoteCancel, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SecurityDefinition(msg SecurityDefinition, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SecurityStatus(msg SecurityStatus, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2TradingSessionStatus(msg TradingSessionStatus, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2MassQuote(msg MassQuote, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2BusinessMessageReject(msg BusinessMessageReject, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2BidRequest(msg BidRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2BidResponse(msg BidResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2ListStrikePrice(msg ListStrikePrice, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2RegistrationInstructions(msg RegistrationInstructions, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2NewOrderCross(msg NewOrderCross, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SecurityTypes(msg SecurityTypes, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SecurityListRequest(msg SecurityListRequest, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2SecurityList(msg SecurityList, sessionID quickfix.SessionID) errors.MessageRejectError
	OnFIX50SP2DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfix.SessionID) errors.MessageRejectError
}
type FIX50SP2MessageCracker struct{}

//OnFIX50SP2IOI is a Callback for IOI messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2IOI(msg IOI, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2Advertisement is a Callback for Advertisement messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2Advertisement(msg Advertisement, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ExecutionReport is a Callback for ExecutionReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2OrderCancelReject is a Callback for OrderCancelReject messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2DerivativeSecurityList is a Callback for DerivativeSecurityList messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2NewOrderMultileg is a Callback for NewOrderMultileg messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2NewOrderMultileg(msg NewOrderMultileg, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2MultilegOrderCancelReplace is a Callback for MultilegOrderCancelReplace messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2TradeCaptureReportRequest is a Callback for TradeCaptureReportRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2TradeCaptureReport is a Callback for TradeCaptureReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradeCaptureReport(msg TradeCaptureReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2OrderMassStatusRequest is a Callback for OrderMassStatusRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2QuoteRequestReject is a Callback for QuoteRequestReject messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteRequestReject(msg QuoteRequestReject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2RFQRequest is a Callback for RFQRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2RFQRequest(msg RFQRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2QuoteStatusReport is a Callback for QuoteStatusReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteStatusReport(msg QuoteStatusReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2QuoteResponse is a Callback for QuoteResponse messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteResponse(msg QuoteResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2Confirmation is a Callback for Confirmation messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2Confirmation(msg Confirmation, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2PositionMaintenanceRequest is a Callback for PositionMaintenanceRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2PositionMaintenanceReport is a Callback for PositionMaintenanceReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2RequestForPositions is a Callback for RequestForPositions messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2RequestForPositions(msg RequestForPositions, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2RequestForPositionsAck is a Callback for RequestForPositionsAck messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2RequestForPositionsAck(msg RequestForPositionsAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2PositionReport is a Callback for PositionReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2PositionReport(msg PositionReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2TradeCaptureReportRequestAck is a Callback for TradeCaptureReportRequestAck messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2TradeCaptureReportAck is a Callback for TradeCaptureReportAck messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2AllocationReport is a Callback for AllocationReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationReport(msg AllocationReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2AllocationReportAck is a Callback for AllocationReportAck messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationReportAck(msg AllocationReportAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ConfirmationAck is a Callback for ConfirmationAck messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ConfirmationAck(msg ConfirmationAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SettlementInstructionRequest is a Callback for SettlementInstructionRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2AssignmentReport is a Callback for AssignmentReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2AssignmentReport(msg AssignmentReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2CollateralRequest is a Callback for CollateralRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralRequest(msg CollateralRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2CollateralAssignment is a Callback for CollateralAssignment messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralAssignment(msg CollateralAssignment, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2CollateralResponse is a Callback for CollateralResponse messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralResponse(msg CollateralResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2News is a Callback for News messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2News(msg News, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2CollateralReport is a Callback for CollateralReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralReport(msg CollateralReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2CollateralInquiry is a Callback for CollateralInquiry messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralInquiry(msg CollateralInquiry, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2NetworkCounterpartySystemStatusRequest is a Callback for NetworkCounterpartySystemStatusRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2NetworkCounterpartySystemStatusResponse is a Callback for NetworkCounterpartySystemStatusResponse messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2UserRequest is a Callback for UserRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2UserRequest(msg UserRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2UserResponse is a Callback for UserResponse messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2UserResponse(msg UserResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2CollateralInquiryAck is a Callback for CollateralInquiryAck messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralInquiryAck(msg CollateralInquiryAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ConfirmationRequest is a Callback for ConfirmationRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ConfirmationRequest(msg ConfirmationRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2TradingSessionListRequest is a Callback for TradingSessionListRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionListRequest(msg TradingSessionListRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2TradingSessionList is a Callback for TradingSessionList messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionList(msg TradingSessionList, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SecurityListUpdateReport is a Callback for SecurityListUpdateReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityListUpdateReport(msg SecurityListUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2AdjustedPositionReport is a Callback for AdjustedPositionReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2AdjustedPositionReport(msg AdjustedPositionReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2AllocationInstructionAlert is a Callback for AllocationInstructionAlert messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationInstructionAlert(msg AllocationInstructionAlert, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ExecutionAcknowledgement is a Callback for ExecutionAcknowledgement messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ContraryIntentionReport is a Callback for ContraryIntentionReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ContraryIntentionReport(msg ContraryIntentionReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SecurityDefinitionUpdateReport is a Callback for SecurityDefinitionUpdateReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SettlementObligationReport is a Callback for SettlementObligationReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SettlementObligationReport(msg SettlementObligationReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2DerivativeSecurityListUpdateReport is a Callback for DerivativeSecurityListUpdateReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2TradingSessionListUpdateReport is a Callback for TradingSessionListUpdateReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2MarketDefinitionRequest is a Callback for MarketDefinitionRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDefinitionRequest(msg MarketDefinitionRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2MarketDefinition is a Callback for MarketDefinition messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDefinition(msg MarketDefinition, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2MarketDefinitionUpdateReport is a Callback for MarketDefinitionUpdateReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ApplicationMessageRequest is a Callback for ApplicationMessageRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ApplicationMessageRequest(msg ApplicationMessageRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ApplicationMessageRequestAck is a Callback for ApplicationMessageRequestAck messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ApplicationMessageReport is a Callback for ApplicationMessageReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ApplicationMessageReport(msg ApplicationMessageReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2OrderMassActionReport is a Callback for OrderMassActionReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassActionReport(msg OrderMassActionReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2Email is a Callback for Email messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2Email(msg Email, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2OrderMassActionRequest is a Callback for OrderMassActionRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassActionRequest(msg OrderMassActionRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2UserNotification is a Callback for UserNotification messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2UserNotification(msg UserNotification, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2StreamAssignmentRequest is a Callback for StreamAssignmentRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2StreamAssignmentRequest(msg StreamAssignmentRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2StreamAssignmentReport is a Callback for StreamAssignmentReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2StreamAssignmentReport(msg StreamAssignmentReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2StreamAssignmentReportACK is a Callback for StreamAssignmentReportACK messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2StreamAssignmentReportACK(msg StreamAssignmentReportACK, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2PartyDetailsListRequest is a Callback for PartyDetailsListRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2PartyDetailsListRequest(msg PartyDetailsListRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2PartyDetailsListReport is a Callback for PartyDetailsListReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2PartyDetailsListReport(msg PartyDetailsListReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2NewOrderSingle is a Callback for NewOrderSingle messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2NewOrderList is a Callback for NewOrderList messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2OrderCancelRequest is a Callback for OrderCancelRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2OrderCancelReplaceRequest is a Callback for OrderCancelReplaceRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2OrderStatusRequest is a Callback for OrderStatusRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2AllocationInstruction is a Callback for AllocationInstruction messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationInstruction(msg AllocationInstruction, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ListCancelRequest is a Callback for ListCancelRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ListExecute is a Callback for ListExecute messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListExecute(msg ListExecute, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ListStatusRequest is a Callback for ListStatusRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ListStatus is a Callback for ListStatus messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListStatus(msg ListStatus, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2AllocationInstructionAck is a Callback for AllocationInstructionAck messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationInstructionAck(msg AllocationInstructionAck, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2DontKnowTrade is a Callback for DontKnowTrade messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2QuoteRequest is a Callback for QuoteRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2Quote is a Callback for Quote messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2Quote(msg Quote, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SettlementInstructions is a Callback for SettlementInstructions messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SettlementInstructions(msg SettlementInstructions, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2MarketDataRequest is a Callback for MarketDataRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDataRequest(msg MarketDataRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2MarketDataSnapshotFullRefresh is a Callback for MarketDataSnapshotFullRefresh messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2MarketDataIncrementalRefresh is a Callback for MarketDataIncrementalRefresh messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2MarketDataRequestReject is a Callback for MarketDataRequestReject messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2QuoteCancel is a Callback for QuoteCancel messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteCancel(msg QuoteCancel, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2QuoteStatusRequest is a Callback for QuoteStatusRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2MassQuoteAcknowledgement is a Callback for MassQuoteAcknowledgement messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SecurityDefinitionRequest is a Callback for SecurityDefinitionRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SecurityDefinition is a Callback for SecurityDefinition messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityDefinition(msg SecurityDefinition, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SecurityStatusRequest is a Callback for SecurityStatusRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SecurityStatus is a Callback for SecurityStatus messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityStatus(msg SecurityStatus, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2TradingSessionStatusRequest is a Callback for TradingSessionStatusRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2TradingSessionStatus is a Callback for TradingSessionStatus messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionStatus(msg TradingSessionStatus, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2MassQuote is a Callback for MassQuote messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2MassQuote(msg MassQuote, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2BusinessMessageReject is a Callback for BusinessMessageReject messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2BusinessMessageReject(msg BusinessMessageReject, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2BidRequest is a Callback for BidRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2BidRequest(msg BidRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2BidResponse is a Callback for BidResponse messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2BidResponse(msg BidResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2ListStrikePrice is a Callback for ListStrikePrice messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListStrikePrice(msg ListStrikePrice, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2RegistrationInstructions is a Callback for RegistrationInstructions messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2RegistrationInstructions(msg RegistrationInstructions, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2RegistrationInstructionsResponse is a Callback for RegistrationInstructionsResponse messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2OrderMassCancelRequest is a Callback for OrderMassCancelRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2OrderMassCancelReport is a Callback for OrderMassCancelReport messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2NewOrderCross is a Callback for NewOrderCross messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2NewOrderCross(msg NewOrderCross, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2CrossOrderCancelReplaceRequest is a Callback for CrossOrderCancelReplaceRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2CrossOrderCancelRequest is a Callback for CrossOrderCancelRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SecurityTypeRequest is a Callback for SecurityTypeRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SecurityTypes is a Callback for SecurityTypes messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityTypes(msg SecurityTypes, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SecurityListRequest is a Callback for SecurityListRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityListRequest(msg SecurityListRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2SecurityList is a Callback for SecurityList messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityList(msg SecurityList, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}

//OnFIX50SP2DerivativeSecurityListRequest is a Callback for DerivativeSecurityListRequest messages.
func (c *FIX50SP2MessageCracker) OnFIX50SP2DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfix.SessionID) errors.MessageRejectError {
	return errors.UnsupportedMessageType()
}
