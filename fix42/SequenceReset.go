package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Body.Set(newseqno)
	return builder
}

//GapFillFlag is a non-required field for SequenceReset.
func (m SequenceReset) GapFillFlag() (field.GapFillFlag, errors.MessageRejectError) {
	var f field.GapFillFlag
	err := m.Body.Get(&f)
	return f, err
}

//NewSeqNo is a required field for SequenceReset.
func (m SequenceReset) NewSeqNo() (field.NewSeqNo, errors.MessageRejectError) {
	var f field.NewSeqNo
	err := m.Body.Get(&f)
	return f, err
}
