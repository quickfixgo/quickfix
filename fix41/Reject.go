package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Reject msg type = 3.
type Reject struct {
	message.Message
}

//RejectBuilder builds Reject messages.
type RejectBuilder struct {
	message.MessageBuilder
}

//CreateRejectBuilder returns an initialized RejectBuilder with specified required fields.
func CreateRejectBuilder(
	refseqnum field.RefSeqNum) RejectBuilder {
	var builder RejectBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(refseqnum)
	return builder
}

//RefSeqNum is a required field for Reject.
func (m Reject) RefSeqNum() (*field.RefSeqNum, errors.MessageRejectError) {
	f := new(field.RefSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetRefSeqNum reads a RefSeqNum from Reject.
func (m Reject) GetRefSeqNum(f *field.RefSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for Reject.
func (m Reject) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Reject.
func (m Reject) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}
