package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AllocationReport msg type = AS.
type AllocationReport struct {
	message.Message
}

//AllocationReportBuilder builds AllocationReport messages.
type AllocationReportBuilder struct {
	message.MessageBuilder
}

//CreateAllocationReportBuilder returns an initialized AllocationReportBuilder with specified required fields.
func CreateAllocationReportBuilder(
	allocreportid field.AllocReportID,
	alloctranstype field.AllocTransType,
	allocreporttype field.AllocReportType,
	allocstatus field.AllocStatus,
	side field.Side,
	quantity field.Quantity,
	avgpx field.AvgPx,
	tradedate field.TradeDate) AllocationReportBuilder {
	var builder AllocationReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(allocreportid)
	builder.Body.Set(alloctranstype)
	builder.Body.Set(allocreporttype)
	builder.Body.Set(allocstatus)
	builder.Body.Set(side)
	builder.Body.Set(quantity)
	builder.Body.Set(avgpx)
	builder.Body.Set(tradedate)
	return builder
}

//AllocReportID is a required field for AllocationReport.
func (m AllocationReport) AllocReportID() (field.AllocReportID, error) {
	var f field.AllocReportID
	err := m.Body.Get(&f)
	return f, err
}

//AllocID is a non-required field for AllocationReport.
func (m AllocationReport) AllocID() (field.AllocID, error) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//AllocTransType is a required field for AllocationReport.
func (m AllocationReport) AllocTransType() (field.AllocTransType, error) {
	var f field.AllocTransType
	err := m.Body.Get(&f)
	return f, err
}

//AllocReportRefID is a non-required field for AllocationReport.
func (m AllocationReport) AllocReportRefID() (field.AllocReportRefID, error) {
	var f field.AllocReportRefID
	err := m.Body.Get(&f)
	return f, err
}

