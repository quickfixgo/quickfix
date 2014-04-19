package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CollateralInquiryAck msg type = BG.
type CollateralInquiryAck struct {
	message.Message
}

//CollateralInquiryAckBuilder builds CollateralInquiryAck messages.
type CollateralInquiryAckBuilder struct {
	message.MessageBuilder
}

//CreateCollateralInquiryAckBuilder returns an initialized CollateralInquiryAckBuilder with specified required fields.
func CreateCollateralInquiryAckBuilder(
	collinquiryid field.CollInquiryID,
	collinquirystatus field.CollInquiryStatus) CollateralInquiryAckBuilder {
	var builder CollateralInquiryAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(collinquiryid)
	builder.Body.Set(collinquirystatus)
	return builder
}

//CollInquiryID is a required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CollInquiryID() (field.CollInquiryID, errors.MessageRejectError) {
	var f field.CollInquiryID
	err := m.Body.Get(&f)
	return f, err
}

//CollInquiryStatus is a required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CollInquiryStatus() (field.CollInquiryStatus, errors.MessageRejectError) {
	var f field.CollInquiryStatus
	err := m.Body.Get(&f)
	return f, err
}

//CollInquiryResult is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CollInquiryResult() (field.CollInquiryResult, errors.MessageRejectError) {
	var f field.CollInquiryResult
	err := m.Body.Get(&f)
	return f, err
}

//NoCollInquiryQualifier is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoCollInquiryQualifier() (field.NoCollInquiryQualifier, errors.MessageRejectError) {
	var f field.NoCollInquiryQualifier
	err := m.Body.Get(&f)
	return f, err
}

//TotNumReports is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TotNumReports() (field.TotNumReports, errors.MessageRejectError) {
	var f field.TotNumReports
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Account() (field.Account, errors.MessageRejectError) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AccountType() (field.AccountType, errors.MessageRejectError) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ClOrdID() (field.ClOrdID, errors.MessageRejectError) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) OrderID() (field.OrderID, errors.MessageRejectError) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecondaryOrderID() (field.SecondaryOrderID, errors.MessageRejectError) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecondaryClOrdID() (field.SecondaryClOrdID, errors.MessageRejectError) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//NoExecs is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoExecs() (field.NoExecs, errors.MessageRejectError) {
	var f field.NoExecs
	err := m.Body.Get(&f)
	return f, err
}

//NoTrades is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoTrades() (field.NoTrades, errors.MessageRejectError) {
	var f field.NoTrades
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecuritySubType() (field.SecuritySubType, errors.MessageRejectError) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StrikeCurrency() (field.StrikeCurrency, errors.MessageRejectError) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Pool() (field.Pool, errors.MessageRejectError) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ContractSettlMonth() (field.ContractSettlMonth, errors.MessageRejectError) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CPProgram() (field.CPProgram, errors.MessageRejectError) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CPRegType() (field.CPRegType, errors.MessageRejectError) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoEvents() (field.NoEvents, errors.MessageRejectError) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) DatedDate() (field.DatedDate, errors.MessageRejectError) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) InterestAccrualDate() (field.InterestAccrualDate, errors.MessageRejectError) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityStatus() (field.SecurityStatus, errors.MessageRejectError) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SettleOnOpenFlag() (field.SettleOnOpenFlag, errors.MessageRejectError) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StrikeMultiplier() (field.StrikeMultiplier, errors.MessageRejectError) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StrikeValue() (field.StrikeValue, errors.MessageRejectError) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) MinPriceIncrement() (field.MinPriceIncrement, errors.MessageRejectError) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) PositionLimit() (field.PositionLimit, errors.MessageRejectError) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NTPositionLimit() (field.NTPositionLimit, errors.MessageRejectError) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoInstrumentParties() (field.NoInstrumentParties, errors.MessageRejectError) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) UnitOfMeasure() (field.UnitOfMeasure, errors.MessageRejectError) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TimeUnit() (field.TimeUnit, errors.MessageRejectError) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) MaturityTime() (field.MaturityTime, errors.MessageRejectError) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityGroup() (field.SecurityGroup, errors.MessageRejectError) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, errors.MessageRejectError) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) UnitOfMeasureQty() (field.UnitOfMeasureQty, errors.MessageRejectError) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityXMLLen() (field.SecurityXMLLen, errors.MessageRejectError) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityXML() (field.SecurityXML, errors.MessageRejectError) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityXMLSchema() (field.SecurityXMLSchema, errors.MessageRejectError) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ProductComplex() (field.ProductComplex, errors.MessageRejectError) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SettlMethod() (field.SettlMethod, errors.MessageRejectError) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ExerciseStyle() (field.ExerciseStyle, errors.MessageRejectError) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) OptPayoutAmount() (field.OptPayoutAmount, errors.MessageRejectError) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) PriceQuoteMethod() (field.PriceQuoteMethod, errors.MessageRejectError) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ListMethod() (field.ListMethod, errors.MessageRejectError) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CapPrice() (field.CapPrice, errors.MessageRejectError) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) FloorPrice() (field.FloorPrice, errors.MessageRejectError) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) PutOrCall() (field.PutOrCall, errors.MessageRejectError) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) FlexibleIndicator() (field.FlexibleIndicator, errors.MessageRejectError) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ValuationMethod() (field.ValuationMethod, errors.MessageRejectError) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ContractMultiplierUnit() (field.ContractMultiplierUnit, errors.MessageRejectError) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) FlowScheduleType() (field.FlowScheduleType, errors.MessageRejectError) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RestructuringType() (field.RestructuringType, errors.MessageRejectError) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Seniority() (field.Seniority, errors.MessageRejectError) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AttachmentPoint() (field.AttachmentPoint, errors.MessageRejectError) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) DetachmentPoint() (field.DetachmentPoint, errors.MessageRejectError) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) OptPayoutType() (field.OptPayoutType, errors.MessageRejectError) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoComplexEvents() (field.NoComplexEvents, errors.MessageRejectError) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AgreementDesc() (field.AgreementDesc, errors.MessageRejectError) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AgreementID() (field.AgreementID, errors.MessageRejectError) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AgreementDate() (field.AgreementDate, errors.MessageRejectError) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AgreementCurrency() (field.AgreementCurrency, errors.MessageRejectError) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TerminationType() (field.TerminationType, errors.MessageRejectError) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StartDate() (field.StartDate, errors.MessageRejectError) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EndDate() (field.EndDate, errors.MessageRejectError) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) DeliveryType() (field.DeliveryType, errors.MessageRejectError) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) MarginRatio() (field.MarginRatio, errors.MessageRejectError) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SettlDate() (field.SettlDate, errors.MessageRejectError) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//Quantity is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Quantity() (field.Quantity, errors.MessageRejectError) {
	var f field.Quantity
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) QtyType() (field.QtyType, errors.MessageRejectError) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoLegs() (field.NoLegs, errors.MessageRejectError) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TradingSessionID() (field.TradingSessionID, errors.MessageRejectError) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TradingSessionSubID() (field.TradingSessionSubID, errors.MessageRejectError) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SettlSessID() (field.SettlSessID, errors.MessageRejectError) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SettlSessSubID() (field.SettlSessSubID, errors.MessageRejectError) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ClearingBusinessDate() (field.ClearingBusinessDate, errors.MessageRejectError) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//ResponseTransportType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ResponseTransportType() (field.ResponseTransportType, errors.MessageRejectError) {
	var f field.ResponseTransportType
	err := m.Body.Get(&f)
	return f, err
}

//ResponseDestination is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ResponseDestination() (field.ResponseDestination, errors.MessageRejectError) {
	var f field.ResponseDestination
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
