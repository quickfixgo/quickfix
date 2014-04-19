package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderMassCancelReport msg type = r.
type OrderMassCancelReport struct {
	message.Message
}

//OrderMassCancelReportBuilder builds OrderMassCancelReport messages.
type OrderMassCancelReportBuilder struct {
	message.MessageBuilder
}

//CreateOrderMassCancelReportBuilder returns an initialized OrderMassCancelReportBuilder with specified required fields.
func CreateOrderMassCancelReportBuilder(
	orderid field.OrderID,
	masscancelrequesttype field.MassCancelRequestType,
	masscancelresponse field.MassCancelResponse,
	massactionreportid field.MassActionReportID) OrderMassCancelReportBuilder {
	var builder OrderMassCancelReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(orderid)
	builder.Body.Set(masscancelrequesttype)
	builder.Body.Set(masscancelresponse)
	builder.Body.Set(massactionreportid)
	return builder
}

//ClOrdID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ClOrdID() (field.ClOrdID, errors.MessageRejectError) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecondaryClOrdID() (field.SecondaryClOrdID, errors.MessageRejectError) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a required field for OrderMassCancelReport.
func (m OrderMassCancelReport) OrderID() (field.OrderID, errors.MessageRejectError) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecondaryOrderID() (field.SecondaryOrderID, errors.MessageRejectError) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//MassCancelRequestType is a required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MassCancelRequestType() (field.MassCancelRequestType, errors.MessageRejectError) {
	var f field.MassCancelRequestType
	err := m.Body.Get(&f)
	return f, err
}

//MassCancelResponse is a required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MassCancelResponse() (field.MassCancelResponse, errors.MessageRejectError) {
	var f field.MassCancelResponse
	err := m.Body.Get(&f)
	return f, err
}

//MassCancelRejectReason is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MassCancelRejectReason() (field.MassCancelRejectReason, errors.MessageRejectError) {
	var f field.MassCancelRejectReason
	err := m.Body.Get(&f)
	return f, err
}

//TotalAffectedOrders is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TotalAffectedOrders() (field.TotalAffectedOrders, errors.MessageRejectError) {
	var f field.TotalAffectedOrders
	err := m.Body.Get(&f)
	return f, err
}

