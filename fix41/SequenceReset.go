package fix41

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type SequenceReset struct {
	quickfixgo.Message
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
