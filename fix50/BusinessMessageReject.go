package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//BusinessMessageReject msg type = j.
type BusinessMessageReject struct {
	message.Message
}

//BusinessMessageRejectBuilder builds BusinessMessageReject messages.
type BusinessMessageRejectBuilder struct {
	message.MessageBuilder
}

//CreateBusinessMessageRejectBuilder returns an initialized BusinessMessageRejectBuilder with specified required fields.
func CreateBusinessMessageRejectBuilder(
	refmsgtype field.RefMsgType,
	businessrejectreason field.BusinessRejectReason) BusinessMessageRejectBuilder {
	var builder BusinessMessageRejectBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(refmsgtype)
	builder.Body.Set(businessrejectreason)
	return builder
}

//RefSeqNum is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) RefSeqNum() (field.RefSeqNum, errors.MessageRejectError) {
	var f field.RefSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//RefMsgType is a required field for BusinessMessageReject.
func (m BusinessMessageReject) RefMsgType() (field.RefMsgType, errors.MessageRejectError) {
	var f field.RefMsgType
	err := m.Body.Get(&f)
	return f, err
}

//BusinessRejectRefID is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) BusinessRejectRefID() (field.BusinessRejectRefID, errors.MessageRejectError) {
	var f field.BusinessRejectRefID
	err := m.Body.Get(&f)
	return f, err
}

//BusinessRejectReason is a required field for BusinessMessageReject.
func (m BusinessMessageReject) BusinessRejectReason() (field.BusinessRejectReason, errors.MessageRejectError) {
	var f field.BusinessRejectReason
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
