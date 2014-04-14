package fix50

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

//NewAllocationReportBuilder returns an initialized AllocationReportBuilder with specified required fields.
func NewAllocationReportBuilder(
	allocreportid field.AllocReportID,
	alloctranstype field.AllocTransType,
	allocreporttype field.AllocReportType,
	allocstatus field.AllocStatus,
	side field.Side,
	quantity field.Quantity,
	avgpx field.AvgPx,
	tradedate field.TradeDate) *AllocationReportBuilder {
	builder := new(AllocationReportBuilder)
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
func (m *AllocationReport) AllocReportID() (*field.AllocReportID, error) {
	f := new(field.AllocReportID)
	err := m.Body.Get(f)
	return f, err
}

//AllocID is a non-required field for AllocationReport.
func (m *AllocationReport) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//AllocTransType is a required field for AllocationReport.
func (m *AllocationReport) AllocTransType() (*field.AllocTransType, error) {
	f := new(field.AllocTransType)
	err := m.Body.Get(f)
	return f, err
}

//AllocReportRefID is a non-required field for AllocationReport.
func (m *AllocationReport) AllocReportRefID() (*field.AllocReportRefID, error) {
	f := new(field.AllocReportRefID)
	err := m.Body.Get(f)
	return f, err
}

//AllocCancReplaceReason is a non-required field for AllocationReport.
func (m *AllocationReport) AllocCancReplaceReason() (*field.AllocCancReplaceReason, error) {
	f := new(field.AllocCancReplaceReason)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryAllocID is a non-required field for AllocationReport.
func (m *AllocationReport) SecondaryAllocID() (*field.SecondaryAllocID, error) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}

//AllocReportType is a required field for AllocationReport.
func (m *AllocationReport) AllocReportType() (*field.AllocReportType, error) {
	f := new(field.AllocReportType)
	err := m.Body.Get(f)
	return f, err
}

//AllocStatus is a required field for AllocationReport.
func (m *AllocationReport) AllocStatus() (*field.AllocStatus, error) {
	f := new(field.AllocStatus)
	err := m.Body.Get(f)
	return f, err
}

//AllocRejCode is a non-required field for AllocationReport.
func (m *AllocationReport) AllocRejCode() (*field.AllocRejCode, error) {
	f := new(field.AllocRejCode)
	err := m.Body.Get(f)
	return f, err
}

//RefAllocID is a non-required field for AllocationReport.
func (m *AllocationReport) RefAllocID() (*field.RefAllocID, error) {
	f := new(field.RefAllocID)
	err := m.Body.Get(f)
	return f, err
}

//AllocIntermedReqType is a non-required field for AllocationReport.
func (m *AllocationReport) AllocIntermedReqType() (*field.AllocIntermedReqType, error) {
	f := new(field.AllocIntermedReqType)
	err := m.Body.Get(f)
	return f, err
}

//AllocLinkID is a non-required field for AllocationReport.
func (m *AllocationReport) AllocLinkID() (*field.AllocLinkID, error) {
	f := new(field.AllocLinkID)
	err := m.Body.Get(f)
	return f, err
}

//AllocLinkType is a non-required field for AllocationReport.
func (m *AllocationReport) AllocLinkType() (*field.AllocLinkType, error) {
	f := new(field.AllocLinkType)
	err := m.Body.Get(f)
	return f, err
}

//BookingRefID is a non-required field for AllocationReport.
func (m *AllocationReport) BookingRefID() (*field.BookingRefID, error) {
	f := new(field.BookingRefID)
	err := m.Body.Get(f)
	return f, err
}

//AllocNoOrdersType is a non-required field for AllocationReport.
func (m *AllocationReport) AllocNoOrdersType() (*field.AllocNoOrdersType, error) {
	f := new(field.AllocNoOrdersType)
	err := m.Body.Get(f)
	return f, err
}

//NoOrders is a non-required field for AllocationReport.
func (m *AllocationReport) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//NoExecs is a non-required field for AllocationReport.
func (m *AllocationReport) NoExecs() (*field.NoExecs, error) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//PreviouslyReported is a non-required field for AllocationReport.
func (m *AllocationReport) PreviouslyReported() (*field.PreviouslyReported, error) {
	f := new(field.PreviouslyReported)
	err := m.Body.Get(f)
	return f, err
}

//ReversalIndicator is a non-required field for AllocationReport.
func (m *AllocationReport) ReversalIndicator() (*field.ReversalIndicator, error) {
	f := new(field.ReversalIndicator)
	err := m.Body.Get(f)
	return f, err
}

//MatchType is a non-required field for AllocationReport.
func (m *AllocationReport) MatchType() (*field.MatchType, error) {
	f := new(field.MatchType)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for AllocationReport.
func (m *AllocationReport) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for AllocationReport.
func (m *AllocationReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for AllocationReport.
func (m *AllocationReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for AllocationReport.
func (m *AllocationReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for AllocationReport.
func (m *AllocationReport) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for AllocationReport.
func (m *AllocationReport) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for AllocationReport.
func (m *AllocationReport) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for AllocationReport.
func (m *AllocationReport) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for AllocationReport.
func (m *AllocationReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for AllocationReport.
func (m *AllocationReport) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for AllocationReport.
func (m *AllocationReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for AllocationReport.
func (m *AllocationReport) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for AllocationReport.
func (m *AllocationReport) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for AllocationReport.
func (m *AllocationReport) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for AllocationReport.
func (m *AllocationReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for AllocationReport.
func (m *AllocationReport) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for AllocationReport.
func (m *AllocationReport) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for AllocationReport.
func (m *AllocationReport) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for AllocationReport.
func (m *AllocationReport) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for AllocationReport.
func (m *AllocationReport) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for AllocationReport.
func (m *AllocationReport) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for AllocationReport.
func (m *AllocationReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for AllocationReport.
func (m *AllocationReport) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for AllocationReport.
func (m *AllocationReport) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for AllocationReport.
func (m *AllocationReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for AllocationReport.
func (m *AllocationReport) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for AllocationReport.
func (m *AllocationReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for AllocationReport.
func (m *AllocationReport) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for AllocationReport.
func (m *AllocationReport) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for AllocationReport.
func (m *AllocationReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for AllocationReport.
func (m *AllocationReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for AllocationReport.
func (m *AllocationReport) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for AllocationReport.
func (m *AllocationReport) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for AllocationReport.
func (m *AllocationReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for AllocationReport.
func (m *AllocationReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for AllocationReport.
func (m *AllocationReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for AllocationReport.
func (m *AllocationReport) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for AllocationReport.
func (m *AllocationReport) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for AllocationReport.
func (m *AllocationReport) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for AllocationReport.
func (m *AllocationReport) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for AllocationReport.
func (m *AllocationReport) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for AllocationReport.
func (m *AllocationReport) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for AllocationReport.
func (m *AllocationReport) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for AllocationReport.
func (m *AllocationReport) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for AllocationReport.
func (m *AllocationReport) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for AllocationReport.
func (m *AllocationReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for AllocationReport.
func (m *AllocationReport) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for AllocationReport.
func (m *AllocationReport) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for AllocationReport.
func (m *AllocationReport) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for AllocationReport.
func (m *AllocationReport) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for AllocationReport.
func (m *AllocationReport) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for AllocationReport.
func (m *AllocationReport) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for AllocationReport.
func (m *AllocationReport) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for AllocationReport.
func (m *AllocationReport) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for AllocationReport.
func (m *AllocationReport) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryForm is a non-required field for AllocationReport.
func (m *AllocationReport) DeliveryForm() (*field.DeliveryForm, error) {
	f := new(field.DeliveryForm)
	err := m.Body.Get(f)
	return f, err
}

//PctAtRisk is a non-required field for AllocationReport.
func (m *AllocationReport) PctAtRisk() (*field.PctAtRisk, error) {
	f := new(field.PctAtRisk)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrAttrib is a non-required field for AllocationReport.
func (m *AllocationReport) NoInstrAttrib() (*field.NoInstrAttrib, error) {
	f := new(field.NoInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDesc is a non-required field for AllocationReport.
func (m *AllocationReport) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//AgreementID is a non-required field for AllocationReport.
func (m *AllocationReport) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDate is a non-required field for AllocationReport.
func (m *AllocationReport) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementCurrency is a non-required field for AllocationReport.
func (m *AllocationReport) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//TerminationType is a non-required field for AllocationReport.
func (m *AllocationReport) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//StartDate is a non-required field for AllocationReport.
func (m *AllocationReport) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//EndDate is a non-required field for AllocationReport.
func (m *AllocationReport) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryType is a non-required field for AllocationReport.
func (m *AllocationReport) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//MarginRatio is a non-required field for AllocationReport.
func (m *AllocationReport) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for AllocationReport.
func (m *AllocationReport) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for AllocationReport.
func (m *AllocationReport) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//Quantity is a required field for AllocationReport.
func (m *AllocationReport) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//QtyType is a non-required field for AllocationReport.
func (m *AllocationReport) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//LastMkt is a non-required field for AllocationReport.
func (m *AllocationReport) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//TradeOriginationDate is a non-required field for AllocationReport.
func (m *AllocationReport) TradeOriginationDate() (*field.TradeOriginationDate, error) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for AllocationReport.
func (m *AllocationReport) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for AllocationReport.
func (m *AllocationReport) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for AllocationReport.
func (m *AllocationReport) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//AvgPx is a required field for AllocationReport.
func (m *AllocationReport) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//AvgParPx is a non-required field for AllocationReport.
func (m *AllocationReport) AvgParPx() (*field.AvgParPx, error) {
	f := new(field.AvgParPx)
	err := m.Body.Get(f)
	return f, err
}

//Spread is a non-required field for AllocationReport.
func (m *AllocationReport) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for AllocationReport.
func (m *AllocationReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveName is a non-required field for AllocationReport.
func (m *AllocationReport) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for AllocationReport.
func (m *AllocationReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPrice is a non-required field for AllocationReport.
func (m *AllocationReport) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPriceType is a non-required field for AllocationReport.
func (m *AllocationReport) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for AllocationReport.
func (m *AllocationReport) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for AllocationReport.
func (m *AllocationReport) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for AllocationReport.
func (m *AllocationReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//AvgPxPrecision is a non-required field for AllocationReport.
func (m *AllocationReport) AvgPxPrecision() (*field.AvgPxPrecision, error) {
	f := new(field.AvgPxPrecision)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for AllocationReport.
func (m *AllocationReport) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a required field for AllocationReport.
func (m *AllocationReport) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for AllocationReport.
func (m *AllocationReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//SettlType is a non-required field for AllocationReport.
func (m *AllocationReport) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate is a non-required field for AllocationReport.
func (m *AllocationReport) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//BookingType is a non-required field for AllocationReport.
func (m *AllocationReport) BookingType() (*field.BookingType, error) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//GrossTradeAmt is a non-required field for AllocationReport.
func (m *AllocationReport) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//Concession is a non-required field for AllocationReport.
func (m *AllocationReport) Concession() (*field.Concession, error) {
	f := new(field.Concession)
	err := m.Body.Get(f)
	return f, err
}

//TotalTakedown is a non-required field for AllocationReport.
func (m *AllocationReport) TotalTakedown() (*field.TotalTakedown, error) {
	f := new(field.TotalTakedown)
	err := m.Body.Get(f)
	return f, err
}

//NetMoney is a non-required field for AllocationReport.
func (m *AllocationReport) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}

//PositionEffect is a non-required field for AllocationReport.
func (m *AllocationReport) PositionEffect() (*field.PositionEffect, error) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//AutoAcceptIndicator is a non-required field for AllocationReport.
func (m *AllocationReport) AutoAcceptIndicator() (*field.AutoAcceptIndicator, error) {
	f := new(field.AutoAcceptIndicator)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for AllocationReport.
func (m *AllocationReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for AllocationReport.
func (m *AllocationReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for AllocationReport.
func (m *AllocationReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//NumDaysInterest is a non-required field for AllocationReport.
func (m *AllocationReport) NumDaysInterest() (*field.NumDaysInterest, error) {
	f := new(field.NumDaysInterest)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestRate is a non-required field for AllocationReport.
func (m *AllocationReport) AccruedInterestRate() (*field.AccruedInterestRate, error) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestAmt is a non-required field for AllocationReport.
func (m *AllocationReport) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//TotalAccruedInterestAmt is a non-required field for AllocationReport.
func (m *AllocationReport) TotalAccruedInterestAmt() (*field.TotalAccruedInterestAmt, error) {
	f := new(field.TotalAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//InterestAtMaturity is a non-required field for AllocationReport.
func (m *AllocationReport) InterestAtMaturity() (*field.InterestAtMaturity, error) {
	f := new(field.InterestAtMaturity)
	err := m.Body.Get(f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for AllocationReport.
func (m *AllocationReport) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, error) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//StartCash is a non-required field for AllocationReport.
func (m *AllocationReport) StartCash() (*field.StartCash, error) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}

//EndCash is a non-required field for AllocationReport.
func (m *AllocationReport) EndCash() (*field.EndCash, error) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}

//LegalConfirm is a non-required field for AllocationReport.
func (m *AllocationReport) LegalConfirm() (*field.LegalConfirm, error) {
	f := new(field.LegalConfirm)
	err := m.Body.Get(f)
	return f, err
}

//NoStipulations is a non-required field for AllocationReport.
func (m *AllocationReport) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//YieldType is a non-required field for AllocationReport.
func (m *AllocationReport) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//Yield is a non-required field for AllocationReport.
func (m *AllocationReport) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//YieldCalcDate is a non-required field for AllocationReport.
func (m *AllocationReport) YieldCalcDate() (*field.YieldCalcDate, error) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionDate is a non-required field for AllocationReport.
func (m *AllocationReport) YieldRedemptionDate() (*field.YieldRedemptionDate, error) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for AllocationReport.
func (m *AllocationReport) YieldRedemptionPrice() (*field.YieldRedemptionPrice, error) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for AllocationReport.
func (m *AllocationReport) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, error) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//TotNoAllocs is a non-required field for AllocationReport.
func (m *AllocationReport) TotNoAllocs() (*field.TotNoAllocs, error) {
	f := new(field.TotNoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//LastFragment is a non-required field for AllocationReport.
func (m *AllocationReport) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//NoAllocs is a non-required field for AllocationReport.
func (m *AllocationReport) NoAllocs() (*field.NoAllocs, error) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a non-required field for AllocationReport.
func (m *AllocationReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//TrdType is a non-required field for AllocationReport.
func (m *AllocationReport) TrdType() (*field.TrdType, error) {
	f := new(field.TrdType)
	err := m.Body.Get(f)
	return f, err
}

//TrdSubType is a non-required field for AllocationReport.
func (m *AllocationReport) TrdSubType() (*field.TrdSubType, error) {
	f := new(field.TrdSubType)
	err := m.Body.Get(f)
	return f, err
}

//MultiLegReportingType is a non-required field for AllocationReport.
func (m *AllocationReport) MultiLegReportingType() (*field.MultiLegReportingType, error) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}

//CustOrderCapacity is a non-required field for AllocationReport.
func (m *AllocationReport) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//TradeInputSource is a non-required field for AllocationReport.
func (m *AllocationReport) TradeInputSource() (*field.TradeInputSource, error) {
	f := new(field.TradeInputSource)
	err := m.Body.Get(f)
	return f, err
}

//RndPx is a non-required field for AllocationReport.
func (m *AllocationReport) RndPx() (*field.RndPx, error) {
	f := new(field.RndPx)
	err := m.Body.Get(f)
	return f, err
}

//MessageEventSource is a non-required field for AllocationReport.
func (m *AllocationReport) MessageEventSource() (*field.MessageEventSource, error) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}

//TradeInputDevice is a non-required field for AllocationReport.
func (m *AllocationReport) TradeInputDevice() (*field.TradeInputDevice, error) {
	f := new(field.TradeInputDevice)
	err := m.Body.Get(f)
	return f, err
}

//AvgPxIndicator is a non-required field for AllocationReport.
func (m *AllocationReport) AvgPxIndicator() (*field.AvgPxIndicator, error) {
	f := new(field.AvgPxIndicator)
	err := m.Body.Get(f)
	return f, err
}

//NoPosAmt is a non-required field for AllocationReport.
func (m *AllocationReport) NoPosAmt() (*field.NoPosAmt, error) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}
