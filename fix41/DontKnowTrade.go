package fix41

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type DontKnowTrade struct {
	quickfixgo.Message
}

func (m *DontKnowTrade) MaturityDay() (*field.MaturityDay, error) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) LastShares() (*field.LastShares, error) {
	f := new(field.LastShares)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) DKReason() (*field.DKReason, error) {
	f := new(field.DKReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}
