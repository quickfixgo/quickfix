package fix40

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type ListStatusRequest struct {
	quickfixgo.Message
}

func (m *ListStatusRequest) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatusRequest) WaveNo() (*field.WaveNo, error) {
	f := new(field.WaveNo)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatusRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