//NoAffectedOrders is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoAffectedOrders() (field.NoAffectedOrders, errors.MessageRejectError) {
	var f field.NoAffectedOrders
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TradingSessionID() (field.TradingSessionID, errors.MessageRejectError) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TradingSessionSubID() (field.TradingSessionSubID, errors.MessageRejectError) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecuritySubType() (field.SecuritySubType, errors.MessageRejectError) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikeCurrency() (field.StrikeCurrency, errors.MessageRejectError) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Pool() (field.Pool, errors.MessageRejectError) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ContractSettlMonth() (field.ContractSettlMonth, errors.MessageRejectError) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CPProgram() (field.CPProgram, errors.MessageRejectError) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CPRegType() (field.CPRegType, errors.MessageRejectError) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoEvents() (field.NoEvents, errors.MessageRejectError) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) DatedDate() (field.DatedDate, errors.MessageRejectError) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) InterestAccrualDate() (field.InterestAccrualDate, errors.MessageRejectError) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityStatus() (field.SecurityStatus, errors.MessageRejectError) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SettleOnOpenFlag() (field.SettleOnOpenFlag, errors.MessageRejectError) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikeMultiplier() (field.StrikeMultiplier, errors.MessageRejectError) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikeValue() (field.StrikeValue, errors.MessageRejectError) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MinPriceIncrement() (field.MinPriceIncrement, errors.MessageRejectError) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) PositionLimit() (field.PositionLimit, errors.MessageRejectError) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NTPositionLimit() (field.NTPositionLimit, errors.MessageRejectError) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoInstrumentParties() (field.NoInstrumentParties, errors.MessageRejectError) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnitOfMeasure() (field.UnitOfMeasure, errors.MessageRejectError) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TimeUnit() (field.TimeUnit, errors.MessageRejectError) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MaturityTime() (field.MaturityTime, errors.MessageRejectError) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityGroup() (field.SecurityGroup, errors.MessageRejectError) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, errors.MessageRejectError) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnitOfMeasureQty() (field.UnitOfMeasureQty, errors.MessageRejectError) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityXMLLen() (field.SecurityXMLLen, errors.MessageRejectError) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityXML() (field.SecurityXML, errors.MessageRejectError) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityXMLSchema() (field.SecurityXMLSchema, errors.MessageRejectError) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ProductComplex() (field.ProductComplex, errors.MessageRejectError) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SettlMethod() (field.SettlMethod, errors.MessageRejectError) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ExerciseStyle() (field.ExerciseStyle, errors.MessageRejectError) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) OptPayoutAmount() (field.OptPayoutAmount, errors.MessageRejectError) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) PriceQuoteMethod() (field.PriceQuoteMethod, errors.MessageRejectError) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ListMethod() (field.ListMethod, errors.MessageRejectError) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CapPrice() (field.CapPrice, errors.MessageRejectError) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) FloorPrice() (field.FloorPrice, errors.MessageRejectError) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) PutOrCall() (field.PutOrCall, errors.MessageRejectError) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) FlexibleIndicator() (field.FlexibleIndicator, errors.MessageRejectError) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ValuationMethod() (field.ValuationMethod, errors.MessageRejectError) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ContractMultiplierUnit() (field.ContractMultiplierUnit, errors.MessageRejectError) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) FlowScheduleType() (field.FlowScheduleType, errors.MessageRejectError) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RestructuringType() (field.RestructuringType, errors.MessageRejectError) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Seniority() (field.Seniority, errors.MessageRejectError) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) AttachmentPoint() (field.AttachmentPoint, errors.MessageRejectError) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) DetachmentPoint() (field.DetachmentPoint, errors.MessageRejectError) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) OptPayoutType() (field.OptPayoutType, errors.MessageRejectError) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoComplexEvents() (field.NoComplexEvents, errors.MessageRejectError) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbol is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSymbol() (field.UnderlyingSymbol, errors.MessageRejectError) {
	var f field.UnderlyingSymbol
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbolSfx is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSymbolSfx() (field.UnderlyingSymbolSfx, errors.MessageRejectError) {
	var f field.UnderlyingSymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityID() (field.UnderlyingSecurityID, errors.MessageRejectError) {
	var f field.UnderlyingSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityIDSource is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityIDSource() (field.UnderlyingSecurityIDSource, errors.MessageRejectError) {
	var f field.UnderlyingSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingSecurityAltID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoUnderlyingSecurityAltID() (field.NoUnderlyingSecurityAltID, errors.MessageRejectError) {
	var f field.NoUnderlyingSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingProduct is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingProduct() (field.UnderlyingProduct, errors.MessageRejectError) {
	var f field.UnderlyingProduct
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCFICode is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCFICode() (field.UnderlyingCFICode, errors.MessageRejectError) {
	var f field.UnderlyingCFICode
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityType() (field.UnderlyingSecurityType, errors.MessageRejectError) {
	var f field.UnderlyingSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecuritySubType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecuritySubType() (field.UnderlyingSecuritySubType, errors.MessageRejectError) {
	var f field.UnderlyingSecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityMonthYear is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingMaturityMonthYear() (field.UnderlyingMaturityMonthYear, errors.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingMaturityDate() (field.UnderlyingMaturityDate, errors.MessageRejectError) {
	var f field.UnderlyingMaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponPaymentDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCouponPaymentDate() (field.UnderlyingCouponPaymentDate, errors.MessageRejectError) {
	var f field.UnderlyingCouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssueDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingIssueDate() (field.UnderlyingIssueDate, errors.MessageRejectError) {
	var f field.UnderlyingIssueDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepoCollateralSecurityType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRepoCollateralSecurityType() (field.UnderlyingRepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.UnderlyingRepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseTerm is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRepurchaseTerm() (field.UnderlyingRepurchaseTerm, errors.MessageRejectError) {
	var f field.UnderlyingRepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRepurchaseRate() (field.UnderlyingRepurchaseRate, errors.MessageRejectError) {
	var f field.UnderlyingRepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFactor is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingFactor() (field.UnderlyingFactor, errors.MessageRejectError) {
	var f field.UnderlyingFactor
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCreditRating is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCreditRating() (field.UnderlyingCreditRating, errors.MessageRejectError) {
	var f field.UnderlyingCreditRating
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingInstrRegistry is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingInstrRegistry() (field.UnderlyingInstrRegistry, errors.MessageRejectError) {
	var f field.UnderlyingInstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCountryOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCountryOfIssue() (field.UnderlyingCountryOfIssue, errors.MessageRejectError) {
	var f field.UnderlyingCountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingStateOrProvinceOfIssue() (field.UnderlyingStateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.UnderlyingStateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingLocaleOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingLocaleOfIssue() (field.UnderlyingLocaleOfIssue, errors.MessageRejectError) {
	var f field.UnderlyingLocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRedemptionDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRedemptionDate() (field.UnderlyingRedemptionDate, errors.MessageRejectError) {
	var f field.UnderlyingRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikePrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingStrikePrice() (field.UnderlyingStrikePrice, errors.MessageRejectError) {
	var f field.UnderlyingStrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikeCurrency is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingStrikeCurrency() (field.UnderlyingStrikeCurrency, errors.MessageRejectError) {
	var f field.UnderlyingStrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingOptAttribute is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingOptAttribute() (field.UnderlyingOptAttribute, errors.MessageRejectError) {
	var f field.UnderlyingOptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingContractMultiplier is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingContractMultiplier() (field.UnderlyingContractMultiplier, errors.MessageRejectError) {
	var f field.UnderlyingContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCouponRate() (field.UnderlyingCouponRate, errors.MessageRejectError) {
	var f field.UnderlyingCouponRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityExchange is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityExchange() (field.UnderlyingSecurityExchange, errors.MessageRejectError) {
	var f field.UnderlyingSecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssuer is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingIssuer() (field.UnderlyingIssuer, errors.MessageRejectError) {
	var f field.UnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuerLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedUnderlyingIssuerLen() (field.EncodedUnderlyingIssuerLen, errors.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuer is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedUnderlyingIssuer() (field.EncodedUnderlyingIssuer, errors.MessageRejectError) {
	var f field.EncodedUnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityDesc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityDesc() (field.UnderlyingSecurityDesc, errors.MessageRejectError) {
	var f field.UnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDescLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedUnderlyingSecurityDescLen() (field.EncodedUnderlyingSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDesc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedUnderlyingSecurityDesc() (field.EncodedUnderlyingSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPProgram is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCPProgram() (field.UnderlyingCPProgram, errors.MessageRejectError) {
	var f field.UnderlyingCPProgram
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPRegType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCPRegType() (field.UnderlyingCPRegType, errors.MessageRejectError) {
	var f field.UnderlyingCPRegType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrency is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCurrency() (field.UnderlyingCurrency, errors.MessageRejectError) {
	var f field.UnderlyingCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingQty is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingQty() (field.UnderlyingQty, errors.MessageRejectError) {
	var f field.UnderlyingQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPx is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingPx() (field.UnderlyingPx, errors.MessageRejectError) {
	var f field.UnderlyingPx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingDirtyPrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingDirtyPrice() (field.UnderlyingDirtyPrice, errors.MessageRejectError) {
	var f field.UnderlyingDirtyPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndPrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingEndPrice() (field.UnderlyingEndPrice, errors.MessageRejectError) {
	var f field.UnderlyingEndPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStartValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingStartValue() (field.UnderlyingStartValue, errors.MessageRejectError) {
	var f field.UnderlyingStartValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrentValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCurrentValue() (field.UnderlyingCurrentValue, errors.MessageRejectError) {
	var f field.UnderlyingCurrentValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingEndValue() (field.UnderlyingEndValue, errors.MessageRejectError) {
	var f field.UnderlyingEndValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingStips is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoUnderlyingStips() (field.NoUnderlyingStips, errors.MessageRejectError) {
	var f field.NoUnderlyingStips
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAllocationPercent is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingAllocationPercent() (field.UnderlyingAllocationPercent, errors.MessageRejectError) {
	var f field.UnderlyingAllocationPercent
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlementType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSettlementType() (field.UnderlyingSettlementType, errors.MessageRejectError) {
	var f field.UnderlyingSettlementType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashAmount is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCashAmount() (field.UnderlyingCashAmount, errors.MessageRejectError) {
	var f field.UnderlyingCashAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCashType() (field.UnderlyingCashType, errors.MessageRejectError) {
	var f field.UnderlyingCashType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasure is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingUnitOfMeasure() (field.UnderlyingUnitOfMeasure, errors.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTimeUnit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingTimeUnit() (field.UnderlyingTimeUnit, errors.MessageRejectError) {
	var f field.UnderlyingTimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCapValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCapValue() (field.UnderlyingCapValue, errors.MessageRejectError) {
	var f field.UnderlyingCapValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUndlyInstrumentParties is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoUndlyInstrumentParties() (field.NoUndlyInstrumentParties, errors.MessageRejectError) {
	var f field.NoUndlyInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSettlMethod() (field.UnderlyingSettlMethod, errors.MessageRejectError) {
	var f field.UnderlyingSettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAdjustedQuantity is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingAdjustedQuantity() (field.UnderlyingAdjustedQuantity, errors.MessageRejectError) {
	var f field.UnderlyingAdjustedQuantity
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingFXRate() (field.UnderlyingFXRate, errors.MessageRejectError) {
	var f field.UnderlyingFXRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRateCalc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingFXRateCalc() (field.UnderlyingFXRateCalc, errors.MessageRejectError) {
	var f field.UnderlyingFXRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityTime is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingMaturityTime() (field.UnderlyingMaturityTime, errors.MessageRejectError) {
	var f field.UnderlyingMaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPutOrCall is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingPutOrCall() (field.UnderlyingPutOrCall, errors.MessageRejectError) {
	var f field.UnderlyingPutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingExerciseStyle is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingExerciseStyle() (field.UnderlyingExerciseStyle, errors.MessageRejectError) {
	var f field.UnderlyingExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasureQty is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingUnitOfMeasureQty() (field.UnderlyingUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceUnitOfMeasure is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingPriceUnitOfMeasure() (field.UnderlyingPriceUnitOfMeasure, errors.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingPriceUnitOfMeasureQty() (field.UnderlyingPriceUnitOfMeasureQty, errors.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingContractMultiplierUnit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingContractMultiplierUnit() (field.UnderlyingContractMultiplierUnit, errors.MessageRejectError) {
	var f field.UnderlyingContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFlowScheduleType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingFlowScheduleType() (field.UnderlyingFlowScheduleType, errors.MessageRejectError) {
	var f field.UnderlyingFlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRestructuringType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRestructuringType() (field.UnderlyingRestructuringType, errors.MessageRejectError) {
	var f field.UnderlyingRestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSeniority is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSeniority() (field.UnderlyingSeniority, errors.MessageRejectError) {
	var f field.UnderlyingSeniority
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingNotionalPercentageOutstanding is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingNotionalPercentageOutstanding() (field.UnderlyingNotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.UnderlyingNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingOriginalNotionalPercentageOutstanding is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingOriginalNotionalPercentageOutstanding() (field.UnderlyingOriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	var f field.UnderlyingOriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAttachmentPoint is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingAttachmentPoint() (field.UnderlyingAttachmentPoint, errors.MessageRejectError) {
	var f field.UnderlyingAttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingDetachmentPoint is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingDetachmentPoint() (field.UnderlyingDetachmentPoint, errors.MessageRejectError) {
	var f field.UnderlyingDetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Side() (field.Side, errors.MessageRejectError) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//MassActionReportID is a required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MassActionReportID() (field.MassActionReportID, errors.MessageRejectError) {
	var f field.MassActionReportID
	err := m.Body.Get(&f)
	return f, err
}

//NoNotAffectedOrders is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoNotAffectedOrders() (field.NoNotAffectedOrders, errors.MessageRejectError) {
	var f field.NoNotAffectedOrders
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MarketID() (field.MarketID, errors.MessageRejectError) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MarketSegmentID() (field.MarketSegmentID, errors.MessageRejectError) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}

//NoTargetPartyIDs is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoTargetPartyIDs() (field.NoTargetPartyIDs, errors.MessageRejectError) {
	var f field.NoTargetPartyIDs
	err := m.Body.Get(&f)
	return f, err
}
