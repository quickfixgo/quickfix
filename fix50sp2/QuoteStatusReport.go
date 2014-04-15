package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteStatusReport msg type = AI.
type QuoteStatusReport struct {
	message.Message
}

//QuoteStatusReportBuilder builds QuoteStatusReport messages.
type QuoteStatusReportBuilder struct {
	message.MessageBuilder
}

//CreateQuoteStatusReportBuilder returns an initialized QuoteStatusReportBuilder with specified required fields.
func CreateQuoteStatusReportBuilder() QuoteStatusReportBuilder {
	var builder QuoteStatusReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//QuoteStatusReqID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteStatusReqID() (field.QuoteStatusReqID, error) {
	var f field.QuoteStatusReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteReqID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteReqID() (field.QuoteReqID, error) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteRespID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteRespID() (field.QuoteRespID, error) {
	var f field.QuoteRespID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteType() (field.QuoteType, error) {
	var f field.QuoteType
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OptPayoutAmount() (field.OptPayoutAmount, error) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ValuationMethod() (field.ValuationMethod, error) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ContractMultiplierUnit() (field.ContractMultiplierUnit, error) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) FlowScheduleType() (field.FlowScheduleType, error) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RestructuringType() (field.RestructuringType, error) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Seniority() (field.Seniority, error) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, error) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, error) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AttachmentPoint() (field.AttachmentPoint, error) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) DetachmentPoint() (field.DetachmentPoint, error) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, error) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, error) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, error) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, error) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OptPayoutType() (field.OptPayoutType, error) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoComplexEvents() (field.NoComplexEvents, error) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrderPercent() (field.OrderPercent, error) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RoundingDirection() (field.RoundingDirection, error) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) RoundingModulus() (field.RoundingModulus, error) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlDate2() (field.SettlDate2, error) {
	var f field.SettlDate2
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrderQty2() (field.OrderQty2, error) {
	var f field.OrderQty2
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoQuoteQualifiers is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoQuoteQualifiers() (field.NoQuoteQualifiers, error) {
	var f field.NoQuoteQualifiers
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) YieldType() (field.YieldType, error) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Yield() (field.Yield, error) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) YieldCalcDate() (field.YieldCalcDate, error) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) YieldRedemptionDate() (field.YieldRedemptionDate, error) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) YieldRedemptionPrice() (field.YieldRedemptionPrice, error) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, error) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BidPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidPx() (field.BidPx, error) {
	var f field.BidPx
	err := m.Body.Get(&f)
	return f, err
}

//OfferPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferPx() (field.OfferPx, error) {
	var f field.OfferPx
	err := m.Body.Get(&f)
	return f, err
}

//MktBidPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MktBidPx() (field.MktBidPx, error) {
	var f field.MktBidPx
	err := m.Body.Get(&f)
	return f, err
}

//MktOfferPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MktOfferPx() (field.MktOfferPx, error) {
	var f field.MktOfferPx
	err := m.Body.Get(&f)
	return f, err
}

//MinBidSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinBidSize() (field.MinBidSize, error) {
	var f field.MinBidSize
	err := m.Body.Get(&f)
	return f, err
}

//BidSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidSize() (field.BidSize, error) {
	var f field.BidSize
	err := m.Body.Get(&f)
	return f, err
}

//MinOfferSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinOfferSize() (field.MinOfferSize, error) {
	var f field.MinOfferSize
	err := m.Body.Get(&f)
	return f, err
}

//OfferSize is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferSize() (field.OfferSize, error) {
	var f field.OfferSize
	err := m.Body.Get(&f)
	return f, err
}

//ValidUntilTime is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ValidUntilTime() (field.ValidUntilTime, error) {
	var f field.ValidUntilTime
	err := m.Body.Get(&f)
	return f, err
}

//BidSpotRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidSpotRate() (field.BidSpotRate, error) {
	var f field.BidSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//OfferSpotRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferSpotRate() (field.OfferSpotRate, error) {
	var f field.OfferSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//BidForwardPoints is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidForwardPoints() (field.BidForwardPoints, error) {
	var f field.BidForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//OfferForwardPoints is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferForwardPoints() (field.OfferForwardPoints, error) {
	var f field.OfferForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//MidPx is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MidPx() (field.MidPx, error) {
	var f field.MidPx
	err := m.Body.Get(&f)
	return f, err
}

//BidYield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidYield() (field.BidYield, error) {
	var f field.BidYield
	err := m.Body.Get(&f)
	return f, err
}

//MidYield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MidYield() (field.MidYield, error) {
	var f field.MidYield
	err := m.Body.Get(&f)
	return f, err
}

//OfferYield is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferYield() (field.OfferYield, error) {
	var f field.OfferYield
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//BidForwardPoints2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BidForwardPoints2() (field.BidForwardPoints2, error) {
	var f field.BidForwardPoints2
	err := m.Body.Get(&f)
	return f, err
}

//OfferForwardPoints2 is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OfferForwardPoints2() (field.OfferForwardPoints2, error) {
	var f field.OfferForwardPoints2
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrBidFxRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlCurrBidFxRate() (field.SettlCurrBidFxRate, error) {
	var f field.SettlCurrBidFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrOfferFxRate is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlCurrOfferFxRate() (field.SettlCurrOfferFxRate, error) {
	var f field.SettlCurrOfferFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) SettlCurrFxRateCalc() (field.SettlCurrFxRateCalc, error) {
	var f field.SettlCurrFxRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) CustOrderCapacity() (field.CustOrderCapacity, error) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//QuoteStatus is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteStatus() (field.QuoteStatus, error) {
	var f field.QuoteStatus
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//ExDestinationIDSource is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) ExDestinationIDSource() (field.ExDestinationIDSource, error) {
	var f field.ExDestinationIDSource
	err := m.Body.Get(&f)
	return f, err
}

//QuoteCancelType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteCancelType() (field.QuoteCancelType, error) {
	var f field.QuoteCancelType
	err := m.Body.Get(&f)
	return f, err
}

//QuoteMsgID is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteMsgID() (field.QuoteMsgID, error) {
	var f field.QuoteMsgID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteRejectReason is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) QuoteRejectReason() (field.QuoteRejectReason, error) {
	var f field.QuoteRejectReason
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) MinQty() (field.MinQty, error) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//BookingType is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) BookingType() (field.BookingType, error) {
	var f field.BookingType
	err := m.Body.Get(&f)
	return f, err
}

//OrderCapacity is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrderCapacity() (field.OrderCapacity, error) {
	var f field.OrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//OrderRestrictions is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) OrderRestrictions() (field.OrderRestrictions, error) {
	var f field.OrderRestrictions
	err := m.Body.Get(&f)
	return f, err
}

//NoTargetPartyIDs is a non-required field for QuoteStatusReport.
func (m QuoteStatusReport) NoTargetPartyIDs() (field.NoTargetPartyIDs, error) {
	var f field.NoTargetPartyIDs
	err := m.Body.Get(&f)
	return f, err
}
