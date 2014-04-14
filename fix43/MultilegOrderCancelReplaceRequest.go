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

//NewMultilegOrderCancelReplaceRequestBuilder returns an initialized MultilegOrderCancelReplaceRequestBuilder with specified required fields.
func NewMultilegOrderCancelReplaceRequestBuilder(
	origclordid field.OrigClOrdID,
	clordid field.ClOrdID,
	handlinst field.HandlInst,
	side field.Side,
	nolegs field.NoLegs,
	transacttime field.TransactTime,
	ordtype field.OrdType) *MultilegOrderCancelReplaceRequestBuilder {
	builder := new(MultilegOrderCancelReplaceRequestBuilder)
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
func (m *MultilegOrderCancelReplaceRequest) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//OrigClOrdID is a required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) OrigClOrdID() (*field.OrigClOrdID, error) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryClOrdID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdLinkID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) ClOrdLinkID() (*field.ClOrdLinkID, error) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}

//OrigOrdModTime is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) OrigOrdModTime() (*field.OrigOrdModTime, error) {
	f := new(field.OrigOrdModTime)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//DayBookingInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) DayBookingInst() (*field.DayBookingInst, error) {
	f := new(field.DayBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//BookingUnit is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) BookingUnit() (*field.BookingUnit, error) {
	f := new(field.BookingUnit)
	err := m.Body.Get(f)
	return f, err
}

//PreallocMethod is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) PreallocMethod() (*field.PreallocMethod, error) {
	f := new(field.PreallocMethod)
	err := m.Body.Get(f)
	return f, err
}

//NoAllocs is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) NoAllocs() (*field.NoAllocs, error) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//SettlmntTyp is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//FutSettDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//CashMargin is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) CashMargin() (*field.CashMargin, error) {
	f := new(field.CashMargin)
	err := m.Body.Get(f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) ClearingFeeIndicator() (*field.ClearingFeeIndicator, error) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}

//HandlInst is a required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//ExecInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//MinQty is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//MaxFloor is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//ExDestination is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//NoTradingSessions is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//ProcessCode is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) ProcessCode() (*field.ProcessCode, error) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//PrevClosePx is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) PrevClosePx() (*field.PrevClosePx, error) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//LocateReqd is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) LocateReqd() (*field.LocateReqd, error) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//QuantityType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) QuantityType() (*field.QuantityType, error) {
	f := new(field.QuantityType)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//CashOrderQty is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//OrderPercent is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) OrderPercent() (*field.OrderPercent, error) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//RoundingDirection is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) RoundingDirection() (*field.RoundingDirection, error) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//RoundingModulus is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) RoundingModulus() (*field.RoundingModulus, error) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//OrdType is a required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//Price is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//StopPx is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//ComplianceID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) ComplianceID() (*field.ComplianceID, error) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//SolicitedFlag is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) SolicitedFlag() (*field.SolicitedFlag, error) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}

//IOIid is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) IOIid() (*field.IOIid, error) {
	f := new(field.IOIid)
	err := m.Body.Get(f)
	return f, err
}

//QuoteID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//TimeInForce is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//EffectiveTime is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//ExpireDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//ExpireTime is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GTBookingInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) GTBookingInst() (*field.GTBookingInst, error) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//Commission is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//CommType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//CommCurrency is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) CommCurrency() (*field.CommCurrency, error) {
	f := new(field.CommCurrency)
	err := m.Body.Get(f)
	return f, err
}

//FundRenewWaiv is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) FundRenewWaiv() (*field.FundRenewWaiv, error) {
	f := new(field.FundRenewWaiv)
	err := m.Body.Get(f)
	return f, err
}

//OrderCapacity is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) OrderCapacity() (*field.OrderCapacity, error) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//OrderRestrictions is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) OrderRestrictions() (*field.OrderRestrictions, error) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}

//CustOrderCapacity is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//ForexReq is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) ForexReq() (*field.ForexReq, error) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrency is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//PositionEffect is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) PositionEffect() (*field.PositionEffect, error) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//CoveredOrUncovered is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) CoveredOrUncovered() (*field.CoveredOrUncovered, error) {
	f := new(field.CoveredOrUncovered)
	err := m.Body.Get(f)
	return f, err
}

//MaxShow is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//PegDifference is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) PegDifference() (*field.PegDifference, error) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) DiscretionInst() (*field.DiscretionInst, error) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionOffset is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) DiscretionOffset() (*field.DiscretionOffset, error) {
	f := new(field.DiscretionOffset)
	err := m.Body.Get(f)
	return f, err
}

//CancellationRights is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//RegistID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//Designation is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) Designation() (*field.Designation, error) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}

//MultiLegRptTypeReq is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) MultiLegRptTypeReq() (*field.MultiLegRptTypeReq, error) {
	f := new(field.MultiLegRptTypeReq)
	err := m.Body.Get(f)
	return f, err
}

//NetMoney is a non-required field for MultilegOrderCancelReplaceRequest.
func (m *MultilegOrderCancelReplaceRequest) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}
