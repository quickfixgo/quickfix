package fix44

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

//CreateTradeCaptureReportBuilder returns an initialized TradeCaptureReportBuilder with specified required fields.
func CreateTradeCaptureReportBuilder(
	tradereportid field.TradeReportID,
	previouslyreported field.PreviouslyReported,
	lastqty field.LastQty,
	lastpx field.LastPx,
	tradedate field.TradeDate,
	transacttime field.TransactTime,
	nosides field.NoSides) TradeCaptureReportBuilder {
	var builder TradeCaptureReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(tradereportid)
	builder.Body.Set(previouslyreported)
	builder.Body.Set(lastqty)
	builder.Body.Set(lastpx)
	builder.Body.Set(tradedate)
	builder.Body.Set(transacttime)
	builder.Body.Set(nosides)
	return builder
}

//TradeReportID is a required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportID() (field.TradeReportID, error) {
	var f field.TradeReportID
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportTransType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportTransType() (field.TradeReportTransType, error) {
	var f field.TradeReportTransType
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportType() (field.TradeReportType, error) {
	var f field.TradeReportType
	err := m.Body.Get(&f)
	return f, err
}

//TradeRequestID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeRequestID() (field.TradeRequestID, error) {
	var f field.TradeRequestID
	err := m.Body.Get(&f)
	return f, err
}

//TrdType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdType() (field.TrdType, error) {
	var f field.TrdType
	err := m.Body.Get(&f)
	return f, err
}

//TrdSubType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdSubType() (field.TrdSubType, error) {
	var f field.TrdSubType
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTrdType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTrdType() (field.SecondaryTrdType, error) {
	var f field.SecondaryTrdType
	err := m.Body.Get(&f)
	return f, err
}

//TransferReason is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TransferReason() (field.TransferReason, error) {
	var f field.TransferReason
	err := m.Body.Get(&f)
	return f, err
}

//ExecType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExecType() (field.ExecType, error) {
	var f field.ExecType
	err := m.Body.Get(&f)
	return f, err
}

//TotNumTradeReports is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TotNumTradeReports() (field.TotNumTradeReports, error) {
	var f field.TotNumTradeReports
	err := m.Body.Get(&f)
	return f, err
}

//LastRptRequested is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastRptRequested() (field.LastRptRequested, error) {
	var f field.LastRptRequested
	err := m.Body.Get(&f)
	return f, err
}

//UnsolicitedIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnsolicitedIndicator() (field.UnsolicitedIndicator, error) {
	var f field.UnsolicitedIndicator
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportRefID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportRefID() (field.TradeReportRefID, error) {
	var f field.TradeReportRefID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeReportRefID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTradeReportRefID() (field.SecondaryTradeReportRefID, error) {
	var f field.SecondaryTradeReportRefID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTradeReportID() (field.SecondaryTradeReportID, error) {
	var f field.SecondaryTradeReportID
	err := m.Body.Get(&f)
	return f, err
}

//TradeLinkID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeLinkID() (field.TradeLinkID, error) {
	var f field.TradeLinkID
	err := m.Body.Get(&f)
	return f, err
}

//TrdMatchID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdMatchID() (field.TrdMatchID, error) {
	var f field.TrdMatchID
	err := m.Body.Get(&f)
	return f, err
}

//ExecID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExecID() (field.ExecID, error) {
	var f field.ExecID
	err := m.Body.Get(&f)
	return f, err
}

//OrdStatus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrdStatus() (field.OrdStatus, error) {
	var f field.OrdStatus
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryExecID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryExecID() (field.SecondaryExecID, error) {
	var f field.SecondaryExecID
	err := m.Body.Get(&f)
	return f, err
}

//ExecRestatementReason is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExecRestatementReason() (field.ExecRestatementReason, error) {
	var f field.ExecRestatementReason
	err := m.Body.Get(&f)
	return f, err
}

//PreviouslyReported is a required field for TradeCaptureReport.
func (m TradeCaptureReport) PreviouslyReported() (field.PreviouslyReported, error) {
	var f field.PreviouslyReported
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrderPercent() (field.OrderPercent, error) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RoundingDirection() (field.RoundingDirection, error) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RoundingModulus() (field.RoundingModulus, error) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) QtyType() (field.QtyType, error) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldType() (field.YieldType, error) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Yield() (field.Yield, error) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldCalcDate() (field.YieldCalcDate, error) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldRedemptionDate() (field.YieldRedemptionDate, error) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldRedemptionPrice() (field.YieldRedemptionPrice, error) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, error) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTradingSessionID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingTradingSessionID() (field.UnderlyingTradingSessionID, error) {
	var f field.UnderlyingTradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingTradingSessionSubID() (field.UnderlyingTradingSessionSubID, error) {
	var f field.UnderlyingTradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//LastQty is a required field for TradeCaptureReport.
func (m TradeCaptureReport) LastQty() (field.LastQty, error) {
	var f field.LastQty
	err := m.Body.Get(&f)
	return f, err
}

//LastPx is a required field for TradeCaptureReport.
func (m TradeCaptureReport) LastPx() (field.LastPx, error) {
	var f field.LastPx
	err := m.Body.Get(&f)
	return f, err
}

//LastParPx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastParPx() (field.LastParPx, error) {
	var f field.LastParPx
	err := m.Body.Get(&f)
	return f, err
}

//LastSpotRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastSpotRate() (field.LastSpotRate, error) {
	var f field.LastSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//LastForwardPoints is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastForwardPoints() (field.LastForwardPoints, error) {
	var f field.LastForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastMkt() (field.LastMkt, error) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AvgPx() (field.AvgPx, error) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AvgPxIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AvgPxIndicator() (field.AvgPxIndicator, error) {
	var f field.AvgPxIndicator
	err := m.Body.Get(&f)
	return f, err
}

//NoPosAmt is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoPosAmt() (field.NoPosAmt, error) {
	var f field.NoPosAmt
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegReportingType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MultiLegReportingType() (field.MultiLegReportingType, error) {
	var f field.MultiLegReportingType
	err := m.Body.Get(&f)
	return f, err
}

//TradeLegRefID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeLegRefID() (field.TradeLegRefID, error) {
	var f field.TradeLegRefID
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for TradeCaptureReport.
func (m TradeCaptureReport) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, error) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//MatchStatus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MatchStatus() (field.MatchStatus, error) {
	var f field.MatchStatus
	err := m.Body.Get(&f)
	return f, err
}

//MatchType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MatchType() (field.MatchType, error) {
	var f field.MatchType
	err := m.Body.Get(&f)
	return f, err
}

//NoSides is a required field for TradeCaptureReport.
func (m TradeCaptureReport) NoSides() (field.NoSides, error) {
	var f field.NoSides
	err := m.Body.Get(&f)
	return f, err
}
