package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	refmsgtype *field.RefMsgTypeField,
	businessrejectreason *field.BusinessRejectReasonField) BusinessMessageRejectBuilder {
	var builder BusinessMessageRejectBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("j"))
	builder.Body().Set(refmsgtype)
	builder.Body().Set(businessrejectreason)
	return builder
}

//RefSeqNum is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) RefSeqNum() (*field.RefSeqNumField, errors.MessageRejectError) {
	f := &field.RefSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefSeqNum reads a RefSeqNum from BusinessMessageReject.
func (m BusinessMessageReject) GetRefSeqNum(f *field.RefSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefMsgType is a required field for BusinessMessageReject.
func (m BusinessMessageReject) RefMsgType() (*field.RefMsgTypeField, errors.MessageRejectError) {
	f := &field.RefMsgTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefMsgType reads a RefMsgType from BusinessMessageReject.
func (m BusinessMessageReject) GetRefMsgType(f *field.RefMsgTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BusinessRejectRefID is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) BusinessRejectRefID() (*field.BusinessRejectRefIDField, errors.MessageRejectError) {
	f := &field.BusinessRejectRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBusinessRejectRefID reads a BusinessRejectRefID from BusinessMessageReject.
func (m BusinessMessageReject) GetBusinessRejectRefID(f *field.BusinessRejectRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BusinessRejectReason is a required field for BusinessMessageReject.
func (m BusinessMessageReject) BusinessRejectReason() (*field.BusinessRejectReasonField, errors.MessageRejectError) {
	f := &field.BusinessRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBusinessRejectReason reads a BusinessRejectReason from BusinessMessageReject.
func (m BusinessMessageReject) GetBusinessRejectReason(f *field.BusinessRejectReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from BusinessMessageReject.
func (m BusinessMessageReject) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from BusinessMessageReject.
func (m BusinessMessageReject) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from BusinessMessageReject.
func (m BusinessMessageReject) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
