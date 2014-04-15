package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CrossOrderCancelReplaceRequest msg type = t.
type CrossOrderCancelReplaceRequest struct {
	message.Message
}

//CrossOrderCancelReplaceRequestBuilder builds CrossOrderCancelReplaceRequest messages.
type CrossOrderCancelReplaceRequestBuilder struct {
	message.MessageBuilder
}

//CreateCrossOrderCancelReplaceRequestBuilder returns an initialized CrossOrderCancelReplaceRequestBuilder with specified required fields.
func CreateCrossOrderCancelReplaceRequestBuilder(
	crossid field.CrossID,
	origcrossid field.OrigCrossID,
	crosstype field.CrossType,
	crossprioritization field.CrossPrioritization,
	nosides field.NoSides,
	transacttime field.TransactTime,
	ordtype field.OrdType) CrossOrderCancelReplaceRequestBuilder {
	var builder CrossOrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(crossid)
	builder.Body.Set(origcrossid)
	builder.Body.Set(crosstype)
	builder.Body.Set(crossprioritization)
	builder.Body.Set(nosides)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//CrossID is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CrossID() (field.CrossID, error) {
	var f field.CrossID
	err := m.Body.Get(&f)
	return f, err
}

//OrigCrossID is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OrigCrossID() (field.OrigCrossID, error) {
	var f field.OrigCrossID
	err := m.Body.Get(&f)
	return f, err
}

//CrossType is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CrossType() (field.CrossType, error) {
	var f field.CrossType
	err := m.Body.Get(&f)
	return f, err
}

//CrossPrioritization is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CrossPrioritization() (field.CrossPrioritization, error) {
	var f field.CrossPrioritization
	err := m.Body.Get(&f)
	return f, err
}

//NoSides is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoSides() (field.NoSides, error) {
	var f field.NoSides
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OptPayoutAmount() (field.OptPayoutAmount, error) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ValuationMethod() (field.ValuationMethod, error) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ContractMultiplierUnit() (field.ContractMultiplierUnit, error) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) FlowScheduleType() (field.FlowScheduleType, error) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RestructuringType() (field.RestructuringType, error) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Seniority() (field.Seniority, error) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, error) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, error) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) AttachmentPoint() (field.AttachmentPoint, error) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DetachmentPoint() (field.DetachmentPoint, error) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, error) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, error) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, error) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, error) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OptPayoutType() (field.OptPayoutType, error) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoComplexEvents() (field.NoComplexEvents, error) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//HandlInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) HandlInst() (field.HandlInst, error) {
	var f field.HandlInst
	err := m.Body.Get(&f)
	return f, err
}

//ExecInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExecInst() (field.ExecInst, error) {
	var f field.ExecInst
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MinQty() (field.MinQty, error) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//MaxFloor is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaxFloor() (field.MaxFloor, error) {
	var f field.MaxFloor
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//NoTradingSessions is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoTradingSessions() (field.NoTradingSessions, error) {
	var f field.NoTradingSessions
	err := m.Body.Get(&f)
	return f, err
}

//ProcessCode is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ProcessCode() (field.ProcessCode, error) {
	var f field.ProcessCode
	err := m.Body.Get(&f)
	return f, err
}

//PrevClosePx is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PrevClosePx() (field.PrevClosePx, error) {
	var f field.PrevClosePx
	err := m.Body.Get(&f)
	return f, err
}

//LocateReqd is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) LocateReqd() (field.LocateReqd, error) {
	var f field.LocateReqd
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//StopPx is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StopPx() (field.StopPx, error) {
	var f field.StopPx
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldType() (field.YieldType, error) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Yield() (field.Yield, error) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldCalcDate() (field.YieldCalcDate, error) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldRedemptionDate() (field.YieldRedemptionDate, error) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldRedemptionPrice() (field.YieldRedemptionPrice, error) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, error) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ComplianceID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ComplianceID() (field.ComplianceID, error) {
	var f field.ComplianceID
	err := m.Body.Get(&f)
	return f, err
}

