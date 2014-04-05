package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type PositionReport struct {
	quickfixgo.Message
}

func (m *PositionReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SettlPrice() (*field.SettlPrice, error) {
	f := new(field.SettlPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) TotalNumPosReports() (*field.TotalNumPosReports, error) {
	f := new(field.TotalNumPosReports)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) UnsolicitedIndicator() (*field.UnsolicitedIndicator, error) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) PriorSettlPrice() (*field.PriorSettlPrice, error) {
	f := new(field.PriorSettlPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) PosMaintRptID() (*field.PosMaintRptID, error) {
	f := new(field.PosMaintRptID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) PosReqID() (*field.PosReqID, error) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) PosReqResult() (*field.PosReqResult, error) {
	f := new(field.PosReqResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) DeliveryDate() (*field.DeliveryDate, error) {
	f := new(field.DeliveryDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) PosReqType() (*field.PosReqType, error) {
	f := new(field.PosReqType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SettlPriceType() (*field.SettlPriceType, error) {
	f := new(field.SettlPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) NoPositions() (*field.NoPositions, error) {
	f := new(field.NoPositions)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) NoPosAmt() (*field.NoPosAmt, error) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) RegistStatus() (*field.RegistStatus, error) {
	f := new(field.RegistStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionReport) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}
