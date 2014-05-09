package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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

//CreateSequenceResetBuilder returns an initialized SequenceResetBuilder with specified required fields.
func CreateSequenceResetBuilder(
	newseqno *field.NewSeqNoField) SequenceResetBuilder {
	var builder SequenceResetBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header.Set(field.NewMsgType("4"))
	builder.Body.Set(newseqno)
	return builder
}

//GapFillFlag is a non-required field for SequenceReset.
func (m SequenceReset) GapFillFlag() (*field.GapFillFlagField, errors.MessageRejectError) {
	f := &field.GapFillFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGapFillFlag reads a GapFillFlag from SequenceReset.
func (m SequenceReset) GetGapFillFlag(f *field.GapFillFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NewSeqNo is a required field for SequenceReset.
func (m SequenceReset) NewSeqNo() (*field.NewSeqNoField, errors.MessageRejectError) {
	f := &field.NewSeqNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNewSeqNo reads a NewSeqNo from SequenceReset.
func (m SequenceReset) GetNewSeqNo(f *field.NewSeqNoField) errors.MessageRejectError {
	return m.Body.Get(f)
}
