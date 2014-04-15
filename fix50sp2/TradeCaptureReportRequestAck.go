package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradeCaptureReportRequestAck msg type = AQ.
type TradeCaptureReportRequestAck struct {
	message.Message
}

//TradeCaptureReportRequestAckBuilder builds TradeCaptureReportRequestAck messages.
type TradeCaptureReportRequestAckBuilder struct {
	message.MessageBuilder
}

//CreateTradeCaptureReportRequestAckBuilder returns an initialized TradeCaptureReportRequestAckBuilder with specified required fields.
func CreateTradeCaptureReportRequestAckBuilder(
	traderequestid field.TradeRequestID,
	traderequesttype field.TradeRequestType,
	traderequestresult field.TradeRequestResult,
	traderequeststatus field.TradeRequestStatus) TradeCaptureReportRequestAckBuilder {
	var builder TradeCaptureReportRequestAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(traderequestid)
	builder.Body.Set(traderequesttype)
	builder.Body.Set(traderequestresult)
	builder.Body.Set(traderequeststatus)
	return builder
}

//TradeRequestID is a required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TradeRequestID() (field.TradeRequestID, error) {
	var f field.TradeRequestID
	err := m.Body.Get(&f)
	return f, err
}

//TradeRequestType is a required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TradeRequestType() (field.TradeRequestType, error) {
	var f field.TradeRequestType
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//TotNumTradeReports is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TotNumTradeReports() (field.TotNumTradeReports, error) {
	var f field.TotNumTradeReports
	err := m.Body.Get(&f)
	return f, err
}

//TradeRequestResult is a required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TradeRequestResult() (field.TradeRequestResult, error) {
	var f field.TradeRequestResult
	err := m.Body.Get(&f)
	return f, err
}

//TradeRequestStatus is a required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TradeRequestStatus() (field.TradeRequestStatus, error) {
	var f field.TradeRequestStatus
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) OptPayoutAmount() (field.OptPayoutAmount, error) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ValuationMethod() (field.ValuationMethod, error) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ContractMultiplierUnit() (field.ContractMultiplierUnit, error) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) FlowScheduleType() (field.FlowScheduleType, error) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) RestructuringType() (field.RestructuringType, error) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Seniority() (field.Seniority, error) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, error) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, error) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) AttachmentPoint() (field.AttachmentPoint, error) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) DetachmentPoint() (field.DetachmentPoint, error) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, error) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, error) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, error) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, error) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) OptPayoutType() (field.OptPayoutType, error) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NoComplexEvents() (field.NoComplexEvents, error) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegReportingType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MultiLegReportingType() (field.MultiLegReportingType, error) {
	var f field.MultiLegReportingType
	err := m.Body.Get(&f)
	return f, err
}

//ResponseTransportType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ResponseTransportType() (field.ResponseTransportType, error) {
	var f field.ResponseTransportType
	err := m.Body.Get(&f)
	return f, err
}

//ResponseDestination is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ResponseDestination() (field.ResponseDestination, error) {
	var f field.ResponseDestination
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//MessageEventSource is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MessageEventSource() (field.MessageEventSource, error) {
	var f field.MessageEventSource
	err := m.Body.Get(&f)
	return f, err
}

//TradeID is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TradeID() (field.TradeID, error) {
	var f field.TradeID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeID is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecondaryTradeID() (field.SecondaryTradeID, error) {
	var f field.SecondaryTradeID
	err := m.Body.Get(&f)
	return f, err
}

//FirmTradeID is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) FirmTradeID() (field.FirmTradeID, error) {
	var f field.FirmTradeID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecondaryFirmTradeID() (field.SecondaryFirmTradeID, error) {
	var f field.SecondaryFirmTradeID
	err := m.Body.Get(&f)
	return f, err
}
