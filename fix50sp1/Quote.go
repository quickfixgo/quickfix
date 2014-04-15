package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Quote msg type = S.
type Quote struct {
	message.Message
}

//QuoteBuilder builds Quote messages.
type QuoteBuilder struct {
	message.MessageBuilder
}

//CreateQuoteBuilder returns an initialized QuoteBuilder with specified required fields.
func CreateQuoteBuilder(
	quoteid field.QuoteID) QuoteBuilder {
	var builder QuoteBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(quoteid)
	return builder
}

//QuoteReqID is a non-required field for Quote.
func (m Quote) QuoteReqID() (field.QuoteReqID, error) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a required field for Quote.
func (m Quote) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteRespID is a non-required field for Quote.
func (m Quote) QuoteRespID() (field.QuoteRespID, error) {
	var f field.QuoteRespID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteType is a non-required field for Quote.
func (m Quote) QuoteType() (field.QuoteType, error) {
	var f field.QuoteType
	err := m.Body.Get(&f)
	return f, err
}

//NoQuoteQualifiers is a non-required field for Quote.
func (m Quote) NoQuoteQualifiers() (field.NoQuoteQualifiers, error) {
	var f field.NoQuoteQualifiers
	err := m.Body.Get(&f)
	return f, err
}

//QuoteResponseLevel is a non-required field for Quote.
func (m Quote) QuoteResponseLevel() (field.QuoteResponseLevel, error) {
	var f field.QuoteResponseLevel
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for Quote.
func (m Quote) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for Quote.
func (m Quote) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for Quote.
func (m Quote) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for Quote.
func (m Quote) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for Quote.
func (m Quote) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for Quote.
func (m Quote) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for Quote.
func (m Quote) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for Quote.
func (m Quote) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for Quote.
func (m Quote) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for Quote.
func (m Quote) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for Quote.
func (m Quote) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for Quote.
func (m Quote) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for Quote.
func (m Quote) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for Quote.
func (m Quote) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for Quote.
func (m Quote) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for Quote.
func (m Quote) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for Quote.
func (m Quote) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for Quote.
func (m Quote) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for Quote.
func (m Quote) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for Quote.
func (m Quote) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for Quote.
func (m Quote) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for Quote.
func (m Quote) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for Quote.
func (m Quote) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for Quote.
func (m Quote) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for Quote.
func (m Quote) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for Quote.
func (m Quote) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for Quote.
func (m Quote) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for Quote.
func (m Quote) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for Quote.
func (m Quote) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for Quote.
func (m Quote) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for Quote.
func (m Quote) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for Quote.
func (m Quote) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for Quote.
func (m Quote) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for Quote.
func (m Quote) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for Quote.
func (m Quote) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for Quote.
func (m Quote) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for Quote.
func (m Quote) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for Quote.
func (m Quote) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for Quote.
func (m Quote) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for Quote.
func (m Quote) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for Quote.
func (m Quote) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for Quote.
func (m Quote) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for Quote.
func (m Quote) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for Quote.
func (m Quote) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for Quote.
func (m Quote) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for Quote.
func (m Quote) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for Quote.
func (m Quote) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for Quote.
func (m Quote) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for Quote.
func (m Quote) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for Quote.
func (m Quote) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for Quote.
func (m Quote) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for Quote.
func (m Quote) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for Quote.
func (m Quote) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for Quote.
func (m Quote) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for Quote.
func (m Quote) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for Quote.
func (m Quote) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for Quote.
func (m Quote) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for Quote.
func (m Quote) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for Quote.
func (m Quote) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for Quote.
func (m Quote) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for Quote.
func (m Quote) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for Quote.
func (m Quote) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for Quote.
func (m Quote) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for Quote.
func (m Quote) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for Quote.
func (m Quote) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for Quote.
func (m Quote) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for Quote.
func (m Quote) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for Quote.
func (m Quote) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayAmount is a non-required field for Quote.
func (m Quote) OptPayAmount() (field.OptPayAmount, error) {
	var f field.OptPayAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for Quote.
func (m Quote) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for Quote.
func (m Quote) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for Quote.
func (m Quote) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for Quote.
func (m Quote) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for Quote.
func (m Quote) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for Quote.
func (m Quote) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for Quote.
func (m Quote) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FuturesValuationMethod is a non-required field for Quote.
func (m Quote) FuturesValuationMethod() (field.FuturesValuationMethod, error) {
	var f field.FuturesValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for Quote.
func (m Quote) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for Quote.
func (m Quote) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for Quote.
func (m Quote) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for Quote.
func (m Quote) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for Quote.
func (m Quote) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for Quote.
func (m Quote) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for Quote.
func (m Quote) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for Quote.
func (m Quote) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for Quote.
func (m Quote) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for Quote.
func (m Quote) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for Quote.
func (m Quote) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for Quote.
func (m Quote) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for Quote.
func (m Quote) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for Quote.
func (m Quote) OrderPercent() (field.OrderPercent, error) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for Quote.
func (m Quote) RoundingDirection() (field.RoundingDirection, error) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for Quote.
func (m Quote) RoundingModulus() (field.RoundingModulus, error) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for Quote.
func (m Quote) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for Quote.
func (m Quote) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate2 is a non-required field for Quote.
func (m Quote) SettlDate2() (field.SettlDate2, error) {
	var f field.SettlDate2
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty2 is a non-required field for Quote.
func (m Quote) OrderQty2() (field.OrderQty2, error) {
	var f field.OrderQty2
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for Quote.
func (m Quote) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for Quote.
func (m Quote) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for Quote.
func (m Quote) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for Quote.
func (m Quote) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for Quote.
func (m Quote) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for Quote.
func (m Quote) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//BidPx is a non-required field for Quote.
func (m Quote) BidPx() (field.BidPx, error) {
	var f field.BidPx
	err := m.Body.Get(&f)
	return f, err
}

//OfferPx is a non-required field for Quote.
func (m Quote) OfferPx() (field.OfferPx, error) {
	var f field.OfferPx
	err := m.Body.Get(&f)
	return f, err
}

//MktBidPx is a non-required field for Quote.
func (m Quote) MktBidPx() (field.MktBidPx, error) {
	var f field.MktBidPx
	err := m.Body.Get(&f)
	return f, err
}

//MktOfferPx is a non-required field for Quote.
func (m Quote) MktOfferPx() (field.MktOfferPx, error) {
	var f field.MktOfferPx
	err := m.Body.Get(&f)
	return f, err
}

//MinBidSize is a non-required field for Quote.
func (m Quote) MinBidSize() (field.MinBidSize, error) {
	var f field.MinBidSize
	err := m.Body.Get(&f)
	return f, err
}

//BidSize is a non-required field for Quote.
func (m Quote) BidSize() (field.BidSize, error) {
	var f field.BidSize
	err := m.Body.Get(&f)
	return f, err
}

//MinOfferSize is a non-required field for Quote.
func (m Quote) MinOfferSize() (field.MinOfferSize, error) {
	var f field.MinOfferSize
	err := m.Body.Get(&f)
	return f, err
}

//OfferSize is a non-required field for Quote.
func (m Quote) OfferSize() (field.OfferSize, error) {
	var f field.OfferSize
	err := m.Body.Get(&f)
	return f, err
}

//ValidUntilTime is a non-required field for Quote.
func (m Quote) ValidUntilTime() (field.ValidUntilTime, error) {
	var f field.ValidUntilTime
	err := m.Body.Get(&f)
	return f, err
}

//BidSpotRate is a non-required field for Quote.
func (m Quote) BidSpotRate() (field.BidSpotRate, error) {
	var f field.BidSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//OfferSpotRate is a non-required field for Quote.
func (m Quote) OfferSpotRate() (field.OfferSpotRate, error) {
	var f field.OfferSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//BidForwardPoints is a non-required field for Quote.
func (m Quote) BidForwardPoints() (field.BidForwardPoints, error) {
	var f field.BidForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//OfferForwardPoints is a non-required field for Quote.
func (m Quote) OfferForwardPoints() (field.OfferForwardPoints, error) {
	var f field.OfferForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//MidPx is a non-required field for Quote.
func (m Quote) MidPx() (field.MidPx, error) {
	var f field.MidPx
	err := m.Body.Get(&f)
	return f, err
}

//BidYield is a non-required field for Quote.
func (m Quote) BidYield() (field.BidYield, error) {
	var f field.BidYield
	err := m.Body.Get(&f)
	return f, err
}

//MidYield is a non-required field for Quote.
func (m Quote) MidYield() (field.MidYield, error) {
	var f field.MidYield
	err := m.Body.Get(&f)
	return f, err
}

//OfferYield is a non-required field for Quote.
func (m Quote) OfferYield() (field.OfferYield, error) {
	var f field.OfferYield
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for Quote.
func (m Quote) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a non-required field for Quote.
func (m Quote) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//BidForwardPoints2 is a non-required field for Quote.
func (m Quote) BidForwardPoints2() (field.BidForwardPoints2, error) {
	var f field.BidForwardPoints2
	err := m.Body.Get(&f)
	return f, err
}

//OfferForwardPoints2 is a non-required field for Quote.
func (m Quote) OfferForwardPoints2() (field.OfferForwardPoints2, error) {
	var f field.OfferForwardPoints2
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrBidFxRate is a non-required field for Quote.
func (m Quote) SettlCurrBidFxRate() (field.SettlCurrBidFxRate, error) {
	var f field.SettlCurrBidFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrOfferFxRate is a non-required field for Quote.
func (m Quote) SettlCurrOfferFxRate() (field.SettlCurrOfferFxRate, error) {
	var f field.SettlCurrOfferFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for Quote.
func (m Quote) SettlCurrFxRateCalc() (field.SettlCurrFxRateCalc, error) {
	var f field.SettlCurrFxRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for Quote.
func (m Quote) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for Quote.
func (m Quote) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for Quote.
func (m Quote) CustOrderCapacity() (field.CustOrderCapacity, error) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for Quote.
func (m Quote) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//OrderCapacity is a non-required field for Quote.
func (m Quote) OrderCapacity() (field.OrderCapacity, error) {
	var f field.OrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for Quote.
func (m Quote) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for Quote.
func (m Quote) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for Quote.
func (m Quote) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for Quote.
func (m Quote) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for Quote.
func (m Quote) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for Quote.
func (m Quote) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for Quote.
func (m Quote) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for Quote.
func (m Quote) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for Quote.
func (m Quote) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for Quote.
func (m Quote) YieldType() (field.YieldType, error) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for Quote.
func (m Quote) Yield() (field.Yield, error) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for Quote.
func (m Quote) YieldCalcDate() (field.YieldCalcDate, error) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for Quote.
func (m Quote) YieldRedemptionDate() (field.YieldRedemptionDate, error) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for Quote.
func (m Quote) YieldRedemptionPrice() (field.YieldRedemptionPrice, error) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for Quote.
func (m Quote) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, error) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for Quote.
func (m Quote) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for Quote.
func (m Quote) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for Quote.
func (m Quote) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//BidSwapPoints is a non-required field for Quote.
func (m Quote) BidSwapPoints() (field.BidSwapPoints, error) {
	var f field.BidSwapPoints
	err := m.Body.Get(&f)
	return f, err
}

//OfferSwapPoints is a non-required field for Quote.
func (m Quote) OfferSwapPoints() (field.OfferSwapPoints, error) {
	var f field.OfferSwapPoints
	err := m.Body.Get(&f)
	return f, err
}

//ExDestinationIDSource is a non-required field for Quote.
func (m Quote) ExDestinationIDSource() (field.ExDestinationIDSource, error) {
	var f field.ExDestinationIDSource
	err := m.Body.Get(&f)
	return f, err
}

//QuoteMsgID is a non-required field for Quote.
func (m Quote) QuoteMsgID() (field.QuoteMsgID, error) {
	var f field.QuoteMsgID
	err := m.Body.Get(&f)
	return f, err
}

//PrivateQuote is a non-required field for Quote.
func (m Quote) PrivateQuote() (field.PrivateQuote, error) {
	var f field.PrivateQuote
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for Quote.
func (m Quote) MinQty() (field.MinQty, error) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}
