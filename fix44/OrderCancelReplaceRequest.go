package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderCancelReplaceRequest msg type = G.
type OrderCancelReplaceRequest struct {
	message.Message
}

//OrderCancelReplaceRequestBuilder builds OrderCancelReplaceRequest messages.
type OrderCancelReplaceRequestBuilder struct {
	message.MessageBuilder
}

//NewOrderCancelReplaceRequestBuilder returns an initialized OrderCancelReplaceRequestBuilder with specified required fields.
func NewOrderCancelReplaceRequestBuilder(
	origclordid field.OrigClOrdID,
	clordid field.ClOrdID,
	side field.Side,
	transacttime field.TransactTime,
	ordtype field.OrdType) *OrderCancelReplaceRequestBuilder {
	builder := new(OrderCancelReplaceRequestBuilder)
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(side)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//TradeOriginationDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) TradeOriginationDate() (*field.TradeOriginationDate, error) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//OrigClOrdID is a required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) OrigClOrdID() (*field.OrigClOrdID, error) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryClOrdID is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdLinkID is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ClOrdLinkID() (*field.ClOrdLinkID, error) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}

//ListID is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//OrigOrdModTime is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) OrigOrdModTime() (*field.OrigOrdModTime, error) {
	f := new(field.OrigOrdModTime)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AcctIDSource is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//DayBookingInst is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) DayBookingInst() (*field.DayBookingInst, error) {
	f := new(field.DayBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//BookingUnit is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) BookingUnit() (*field.BookingUnit, error) {
	f := new(field.BookingUnit)
	err := m.Body.Get(f)
	return f, err
}

//PreallocMethod is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) PreallocMethod() (*field.PreallocMethod, error) {
	f := new(field.PreallocMethod)
	err := m.Body.Get(f)
	return f, err
}

//AllocID is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//NoAllocs is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) NoAllocs() (*field.NoAllocs, error) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//SettlType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//CashMargin is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CashMargin() (*field.CashMargin, error) {
	f := new(field.CashMargin)
	err := m.Body.Get(f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ClearingFeeIndicator() (*field.ClearingFeeIndicator, error) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}

//HandlInst is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//ExecInst is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//MinQty is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//MaxFloor is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//ExDestination is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//NoTradingSessions is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDesc is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//AgreementID is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementCurrency is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//TerminationType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//StartDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//EndDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//MarginRatio is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//QtyType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//CashOrderQty is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//OrderPercent is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) OrderPercent() (*field.OrderPercent, error) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//RoundingDirection is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) RoundingDirection() (*field.RoundingDirection, error) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//RoundingModulus is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) RoundingModulus() (*field.RoundingModulus, error) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//OrdType is a required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//Price is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//StopPx is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//Spread is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveName is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPrice is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPriceType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//YieldType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//Yield is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//YieldCalcDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) YieldCalcDate() (*field.YieldCalcDate, error) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) YieldRedemptionDate() (*field.YieldRedemptionDate, error) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) YieldRedemptionPrice() (*field.YieldRedemptionPrice, error) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, error) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//PegOffsetValue is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) PegOffsetValue() (*field.PegOffsetValue, error) {
	f := new(field.PegOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//PegMoveType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) PegMoveType() (*field.PegMoveType, error) {
	f := new(field.PegMoveType)
	err := m.Body.Get(f)
	return f, err
}

//PegOffsetType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) PegOffsetType() (*field.PegOffsetType, error) {
	f := new(field.PegOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//PegLimitType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) PegLimitType() (*field.PegLimitType, error) {
	f := new(field.PegLimitType)
	err := m.Body.Get(f)
	return f, err
}

//PegRoundDirection is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) PegRoundDirection() (*field.PegRoundDirection, error) {
	f := new(field.PegRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//PegScope is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) PegScope() (*field.PegScope, error) {
	f := new(field.PegScope)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionInst is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) DiscretionInst() (*field.DiscretionInst, error) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionOffsetValue is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) DiscretionOffsetValue() (*field.DiscretionOffsetValue, error) {
	f := new(field.DiscretionOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionMoveType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) DiscretionMoveType() (*field.DiscretionMoveType, error) {
	f := new(field.DiscretionMoveType)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionOffsetType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) DiscretionOffsetType() (*field.DiscretionOffsetType, error) {
	f := new(field.DiscretionOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionLimitType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) DiscretionLimitType() (*field.DiscretionLimitType, error) {
	f := new(field.DiscretionLimitType)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionRoundDirection is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) DiscretionRoundDirection() (*field.DiscretionRoundDirection, error) {
	f := new(field.DiscretionRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionScope is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) DiscretionScope() (*field.DiscretionScope, error) {
	f := new(field.DiscretionScope)
	err := m.Body.Get(f)
	return f, err
}

//TargetStrategy is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) TargetStrategy() (*field.TargetStrategy, error) {
	f := new(field.TargetStrategy)
	err := m.Body.Get(f)
	return f, err
}

//TargetStrategyParameters is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) TargetStrategyParameters() (*field.TargetStrategyParameters, error) {
	f := new(field.TargetStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//ParticipationRate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ParticipationRate() (*field.ParticipationRate, error) {
	f := new(field.ParticipationRate)
	err := m.Body.Get(f)
	return f, err
}

//ComplianceID is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ComplianceID() (*field.ComplianceID, error) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//SolicitedFlag is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SolicitedFlag() (*field.SolicitedFlag, error) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//TimeInForce is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//EffectiveTime is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//ExpireDate is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//ExpireTime is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GTBookingInst is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) GTBookingInst() (*field.GTBookingInst, error) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//Commission is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//CommType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//CommCurrency is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CommCurrency() (*field.CommCurrency, error) {
	f := new(field.CommCurrency)
	err := m.Body.Get(f)
	return f, err
}

//FundRenewWaiv is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) FundRenewWaiv() (*field.FundRenewWaiv, error) {
	f := new(field.FundRenewWaiv)
	err := m.Body.Get(f)
	return f, err
}

//OrderCapacity is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) OrderCapacity() (*field.OrderCapacity, error) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//OrderRestrictions is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) OrderRestrictions() (*field.OrderRestrictions, error) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}

//CustOrderCapacity is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//ForexReq is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) ForexReq() (*field.ForexReq, error) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrency is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BookingType is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) BookingType() (*field.BookingType, error) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate2 is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) SettlDate2() (*field.SettlDate2, error) {
	f := new(field.SettlDate2)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty2 is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) OrderQty2() (*field.OrderQty2, error) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//Price2 is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Price2() (*field.Price2, error) {
	f := new(field.Price2)
	err := m.Body.Get(f)
	return f, err
}

//PositionEffect is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) PositionEffect() (*field.PositionEffect, error) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//CoveredOrUncovered is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CoveredOrUncovered() (*field.CoveredOrUncovered, error) {
	f := new(field.CoveredOrUncovered)
	err := m.Body.Get(f)
	return f, err
}

//MaxShow is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//LocateReqd is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) LocateReqd() (*field.LocateReqd, error) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//CancellationRights is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//RegistID is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//Designation is a non-required field for OrderCancelReplaceRequest.
func (m *OrderCancelReplaceRequest) Designation() (*field.Designation, error) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}
