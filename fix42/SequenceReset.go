package fix42

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type SequenceReset struct {
	quickfix.Message
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