//AllocCancReplaceReason is a non-required field for AllocationReport.
func (m AllocationReport) AllocCancReplaceReason() (field.AllocCancReplaceReason, error) {
	var f field.AllocCancReplaceReason
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryAllocID is a non-required field for AllocationReport.
func (m AllocationReport) SecondaryAllocID() (field.SecondaryAllocID, error) {
	var f field.SecondaryAllocID
	err := m.Body.Get(&f)
	return f, err
}

//AllocReportType is a required field for AllocationReport.
func (m AllocationReport) AllocReportType() (field.AllocReportType, error) {
	var f field.AllocReportType
	err := m.Body.Get(&f)
	return f, err
}

//AllocStatus is a required field for AllocationReport.
func (m AllocationReport) AllocStatus() (field.AllocStatus, error) {
	var f field.AllocStatus
	err := m.Body.Get(&f)
	return f, err
}

//AllocRejCode is a non-required field for AllocationReport.
func (m AllocationReport) AllocRejCode() (field.AllocRejCode, error) {
	var f field.AllocRejCode
	err := m.Body.Get(&f)
	return f, err
}

//RefAllocID is a non-required field for AllocationReport.
func (m AllocationReport) RefAllocID() (field.RefAllocID, error) {
	var f field.RefAllocID
	err := m.Body.Get(&f)
	return f, err
}

//AllocIntermedReqType is a non-required field for AllocationReport.
func (m AllocationReport) AllocIntermedReqType() (field.AllocIntermedReqType, error) {
	var f field.AllocIntermedReqType
	err := m.Body.Get(&f)
	return f, err
}

//AllocLinkID is a non-required field for AllocationReport.
func (m AllocationReport) AllocLinkID() (field.AllocLinkID, error) {
	var f field.AllocLinkID
	err := m.Body.Get(&f)
	return f, err
}

//AllocLinkType is a non-required field for AllocationReport.
func (m AllocationReport) AllocLinkType() (field.AllocLinkType, error) {
	var f field.AllocLinkType
	err := m.Body.Get(&f)
	return f, err
}

//BookingRefID is a non-required field for AllocationReport.
func (m AllocationReport) BookingRefID() (field.BookingRefID, error) {
	var f field.BookingRefID
	err := m.Body.Get(&f)
	return f, err
}

//AllocNoOrdersType is a non-required field for AllocationReport.
func (m AllocationReport) AllocNoOrdersType() (field.AllocNoOrdersType, error) {
	var f field.AllocNoOrdersType
	err := m.Body.Get(&f)
	return f, err
}

//NoOrders is a non-required field for AllocationReport.
func (m AllocationReport) NoOrders() (field.NoOrders, error) {
	var f field.NoOrders
	err := m.Body.Get(&f)
	return f, err
}

//NoExecs is a non-required field for AllocationReport.
func (m AllocationReport) NoExecs() (field.NoExecs, error) {
	var f field.NoExecs
	err := m.Body.Get(&f)
	return f, err
}

//PreviouslyReported is a non-required field for AllocationReport.
func (m AllocationReport) PreviouslyReported() (field.PreviouslyReported, error) {
	var f field.PreviouslyReported
	err := m.Body.Get(&f)
	return f, err
}

//ReversalIndicator is a non-required field for AllocationReport.
func (m AllocationReport) ReversalIndicator() (field.ReversalIndicator, error) {
	var f field.ReversalIndicator
	err := m.Body.Get(&f)
	return f, err
}

//MatchType is a non-required field for AllocationReport.
func (m AllocationReport) MatchType() (field.MatchType, error) {
	var f field.MatchType
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for AllocationReport.
func (m AllocationReport) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for AllocationReport.
func (m AllocationReport) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for AllocationReport.
func (m AllocationReport) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for AllocationReport.
func (m AllocationReport) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for AllocationReport.
func (m AllocationReport) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for AllocationReport.
func (m AllocationReport) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for AllocationReport.
func (m AllocationReport) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for AllocationReport.
func (m AllocationReport) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for AllocationReport.
func (m AllocationReport) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for AllocationReport.
func (m AllocationReport) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for AllocationReport.
func (m AllocationReport) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for AllocationReport.
func (m AllocationReport) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for AllocationReport.
func (m AllocationReport) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for AllocationReport.
func (m AllocationReport) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for AllocationReport.
func (m AllocationReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for AllocationReport.
func (m AllocationReport) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for AllocationReport.
func (m AllocationReport) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for AllocationReport.
func (m AllocationReport) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for AllocationReport.
func (m AllocationReport) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for AllocationReport.
func (m AllocationReport) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for AllocationReport.
func (m AllocationReport) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for AllocationReport.
func (m AllocationReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for AllocationReport.
func (m AllocationReport) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for AllocationReport.
func (m AllocationReport) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for AllocationReport.
func (m AllocationReport) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for AllocationReport.
func (m AllocationReport) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for AllocationReport.
func (m AllocationReport) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for AllocationReport.
func (m AllocationReport) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for AllocationReport.
func (m AllocationReport) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for AllocationReport.
func (m AllocationReport) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for AllocationReport.
func (m AllocationReport) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for AllocationReport.
func (m AllocationReport) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for AllocationReport.
func (m AllocationReport) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for AllocationReport.
func (m AllocationReport) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for AllocationReport.
func (m AllocationReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for AllocationReport.
func (m AllocationReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for AllocationReport.
func (m AllocationReport) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for AllocationReport.
func (m AllocationReport) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for AllocationReport.
func (m AllocationReport) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for AllocationReport.
func (m AllocationReport) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for AllocationReport.
func (m AllocationReport) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for AllocationReport.
func (m AllocationReport) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for AllocationReport.
func (m AllocationReport) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for AllocationReport.
func (m AllocationReport) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for AllocationReport.
func (m AllocationReport) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for AllocationReport.
func (m AllocationReport) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for AllocationReport.
func (m AllocationReport) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for AllocationReport.
func (m AllocationReport) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for AllocationReport.
func (m AllocationReport) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for AllocationReport.
func (m AllocationReport) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for AllocationReport.
func (m AllocationReport) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for AllocationReport.
func (m AllocationReport) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for AllocationReport.
func (m AllocationReport) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for AllocationReport.
func (m AllocationReport) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for AllocationReport.
func (m AllocationReport) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for AllocationReport.
func (m AllocationReport) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for AllocationReport.
func (m AllocationReport) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for AllocationReport.
func (m AllocationReport) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for AllocationReport.
func (m AllocationReport) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for AllocationReport.
func (m AllocationReport) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for AllocationReport.
func (m AllocationReport) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for AllocationReport.
func (m AllocationReport) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for AllocationReport.
func (m AllocationReport) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for AllocationReport.
func (m AllocationReport) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for AllocationReport.
func (m AllocationReport) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for AllocationReport.
func (m AllocationReport) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for AllocationReport.
func (m AllocationReport) OptPayoutAmount() (field.OptPayoutAmount, error) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for AllocationReport.
func (m AllocationReport) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for AllocationReport.
func (m AllocationReport) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for AllocationReport.
func (m AllocationReport) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for AllocationReport.
func (m AllocationReport) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for AllocationReport.
func (m AllocationReport) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for AllocationReport.
func (m AllocationReport) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for AllocationReport.
func (m AllocationReport) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for AllocationReport.
func (m AllocationReport) ValuationMethod() (field.ValuationMethod, error) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for AllocationReport.
func (m AllocationReport) ContractMultiplierUnit() (field.ContractMultiplierUnit, error) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for AllocationReport.
func (m AllocationReport) FlowScheduleType() (field.FlowScheduleType, error) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for AllocationReport.
func (m AllocationReport) RestructuringType() (field.RestructuringType, error) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for AllocationReport.
func (m AllocationReport) Seniority() (field.Seniority, error) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for AllocationReport.
func (m AllocationReport) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, error) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for AllocationReport.
func (m AllocationReport) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, error) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for AllocationReport.
func (m AllocationReport) AttachmentPoint() (field.AttachmentPoint, error) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for AllocationReport.
func (m AllocationReport) DetachmentPoint() (field.DetachmentPoint, error) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for AllocationReport.
func (m AllocationReport) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, error) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for AllocationReport.
func (m AllocationReport) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, error) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for AllocationReport.
func (m AllocationReport) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, error) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for AllocationReport.
func (m AllocationReport) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, error) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for AllocationReport.
func (m AllocationReport) OptPayoutType() (field.OptPayoutType, error) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for AllocationReport.
func (m AllocationReport) NoComplexEvents() (field.NoComplexEvents, error) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryForm is a non-required field for AllocationReport.
func (m AllocationReport) DeliveryForm() (field.DeliveryForm, error) {
	var f field.DeliveryForm
	err := m.Body.Get(&f)
	return f, err
}

