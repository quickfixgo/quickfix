package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MultilegOrderCancelReplace msg type = AC.
type MultilegOrderCancelReplace struct {
	message.Message
}

//MultilegOrderCancelReplaceBuilder builds MultilegOrderCancelReplace messages.
type MultilegOrderCancelReplaceBuilder struct {
	message.MessageBuilder
}

//CreateMultilegOrderCancelReplaceBuilder returns an initialized MultilegOrderCancelReplaceBuilder with specified required fields.
func CreateMultilegOrderCancelReplaceBuilder(
	origclordid field.OrigClOrdID,
	clordid field.ClOrdID,
	side field.Side,
	nolegs field.NoLegs,
	transacttime field.TransactTime,
	ordtype field.OrdType) MultilegOrderCancelReplaceBuilder {
	var builder MultilegOrderCancelReplaceBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(side)
	builder.Body.Set(nolegs)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//OrigClOrdID is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrigClOrdID() (field.OrigClOrdID, error) {
	var f field.OrigClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecondaryClOrdID() (field.SecondaryClOrdID, error) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdLinkID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ClOrdLinkID() (field.ClOrdLinkID, error) {
	var f field.ClOrdLinkID
	err := m.Body.Get(&f)
	return f, err
}

//OrigOrdModTime is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrigOrdModTime() (field.OrigOrdModTime, error) {
	var f field.OrigOrdModTime
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradeOriginationDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TradeOriginationDate() (field.TradeOriginationDate, error) {
	var f field.TradeOriginationDate
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//DayBookingInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DayBookingInst() (field.DayBookingInst, error) {
	var f field.DayBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//BookingUnit is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) BookingUnit() (field.BookingUnit, error) {
	var f field.BookingUnit
	err := m.Body.Get(&f)
	return f, err
}

//PreallocMethod is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PreallocMethod() (field.PreallocMethod, error) {
	var f field.PreallocMethod
	err := m.Body.Get(&f)
	return f, err
}

//AllocID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) AllocID() (field.AllocID, error) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoAllocs() (field.NoAllocs, error) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//CashMargin is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CashMargin() (field.CashMargin, error) {
	var f field.CashMargin
	err := m.Body.Get(&f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ClearingFeeIndicator() (field.ClearingFeeIndicator, error) {
	var f field.ClearingFeeIndicator
	err := m.Body.Get(&f)
	return f, err
}

//HandlInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) HandlInst() (field.HandlInst, error) {
	var f field.HandlInst
	err := m.Body.Get(&f)
	return f, err
}

//ExecInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExecInst() (field.ExecInst, error) {
	var f field.ExecInst
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MinQty() (field.MinQty, error) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//MaxFloor is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaxFloor() (field.MaxFloor, error) {
	var f field.MaxFloor
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//NoTradingSessions is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoTradingSessions() (field.NoTradingSessions, error) {
	var f field.NoTradingSessions
	err := m.Body.Get(&f)
	return f, err
}

//ProcessCode is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ProcessCode() (field.ProcessCode, error) {
	var f field.ProcessCode
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//PrevClosePx is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PrevClosePx() (field.PrevClosePx, error) {
	var f field.PrevClosePx
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//LocateReqd is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) LocateReqd() (field.LocateReqd, error) {
	var f field.LocateReqd
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) QtyType() (field.QtyType, error) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderPercent() (field.OrderPercent, error) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RoundingDirection() (field.RoundingDirection, error) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RoundingModulus() (field.RoundingModulus, error) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//StopPx is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StopPx() (field.StopPx, error) {
	var f field.StopPx
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ComplianceID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ComplianceID() (field.ComplianceID, error) {
	var f field.ComplianceID
	err := m.Body.Get(&f)
	return f, err
}

//SolicitedFlag is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SolicitedFlag() (field.SolicitedFlag, error) {
	var f field.SolicitedFlag
	err := m.Body.Get(&f)
	return f, err
}

