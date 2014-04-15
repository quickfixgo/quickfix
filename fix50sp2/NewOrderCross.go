package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderCross msg type = s.
type NewOrderCross struct {
	message.Message
}

//NewOrderCrossBuilder builds NewOrderCross messages.
type NewOrderCrossBuilder struct {
	message.MessageBuilder
}

//CreateNewOrderCrossBuilder returns an initialized NewOrderCrossBuilder with specified required fields.
func CreateNewOrderCrossBuilder(
	crossid field.CrossID,
	crosstype field.CrossType,
	crossprioritization field.CrossPrioritization,
	nosides field.NoSides,
	transacttime field.TransactTime,
	ordtype field.OrdType) NewOrderCrossBuilder {
	var builder NewOrderCrossBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(crossid)
	builder.Body.Set(crosstype)
	builder.Body.Set(crossprioritization)
	builder.Body.Set(nosides)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//CrossID is a required field for NewOrderCross.
func (m NewOrderCross) CrossID() (field.CrossID, error) {
	var f field.CrossID
	err := m.Body.Get(&f)
	return f, err
}

//CrossType is a required field for NewOrderCross.
func (m NewOrderCross) CrossType() (field.CrossType, error) {
	var f field.CrossType
	err := m.Body.Get(&f)
	return f, err
}

//CrossPrioritization is a required field for NewOrderCross.
func (m NewOrderCross) CrossPrioritization() (field.CrossPrioritization, error) {
	var f field.CrossPrioritization
	err := m.Body.Get(&f)
	return f, err
}

//NoSides is a required field for NewOrderCross.
func (m NewOrderCross) NoSides() (field.NoSides, error) {
	var f field.NoSides
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for NewOrderCross.
func (m NewOrderCross) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for NewOrderCross.
func (m NewOrderCross) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for NewOrderCross.
func (m NewOrderCross) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for NewOrderCross.
func (m NewOrderCross) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for NewOrderCross.
func (m NewOrderCross) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for NewOrderCross.
func (m NewOrderCross) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for NewOrderCross.
func (m NewOrderCross) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for NewOrderCross.
func (m NewOrderCross) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for NewOrderCross.
func (m NewOrderCross) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for NewOrderCross.
func (m NewOrderCross) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for NewOrderCross.
func (m NewOrderCross) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for NewOrderCross.
func (m NewOrderCross) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for NewOrderCross.
func (m NewOrderCross) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for NewOrderCross.
func (m NewOrderCross) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for NewOrderCross.
func (m NewOrderCross) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for NewOrderCross.
func (m NewOrderCross) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for NewOrderCross.
func (m NewOrderCross) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for NewOrderCross.
func (m NewOrderCross) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for NewOrderCross.
func (m NewOrderCross) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for NewOrderCross.
func (m NewOrderCross) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for NewOrderCross.
func (m NewOrderCross) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for NewOrderCross.
func (m NewOrderCross) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for NewOrderCross.
func (m NewOrderCross) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for NewOrderCross.
func (m NewOrderCross) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for NewOrderCross.
func (m NewOrderCross) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for NewOrderCross.
func (m NewOrderCross) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for NewOrderCross.
func (m NewOrderCross) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for NewOrderCross.
func (m NewOrderCross) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for NewOrderCross.
func (m NewOrderCross) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for NewOrderCross.
func (m NewOrderCross) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for NewOrderCross.
func (m NewOrderCross) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for NewOrderCross.
func (m NewOrderCross) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for NewOrderCross.
func (m NewOrderCross) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for NewOrderCross.
func (m NewOrderCross) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for NewOrderCross.
func (m NewOrderCross) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for NewOrderCross.
func (m NewOrderCross) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for NewOrderCross.
func (m NewOrderCross) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for NewOrderCross.
func (m NewOrderCross) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for NewOrderCross.
func (m NewOrderCross) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for NewOrderCross.
func (m NewOrderCross) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for NewOrderCross.
func (m NewOrderCross) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for NewOrderCross.
func (m NewOrderCross) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for NewOrderCross.
func (m NewOrderCross) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for NewOrderCross.
func (m NewOrderCross) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for NewOrderCross.
func (m NewOrderCross) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for NewOrderCross.
func (m NewOrderCross) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for NewOrderCross.
func (m NewOrderCross) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for NewOrderCross.
func (m NewOrderCross) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for NewOrderCross.
func (m NewOrderCross) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for NewOrderCross.
func (m NewOrderCross) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for NewOrderCross.
func (m NewOrderCross) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for NewOrderCross.
func (m NewOrderCross) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for NewOrderCross.
func (m NewOrderCross) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for NewOrderCross.
func (m NewOrderCross) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for NewOrderCross.
func (m NewOrderCross) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for NewOrderCross.
func (m NewOrderCross) OptPayoutAmount() (field.OptPayoutAmount, error) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for NewOrderCross.
func (m NewOrderCross) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for NewOrderCross.
func (m NewOrderCross) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for NewOrderCross.
func (m NewOrderCross) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for NewOrderCross.
func (m NewOrderCross) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for NewOrderCross.
func (m NewOrderCross) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for NewOrderCross.
func (m NewOrderCross) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for NewOrderCross.
func (m NewOrderCross) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for NewOrderCross.
func (m NewOrderCross) ValuationMethod() (field.ValuationMethod, error) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for NewOrderCross.
func (m NewOrderCross) ContractMultiplierUnit() (field.ContractMultiplierUnit, error) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for NewOrderCross.
func (m NewOrderCross) FlowScheduleType() (field.FlowScheduleType, error) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for NewOrderCross.
func (m NewOrderCross) RestructuringType() (field.RestructuringType, error) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for NewOrderCross.
func (m NewOrderCross) Seniority() (field.Seniority, error) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for NewOrderCross.
func (m NewOrderCross) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, error) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for NewOrderCross.
func (m NewOrderCross) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, error) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for NewOrderCross.
func (m NewOrderCross) AttachmentPoint() (field.AttachmentPoint, error) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for NewOrderCross.
func (m NewOrderCross) DetachmentPoint() (field.DetachmentPoint, error) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for NewOrderCross.
func (m NewOrderCross) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, error) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for NewOrderCross.
func (m NewOrderCross) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, error) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for NewOrderCross.
func (m NewOrderCross) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, error) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for NewOrderCross.
func (m NewOrderCross) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, error) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for NewOrderCross.
func (m NewOrderCross) OptPayoutType() (field.OptPayoutType, error) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for NewOrderCross.
func (m NewOrderCross) NoComplexEvents() (field.NoComplexEvents, error) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for NewOrderCross.
func (m NewOrderCross) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for NewOrderCross.
func (m NewOrderCross) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for NewOrderCross.
func (m NewOrderCross) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for NewOrderCross.
func (m NewOrderCross) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//HandlInst is a non-required field for NewOrderCross.
func (m NewOrderCross) HandlInst() (field.HandlInst, error) {
	var f field.HandlInst
	err := m.Body.Get(&f)
	return f, err
}

