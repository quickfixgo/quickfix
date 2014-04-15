package fix40

import (
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
func (m Reject) RefSeqNum() (field.RefSeqNum, error) {
	var f field.RefSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for Reject.
func (m Reject) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}
