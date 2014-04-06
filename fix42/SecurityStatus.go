package fix42

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type SecurityStatus struct {
	quickfix.Message
}

func (m *SecurityStatus) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) MaturityDay() (*field.MaturityDay, error) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) SecurityTradingStatus() (*field.SecurityTradingStatus, error) {
	f := new(field.SecurityTradingStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) BuyVolume() (*field.BuyVolume, error) {
	f := new(field.BuyVolume)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) HighPx() (*field.HighPx, error) {
	f := new(field.HighPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) UnsolicitedIndicator() (*field.UnsolicitedIndicator, error) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) DueToRelated() (*field.DueToRelated, error) {
	f := new(field.DueToRelated)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) Adjustment() (*field.Adjustment, error) {
	f := new(field.Adjustment)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) SecurityStatusReqID() (*field.SecurityStatusReqID, error) {
	f := new(field.SecurityStatusReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) SellVolume() (*field.SellVolume, error) {
	f := new(field.SellVolume)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) FinancialStatus() (*field.FinancialStatus, error) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) HaltReasonChar() (*field.HaltReasonChar, error) {
	f := new(field.HaltReasonChar)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) InViewOfCommon() (*field.InViewOfCommon, error) {
	f := new(field.InViewOfCommon)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) CorporateAction() (*field.CorporateAction, error) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) LowPx() (*field.LowPx, error) {
	f := new(field.LowPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityStatus) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
