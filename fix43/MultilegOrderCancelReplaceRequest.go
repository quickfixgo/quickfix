package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MultilegOrderCancelReplaceRequest msg type = AC.
type MultilegOrderCancelReplaceRequest struct {
	message.Message
}

//MultilegOrderCancelReplaceRequestBuilder builds MultilegOrderCancelReplaceRequest messages.
type MultilegOrderCancelReplaceRequestBuilder struct {
	message.MessageBuilder
}

//CreateMultilegOrderCancelReplaceRequestBuilder returns an initialized MultilegOrderCancelReplaceRequestBuilder with specified required fields.
func CreateMultilegOrderCancelReplaceRequestBuilder(
	origclordid field.OrigClOrdID,
	clordid field.ClOrdID,
	handlinst field.HandlInst,
	side field.Side,
	nolegs field.NoLegs,
	transacttime field.TransactTime,
	ordtype field.OrdType) MultilegOrderCancelReplaceRequestBuilder {
	var builder MultilegOrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(handlinst)
	builder.Body.Set(side)
	builder.Body.Set(nolegs)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//OrigClOrdID is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrigClOrdID() (field.OrigClOrdID, error) {
	var f field.OrigClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecondaryClOrdID() (field.SecondaryClOrdID, error) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdLinkID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ClOrdLinkID() (field.ClOrdLinkID, error) {
	var f field.ClOrdLinkID
	err := m.Body.Get(&f)
	return f, err
}

//OrigOrdModTime is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrigOrdModTime() (field.OrigOrdModTime, error) {
	var f field.OrigOrdModTime
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//DayBookingInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) DayBookingInst() (field.DayBookingInst, error) {
	var f field.DayBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//BookingUnit is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) BookingUnit() (field.BookingUnit, error) {
	var f field.BookingUnit
	err := m.Body.Get(&f)
	return f, err
}

//PreallocMethod is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PreallocMethod() (field.PreallocMethod, error) {
	var f field.PreallocMethod
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoAllocs() (field.NoAllocs, error) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}

//SettlmntTyp is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SettlmntTyp() (field.SettlmntTyp, error) {
	var f field.SettlmntTyp
	err := m.Body.Get(&f)
	return f, err
}

//FutSettDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) FutSettDate() (field.FutSettDate, error) {
	var f field.FutSettDate
	err := m.Body.Get(&f)
	return f, err
}

//CashMargin is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CashMargin() (field.CashMargin, error) {
	var f field.CashMargin
	err := m.Body.Get(&f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ClearingFeeIndicator() (field.ClearingFeeIndicator, error) {
	var f field.ClearingFeeIndicator
	err := m.Body.Get(&f)
	return f, err
}

//HandlInst is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) HandlInst() (field.HandlInst, error) {
	var f field.HandlInst
	err := m.Body.Get(&f)
	return f, err
}

//ExecInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ExecInst() (field.ExecInst, error) {
	var f field.ExecInst
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MinQty() (field.MinQty, error) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//MaxFloor is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MaxFloor() (field.MaxFloor, error) {
	var f field.MaxFloor
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//NoTradingSessions is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoTradingSessions() (field.NoTradingSessions, error) {
	var f field.NoTradingSessions
	err := m.Body.Get(&f)
	return f, err
}

//ProcessCode is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ProcessCode() (field.ProcessCode, error) {
	var f field.ProcessCode
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//PrevClosePx is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PrevClosePx() (field.PrevClosePx, error) {
	var f field.PrevClosePx
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//LocateReqd is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) LocateReqd() (field.LocateReqd, error) {
	var f field.LocateReqd
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//QuantityType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) QuantityType() (field.QuantityType, error) {
	var f field.QuantityType
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderPercent() (field.OrderPercent, error) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RoundingDirection() (field.RoundingDirection, error) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RoundingModulus() (field.RoundingModulus, error) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//StopPx is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) StopPx() (field.StopPx, error) {
	var f field.StopPx
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ComplianceID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ComplianceID() (field.ComplianceID, error) {
	var f field.ComplianceID
	err := m.Body.Get(&f)
	return f, err
}

//SolicitedFlag is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SolicitedFlag() (field.SolicitedFlag, error) {
	var f field.SolicitedFlag
	err := m.Body.Get(&f)
	return f, err
}

//IOIid is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) IOIid() (field.IOIid, error) {
	var f field.IOIid
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//TimeInForce is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) TimeInForce() (field.TimeInForce, error) {
	var f field.TimeInForce
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EffectiveTime() (field.EffectiveTime, error) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ExpireDate() (field.ExpireDate, error) {
	var f field.ExpireDate
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//GTBookingInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GTBookingInst() (field.GTBookingInst, error) {
	var f field.GTBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//CommCurrency is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CommCurrency() (field.CommCurrency, error) {
	var f field.CommCurrency
	err := m.Body.Get(&f)
	return f, err
}

//FundRenewWaiv is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) FundRenewWaiv() (field.FundRenewWaiv, error) {
	var f field.FundRenewWaiv
	err := m.Body.Get(&f)
	return f, err
}

//OrderCapacity is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderCapacity() (field.OrderCapacity, error) {
	var f field.OrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//OrderRestrictions is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderRestrictions() (field.OrderRestrictions, error) {
	var f field.OrderRestrictions
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CustOrderCapacity() (field.CustOrderCapacity, error) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//ForexReq is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ForexReq() (field.ForexReq, error) {
	var f field.ForexReq
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SettlCurrency() (field.SettlCurrency, error) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//PositionEffect is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PositionEffect() (field.PositionEffect, error) {
	var f field.PositionEffect
	err := m.Body.Get(&f)
	return f, err
}

//CoveredOrUncovered is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CoveredOrUncovered() (field.CoveredOrUncovered, error) {
	var f field.CoveredOrUncovered
	err := m.Body.Get(&f)
	return f, err
}

//MaxShow is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MaxShow() (field.MaxShow, error) {
	var f field.MaxShow
	err := m.Body.Get(&f)
	return f, err
}

//PegDifference is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PegDifference() (field.PegDifference, error) {
	var f field.PegDifference
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) DiscretionInst() (field.DiscretionInst, error) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffset is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) DiscretionOffset() (field.DiscretionOffset, error) {
	var f field.DiscretionOffset
	err := m.Body.Get(&f)
	return f, err
}

//CancellationRights is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CancellationRights() (field.CancellationRights, error) {
	var f field.CancellationRights
	err := m.Body.Get(&f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MoneyLaunderingStatus() (field.MoneyLaunderingStatus, error) {
	var f field.MoneyLaunderingStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RegistID() (field.RegistID, error) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//Designation is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Designation() (field.Designation, error) {
	var f field.Designation
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegRptTypeReq is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MultiLegRptTypeReq() (field.MultiLegRptTypeReq, error) {
	var f field.MultiLegRptTypeReq
	err := m.Body.Get(&f)
	return f, err
}

//NetMoney is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NetMoney() (field.NetMoney, error) {
	var f field.NetMoney
	err := m.Body.Get(&f)
	return f, err
}
