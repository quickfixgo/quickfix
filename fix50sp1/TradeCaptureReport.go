package fix50sp1

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
func (m *TradeCaptureReport) CopyMsgIndicator() (*field.CopyMsgIndicator, error) {
	f := new(field.CopyMsgIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) PublishTrdIndicator() (*field.PublishTrdIndicator, error) {
	f := new(field.PublishTrdIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) ShortSaleReason() (*field.ShortSaleReason, error) {
	f := new(field.ShortSaleReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TrdRptStatus() (*field.TrdRptStatus, error) {
	f := new(field.TrdRptStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) AsOfIndicator() (*field.AsOfIndicator, error) {
	f := new(field.AsOfIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TierCode() (*field.TierCode, error) {
	f := new(field.TierCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) MessageEventSource() (*field.MessageEventSource, error) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastUpdateTime() (*field.LastUpdateTime, error) {
	f := new(field.LastUpdateTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) RndPx() (*field.RndPx, error) {
	f := new(field.RndPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeID() (*field.TradeID, error) {
	f := new(field.TradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SecondaryTradeID() (*field.SecondaryTradeID, error) {
	f := new(field.SecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) FirmTradeID() (*field.FirmTradeID, error) {
	f := new(field.FirmTradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SecondaryFirmTradeID() (*field.SecondaryFirmTradeID, error) {
	f := new(field.SecondaryFirmTradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) CalculatedCcyLastQty() (*field.CalculatedCcyLastQty, error) {
	f := new(field.CalculatedCcyLastQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastSwapPoints() (*field.LastSwapPoints, error) {
	f := new(field.LastSwapPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) UnderlyingSettlementDate() (*field.UnderlyingSettlementDate, error) {
	f := new(field.UnderlyingSettlementDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) OrderCategory() (*field.OrderCategory, error) {
	f := new(field.OrderCategory)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeHandlingInstr() (*field.TradeHandlingInstr, error) {
	f := new(field.TradeHandlingInstr)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) OrigTradeHandlingInstr() (*field.OrigTradeHandlingInstr, error) {
	f := new(field.OrigTradeHandlingInstr)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) OrigTradeDate() (*field.OrigTradeDate, error) {
	f := new(field.OrigTradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) OrigTradeID() (*field.OrigTradeID, error) {
	f := new(field.OrigTradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) OrigSecondaryTradeID() (*field.OrigSecondaryTradeID, error) {
	f := new(field.OrigSecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TZTransactTime() (*field.TZTransactTime, error) {
	f := new(field.TZTransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) ReportedPxDiff() (*field.ReportedPxDiff, error) {
	f := new(field.ReportedPxDiff)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) RejectText() (*field.RejectText, error) {
	f := new(field.RejectText)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) FeeMultiplier() (*field.FeeMultiplier, error) {
	f := new(field.FeeMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) Volatility() (*field.Volatility, error) {
	f := new(field.Volatility)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) DividendYield() (*field.DividendYield, error) {
	f := new(field.DividendYield)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) RiskFreeRate() (*field.RiskFreeRate, error) {
	f := new(field.RiskFreeRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) CurrencyRatio() (*field.CurrencyRatio, error) {
	f := new(field.CurrencyRatio)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradePublishIndicator() (*field.TradePublishIndicator, error) {
	f := new(field.TradePublishIndicator)
	err := m.Body.Get(f)
	return f, err
}
