package fix41

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type ListExecute struct {
	quickfix.Message
}

func (m *ListExecute) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListExecute) WaveNo() (*field.WaveNo, error) {
	f := new(field.WaveNo)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListExecute) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
