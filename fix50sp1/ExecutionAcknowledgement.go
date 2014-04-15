package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ExecutionAcknowledgement msg type = BN.
type ExecutionAcknowledgement struct {
	message.Message
}

//ExecutionAcknowledgementBuilder builds ExecutionAcknowledgement messages.
type ExecutionAcknowledgementBuilder struct {
	message.MessageBuilder
}

//CreateExecutionAcknowledgementBuilder returns an initialized ExecutionAcknowledgementBuilder with specified required fields.
func CreateExecutionAcknowledgementBuilder(
	orderid field.OrderID,
	execackstatus field.ExecAckStatus,
	execid field.ExecID,
	side field.Side) ExecutionAcknowledgementBuilder {
	var builder ExecutionAcknowledgementBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(orderid)
	builder.Body.Set(execackstatus)
	builder.Body.Set(execid)
	builder.Body.Set(side)
	return builder
}

//OrderID is a required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecondaryOrderID() (field.SecondaryOrderID, error) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ExecAckStatus is a required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ExecAckStatus() (field.ExecAckStatus, error) {
	var f field.ExecAckStatus
	err := m.Body.Get(&f)
	return f, err
}

//ExecID is a required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ExecID() (field.ExecID, error) {
	var f field.ExecID
	err := m.Body.Get(&f)
	return f, err
}

//DKReason is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) DKReason() (field.DKReason, error) {
	var f field.DKReason
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayAmount is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) OptPayAmount() (field.OptPayAmount, error) {
	var f field.OptPayAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FuturesValuationMethod is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) FuturesValuationMethod() (field.FuturesValuationMethod, error) {
	var f field.FuturesValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) OrderPercent() (field.OrderPercent, error) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) RoundingDirection() (field.RoundingDirection, error) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) RoundingModulus() (field.RoundingModulus, error) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//LastQty is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) LastQty() (field.LastQty, error) {
	var f field.LastQty
	err := m.Body.Get(&f)
	return f, err
}

//LastPx is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) LastPx() (field.LastPx, error) {
	var f field.LastPx
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//LastParPx is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) LastParPx() (field.LastParPx, error) {
	var f field.LastParPx
	err := m.Body.Get(&f)
	return f, err
}

//CumQty is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CumQty() (field.CumQty, error) {
	var f field.CumQty
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) AvgPx() (field.AvgPx, error) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
