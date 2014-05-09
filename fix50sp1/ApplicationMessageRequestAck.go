package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//ApplicationMessageRequestAck msg type = BX.
type ApplicationMessageRequestAck struct {
	message.Message
}

//ApplicationMessageRequestAckBuilder builds ApplicationMessageRequestAck messages.
type ApplicationMessageRequestAckBuilder struct {
	message.MessageBuilder
}

//CreateApplicationMessageRequestAckBuilder returns an initialized ApplicationMessageRequestAckBuilder with specified required fields.
func CreateApplicationMessageRequestAckBuilder(
	applresponseid *field.ApplResponseIDField) ApplicationMessageRequestAckBuilder {
	var builder ApplicationMessageRequestAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("BX"))
	builder.Body.Set(applresponseid)
	return builder
}

//ApplResponseID is a required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplResponseID() (*field.ApplResponseIDField, errors.MessageRejectError) {
	f := &field.ApplResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResponseID reads a ApplResponseID from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetApplResponseID(f *field.ApplResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReqID is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplReqID() (*field.ApplReqIDField, errors.MessageRejectError) {
	f := &field.ApplReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqID reads a ApplReqID from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetApplReqID(f *field.ApplReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReqType is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplReqType() (*field.ApplReqTypeField, errors.MessageRejectError) {
	f := &field.ApplReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqType reads a ApplReqType from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetApplReqType(f *field.ApplReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResponseType is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplResponseType() (*field.ApplResponseTypeField, errors.MessageRejectError) {
	f := &field.ApplResponseTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResponseType reads a ApplResponseType from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetApplResponseType(f *field.ApplResponseTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplTotalMessageCount is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplTotalMessageCount() (*field.ApplTotalMessageCountField, errors.MessageRejectError) {
	f := &field.ApplTotalMessageCountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplTotalMessageCount reads a ApplTotalMessageCount from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetApplTotalMessageCount(f *field.ApplTotalMessageCountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoApplIDs is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) NoApplIDs() (*field.NoApplIDsField, errors.MessageRejectError) {
	f := &field.NoApplIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoApplIDs reads a NoApplIDs from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetNoApplIDs(f *field.NoApplIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
