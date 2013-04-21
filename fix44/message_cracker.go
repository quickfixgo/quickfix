package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/tag"
)

func Crack(msg message.Message, sessionID quickfixgo.SessionID, router MessageRouter) reject.MessageReject {
	switch msgType, _ := msg.Header.StringValue(tag.MsgType); msgType {
	case "0":
		return router.OnFIX44Heartbeat(Heartbeat{msg}, sessionID)
	case "1":
		return router.OnFIX44TestRequest(TestRequest{msg}, sessionID)
	case "2":
		return router.OnFIX44ResendRequest(ResendRequest{msg}, sessionID)
	case "3":
		return router.OnFIX44Reject(Reject{msg}, sessionID)
	case "4":
		return router.OnFIX44SequenceReset(SequenceReset{msg}, sessionID)
	case "5":
		return router.OnFIX44Logout(Logout{msg}, sessionID)
	case "6":
		return router.OnFIX44IOI(IOI{msg}, sessionID)
	case "7":
		return router.OnFIX44Advertisement(Advertisement{msg}, sessionID)
	case "8":
		return router.OnFIX44ExecutionReport(ExecutionReport{msg}, sessionID)
	case "9":
		return router.OnFIX44OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "A":
		return router.OnFIX44Logon(Logon{msg}, sessionID)
	case "B":
		return router.OnFIX44News(News{msg}, sessionID)
	case "C":
		return router.OnFIX44Email(Email{msg}, sessionID)
	case "D":
		return router.OnFIX44NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "E":
		return router.OnFIX44NewOrderList(NewOrderList{msg}, sessionID)
	case "F":
		return router.OnFIX44OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "G":
		return router.OnFIX44OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "H":
		return router.OnFIX44OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "J":
		return router.OnFIX44AllocationInstruction(AllocationInstruction{msg}, sessionID)
	case "K":
		return router.OnFIX44ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "L":
		return router.OnFIX44ListExecute(ListExecute{msg}, sessionID)
	case "M":
		return router.OnFIX44ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "N":
		return router.OnFIX44ListStatus(ListStatus{msg}, sessionID)
	case "P":
		return router.OnFIX44AllocationInstructionAck(AllocationInstructionAck{msg}, sessionID)
	case "Q":
		return router.OnFIX44DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "R":
		return router.OnFIX44QuoteRequest(QuoteRequest{msg}, sessionID)
	case "S":
		return router.OnFIX44Quote(Quote{msg}, sessionID)
	case "T":
		return router.OnFIX44SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "V":
		return router.OnFIX44MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "W":
		return router.OnFIX44MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "X":
		return router.OnFIX44MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "Y":
		return router.OnFIX44MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "Z":
		return router.OnFIX44QuoteCancel(QuoteCancel{msg}, sessionID)
	case "a":
		return router.OnFIX44QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "b":
		return router.OnFIX44MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "c":
		return router.OnFIX44SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "d":
		return router.OnFIX44SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "e":
		return router.OnFIX44SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "f":
		return router.OnFIX44SecurityStatus(SecurityStatus{msg}, sessionID)
	case "g":
		return router.OnFIX44TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "h":
		return router.OnFIX44TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "i":
		return router.OnFIX44MassQuote(MassQuote{msg}, sessionID)
	case "j":
		return router.OnFIX44BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "k":
		return router.OnFIX44BidRequest(BidRequest{msg}, sessionID)
	case "l":
		return router.OnFIX44BidResponse(BidResponse{msg}, sessionID)
	case "m":
		return router.OnFIX44ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "o":
		return router.OnFIX44RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "p":
		return router.OnFIX44RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "q":
		return router.OnFIX44OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "r":
		return router.OnFIX44OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "s":
		return router.OnFIX44NewOrderCross(NewOrderCross{msg}, sessionID)
	case "t":
		return router.OnFIX44CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "u":
		return router.OnFIX44CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "v":
		return router.OnFIX44SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "w":
		return router.OnFIX44SecurityTypes(SecurityTypes{msg}, sessionID)
	case "x":
		return router.OnFIX44SecurityListRequest(SecurityListRequest{msg}, sessionID)
	case "y":
		return router.OnFIX44SecurityList(SecurityList{msg}, sessionID)
	case "z":
		return router.OnFIX44DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	case "AA":
		return router.OnFIX44DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "AB":
		return router.OnFIX44NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "AC":
		return router.OnFIX44MultilegOrderCancelReplace(MultilegOrderCancelReplace{msg}, sessionID)
	case "AD":
		return router.OnFIX44TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "AE":
		return router.OnFIX44TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "AF":
		return router.OnFIX44OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "AG":
		return router.OnFIX44QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "AH":
		return router.OnFIX44RFQRequest(RFQRequest{msg}, sessionID)
	case "AI":
		return router.OnFIX44QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	case "AJ":
		return router.OnFIX44QuoteResponse(QuoteResponse{msg}, sessionID)
	case "AK":
		return router.OnFIX44Confirmation(Confirmation{msg}, sessionID)
	case "AL":
		return router.OnFIX44PositionMaintenanceRequest(PositionMaintenanceRequest{msg}, sessionID)
	case "AM":
		return router.OnFIX44PositionMaintenanceReport(PositionMaintenanceReport{msg}, sessionID)
	case "AN":
		return router.OnFIX44RequestForPositions(RequestForPositions{msg}, sessionID)
	case "AO":
		return router.OnFIX44RequestForPositionsAck(RequestForPositionsAck{msg}, sessionID)
	case "AP":
		return router.OnFIX44PositionReport(PositionReport{msg}, sessionID)
	case "AQ":
		return router.OnFIX44TradeCaptureReportRequestAck(TradeCaptureReportRequestAck{msg}, sessionID)
	case "AR":
		return router.OnFIX44TradeCaptureReportAck(TradeCaptureReportAck{msg}, sessionID)
	case "AS":
		return router.OnFIX44AllocationReport(AllocationReport{msg}, sessionID)
	case "AT":
		return router.OnFIX44AllocationReportAck(AllocationReportAck{msg}, sessionID)
	case "AU":
		return router.OnFIX44ConfirmationAck(ConfirmationAck{msg}, sessionID)
	case "AV":
		return router.OnFIX44SettlementInstructionRequest(SettlementInstructionRequest{msg}, sessionID)
	case "AW":
		return router.OnFIX44AssignmentReport(AssignmentReport{msg}, sessionID)
	case "AX":
		return router.OnFIX44CollateralRequest(CollateralRequest{msg}, sessionID)
	case "AY":
		return router.OnFIX44CollateralAssignment(CollateralAssignment{msg}, sessionID)
	case "AZ":
		return router.OnFIX44CollateralResponse(CollateralResponse{msg}, sessionID)
	case "BA":
		return router.OnFIX44CollateralReport(CollateralReport{msg}, sessionID)
	case "BB":
		return router.OnFIX44CollateralInquiry(CollateralInquiry{msg}, sessionID)
	case "BC":
		return router.OnFIX44NetworkCounterpartySystemStatusRequest(NetworkCounterpartySystemStatusRequest{msg}, sessionID)
	case "BD":
		return router.OnFIX44NetworkCounterpartySystemStatusResponse(NetworkCounterpartySystemStatusResponse{msg}, sessionID)
	case "BE":
		return router.OnFIX44UserRequest(UserRequest{msg}, sessionID)
	case "BF":
		return router.OnFIX44UserResponse(UserResponse{msg}, sessionID)
	case "BG":
		return router.OnFIX44CollateralInquiryAck(CollateralInquiryAck{msg}, sessionID)
	case "BH":
		return router.OnFIX44ConfirmationRequest(ConfirmationRequest{msg}, sessionID)
	}
	return reject.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX44Heartbeat(msg Heartbeat, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44TestRequest(msg TestRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44ResendRequest(msg ResendRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44Reject(msg Reject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44SequenceReset(msg SequenceReset, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44Logout(msg Logout, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44IOI(msg IOI, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44Advertisement(msg Advertisement, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44ExecutionReport(msg ExecutionReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44OrderCancelReject(msg OrderCancelReject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44Logon(msg Logon, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44News(msg News, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44Email(msg Email, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44NewOrderSingle(msg NewOrderSingle, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44NewOrderList(msg NewOrderList, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44OrderCancelRequest(msg OrderCancelRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44OrderStatusRequest(msg OrderStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44AllocationInstruction(msg AllocationInstruction, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44ListCancelRequest(msg ListCancelRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44ListExecute(msg ListExecute, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44ListStatusRequest(msg ListStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44ListStatus(msg ListStatus, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44AllocationInstructionAck(msg AllocationInstructionAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44DontKnowTrade(msg DontKnowTrade, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44QuoteRequest(msg QuoteRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44Quote(msg Quote, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44SettlementInstructions(msg SettlementInstructions, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44MarketDataRequest(msg MarketDataRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44QuoteCancel(msg QuoteCancel, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44SecurityDefinition(msg SecurityDefinition, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44SecurityStatus(msg SecurityStatus, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44TradingSessionStatus(msg TradingSessionStatus, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44MassQuote(msg MassQuote, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44BusinessMessageReject(msg BusinessMessageReject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44BidRequest(msg BidRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44BidResponse(msg BidResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44ListStrikePrice(msg ListStrikePrice, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44RegistrationInstructions(msg RegistrationInstructions, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44NewOrderCross(msg NewOrderCross, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44SecurityTypes(msg SecurityTypes, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44SecurityListRequest(msg SecurityListRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44SecurityList(msg SecurityList, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44NewOrderMultileg(msg NewOrderMultileg, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44TradeCaptureReport(msg TradeCaptureReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44QuoteRequestReject(msg QuoteRequestReject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44RFQRequest(msg RFQRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44QuoteStatusReport(msg QuoteStatusReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44QuoteResponse(msg QuoteResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44Confirmation(msg Confirmation, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44PositionMaintenanceReport(msg PositionMaintenanceReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44RequestForPositions(msg RequestForPositions, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44RequestForPositionsAck(msg RequestForPositionsAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44PositionReport(msg PositionReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44TradeCaptureReportAck(msg TradeCaptureReportAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44AllocationReport(msg AllocationReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44AllocationReportAck(msg AllocationReportAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44ConfirmationAck(msg ConfirmationAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44SettlementInstructionRequest(msg SettlementInstructionRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44AssignmentReport(msg AssignmentReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44CollateralRequest(msg CollateralRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44CollateralAssignment(msg CollateralAssignment, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44CollateralResponse(msg CollateralResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44CollateralReport(msg CollateralReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44CollateralInquiry(msg CollateralInquiry, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44UserRequest(msg UserRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44UserResponse(msg UserResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44CollateralInquiryAck(msg CollateralInquiryAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX44ConfirmationRequest(msg ConfirmationRequest, sessionID quickfixgo.SessionID) reject.MessageReject
}
type FIX44MessageCracker struct{}

func (c *FIX44MessageCracker) OnFIX44Heartbeat(msg Heartbeat, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TestRequest(msg TestRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ResendRequest(msg ResendRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Reject(msg Reject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SequenceReset(msg SequenceReset, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Logout(msg Logout, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44IOI(msg IOI, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Advertisement(msg Advertisement, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ExecutionReport(msg ExecutionReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderCancelReject(msg OrderCancelReject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Logon(msg Logon, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44News(msg News, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Email(msg Email, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44NewOrderSingle(msg NewOrderSingle, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44NewOrderList(msg NewOrderList, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderCancelRequest(msg OrderCancelRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderStatusRequest(msg OrderStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44AllocationInstruction(msg AllocationInstruction, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ListCancelRequest(msg ListCancelRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ListExecute(msg ListExecute, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ListStatusRequest(msg ListStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ListStatus(msg ListStatus, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44AllocationInstructionAck(msg AllocationInstructionAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44DontKnowTrade(msg DontKnowTrade, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44QuoteRequest(msg QuoteRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Quote(msg Quote, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SettlementInstructions(msg SettlementInstructions, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MarketDataRequest(msg MarketDataRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MarketDataRequestReject(msg MarketDataRequestReject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44QuoteCancel(msg QuoteCancel, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44QuoteStatusRequest(msg QuoteStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityDefinition(msg SecurityDefinition, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityStatusRequest(msg SecurityStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityStatus(msg SecurityStatus, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TradingSessionStatus(msg TradingSessionStatus, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MassQuote(msg MassQuote, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44BusinessMessageReject(msg BusinessMessageReject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44BidRequest(msg BidRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44BidResponse(msg BidResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ListStrikePrice(msg ListStrikePrice, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44RegistrationInstructions(msg RegistrationInstructions, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderMassCancelReport(msg OrderMassCancelReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44NewOrderCross(msg NewOrderCross, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityTypeRequest(msg SecurityTypeRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityTypes(msg SecurityTypes, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityListRequest(msg SecurityListRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SecurityList(msg SecurityList, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44DerivativeSecurityList(msg DerivativeSecurityList, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44NewOrderMultileg(msg NewOrderMultileg, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44MultilegOrderCancelReplace(msg MultilegOrderCancelReplace, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TradeCaptureReport(msg TradeCaptureReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44QuoteRequestReject(msg QuoteRequestReject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44RFQRequest(msg RFQRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44QuoteStatusReport(msg QuoteStatusReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44QuoteResponse(msg QuoteResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44Confirmation(msg Confirmation, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44PositionMaintenanceRequest(msg PositionMaintenanceRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44PositionMaintenanceReport(msg PositionMaintenanceReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44RequestForPositions(msg RequestForPositions, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44RequestForPositionsAck(msg RequestForPositionsAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44PositionReport(msg PositionReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TradeCaptureReportRequestAck(msg TradeCaptureReportRequestAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44TradeCaptureReportAck(msg TradeCaptureReportAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44AllocationReport(msg AllocationReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44AllocationReportAck(msg AllocationReportAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ConfirmationAck(msg ConfirmationAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44SettlementInstructionRequest(msg SettlementInstructionRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44AssignmentReport(msg AssignmentReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CollateralRequest(msg CollateralRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CollateralAssignment(msg CollateralAssignment, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CollateralResponse(msg CollateralResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CollateralReport(msg CollateralReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CollateralInquiry(msg CollateralInquiry, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44NetworkCounterpartySystemStatusRequest(msg NetworkCounterpartySystemStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44NetworkCounterpartySystemStatusResponse(msg NetworkCounterpartySystemStatusResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44UserRequest(msg UserRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44UserResponse(msg UserResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44CollateralInquiryAck(msg CollateralInquiryAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX44MessageCracker) OnFIX44ConfirmationRequest(msg ConfirmationRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
