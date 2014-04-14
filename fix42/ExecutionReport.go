package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ExecutionReport msg type = 8.
type ExecutionReport struct {
	message.Message
}

//ExecutionReportBuilder builds ExecutionReport messages.
type ExecutionReportBuilder struct {
	message.MessageBuilder
}

//NewExecutionReportBuilder returns an initialized ExecutionReportBuilder with specified required fields.
func NewExecutionReportBuilder(
	orderid field.OrderID,
	execid field.ExecID,
	exectranstype field.ExecTransType,
	exectype field.ExecType,
	ordstatus field.OrdStatus,
	symbol field.Symbol,
	side field.Side,
	leavesqty field.LeavesQty,
	cumqty field.CumQty,
	avgpx field.AvgPx) *ExecutionReportBuilder {
	builder := new(ExecutionReportBuilder)
	builder.Body.Set(orderid)
	builder.Body.Set(execid)
	builder.Body.Set(exectranstype)
	builder.Body.Set(exectype)
	builder.Body.Set(ordstatus)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	builder.Body.Set(leavesqty)
	builder.Body.Set(cumqty)
	builder.Body.Set(avgpx)
	return builder
}

//OrderID is a required field for ExecutionReport.
func (m *ExecutionReport) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryOrderID is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a non-required field for ExecutionReport.
func (m *ExecutionReport) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//OrigClOrdID is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrigClOrdID() (*field.OrigClOrdID, error) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//ClientID is a non-required field for ExecutionReport.
func (m *ExecutionReport) ClientID() (*field.ClientID, error) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}

//ExecBroker is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExecBroker() (*field.ExecBroker, error) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}

//NoContraBrokers is a non-required field for ExecutionReport.
func (m *ExecutionReport) NoContraBrokers() (*field.NoContraBrokers, error) {
	f := new(field.NoContraBrokers)
	err := m.Body.Get(f)
	return f, err
}

//ListID is a non-required field for ExecutionReport.
func (m *ExecutionReport) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//ExecID is a required field for ExecutionReport.
func (m *ExecutionReport) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//ExecTransType is a required field for ExecutionReport.
func (m *ExecutionReport) ExecTransType() (*field.ExecTransType, error) {
	f := new(field.ExecTransType)
	err := m.Body.Get(f)
	return f, err
}

//ExecRefID is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExecRefID() (*field.ExecRefID, error) {
	f := new(field.ExecRefID)
	err := m.Body.Get(f)
	return f, err
}

//ExecType is a required field for ExecutionReport.
func (m *ExecutionReport) ExecType() (*field.ExecType, error) {
	f := new(field.ExecType)
	err := m.Body.Get(f)
	return f, err
}

//OrdStatus is a required field for ExecutionReport.
func (m *ExecutionReport) OrdStatus() (*field.OrdStatus, error) {
	f := new(field.OrdStatus)
	err := m.Body.Get(f)
	return f, err
}

//OrdRejReason is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrdRejReason() (*field.OrdRejReason, error) {
	f := new(field.OrdRejReason)
	err := m.Body.Get(f)
	return f, err
}

//ExecRestatementReason is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExecRestatementReason() (*field.ExecRestatementReason, error) {
	f := new(field.ExecRestatementReason)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for ExecutionReport.
