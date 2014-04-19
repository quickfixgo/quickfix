package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	builder.Body.Set(applreqid)
	builder.Body.Set(applreqtype)
	return builder
}

//ApplReqID is a required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) ApplReqID() (field.ApplReqID, errors.MessageRejectError) {
	var f field.ApplReqID
	err := m.Body.Get(&f)
	return f, err
}

//ApplReqType is a required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) ApplReqType() (field.ApplReqType, errors.MessageRejectError) {
	var f field.ApplReqType
	err := m.Body.Get(&f)
	return f, err
}

//NoApplIDs is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) NoApplIDs() (field.NoApplIDs, errors.MessageRejectError) {
	var f field.NoApplIDs
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for ApplicationMessageRequest.
func (m ApplicationMessageRequest) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
