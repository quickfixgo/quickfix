package fix42

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type IndicationofInterest struct {
	quickfix.Message
}

func (m *IndicationofInterest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) NoRoutingIDs() (*field.NoRoutingIDs, error) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOIQltyInd() (*field.IOIQltyInd, error) {
	f := new(field.IOIQltyInd)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) SpreadToBenchmark() (*field.SpreadToBenchmark, error) {
	f := new(field.SpreadToBenchmark)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOIShares() (*field.IOIShares, error) {
	f := new(field.IOIShares)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOIRefID() (*field.IOIRefID, error) {
	f := new(field.IOIRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) ValidUntilTime() (*field.ValidUntilTime, error) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOIid() (*field.IOIid, error) {
	f := new(field.IOIid)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOINaturalFlag() (*field.IOINaturalFlag, error) {
	f := new(field.IOINaturalFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) NoIOIQualifiers() (*field.NoIOIQualifiers, error) {
	f := new(field.NoIOIQualifiers)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) Benchmark() (*field.Benchmark, error) {
	f := new(field.Benchmark)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) MaturityDay() (*field.MaturityDay, error) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) URLLink() (*field.URLLink, error) {
	f := new(field.URLLink)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOITransType() (*field.IOITransType, error) {
	f := new(field.IOITransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
