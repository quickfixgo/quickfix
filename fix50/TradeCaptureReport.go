package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradeCaptureReport msg type = AE.
type TradeCaptureReport struct {
	message.Message
}

//TradeCaptureReportBuilder builds TradeCaptureReport messages.
type TradeCaptureReportBuilder struct {
	message.MessageBuilder
}

//NewTradeCaptureReportBuilder returns an initialized TradeCaptureReportBuilder with specified required fields.
func NewTradeCaptureReportBuilder(
	lastqty field.LastQty,
	lastpx field.LastPx,
	tradedate field.TradeDate,
	nosides field.NoSides) *TradeCaptureReportBuilder {
	builder := new(TradeCaptureReportBuilder)
	builder.Body.Set(lastqty)
	builder.Body.Set(lastpx)
	builder.Body.Set(tradedate)
	builder.Body.Set(nosides)
	return builder
}

//TradeReportID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeReportID() (*field.TradeReportID, error) {
	f := new(field.TradeReportID)
	err := m.Body.Get(f)
	return f, err
}

//TradeReportTransType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeReportTransType() (*field.TradeReportTransType, error) {
	f := new(field.TradeReportTransType)
	err := m.Body.Get(f)
	return f, err
}

//TradeReportType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeReportType() (*field.TradeReportType, error) {
	f := new(field.TradeReportType)
	err := m.Body.Get(f)
	return f, err
}

//TradeRequestID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeRequestID() (*field.TradeRequestID, error) {
	f := new(field.TradeRequestID)
	err := m.Body.Get(f)
	return f, err
}

//TrdType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TrdType() (*field.TrdType, error) {
	f := new(field.TrdType)
	err := m.Body.Get(f)
	return f, err
}

//TrdSubType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TrdSubType() (*field.TrdSubType, error) {
	f := new(field.TrdSubType)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryTrdType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecondaryTrdType() (*field.SecondaryTrdType, error) {
	f := new(field.SecondaryTrdType)
	err := m.Body.Get(f)
	return f, err
}

//TransferReason is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TransferReason() (*field.TransferReason, error) {
	f := new(field.TransferReason)
	err := m.Body.Get(f)
	return f, err
}

//ExecType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) ExecType() (*field.ExecType, error) {
	f := new(field.ExecType)
	err := m.Body.Get(f)
	return f, err
}

//TotNumTradeReports is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TotNumTradeReports() (*field.TotNumTradeReports, error) {
	f := new(field.TotNumTradeReports)
	err := m.Body.Get(f)
	return f, err
}

//LastRptRequested is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastRptRequested() (*field.LastRptRequested, error) {
	f := new(field.LastRptRequested)
	err := m.Body.Get(f)
	return f, err
}

//UnsolicitedIndicator is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) UnsolicitedIndicator() (*field.UnsolicitedIndicator, error) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//TradeReportRefID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeReportRefID() (*field.TradeReportRefID, error) {
	f := new(field.TradeReportRefID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryTradeReportRefID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecondaryTradeReportRefID() (*field.SecondaryTradeReportRefID, error) {
	f := new(field.SecondaryTradeReportRefID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecondaryTradeReportID() (*field.SecondaryTradeReportID, error) {
	f := new(field.SecondaryTradeReportID)
	err := m.Body.Get(f)
	return f, err
}

//TradeLinkID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeLinkID() (*field.TradeLinkID, error) {
	f := new(field.TradeLinkID)
	err := m.Body.Get(f)
	return f, err
}

//TrdMatchID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TrdMatchID() (*field.TrdMatchID, error) {
	f := new(field.TrdMatchID)
	err := m.Body.Get(f)
	return f, err
}

//ExecID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//OrdStatus is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) OrdStatus() (*field.OrdStatus, error) {
	f := new(field.OrdStatus)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryExecID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecondaryExecID() (*field.SecondaryExecID, error) {
	f := new(field.SecondaryExecID)
	err := m.Body.Get(f)
	return f, err
}

//ExecRestatementReason is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) ExecRestatementReason() (*field.ExecRestatementReason, error) {
	f := new(field.ExecRestatementReason)
	err := m.Body.Get(f)
	return f, err
}

//PreviouslyReported is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) PreviouslyReported() (*field.PreviouslyReported, error) {
	f := new(field.PreviouslyReported)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDesc is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//AgreementID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementCurrency is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//TerminationType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//StartDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//EndDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//MarginRatio is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//CashOrderQty is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//OrderPercent is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) OrderPercent() (*field.OrderPercent, error) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//RoundingDirection is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RoundingDirection() (*field.RoundingDirection, error) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//RoundingModulus is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RoundingModulus() (*field.RoundingModulus, error) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//QtyType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//YieldType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//Yield is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//YieldCalcDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) YieldCalcDate() (*field.YieldCalcDate, error) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) YieldRedemptionDate() (*field.YieldRedemptionDate, error) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) YieldRedemptionPrice() (*field.YieldRedemptionPrice, error) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, error) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingTradingSessionID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) UnderlyingTradingSessionID() (*field.UnderlyingTradingSessionID, error) {
	f := new(field.UnderlyingTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) UnderlyingTradingSessionSubID() (*field.UnderlyingTradingSessionSubID, error) {
	f := new(field.UnderlyingTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//LastQty is a required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastQty() (*field.LastQty, error) {
	f := new(field.LastQty)
	err := m.Body.Get(f)
	return f, err
}

//LastPx is a required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//LastParPx is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastParPx() (*field.LastParPx, error) {
	f := new(field.LastParPx)
	err := m.Body.Get(f)
	return f, err
}

