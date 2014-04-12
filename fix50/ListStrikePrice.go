package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type ListStrikePrice struct {
	message.Message
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
func (m *ListStrikePrice) NoStrikes() (*field.NoStrikes, error) {
	f := new(field.NoStrikes)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStrikePrice) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}
