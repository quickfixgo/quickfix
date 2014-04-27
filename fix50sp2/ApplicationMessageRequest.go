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
	applreqid field.ApplReqID,
	applreqtype field.ApplReqType) ApplicationMessageRequestBuilder {
	var builder ApplicationMessageRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.BuildMsgType("BW"))
	builder.Body.Set(applreqid)
	builder.Body.Set(applreqtype)
	return builder
}

//ApplReqID is a required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) ApplReqID() (*field.ApplReqID, errors.MessageRejectError) {
	f := new(field.ApplReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqID reads a ApplReqID from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetApplReqID(f *field.ApplReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReqType is a required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) ApplReqType() (*field.ApplReqType, errors.MessageRejectError) {
	f := new(field.ApplReqType)
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqType reads a ApplReqType from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetApplReqType(f *field.ApplReqType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoApplIDs is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) NoApplIDs() (*field.NoApplIDs, errors.MessageRejectError) {
	f := new(field.NoApplIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoApplIDs reads a NoApplIDs from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetNoApplIDs(f *field.NoApplIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from ApplicationMessageRequest.
func (m ApplicationMessageRequest) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}
