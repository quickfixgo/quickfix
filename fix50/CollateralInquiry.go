package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CollateralInquiry msg type = BB.
type CollateralInquiry struct {
	message.Message
}

//CollateralInquiryBuilder builds CollateralInquiry messages.
type CollateralInquiryBuilder struct {
	message.MessageBuilder
}

//CreateCollateralInquiryBuilder returns an initialized CollateralInquiryBuilder with specified required fields.
func CreateCollateralInquiryBuilder() CollateralInquiryBuilder {
	var builder CollateralInquiryBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//CollInquiryID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CollInquiryID() (field.CollInquiryID, error) {
	var f field.CollInquiryID
	err := m.Body.Get(&f)
	return f, err
}

//NoCollInquiryQualifier is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoCollInquiryQualifier() (field.NoCollInquiryQualifier, error) {
	var f field.NoCollInquiryQualifier
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//ResponseTransportType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ResponseTransportType() (field.ResponseTransportType, error) {
	var f field.ResponseTransportType
	err := m.Body.Get(&f)
	return f, err
}

//ResponseDestination is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ResponseDestination() (field.ResponseDestination, error) {
	var f field.ResponseDestination
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecondaryOrderID() (field.SecondaryOrderID, error) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecondaryClOrdID() (field.SecondaryClOrdID, error) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//NoExecs is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoExecs() (field.NoExecs, error) {
	var f field.NoExecs
	err := m.Body.Get(&f)
	return f, err
}

//NoTrades is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoTrades() (field.NoTrades, error) {
	var f field.NoTrades
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for CollateralInquiry.
func (m CollateralInquiry) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for CollateralInquiry.
func (m CollateralInquiry) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for CollateralInquiry.
func (m CollateralInquiry) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for CollateralInquiry.
func (m CollateralInquiry) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for CollateralInquiry.
func (m CollateralInquiry) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for CollateralInquiry.
func (m CollateralInquiry) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for CollateralInquiry.
func (m CollateralInquiry) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for CollateralInquiry.
func (m CollateralInquiry) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for CollateralInquiry.
func (m CollateralInquiry) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for CollateralInquiry.
func (m CollateralInquiry) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//Quantity is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Quantity() (field.Quantity, error) {
	var f field.Quantity
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) QtyType() (field.QtyType, error) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//MarginExcess is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MarginExcess() (field.MarginExcess, error) {
	var f field.MarginExcess
	err := m.Body.Get(&f)
	return f, err
}

//TotalNetValue is a non-required field for CollateralInquiry.
func (m CollateralInquiry) TotalNetValue() (field.TotalNetValue, error) {
	var f field.TotalNetValue
	err := m.Body.Get(&f)
	return f, err
}

//CashOutstanding is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CashOutstanding() (field.CashOutstanding, error) {
	var f field.CashOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, error) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestAmt is a non-required field for CollateralInquiry.
func (m CollateralInquiry) AccruedInterestAmt() (field.AccruedInterestAmt, error) {
	var f field.AccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EndAccruedInterestAmt() (field.EndAccruedInterestAmt, error) {
	var f field.EndAccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//StartCash is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StartCash() (field.StartCash, error) {
	var f field.StartCash
	err := m.Body.Get(&f)
	return f, err
}

//EndCash is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EndCash() (field.EndCash, error) {
	var f field.EndCash
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//SettlDeliveryType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SettlDeliveryType() (field.SettlDeliveryType, error) {
	var f field.SettlDeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StandInstDbType() (field.StandInstDbType, error) {
	var f field.StandInstDbType
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbName is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StandInstDbName() (field.StandInstDbName, error) {
	var f field.StandInstDbName
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StandInstDbID() (field.StandInstDbID, error) {
	var f field.StandInstDbID
	err := m.Body.Get(&f)
	return f, err
}

//NoDlvyInst is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoDlvyInst() (field.NoDlvyInst, error) {
	var f field.NoDlvyInst
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SettlSessID() (field.SettlSessID, error) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SettlSessSubID() (field.SettlSessSubID, error) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