//PctAtRisk is a non-required field for AllocationReport.
func (m AllocationReport) PctAtRisk() (field.PctAtRisk, error) {
	var f field.PctAtRisk
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrAttrib is a non-required field for AllocationReport.
func (m AllocationReport) NoInstrAttrib() (field.NoInstrAttrib, error) {
	var f field.NoInstrAttrib
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for AllocationReport.
func (m AllocationReport) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for AllocationReport.
func (m AllocationReport) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for AllocationReport.
func (m AllocationReport) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for AllocationReport.
func (m AllocationReport) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for AllocationReport.
func (m AllocationReport) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for AllocationReport.
func (m AllocationReport) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for AllocationReport.
func (m AllocationReport) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for AllocationReport.
func (m AllocationReport) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for AllocationReport.
func (m AllocationReport) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for AllocationReport.
func (m AllocationReport) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for AllocationReport.
func (m AllocationReport) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//Quantity is a required field for AllocationReport.
func (m AllocationReport) Quantity() (field.Quantity, error) {
	var f field.Quantity
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for AllocationReport.
func (m AllocationReport) QtyType() (field.QtyType, error) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for AllocationReport.
func (m AllocationReport) LastMkt() (field.LastMkt, error) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//TradeOriginationDate is a non-required field for AllocationReport.
func (m AllocationReport) TradeOriginationDate() (field.TradeOriginationDate, error) {
	var f field.TradeOriginationDate
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for AllocationReport.
func (m AllocationReport) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for AllocationReport.
func (m AllocationReport) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for AllocationReport.
func (m AllocationReport) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a required field for AllocationReport.
func (m AllocationReport) AvgPx() (field.AvgPx, error) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//AvgParPx is a non-required field for AllocationReport.
func (m AllocationReport) AvgParPx() (field.AvgParPx, error) {
	var f field.AvgParPx
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for AllocationReport.
func (m AllocationReport) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for AllocationReport.
func (m AllocationReport) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for AllocationReport.
func (m AllocationReport) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//AvgPxPrecision is a non-required field for AllocationReport.
func (m AllocationReport) AvgPxPrecision() (field.AvgPxPrecision, error) {
	var f field.AvgPxPrecision
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for AllocationReport.
func (m AllocationReport) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a required field for AllocationReport.
func (m AllocationReport) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for AllocationReport.
func (m AllocationReport) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for AllocationReport.
func (m AllocationReport) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for AllocationReport.
func (m AllocationReport) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//BookingType is a non-required field for AllocationReport.
func (m AllocationReport) BookingType() (field.BookingType, error) {
	var f field.BookingType
	err := m.Body.Get(&f)
	return f, err
}

//GrossTradeAmt is a non-required field for AllocationReport.
func (m AllocationReport) GrossTradeAmt() (field.GrossTradeAmt, error) {
	var f field.GrossTradeAmt
	err := m.Body.Get(&f)
	return f, err
}

//Concession is a non-required field for AllocationReport.
func (m AllocationReport) Concession() (field.Concession, error) {
	var f field.Concession
	err := m.Body.Get(&f)
	return f, err
}

//TotalTakedown is a non-required field for AllocationReport.
func (m AllocationReport) TotalTakedown() (field.TotalTakedown, error) {
	var f field.TotalTakedown
	err := m.Body.Get(&f)
	return f, err
}

//NetMoney is a non-required field for AllocationReport.
func (m AllocationReport) NetMoney() (field.NetMoney, error) {
	var f field.NetMoney
	err := m.Body.Get(&f)
	return f, err
}

//PositionEffect is a non-required field for AllocationReport.
func (m AllocationReport) PositionEffect() (field.PositionEffect, error) {
	var f field.PositionEffect
	err := m.Body.Get(&f)
	return f, err
}

//AutoAcceptIndicator is a non-required field for AllocationReport.
func (m AllocationReport) AutoAcceptIndicator() (field.AutoAcceptIndicator, error) {
	var f field.AutoAcceptIndicator
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for AllocationReport.
func (m AllocationReport) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for AllocationReport.
func (m AllocationReport) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for AllocationReport.
func (m AllocationReport) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NumDaysInterest is a non-required field for AllocationReport.
func (m AllocationReport) NumDaysInterest() (field.NumDaysInterest, error) {
	var f field.NumDaysInterest
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestRate is a non-required field for AllocationReport.
func (m AllocationReport) AccruedInterestRate() (field.AccruedInterestRate, error) {
	var f field.AccruedInterestRate
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestAmt is a non-required field for AllocationReport.
func (m AllocationReport) AccruedInterestAmt() (field.AccruedInterestAmt, error) {
	var f field.AccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//TotalAccruedInterestAmt is a non-required field for AllocationReport.
func (m AllocationReport) TotalAccruedInterestAmt() (field.TotalAccruedInterestAmt, error) {
	var f field.TotalAccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//InterestAtMaturity is a non-required field for AllocationReport.
func (m AllocationReport) InterestAtMaturity() (field.InterestAtMaturity, error) {
	var f field.InterestAtMaturity
	err := m.Body.Get(&f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for AllocationReport.
func (m AllocationReport) EndAccruedInterestAmt() (field.EndAccruedInterestAmt, error) {
	var f field.EndAccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//StartCash is a non-required field for AllocationReport.
func (m AllocationReport) StartCash() (field.StartCash, error) {
	var f field.StartCash
	err := m.Body.Get(&f)
	return f, err
}

//EndCash is a non-required field for AllocationReport.
func (m AllocationReport) EndCash() (field.EndCash, error) {
	var f field.EndCash
	err := m.Body.Get(&f)
	return f, err
}

//LegalConfirm is a non-required field for AllocationReport.
func (m AllocationReport) LegalConfirm() (field.LegalConfirm, error) {
	var f field.LegalConfirm
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for AllocationReport.
func (m AllocationReport) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for AllocationReport.
func (m AllocationReport) YieldType() (field.YieldType, error) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for AllocationReport.
func (m AllocationReport) Yield() (field.Yield, error) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for AllocationReport.
func (m AllocationReport) YieldCalcDate() (field.YieldCalcDate, error) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for AllocationReport.
func (m AllocationReport) YieldRedemptionDate() (field.YieldRedemptionDate, error) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for AllocationReport.
func (m AllocationReport) YieldRedemptionPrice() (field.YieldRedemptionPrice, error) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for AllocationReport.
func (m AllocationReport) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, error) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//TotNoAllocs is a non-required field for AllocationReport.
func (m AllocationReport) TotNoAllocs() (field.TotNoAllocs, error) {
	var f field.TotNoAllocs
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for AllocationReport.
func (m AllocationReport) LastFragment() (field.LastFragment, error) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for AllocationReport.
func (m AllocationReport) NoAllocs() (field.NoAllocs, error) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for AllocationReport.
func (m AllocationReport) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//TrdType is a non-required field for AllocationReport.
func (m AllocationReport) TrdType() (field.TrdType, error) {
	var f field.TrdType
	err := m.Body.Get(&f)
	return f, err
}

//TrdSubType is a non-required field for AllocationReport.
func (m AllocationReport) TrdSubType() (field.TrdSubType, error) {
	var f field.TrdSubType
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegReportingType is a non-required field for AllocationReport.
func (m AllocationReport) MultiLegReportingType() (field.MultiLegReportingType, error) {
	var f field.MultiLegReportingType
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for AllocationReport.
func (m AllocationReport) CustOrderCapacity() (field.CustOrderCapacity, error) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//TradeInputSource is a non-required field for AllocationReport.
func (m AllocationReport) TradeInputSource() (field.TradeInputSource, error) {
	var f field.TradeInputSource
	err := m.Body.Get(&f)
	return f, err
}

//RndPx is a non-required field for AllocationReport.
func (m AllocationReport) RndPx() (field.RndPx, error) {
	var f field.RndPx
	err := m.Body.Get(&f)
	return f, err
}

//MessageEventSource is a non-required field for AllocationReport.
func (m AllocationReport) MessageEventSource() (field.MessageEventSource, error) {
	var f field.MessageEventSource
	err := m.Body.Get(&f)
	return f, err
}

//TradeInputDevice is a non-required field for AllocationReport.
func (m AllocationReport) TradeInputDevice() (field.TradeInputDevice, error) {
	var f field.TradeInputDevice
	err := m.Body.Get(&f)
	return f, err
}

//AvgPxIndicator is a non-required field for AllocationReport.
func (m AllocationReport) AvgPxIndicator() (field.AvgPxIndicator, error) {
	var f field.AvgPxIndicator
	err := m.Body.Get(&f)
	return f, err
}

//NoPosAmt is a non-required field for AllocationReport.
func (m AllocationReport) NoPosAmt() (field.NoPosAmt, error) {
	var f field.NoPosAmt
	err := m.Body.Get(&f)
	return f, err
}

//NoRateSources is a non-required field for AllocationReport.
func (m AllocationReport) NoRateSources() (field.NoRateSources, error) {
	var f field.NoRateSources
	err := m.Body.Get(&f)
	return f, err
}
