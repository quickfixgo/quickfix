package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//ApplicationMessageRequest msg type = BW.
type ApplicationMessageRequest struct {
	message.Message
}

//ApplicationMessageRequestBuilder builds ApplicationMessageRequest messages.
type ApplicationMessageRequestBuilder struct {
	message.MessageBuilder
}

//CreateApplicationMessageRequestBuilder returns an initialized ApplicationMessageRequestBuilder with specified required fields.
func CreateApplicationMessageRequestBuilder(
	applreqid *field.ApplReqIDField,
	applreqtype *field.ApplReqTypeField) ApplicationMessageRequestBuilder {
	var builder ApplicationMessageRequestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("BW"))
	builder.Body.Set(applreqid)
	builder.Body.Set(applreqtype)
	return builder
}

//ApplReqID is a required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) ApplReqID() (*field.ApplReqIDField, errors.MessageRejectError) {
	f := &field.ApplReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqID reads a ApplReqID from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetApplReqID(f *field.ApplReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReqType is a required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) ApplReqType() (*field.ApplReqTypeField, errors.MessageRejectError) {
	f := &field.ApplReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqType reads a ApplReqType from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetApplReqType(f *field.ApplReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoApplIDs is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) NoApplIDs() (*field.NoApplIDsField, errors.MessageRejectError) {
	f := &field.NoApplIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoApplIDs reads a NoApplIDs from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetNoApplIDs(f *field.NoApplIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
