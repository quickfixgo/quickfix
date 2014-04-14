package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderMassActionRequest msg type = CA.
type OrderMassActionRequest struct {
	message.Message
}

//OrderMassActionRequestBuilder builds OrderMassActionRequest messages.
type OrderMassActionRequestBuilder struct {
	message.MessageBuilder
}

//NewOrderMassActionRequestBuilder returns an initialized OrderMassActionRequestBuilder with specified required fields.
func NewOrderMassActionRequestBuilder(
	clordid field.ClOrdID,
	massactiontype field.MassActionType,
	massactionscope field.MassActionScope,
	transacttime field.TransactTime) *OrderMassActionRequestBuilder {
	builder := new(OrderMassActionRequestBuilder)
	builder.Body.Set(clordid)
	builder.Body.Set(massactiontype)
	builder.Body.Set(massactionscope)
	builder.Body.Set(transacttime)
	return builder
}

//ClOrdID is a required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryClOrdID is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//MassActionType is a required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) MassActionType() (*field.MassActionType, error) {
	f := new(field.MassActionType)
	err := m.Body.Get(f)
	return f, err
}

//MassActionScope is a required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) MassActionScope() (*field.MassActionScope, error) {
	f := new(field.MassActionScope)
	err := m.Body.Get(f)
	return f, err
}

