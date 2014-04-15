package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderMassStatusRequest msg type = AF.
type OrderMassStatusRequest struct {
	message.Message
}

//OrderMassStatusRequestBuilder builds OrderMassStatusRequest messages.
type OrderMassStatusRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderMassStatusRequestBuilder returns an initialized OrderMassStatusRequestBuilder with specified required fields.
func CreateOrderMassStatusRequestBuilder(
	massstatusreqid field.MassStatusReqID,
	massstatusreqtype field.MassStatusReqType) OrderMassStatusRequestBuilder {
	var builder OrderMassStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(massstatusreqid)
	builder.Body.Set(massstatusreqtype)
	return builder
}

//MassStatusReqID is a required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MassStatusReqID() (field.MassStatusReqID, error) {
	var f field.MassStatusReqID
	err := m.Body.Get(&f)
	return f, err
}

//MassStatusReqType is a required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MassStatusReqType() (field.MassStatusReqType, error) {
	var f field.MassStatusReqType
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayAmount is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) OptPayAmount() (field.OptPayAmount, error) {
	var f field.OptPayAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FuturesValuationMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) FuturesValuationMethod() (field.FuturesValuationMethod, error) {
	var f field.FuturesValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbol is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSymbol() (field.UnderlyingSymbol, error) {
	var f field.UnderlyingSymbol
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbolSfx is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSymbolSfx() (field.UnderlyingSymbolSfx, error) {
	var f field.UnderlyingSymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityID is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSecurityID() (field.UnderlyingSecurityID, error) {
	var f field.UnderlyingSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityIDSource is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSecurityIDSource() (field.UnderlyingSecurityIDSource, error) {
	var f field.UnderlyingSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingSecurityAltID is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoUnderlyingSecurityAltID() (field.NoUnderlyingSecurityAltID, error) {
	var f field.NoUnderlyingSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingProduct is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingProduct() (field.UnderlyingProduct, error) {
	var f field.UnderlyingProduct
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCFICode is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCFICode() (field.UnderlyingCFICode, error) {
	var f field.UnderlyingCFICode
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSecurityType() (field.UnderlyingSecurityType, error) {
	var f field.UnderlyingSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecuritySubType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSecuritySubType() (field.UnderlyingSecuritySubType, error) {
	var f field.UnderlyingSecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityMonthYear is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingMaturityMonthYear() (field.UnderlyingMaturityMonthYear, error) {
	var f field.UnderlyingMaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingMaturityDate() (field.UnderlyingMaturityDate, error) {
	var f field.UnderlyingMaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponPaymentDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCouponPaymentDate() (field.UnderlyingCouponPaymentDate, error) {
	var f field.UnderlyingCouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssueDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingIssueDate() (field.UnderlyingIssueDate, error) {
	var f field.UnderlyingIssueDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepoCollateralSecurityType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingRepoCollateralSecurityType() (field.UnderlyingRepoCollateralSecurityType, error) {
	var f field.UnderlyingRepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseTerm is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingRepurchaseTerm() (field.UnderlyingRepurchaseTerm, error) {
	var f field.UnderlyingRepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseRate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingRepurchaseRate() (field.UnderlyingRepurchaseRate, error) {
	var f field.UnderlyingRepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFactor is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingFactor() (field.UnderlyingFactor, error) {
	var f field.UnderlyingFactor
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCreditRating is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCreditRating() (field.UnderlyingCreditRating, error) {
	var f field.UnderlyingCreditRating
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingInstrRegistry is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingInstrRegistry() (field.UnderlyingInstrRegistry, error) {
	var f field.UnderlyingInstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCountryOfIssue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCountryOfIssue() (field.UnderlyingCountryOfIssue, error) {
	var f field.UnderlyingCountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingStateOrProvinceOfIssue() (field.UnderlyingStateOrProvinceOfIssue, error) {
	var f field.UnderlyingStateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingLocaleOfIssue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingLocaleOfIssue() (field.UnderlyingLocaleOfIssue, error) {
	var f field.UnderlyingLocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRedemptionDate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingRedemptionDate() (field.UnderlyingRedemptionDate, error) {
	var f field.UnderlyingRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikePrice is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingStrikePrice() (field.UnderlyingStrikePrice, error) {
	var f field.UnderlyingStrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikeCurrency is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingStrikeCurrency() (field.UnderlyingStrikeCurrency, error) {
	var f field.UnderlyingStrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingOptAttribute is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingOptAttribute() (field.UnderlyingOptAttribute, error) {
	var f field.UnderlyingOptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingContractMultiplier is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingContractMultiplier() (field.UnderlyingContractMultiplier, error) {
	var f field.UnderlyingContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponRate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCouponRate() (field.UnderlyingCouponRate, error) {
	var f field.UnderlyingCouponRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityExchange is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSecurityExchange() (field.UnderlyingSecurityExchange, error) {
	var f field.UnderlyingSecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssuer is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingIssuer() (field.UnderlyingIssuer, error) {
	var f field.UnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuerLen is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedUnderlyingIssuerLen() (field.EncodedUnderlyingIssuerLen, error) {
	var f field.EncodedUnderlyingIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuer is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedUnderlyingIssuer() (field.EncodedUnderlyingIssuer, error) {
	var f field.EncodedUnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityDesc is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSecurityDesc() (field.UnderlyingSecurityDesc, error) {
	var f field.UnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDescLen is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedUnderlyingSecurityDescLen() (field.EncodedUnderlyingSecurityDescLen, error) {
	var f field.EncodedUnderlyingSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDesc is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) EncodedUnderlyingSecurityDesc() (field.EncodedUnderlyingSecurityDesc, error) {
	var f field.EncodedUnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPProgram is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCPProgram() (field.UnderlyingCPProgram, error) {
	var f field.UnderlyingCPProgram
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPRegType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCPRegType() (field.UnderlyingCPRegType, error) {
	var f field.UnderlyingCPRegType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrency is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCurrency() (field.UnderlyingCurrency, error) {
	var f field.UnderlyingCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingQty is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingQty() (field.UnderlyingQty, error) {
	var f field.UnderlyingQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPx is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingPx() (field.UnderlyingPx, error) {
	var f field.UnderlyingPx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingDirtyPrice is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingDirtyPrice() (field.UnderlyingDirtyPrice, error) {
	var f field.UnderlyingDirtyPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndPrice is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingEndPrice() (field.UnderlyingEndPrice, error) {
	var f field.UnderlyingEndPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStartValue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingStartValue() (field.UnderlyingStartValue, error) {
	var f field.UnderlyingStartValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrentValue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCurrentValue() (field.UnderlyingCurrentValue, error) {
	var f field.UnderlyingCurrentValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndValue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingEndValue() (field.UnderlyingEndValue, error) {
	var f field.UnderlyingEndValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingStips is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoUnderlyingStips() (field.NoUnderlyingStips, error) {
	var f field.NoUnderlyingStips
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAllocationPercent is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingAllocationPercent() (field.UnderlyingAllocationPercent, error) {
	var f field.UnderlyingAllocationPercent
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlementType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSettlementType() (field.UnderlyingSettlementType, error) {
	var f field.UnderlyingSettlementType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashAmount is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCashAmount() (field.UnderlyingCashAmount, error) {
	var f field.UnderlyingCashAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashType is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCashType() (field.UnderlyingCashType, error) {
	var f field.UnderlyingCashType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasure is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingUnitOfMeasure() (field.UnderlyingUnitOfMeasure, error) {
	var f field.UnderlyingUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTimeUnit is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingTimeUnit() (field.UnderlyingTimeUnit, error) {
	var f field.UnderlyingTimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCapValue is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingCapValue() (field.UnderlyingCapValue, error) {
	var f field.UnderlyingCapValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUndlyInstrumentParties is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) NoUndlyInstrumentParties() (field.NoUndlyInstrumentParties, error) {
	var f field.NoUndlyInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlMethod is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingSettlMethod() (field.UnderlyingSettlMethod, error) {
	var f field.UnderlyingSettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAdjustedQuantity is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingAdjustedQuantity() (field.UnderlyingAdjustedQuantity, error) {
	var f field.UnderlyingAdjustedQuantity
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRate is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingFXRate() (field.UnderlyingFXRate, error) {
	var f field.UnderlyingFXRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRateCalc is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingFXRateCalc() (field.UnderlyingFXRateCalc, error) {
	var f field.UnderlyingFXRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityTime is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingMaturityTime() (field.UnderlyingMaturityTime, error) {
	var f field.UnderlyingMaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPutOrCall is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingPutOrCall() (field.UnderlyingPutOrCall, error) {
	var f field.UnderlyingPutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingExerciseStyle is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingExerciseStyle() (field.UnderlyingExerciseStyle, error) {
	var f field.UnderlyingExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasureQty is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingUnitOfMeasureQty() (field.UnderlyingUnitOfMeasureQty, error) {
	var f field.UnderlyingUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceUnitOfMeasure is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingPriceUnitOfMeasure() (field.UnderlyingPriceUnitOfMeasure, error) {
	var f field.UnderlyingPriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) UnderlyingPriceUnitOfMeasureQty() (field.UnderlyingPriceUnitOfMeasureQty, error) {
	var f field.UnderlyingPriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for OrderMassStatusRequest.
func (m OrderMassStatusRequest) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}