//IOIID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) IOIID() (field.IOIID, error) {
	var f field.IOIID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//TimeInForce is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TimeInForce() (field.TimeInForce, error) {
	var f field.TimeInForce
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EffectiveTime() (field.EffectiveTime, error) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExpireDate() (field.ExpireDate, error) {
	var f field.ExpireDate
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//GTBookingInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GTBookingInst() (field.GTBookingInst, error) {
	var f field.GTBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//CommCurrency is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CommCurrency() (field.CommCurrency, error) {
	var f field.CommCurrency
	err := m.Body.Get(&f)
	return f, err
}

//FundRenewWaiv is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) FundRenewWaiv() (field.FundRenewWaiv, error) {
	var f field.FundRenewWaiv
	err := m.Body.Get(&f)
	return f, err
}

//OrderCapacity is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderCapacity() (field.OrderCapacity, error) {
	var f field.OrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//OrderRestrictions is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderRestrictions() (field.OrderRestrictions, error) {
	var f field.OrderRestrictions
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CustOrderCapacity() (field.CustOrderCapacity, error) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//ForexReq is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ForexReq() (field.ForexReq, error) {
	var f field.ForexReq
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SettlCurrency() (field.SettlCurrency, error) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BookingType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) BookingType() (field.BookingType, error) {
	var f field.BookingType
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//PositionEffect is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PositionEffect() (field.PositionEffect, error) {
	var f field.PositionEffect
	err := m.Body.Get(&f)
	return f, err
}

//CoveredOrUncovered is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CoveredOrUncovered() (field.CoveredOrUncovered, error) {
	var f field.CoveredOrUncovered
	err := m.Body.Get(&f)
	return f, err
}

//MaxShow is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaxShow() (field.MaxShow, error) {
	var f field.MaxShow
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetValue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegOffsetValue() (field.PegOffsetValue, error) {
	var f field.PegOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//PegMoveType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegMoveType() (field.PegMoveType, error) {
	var f field.PegMoveType
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegOffsetType() (field.PegOffsetType, error) {
	var f field.PegOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//PegLimitType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegLimitType() (field.PegLimitType, error) {
	var f field.PegLimitType
	err := m.Body.Get(&f)
	return f, err
}

//PegRoundDirection is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegRoundDirection() (field.PegRoundDirection, error) {
	var f field.PegRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//PegScope is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegScope() (field.PegScope, error) {
	var f field.PegScope
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionInst() (field.DiscretionInst, error) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetValue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionOffsetValue() (field.DiscretionOffsetValue, error) {
	var f field.DiscretionOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionMoveType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionMoveType() (field.DiscretionMoveType, error) {
	var f field.DiscretionMoveType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionOffsetType() (field.DiscretionOffsetType, error) {
	var f field.DiscretionOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionLimitType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionLimitType() (field.DiscretionLimitType, error) {
	var f field.DiscretionLimitType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionRoundDirection is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionRoundDirection() (field.DiscretionRoundDirection, error) {
	var f field.DiscretionRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionScope is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionScope() (field.DiscretionScope, error) {
	var f field.DiscretionScope
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategy is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TargetStrategy() (field.TargetStrategy, error) {
	var f field.TargetStrategy
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategyParameters is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TargetStrategyParameters() (field.TargetStrategyParameters, error) {
	var f field.TargetStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//ParticipationRate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ParticipationRate() (field.ParticipationRate, error) {
	var f field.ParticipationRate
	err := m.Body.Get(&f)
	return f, err
}

//CancellationRights is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CancellationRights() (field.CancellationRights, error) {
	var f field.CancellationRights
	err := m.Body.Get(&f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MoneyLaunderingStatus() (field.MoneyLaunderingStatus, error) {
	var f field.MoneyLaunderingStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RegistID() (field.RegistID, error) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//Designation is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Designation() (field.Designation, error) {
	var f field.Designation
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegRptTypeReq is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MultiLegRptTypeReq() (field.MultiLegRptTypeReq, error) {
	var f field.MultiLegRptTypeReq
	err := m.Body.Get(&f)
	return f, err
}