func (m *ExecutionReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//SettlmntTyp is a non-required field for ExecutionReport.
func (m *ExecutionReport) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//FutSettDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a required field for ExecutionReport.
func (m *ExecutionReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for ExecutionReport.
func (m *ExecutionReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//IDSource is a non-required field for ExecutionReport.
func (m *ExecutionReport) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for ExecutionReport.
func (m *ExecutionReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDay is a non-required field for ExecutionReport.
func (m *ExecutionReport) MaturityDay() (*field.MaturityDay, error) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for ExecutionReport.
func (m *ExecutionReport) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for ExecutionReport.
func (m *ExecutionReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for ExecutionReport.
func (m *ExecutionReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for ExecutionReport.
func (m *ExecutionReport) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for ExecutionReport.
func (m *ExecutionReport) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for ExecutionReport.
func (m *ExecutionReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for ExecutionReport.
func (m *ExecutionReport) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for ExecutionReport.
func (m *ExecutionReport) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for ExecutionReport.
func (m *ExecutionReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for ExecutionReport.
func (m *ExecutionReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for ExecutionReport.
func (m *ExecutionReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for ExecutionReport.
func (m *ExecutionReport) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//CashOrderQty is a non-required field for ExecutionReport.
func (m *ExecutionReport) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//OrdType is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//Price is a non-required field for ExecutionReport.
func (m *ExecutionReport) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//StopPx is a non-required field for ExecutionReport.
func (m *ExecutionReport) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//PegDifference is a non-required field for ExecutionReport.
func (m *ExecutionReport) PegDifference() (*field.PegDifference, error) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionInst is a non-required field for ExecutionReport.
func (m *ExecutionReport) DiscretionInst() (*field.DiscretionInst, error) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionOffset is a non-required field for ExecutionReport.
func (m *ExecutionReport) DiscretionOffset() (*field.DiscretionOffset, error) {
	f := new(field.DiscretionOffset)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for ExecutionReport.
func (m *ExecutionReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//ComplianceID is a non-required field for ExecutionReport.
func (m *ExecutionReport) ComplianceID() (*field.ComplianceID, error) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//SolicitedFlag is a non-required field for ExecutionReport.
func (m *ExecutionReport) SolicitedFlag() (*field.SolicitedFlag, error) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}

//TimeInForce is a non-required field for ExecutionReport.
func (m *ExecutionReport) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//EffectiveTime is a non-required field for ExecutionReport.
func (m *ExecutionReport) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//ExpireDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//ExpireTime is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//ExecInst is a non-required field for ExecutionReport.
func (m *ExecutionReport) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//Rule80A is a non-required field for ExecutionReport.
func (m *ExecutionReport) Rule80A() (*field.Rule80A, error) {
	f := new(field.Rule80A)
	err := m.Body.Get(f)
	return f, err
}

//LastShares is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastShares() (*field.LastShares, error) {
	f := new(field.LastShares)
	err := m.Body.Get(f)
	return f, err
}

//LastPx is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//LastSpotRate is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastSpotRate() (*field.LastSpotRate, error) {
	f := new(field.LastSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//LastForwardPoints is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastForwardPoints() (*field.LastForwardPoints, error) {
	f := new(field.LastForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//LastMkt is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for ExecutionReport.
func (m *ExecutionReport) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//LastCapacity is a non-required field for ExecutionReport.
func (m *ExecutionReport) LastCapacity() (*field.LastCapacity, error) {
	f := new(field.LastCapacity)
	err := m.Body.Get(f)
	return f, err
}

//LeavesQty is a required field for ExecutionReport.
func (m *ExecutionReport) LeavesQty() (*field.LeavesQty, error) {
	f := new(field.LeavesQty)
	err := m.Body.Get(f)
	return f, err
}

//CumQty is a required field for ExecutionReport.
func (m *ExecutionReport) CumQty() (*field.CumQty, error) {
	f := new(field.CumQty)
	err := m.Body.Get(f)
	return f, err
}

//AvgPx is a required field for ExecutionReport.
func (m *ExecutionReport) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//DayOrderQty is a non-required field for ExecutionReport.
func (m *ExecutionReport) DayOrderQty() (*field.DayOrderQty, error) {
	f := new(field.DayOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//DayCumQty is a non-required field for ExecutionReport.
func (m *ExecutionReport) DayCumQty() (*field.DayCumQty, error) {
	f := new(field.DayCumQty)
	err := m.Body.Get(f)
	return f, err
}

//DayAvgPx is a non-required field for ExecutionReport.
func (m *ExecutionReport) DayAvgPx() (*field.DayAvgPx, error) {
	f := new(field.DayAvgPx)
	err := m.Body.Get(f)
	return f, err
}

//GTBookingInst is a non-required field for ExecutionReport.
func (m *ExecutionReport) GTBookingInst() (*field.GTBookingInst, error) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a non-required field for ExecutionReport.
func (m *ExecutionReport) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for ExecutionReport.
func (m *ExecutionReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//ReportToExch is a non-required field for ExecutionReport.
func (m *ExecutionReport) ReportToExch() (*field.ReportToExch, error) {
	f := new(field.ReportToExch)
	err := m.Body.Get(f)
	return f, err
}

//Commission is a non-required field for ExecutionReport.
func (m *ExecutionReport) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//CommType is a non-required field for ExecutionReport.
func (m *ExecutionReport) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//GrossTradeAmt is a non-required field for ExecutionReport.
func (m *ExecutionReport) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrAmt is a non-required field for ExecutionReport.
func (m *ExecutionReport) SettlCurrAmt() (*field.SettlCurrAmt, error) {
	f := new(field.SettlCurrAmt)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrency is a non-required field for ExecutionReport.
func (m *ExecutionReport) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrFxRate is a non-required field for ExecutionReport.
func (m *ExecutionReport) SettlCurrFxRate() (*field.SettlCurrFxRate, error) {
	f := new(field.SettlCurrFxRate)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for ExecutionReport.
func (m *ExecutionReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, error) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//HandlInst is a non-required field for ExecutionReport.
func (m *ExecutionReport) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//MinQty is a non-required field for ExecutionReport.
func (m *ExecutionReport) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//MaxFloor is a non-required field for ExecutionReport.
func (m *ExecutionReport) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//OpenClose is a non-required field for ExecutionReport.
func (m *ExecutionReport) OpenClose() (*field.OpenClose, error) {
	f := new(field.OpenClose)
	err := m.Body.Get(f)
	return f, err
}

//MaxShow is a non-required field for ExecutionReport.
func (m *ExecutionReport) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for ExecutionReport.
func (m *ExecutionReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for ExecutionReport.
func (m *ExecutionReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for ExecutionReport.
func (m *ExecutionReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//FutSettDate2 is a non-required field for ExecutionReport.
func (m *ExecutionReport) FutSettDate2() (*field.FutSettDate2, error) {
	f := new(field.FutSettDate2)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty2 is a non-required field for ExecutionReport.
func (m *ExecutionReport) OrderQty2() (*field.OrderQty2, error) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//ClearingFirm is a non-required field for ExecutionReport.
func (m *ExecutionReport) ClearingFirm() (*field.ClearingFirm, error) {
	f := new(field.ClearingFirm)
	err := m.Body.Get(f)
	return f, err
}

//ClearingAccount is a non-required field for ExecutionReport.
func (m *ExecutionReport) ClearingAccount() (*field.ClearingAccount, error) {
	f := new(field.ClearingAccount)
	err := m.Body.Get(f)
	return f, err
}

//MultiLegReportingType is a non-required field for ExecutionReport.
func (m *ExecutionReport) MultiLegReportingType() (*field.MultiLegReportingType, error) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}
