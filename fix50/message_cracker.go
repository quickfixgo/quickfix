package fix50

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

func Crack(msg message.Message, sessionID quickfix.SessionID, router MessageRouter) message.MessageReject {

	msgType := new(field.MsgType)
	switch msg.Header.Get(msgType); msgType.Value {
	case "C":
		return router.OnFIX50Email(Email{msg}, sessionID)
	case "o":
		return router.OnFIX50RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "AA":
		return router.OnFIX50DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "AK":
		return router.OnFIX50Confirmation(Confirmation{msg}, sessionID)
	case "AM":
		return router.OnFIX50PositionMaintenanceReport(PositionMaintenanceReport{msg}, sessionID)
	case "D":
		return router.OnFIX50NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "k":
		return router.OnFIX50BidRequest(BidRequest{msg}, sessionID)
	case "m":
		return router.OnFIX50ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "t":
		return router.OnFIX50CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "AX":
		return router.OnFIX50CollateralRequest(CollateralRequest{msg}, sessionID)
	case "BI":
		return router.OnFIX50TradingSessionListRequest(TradingSessionListRequest{msg}, sessionID)
	case "F":
		return router.OnFIX50OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "a":
		return router.OnFIX50QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "b":
		return router.OnFIX50MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "h":
		return router.OnFIX50TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "j":
		return router.OnFIX50BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "r":
		return router.OnFIX50OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "y":
		return router.OnFIX50SecurityList(SecurityList{msg}, sessionID)
	case "AV":
		return router.OnFIX50SettlementInstructionRequest(SettlementInstructionRequest{msg}, sessionID)
	case "BA":
		return router.OnFIX50CollateralReport(CollateralReport{msg}, sessionID)
	case "BB":
		return router.OnFIX50CollateralInquiry(CollateralInquiry{msg}, sessionID)
	case "BM":
		return router.OnFIX50AllocationInstructionAlert(AllocationInstructionAlert{msg}, sessionID)
	case "S":
		return router.OnFIX50Quote(Quote{msg}, sessionID)
	case "c":
		return router.OnFIX50SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "g":
		return router.OnFIX50TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "i":
		return router.OnFIX50MassQuote(MassQuote{msg}, sessionID)
	case "AN":
		return router.OnFIX50RequestForPositions(RequestForPositions{msg}, sessionID)
	case "G":
		return router.OnFIX50OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "P":
		return router.OnFIX50AllocationInstructionAck(AllocationInstructionAck{msg}, sessionID)
	case "Q":
		return router.OnFIX50DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "T":
		return router.OnFIX50SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "BO":
		return router.OnFIX50ContraryIntentionReport(ContraryIntentionReport{msg}, sessionID)
	case "J":
		return router.OnFIX50AllocationInstruction(AllocationInstruction{msg}, sessionID)
	case "w":
		return router.OnFIX50SecurityTypes(SecurityTypes{msg}, sessionID)
	case "AB":
		return router.OnFIX50NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "BH":
		return router.OnFIX50ConfirmationRequest(ConfirmationRequest{msg}, sessionID)
	case "8":
		return router.OnFIX50ExecutionReport(ExecutionReport{msg}, sessionID)
	case "W":
		return router.OnFIX50MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "Z":
		return router.OnFIX50QuoteCancel(QuoteCancel{msg}, sessionID)
	case "v":
		return router.OnFIX50SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "x":
		return router.OnFIX50SecurityListRequest(SecurityListRequest{msg}, sessionID)
	case "z":
		return router.OnFIX50DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	case "AJ":
		return router.OnFIX50QuoteResponse(QuoteResponse{msg}, sessionID)
	case "AL":
		return router.OnFIX50PositionMaintenanceRequest(PositionMaintenanceRequest{msg}, sessionID)
	case "AW":
		return router.OnFIX50AssignmentReport(AssignmentReport{msg}, sessionID)
	case "BC":
		return router.OnFIX50NetworkCounterpartySystemStatusRequest(NetworkCounterpartySystemStatusRequest{msg}, sessionID)
	case "BN":
		return router.OnFIX50ExecutionAcknowledgement(ExecutionAcknowledgement{msg}, sessionID)
	case "6":
		return router.OnFIX50IOI(IOI{msg}, sessionID)
	case "E":
		return router.OnFIX50NewOrderList(NewOrderList{msg}, sessionID)
	case "K":
		return router.OnFIX50ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "AR":
		return router.OnFIX50TradeCaptureReportAck(TradeCaptureReportAck{msg}, sessionID)
	case "AS":
		return router.OnFIX50AllocationReport(AllocationReport{msg}, sessionID)
	case "AG":
		return router.OnFIX50QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "AO":
		return router.OnFIX50RequestForPositionsAck(RequestForPositionsAck{msg}, sessionID)
	case "BG":
		return router.OnFIX50CollateralInquiryAck(CollateralInquiryAck{msg}, sessionID)
	case "BP":
		return router.OnFIX50SecurityDefinitionUpdateReport(SecurityDefinitionUpdateReport{msg}, sessionID)
	case "BK":
		return router.OnFIX50SecurityListUpdateReport(SecurityListUpdateReport{msg}, sessionID)
	case "H":
		return router.OnFIX50OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "L":
		return router.OnFIX50ListExecute(ListExecute{msg}, sessionID)
	case "d":
		return router.OnFIX50SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "e":
		return router.OnFIX50SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "s":
		return router.OnFIX50NewOrderCross(NewOrderCross{msg}, sessionID)
	case "AF":
		return router.OnFIX50OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "AU":
		return router.OnFIX50ConfirmationAck(ConfirmationAck{msg}, sessionID)
	case "p":
		return router.OnFIX50RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "AI":
		return router.OnFIX50QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	case "7":
		return router.OnFIX50Advertisement(Advertisement{msg}, sessionID)
	case "M":
		return router.OnFIX50ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "AQ":
		return router.OnFIX50TradeCaptureReportRequestAck(TradeCaptureReportRequestAck{msg}, sessionID)
	case "BD":
		return router.OnFIX50NetworkCounterpartySystemStatusResponse(NetworkCounterpartySystemStatusResponse{msg}, sessionID)
	case "N":
		return router.OnFIX50ListStatus(ListStatus{msg}, sessionID)
	case "V":
		return router.OnFIX50MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "l":
		return router.OnFIX50BidResponse(BidResponse{msg}, sessionID)
	case "u":
		return router.OnFIX50CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "AC":
		return router.OnFIX50MultilegOrderCancelReplace(MultilegOrderCancelReplace{msg}, sessionID)
	case "AE":
		return router.OnFIX50TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "AZ":
		return router.OnFIX50CollateralResponse(CollateralResponse{msg}, sessionID)
	case "BE":
		return router.OnFIX50UserRequest(UserRequest{msg}, sessionID)
	case "9":
		return router.OnFIX50OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "f":
		return router.OnFIX50SecurityStatus(SecurityStatus{msg}, sessionID)
	case "AD":
		return router.OnFIX50TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "AP":
		return router.OnFIX50PositionReport(PositionReport{msg}, sessionID)
	case "BF":
		return router.OnFIX50UserResponse(UserResponse{msg}, sessionID)
	case "BJ":
		return router.OnFIX50TradingSessionList(TradingSessionList{msg}, sessionID)
	case "B":
		return router.OnFIX50News(News{msg}, sessionID)
	case "X":
		return router.OnFIX50MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "Y":
		return router.OnFIX50MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "q":
		return router.OnFIX50OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "AT":
		return router.OnFIX50AllocationReportAck(AllocationReportAck{msg}, sessionID)
	case "AY":
		return router.OnFIX50CollateralAssignment(CollateralAssignment{msg}, sessionID)
	case "BL":
		return router.OnFIX50AdjustedPositionReport(AdjustedPositionReport{msg}, sessionID)
	case "R":
		return router.OnFIX50QuoteRequest(QuoteRequest{msg}, sessionID)
	case "AH":
		return router.OnFIX50RFQRequest(RFQRequest{msg}, sessionID)
	}
	return message.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX50RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50QuoteStatusReport(msg QuoteStatusReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50Advertisement(msg Advertisement, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50ListStatus(msg ListStatus, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50MarketDataRequest(msg MarketDataRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50BidResponse(msg BidResponse, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50TradeCaptureReport(msg TradeCaptureReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50CollateralResponse(msg CollateralResponse, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50UserRequest(msg UserRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50SecurityStatus(msg SecurityStatus, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50PositionReport(msg PositionReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50UserResponse(msg UserResponse, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50TradingSessionList(msg TradingSessionList, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50News(msg News, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50AllocationReportAck(msg AllocationReportAck, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50CollateralAssignment(msg CollateralAssignment, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50AdjustedPositionReport(msg AdjustedPositionReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50RFQRequest(msg RFQRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50Email(msg Email, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50RegistrationInstructions(msg RegistrationInstructions, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50Confirmation(msg Confirmation, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50BidRequest(msg BidRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50ListStrikePrice(msg ListStrikePrice, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50CollateralRequest(msg CollateralRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50TradingSessionListRequest(msg TradingSessionListRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50TradingSessionStatus(msg TradingSessionStatus, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50BusinessMessageReject(msg BusinessMessageReject, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50SecurityList(msg SecurityList, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50CollateralReport(msg CollateralReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50CollateralInquiry(msg CollateralInquiry, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50AllocationInstructionAlert(msg AllocationInstructionAlert, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50Quote(msg Quote, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50MassQuote(msg MassQuote, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50RequestForPositions(msg RequestForPositions, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50AllocationInstructionAck(msg AllocationInstructionAck, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50SettlementInstructions(msg SettlementInstructions, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50ContraryIntentionReport(msg ContraryIntentionReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50AllocationInstruction(msg AllocationInstruction, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50SecurityTypes(msg SecurityTypes, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50NewOrderMultileg(msg NewOrderMultileg, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50ConfirmationRequest(msg ConfirmationRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50QuoteCancel(msg QuoteCancel, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50SecurityListRequest(msg SecurityListRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50QuoteResponse(msg QuoteResponse, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50AssignmentReport(msg AssignmentReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50IOI(msg IOI, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50AllocationReport(msg AllocationReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50QuoteRequestReject(msg QuoteRequestReject, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50RequestForPositionsAck(msg RequestForPositionsAck, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50CollateralInquiryAck(msg CollateralInquiryAck, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50SecurityListUpdateReport(msg SecurityListUpdateReport, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50ListExecute(msg ListExecute, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50SecurityDefinition(msg SecurityDefinition, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50NewOrderCross(msg NewOrderCross, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfix.SessionID) message.MessageReject
	OnFIX50ConfirmationAck(msg ConfirmationAck, sessionID quickfix.SessionID) message.MessageReject
}
type FIX50MessageCracker struct{}

func (c *FIX50MessageCracker) OnFIX50OrderCancelRequest(msg OrderCancelRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50QuoteStatusRequest(msg QuoteStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradingSessionStatus(msg TradingSessionStatus, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50BusinessMessageReject(msg BusinessMessageReject, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderMassCancelReport(msg OrderMassCancelReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityList(msg SecurityList, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SettlementInstructionRequest(msg SettlementInstructionRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CollateralReport(msg CollateralReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CollateralInquiry(msg CollateralInquiry, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AllocationInstructionAlert(msg AllocationInstructionAlert, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50Quote(msg Quote, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MassQuote(msg MassQuote, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50RequestForPositions(msg RequestForPositions, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AllocationInstructionAck(msg AllocationInstructionAck, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50DontKnowTrade(msg DontKnowTrade, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SettlementInstructions(msg SettlementInstructions, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ContraryIntentionReport(msg ContraryIntentionReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AllocationInstruction(msg AllocationInstruction, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityTypes(msg SecurityTypes, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50NewOrderMultileg(msg NewOrderMultileg, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ConfirmationRequest(msg ConfirmationRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ExecutionReport(msg ExecutionReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50QuoteCancel(msg QuoteCancel, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityTypeRequest(msg SecurityTypeRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityListRequest(msg SecurityListRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50QuoteResponse(msg QuoteResponse, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AssignmentReport(msg AssignmentReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50IOI(msg IOI, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50NewOrderList(msg NewOrderList, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ListCancelRequest(msg ListCancelRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradeCaptureReportAck(msg TradeCaptureReportAck, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AllocationReport(msg AllocationReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50QuoteRequestReject(msg QuoteRequestReject, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50RequestForPositionsAck(msg RequestForPositionsAck, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CollateralInquiryAck(msg CollateralInquiryAck, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityListUpdateReport(msg SecurityListUpdateReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderStatusRequest(msg OrderStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ListExecute(msg ListExecute, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityDefinition(msg SecurityDefinition, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityStatusRequest(msg SecurityStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50NewOrderCross(msg NewOrderCross, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ConfirmationAck(msg ConfirmationAck, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50QuoteStatusReport(msg QuoteStatusReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50Advertisement(msg Advertisement, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ListStatusRequest(msg ListStatusRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ListStatus(msg ListStatus, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MarketDataRequest(msg MarketDataRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50BidResponse(msg BidResponse, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradeCaptureReport(msg TradeCaptureReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CollateralResponse(msg CollateralResponse, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50UserRequest(msg UserRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderCancelReject(msg OrderCancelReject, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityStatus(msg SecurityStatus, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50PositionReport(msg PositionReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50UserResponse(msg UserResponse, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradingSessionList(msg TradingSessionList, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50News(msg News, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MarketDataRequestReject(msg MarketDataRequestReject, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AllocationReportAck(msg AllocationReportAck, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CollateralAssignment(msg CollateralAssignment, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AdjustedPositionReport(msg AdjustedPositionReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50QuoteRequest(msg QuoteRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50RFQRequest(msg RFQRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50Email(msg Email, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50RegistrationInstructions(msg RegistrationInstructions, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50DerivativeSecurityList(msg DerivativeSecurityList, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50Confirmation(msg Confirmation, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50PositionMaintenanceReport(msg PositionMaintenanceReport, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50NewOrderSingle(msg NewOrderSingle, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50BidRequest(msg BidRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ListStrikePrice(msg ListStrikePrice, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CollateralRequest(msg CollateralRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradingSessionListRequest(msg TradingSessionListRequest, sessionId quickfix.SessionID) message.MessageReject {
	return message.NewUnsupportedMessageType(msg.Message)
}
