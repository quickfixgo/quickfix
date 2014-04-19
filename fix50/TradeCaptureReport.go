package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
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

//CreateTradeCaptureReportBuilder returns an initialized TradeCaptureReportBuilder with specified required fields.
func CreateTradeCaptureReportBuilder(
	lastqty field.LastQty,
	lastpx field.LastPx,
	tradedate field.TradeDate,
	nosides field.NoSides) TradeCaptureReportBuilder {
	var builder TradeCaptureReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(lastqty)
	builder.Body.Set(lastpx)
	builder.Body.Set(tradedate)
	builder.Body.Set(nosides)
	return builder
}

//TradeReportID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportID() (field.TradeReportID, errors.MessageRejectError) {
	var f field.TradeReportID
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportTransType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportTransType() (field.TradeReportTransType, errors.MessageRejectError) {
	var f field.TradeReportTransType
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportType() (field.TradeReportType, errors.MessageRejectError) {
	var f field.TradeReportType
	err := m.Body.Get(&f)
	return f, err
}

//TradeRequestID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeRequestID() (field.TradeRequestID, errors.MessageRejectError) {
	var f field.TradeRequestID
	err := m.Body.Get(&f)
	return f, err
}

//TrdType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdType() (field.TrdType, errors.MessageRejectError) {
	var f field.TrdType
	err := m.Body.Get(&f)
	return f, err
}

//TrdSubType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdSubType() (field.TrdSubType, errors.MessageRejectError) {
	var f field.TrdSubType
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTrdType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTrdType() (field.SecondaryTrdType, errors.MessageRejectError) {
	var f field.SecondaryTrdType
	err := m.Body.Get(&f)
	return f, err
}

//TransferReason is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TransferReason() (field.TransferReason, errors.MessageRejectError) {
	var f field.TransferReason
	err := m.Body.Get(&f)
	return f, err
}

//ExecType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExecType() (field.ExecType, errors.MessageRejectError) {
	var f field.ExecType
	err := m.Body.Get(&f)
	return f, err
}

//TotNumTradeReports is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TotNumTradeReports() (field.TotNumTradeReports, errors.MessageRejectError) {
	var f field.TotNumTradeReports
	err := m.Body.Get(&f)
	return f, err
}

//LastRptRequested is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastRptRequested() (field.LastRptRequested, errors.MessageRejectError) {
	var f field.LastRptRequested
	err := m.Body.Get(&f)
	return f, err
}

//UnsolicitedIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnsolicitedIndicator() (field.UnsolicitedIndicator, errors.MessageRejectError) {
	var f field.UnsolicitedIndicator
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SubscriptionRequestType() (field.SubscriptionRequestType, errors.MessageRejectError) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportRefID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportRefID() (field.TradeReportRefID, errors.MessageRejectError) {
	var f field.TradeReportRefID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeReportRefID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTradeReportRefID() (field.SecondaryTradeReportRefID, errors.MessageRejectError) {
	var f field.SecondaryTradeReportRefID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTradeReportID() (field.SecondaryTradeReportID, errors.MessageRejectError) {
	var f field.SecondaryTradeReportID
	err := m.Body.Get(&f)
	return f, err
}

//TradeLinkID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeLinkID() (field.TradeLinkID, errors.MessageRejectError) {
	var f field.TradeLinkID
	err := m.Body.Get(&f)
	return f, err
}

//TrdMatchID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdMatchID() (field.TrdMatchID, errors.MessageRejectError) {
	var f field.TrdMatchID
	err := m.Body.Get(&f)
	return f, err
}

//ExecID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExecID() (field.ExecID, errors.MessageRejectError) {
	var f field.ExecID
	err := m.Body.Get(&f)
	return f, err
}

//OrdStatus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrdStatus() (field.OrdStatus, errors.MessageRejectError) {
	var f field.OrdStatus
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryExecID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryExecID() (field.SecondaryExecID, errors.MessageRejectError) {
	var f field.SecondaryExecID
	err := m.Body.Get(&f)
	return f, err
}

