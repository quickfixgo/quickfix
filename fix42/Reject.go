package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	refseqnum *field.RefSeqNumField) RejectBuilder {
	var builder RejectBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header().Set(field.NewMsgType("3"))
	builder.Body().Set(refseqnum)
	return builder
}

//RefSeqNum is a required field for Reject.
func (m Reject) RefSeqNum() (*field.RefSeqNumField, errors.MessageRejectError) {
	f := &field.RefSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefSeqNum reads a RefSeqNum from Reject.
func (m Reject) GetRefSeqNum(f *field.RefSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefTagID is a non-required field for Reject.
func (m Reject) RefTagID() (*field.RefTagIDField, errors.MessageRejectError) {
	f := &field.RefTagIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefTagID reads a RefTagID from Reject.
func (m Reject) GetRefTagID(f *field.RefTagIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefMsgType is a non-required field for Reject.
func (m Reject) RefMsgType() (*field.RefMsgTypeField, errors.MessageRejectError) {
	f := &field.RefMsgTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefMsgType reads a RefMsgType from Reject.
func (m Reject) GetRefMsgType(f *field.RefMsgTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SessionRejectReason is a non-required field for Reject.
func (m Reject) SessionRejectReason() (*field.SessionRejectReasonField, errors.MessageRejectError) {
	f := &field.SessionRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSessionRejectReason reads a SessionRejectReason from Reject.
func (m Reject) GetSessionRejectReason(f *field.SessionRejectReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for Reject.
func (m Reject) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Reject.
func (m Reject) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for Reject.
func (m Reject) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from Reject.
func (m Reject) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for Reject.
func (m Reject) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from Reject.
func (m Reject) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
