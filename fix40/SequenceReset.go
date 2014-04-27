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
	newseqno field.NewSeqNo) SequenceResetBuilder {
	var builder SequenceResetBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIX40))
	builder.Header.Set(field.BuildMsgType("4"))
	builder.Body.Set(newseqno)
	return builder
}

//GapFillFlag is a non-required field for SequenceReset.
func (m SequenceReset) GapFillFlag() (*field.GapFillFlag, errors.MessageRejectError) {
	f := new(field.GapFillFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetGapFillFlag reads a GapFillFlag from SequenceReset.
func (m SequenceReset) GetGapFillFlag(f *field.GapFillFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NewSeqNo is a required field for SequenceReset.
func (m SequenceReset) NewSeqNo() (*field.NewSeqNo, errors.MessageRejectError) {
	f := new(field.NewSeqNo)
	err := m.Body.Get(f)
	return f, err
}

//GetNewSeqNo reads a NewSeqNo from SequenceReset.
func (m SequenceReset) GetNewSeqNo(f *field.NewSeqNo) errors.MessageRejectError {
	return m.Body.Get(f)
}
