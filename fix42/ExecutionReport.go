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

//CreateExecutionReportBuilder returns an initialized ExecutionReportBuilder with specified required fields.
func CreateExecutionReportBuilder(
	orderid field.OrderID,
	execid field.ExecID,
	exectranstype field.ExecTransType,
	exectype field.ExecType,
	ordstatus field.OrdStatus,
	symbol field.Symbol,
	side field.Side,
	leavesqty field.LeavesQty,
	cumqty field.CumQty,
	avgpx field.AvgPx) ExecutionReportBuilder {
	var builder ExecutionReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
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
func (m ExecutionReport) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecondaryOrderID() (field.SecondaryOrderID, error) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrigClOrdID is a non-required field for ExecutionReport.
func (m ExecutionReport) OrigClOrdID() (field.OrigClOrdID, error) {
	var f field.OrigClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClientID is a non-required field for ExecutionReport.
func (m ExecutionReport) ClientID() (field.ClientID, error) {
	var f field.ClientID
	err := m.Body.Get(&f)
	return f, err
}

//ExecBroker is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecBroker() (field.ExecBroker, error) {
	var f field.ExecBroker
	err := m.Body.Get(&f)
	return f, err
}

//NoContraBrokers is a non-required field for ExecutionReport.
func (m ExecutionReport) NoContraBrokers() (field.NoContraBrokers, error) {
	var f field.NoContraBrokers
	err := m.Body.Get(&f)
	return f, err
}

//ListID is a non-required field for ExecutionReport.
func (m ExecutionReport) ListID() (field.ListID, error) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//ExecID is a required field for ExecutionReport.
func (m ExecutionReport) ExecID() (field.ExecID, error) {
	var f field.ExecID
	err := m.Body.Get(&f)
	return f, err
}

//ExecTransType is a required field for ExecutionReport.
func (m ExecutionReport) ExecTransType() (field.ExecTransType, error) {
	var f field.ExecTransType
	err := m.Body.Get(&f)
	return f, err
}

//ExecRefID is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRefID() (field.ExecRefID, error) {
	var f field.ExecRefID
	err := m.Body.Get(&f)
	return f, err
}

//ExecType is a required field for ExecutionReport.
func (m ExecutionReport) ExecType() (field.ExecType, error) {
	var f field.ExecType
	err := m.Body.Get(&f)
	return f, err
}

//OrdStatus is a required field for ExecutionReport.
func (m ExecutionReport) OrdStatus() (field.OrdStatus, error) {
	var f field.OrdStatus
	err := m.Body.Get(&f)
	return f, err
}

//OrdRejReason is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdRejReason() (field.OrdRejReason, error) {
	var f field.OrdRejReason
	err := m.Body.Get(&f)
	return f, err
}