//MarketID is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//MarketSegmentID is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//SecurityGroup is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLLen is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXML is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLSchema is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//ProductComplex is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SettlMethod is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//ExerciseStyle is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//OptPayoutAmount is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) OptPayoutAmount() (*field.OptPayoutAmount, error) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//PriceQuoteMethod is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//ListMethod is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//CapPrice is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//FloorPrice is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//FlexibleIndicator is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//ValuationMethod is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) ValuationMethod() (*field.ValuationMethod, error) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) ContractMultiplierUnit() (*field.ContractMultiplierUnit, error) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//FlowScheduleType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) FlowScheduleType() (*field.FlowScheduleType, error) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//RestructuringType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) RestructuringType() (*field.RestructuringType, error) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//Seniority is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) Seniority() (*field.Seniority, error) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, error) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, error) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//AttachmentPoint is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) AttachmentPoint() (*field.AttachmentPoint, error) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//DetachmentPoint is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) DetachmentPoint() (*field.DetachmentPoint, error) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, error) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, error) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, error) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, error) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//OptPayoutType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) OptPayoutType() (*field.OptPayoutType, error) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//NoComplexEvents is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) NoComplexEvents() (*field.NoComplexEvents, error) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSymbol is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingSymbol() (*field.UnderlyingSymbol, error) {
	f := new(field.UnderlyingSymbol)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSymbolSfx is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfx, error) {
	f := new(field.UnderlyingSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityID is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingSecurityID() (*field.UnderlyingSecurityID, error) {
	f := new(field.UnderlyingSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityIDSource is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSource, error) {
	f := new(field.UnderlyingSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyingSecurityAltID is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltID, error) {
	f := new(field.NoUnderlyingSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingProduct is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingProduct() (*field.UnderlyingProduct, error) {
	f := new(field.UnderlyingProduct)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCFICode is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingCFICode() (*field.UnderlyingCFICode, error) {
	f := new(field.UnderlyingCFICode)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingSecurityType() (*field.UnderlyingSecurityType, error) {
	f := new(field.UnderlyingSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecuritySubType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubType, error) {
	f := new(field.UnderlyingSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingMaturityMonthYear is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYear, error) {
	f := new(field.UnderlyingMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingMaturityDate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingMaturityDate() (*field.UnderlyingMaturityDate, error) {
	f := new(field.UnderlyingMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCouponPaymentDate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDate, error) {
	f := new(field.UnderlyingCouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingIssueDate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingIssueDate() (*field.UnderlyingIssueDate, error) {
	f := new(field.UnderlyingIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRepoCollateralSecurityType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityType, error) {
	f := new(field.UnderlyingRepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRepurchaseTerm is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTerm, error) {
	f := new(field.UnderlyingRepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRepurchaseRate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRate, error) {
	f := new(field.UnderlyingRepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFactor is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingFactor() (*field.UnderlyingFactor, error) {
	f := new(field.UnderlyingFactor)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCreditRating is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingCreditRating() (*field.UnderlyingCreditRating, error) {
	f := new(field.UnderlyingCreditRating)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingInstrRegistry is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistry, error) {
	f := new(field.UnderlyingInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCountryOfIssue is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssue, error) {
	f := new(field.UnderlyingCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssue, error) {
	f := new(field.UnderlyingStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingLocaleOfIssue is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssue, error) {
	f := new(field.UnderlyingLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRedemptionDate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDate, error) {
	f := new(field.UnderlyingRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStrikePrice is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingStrikePrice() (*field.UnderlyingStrikePrice, error) {
	f := new(field.UnderlyingStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStrikeCurrency is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrency, error) {
	f := new(field.UnderlyingStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingOptAttribute is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingOptAttribute() (*field.UnderlyingOptAttribute, error) {
	f := new(field.UnderlyingOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingContractMultiplier is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplier, error) {
	f := new(field.UnderlyingContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCouponRate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingCouponRate() (*field.UnderlyingCouponRate, error) {
	f := new(field.UnderlyingCouponRate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityExchange is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchange, error) {
	f := new(field.UnderlyingSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingIssuer is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingIssuer() (*field.UnderlyingIssuer, error) {
	f := new(field.UnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingIssuerLen is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLen, error) {
	f := new(field.EncodedUnderlyingIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingIssuer is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuer, error) {
	f := new(field.EncodedUnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSecurityDesc is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDesc, error) {
	f := new(field.UnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingSecurityDescLen is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLen, error) {
	f := new(field.EncodedUnderlyingSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedUnderlyingSecurityDesc is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDesc, error) {
	f := new(field.EncodedUnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCPProgram is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingCPProgram() (*field.UnderlyingCPProgram, error) {
	f := new(field.UnderlyingCPProgram)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCPRegType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingCPRegType() (*field.UnderlyingCPRegType, error) {
	f := new(field.UnderlyingCPRegType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCurrency is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingCurrency() (*field.UnderlyingCurrency, error) {
	f := new(field.UnderlyingCurrency)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingQty is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingQty() (*field.UnderlyingQty, error) {
	f := new(field.UnderlyingQty)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPx is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingPx() (*field.UnderlyingPx, error) {
	f := new(field.UnderlyingPx)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingDirtyPrice is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPrice, error) {
	f := new(field.UnderlyingDirtyPrice)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingEndPrice is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingEndPrice() (*field.UnderlyingEndPrice, error) {
	f := new(field.UnderlyingEndPrice)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingStartValue is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingStartValue() (*field.UnderlyingStartValue, error) {
	f := new(field.UnderlyingStartValue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCurrentValue is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingCurrentValue() (*field.UnderlyingCurrentValue, error) {
	f := new(field.UnderlyingCurrentValue)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingEndValue is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingEndValue() (*field.UnderlyingEndValue, error) {
	f := new(field.UnderlyingEndValue)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyingStips is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) NoUnderlyingStips() (*field.NoUnderlyingStips, error) {
	f := new(field.NoUnderlyingStips)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingAllocationPercent is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercent, error) {
	f := new(field.UnderlyingAllocationPercent)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSettlementType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingSettlementType() (*field.UnderlyingSettlementType, error) {
	f := new(field.UnderlyingSettlementType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCashAmount is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingCashAmount() (*field.UnderlyingCashAmount, error) {
	f := new(field.UnderlyingCashAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCashType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingCashType() (*field.UnderlyingCashType, error) {
	f := new(field.UnderlyingCashType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingUnitOfMeasure is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasure, error) {
	f := new(field.UnderlyingUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingTimeUnit is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingTimeUnit() (*field.UnderlyingTimeUnit, error) {
	f := new(field.UnderlyingTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingCapValue is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingCapValue() (*field.UnderlyingCapValue, error) {
	f := new(field.UnderlyingCapValue)
	err := m.Body.Get(f)
	return f, err
}

//NoUndlyInstrumentParties is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentParties, error) {
	f := new(field.NoUndlyInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSettlMethod is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingSettlMethod() (*field.UnderlyingSettlMethod, error) {
	f := new(field.UnderlyingSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingAdjustedQuantity is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantity, error) {
	f := new(field.UnderlyingAdjustedQuantity)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFXRate is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingFXRate() (*field.UnderlyingFXRate, error) {
	f := new(field.UnderlyingFXRate)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFXRateCalc is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalc, error) {
	f := new(field.UnderlyingFXRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingMaturityTime is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingMaturityTime() (*field.UnderlyingMaturityTime, error) {
	f := new(field.UnderlyingMaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPutOrCall is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingPutOrCall() (*field.UnderlyingPutOrCall, error) {
	f := new(field.UnderlyingPutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingExerciseStyle is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingExerciseStyle() (*field.UnderlyingExerciseStyle, error) {
	f := new(field.UnderlyingExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingUnitOfMeasureQty is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingUnitOfMeasureQty() (*field.UnderlyingUnitOfMeasureQty, error) {
	f := new(field.UnderlyingUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPriceUnitOfMeasure is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingPriceUnitOfMeasure() (*field.UnderlyingPriceUnitOfMeasure, error) {
	f := new(field.UnderlyingPriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingPriceUnitOfMeasureQty() (*field.UnderlyingPriceUnitOfMeasureQty, error) {
	f := new(field.UnderlyingPriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingContractMultiplierUnit is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingContractMultiplierUnit() (*field.UnderlyingContractMultiplierUnit, error) {
	f := new(field.UnderlyingContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingFlowScheduleType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingFlowScheduleType() (*field.UnderlyingFlowScheduleType, error) {
	f := new(field.UnderlyingFlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingRestructuringType is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingRestructuringType() (*field.UnderlyingRestructuringType, error) {
	f := new(field.UnderlyingRestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingSeniority is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingSeniority() (*field.UnderlyingSeniority, error) {
	f := new(field.UnderlyingSeniority)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingNotionalPercentageOutstanding is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingNotionalPercentageOutstanding() (*field.UnderlyingNotionalPercentageOutstanding, error) {
	f := new(field.UnderlyingNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingOriginalNotionalPercentageOutstanding is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingOriginalNotionalPercentageOutstanding() (*field.UnderlyingOriginalNotionalPercentageOutstanding, error) {
	f := new(field.UnderlyingOriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingAttachmentPoint is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingAttachmentPoint() (*field.UnderlyingAttachmentPoint, error) {
	f := new(field.UnderlyingAttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingDetachmentPoint is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) UnderlyingDetachmentPoint() (*field.UnderlyingDetachmentPoint, error) {
	f := new(field.UnderlyingDetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//Side is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//NoTargetPartyIDs is a non-required field for OrderMassActionRequest.
func (m *OrderMassActionRequest) NoTargetPartyIDs() (*field.NoTargetPartyIDs, error) {
	f := new(field.NoTargetPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
