package fix42

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
func (m *ListStrikePrice) NoStrikes() (*field.NoStrikes, error) {
	f := new(field.NoStrikes)
	err := m.Body.Get(f)
	return f, err
}
