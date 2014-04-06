package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

func Crack(msg quickfix.Message, sessionID quickfix.SessionID, router MessageRouter) quickfix.MessageReject {

	msgType := new(field.MsgType)
	switch msg.Header.Get(msgType); msgType.Value {
	case "N":
		return router.OnFIX50SP2ListStatus(ListStatus{msg}, sessionID)
	case "Y":
		return router.OnFIX50SP2MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "a":
		return router.OnFIX50SP2QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "AJ":
		return router.OnFIX50SP2QuoteResponse(QuoteResponse{msg}, sessionID)
	case "BC":
		return router.OnFIX50SP2NetworkCounterpartySystemStatusRequest(NetworkCounterpartySystemStatusRequest{msg}, sessionID)
	case "BJ":
		return router.OnFIX50SP2TradingSessionList(TradingSessionList{msg}, sessionID)
	case "i":
		return router.OnFIX50SP2MassQuote(MassQuote{msg}, sessionID)
	case "AI":
		return router.OnFIX50SP2QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	case "BP":
		return router.OnFIX50SP2SecurityDefinitionUpdateReport(SecurityDefinitionUpdateReport{msg}, sessionID)
	case "M":
		return router.OnFIX50SP2ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "c":
		return router.OnFIX50SP2SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "m":
		return router.OnFIX50SP2ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "w":
		return router.OnFIX50SP2SecurityTypes(SecurityTypes{msg}, sessionID)
	case "z":
		return router.OnFIX50SP2DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	case "AF":
		return router.OnFIX50SP2OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "BO":
		return router.OnFIX50SP2ContraryIntentionReport(ContraryIntentionReport{msg}, sessionID)
	case "BL":
		return router.OnFIX50SP2AdjustedPositionReport(AdjustedPositionReport{msg}, sessionID)
	case "y":
		return router.OnFIX50SP2SecurityList(SecurityList{msg}, sessionID)
	case "R":
		return router.OnFIX50SP2QuoteRequest(QuoteRequest{msg}, sessionID)
	case "CD":
		return router.OnFIX50SP2StreamAssignmentReport(StreamAssignmentReport{msg}, sessionID)
	case "BH":
		return router.OnFIX50SP2ConfirmationRequest(ConfirmationRequest{msg}, sessionID)
	case "BI":
		return router.OnFIX50SP2TradingSessionListRequest(TradingSessionListRequest{msg}, sessionID)
	case "D":
		return router.OnFIX50SP2NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "s":
		return router.OnFIX50SP2NewOrderCross(NewOrderCross{msg}, sessionID)
	case "AL":
		return router.OnFIX50SP2PositionMaintenanceRequest(PositionMaintenanceRequest{msg}, sessionID)
	case "BZ":
		return router.OnFIX50SP2OrderMassActionReport(OrderMassActionReport{msg}, sessionID)
	case "9":
		return router.OnFIX50SP2OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "BG":
		return router.OnFIX50SP2CollateralInquiryAck(CollateralInquiryAck{msg}, sessionID)
	case "AN":
		return router.OnFIX50SP2RequestForPositions(RequestForPositions{msg}, sessionID)
	case "AO":
		return router.OnFIX50SP2RequestForPositionsAck(RequestForPositionsAck{msg}, sessionID)
	case "AQ":
		return router.OnFIX50SP2TradeCaptureReportRequestAck(TradeCaptureReportRequestAck{msg}, sessionID)
	case "AT":
		return router.OnFIX50SP2AllocationReportAck(AllocationReportAck{msg}, sessionID)
	case "CA":
		return router.OnFIX50SP2OrderMassActionRequest(OrderMassActionRequest{msg}, sessionID)
	case "g":
		return router.OnFIX50SP2TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "q":
		return router.OnFIX50SP2OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "AB":
		return router.OnFIX50SP2NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "H":
		return router.OnFIX50SP2OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "X":
		return router.OnFIX50SP2MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "B":
		return router.OnFIX50SP2News(News{msg}, sessionID)
	case "L":
		return router.OnFIX50SP2ListExecute(ListExecute{msg}, sessionID)
	case "o":
		return router.OnFIX50SP2RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "u":
		return router.OnFIX50SP2CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "AK":
		return router.OnFIX50SP2Confirmation(Confirmation{msg}, sessionID)
	case "AP":
		return router.OnFIX50SP2PositionReport(PositionReport{msg}, sessionID)
	case "CF":
		return router.OnFIX50SP2PartyDetailsListRequest(PartyDetailsListRequest{msg}, sessionID)
	case "AA":
		return router.OnFIX50SP2DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "AY":
		return router.OnFIX50SP2CollateralAssignment(CollateralAssignment{msg}, sessionID)
	case "BB":
		return router.OnFIX50SP2CollateralInquiry(CollateralInquiry{msg}, sessionID)
	case "BR":
		return router.OnFIX50SP2DerivativeSecurityListUpdateReport(DerivativeSecurityListUpdateReport{msg}, sessionID)
	case "BS":
		return router.OnFIX50SP2TradingSessionListUpdateReport(TradingSessionListUpdateReport{msg}, sessionID)
	case "BW":
		return router.OnFIX50SP2ApplicationMessageRequest(ApplicationMessageRequest{msg}, sessionID)
	case "8":
		return router.OnFIX50SP2ExecutionReport(ExecutionReport{msg}, sessionID)
	case "p":
		return router.OnFIX50SP2RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "AZ":
		return router.OnFIX50SP2CollateralResponse(CollateralResponse{msg}, sessionID)
	case "BK":
		return router.OnFIX50SP2SecurityListUpdateReport(SecurityListUpdateReport{msg}, sessionID)
	case "f":
		return router.OnFIX50SP2SecurityStatus(SecurityStatus{msg}, sessionID)
	case "h":
		return router.OnFIX50SP2TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "AH":
		return router.OnFIX50SP2RFQRequest(RFQRequest{msg}, sessionID)
	case "BV":
		return router.OnFIX50SP2MarketDefinitionUpdateReport(MarketDefinitionUpdateReport{msg}, sessionID)
	case "BY":
		return router.OnFIX50SP2ApplicationMessageReport(ApplicationMessageReport{msg}, sessionID)
	case "6":
		return router.OnFIX50SP2IOI(IOI{msg}, sessionID)
	case "BM":
		return router.OnFIX50SP2AllocationInstructionAlert(AllocationInstructionAlert{msg}, sessionID)
	case "CB":
		return router.OnFIX50SP2UserNotification(UserNotification{msg}, sessionID)
	case "F":
		return router.OnFIX50SP2OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "Z":
		return router.OnFIX50SP2QuoteCancel(QuoteCancel{msg}, sessionID)
	case "AG":
		return router.OnFIX50SP2QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "AW":
		return router.OnFIX50SP2AssignmentReport(AssignmentReport{msg}, sessionID)
	case "G":
		return router.OnFIX50SP2OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "d":
		return router.OnFIX50SP2SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "BE":
		return router.OnFIX50SP2UserRequest(UserRequest{msg}, sessionID)
	case "7":
		return router.OnFIX50SP2Advertisement(Advertisement{msg}, sessionID)
	case "e":
		return router.OnFIX50SP2SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "r":
		return router.OnFIX50SP2OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "BX":
		return router.OnFIX50SP2ApplicationMessageRequestAck(ApplicationMessageRequestAck{msg}, sessionID)
	case "j":
		return router.OnFIX50SP2BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "BD":
		return router.OnFIX50SP2NetworkCounterpartySystemStatusResponse(NetworkCounterpartySystemStatusResponse{msg}, sessionID)
	case "BN":
		return router.OnFIX50SP2ExecutionAcknowledgement(ExecutionAcknowledgement{msg}, sessionID)
	case "AS":
		return router.OnFIX50SP2AllocationReport(AllocationReport{msg}, sessionID)
	case "BF":
		return router.OnFIX50SP2UserResponse(UserResponse{msg}, sessionID)
	case "C":
		return router.OnFIX50SP2Email(Email{msg}, sessionID)
	case "b":
		return router.OnFIX50SP2MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "v":
		return router.OnFIX50SP2SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "AU":
		return router.OnFIX50SP2ConfirmationAck(ConfirmationAck{msg}, sessionID)
	case "AX":
		return router.OnFIX50SP2CollateralRequest(CollateralRequest{msg}, sessionID)
	case "AD":
		return router.OnFIX50SP2TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "CC":
		return router.OnFIX50SP2StreamAssignmentRequest(StreamAssignmentRequest{msg}, sessionID)
	case "W":
		return router.OnFIX50SP2MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "AE":
		return router.OnFIX50SP2TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "AM":
		return router.OnFIX50SP2PositionMaintenanceReport(PositionMaintenanceReport{msg}, sessionID)
	case "BU":
		return router.OnFIX50SP2MarketDefinition(MarketDefinition{msg}, sessionID)
	case "E":
		return router.OnFIX50SP2NewOrderList(NewOrderList{msg}, sessionID)
	case "k":
		return router.OnFIX50SP2BidRequest(BidRequest{msg}, sessionID)
	case "AC":
		return router.OnFIX50SP2MultilegOrderCancelReplace(MultilegOrderCancelReplace{msg}, sessionID)
	case "K":
		return router.OnFIX50SP2ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "t":
		return router.OnFIX50SP2CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "x":
		return router.OnFIX50SP2SecurityListRequest(SecurityListRequest{msg}, sessionID)
	case "BA":
		return router.OnFIX50SP2CollateralReport(CollateralReport{msg}, sessionID)
	case "S":
		return router.OnFIX50SP2Quote(Quote{msg}, sessionID)
	case "T":
		return router.OnFIX50SP2SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "V":
		return router.OnFIX50SP2MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "l":
		return router.OnFIX50SP2BidResponse(BidResponse{msg}, sessionID)
	case "AR":
		return router.OnFIX50SP2TradeCaptureReportAck(TradeCaptureReportAck{msg}, sessionID)
	case "BT":
		return router.OnFIX50SP2MarketDefinitionRequest(MarketDefinitionRequest{msg}, sessionID)
	case "CG":
		return router.OnFIX50SP2PartyDetailsListReport(PartyDetailsListReport{msg}, sessionID)
	case "P":
		return router.OnFIX50SP2AllocationInstructionAck(AllocationInstructionAck{msg}, sessionID)
	case "Q":
		return router.OnFIX50SP2DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "AV":
		return router.OnFIX50SP2SettlementInstructionRequest(SettlementInstructionRequest{msg}, sessionID)
	case "CE":
		return router.OnFIX50SP2StreamAssignmentReportACK(StreamAssignmentReportACK{msg}, sessionID)
	case "J":
		return router.OnFIX50SP2AllocationInstruction(AllocationInstruction{msg}, sessionID)
	case "BQ":
		return router.OnFIX50SP2SettlementObligationReport(SettlementObligationReport{msg}, sessionID)
	}
	return quickfix.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX50SP2TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2StreamAssignmentRequest(msg StreamAssignmentRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2TradeCaptureReport(msg TradeCaptureReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2MarketDefinition(msg MarketDefinition, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2BidRequest(msg BidRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SecurityListRequest(msg SecurityListRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2CollateralReport(msg CollateralReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2Quote(msg Quote, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SettlementInstructions(msg SettlementInstructions, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2MarketDataRequest(msg MarketDataRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2BidResponse(msg BidResponse, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2MarketDefinitionRequest(msg MarketDefinitionRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2PartyDetailsListReport(msg PartyDetailsListReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2AllocationInstructionAck(msg AllocationInstructionAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2StreamAssignmentReportACK(msg StreamAssignmentReportACK, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2AllocationInstruction(msg AllocationInstruction, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SettlementObligationReport(msg SettlementObligationReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ListStatus(msg ListStatus, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2QuoteResponse(msg QuoteResponse, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2TradingSessionList(msg TradingSessionList, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2MassQuote(msg MassQuote, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2QuoteStatusReport(msg QuoteStatusReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ListStrikePrice(msg ListStrikePrice, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SecurityTypes(msg SecurityTypes, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ContraryIntentionReport(msg ContraryIntentionReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2AdjustedPositionReport(msg AdjustedPositionReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SecurityList(msg SecurityList, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2StreamAssignmentReport(msg StreamAssignmentReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ConfirmationRequest(msg ConfirmationRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2TradingSessionListRequest(msg TradingSessionListRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2NewOrderCross(msg NewOrderCross, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2OrderMassActionReport(msg OrderMassActionReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2CollateralInquiryAck(msg CollateralInquiryAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2RequestForPositions(msg RequestForPositions, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2RequestForPositionsAck(msg RequestForPositionsAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2AllocationReportAck(msg AllocationReportAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2OrderMassActionRequest(msg OrderMassActionRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2NewOrderMultileg(msg NewOrderMultileg, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2News(msg News, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ListExecute(msg ListExecute, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2RegistrationInstructions(msg RegistrationInstructions, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2Confirmation(msg Confirmation, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2PositionReport(msg PositionReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2PartyDetailsListRequest(msg PartyDetailsListRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2CollateralAssignment(msg CollateralAssignment, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2CollateralInquiry(msg CollateralInquiry, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ApplicationMessageRequest(msg ApplicationMessageRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2CollateralResponse(msg CollateralResponse, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SecurityListUpdateReport(msg SecurityListUpdateReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SecurityStatus(msg SecurityStatus, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2TradingSessionStatus(msg TradingSessionStatus, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2RFQRequest(msg RFQRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ApplicationMessageReport(msg ApplicationMessageReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2IOI(msg IOI, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2AllocationInstructionAlert(msg AllocationInstructionAlert, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2UserNotification(msg UserNotification, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2QuoteCancel(msg QuoteCancel, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2QuoteRequestReject(msg QuoteRequestReject, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2AssignmentReport(msg AssignmentReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SecurityDefinition(msg SecurityDefinition, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2UserRequest(msg UserRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2Advertisement(msg Advertisement, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2BusinessMessageReject(msg BusinessMessageReject, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2AllocationReport(msg AllocationReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2UserResponse(msg UserResponse, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2Email(msg Email, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2ConfirmationAck(msg ConfirmationAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX50SP2CollateralRequest(msg CollateralRequest, sessionID quickfix.SessionID) quickfix.MessageReject
}
type FIX50SP2MessageCracker struct{}

func (c *FIX50SP2MessageCracker) OnFIX50SP2ListStatus(msg ListStatus, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDataRequestReject(msg MarketDataRequestReject, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteStatusRequest(msg QuoteStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteResponse(msg QuoteResponse, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionList(msg TradingSessionList, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MassQuote(msg MassQuote, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteStatusReport(msg QuoteStatusReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListStatusRequest(msg ListStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListStrikePrice(msg ListStrikePrice, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityTypes(msg SecurityTypes, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ContraryIntentionReport(msg ContraryIntentionReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AdjustedPositionReport(msg AdjustedPositionReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityList(msg SecurityList, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteRequest(msg QuoteRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2StreamAssignmentReport(msg StreamAssignmentReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ConfirmationRequest(msg ConfirmationRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionListRequest(msg TradingSessionListRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2NewOrderSingle(msg NewOrderSingle, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2NewOrderCross(msg NewOrderCross, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassActionReport(msg OrderMassActionReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderCancelReject(msg OrderCancelReject, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralInquiryAck(msg CollateralInquiryAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2RequestForPositions(msg RequestForPositions, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2RequestForPositionsAck(msg RequestForPositionsAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationReportAck(msg AllocationReportAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassActionRequest(msg OrderMassActionRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2NewOrderMultileg(msg NewOrderMultileg, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderStatusRequest(msg OrderStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2News(msg News, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListExecute(msg ListExecute, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2RegistrationInstructions(msg RegistrationInstructions, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2Confirmation(msg Confirmation, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2PositionReport(msg PositionReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2PartyDetailsListRequest(msg PartyDetailsListRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2DerivativeSecurityList(msg DerivativeSecurityList, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralAssignment(msg CollateralAssignment, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralInquiry(msg CollateralInquiry, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2DerivativeSecurityListUpdateReport(msg DerivativeSecurityListUpdateReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionListUpdateReport(msg TradingSessionListUpdateReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ApplicationMessageRequest(msg ApplicationMessageRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ExecutionReport(msg ExecutionReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralResponse(msg CollateralResponse, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityListUpdateReport(msg SecurityListUpdateReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityStatus(msg SecurityStatus, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradingSessionStatus(msg TradingSessionStatus, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2RFQRequest(msg RFQRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDefinitionUpdateReport(msg MarketDefinitionUpdateReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ApplicationMessageReport(msg ApplicationMessageReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2IOI(msg IOI, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationInstructionAlert(msg AllocationInstructionAlert, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2UserNotification(msg UserNotification, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderCancelRequest(msg OrderCancelRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteCancel(msg QuoteCancel, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2QuoteRequestReject(msg QuoteRequestReject, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AssignmentReport(msg AssignmentReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityDefinition(msg SecurityDefinition, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2UserRequest(msg UserRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2Advertisement(msg Advertisement, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityStatusRequest(msg SecurityStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2OrderMassCancelReport(msg OrderMassCancelReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ApplicationMessageRequestAck(msg ApplicationMessageRequestAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2BusinessMessageReject(msg BusinessMessageReject, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationReport(msg AllocationReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2UserResponse(msg UserResponse, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2Email(msg Email, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityTypeRequest(msg SecurityTypeRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ConfirmationAck(msg ConfirmationAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralRequest(msg CollateralRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2StreamAssignmentRequest(msg StreamAssignmentRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradeCaptureReport(msg TradeCaptureReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2PositionMaintenanceReport(msg PositionMaintenanceReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDefinition(msg MarketDefinition, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2NewOrderList(msg NewOrderList, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2BidRequest(msg BidRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2ListCancelRequest(msg ListCancelRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SecurityListRequest(msg SecurityListRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2CollateralReport(msg CollateralReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2Quote(msg Quote, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SettlementInstructions(msg SettlementInstructions, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDataRequest(msg MarketDataRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2BidResponse(msg BidResponse, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2TradeCaptureReportAck(msg TradeCaptureReportAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2MarketDefinitionRequest(msg MarketDefinitionRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2PartyDetailsListReport(msg PartyDetailsListReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationInstructionAck(msg AllocationInstructionAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2DontKnowTrade(msg DontKnowTrade, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SettlementInstructionRequest(msg SettlementInstructionRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2StreamAssignmentReportACK(msg StreamAssignmentReportACK, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2AllocationInstruction(msg AllocationInstruction, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50SP2MessageCracker) OnFIX50SP2SettlementObligationReport(msg SettlementObligationReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
