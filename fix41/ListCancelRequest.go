package fix41

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type ListCancelRequest struct {
	quickfixgo.Message
}

func (m *ListCancelRequest) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListCancelRequest) WaveNo() (*field.WaveNo, error) {
	f := new(field.WaveNo)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListCancelRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
