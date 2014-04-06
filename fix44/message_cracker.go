package fix44

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

func Crack(msg quickfix.Message, sessionID quickfix.SessionID, router MessageRouter) quickfix.MessageReject {

	msgType := new(field.MsgType)
	switch msg.Header.Get(msgType); msgType.Value {
	case "4":
		return router.OnFIX44SequenceReset(SequenceReset{msg}, sessionID)
	case "7":
		return router.OnFIX44Advertisement(Advertisement{msg}, sessionID)
	case "L":
		return router.OnFIX44ListExecute(ListExecute{msg}, sessionID)
	case "h":
		return router.OnFIX44TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "k":
		return router.OnFIX44BidRequest(BidRequest{msg}, sessionID)
	case "q":
		return router.OnFIX44OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "AA":
		return router.OnFIX44DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "AE":
		return router.OnFIX44TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "BG":
		return router.OnFIX44CollateralInquiryAck(CollateralInquiryAck{msg}, sessionID)
	case "A":
		return router.OnFIX44Logon(Logon{msg}, sessionID)
	case "R":
		return router.OnFIX44QuoteRequest(QuoteRequest{msg}, sessionID)
	case "V":
		return router.OnFIX44MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "i":
		return router.OnFIX44MassQuote(MassQuote{msg}, sessionID)
	case "u":
		return router.OnFIX44CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "AG":
		return router.OnFIX44QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "AQ":
		return router.OnFIX44TradeCaptureReportRequestAck(TradeCaptureReportRequestAck{msg}, sessionID)
	case "BH":
		return router.OnFIX44ConfirmationRequest(ConfirmationRequest{msg}, sessionID)
	case "9":
		return router.OnFIX44OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "w":
		return router.OnFIX44SecurityTypes(SecurityTypes{msg}, sessionID)
	case "AK":
		return router.OnFIX44Confirmation(Confirmation{msg}, sessionID)
	case "AS":
		return router.OnFIX44AllocationReport(AllocationReport{msg}, sessionID)
	case "AW":
		return router.OnFIX44AssignmentReport(AssignmentReport{msg}, sessionID)
	case "5":
		return router.OnFIX44Logout(Logout{msg}, sessionID)
	case "D":
		return router.OnFIX44NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "P":
		return router.OnFIX44AllocationInstructionAck(AllocationInstructionAck{msg}, sessionID)
	case "d":
		return router.OnFIX44SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "y":
		return router.OnFIX44SecurityList(SecurityList{msg}, sessionID)
	case "AU":
		return router.OnFIX44ConfirmationAck(ConfirmationAck{msg}, sessionID)
	case "BD":
		return router.OnFIX44NetworkCounterpartySystemStatusResponse(NetworkCounterpartySystemStatusResponse{msg}, sessionID)
	case "H":
		return router.OnFIX44OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "Q":
		return router.OnFIX44DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "S":
		return router.OnFIX44Quote(Quote{msg}, sessionID)
	case "c":
		return router.OnFIX44SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "l":
		return router.OnFIX44BidResponse(BidResponse{msg}, sessionID)
	case "AL":
		return router.OnFIX44PositionMaintenanceRequest(PositionMaintenanceRequest{msg}, sessionID)
	case "AM":
		return router.OnFIX44PositionMaintenanceReport(PositionMaintenanceReport{msg}, sessionID)
	case "BC":
		return router.OnFIX44NetworkCounterpartySystemStatusRequest(NetworkCounterpartySystemStatusRequest{msg}, sessionID)
	case "r":
		return router.OnFIX44OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "s":
		return router.OnFIX44NewOrderCross(NewOrderCross{msg}, sessionID)
	case "AB":
		return router.OnFIX44NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "AR":
		return router.OnFIX44TradeCaptureReportAck(TradeCaptureReportAck{msg}, sessionID)
	case "T":
		return router.OnFIX44SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "W":
		return router.OnFIX44MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "BB":
		return router.OnFIX44CollateralInquiry(CollateralInquiry{msg}, sessionID)
	case "BE":
		return router.OnFIX44UserRequest(UserRequest{msg}, sessionID)
	case "G":
		return router.OnFIX44OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "Y":
		return router.OnFIX44MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "AP":
		return router.OnFIX44PositionReport(PositionReport{msg}, sessionID)
	case "AT":
		return router.OnFIX44AllocationReportAck(AllocationReportAck{msg}, sessionID)
	case "0":
		return router.OnFIX44Heartbeat(Heartbeat{msg}, sessionID)
	case "C":
		return router.OnFIX44Email(Email{msg}, sessionID)
	case "N":
		return router.OnFIX44ListStatus(ListStatus{msg}, sessionID)
	case "e":
		return router.OnFIX44SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "z":
		return router.OnFIX44DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	case "AY":
		return router.OnFIX44CollateralAssignment(CollateralAssignment{msg}, sessionID)
	case "AZ":
		return router.OnFIX44CollateralResponse(CollateralResponse{msg}, sessionID)
	case "2":
		return router.OnFIX44ResendRequest(ResendRequest{msg}, sessionID)
	case "J":
		return router.OnFIX44AllocationInstruction(AllocationInstruction{msg}, sessionID)
	case "p":
		return router.OnFIX44RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "E":
		return router.OnFIX44NewOrderList(NewOrderList{msg}, sessionID)
	case "g":
		return router.OnFIX44TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "AX":
		return router.OnFIX44CollateralRequest(CollateralRequest{msg}, sessionID)
	case "1":
		return router.OnFIX44TestRequest(TestRequest{msg}, sessionID)
	case "6":
		return router.OnFIX44IOI(IOI{msg}, sessionID)
	case "F":
		return router.OnFIX44OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "b":
		return router.OnFIX44MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "AC":
		return router.OnFIX44MultilegOrderCancelReplace(MultilegOrderCancelReplace{msg}, sessionID)
	case "AI":
		return router.OnFIX44QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	case "8":
		return router.OnFIX44ExecutionReport(ExecutionReport{msg}, sessionID)
	case "B":
		return router.OnFIX44News(News{msg}, sessionID)
	case "K":
		return router.OnFIX44ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "X":
		return router.OnFIX44MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "o":
		return router.OnFIX44RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "v":
		return router.OnFIX44SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "AH":
		return router.OnFIX44RFQRequest(RFQRequest{msg}, sessionID)
	case "AJ":
		return router.OnFIX44QuoteResponse(QuoteResponse{msg}, sessionID)
	case "AO":
		return router.OnFIX44RequestForPositionsAck(RequestForPositionsAck{msg}, sessionID)
	case "M":
		return router.OnFIX44ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "a":
		return router.OnFIX44QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "m":
		return router.OnFIX44ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "t":
		return router.OnFIX44CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "AD":
		return router.OnFIX44TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "AV":
		return router.OnFIX44SettlementInstructionRequest(SettlementInstructionRequest{msg}, sessionID)
	case "BA":
		return router.OnFIX44CollateralReport(CollateralReport{msg}, sessionID)
	case "BF":
		return router.OnFIX44UserResponse(UserResponse{msg}, sessionID)
	case "Z":
		return router.OnFIX44QuoteCancel(QuoteCancel{msg}, sessionID)
	case "AF":
		return router.OnFIX44OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "AN":
		return router.OnFIX44RequestForPositions(RequestForPositions{msg}, sessionID)
	case "3":
		return router.OnFIX44Reject(Reject{msg}, sessionID)
	case "f":
		return router.OnFIX44SecurityStatus(SecurityStatus{msg}, sessionID)
	case "j":
		return router.OnFIX44BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "x":
		return router.OnFIX44SecurityListRequest(SecurityListRequest{msg}, sessionID)
	}
	return quickfix.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX44OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44NewOrderCross(msg NewOrderCross, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44NewOrderMultileg(msg NewOrderMultileg, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44SettlementInstructions(msg SettlementInstructions, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44CollateralInquiry(msg CollateralInquiry, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44UserRequest(msg UserRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44PositionReport(msg PositionReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44AllocationReportAck(msg AllocationReportAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44Heartbeat(msg Heartbeat, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44Email(msg Email, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44ListStatus(msg ListStatus, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44CollateralAssignment(msg CollateralAssignment, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44CollateralResponse(msg CollateralResponse, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44ResendRequest(msg ResendRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44AllocationInstruction(msg AllocationInstruction, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44NewOrderList(msg NewOrderList, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44CollateralRequest(msg CollateralRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44TestRequest(msg TestRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44IOI(msg IOI, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44OrderCancelRequest(msg OrderCancelRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44QuoteStatusReport(msg QuoteStatusReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44ExecutionReport(msg ExecutionReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44News(msg News, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44ListCancelRequest(msg ListCancelRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44RegistrationInstructions(msg RegistrationInstructions, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44RFQRequest(msg RFQRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44QuoteResponse(msg QuoteResponse, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44RequestForPositionsAck(msg RequestForPositionsAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44ListStatusRequest(msg ListStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44ListStrikePrice(msg ListStrikePrice, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44CollateralReport(msg CollateralReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44UserResponse(msg UserResponse, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44QuoteCancel(msg QuoteCancel, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44RequestForPositions(msg RequestForPositions, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44Reject(msg Reject, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44SecurityStatus(msg SecurityStatus, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44BusinessMessageReject(msg BusinessMessageReject, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44SecurityListRequest(msg SecurityListRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44SequenceReset(msg SequenceReset, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44Advertisement(msg Advertisement, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44ListExecute(msg ListExecute, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44TradingSessionStatus(msg TradingSessionStatus, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44BidRequest(msg BidRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44TradeCaptureReport(msg TradeCaptureReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44CollateralInquiryAck(msg CollateralInquiryAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44Logon(msg Logon, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44QuoteRequest(msg QuoteRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44MarketDataRequest(msg MarketDataRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44MassQuote(msg MassQuote, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44QuoteRequestReject(msg QuoteRequestReject, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44ConfirmationRequest(msg ConfirmationRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44OrderCancelReject(msg OrderCancelReject, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44SecurityTypes(msg SecurityTypes, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44Confirmation(msg Confirmation, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44AllocationReport(msg AllocationReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44AssignmentReport(msg AssignmentReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44Logout(msg Logout, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44NewOrderSingle(msg NewOrderSingle, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44AllocationInstructionAck(msg AllocationInstructionAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44SecurityDefinition(msg SecurityDefinition, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44SecurityList(msg SecurityList, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44ConfirmationAck(msg ConfirmationAck, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44OrderStatusRequest(msg OrderStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44DontKnowTrade(msg DontKnowTrade, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44Quote(msg Quote, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44BidResponse(msg BidResponse, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID quickfix.SessionID) quickfix.MessageReject
	OnFIX44NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID quickfix.SessionID) quickfix.MessageReject
}
type FIX44MessageCracker struct{}

func (c *FIX44MessageCracker) OnFIX44OrderCancelReject(msg OrderCancelReject, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityTypes(msg SecurityTypes, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Confirmation(msg Confirmation, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44AllocationReport(msg AllocationReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44AssignmentReport(msg AssignmentReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Logout(msg Logout, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44NewOrderSingle(msg NewOrderSingle, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44AllocationInstructionAck(msg AllocationInstructionAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityDefinition(msg SecurityDefinition, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityList(msg SecurityList, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ConfirmationAck(msg ConfirmationAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderStatusRequest(msg OrderStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44DontKnowTrade(msg DontKnowTrade, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Quote(msg Quote, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44BidResponse(msg BidResponse, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44PositionMaintenanceReport(msg PositionMaintenanceReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderMassCancelReport(msg OrderMassCancelReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44NewOrderCross(msg NewOrderCross, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44NewOrderMultileg(msg NewOrderMultileg, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TradeCaptureReportAck(msg TradeCaptureReportAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SettlementInstructions(msg SettlementInstructions, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CollateralInquiry(msg CollateralInquiry, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44UserRequest(msg UserRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MarketDataRequestReject(msg MarketDataRequestReject, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44PositionReport(msg PositionReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44AllocationReportAck(msg AllocationReportAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Heartbeat(msg Heartbeat, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Email(msg Email, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ListStatus(msg ListStatus, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityStatusRequest(msg SecurityStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CollateralAssignment(msg CollateralAssignment, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CollateralResponse(msg CollateralResponse, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ResendRequest(msg ResendRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44AllocationInstruction(msg AllocationInstruction, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44NewOrderList(msg NewOrderList, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CollateralRequest(msg CollateralRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TestRequest(msg TestRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44IOI(msg IOI, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderCancelRequest(msg OrderCancelRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44QuoteStatusReport(msg QuoteStatusReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ExecutionReport(msg ExecutionReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44News(msg News, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ListCancelRequest(msg ListCancelRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44RegistrationInstructions(msg RegistrationInstructions, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityTypeRequest(msg SecurityTypeRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44RFQRequest(msg RFQRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44QuoteResponse(msg QuoteResponse, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44RequestForPositionsAck(msg RequestForPositionsAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ListStatusRequest(msg ListStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44QuoteStatusRequest(msg QuoteStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ListStrikePrice(msg ListStrikePrice, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SettlementInstructionRequest(msg SettlementInstructionRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CollateralReport(msg CollateralReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44UserResponse(msg UserResponse, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44QuoteCancel(msg QuoteCancel, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44RequestForPositions(msg RequestForPositions, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Reject(msg Reject, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityStatus(msg SecurityStatus, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44BusinessMessageReject(msg BusinessMessageReject, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityListRequest(msg SecurityListRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SequenceReset(msg SequenceReset, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Advertisement(msg Advertisement, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ListExecute(msg ListExecute, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TradingSessionStatus(msg TradingSessionStatus, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44BidRequest(msg BidRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44DerivativeSecurityList(msg DerivativeSecurityList, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TradeCaptureReport(msg TradeCaptureReport, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CollateralInquiryAck(msg CollateralInquiryAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Logon(msg Logon, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44QuoteRequest(msg QuoteRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MarketDataRequest(msg MarketDataRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MassQuote(msg MassQuote, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44QuoteRequestReject(msg QuoteRequestReject, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ConfirmationRequest(msg ConfirmationRequest, sessionId quickfix.SessionID) quickfix.MessageReject {
	return quickfix.NewUnsupportedMessageType(msg.Message)
}
