package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderSingle msg type = D.
type NewOrderSingle struct {
	message.Message
}

//NewOrderSingleBuilder builds NewOrderSingle messages.
type NewOrderSingleBuilder struct {
	message.MessageBuilder
}

//CreateNewOrderSingleBuilder returns an initialized NewOrderSingleBuilder with specified required fields.
func CreateNewOrderSingleBuilder(
	clordid field.ClOrdID,
	handlinst field.HandlInst,
	side field.Side,
	transacttime field.TransactTime,
	ordtype field.OrdType) NewOrderSingleBuilder {
	var builder NewOrderSingleBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(clordid)
	builder.Body.Set(handlinst)
	builder.Body.Set(side)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//ClOrdID is a required field for NewOrderSingle.
func (m NewOrderSingle) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecondaryClOrdID() (field.SecondaryClOrdID, error) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdLinkID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ClOrdLinkID() (field.ClOrdLinkID, error) {
	var f field.ClOrdLinkID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradeOriginationDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TradeOriginationDate() (field.TradeOriginationDate, error) {
	var f field.TradeOriginationDate
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//DayBookingInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DayBookingInst() (field.DayBookingInst, error) {
	var f field.DayBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//BookingUnit is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BookingUnit() (field.BookingUnit, error) {
	var f field.BookingUnit
	err := m.Body.Get(&f)
	return f, err
}

//PreallocMethod is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PreallocMethod() (field.PreallocMethod, error) {
	var f field.PreallocMethod
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoAllocs() (field.NoAllocs, error) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}

//SettlmntTyp is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlmntTyp() (field.SettlmntTyp, error) {
	var f field.SettlmntTyp
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) FutSettDate() (field.FutSettDate, error) {
	var f field.FutSettDate
	err := m.Body.Get(&f)
	return f, err
}

//CashMargin is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CashMargin() (field.CashMargin, error) {
	var f field.CashMargin
	err := m.Body.Get(&f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ClearingFeeIndicator() (field.ClearingFeeIndicator, error) {
	var f field.ClearingFeeIndicator
	err := m.Body.Get(&f)
	return f, err
}

//HandlInst is a required field for NewOrderSingle.
func (m NewOrderSingle) HandlInst() (field.HandlInst, error) {
	var f field.HandlInst
	err := m.Body.Get(&f)
	return f, err
}

//ExecInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExecInst() (field.ExecInst, error) {
	var f field.ExecInst
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MinQty() (field.MinQty, error) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//MaxFloor is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaxFloor() (field.MaxFloor, error) {
	var f field.MaxFloor
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//NoTradingSessions is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoTradingSessions() (field.NoTradingSessions, error) {
	var f field.NoTradingSessions
	err := m.Body.Get(&f)
	return f, err
}

//ProcessCode is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ProcessCode() (field.ProcessCode, error) {
	var f field.ProcessCode
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for NewOrderSingle.
func (m NewOrderSingle) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//PrevClosePx is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PrevClosePx() (field.PrevClosePx, error) {
	var f field.PrevClosePx
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for NewOrderSingle.
func (m NewOrderSingle) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//LocateReqd is a non-required field for NewOrderSingle.
func (m NewOrderSingle) LocateReqd() (field.LocateReqd, error) {
	var f field.LocateReqd
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for NewOrderSingle.
func (m NewOrderSingle) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//QuantityType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) QuantityType() (field.QuantityType, error) {
	var f field.QuantityType
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderPercent() (field.OrderPercent, error) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RoundingDirection() (field.RoundingDirection, error) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RoundingModulus() (field.RoundingModulus, error) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a required field for NewOrderSingle.
func (m NewOrderSingle) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//StopPx is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StopPx() (field.StopPx, error) {
	var f field.StopPx
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) YieldType() (field.YieldType, error) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Yield() (field.Yield, error) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ComplianceID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ComplianceID() (field.ComplianceID, error) {
	var f field.ComplianceID
	err := m.Body.Get(&f)
	return f, err
}

//SolicitedFlag is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SolicitedFlag() (field.SolicitedFlag, error) {
	var f field.SolicitedFlag
	err := m.Body.Get(&f)
	return f, err
}

//IOIid is a non-required field for NewOrderSingle.
func (m NewOrderSingle) IOIid() (field.IOIid, error) {
	var f field.IOIid
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//TimeInForce is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TimeInForce() (field.TimeInForce, error) {
	var f field.TimeInForce
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EffectiveTime() (field.EffectiveTime, error) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExpireDate() (field.ExpireDate, error) {
	var f field.ExpireDate
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//GTBookingInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) GTBookingInst() (field.GTBookingInst, error) {
	var f field.GTBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//CommCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CommCurrency() (field.CommCurrency, error) {
	var f field.CommCurrency
	err := m.Body.Get(&f)
	return f, err
}

//FundRenewWaiv is a non-required field for NewOrderSingle.
func (m NewOrderSingle) FundRenewWaiv() (field.FundRenewWaiv, error) {
	var f field.FundRenewWaiv
	err := m.Body.Get(&f)
	return f, err
}

//OrderCapacity is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderCapacity() (field.OrderCapacity, error) {
	var f field.OrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//OrderRestrictions is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderRestrictions() (field.OrderRestrictions, error) {
	var f field.OrderRestrictions
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CustOrderCapacity() (field.CustOrderCapacity, error) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//Rule80A is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Rule80A() (field.Rule80A, error) {
	var f field.Rule80A
	err := m.Body.Get(&f)
	return f, err
}

//ForexReq is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ForexReq() (field.ForexReq, error) {
	var f field.ForexReq
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlCurrency() (field.SettlCurrency, error) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate2 is a non-required field for NewOrderSingle.
func (m NewOrderSingle) FutSettDate2() (field.FutSettDate2, error) {
	var f field.FutSettDate2
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty2 is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderQty2() (field.OrderQty2, error) {
	var f field.OrderQty2
	err := m.Body.Get(&f)
	return f, err
}

//Price2 is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Price2() (field.Price2, error) {
	var f field.Price2
	err := m.Body.Get(&f)
	return f, err
}

//PositionEffect is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PositionEffect() (field.PositionEffect, error) {
	var f field.PositionEffect
	err := m.Body.Get(&f)
	return f, err
}

//CoveredOrUncovered is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CoveredOrUncovered() (field.CoveredOrUncovered, error) {
	var f field.CoveredOrUncovered
	err := m.Body.Get(&f)
	return f, err
}

//MaxShow is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaxShow() (field.MaxShow, error) {
	var f field.MaxShow
	err := m.Body.Get(&f)
	return f, err
}

//PegDifference is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegDifference() (field.PegDifference, error) {
	var f field.PegDifference
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionInst() (field.DiscretionInst, error) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffset is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionOffset() (field.DiscretionOffset, error) {
	var f field.DiscretionOffset
	err := m.Body.Get(&f)
	return f, err
}

//CancellationRights is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CancellationRights() (field.CancellationRights, error) {
	var f field.CancellationRights
	err := m.Body.Get(&f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MoneyLaunderingStatus() (field.MoneyLaunderingStatus, error) {
	var f field.MoneyLaunderingStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RegistID() (field.RegistID, error) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//Designation is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Designation() (field.Designation, error) {
	var f field.Designation
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestRate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AccruedInterestRate() (field.AccruedInterestRate, error) {
	var f field.AccruedInterestRate
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestAmt is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AccruedInterestAmt() (field.AccruedInterestAmt, error) {
	var f field.AccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//NetMoney is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NetMoney() (field.NetMoney, error) {
	var f field.NetMoney
	err := m.Body.Get(&f)
	return f, err
}
