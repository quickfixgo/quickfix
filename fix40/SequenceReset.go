package fix40

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type SequenceReset struct {
	message.Message
}

func (m *SequenceReset) GapFillFlag() (*field.GapFillFlag, error) {
	f := new(field.GapFillFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *SequenceReset) NewSeqNo() (*field.NewSeqNo, error) {
	f := new(field.NewSeqNo)
	err := m.Body.Get(f)
	return f, err
}
