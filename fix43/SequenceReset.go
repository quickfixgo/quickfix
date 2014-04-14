package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SequenceReset msg type = 4.
type SequenceReset struct {
	message.Message
}

//SequenceResetBuilder builds SequenceReset messages.
type SequenceResetBuilder struct {
	message.MessageBuilder
}

//NewSequenceResetBuilder returns an initialized SequenceResetBuilder with specified required fields.
func NewSequenceResetBuilder(
	newseqno field.NewSeqNo) *SequenceResetBuilder {
	builder := new(SequenceResetBuilder)
	builder.Body.Set(newseqno)
	return builder
}

//GapFillFlag is a non-required field for SequenceReset.
func (m *SequenceReset) GapFillFlag() (*field.GapFillFlag, error) {
	f := new(field.GapFillFlag)
	err := m.Body.Get(f)
	return f, err
}

//NewSeqNo is a required field for SequenceReset.
func (m *SequenceReset) NewSeqNo() (*field.NewSeqNo, error) {
	f := new(field.NewSeqNo)
	err := m.Body.Get(f)
	return f, err
}
