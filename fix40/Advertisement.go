package fix40

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type Advertisement struct {
	quickfix.Message
}

func (m *Advertisement) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) AdvId() (*field.AdvId, error) {
	f := new(field.AdvId)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) AdvTransType() (*field.AdvTransType, error) {
	f := new(field.AdvTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) AdvSide() (*field.AdvSide, error) {
	f := new(field.AdvSide)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) Shares() (*field.Shares, error) {
	f := new(field.Shares)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) AdvRefID() (*field.AdvRefID, error) {
	f := new(field.AdvRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Advertisement) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
