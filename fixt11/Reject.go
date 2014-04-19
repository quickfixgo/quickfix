package fixt11

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
func (m Reject) RefSeqNum() (field.RefSeqNum, errors.MessageRejectError) {
	var f field.RefSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//RefTagID is a non-required field for Reject.
func (m Reject) RefTagID() (field.RefTagID, errors.MessageRejectError) {
	var f field.RefTagID
	err := m.Body.Get(&f)
	return f, err
}

//RefMsgType is a non-required field for Reject.
func (m Reject) RefMsgType() (field.RefMsgType, errors.MessageRejectError) {
	var f field.RefMsgType
	err := m.Body.Get(&f)
	return f, err
}

//SessionRejectReason is a non-required field for Reject.
func (m Reject) SessionRejectReason() (field.SessionRejectReason, errors.MessageRejectError) {
	var f field.SessionRejectReason
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for Reject.
func (m Reject) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for Reject.
func (m Reject) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for Reject.
func (m Reject) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