//ExecRestatementReason is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecRestatementReason() (field.ExecRestatementReason, error) {
	var f field.ExecRestatementReason
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for ExecutionReport.
func (m ExecutionReport) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//SettlmntTyp is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlmntTyp() (field.SettlmntTyp, error) {
	var f field.SettlmntTyp
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate is a non-required field for ExecutionReport.
func (m ExecutionReport) FutSettDate() (field.FutSettDate, error) {
	var f field.FutSettDate
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a required field for ExecutionReport.
func (m ExecutionReport) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for ExecutionReport.
func (m ExecutionReport) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//IDSource is a non-required field for ExecutionReport.
func (m ExecutionReport) IDSource() (field.IDSource, error) {
	var f field.IDSource
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDay is a non-required field for ExecutionReport.
func (m ExecutionReport) MaturityDay() (field.MaturityDay, error) {
	var f field.MaturityDay
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for ExecutionReport.
func (m ExecutionReport) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for ExecutionReport.
func (m ExecutionReport) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for ExecutionReport.
func (m ExecutionReport) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for ExecutionReport.
func (m ExecutionReport) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for ExecutionReport.
func (m ExecutionReport) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for ExecutionReport.
func (m ExecutionReport) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for ExecutionReport.
func (m ExecutionReport) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a non-required field for ExecutionReport.
func (m ExecutionReport) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for ExecutionReport.
func (m ExecutionReport) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//StopPx is a non-required field for ExecutionReport.
func (m ExecutionReport) StopPx() (field.StopPx, error) {
	var f field.StopPx
	err := m.Body.Get(&f)
	return f, err
}

//PegDifference is a non-required field for ExecutionReport.
func (m ExecutionReport) PegDifference() (field.PegDifference, error) {
	var f field.PegDifference
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionInst() (field.DiscretionInst, error) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffset is a non-required field for ExecutionReport.
func (m ExecutionReport) DiscretionOffset() (field.DiscretionOffset, error) {
	var f field.DiscretionOffset
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for ExecutionReport.
func (m ExecutionReport) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ComplianceID is a non-required field for ExecutionReport.
func (m ExecutionReport) ComplianceID() (field.ComplianceID, error) {
	var f field.ComplianceID
	err := m.Body.Get(&f)
	return f, err
}

//SolicitedFlag is a non-required field for ExecutionReport.
func (m ExecutionReport) SolicitedFlag() (field.SolicitedFlag, error) {
	var f field.SolicitedFlag
	err := m.Body.Get(&f)
	return f, err
}

//TimeInForce is a non-required field for ExecutionReport.
func (m ExecutionReport) TimeInForce() (field.TimeInForce, error) {
	var f field.TimeInForce
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for ExecutionReport.
func (m ExecutionReport) EffectiveTime() (field.EffectiveTime, error) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireDate is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireDate() (field.ExpireDate, error) {
	var f field.ExpireDate
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for ExecutionReport.
func (m ExecutionReport) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//ExecInst is a non-required field for ExecutionReport.
func (m ExecutionReport) ExecInst() (field.ExecInst, error) {
	var f field.ExecInst
	err := m.Body.Get(&f)
	return f, err
}

//Rule80A is a non-required field for ExecutionReport.
func (m ExecutionReport) Rule80A() (field.Rule80A, error) {
	var f field.Rule80A
	err := m.Body.Get(&f)
	return f, err
}

//LastShares is a non-required field for ExecutionReport.
func (m ExecutionReport) LastShares() (field.LastShares, error) {
	var f field.LastShares
	err := m.Body.Get(&f)
	return f, err
}

//LastPx is a non-required field for ExecutionReport.
func (m ExecutionReport) LastPx() (field.LastPx, error) {
	var f field.LastPx
	err := m.Body.Get(&f)
	return f, err
}

//LastSpotRate is a non-required field for ExecutionReport.
func (m ExecutionReport) LastSpotRate() (field.LastSpotRate, error) {
	var f field.LastSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//LastForwardPoints is a non-required field for ExecutionReport.
func (m ExecutionReport) LastForwardPoints() (field.LastForwardPoints, error) {
	var f field.LastForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for ExecutionReport.
func (m ExecutionReport) LastMkt() (field.LastMkt, error) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for ExecutionReport.
func (m ExecutionReport) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//LastCapacity is a non-required field for ExecutionReport.
func (m ExecutionReport) LastCapacity() (field.LastCapacity, error) {
	var f field.LastCapacity
	err := m.Body.Get(&f)
	return f, err
}

//LeavesQty is a required field for ExecutionReport.
func (m ExecutionReport) LeavesQty() (field.LeavesQty, error) {
	var f field.LeavesQty
	err := m.Body.Get(&f)
	return f, err
}

//CumQty is a required field for ExecutionReport.
func (m ExecutionReport) CumQty() (field.CumQty, error) {
	var f field.CumQty
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a required field for ExecutionReport.
func (m ExecutionReport) AvgPx() (field.AvgPx, error) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//DayOrderQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayOrderQty() (field.DayOrderQty, error) {
	var f field.DayOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//DayCumQty is a non-required field for ExecutionReport.
func (m ExecutionReport) DayCumQty() (field.DayCumQty, error) {
	var f field.DayCumQty
	err := m.Body.Get(&f)
	return f, err
}

//DayAvgPx is a non-required field for ExecutionReport.
func (m ExecutionReport) DayAvgPx() (field.DayAvgPx, error) {
	var f field.DayAvgPx
	err := m.Body.Get(&f)
	return f, err
}

//GTBookingInst is a non-required field for ExecutionReport.
func (m ExecutionReport) GTBookingInst() (field.GTBookingInst, error) {
	var f field.GTBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for ExecutionReport.
func (m ExecutionReport) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for ExecutionReport.
func (m ExecutionReport) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//ReportToExch is a non-required field for ExecutionReport.
func (m ExecutionReport) ReportToExch() (field.ReportToExch, error) {
	var f field.ReportToExch
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for ExecutionReport.
func (m ExecutionReport) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for ExecutionReport.
func (m ExecutionReport) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//GrossTradeAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) GrossTradeAmt() (field.GrossTradeAmt, error) {
	var f field.GrossTradeAmt
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrAmt is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrAmt() (field.SettlCurrAmt, error) {
	var f field.SettlCurrAmt
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrency() (field.SettlCurrency, error) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRate is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRate() (field.SettlCurrFxRate, error) {
	var f field.SettlCurrFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for ExecutionReport.
func (m ExecutionReport) SettlCurrFxRateCalc() (field.SettlCurrFxRateCalc, error) {
	var f field.SettlCurrFxRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//HandlInst is a non-required field for ExecutionReport.
func (m ExecutionReport) HandlInst() (field.HandlInst, error) {
	var f field.HandlInst
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for ExecutionReport.
func (m ExecutionReport) MinQty() (field.MinQty, error) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//MaxFloor is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxFloor() (field.MaxFloor, error) {
	var f field.MaxFloor
	err := m.Body.Get(&f)
	return f, err
}

//OpenClose is a non-required field for ExecutionReport.
func (m ExecutionReport) OpenClose() (field.OpenClose, error) {
	var f field.OpenClose
	err := m.Body.Get(&f)
	return f, err
}

//MaxShow is a non-required field for ExecutionReport.
func (m ExecutionReport) MaxShow() (field.MaxShow, error) {
	var f field.MaxShow
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for ExecutionReport.
func (m ExecutionReport) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for ExecutionReport.
func (m ExecutionReport) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate2 is a non-required field for ExecutionReport.
func (m ExecutionReport) FutSettDate2() (field.FutSettDate2, error) {
	var f field.FutSettDate2
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty2 is a non-required field for ExecutionReport.
func (m ExecutionReport) OrderQty2() (field.OrderQty2, error) {
	var f field.OrderQty2
	err := m.Body.Get(&f)
	return f, err
}

//ClearingFirm is a non-required field for ExecutionReport.
func (m ExecutionReport) ClearingFirm() (field.ClearingFirm, error) {
	var f field.ClearingFirm
	err := m.Body.Get(&f)
	return f, err
}

//ClearingAccount is a non-required field for ExecutionReport.
func (m ExecutionReport) ClearingAccount() (field.ClearingAccount, error) {
	var f field.ClearingAccount
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegReportingType is a non-required field for ExecutionReport.
func (m ExecutionReport) MultiLegReportingType() (field.MultiLegReportingType, error) {
	var f field.MultiLegReportingType
	err := m.Body.Get(&f)
	return f, err
}
