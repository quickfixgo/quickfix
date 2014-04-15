package fix44

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

//CreateCrossOrderCancelReplaceRequestBuilder returns an initialized CrossOrderCancelReplaceRequestBuilder with specified required fields.
func CreateCrossOrderCancelReplaceRequestBuilder(
	crossid field.CrossID,
	origcrossid field.OrigCrossID,
	crosstype field.CrossType,
	crossprioritization field.CrossPrioritization,
	nosides field.NoSides,
	transacttime field.TransactTime,
	ordtype field.OrdType) CrossOrderCancelReplaceRequestBuilder {
	var builder CrossOrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
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
func (m CrossOrderCancelReplaceRequest) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//CrossID is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CrossID() (field.CrossID, error) {
	var f field.CrossID
	err := m.Body.Get(&f)
	return f, err
}

//OrigCrossID is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OrigCrossID() (field.OrigCrossID, error) {
	var f field.OrigCrossID
	err := m.Body.Get(&f)
	return f, err
}

//CrossType is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CrossType() (field.CrossType, error) {
	var f field.CrossType
	err := m.Body.Get(&f)
	return f, err
}

//CrossPrioritization is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CrossPrioritization() (field.CrossPrioritization, error) {
	var f field.CrossPrioritization
	err := m.Body.Get(&f)
	return f, err
}

//NoSides is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoSides() (field.NoSides, error) {
	var f field.NoSides
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//HandlInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) HandlInst() (field.HandlInst, error) {
	var f field.HandlInst
	err := m.Body.Get(&f)
	return f, err
}

//ExecInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExecInst() (field.ExecInst, error) {
	var f field.ExecInst
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MinQty() (field.MinQty, error) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//MaxFloor is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaxFloor() (field.MaxFloor, error) {
	var f field.MaxFloor
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//NoTradingSessions is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoTradingSessions() (field.NoTradingSessions, error) {
	var f field.NoTradingSessions
	err := m.Body.Get(&f)
	return f, err
}

//ProcessCode is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ProcessCode() (field.ProcessCode, error) {
	var f field.ProcessCode
	err := m.Body.Get(&f)
	return f, err
}

//PrevClosePx is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PrevClosePx() (field.PrevClosePx, error) {
	var f field.PrevClosePx
	err := m.Body.Get(&f)
	return f, err
}

//LocateReqd is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) LocateReqd() (field.LocateReqd, error) {
	var f field.LocateReqd
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//StopPx is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StopPx() (field.StopPx, error) {
	var f field.StopPx
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldType() (field.YieldType, error) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Yield() (field.Yield, error) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldCalcDate() (field.YieldCalcDate, error) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldRedemptionDate() (field.YieldRedemptionDate, error) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldRedemptionPrice() (field.YieldRedemptionPrice, error) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, error) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ComplianceID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ComplianceID() (field.ComplianceID, error) {
	var f field.ComplianceID
	err := m.Body.Get(&f)
	return f, err
}

//IOIID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) IOIID() (field.IOIID, error) {
	var f field.IOIID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//TimeInForce is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TimeInForce() (field.TimeInForce, error) {
	var f field.TimeInForce
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EffectiveTime() (field.EffectiveTime, error) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExpireDate() (field.ExpireDate, error) {
	var f field.ExpireDate
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//GTBookingInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GTBookingInst() (field.GTBookingInst, error) {
	var f field.GTBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//MaxShow is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaxShow() (field.MaxShow, error) {
	var f field.MaxShow
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetValue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegOffsetValue() (field.PegOffsetValue, error) {
	var f field.PegOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//PegMoveType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegMoveType() (field.PegMoveType, error) {
	var f field.PegMoveType
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegOffsetType() (field.PegOffsetType, error) {
	var f field.PegOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//PegLimitType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegLimitType() (field.PegLimitType, error) {
	var f field.PegLimitType
	err := m.Body.Get(&f)
	return f, err
}

//PegRoundDirection is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegRoundDirection() (field.PegRoundDirection, error) {
	var f field.PegRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//PegScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegScope() (field.PegScope, error) {
	var f field.PegScope
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionInst() (field.DiscretionInst, error) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetValue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionOffsetValue() (field.DiscretionOffsetValue, error) {
	var f field.DiscretionOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionMoveType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionMoveType() (field.DiscretionMoveType, error) {
	var f field.DiscretionMoveType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionOffsetType() (field.DiscretionOffsetType, error) {
	var f field.DiscretionOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionLimitType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionLimitType() (field.DiscretionLimitType, error) {
	var f field.DiscretionLimitType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionRoundDirection is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionRoundDirection() (field.DiscretionRoundDirection, error) {
	var f field.DiscretionRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionScope is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionScope() (field.DiscretionScope, error) {
	var f field.DiscretionScope
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategy is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TargetStrategy() (field.TargetStrategy, error) {
	var f field.TargetStrategy
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategyParameters is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TargetStrategyParameters() (field.TargetStrategyParameters, error) {
	var f field.TargetStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//ParticipationRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ParticipationRate() (field.ParticipationRate, error) {
	var f field.ParticipationRate
	err := m.Body.Get(&f)
	return f, err
}

//CancellationRights is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CancellationRights() (field.CancellationRights, error) {
	var f field.CancellationRights
	err := m.Body.Get(&f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MoneyLaunderingStatus() (field.MoneyLaunderingStatus, error) {
	var f field.MoneyLaunderingStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RegistID() (field.RegistID, error) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//Designation is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Designation() (field.Designation, error) {
	var f field.Designation
	err := m.Body.Get(&f)
	return f, err
}
