package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CrossOrderCancelReplaceRequest msg type = t.
type CrossOrderCancelReplaceRequest struct {
	message.Message
}

//CrossOrderCancelReplaceRequestBuilder builds CrossOrderCancelReplaceRequest messages.
type CrossOrderCancelReplaceRequestBuilder struct {
	message.MessageBuilder
}

//NewCrossOrderCancelReplaceRequestBuilder returns an initialized CrossOrderCancelReplaceRequestBuilder with specified required fields.
func NewCrossOrderCancelReplaceRequestBuilder(
	crossid field.CrossID,
	origcrossid field.OrigCrossID,
	crosstype field.CrossType,
	crossprioritization field.CrossPrioritization,
	nosides field.NoSides,
	transacttime field.TransactTime,
	ordtype field.OrdType) *CrossOrderCancelReplaceRequestBuilder {
	builder := new(CrossOrderCancelReplaceRequestBuilder)
	builder.Body.Set(crossid)
	builder.Body.Set(origcrossid)
	builder.Body.Set(crosstype)
	builder.Body.Set(crossprioritization)
	builder.Body.Set(nosides)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//CrossID is a required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) CrossID() (*field.CrossID, error) {
	f := new(field.CrossID)
	err := m.Body.Get(f)
	return f, err
}

//OrigCrossID is a required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) OrigCrossID() (*field.OrigCrossID, error) {
	f := new(field.OrigCrossID)
	err := m.Body.Get(f)
	return f, err
}

//CrossType is a required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) CrossType() (*field.CrossType, error) {
	f := new(field.CrossType)
	err := m.Body.Get(f)
	return f, err
}

//CrossPrioritization is a required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) CrossPrioritization() (*field.CrossPrioritization, error) {
	f := new(field.CrossPrioritization)
	err := m.Body.Get(f)
	return f, err
}

//NoSides is a required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) NoSides() (*field.NoSides, error) {
	f := new(field.NoSides)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//SettlType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//HandlInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//ExecInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//MinQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//MaxFloor is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//ExDestination is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//NoTradingSessions is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//ProcessCode is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) ProcessCode() (*field.ProcessCode, error) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}

//PrevClosePx is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PrevClosePx() (*field.PrevClosePx, error) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}

//LocateReqd is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) LocateReqd() (*field.LocateReqd, error) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//NoStipulations is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//OrdType is a required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//Price is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//StopPx is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//Spread is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveName is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//YieldType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//Yield is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//YieldCalcDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) YieldCalcDate() (*field.YieldCalcDate, error) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) YieldRedemptionDate() (*field.YieldRedemptionDate, error) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) YieldRedemptionPrice() (*field.YieldRedemptionPrice, error) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, error) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//ComplianceID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) ComplianceID() (*field.ComplianceID, error) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//IOIID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) IOIID() (*field.IOIID, error) {
	f := new(field.IOIID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//TimeInForce is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//EffectiveTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//ExpireDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//ExpireTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GTBookingInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) GTBookingInst() (*field.GTBookingInst, error) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//MaxShow is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//PegOffsetValue is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PegOffsetValue() (*field.PegOffsetValue, error) {
	f := new(field.PegOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//PegMoveType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PegMoveType() (*field.PegMoveType, error) {
	f := new(field.PegMoveType)
	err := m.Body.Get(f)
	return f, err
}

//PegOffsetType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PegOffsetType() (*field.PegOffsetType, error) {
	f := new(field.PegOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//PegLimitType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PegLimitType() (*field.PegLimitType, error) {
	f := new(field.PegLimitType)
	err := m.Body.Get(f)
	return f, err
}

//PegRoundDirection is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PegRoundDirection() (*field.PegRoundDirection, error) {
	f := new(field.PegRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//PegScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PegScope() (*field.PegScope, error) {
	f := new(field.PegScope)
	err := m.Body.Get(f)
	return f, err
}

//PegPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PegPriceType() (*field.PegPriceType, error) {
	f := new(field.PegPriceType)
	err := m.Body.Get(f)
	return f, err
}

//PegSecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PegSecurityIDSource() (*field.PegSecurityIDSource, error) {
	f := new(field.PegSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//PegSecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PegSecurityID() (*field.PegSecurityID, error) {
	f := new(field.PegSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//PegSymbol is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PegSymbol() (*field.PegSymbol, error) {
	f := new(field.PegSymbol)
	err := m.Body.Get(f)
	return f, err
}

//PegSecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PegSecurityDesc() (*field.PegSecurityDesc, error) {
	f := new(field.PegSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DiscretionInst() (*field.DiscretionInst, error) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionOffsetValue is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DiscretionOffsetValue() (*field.DiscretionOffsetValue, error) {
	f := new(field.DiscretionOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionMoveType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DiscretionMoveType() (*field.DiscretionMoveType, error) {
	f := new(field.DiscretionMoveType)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionOffsetType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DiscretionOffsetType() (*field.DiscretionOffsetType, error) {
	f := new(field.DiscretionOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionLimitType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DiscretionLimitType() (*field.DiscretionLimitType, error) {
	f := new(field.DiscretionLimitType)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionRoundDirection is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DiscretionRoundDirection() (*field.DiscretionRoundDirection, error) {
	f := new(field.DiscretionRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DiscretionScope() (*field.DiscretionScope, error) {
	f := new(field.DiscretionScope)
	err := m.Body.Get(f)
	return f, err
}

//TargetStrategy is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TargetStrategy() (*field.TargetStrategy, error) {
	f := new(field.TargetStrategy)
	err := m.Body.Get(f)
	return f, err
}

//TargetStrategyParameters is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TargetStrategyParameters() (*field.TargetStrategyParameters, error) {
	f := new(field.TargetStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//ParticipationRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) ParticipationRate() (*field.ParticipationRate, error) {
	f := new(field.ParticipationRate)
	err := m.Body.Get(f)
	return f, err
}

//CancellationRights is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//RegistID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//Designation is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) Designation() (*field.Designation, error) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}

//NoStrategyParameters is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) NoStrategyParameters() (*field.NoStrategyParameters, error) {
	f := new(field.NoStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//HostCrossID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) HostCrossID() (*field.HostCrossID, error) {
	f := new(field.HostCrossID)
	err := m.Body.Get(f)
	return f, err
}

//TransBkdTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TransBkdTime() (*field.TransBkdTime, error) {
	f := new(field.TransBkdTime)
	err := m.Body.Get(f)
	return f, err
}

//NoRootPartyIDs is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) NoRootPartyIDs() (*field.NoRootPartyIDs, error) {
	f := new(field.NoRootPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//MatchIncrement is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) MatchIncrement() (*field.MatchIncrement, error) {
	f := new(field.MatchIncrement)
	err := m.Body.Get(f)
	return f, err
}

//MaxPriceLevels is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) MaxPriceLevels() (*field.MaxPriceLevels, error) {
	f := new(field.MaxPriceLevels)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryDisplayQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) SecondaryDisplayQty() (*field.SecondaryDisplayQty, error) {
	f := new(field.SecondaryDisplayQty)
	err := m.Body.Get(f)
	return f, err
}

//DisplayWhen is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DisplayWhen() (*field.DisplayWhen, error) {
	f := new(field.DisplayWhen)
	err := m.Body.Get(f)
	return f, err
}

//DisplayMethod is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DisplayMethod() (*field.DisplayMethod, error) {
	f := new(field.DisplayMethod)
	err := m.Body.Get(f)
	return f, err
}

//DisplayLowQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DisplayLowQty() (*field.DisplayLowQty, error) {
	f := new(field.DisplayLowQty)
	err := m.Body.Get(f)
	return f, err
}

//DisplayHighQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DisplayHighQty() (*field.DisplayHighQty, error) {
	f := new(field.DisplayHighQty)
	err := m.Body.Get(f)
	return f, err
}

//DisplayMinIncr is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DisplayMinIncr() (*field.DisplayMinIncr, error) {
	f := new(field.DisplayMinIncr)
	err := m.Body.Get(f)
	return f, err
}

//RefreshQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) RefreshQty() (*field.RefreshQty, error) {
	f := new(field.RefreshQty)
	err := m.Body.Get(f)
	return f, err
}

//DisplayQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) DisplayQty() (*field.DisplayQty, error) {
	f := new(field.DisplayQty)
	err := m.Body.Get(f)
	return f, err
}

//PriceProtectionScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) PriceProtectionScope() (*field.PriceProtectionScope, error) {
	f := new(field.PriceProtectionScope)
	err := m.Body.Get(f)
	return f, err
}

//TriggerType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerType() (*field.TriggerType, error) {
	f := new(field.TriggerType)
	err := m.Body.Get(f)
	return f, err
}

//TriggerAction is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerAction() (*field.TriggerAction, error) {
	f := new(field.TriggerAction)
	err := m.Body.Get(f)
	return f, err
}

//TriggerPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerPrice() (*field.TriggerPrice, error) {
	f := new(field.TriggerPrice)
	err := m.Body.Get(f)
	return f, err
}

//TriggerSymbol is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerSymbol() (*field.TriggerSymbol, error) {
	f := new(field.TriggerSymbol)
	err := m.Body.Get(f)
	return f, err
}

//TriggerSecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerSecurityID() (*field.TriggerSecurityID, error) {
	f := new(field.TriggerSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//TriggerSecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerSecurityIDSource() (*field.TriggerSecurityIDSource, error) {
	f := new(field.TriggerSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//TriggerSecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerSecurityDesc() (*field.TriggerSecurityDesc, error) {
	f := new(field.TriggerSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//TriggerPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerPriceType() (*field.TriggerPriceType, error) {
	f := new(field.TriggerPriceType)
	err := m.Body.Get(f)
	return f, err
}

//TriggerPriceTypeScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerPriceTypeScope() (*field.TriggerPriceTypeScope, error) {
	f := new(field.TriggerPriceTypeScope)
	err := m.Body.Get(f)
	return f, err
}

//TriggerPriceDirection is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerPriceDirection() (*field.TriggerPriceDirection, error) {
	f := new(field.TriggerPriceDirection)
	err := m.Body.Get(f)
	return f, err
}

//TriggerNewPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerNewPrice() (*field.TriggerNewPrice, error) {
	f := new(field.TriggerNewPrice)
	err := m.Body.Get(f)
	return f, err
}

//TriggerOrderType is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerOrderType() (*field.TriggerOrderType, error) {
	f := new(field.TriggerOrderType)
	err := m.Body.Get(f)
	return f, err
}

//TriggerNewQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerNewQty() (*field.TriggerNewQty, error) {
	f := new(field.TriggerNewQty)
	err := m.Body.Get(f)
	return f, err
}

//TriggerTradingSessionID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerTradingSessionID() (*field.TriggerTradingSessionID, error) {
	f := new(field.TriggerTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TriggerTradingSessionSubID is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubID, error) {
	f := new(field.TriggerTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//ExDestinationIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m *CrossOrderCancelReplaceRequest) ExDestinationIDSource() (*field.ExDestinationIDSource, error) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}
