package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type TradeCaptureReport struct {
	quickfixgo.Message
}

func (m *TradeCaptureReport) TradeReportID() (*field.TradeReportID, error) {
	f := new(field.TradeReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeReportTransType() (*field.TradeReportTransType, error) {
	f := new(field.TradeReportTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeReportType() (*field.TradeReportType, error) {
	f := new(field.TradeReportType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeRequestID() (*field.TradeRequestID, error) {
	f := new(field.TradeRequestID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TrdType() (*field.TrdType, error) {
	f := new(field.TrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TrdSubType() (*field.TrdSubType, error) {
	f := new(field.TrdSubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SecondaryTrdType() (*field.SecondaryTrdType, error) {
	f := new(field.SecondaryTrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TransferReason() (*field.TransferReason, error) {
	f := new(field.TransferReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) ExecType() (*field.ExecType, error) {
	f := new(field.ExecType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TotNumTradeReports() (*field.TotNumTradeReports, error) {
	f := new(field.TotNumTradeReports)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastRptRequested() (*field.LastRptRequested, error) {
	f := new(field.LastRptRequested)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) UnsolicitedIndicator() (*field.UnsolicitedIndicator, error) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeReportRefID() (*field.TradeReportRefID, error) {
	f := new(field.TradeReportRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SecondaryTradeReportRefID() (*field.SecondaryTradeReportRefID, error) {
	f := new(field.SecondaryTradeReportRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SecondaryTradeReportID() (*field.SecondaryTradeReportID, error) {
	f := new(field.SecondaryTradeReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeLinkID() (*field.TradeLinkID, error) {
	f := new(field.TradeLinkID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TrdMatchID() (*field.TrdMatchID, error) {
	f := new(field.TrdMatchID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) OrdStatus() (*field.OrdStatus, error) {
	f := new(field.OrdStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SecondaryExecID() (*field.SecondaryExecID, error) {
	f := new(field.SecondaryExecID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) ExecRestatementReason() (*field.ExecRestatementReason, error) {
	f := new(field.ExecRestatementReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) PreviouslyReported() (*field.PreviouslyReported, error) {
	f := new(field.PreviouslyReported)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) UnderlyingTradingSessionID() (*field.UnderlyingTradingSessionID, error) {
	f := new(field.UnderlyingTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) UnderlyingTradingSessionSubID() (*field.UnderlyingTradingSessionSubID, error) {
	f := new(field.UnderlyingTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastQty() (*field.LastQty, error) {
	f := new(field.LastQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastParPx() (*field.LastParPx, error) {
	f := new(field.LastParPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastSpotRate() (*field.LastSpotRate, error) {
	f := new(field.LastSpotRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastForwardPoints() (*field.LastForwardPoints, error) {
	f := new(field.LastForwardPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) AvgPxIndicator() (*field.AvgPxIndicator, error) {
	f := new(field.AvgPxIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) MultiLegReportingType() (*field.MultiLegReportingType, error) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeLegRefID() (*field.TradeLegRefID, error) {
	f := new(field.TradeLegRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) MatchType() (*field.MatchType, error) {
	f := new(field.MatchType)
	err := m.Body.Get(f)
	return f, err
}
