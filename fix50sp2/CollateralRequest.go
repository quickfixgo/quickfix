package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CollateralRequest msg type = AX.
type CollateralRequest struct {
	message.Message
}

//CollateralRequestBuilder builds CollateralRequest messages.
type CollateralRequestBuilder struct {
	message.MessageBuilder
}

//CreateCollateralRequestBuilder returns an initialized CollateralRequestBuilder with specified required fields.
func CreateCollateralRequestBuilder(
	collreqid field.CollReqID,
	collasgnreason field.CollAsgnReason,
	transacttime field.TransactTime) CollateralRequestBuilder {
	var builder CollateralRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(collreqid)
	builder.Body.Set(collasgnreason)
	builder.Body.Set(transacttime)
	return builder
}

//CollReqID is a required field for CollateralRequest.
func (m CollateralRequest) CollReqID() (field.CollReqID, error) {
	var f field.CollReqID
	err := m.Body.Get(&f)
	return f, err
}

//CollAsgnReason is a required field for CollateralRequest.
func (m CollateralRequest) CollAsgnReason() (field.CollAsgnReason, error) {
	var f field.CollAsgnReason
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for CollateralRequest.
func (m CollateralRequest) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for CollateralRequest.
func (m CollateralRequest) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for CollateralRequest.
func (m CollateralRequest) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for CollateralRequest.
func (m CollateralRequest) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for CollateralRequest.
func (m CollateralRequest) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for CollateralRequest.
func (m CollateralRequest) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a non-required field for CollateralRequest.
func (m CollateralRequest) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for CollateralRequest.
func (m CollateralRequest) SecondaryOrderID() (field.SecondaryOrderID, error) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for CollateralRequest.
func (m CollateralRequest) SecondaryClOrdID() (field.SecondaryClOrdID, error) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//NoExecs is a non-required field for CollateralRequest.
func (m CollateralRequest) NoExecs() (field.NoExecs, error) {
	var f field.NoExecs
	err := m.Body.Get(&f)
	return f, err
}

//NoTrades is a non-required field for CollateralRequest.
func (m CollateralRequest) NoTrades() (field.NoTrades, error) {
	var f field.NoTrades
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for CollateralRequest.
func (m CollateralRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for CollateralRequest.
func (m CollateralRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for CollateralRequest.
func (m CollateralRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for CollateralRequest.
func (m CollateralRequest) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for CollateralRequest.
func (m CollateralRequest) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for CollateralRequest.
func (m CollateralRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for CollateralRequest.
func (m CollateralRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for CollateralRequest.
func (m CollateralRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for CollateralRequest.
func (m CollateralRequest) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for CollateralRequest.
func (m CollateralRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for CollateralRequest.
func (m CollateralRequest) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for CollateralRequest.
func (m CollateralRequest) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for CollateralRequest.
func (m CollateralRequest) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for CollateralRequest.
func (m CollateralRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for CollateralRequest.
func (m CollateralRequest) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for CollateralRequest.
func (m CollateralRequest) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for CollateralRequest.
func (m CollateralRequest) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for CollateralRequest.
func (m CollateralRequest) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for CollateralRequest.
func (m CollateralRequest) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for CollateralRequest.
func (m CollateralRequest) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for CollateralRequest.
func (m CollateralRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for CollateralRequest.
func (m CollateralRequest) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for CollateralRequest.
func (m CollateralRequest) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for CollateralRequest.
func (m CollateralRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for CollateralRequest.
func (m CollateralRequest) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for CollateralRequest.
func (m CollateralRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for CollateralRequest.
func (m CollateralRequest) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for CollateralRequest.
func (m CollateralRequest) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for CollateralRequest.
func (m CollateralRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for CollateralRequest.
func (m CollateralRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for CollateralRequest.
func (m CollateralRequest) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for CollateralRequest.
func (m CollateralRequest) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for CollateralRequest.
func (m CollateralRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for CollateralRequest.
func (m CollateralRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for CollateralRequest.
func (m CollateralRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for CollateralRequest.
func (m CollateralRequest) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for CollateralRequest.
func (m CollateralRequest) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for CollateralRequest.
func (m CollateralRequest) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for CollateralRequest.
func (m CollateralRequest) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for CollateralRequest.
func (m CollateralRequest) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for CollateralRequest.
func (m CollateralRequest) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for CollateralRequest.
func (m CollateralRequest) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for CollateralRequest.
func (m CollateralRequest) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for CollateralRequest.
func (m CollateralRequest) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for CollateralRequest.
func (m CollateralRequest) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for CollateralRequest.
func (m CollateralRequest) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for CollateralRequest.
func (m CollateralRequest) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for CollateralRequest.
func (m CollateralRequest) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for CollateralRequest.
func (m CollateralRequest) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for CollateralRequest.
func (m CollateralRequest) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for CollateralRequest.
func (m CollateralRequest) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for CollateralRequest.
func (m CollateralRequest) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for CollateralRequest.
func (m CollateralRequest) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for CollateralRequest.
func (m CollateralRequest) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for CollateralRequest.
func (m CollateralRequest) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for CollateralRequest.
func (m CollateralRequest) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for CollateralRequest.
func (m CollateralRequest) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for CollateralRequest.
func (m CollateralRequest) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for CollateralRequest.
func (m CollateralRequest) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for CollateralRequest.
func (m CollateralRequest) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for CollateralRequest.
func (m CollateralRequest) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for CollateralRequest.
func (m CollateralRequest) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for CollateralRequest.
func (m CollateralRequest) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for CollateralRequest.
func (m CollateralRequest) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for CollateralRequest.
func (m CollateralRequest) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for CollateralRequest.
func (m CollateralRequest) OptPayoutAmount() (field.OptPayoutAmount, error) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for CollateralRequest.
func (m CollateralRequest) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for CollateralRequest.
func (m CollateralRequest) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for CollateralRequest.
func (m CollateralRequest) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for CollateralRequest.
func (m CollateralRequest) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for CollateralRequest.
func (m CollateralRequest) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for CollateralRequest.
func (m CollateralRequest) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for CollateralRequest.
func (m CollateralRequest) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for CollateralRequest.
func (m CollateralRequest) ValuationMethod() (field.ValuationMethod, error) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for CollateralRequest.
func (m CollateralRequest) ContractMultiplierUnit() (field.ContractMultiplierUnit, error) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for CollateralRequest.
func (m CollateralRequest) FlowScheduleType() (field.FlowScheduleType, error) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for CollateralRequest.
func (m CollateralRequest) RestructuringType() (field.RestructuringType, error) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for CollateralRequest.
func (m CollateralRequest) Seniority() (field.Seniority, error) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for CollateralRequest.
func (m CollateralRequest) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, error) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for CollateralRequest.
func (m CollateralRequest) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, error) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for CollateralRequest.
func (m CollateralRequest) AttachmentPoint() (field.AttachmentPoint, error) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for CollateralRequest.
func (m CollateralRequest) DetachmentPoint() (field.DetachmentPoint, error) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for CollateralRequest.
func (m CollateralRequest) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, error) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for CollateralRequest.
func (m CollateralRequest) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, error) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for CollateralRequest.
func (m CollateralRequest) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, error) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for CollateralRequest.
func (m CollateralRequest) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, error) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for CollateralRequest.
func (m CollateralRequest) OptPayoutType() (field.OptPayoutType, error) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for CollateralRequest.
func (m CollateralRequest) NoComplexEvents() (field.NoComplexEvents, error) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for CollateralRequest.
func (m CollateralRequest) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for CollateralRequest.
func (m CollateralRequest) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for CollateralRequest.
func (m CollateralRequest) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for CollateralRequest.
func (m CollateralRequest) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for CollateralRequest.
func (m CollateralRequest) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for CollateralRequest.
func (m CollateralRequest) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for CollateralRequest.
func (m CollateralRequest) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for CollateralRequest.
func (m CollateralRequest) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for CollateralRequest.
func (m CollateralRequest) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for CollateralRequest.
func (m CollateralRequest) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//Quantity is a non-required field for CollateralRequest.
func (m CollateralRequest) Quantity() (field.Quantity, error) {
	var f field.Quantity
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for CollateralRequest.
func (m CollateralRequest) QtyType() (field.QtyType, error) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for CollateralRequest.
func (m CollateralRequest) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for CollateralRequest.
func (m CollateralRequest) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for CollateralRequest.
func (m CollateralRequest) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//MarginExcess is a non-required field for CollateralRequest.
func (m CollateralRequest) MarginExcess() (field.MarginExcess, error) {
	var f field.MarginExcess
	err := m.Body.Get(&f)
	return f, err
}

//TotalNetValue is a non-required field for CollateralRequest.
func (m CollateralRequest) TotalNetValue() (field.TotalNetValue, error) {
	var f field.TotalNetValue
	err := m.Body.Get(&f)
	return f, err
}

//CashOutstanding is a non-required field for CollateralRequest.
func (m CollateralRequest) CashOutstanding() (field.CashOutstanding, error) {
	var f field.CashOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for CollateralRequest.
func (m CollateralRequest) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, error) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for CollateralRequest.
func (m CollateralRequest) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//NoMiscFees is a non-required field for CollateralRequest.
func (m CollateralRequest) NoMiscFees() (field.NoMiscFees, error) {
	var f field.NoMiscFees
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for CollateralRequest.
func (m CollateralRequest) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for CollateralRequest.
func (m CollateralRequest) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestAmt is a non-required field for CollateralRequest.
func (m CollateralRequest) AccruedInterestAmt() (field.AccruedInterestAmt, error) {
	var f field.AccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for CollateralRequest.
func (m CollateralRequest) EndAccruedInterestAmt() (field.EndAccruedInterestAmt, error) {
	var f field.EndAccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//StartCash is a non-required field for CollateralRequest.
func (m CollateralRequest) StartCash() (field.StartCash, error) {
	var f field.StartCash
	err := m.Body.Get(&f)
	return f, err
}

//EndCash is a non-required field for CollateralRequest.
func (m CollateralRequest) EndCash() (field.EndCash, error) {
	var f field.EndCash
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for CollateralRequest.
func (m CollateralRequest) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for CollateralRequest.
func (m CollateralRequest) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for CollateralRequest.
func (m CollateralRequest) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for CollateralRequest.
func (m CollateralRequest) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for CollateralRequest.
func (m CollateralRequest) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for CollateralRequest.
func (m CollateralRequest) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for CollateralRequest.
func (m CollateralRequest) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for CollateralRequest.
func (m CollateralRequest) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for CollateralRequest.
func (m CollateralRequest) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for CollateralRequest.
func (m CollateralRequest) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for CollateralRequest.
func (m CollateralRequest) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for CollateralRequest.
func (m CollateralRequest) SettlSessID() (field.SettlSessID, error) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a non-required field for CollateralRequest.
func (m CollateralRequest) SettlSessSubID() (field.SettlSessSubID, error) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for CollateralRequest.
func (m CollateralRequest) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for CollateralRequest.
func (m CollateralRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for CollateralRequest.
func (m CollateralRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for CollateralRequest.
func (m CollateralRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
