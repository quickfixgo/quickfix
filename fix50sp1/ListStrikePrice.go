package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type ListStrikePrice struct {
	quickfixgo.Message
}

func (m *ListStrikePrice) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStrikePrice) TotNoStrikes() (*field.TotNoStrikes, error) {
	f := new(field.TotNoStrikes)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStrikePrice) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}
