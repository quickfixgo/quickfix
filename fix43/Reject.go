package fix43

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
	builder.Header.Set(field.BuildMsgType("3"))
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

//RefTagID is a non-required field for Reject.
func (m Reject) RefTagID() (*field.RefTagID, errors.MessageRejectError) {
	f := new(field.RefTagID)
	err := m.Body.Get(f)
	return f, err
}

//GetRefTagID reads a RefTagID from Reject.
func (m Reject) GetRefTagID(f *field.RefTagID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefMsgType is a non-required field for Reject.
func (m Reject) RefMsgType() (*field.RefMsgType, errors.MessageRejectError) {
	f := new(field.RefMsgType)
	err := m.Body.Get(f)
	return f, err
}

//GetRefMsgType reads a RefMsgType from Reject.
func (m Reject) GetRefMsgType(f *field.RefMsgType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SessionRejectReason is a non-required field for Reject.
func (m Reject) SessionRejectReason() (*field.SessionRejectReason, errors.MessageRejectError) {
	f := new(field.SessionRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//GetSessionRejectReason reads a SessionRejectReason from Reject.
func (m Reject) GetSessionRejectReason(f *field.SessionRejectReason) errors.MessageRejectError {
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

//EncodedTextLen is a non-required field for Reject.
func (m Reject) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from Reject.
func (m Reject) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for Reject.
func (m Reject) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from Reject.
func (m Reject) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
