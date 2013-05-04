package fix40

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type ResendRequest struct {
	quickfixgo.Message
}

func (m *ResendRequest) BeginSeqNo() (*field.BeginSeqNo, error) {
	f := new(field.BeginSeqNo)
	err := m.Body.Get(f)
	return f, err
}
func (m *ResendRequest) EndSeqNo() (*field.EndSeqNo, error) {
	f := new(field.EndSeqNo)
	err := m.Body.Get(f)
	return f, err
}
