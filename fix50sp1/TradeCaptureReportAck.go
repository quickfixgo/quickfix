package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradeCaptureReportAck msg type = AR.
type TradeCaptureReportAck struct {
	message.Message
}

//TradeCaptureReportAckBuilder builds TradeCaptureReportAck messages.
type TradeCaptureReportAckBuilder struct {
	message.MessageBuilder
}

//CreateTradeCaptureReportAckBuilder returns an initialized TradeCaptureReportAckBuilder with specified required fields.
func CreateTradeCaptureReportAckBuilder(
	nosides field.NoSides) TradeCaptureReportAckBuilder {
	var builder TradeCaptureReportAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(nosides)
	return builder
}

//TradeReportID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportID() (field.TradeReportID, error) {
	var f field.TradeReportID
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportTransType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportTransType() (field.TradeReportTransType, error) {
	var f field.TradeReportTransType
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportType() (field.TradeReportType, error) {
	var f field.TradeReportType
	err := m.Body.Get(&f)
	return f, err
}

//TrdType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdType() (field.TrdType, error) {
	var f field.TrdType
	err := m.Body.Get(&f)
	return f, err
}

//TrdSubType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdSubType() (field.TrdSubType, error) {
	var f field.TrdSubType
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTrdType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTrdType() (field.SecondaryTrdType, error) {
	var f field.SecondaryTrdType
	err := m.Body.Get(&f)
	return f, err
}

//TransferReason is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TransferReason() (field.TransferReason, error) {
	var f field.TransferReason
	err := m.Body.Get(&f)
	return f, err
}

//ExecType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ExecType() (field.ExecType, error) {
	var f field.ExecType
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportRefID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportRefID() (field.TradeReportRefID, error) {
	var f field.TradeReportRefID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeReportRefID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTradeReportRefID() (field.SecondaryTradeReportRefID, error) {
	var f field.SecondaryTradeReportRefID
	err := m.Body.Get(&f)
	return f, err
}

//TrdRptStatus is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdRptStatus() (field.TrdRptStatus, error) {
	var f field.TrdRptStatus
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportRejectReason is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeReportRejectReason() (field.TradeReportRejectReason, error) {
	var f field.TradeReportRejectReason
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTradeReportID() (field.SecondaryTradeReportID, error) {
	var f field.SecondaryTradeReportID
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//TradeLinkID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeLinkID() (field.TradeLinkID, error) {
	var f field.TradeLinkID
	err := m.Body.Get(&f)
	return f, err
}

//TrdMatchID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TrdMatchID() (field.TrdMatchID, error) {
	var f field.TrdMatchID
	err := m.Body.Get(&f)
	return f, err
}

//ExecID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ExecID() (field.ExecID, error) {
	var f field.ExecID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryExecID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryExecID() (field.SecondaryExecID, error) {
	var f field.SecondaryExecID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayAmount is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OptPayAmount() (field.OptPayAmount, error) {
	var f field.OptPayAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FuturesValuationMethod is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) FuturesValuationMethod() (field.FuturesValuationMethod, error) {
	var f field.FuturesValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, error) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//ResponseTransportType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ResponseTransportType() (field.ResponseTransportType, error) {
	var f field.ResponseTransportType
	err := m.Body.Get(&f)
	return f, err
}

//ResponseDestination is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ResponseDestination() (field.ResponseDestination, error) {
	var f field.ResponseDestination
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ClearingFeeIndicator() (field.ClearingFeeIndicator, error) {
	var f field.ClearingFeeIndicator
	err := m.Body.Get(&f)
	return f, err
}

//OrdStatus is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrdStatus() (field.OrdStatus, error) {
	var f field.OrdStatus
	err := m.Body.Get(&f)
	return f, err
}

//ExecRestatementReason is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ExecRestatementReason() (field.ExecRestatementReason, error) {
	var f field.ExecRestatementReason
	err := m.Body.Get(&f)
	return f, err
}

//PreviouslyReported is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PreviouslyReported() (field.PreviouslyReported, error) {
	var f field.PreviouslyReported
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTradingSessionID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) UnderlyingTradingSessionID() (field.UnderlyingTradingSessionID, error) {
	var f field.UnderlyingTradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) QtyType() (field.QtyType, error) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) UnderlyingTradingSessionSubID() (field.UnderlyingTradingSessionSubID, error) {
	var f field.UnderlyingTradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//LastQty is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastQty() (field.LastQty, error) {
	var f field.LastQty
	err := m.Body.Get(&f)
	return f, err
}

//LastPx is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastPx() (field.LastPx, error) {
	var f field.LastPx
	err := m.Body.Get(&f)
	return f, err
}

//LastParPx is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastParPx() (field.LastParPx, error) {
	var f field.LastParPx
	err := m.Body.Get(&f)
	return f, err
}

//LastSpotRate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastSpotRate() (field.LastSpotRate, error) {
	var f field.LastSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//LastForwardPoints is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastForwardPoints() (field.LastForwardPoints, error) {
	var f field.LastForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastMkt() (field.LastMkt, error) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) AvgPx() (field.AvgPx, error) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//AvgPxIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) AvgPxIndicator() (field.AvgPxIndicator, error) {
	var f field.AvgPxIndicator
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegReportingType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MultiLegReportingType() (field.MultiLegReportingType, error) {
	var f field.MultiLegReportingType
	err := m.Body.Get(&f)
	return f, err
}

