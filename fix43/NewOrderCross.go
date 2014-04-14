package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderCross msg type = s.
type NewOrderCross struct {
	message.Message
}

//NewOrderCrossBuilder builds NewOrderCross messages.
type NewOrderCrossBuilder struct {
	message.MessageBuilder
}

//NewNewOrderCrossBuilder returns an initialized NewOrderCrossBuilder with specified required fields.
func NewNewOrderCrossBuilder(
	crossid field.CrossID,
	crosstype field.CrossType,
	crossprioritization field.CrossPrioritization,
	nosides field.NoSides,
	handlinst field.HandlInst,
	transacttime field.TransactTime,
	ordtype field.OrdType) *NewOrderCrossBuilder {
	builder := new(NewOrderCrossBuilder)
	builder.Body.Set(crossid)
	builder.Body.Set(crosstype)
	builder.Body.Set(crossprioritization)
	builder.Body.Set(nosides)
	builder.Body.Set(handlinst)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//CrossID is a required field for NewOrderCross.
func (m *NewOrderCross) CrossID() (*field.CrossID, error) {
	f := new(field.CrossID)
	err := m.Body.Get(f)
	return f, err
}

//CrossType is a required field for NewOrderCross.
func (m *NewOrderCross) CrossType() (*field.CrossType, error) {
	f := new(field.CrossType)
	err := m.Body.Get(f)
	return f, err
}

//CrossPrioritization is a required field for NewOrderCross.
func (m *NewOrderCross) CrossPrioritization() (*field.CrossPrioritization, error) {
	f := new(field.CrossPrioritization)
	err := m.Body.Get(f)
	return f, err
}

//NoSides is a required field for NewOrderCross.
func (m *NewOrderCross) NoSides() (*field.NoSides, error) {
	f := new(field.NoSides)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for NewOrderCross.
func (m *NewOrderCross) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for NewOrderCross.
func (m *NewOrderCross) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for NewOrderCross.
func (m *NewOrderCross) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for NewOrderCross.
func (m *NewOrderCross) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for NewOrderCross.
func (m *NewOrderCross) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for NewOrderCross.
func (m *NewOrderCross) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for NewOrderCross.
func (m *NewOrderCross) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for NewOrderCross.
func (m *NewOrderCross) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for NewOrderCross.
func (m *NewOrderCross) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for NewOrderCross.
func (m *NewOrderCross) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for NewOrderCross.
func (m *NewOrderCross) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for NewOrderCross.
func (m *NewOrderCross) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for NewOrderCross.
func (m *NewOrderCross) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for NewOrderCross.
func (m *NewOrderCross) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for NewOrderCross.
func (m *NewOrderCross) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for NewOrderCross.
func (m *NewOrderCross) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for NewOrderCross.
func (m *NewOrderCross) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for NewOrderCross.
func (m *NewOrderCross) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for NewOrderCross.
func (m *NewOrderCross) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for NewOrderCross.
func (m *NewOrderCross) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for NewOrderCross.
func (m *NewOrderCross) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for NewOrderCross.
func (m *NewOrderCross) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for NewOrderCross.
func (m *NewOrderCross) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for NewOrderCross.
func (m *NewOrderCross) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for NewOrderCross.
func (m *NewOrderCross) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for NewOrderCross.
func (m *NewOrderCross) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for NewOrderCross.
func (m *NewOrderCross) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for NewOrderCross.
func (m *NewOrderCross) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for NewOrderCross.
func (m *NewOrderCross) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for NewOrderCross.
func (m *NewOrderCross) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for NewOrderCross.
func (m *NewOrderCross) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for NewOrderCross.
func (m *NewOrderCross) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for NewOrderCross.
func (m *NewOrderCross) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//SettlmntTyp is a non-required field for NewOrderCross.
func (m *NewOrderCross) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//FutSettDate is a non-required field for NewOrderCross.
func (m *NewOrderCross) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//HandlInst is a required field for NewOrderCross.
func (m *NewOrderCross) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//ExecInst is a non-required field for NewOrderCross.
func (m *NewOrderCross) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//MinQty is a non-required field for NewOrderCross.
func (m *NewOrderCross) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//MaxFloor is a non-required field for NewOrderCross.
func (m *NewOrderCross) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//ExDestination is a non-required field for NewOrderCross.
func (m *NewOrderCross) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//NoTradingSessions is a non-required field for NewOrderCross.
func (m *NewOrderCross) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//ProcessCode is a non-required field for NewOrderCross.
func (m *NewOrderCross) ProcessCode() (*field.ProcessCode, error) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}

//PrevClosePx is a non-required field for NewOrderCross.
func (m *NewOrderCross) PrevClosePx() (*field.PrevClosePx, error) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}

