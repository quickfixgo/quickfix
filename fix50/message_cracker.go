package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/tag"
)

func Crack(msg message.Message, sessionID quickfixgo.SessionID, router MessageRouter) reject.MessageReject {
	switch msgType, _ := msg.Header.StringValue(tag.MsgType); msgType {
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
	OnFIX50IOI(msg IOI, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50Advertisement(msg Advertisement, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50ExecutionReport(msg ExecutionReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50OrderCancelReject(msg OrderCancelReject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50News(msg News, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50Email(msg Email, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50NewOrderSingle(msg NewOrderSingle, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50NewOrderList(msg NewOrderList, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50OrderCancelRequest(msg OrderCancelRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50OrderStatusRequest(msg OrderStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50AllocationInstruction(msg AllocationInstruction, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50ListCancelRequest(msg ListCancelRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50ListExecute(msg ListExecute, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50ListStatusRequest(msg ListStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50ListStatus(msg ListStatus, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50AllocationInstructionAck(msg AllocationInstructionAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50DontKnowTrade(msg DontKnowTrade, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50QuoteRequest(msg QuoteRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50Quote(msg Quote, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50SettlementInstructions(msg SettlementInstructions, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50MarketDataRequest(msg MarketDataRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50QuoteCancel(msg QuoteCancel, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50SecurityDefinition(msg SecurityDefinition, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50SecurityStatus(msg SecurityStatus, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50TradingSessionStatus(msg TradingSessionStatus, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50MassQuote(msg MassQuote, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50BusinessMessageReject(msg BusinessMessageReject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50BidRequest(msg BidRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50BidResponse(msg BidResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50ListStrikePrice(msg ListStrikePrice, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50RegistrationInstructions(msg RegistrationInstructions, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50NewOrderCross(msg NewOrderCross, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50SecurityTypes(msg SecurityTypes, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50SecurityListRequest(msg SecurityListRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50SecurityList(msg SecurityList, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50NewOrderMultileg(msg NewOrderMultileg, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50TradeCaptureReport(msg TradeCaptureReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50QuoteRequestReject(msg QuoteRequestReject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50RFQRequest(msg RFQRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50QuoteStatusReport(msg QuoteStatusReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50QuoteResponse(msg QuoteResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50Confirmation(msg Confirmation, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50RequestForPositions(msg RequestForPositions, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50RequestForPositionsAck(msg RequestForPositionsAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50PositionReport(msg PositionReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50AllocationReport(msg AllocationReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50AllocationReportAck(msg AllocationReportAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50ConfirmationAck(msg ConfirmationAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50AssignmentReport(msg AssignmentReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50CollateralRequest(msg CollateralRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50CollateralAssignment(msg CollateralAssignment, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50CollateralResponse(msg CollateralResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50CollateralReport(msg CollateralReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50CollateralInquiry(msg CollateralInquiry, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50UserRequest(msg UserRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50UserResponse(msg UserResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50CollateralInquiryAck(msg CollateralInquiryAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50ConfirmationRequest(msg ConfirmationRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50ContraryIntentionReport(msg ContraryIntentionReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50SecurityListUpdateReport(msg SecurityListUpdateReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50AdjustedPositionReport(msg AdjustedPositionReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50AllocationInstructionAlert(msg AllocationInstructionAlert, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50TradingSessionList(msg TradingSessionList, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX50TradingSessionListRequest(msg TradingSessionListRequest, sessionID quickfixgo.SessionID) reject.MessageReject
}
type FIX50MessageCracker struct{}

func (c *FIX50MessageCracker) OnFIX50IOI(msg IOI, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50Advertisement(msg Advertisement, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ExecutionReport(msg ExecutionReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderCancelReject(msg OrderCancelReject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50News(msg News, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50Email(msg Email, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50NewOrderSingle(msg NewOrderSingle, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50NewOrderList(msg NewOrderList, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderCancelRequest(msg OrderCancelRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderStatusRequest(msg OrderStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AllocationInstruction(msg AllocationInstruction, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ListCancelRequest(msg ListCancelRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ListExecute(msg ListExecute, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ListStatusRequest(msg ListStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ListStatus(msg ListStatus, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AllocationInstructionAck(msg AllocationInstructionAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50DontKnowTrade(msg DontKnowTrade, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50QuoteRequest(msg QuoteRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50Quote(msg Quote, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SettlementInstructions(msg SettlementInstructions, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MarketDataRequest(msg MarketDataRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MarketDataRequestReject(msg MarketDataRequestReject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50QuoteCancel(msg QuoteCancel, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50QuoteStatusRequest(msg QuoteStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityDefinition(msg SecurityDefinition, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityStatusRequest(msg SecurityStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityStatus(msg SecurityStatus, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradingSessionStatus(msg TradingSessionStatus, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MassQuote(msg MassQuote, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50BusinessMessageReject(msg BusinessMessageReject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50BidRequest(msg BidRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50BidResponse(msg BidResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ListStrikePrice(msg ListStrikePrice, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50RegistrationInstructions(msg RegistrationInstructions, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderMassCancelReport(msg OrderMassCancelReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50NewOrderCross(msg NewOrderCross, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityTypeRequest(msg SecurityTypeRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityTypes(msg SecurityTypes, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityListRequest(msg SecurityListRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityList(msg SecurityList, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50DerivativeSecurityList(msg DerivativeSecurityList, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50NewOrderMultileg(msg NewOrderMultileg, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradeCaptureReport(msg TradeCaptureReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50QuoteRequestReject(msg QuoteRequestReject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50RFQRequest(msg RFQRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50QuoteStatusReport(msg QuoteStatusReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50QuoteResponse(msg QuoteResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50Confirmation(msg Confirmation, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50PositionMaintenanceReport(msg PositionMaintenanceReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50RequestForPositions(msg RequestForPositions, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50RequestForPositionsAck(msg RequestForPositionsAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50PositionReport(msg PositionReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradeCaptureReportAck(msg TradeCaptureReportAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AllocationReport(msg AllocationReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AllocationReportAck(msg AllocationReportAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ConfirmationAck(msg ConfirmationAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SettlementInstructionRequest(msg SettlementInstructionRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AssignmentReport(msg AssignmentReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CollateralRequest(msg CollateralRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CollateralAssignment(msg CollateralAssignment, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CollateralResponse(msg CollateralResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CollateralReport(msg CollateralReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CollateralInquiry(msg CollateralInquiry, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50UserRequest(msg UserRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50UserResponse(msg UserResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50CollateralInquiryAck(msg CollateralInquiryAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ConfirmationRequest(msg ConfirmationRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ContraryIntentionReport(msg ContraryIntentionReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityDefinitionUpdateReport(msg SecurityDefinitionUpdateReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50SecurityListUpdateReport(msg SecurityListUpdateReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AdjustedPositionReport(msg AdjustedPositionReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50AllocationInstructionAlert(msg AllocationInstructionAlert, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50ExecutionAcknowledgement(msg ExecutionAcknowledgement, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradingSessionList(msg TradingSessionList, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX50MessageCracker) OnFIX50TradingSessionListRequest(msg TradingSessionListRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
