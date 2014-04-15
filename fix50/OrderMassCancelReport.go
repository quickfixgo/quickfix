package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderMassCancelReport msg type = r.
type OrderMassCancelReport struct {
	message.Message
}

//OrderMassCancelReportBuilder builds OrderMassCancelReport messages.
type OrderMassCancelReportBuilder struct {
	message.MessageBuilder
}

//CreateOrderMassCancelReportBuilder returns an initialized OrderMassCancelReportBuilder with specified required fields.
func CreateOrderMassCancelReportBuilder(
	orderid field.OrderID,
	masscancelrequesttype field.MassCancelRequestType,
	masscancelresponse field.MassCancelResponse) OrderMassCancelReportBuilder {
	var builder OrderMassCancelReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(orderid)
	builder.Body.Set(masscancelrequesttype)
	builder.Body.Set(masscancelresponse)
	return builder
}

//ClOrdID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecondaryClOrdID() (field.SecondaryClOrdID, error) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a required field for OrderMassCancelReport.
func (m OrderMassCancelReport) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecondaryOrderID() (field.SecondaryOrderID, error) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//MassCancelRequestType is a required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MassCancelRequestType() (field.MassCancelRequestType, error) {
	var f field.MassCancelRequestType
	err := m.Body.Get(&f)
	return f, err
}

//MassCancelResponse is a required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MassCancelResponse() (field.MassCancelResponse, error) {
	var f field.MassCancelResponse
	err := m.Body.Get(&f)
	return f, err
}

//MassCancelRejectReason is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MassCancelRejectReason() (field.MassCancelRejectReason, error) {
	var f field.MassCancelRejectReason
	err := m.Body.Get(&f)
	return f, err
}

//TotalAffectedOrders is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TotalAffectedOrders() (field.TotalAffectedOrders, error) {
	var f field.TotalAffectedOrders
	err := m.Body.Get(&f)
	return f, err
}

