package fix44

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
	builder.Header.Set(field.BuildMsgType("j"))
	builder.Body.Set(refmsgtype)
	builder.Body.Set(businessrejectreason)
	return builder
}

//RefSeqNum is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) RefSeqNum() (*field.RefSeqNum, errors.MessageRejectError) {
	f := new(field.RefSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetRefSeqNum reads a RefSeqNum from BusinessMessageReject.
func (m BusinessMessageReject) GetRefSeqNum(f *field.RefSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefMsgType is a required field for BusinessMessageReject.
func (m BusinessMessageReject) RefMsgType() (*field.RefMsgType, errors.MessageRejectError) {
	f := new(field.RefMsgType)
	err := m.Body.Get(f)
	return f, err
}

//GetRefMsgType reads a RefMsgType from BusinessMessageReject.
func (m BusinessMessageReject) GetRefMsgType(f *field.RefMsgType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BusinessRejectRefID is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) BusinessRejectRefID() (*field.BusinessRejectRefID, errors.MessageRejectError) {
	f := new(field.BusinessRejectRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetBusinessRejectRefID reads a BusinessRejectRefID from BusinessMessageReject.
func (m BusinessMessageReject) GetBusinessRejectRefID(f *field.BusinessRejectRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BusinessRejectReason is a required field for BusinessMessageReject.
func (m BusinessMessageReject) BusinessRejectReason() (*field.BusinessRejectReason, errors.MessageRejectError) {
	f := new(field.BusinessRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//GetBusinessRejectReason reads a BusinessRejectReason from BusinessMessageReject.
func (m BusinessMessageReject) GetBusinessRejectReason(f *field.BusinessRejectReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from BusinessMessageReject.
func (m BusinessMessageReject) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from BusinessMessageReject.
func (m BusinessMessageReject) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for BusinessMessageReject.
func (m BusinessMessageReject) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from BusinessMessageReject.
func (m BusinessMessageReject) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
