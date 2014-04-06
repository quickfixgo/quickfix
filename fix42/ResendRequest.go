package fix42

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type ResendRequest struct {
	quickfix.Message
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
