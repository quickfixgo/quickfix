package fix40

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type DontKnowTrade struct {
	quickfixgo.Message
}

func (m *DontKnowTrade) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
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
func (m *DontKnowTrade) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) LastShares() (*field.LastShares, error) {
	f := new(field.LastShares)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