//ExecInst is a non-required field for NewOrderCross.
func (m NewOrderCross) ExecInst() (field.ExecInst, error) {
	var f field.ExecInst
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for NewOrderCross.
func (m NewOrderCross) MinQty() (field.MinQty, error) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//MaxFloor is a non-required field for NewOrderCross.
func (m NewOrderCross) MaxFloor() (field.MaxFloor, error) {
	var f field.MaxFloor
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for NewOrderCross.
func (m NewOrderCross) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//NoTradingSessions is a non-required field for NewOrderCross.
func (m NewOrderCross) NoTradingSessions() (field.NoTradingSessions, error) {
	var f field.NoTradingSessions
	err := m.Body.Get(&f)
	return f, err
}

//ProcessCode is a non-required field for NewOrderCross.
func (m NewOrderCross) ProcessCode() (field.ProcessCode, error) {
	var f field.ProcessCode
	err := m.Body.Get(&f)
	return f, err
}

//PrevClosePx is a non-required field for NewOrderCross.
func (m NewOrderCross) PrevClosePx() (field.PrevClosePx, error) {
	var f field.PrevClosePx
	err := m.Body.Get(&f)
	return f, err
}

//LocateReqd is a non-required field for NewOrderCross.
func (m NewOrderCross) LocateReqd() (field.LocateReqd, error) {
	var f field.LocateReqd
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for NewOrderCross.
func (m NewOrderCross) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for NewOrderCross.
func (m NewOrderCross) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a required field for NewOrderCross.
func (m NewOrderCross) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for NewOrderCross.
func (m NewOrderCross) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for NewOrderCross.
func (m NewOrderCross) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//StopPx is a non-required field for NewOrderCross.
func (m NewOrderCross) StopPx() (field.StopPx, error) {
	var f field.StopPx
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for NewOrderCross.
func (m NewOrderCross) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for NewOrderCross.
func (m NewOrderCross) YieldType() (field.YieldType, error) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for NewOrderCross.
func (m NewOrderCross) Yield() (field.Yield, error) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for NewOrderCross.
func (m NewOrderCross) YieldCalcDate() (field.YieldCalcDate, error) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for NewOrderCross.
func (m NewOrderCross) YieldRedemptionDate() (field.YieldRedemptionDate, error) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for NewOrderCross.
func (m NewOrderCross) YieldRedemptionPrice() (field.YieldRedemptionPrice, error) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for NewOrderCross.
func (m NewOrderCross) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, error) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for NewOrderCross.
func (m NewOrderCross) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ComplianceID is a non-required field for NewOrderCross.
func (m NewOrderCross) ComplianceID() (field.ComplianceID, error) {
	var f field.ComplianceID
	err := m.Body.Get(&f)
	return f, err
}

//IOIID is a non-required field for NewOrderCross.
func (m NewOrderCross) IOIID() (field.IOIID, error) {
	var f field.IOIID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for NewOrderCross.
func (m NewOrderCross) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//TimeInForce is a non-required field for NewOrderCross.
func (m NewOrderCross) TimeInForce() (field.TimeInForce, error) {
	var f field.TimeInForce
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for NewOrderCross.
func (m NewOrderCross) EffectiveTime() (field.EffectiveTime, error) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireDate is a non-required field for NewOrderCross.
func (m NewOrderCross) ExpireDate() (field.ExpireDate, error) {
	var f field.ExpireDate
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for NewOrderCross.
func (m NewOrderCross) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//GTBookingInst is a non-required field for NewOrderCross.
func (m NewOrderCross) GTBookingInst() (field.GTBookingInst, error) {
	var f field.GTBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//MaxShow is a non-required field for NewOrderCross.
func (m NewOrderCross) MaxShow() (field.MaxShow, error) {
	var f field.MaxShow
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetValue is a non-required field for NewOrderCross.
func (m NewOrderCross) PegOffsetValue() (field.PegOffsetValue, error) {
	var f field.PegOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//PegMoveType is a non-required field for NewOrderCross.
func (m NewOrderCross) PegMoveType() (field.PegMoveType, error) {
	var f field.PegMoveType
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetType is a non-required field for NewOrderCross.
func (m NewOrderCross) PegOffsetType() (field.PegOffsetType, error) {
	var f field.PegOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//PegLimitType is a non-required field for NewOrderCross.
func (m NewOrderCross) PegLimitType() (field.PegLimitType, error) {
	var f field.PegLimitType
	err := m.Body.Get(&f)
	return f, err
}

//PegRoundDirection is a non-required field for NewOrderCross.
func (m NewOrderCross) PegRoundDirection() (field.PegRoundDirection, error) {
	var f field.PegRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//PegScope is a non-required field for NewOrderCross.
func (m NewOrderCross) PegScope() (field.PegScope, error) {
	var f field.PegScope
	err := m.Body.Get(&f)
	return f, err
}

//PegPriceType is a non-required field for NewOrderCross.
func (m NewOrderCross) PegPriceType() (field.PegPriceType, error) {
	var f field.PegPriceType
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityIDSource is a non-required field for NewOrderCross.
func (m NewOrderCross) PegSecurityIDSource() (field.PegSecurityIDSource, error) {
	var f field.PegSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityID is a non-required field for NewOrderCross.
func (m NewOrderCross) PegSecurityID() (field.PegSecurityID, error) {
	var f field.PegSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//PegSymbol is a non-required field for NewOrderCross.
func (m NewOrderCross) PegSymbol() (field.PegSymbol, error) {
	var f field.PegSymbol
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityDesc is a non-required field for NewOrderCross.
func (m NewOrderCross) PegSecurityDesc() (field.PegSecurityDesc, error) {
	var f field.PegSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionInst() (field.DiscretionInst, error) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetValue is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionOffsetValue() (field.DiscretionOffsetValue, error) {
	var f field.DiscretionOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionMoveType is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionMoveType() (field.DiscretionMoveType, error) {
	var f field.DiscretionMoveType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetType is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionOffsetType() (field.DiscretionOffsetType, error) {
	var f field.DiscretionOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionLimitType is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionLimitType() (field.DiscretionLimitType, error) {
	var f field.DiscretionLimitType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionRoundDirection is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionRoundDirection() (field.DiscretionRoundDirection, error) {
	var f field.DiscretionRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionScope is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionScope() (field.DiscretionScope, error) {
	var f field.DiscretionScope
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategy is a non-required field for NewOrderCross.
func (m NewOrderCross) TargetStrategy() (field.TargetStrategy, error) {
	var f field.TargetStrategy
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategyParameters is a non-required field for NewOrderCross.
func (m NewOrderCross) TargetStrategyParameters() (field.TargetStrategyParameters, error) {
	var f field.TargetStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//ParticipationRate is a non-required field for NewOrderCross.
func (m NewOrderCross) ParticipationRate() (field.ParticipationRate, error) {
	var f field.ParticipationRate
	err := m.Body.Get(&f)
	return f, err
}

//CancellationRights is a non-required field for NewOrderCross.
func (m NewOrderCross) CancellationRights() (field.CancellationRights, error) {
	var f field.CancellationRights
	err := m.Body.Get(&f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for NewOrderCross.
func (m NewOrderCross) MoneyLaunderingStatus() (field.MoneyLaunderingStatus, error) {
	var f field.MoneyLaunderingStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistID is a non-required field for NewOrderCross.
func (m NewOrderCross) RegistID() (field.RegistID, error) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//Designation is a non-required field for NewOrderCross.
func (m NewOrderCross) Designation() (field.Designation, error) {
	var f field.Designation
	err := m.Body.Get(&f)
	return f, err
}

//NoStrategyParameters is a non-required field for NewOrderCross.
func (m NewOrderCross) NoStrategyParameters() (field.NoStrategyParameters, error) {
	var f field.NoStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//TransBkdTime is a non-required field for NewOrderCross.
func (m NewOrderCross) TransBkdTime() (field.TransBkdTime, error) {
	var f field.TransBkdTime
	err := m.Body.Get(&f)
	return f, err
}

//NoRootPartyIDs is a non-required field for NewOrderCross.
func (m NewOrderCross) NoRootPartyIDs() (field.NoRootPartyIDs, error) {
	var f field.NoRootPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//MatchIncrement is a non-required field for NewOrderCross.
func (m NewOrderCross) MatchIncrement() (field.MatchIncrement, error) {
	var f field.MatchIncrement
	err := m.Body.Get(&f)
	return f, err
}

//MaxPriceLevels is a non-required field for NewOrderCross.
func (m NewOrderCross) MaxPriceLevels() (field.MaxPriceLevels, error) {
	var f field.MaxPriceLevels
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryDisplayQty is a non-required field for NewOrderCross.
func (m NewOrderCross) SecondaryDisplayQty() (field.SecondaryDisplayQty, error) {
	var f field.SecondaryDisplayQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayWhen is a non-required field for NewOrderCross.
func (m NewOrderCross) DisplayWhen() (field.DisplayWhen, error) {
	var f field.DisplayWhen
	err := m.Body.Get(&f)
	return f, err
}

//DisplayMethod is a non-required field for NewOrderCross.
func (m NewOrderCross) DisplayMethod() (field.DisplayMethod, error) {
	var f field.DisplayMethod
	err := m.Body.Get(&f)
	return f, err
}

//DisplayLowQty is a non-required field for NewOrderCross.
func (m NewOrderCross) DisplayLowQty() (field.DisplayLowQty, error) {
	var f field.DisplayLowQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayHighQty is a non-required field for NewOrderCross.
func (m NewOrderCross) DisplayHighQty() (field.DisplayHighQty, error) {
	var f field.DisplayHighQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayMinIncr is a non-required field for NewOrderCross.
func (m NewOrderCross) DisplayMinIncr() (field.DisplayMinIncr, error) {
	var f field.DisplayMinIncr
	err := m.Body.Get(&f)
	return f, err
}

//RefreshQty is a non-required field for NewOrderCross.
func (m NewOrderCross) RefreshQty() (field.RefreshQty, error) {
	var f field.RefreshQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayQty is a non-required field for NewOrderCross.
func (m NewOrderCross) DisplayQty() (field.DisplayQty, error) {
	var f field.DisplayQty
	err := m.Body.Get(&f)
	return f, err
}

//PriceProtectionScope is a non-required field for NewOrderCross.
func (m NewOrderCross) PriceProtectionScope() (field.PriceProtectionScope, error) {
	var f field.PriceProtectionScope
	err := m.Body.Get(&f)
	return f, err
}

//TriggerType is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerType() (field.TriggerType, error) {
	var f field.TriggerType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerAction is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerAction() (field.TriggerAction, error) {
	var f field.TriggerAction
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPrice is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerPrice() (field.TriggerPrice, error) {
	var f field.TriggerPrice
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSymbol is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerSymbol() (field.TriggerSymbol, error) {
	var f field.TriggerSymbol
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityID is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerSecurityID() (field.TriggerSecurityID, error) {
	var f field.TriggerSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityIDSource is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerSecurityIDSource() (field.TriggerSecurityIDSource, error) {
	var f field.TriggerSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityDesc is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerSecurityDesc() (field.TriggerSecurityDesc, error) {
	var f field.TriggerSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceType is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerPriceType() (field.TriggerPriceType, error) {
	var f field.TriggerPriceType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceTypeScope is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerPriceTypeScope() (field.TriggerPriceTypeScope, error) {
	var f field.TriggerPriceTypeScope
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceDirection is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerPriceDirection() (field.TriggerPriceDirection, error) {
	var f field.TriggerPriceDirection
	err := m.Body.Get(&f)
	return f, err
}

//TriggerNewPrice is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerNewPrice() (field.TriggerNewPrice, error) {
	var f field.TriggerNewPrice
	err := m.Body.Get(&f)
	return f, err
}

//TriggerOrderType is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerOrderType() (field.TriggerOrderType, error) {
	var f field.TriggerOrderType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerNewQty is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerNewQty() (field.TriggerNewQty, error) {
	var f field.TriggerNewQty
	err := m.Body.Get(&f)
	return f, err
}

//TriggerTradingSessionID is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerTradingSessionID() (field.TriggerTradingSessionID, error) {
	var f field.TriggerTradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TriggerTradingSessionSubID is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerTradingSessionSubID() (field.TriggerTradingSessionSubID, error) {
	var f field.TriggerTradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//ExDestinationIDSource is a non-required field for NewOrderCross.
func (m NewOrderCross) ExDestinationIDSource() (field.ExDestinationIDSource, error) {
	var f field.ExDestinationIDSource
	err := m.Body.Get(&f)
	return f, err
}