//NoAffectedOrders is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoAffectedOrders() (field.NoAffectedOrders, error) {
	var f field.NoAffectedOrders
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbol is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSymbol() (field.UnderlyingSymbol, error) {
	var f field.UnderlyingSymbol
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbolSfx is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSymbolSfx() (field.UnderlyingSymbolSfx, error) {
	var f field.UnderlyingSymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityID() (field.UnderlyingSecurityID, error) {
	var f field.UnderlyingSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityIDSource is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityIDSource() (field.UnderlyingSecurityIDSource, error) {
	var f field.UnderlyingSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingSecurityAltID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoUnderlyingSecurityAltID() (field.NoUnderlyingSecurityAltID, error) {
	var f field.NoUnderlyingSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingProduct is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingProduct() (field.UnderlyingProduct, error) {
	var f field.UnderlyingProduct
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCFICode is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCFICode() (field.UnderlyingCFICode, error) {
	var f field.UnderlyingCFICode
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityType() (field.UnderlyingSecurityType, error) {
	var f field.UnderlyingSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecuritySubType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecuritySubType() (field.UnderlyingSecuritySubType, error) {
	var f field.UnderlyingSecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityMonthYear is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingMaturityMonthYear() (field.UnderlyingMaturityMonthYear, error) {
	var f field.UnderlyingMaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingMaturityDate() (field.UnderlyingMaturityDate, error) {
	var f field.UnderlyingMaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponPaymentDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCouponPaymentDate() (field.UnderlyingCouponPaymentDate, error) {
	var f field.UnderlyingCouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssueDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingIssueDate() (field.UnderlyingIssueDate, error) {
	var f field.UnderlyingIssueDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepoCollateralSecurityType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRepoCollateralSecurityType() (field.UnderlyingRepoCollateralSecurityType, error) {
	var f field.UnderlyingRepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseTerm is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRepurchaseTerm() (field.UnderlyingRepurchaseTerm, error) {
	var f field.UnderlyingRepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRepurchaseRate() (field.UnderlyingRepurchaseRate, error) {
	var f field.UnderlyingRepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFactor is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingFactor() (field.UnderlyingFactor, error) {
	var f field.UnderlyingFactor
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCreditRating is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCreditRating() (field.UnderlyingCreditRating, error) {
	var f field.UnderlyingCreditRating
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingInstrRegistry is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingInstrRegistry() (field.UnderlyingInstrRegistry, error) {
	var f field.UnderlyingInstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCountryOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCountryOfIssue() (field.UnderlyingCountryOfIssue, error) {
	var f field.UnderlyingCountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingStateOrProvinceOfIssue() (field.UnderlyingStateOrProvinceOfIssue, error) {
	var f field.UnderlyingStateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingLocaleOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingLocaleOfIssue() (field.UnderlyingLocaleOfIssue, error) {
	var f field.UnderlyingLocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRedemptionDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRedemptionDate() (field.UnderlyingRedemptionDate, error) {
	var f field.UnderlyingRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikePrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingStrikePrice() (field.UnderlyingStrikePrice, error) {
	var f field.UnderlyingStrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikeCurrency is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingStrikeCurrency() (field.UnderlyingStrikeCurrency, error) {
	var f field.UnderlyingStrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingOptAttribute is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingOptAttribute() (field.UnderlyingOptAttribute, error) {
	var f field.UnderlyingOptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingContractMultiplier is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingContractMultiplier() (field.UnderlyingContractMultiplier, error) {
	var f field.UnderlyingContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCouponRate() (field.UnderlyingCouponRate, error) {
	var f field.UnderlyingCouponRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityExchange is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityExchange() (field.UnderlyingSecurityExchange, error) {
	var f field.UnderlyingSecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssuer is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingIssuer() (field.UnderlyingIssuer, error) {
	var f field.UnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuerLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedUnderlyingIssuerLen() (field.EncodedUnderlyingIssuerLen, error) {
	var f field.EncodedUnderlyingIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuer is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedUnderlyingIssuer() (field.EncodedUnderlyingIssuer, error) {
	var f field.EncodedUnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityDesc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityDesc() (field.UnderlyingSecurityDesc, error) {
	var f field.UnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDescLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedUnderlyingSecurityDescLen() (field.EncodedUnderlyingSecurityDescLen, error) {
	var f field.EncodedUnderlyingSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDesc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedUnderlyingSecurityDesc() (field.EncodedUnderlyingSecurityDesc, error) {
	var f field.EncodedUnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPProgram is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCPProgram() (field.UnderlyingCPProgram, error) {
	var f field.UnderlyingCPProgram
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPRegType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCPRegType() (field.UnderlyingCPRegType, error) {
	var f field.UnderlyingCPRegType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrency is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCurrency() (field.UnderlyingCurrency, error) {
	var f field.UnderlyingCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingQty is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingQty() (field.UnderlyingQty, error) {
	var f field.UnderlyingQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPx is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingPx() (field.UnderlyingPx, error) {
	var f field.UnderlyingPx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingDirtyPrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingDirtyPrice() (field.UnderlyingDirtyPrice, error) {
	var f field.UnderlyingDirtyPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndPrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingEndPrice() (field.UnderlyingEndPrice, error) {
	var f field.UnderlyingEndPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStartValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingStartValue() (field.UnderlyingStartValue, error) {
	var f field.UnderlyingStartValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrentValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCurrentValue() (field.UnderlyingCurrentValue, error) {
	var f field.UnderlyingCurrentValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingEndValue() (field.UnderlyingEndValue, error) {
	var f field.UnderlyingEndValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingStips is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoUnderlyingStips() (field.NoUnderlyingStips, error) {
	var f field.NoUnderlyingStips
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAllocationPercent is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingAllocationPercent() (field.UnderlyingAllocationPercent, error) {
	var f field.UnderlyingAllocationPercent
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlementType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSettlementType() (field.UnderlyingSettlementType, error) {
	var f field.UnderlyingSettlementType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashAmount is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCashAmount() (field.UnderlyingCashAmount, error) {
	var f field.UnderlyingCashAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCashType() (field.UnderlyingCashType, error) {
	var f field.UnderlyingCashType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasure is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingUnitOfMeasure() (field.UnderlyingUnitOfMeasure, error) {
	var f field.UnderlyingUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTimeUnit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingTimeUnit() (field.UnderlyingTimeUnit, error) {
	var f field.UnderlyingTimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCapValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCapValue() (field.UnderlyingCapValue, error) {
	var f field.UnderlyingCapValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUndlyInstrumentParties is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoUndlyInstrumentParties() (field.NoUndlyInstrumentParties, error) {
	var f field.NoUndlyInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSettlMethod() (field.UnderlyingSettlMethod, error) {
	var f field.UnderlyingSettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAdjustedQuantity is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingAdjustedQuantity() (field.UnderlyingAdjustedQuantity, error) {
	var f field.UnderlyingAdjustedQuantity
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingFXRate() (field.UnderlyingFXRate, error) {
	var f field.UnderlyingFXRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRateCalc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingFXRateCalc() (field.UnderlyingFXRateCalc, error) {
	var f field.UnderlyingFXRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}
