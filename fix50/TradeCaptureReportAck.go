package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type TradeCaptureReportAck struct {
	quickfixgo.Message
}

func (m *TradeCaptureReportAck) TradeReportID() (*field.TradeReportID, error) {
	f := new(field.TradeReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TradeReportTransType() (*field.TradeReportTransType, error) {
	f := new(field.TradeReportTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TradeReportType() (*field.TradeReportType, error) {
	f := new(field.TradeReportType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TrdType() (*field.TrdType, error) {
	f := new(field.TrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TrdSubType() (*field.TrdSubType, error) {
	f := new(field.TrdSubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) SecondaryTrdType() (*field.SecondaryTrdType, error) {
	f := new(field.SecondaryTrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TransferReason() (*field.TransferReason, error) {
	f := new(field.TransferReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) ExecType() (*field.ExecType, error) {
	f := new(field.ExecType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TradeReportRefID() (*field.TradeReportRefID, error) {
	f := new(field.TradeReportRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) SecondaryTradeReportRefID() (*field.SecondaryTradeReportRefID, error) {
	f := new(field.SecondaryTradeReportRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TrdRptStatus() (*field.TrdRptStatus, error) {
	f := new(field.TrdRptStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TradeReportRejectReason() (*field.TradeReportRejectReason, error) {
	f := new(field.TradeReportRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) SecondaryTradeReportID() (*field.SecondaryTradeReportID, error) {
	f := new(field.SecondaryTradeReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TradeLinkID() (*field.TradeLinkID, error) {
	f := new(field.TradeLinkID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TrdMatchID() (*field.TrdMatchID, error) {
	f := new(field.TrdMatchID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) SecondaryExecID() (*field.SecondaryExecID, error) {
	f := new(field.SecondaryExecID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) ResponseTransportType() (*field.ResponseTransportType, error) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) ResponseDestination() (*field.ResponseDestination, error) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) ClearingFeeIndicator() (*field.ClearingFeeIndicator, error) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) OrdStatus() (*field.OrdStatus, error) {
	f := new(field.OrdStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) ExecRestatementReason() (*field.ExecRestatementReason, error) {
	f := new(field.ExecRestatementReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) PreviouslyReported() (*field.PreviouslyReported, error) {
	f := new(field.PreviouslyReported)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) UnderlyingTradingSessionID() (*field.UnderlyingTradingSessionID, error) {
	f := new(field.UnderlyingTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) UnderlyingTradingSessionSubID() (*field.UnderlyingTradingSessionSubID, error) {
	f := new(field.UnderlyingTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) LastQty() (*field.LastQty, error) {
	f := new(field.LastQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) LastParPx() (*field.LastParPx, error) {
	f := new(field.LastParPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) LastSpotRate() (*field.LastSpotRate, error) {
	f := new(field.LastSpotRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) LastForwardPoints() (*field.LastForwardPoints, error) {
	f := new(field.LastForwardPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) AvgPxIndicator() (*field.AvgPxIndicator, error) {
	f := new(field.AvgPxIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) MultiLegReportingType() (*field.MultiLegReportingType, error) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TradeLegRefID() (*field.TradeLegRefID, error) {
	f := new(field.TradeLegRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) MatchType() (*field.MatchType, error) {
	f := new(field.MatchType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) CopyMsgIndicator() (*field.CopyMsgIndicator, error) {
	f := new(field.CopyMsgIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) PublishTrdIndicator() (*field.PublishTrdIndicator, error) {
	f := new(field.PublishTrdIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) ShortSaleReason() (*field.ShortSaleReason, error) {
	f := new(field.ShortSaleReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TierCode() (*field.TierCode, error) {
	f := new(field.TierCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) MessageEventSource() (*field.MessageEventSource, error) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) LastUpdateTime() (*field.LastUpdateTime, error) {
	f := new(field.LastUpdateTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) RndPx() (*field.RndPx, error) {
	f := new(field.RndPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) AsOfIndicator() (*field.AsOfIndicator, error) {
	f := new(field.AsOfIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TradeID() (*field.TradeID, error) {
	f := new(field.TradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) SecondaryTradeID() (*field.SecondaryTradeID, error) {
	f := new(field.SecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) FirmTradeID() (*field.FirmTradeID, error) {
	f := new(field.FirmTradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) SecondaryFirmTradeID() (*field.SecondaryFirmTradeID, error) {
	f := new(field.SecondaryFirmTradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) CalculatedCcyLastQty() (*field.CalculatedCcyLastQty, error) {
	f := new(field.CalculatedCcyLastQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) LastSwapPoints() (*field.LastSwapPoints, error) {
	f := new(field.LastSwapPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) TradeHandlingInstr() (*field.TradeHandlingInstr, error) {
	f := new(field.TradeHandlingInstr)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) OrigTradeHandlingInstr() (*field.OrigTradeHandlingInstr, error) {
	f := new(field.OrigTradeHandlingInstr)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) OrigTradeDate() (*field.OrigTradeDate, error) {
	f := new(field.OrigTradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) OrigTradeID() (*field.OrigTradeID, error) {
	f := new(field.OrigTradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) OrigSecondaryTradeID() (*field.OrigSecondaryTradeID, error) {
	f := new(field.OrigSecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportAck) RptSys() (*field.RptSys, error) {
	f := new(field.RptSys)
	err := m.Body.Get(f)
	return f, err
}
