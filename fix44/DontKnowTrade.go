package fix44

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
func (m *DontKnowTrade) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
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
func (m *DontKnowTrade) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) LastQty() (*field.LastQty, error) {
	f := new(field.LastQty)
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
func (m *DontKnowTrade) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *DontKnowTrade) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