//ExecRestatementReason is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExecRestatementReason() (field.ExecRestatementReason, errors.MessageRejectError) {
	var f field.ExecRestatementReason
	err := m.Body.Get(&f)
	return f, err
}

//PreviouslyReported is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PreviouslyReported() (field.PreviouslyReported, errors.MessageRejectError) {
	var f field.PreviouslyReported
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PriceType() (field.PriceType, errors.MessageRejectError) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecuritySubType() (field.SecuritySubType, errors.MessageRejectError) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikeCurrency() (field.StrikeCurrency, errors.MessageRejectError) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Pool() (field.Pool, errors.MessageRejectError) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ContractSettlMonth() (field.ContractSettlMonth, errors.MessageRejectError) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CPProgram() (field.CPProgram, errors.MessageRejectError) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CPRegType() (field.CPRegType, errors.MessageRejectError) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoEvents() (field.NoEvents, errors.MessageRejectError) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) DatedDate() (field.DatedDate, errors.MessageRejectError) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) InterestAccrualDate() (field.InterestAccrualDate, errors.MessageRejectError) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityStatus() (field.SecurityStatus, errors.MessageRejectError) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettleOnOpenFlag() (field.SettleOnOpenFlag, errors.MessageRejectError) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikeMultiplier() (field.StrikeMultiplier, errors.MessageRejectError) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikeValue() (field.StrikeValue, errors.MessageRejectError) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MinPriceIncrement() (field.MinPriceIncrement, errors.MessageRejectError) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PositionLimit() (field.PositionLimit, errors.MessageRejectError) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NTPositionLimit() (field.NTPositionLimit, errors.MessageRejectError) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoInstrumentParties() (field.NoInstrumentParties, errors.MessageRejectError) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnitOfMeasure() (field.UnitOfMeasure, errors.MessageRejectError) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TimeUnit() (field.TimeUnit, errors.MessageRejectError) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MaturityTime() (field.MaturityTime, errors.MessageRejectError) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementDesc() (field.AgreementDesc, errors.MessageRejectError) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementID() (field.AgreementID, errors.MessageRejectError) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementDate() (field.AgreementDate, errors.MessageRejectError) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementCurrency() (field.AgreementCurrency, errors.MessageRejectError) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TerminationType() (field.TerminationType, errors.MessageRejectError) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StartDate() (field.StartDate, errors.MessageRejectError) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EndDate() (field.EndDate, errors.MessageRejectError) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) DeliveryType() (field.DeliveryType, errors.MessageRejectError) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MarginRatio() (field.MarginRatio, errors.MessageRejectError) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrderQty() (field.OrderQty, errors.MessageRejectError) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CashOrderQty() (field.CashOrderQty, errors.MessageRejectError) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrderPercent() (field.OrderPercent, errors.MessageRejectError) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RoundingDirection() (field.RoundingDirection, errors.MessageRejectError) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RoundingModulus() (field.RoundingModulus, errors.MessageRejectError) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) QtyType() (field.QtyType, errors.MessageRejectError) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldType() (field.YieldType, errors.MessageRejectError) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Yield() (field.Yield, errors.MessageRejectError) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldCalcDate() (field.YieldCalcDate, errors.MessageRejectError) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldRedemptionDate() (field.YieldRedemptionDate, errors.MessageRejectError) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldRedemptionPrice() (field.YieldRedemptionPrice, errors.MessageRejectError) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, errors.MessageRejectError) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTradingSessionID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingTradingSessionID() (field.UnderlyingTradingSessionID, errors.MessageRejectError) {
	var f field.UnderlyingTradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingTradingSessionSubID() (field.UnderlyingTradingSessionSubID, errors.MessageRejectError) {
	var f field.UnderlyingTradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//LastQty is a required field for TradeCaptureReport.
func (m TradeCaptureReport) LastQty() (field.LastQty, errors.MessageRejectError) {
	var f field.LastQty
	err := m.Body.Get(&f)
	return f, err
}

//LastPx is a required field for TradeCaptureReport.
func (m TradeCaptureReport) LastPx() (field.LastPx, errors.MessageRejectError) {
	var f field.LastPx
	err := m.Body.Get(&f)
	return f, err
}

//LastParPx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastParPx() (field.LastParPx, errors.MessageRejectError) {
	var f field.LastParPx
	err := m.Body.Get(&f)
	return f, err
}

//LastSpotRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastSpotRate() (field.LastSpotRate, errors.MessageRejectError) {
	var f field.LastSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//LastForwardPoints is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastForwardPoints() (field.LastForwardPoints, errors.MessageRejectError) {
	var f field.LastForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastMkt() (field.LastMkt, errors.MessageRejectError) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeDate() (field.TradeDate, errors.MessageRejectError) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ClearingBusinessDate() (field.ClearingBusinessDate, errors.MessageRejectError) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AvgPx() (field.AvgPx, errors.MessageRejectError) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Spread() (field.Spread, errors.MessageRejectError) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkCurveName() (field.BenchmarkCurveName, errors.MessageRejectError) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, errors.MessageRejectError) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkPrice() (field.BenchmarkPrice, errors.MessageRejectError) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkPriceType() (field.BenchmarkPriceType, errors.MessageRejectError) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkSecurityID() (field.BenchmarkSecurityID, errors.MessageRejectError) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AvgPxIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AvgPxIndicator() (field.AvgPxIndicator, errors.MessageRejectError) {
	var f field.AvgPxIndicator
	err := m.Body.Get(&f)
	return f, err
}