//LastSpotRate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastSpotRate() (*field.LastSpotRate, error) {
	f := new(field.LastSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//LastForwardPoints is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastForwardPoints() (*field.LastForwardPoints, error) {
	f := new(field.LastForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//LastMkt is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//AvgPx is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//Spread is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveName is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPrice is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPriceType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//AvgPxIndicator is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) AvgPxIndicator() (*field.AvgPxIndicator, error) {
	f := new(field.AvgPxIndicator)
	err := m.Body.Get(f)
	return f, err
}

//NoPosAmt is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) NoPosAmt() (*field.NoPosAmt, error) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}

//MultiLegReportingType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MultiLegReportingType() (*field.MultiLegReportingType, error) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}

//TradeLegRefID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeLegRefID() (*field.TradeLegRefID, error) {
	f := new(field.TradeLegRefID)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, error) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}

//SettlType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//MatchStatus is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//MatchType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MatchType() (*field.MatchType, error) {
	f := new(field.MatchType)
	err := m.Body.Get(f)
	return f, err
}

//NoSides is a required field for TradeCaptureReport.
func (m *TradeCaptureReport) NoSides() (*field.NoSides, error) {
	f := new(field.NoSides)
	err := m.Body.Get(f)
	return f, err
}

//CopyMsgIndicator is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CopyMsgIndicator() (*field.CopyMsgIndicator, error) {
	f := new(field.CopyMsgIndicator)
	err := m.Body.Get(f)
	return f, err
}

//PublishTrdIndicator is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) PublishTrdIndicator() (*field.PublishTrdIndicator, error) {
	f := new(field.PublishTrdIndicator)
	err := m.Body.Get(f)
	return f, err
}

//ShortSaleReason is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) ShortSaleReason() (*field.ShortSaleReason, error) {
	f := new(field.ShortSaleReason)
	err := m.Body.Get(f)
	return f, err
}

//TrdRptStatus is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TrdRptStatus() (*field.TrdRptStatus, error) {
	f := new(field.TrdRptStatus)
	err := m.Body.Get(f)
	return f, err
}

//AsOfIndicator is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) AsOfIndicator() (*field.AsOfIndicator, error) {
	f := new(field.AsOfIndicator)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessSubID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//TierCode is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TierCode() (*field.TierCode, error) {
	f := new(field.TierCode)
	err := m.Body.Get(f)
	return f, err
}

//MessageEventSource is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MessageEventSource() (*field.MessageEventSource, error) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}

//LastUpdateTime is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastUpdateTime() (*field.LastUpdateTime, error) {
	f := new(field.LastUpdateTime)
	err := m.Body.Get(f)
	return f, err
}

//RndPx is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RndPx() (*field.RndPx, error) {
	f := new(field.RndPx)
	err := m.Body.Get(f)
	return f, err
}

//TradeID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeID() (*field.TradeID, error) {
	f := new(field.TradeID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryTradeID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecondaryTradeID() (*field.SecondaryTradeID, error) {
	f := new(field.SecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}

//FirmTradeID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) FirmTradeID() (*field.FirmTradeID, error) {
	f := new(field.FirmTradeID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecondaryFirmTradeID() (*field.SecondaryFirmTradeID, error) {
	f := new(field.SecondaryFirmTradeID)
	err := m.Body.Get(f)
	return f, err
}

//CalculatedCcyLastQty is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CalculatedCcyLastQty() (*field.CalculatedCcyLastQty, error) {
	f := new(field.CalculatedCcyLastQty)
	err := m.Body.Get(f)
	return f, err
}

//LastSwapPoints is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastSwapPoints() (*field.LastSwapPoints, error) {
	f := new(field.LastSwapPoints)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSettlementDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) UnderlyingSettlementDate() (*field.UnderlyingSettlementDate, error) {
	f := new(field.UnderlyingSettlementDate)
	err := m.Body.Get(f)
	return f, err
}

//GrossTradeAmt is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//NoRootPartyIDs is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) NoRootPartyIDs() (*field.NoRootPartyIDs, error) {
	f := new(field.NoRootPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//OrderCategory is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) OrderCategory() (*field.OrderCategory, error) {
	f := new(field.OrderCategory)
	err := m.Body.Get(f)
	return f, err
}

//TradeHandlingInstr is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeHandlingInstr() (*field.TradeHandlingInstr, error) {
	f := new(field.TradeHandlingInstr)
	err := m.Body.Get(f)
	return f, err
}

//OrigTradeHandlingInstr is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) OrigTradeHandlingInstr() (*field.OrigTradeHandlingInstr, error) {
	f := new(field.OrigTradeHandlingInstr)
	err := m.Body.Get(f)
	return f, err
}

//OrigTradeDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) OrigTradeDate() (*field.OrigTradeDate, error) {
	f := new(field.OrigTradeDate)
	err := m.Body.Get(f)
	return f, err
}

//OrigTradeID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) OrigTradeID() (*field.OrigTradeID, error) {
	f := new(field.OrigTradeID)
	err := m.Body.Get(f)
	return f, err
}

//OrigSecondaryTradeID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) OrigSecondaryTradeID() (*field.OrigSecondaryTradeID, error) {
	f := new(field.OrigSecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}

//TZTransactTime is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TZTransactTime() (*field.TZTransactTime, error) {
	f := new(field.TZTransactTime)
	err := m.Body.Get(f)
	return f, err
}

//ReportedPxDiff is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) ReportedPxDiff() (*field.ReportedPxDiff, error) {
	f := new(field.ReportedPxDiff)
	err := m.Body.Get(f)
	return f, err
}
