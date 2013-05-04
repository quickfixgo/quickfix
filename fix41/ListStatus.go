package fix41

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type ListStatus struct {
	quickfixgo.Message
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