//NoPosAmt is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoPosAmt() (field.NoPosAmt, errors.MessageRejectError) {
	var f field.NoPosAmt
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegReportingType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MultiLegReportingType() (field.MultiLegReportingType, errors.MessageRejectError) {
	var f field.MultiLegReportingType
	err := m.Body.Get(&f)
	return f, err
}

//TradeLegRefID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeLegRefID() (field.TradeLegRefID, errors.MessageRejectError) {
	var f field.TradeLegRefID
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoLegs() (field.NoLegs, errors.MessageRejectError) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, errors.MessageRejectError) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlType() (field.SettlType, errors.MessageRejectError) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlDate() (field.SettlDate, errors.MessageRejectError) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//MatchStatus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MatchStatus() (field.MatchStatus, errors.MessageRejectError) {
	var f field.MatchStatus
	err := m.Body.Get(&f)
	return f, err
}

//MatchType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MatchType() (field.MatchType, errors.MessageRejectError) {
	var f field.MatchType
	err := m.Body.Get(&f)
	return f, err
}

//NoSides is a required field for TradeCaptureReport.
func (m TradeCaptureReport) NoSides() (field.NoSides, errors.MessageRejectError) {
	var f field.NoSides
	err := m.Body.Get(&f)
	return f, err
}

//CopyMsgIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CopyMsgIndicator() (field.CopyMsgIndicator, errors.MessageRejectError) {
	var f field.CopyMsgIndicator
	err := m.Body.Get(&f)
	return f, err
}

//PublishTrdIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PublishTrdIndicator() (field.PublishTrdIndicator, errors.MessageRejectError) {
	var f field.PublishTrdIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ShortSaleReason is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ShortSaleReason() (field.ShortSaleReason, errors.MessageRejectError) {
	var f field.ShortSaleReason
	err := m.Body.Get(&f)
	return f, err
}

//TrdRptStatus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdRptStatus() (field.TrdRptStatus, errors.MessageRejectError) {
	var f field.TrdRptStatus
	err := m.Body.Get(&f)
	return f, err
}

//AsOfIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AsOfIndicator() (field.AsOfIndicator, errors.MessageRejectError) {
	var f field.AsOfIndicator
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlSessID() (field.SettlSessID, errors.MessageRejectError) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlSessSubID() (field.SettlSessSubID, errors.MessageRejectError) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//TierCode is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TierCode() (field.TierCode, errors.MessageRejectError) {
	var f field.TierCode
	err := m.Body.Get(&f)
	return f, err
}

