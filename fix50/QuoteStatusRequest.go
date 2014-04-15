package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteStatusRequest msg type = a.
type QuoteStatusRequest struct {
	message.Message
}

//QuoteStatusRequestBuilder builds QuoteStatusRequest messages.
type QuoteStatusRequestBuilder struct {
	message.MessageBuilder
}

//CreateQuoteStatusRequestBuilder returns an initialized QuoteStatusRequestBuilder with specified required fields.
func CreateQuoteStatusRequestBuilder() QuoteStatusRequestBuilder {
	var builder QuoteStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//QuoteStatusReqID is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) QuoteStatusReqID() (field.QuoteStatusReqID, error) {
	var f field.QuoteStatusReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for QuoteStatusRequest.
func (m QuoteStatusRequest) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}
