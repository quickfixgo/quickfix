package fix41

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type ListStatus struct {
	quickfix.Message
}

func (m *ListStatus) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) WaveNo() (*field.WaveNo, error) {
	f := new(field.WaveNo)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) NoRpts() (*field.NoRpts, error) {
	f := new(field.NoRpts)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) RptSeq() (*field.RptSeq, error) {
	f := new(field.RptSeq)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}
