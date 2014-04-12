package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type ResendRequest struct {
	message.Message
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