//TradeLegRefID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeLegRefID() (field.TradeLegRefID, error) {
	var f field.TradeLegRefID
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//MatchStatus is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MatchStatus() (field.MatchStatus, error) {
	var f field.MatchStatus
	err := m.Body.Get(&f)
	return f, err
}

//MatchType is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MatchType() (field.MatchType, error) {
	var f field.MatchType
	err := m.Body.Get(&f)
	return f, err
}

//CopyMsgIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CopyMsgIndicator() (field.CopyMsgIndicator, error) {
	var f field.CopyMsgIndicator
	err := m.Body.Get(&f)
	return f, err
}

//PublishTrdIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) PublishTrdIndicator() (field.PublishTrdIndicator, error) {
	var f field.PublishTrdIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ShortSaleReason is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) ShortSaleReason() (field.ShortSaleReason, error) {
	var f field.ShortSaleReason
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SettlSessID() (field.SettlSessID, error) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SettlSessSubID() (field.SettlSessSubID, error) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//NoPosAmt is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoPosAmt() (field.NoPosAmt, error) {
	var f field.NoPosAmt
	err := m.Body.Get(&f)
	return f, err
}

//TierCode is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TierCode() (field.TierCode, error) {
	var f field.TierCode
	err := m.Body.Get(&f)
	return f, err
}

//MessageEventSource is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) MessageEventSource() (field.MessageEventSource, error) {
	var f field.MessageEventSource
	err := m.Body.Get(&f)
	return f, err
}

//LastUpdateTime is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastUpdateTime() (field.LastUpdateTime, error) {
	var f field.LastUpdateTime
	err := m.Body.Get(&f)
	return f, err
}

//RndPx is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RndPx() (field.RndPx, error) {
	var f field.RndPx
	err := m.Body.Get(&f)
	return f, err
}

//NoSides is a required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoSides() (field.NoSides, error) {
	var f field.NoSides
	err := m.Body.Get(&f)
	return f, err
}

//AsOfIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) AsOfIndicator() (field.AsOfIndicator, error) {
	var f field.AsOfIndicator
	err := m.Body.Get(&f)
	return f, err
}

//TradeID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeID() (field.TradeID, error) {
	var f field.TradeID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryTradeID() (field.SecondaryTradeID, error) {
	var f field.SecondaryTradeID
	err := m.Body.Get(&f)
	return f, err
}

//FirmTradeID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) FirmTradeID() (field.FirmTradeID, error) {
	var f field.FirmTradeID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SecondaryFirmTradeID() (field.SecondaryFirmTradeID, error) {
	var f field.SecondaryFirmTradeID
	err := m.Body.Get(&f)
	return f, err
}

//CalculatedCcyLastQty is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) CalculatedCcyLastQty() (field.CalculatedCcyLastQty, error) {
	var f field.CalculatedCcyLastQty
	err := m.Body.Get(&f)
	return f, err
}

//LastSwapPoints is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) LastSwapPoints() (field.LastSwapPoints, error) {
	var f field.LastSwapPoints
	err := m.Body.Get(&f)
	return f, err
}

//GrossTradeAmt is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) GrossTradeAmt() (field.GrossTradeAmt, error) {
	var f field.GrossTradeAmt
	err := m.Body.Get(&f)
	return f, err
}

//NoRootPartyIDs is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoRootPartyIDs() (field.NoRootPartyIDs, error) {
	var f field.NoRootPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradeHandlingInstr is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradeHandlingInstr() (field.TradeHandlingInstr, error) {
	var f field.TradeHandlingInstr
	err := m.Body.Get(&f)
	return f, err
}

//OrigTradeHandlingInstr is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrigTradeHandlingInstr() (field.OrigTradeHandlingInstr, error) {
	var f field.OrigTradeHandlingInstr
	err := m.Body.Get(&f)
	return f, err
}

//OrigTradeDate is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrigTradeDate() (field.OrigTradeDate, error) {
	var f field.OrigTradeDate
	err := m.Body.Get(&f)
	return f, err
}

//OrigTradeID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrigTradeID() (field.OrigTradeID, error) {
	var f field.OrigTradeID
	err := m.Body.Get(&f)
	return f, err
}

//OrigSecondaryTradeID is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) OrigSecondaryTradeID() (field.OrigSecondaryTradeID, error) {
	var f field.OrigSecondaryTradeID
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//RptSys is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) RptSys() (field.RptSys, error) {
	var f field.RptSys
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) SettlCurrency() (field.SettlCurrency, error) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//FeeMultiplier is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) FeeMultiplier() (field.FeeMultiplier, error) {
	var f field.FeeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRepIndicators is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) NoTrdRepIndicators() (field.NoTrdRepIndicators, error) {
	var f field.NoTrdRepIndicators
	err := m.Body.Get(&f)
	return f, err
}

//TradePublishIndicator is a non-required field for TradeCaptureReportAck.
func (m TradeCaptureReportAck) TradePublishIndicator() (field.TradePublishIndicator, error) {
	var f field.TradePublishIndicator
	err := m.Body.Get(&f)
	return f, err
}
