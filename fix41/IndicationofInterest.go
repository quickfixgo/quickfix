package fix41

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type IndicationofInterest struct {
	quickfix.Message
}

func (m *IndicationofInterest) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) NoIOIQualifiers() (*field.NoIOIQualifiers, error) {
	f := new(field.NoIOIQualifiers)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOITransType() (*field.IOITransType, error) {
	f := new(field.IOITransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOIQltyInd() (*field.IOIQltyInd, error) {
	f := new(field.IOIQltyInd)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOIOthSvc() (*field.IOIOthSvc, error) {
	f := new(field.IOIOthSvc)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) MaturityDay() (*field.MaturityDay, error) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) ValidUntilTime() (*field.ValidUntilTime, error) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOIRefID() (*field.IOIRefID, error) {
	f := new(field.IOIRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) URLLink() (*field.URLLink, error) {
	f := new(field.URLLink)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOIid() (*field.IOIid, error) {
	f := new(field.IOIid)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOINaturalFlag() (*field.IOINaturalFlag, error) {
	f := new(field.IOINaturalFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) IOIShares() (*field.IOIShares, error) {
	f := new(field.IOIShares)
	err := m.Body.Get(f)
	return f, err
}
func (m *IndicationofInterest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