//LocateReqd is a non-required field for NewOrderCross.
func (m *NewOrderCross) LocateReqd() (*field.LocateReqd, error) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for NewOrderCross.
func (m *NewOrderCross) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//NoStipulations is a non-required field for NewOrderCross.
func (m *NewOrderCross) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//OrdType is a required field for NewOrderCross.
func (m *NewOrderCross) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for NewOrderCross.
func (m *NewOrderCross) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//Price is a non-required field for NewOrderCross.
func (m *NewOrderCross) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//StopPx is a non-required field for NewOrderCross.
func (m *NewOrderCross) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//Spread is a non-required field for NewOrderCross.
func (m *NewOrderCross) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for NewOrderCross.
func (m *NewOrderCross) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveName is a non-required field for NewOrderCross.
func (m *NewOrderCross) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for NewOrderCross.
func (m *NewOrderCross) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//YieldType is a non-required field for NewOrderCross.
func (m *NewOrderCross) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//Yield is a non-required field for NewOrderCross.
func (m *NewOrderCross) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for NewOrderCross.
func (m *NewOrderCross) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//ComplianceID is a non-required field for NewOrderCross.
func (m *NewOrderCross) ComplianceID() (*field.ComplianceID, error) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//IOIid is a non-required field for NewOrderCross.
func (m *NewOrderCross) IOIid() (*field.IOIid, error) {
	f := new(field.IOIid)
	err := m.Body.Get(f)
	return f, err
}

//QuoteID is a non-required field for NewOrderCross.
func (m *NewOrderCross) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//TimeInForce is a non-required field for NewOrderCross.
func (m *NewOrderCross) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//EffectiveTime is a non-required field for NewOrderCross.
func (m *NewOrderCross) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//ExpireDate is a non-required field for NewOrderCross.
func (m *NewOrderCross) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//ExpireTime is a non-required field for NewOrderCross.
func (m *NewOrderCross) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GTBookingInst is a non-required field for NewOrderCross.
func (m *NewOrderCross) GTBookingInst() (*field.GTBookingInst, error) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//MaxShow is a non-required field for NewOrderCross.
func (m *NewOrderCross) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//PegDifference is a non-required field for NewOrderCross.
func (m *NewOrderCross) PegDifference() (*field.PegDifference, error) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionInst is a non-required field for NewOrderCross.
func (m *NewOrderCross) DiscretionInst() (*field.DiscretionInst, error) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//DiscretionOffset is a non-required field for NewOrderCross.
func (m *NewOrderCross) DiscretionOffset() (*field.DiscretionOffset, error) {
	f := new(field.DiscretionOffset)
	err := m.Body.Get(f)
	return f, err
}

//CancellationRights is a non-required field for NewOrderCross.
func (m *NewOrderCross) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for NewOrderCross.
func (m *NewOrderCross) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//RegistID is a non-required field for NewOrderCross.
func (m *NewOrderCross) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//Designation is a non-required field for NewOrderCross.
func (m *NewOrderCross) Designation() (*field.Designation, error) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestRate is a non-required field for NewOrderCross.
func (m *NewOrderCross) AccruedInterestRate() (*field.AccruedInterestRate, error) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestAmt is a non-required field for NewOrderCross.
func (m *NewOrderCross) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//NetMoney is a non-required field for NewOrderCross.
func (m *NewOrderCross) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}