//IOIID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) IOIID() (field.IOIID, error) {
	var f field.IOIID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//TimeInForce is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TimeInForce() (field.TimeInForce, error) {
	var f field.TimeInForce
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EffectiveTime() (field.EffectiveTime, error) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExpireDate() (field.ExpireDate, error) {
	var f field.ExpireDate
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//GTBookingInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GTBookingInst() (field.GTBookingInst, error) {
	var f field.GTBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//MaxShow is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaxShow() (field.MaxShow, error) {
	var f field.MaxShow
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetValue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegOffsetValue() (field.PegOffsetValue, error) {
	var f field.PegOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//PegMoveType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegMoveType() (field.PegMoveType, error) {
	var f field.PegMoveType
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegOffsetType() (field.PegOffsetType, error) {
	var f field.PegOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//PegLimitType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegLimitType() (field.PegLimitType, error) {
	var f field.PegLimitType
	err := m.Body.Get(&f)
	return f, err
}

//PegRoundDirection is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegRoundDirection() (field.PegRoundDirection, error) {
	var f field.PegRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//PegScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegScope() (field.PegScope, error) {
	var f field.PegScope
	err := m.Body.Get(&f)
	return f, err
}

//PegPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegPriceType() (field.PegPriceType, error) {
	var f field.PegPriceType
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegSecurityIDSource() (field.PegSecurityIDSource, error) {
	var f field.PegSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegSecurityID() (field.PegSecurityID, error) {
	var f field.PegSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//PegSymbol is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegSymbol() (field.PegSymbol, error) {
	var f field.PegSymbol
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegSecurityDesc() (field.PegSecurityDesc, error) {
	var f field.PegSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionInst() (field.DiscretionInst, error) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetValue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionOffsetValue() (field.DiscretionOffsetValue, error) {
	var f field.DiscretionOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionMoveType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionMoveType() (field.DiscretionMoveType, error) {
	var f field.DiscretionMoveType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionOffsetType() (field.DiscretionOffsetType, error) {
	var f field.DiscretionOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionLimitType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionLimitType() (field.DiscretionLimitType, error) {
	var f field.DiscretionLimitType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionRoundDirection is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionRoundDirection() (field.DiscretionRoundDirection, error) {
	var f field.DiscretionRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionScope() (field.DiscretionScope, error) {
	var f field.DiscretionScope
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategy is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TargetStrategy() (field.TargetStrategy, error) {
	var f field.TargetStrategy
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategyParameters is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TargetStrategyParameters() (field.TargetStrategyParameters, error) {
	var f field.TargetStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//ParticipationRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ParticipationRate() (field.ParticipationRate, error) {
	var f field.ParticipationRate
	err := m.Body.Get(&f)
	return f, err
}

//CancellationRights is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CancellationRights() (field.CancellationRights, error) {
	var f field.CancellationRights
	err := m.Body.Get(&f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MoneyLaunderingStatus() (field.MoneyLaunderingStatus, error) {
	var f field.MoneyLaunderingStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RegistID() (field.RegistID, error) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//Designation is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Designation() (field.Designation, error) {
	var f field.Designation
	err := m.Body.Get(&f)
	return f, err
}

//NoStrategyParameters is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoStrategyParameters() (field.NoStrategyParameters, error) {
	var f field.NoStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//HostCrossID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) HostCrossID() (field.HostCrossID, error) {
	var f field.HostCrossID
	err := m.Body.Get(&f)
	return f, err
}

//TransBkdTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TransBkdTime() (field.TransBkdTime, error) {
	var f field.TransBkdTime
	err := m.Body.Get(&f)
	return f, err
}

//NoRootPartyIDs is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoRootPartyIDs() (field.NoRootPartyIDs, error) {
	var f field.NoRootPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//MatchIncrement is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MatchIncrement() (field.MatchIncrement, error) {
	var f field.MatchIncrement
	err := m.Body.Get(&f)
	return f, err
}

//MaxPriceLevels is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaxPriceLevels() (field.MaxPriceLevels, error) {
	var f field.MaxPriceLevels
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryDisplayQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecondaryDisplayQty() (field.SecondaryDisplayQty, error) {
	var f field.SecondaryDisplayQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayWhen is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DisplayWhen() (field.DisplayWhen, error) {
	var f field.DisplayWhen
	err := m.Body.Get(&f)
	return f, err
}

//DisplayMethod is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DisplayMethod() (field.DisplayMethod, error) {
	var f field.DisplayMethod
	err := m.Body.Get(&f)
	return f, err
}

//DisplayLowQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DisplayLowQty() (field.DisplayLowQty, error) {
	var f field.DisplayLowQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayHighQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DisplayHighQty() (field.DisplayHighQty, error) {
	var f field.DisplayHighQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayMinIncr is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DisplayMinIncr() (field.DisplayMinIncr, error) {
	var f field.DisplayMinIncr
	err := m.Body.Get(&f)
	return f, err
}

//RefreshQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RefreshQty() (field.RefreshQty, error) {
	var f field.RefreshQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DisplayQty() (field.DisplayQty, error) {
	var f field.DisplayQty
	err := m.Body.Get(&f)
	return f, err
}

//PriceProtectionScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PriceProtectionScope() (field.PriceProtectionScope, error) {
	var f field.PriceProtectionScope
	err := m.Body.Get(&f)
	return f, err
}

//TriggerType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerType() (field.TriggerType, error) {
	var f field.TriggerType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerAction is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerAction() (field.TriggerAction, error) {
	var f field.TriggerAction
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerPrice() (field.TriggerPrice, error) {
	var f field.TriggerPrice
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSymbol is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerSymbol() (field.TriggerSymbol, error) {
	var f field.TriggerSymbol
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerSecurityID() (field.TriggerSecurityID, error) {
	var f field.TriggerSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerSecurityIDSource() (field.TriggerSecurityIDSource, error) {
	var f field.TriggerSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerSecurityDesc() (field.TriggerSecurityDesc, error) {
	var f field.TriggerSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerPriceType() (field.TriggerPriceType, error) {
	var f field.TriggerPriceType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceTypeScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerPriceTypeScope() (field.TriggerPriceTypeScope, error) {
	var f field.TriggerPriceTypeScope
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceDirection is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerPriceDirection() (field.TriggerPriceDirection, error) {
	var f field.TriggerPriceDirection
	err := m.Body.Get(&f)
	return f, err
}

//TriggerNewPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerNewPrice() (field.TriggerNewPrice, error) {
	var f field.TriggerNewPrice
	err := m.Body.Get(&f)
	return f, err
}

//TriggerOrderType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerOrderType() (field.TriggerOrderType, error) {
	var f field.TriggerOrderType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerNewQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerNewQty() (field.TriggerNewQty, error) {
	var f field.TriggerNewQty
	err := m.Body.Get(&f)
	return f, err
}

//TriggerTradingSessionID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerTradingSessionID() (field.TriggerTradingSessionID, error) {
	var f field.TriggerTradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TriggerTradingSessionSubID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TriggerTradingSessionSubID() (field.TriggerTradingSessionSubID, error) {
	var f field.TriggerTradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//ExDestinationIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExDestinationIDSource() (field.ExDestinationIDSource, error) {
	var f field.ExDestinationIDSource
	err := m.Body.Get(&f)
	return f, err
}