//MessageEventSource is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MessageEventSource() (field.MessageEventSource, errors.MessageRejectError) {
	var f field.MessageEventSource
	err := m.Body.Get(&f)
	return f, err
}

//LastUpdateTime is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastUpdateTime() (field.LastUpdateTime, errors.MessageRejectError) {
	var f field.LastUpdateTime
	err := m.Body.Get(&f)
	return f, err
}

//RndPx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RndPx() (field.RndPx, errors.MessageRejectError) {
	var f field.RndPx
	err := m.Body.Get(&f)
	return f, err
}

//TradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeID() (field.TradeID, errors.MessageRejectError) {
	var f field.TradeID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTradeID() (field.SecondaryTradeID, errors.MessageRejectError) {
	var f field.SecondaryTradeID
	err := m.Body.Get(&f)
	return f, err
}

//FirmTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FirmTradeID() (field.FirmTradeID, errors.MessageRejectError) {
	var f field.FirmTradeID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryFirmTradeID() (field.SecondaryFirmTradeID, errors.MessageRejectError) {
	var f field.SecondaryFirmTradeID
	err := m.Body.Get(&f)
	return f, err
}

//CalculatedCcyLastQty is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CalculatedCcyLastQty() (field.CalculatedCcyLastQty, errors.MessageRejectError) {
	var f field.CalculatedCcyLastQty
	err := m.Body.Get(&f)
	return f, err
}

//LastSwapPoints is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastSwapPoints() (field.LastSwapPoints, errors.MessageRejectError) {
	var f field.LastSwapPoints
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlementDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingSettlementDate() (field.UnderlyingSettlementDate, errors.MessageRejectError) {
	var f field.UnderlyingSettlementDate
	err := m.Body.Get(&f)
	return f, err
}

//GrossTradeAmt is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) GrossTradeAmt() (field.GrossTradeAmt, errors.MessageRejectError) {
	var f field.GrossTradeAmt
	err := m.Body.Get(&f)
	return f, err
}

//NoRootPartyIDs is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoRootPartyIDs() (field.NoRootPartyIDs, errors.MessageRejectError) {
	var f field.NoRootPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//OrderCategory is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrderCategory() (field.OrderCategory, errors.MessageRejectError) {
	var f field.OrderCategory
	err := m.Body.Get(&f)
	return f, err
}

//TradeHandlingInstr is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeHandlingInstr() (field.TradeHandlingInstr, errors.MessageRejectError) {
	var f field.TradeHandlingInstr
	err := m.Body.Get(&f)
	return f, err
}

//OrigTradeHandlingInstr is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrigTradeHandlingInstr() (field.OrigTradeHandlingInstr, errors.MessageRejectError) {
	var f field.OrigTradeHandlingInstr
	err := m.Body.Get(&f)
	return f, err
}

//OrigTradeDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrigTradeDate() (field.OrigTradeDate, errors.MessageRejectError) {
	var f field.OrigTradeDate
	err := m.Body.Get(&f)
	return f, err
}

//OrigTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrigTradeID() (field.OrigTradeID, errors.MessageRejectError) {
	var f field.OrigTradeID
	err := m.Body.Get(&f)
	return f, err
}

//OrigSecondaryTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrigSecondaryTradeID() (field.OrigSecondaryTradeID, errors.MessageRejectError) {
	var f field.OrigSecondaryTradeID
	err := m.Body.Get(&f)
	return f, err
}

//TZTransactTime is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TZTransactTime() (field.TZTransactTime, errors.MessageRejectError) {
	var f field.TZTransactTime
	err := m.Body.Get(&f)
	return f, err
}

//ReportedPxDiff is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ReportedPxDiff() (field.ReportedPxDiff, errors.MessageRejectError) {
	var f field.ReportedPxDiff
	err := m.Body.Get(&f)
	return f, err
}
