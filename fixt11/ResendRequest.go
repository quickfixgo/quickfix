package fixt11

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ResendRequest msg type = 2.
type ResendRequest struct {
	message.Message
}

//ResendRequestBuilder builds ResendRequest messages.
type ResendRequestBuilder struct {
	message.MessageBuilder
}

//CreateResendRequestBuilder returns an initialized ResendRequestBuilder with specified required fields.
func CreateResendRequestBuilder(
	beginseqno *field.BeginSeqNoField,
	endseqno *field.EndSeqNoField) ResendRequestBuilder {
	var builder ResendRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewMsgType("2"))
	builder.Body.Set(beginseqno)
	builder.Body.Set(endseqno)
	return builder
}

//BeginSeqNo is a required field for ResendRequest.
func (m ResendRequest) BeginSeqNo() (*field.BeginSeqNoField, errors.MessageRejectError) {
	f := &field.BeginSeqNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBeginSeqNo reads a BeginSeqNo from ResendRequest.
func (m ResendRequest) GetBeginSeqNo(f *field.BeginSeqNoField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndSeqNo is a required field for ResendRequest.
func (m ResendRequest) EndSeqNo() (*field.EndSeqNoField, errors.MessageRejectError) {
	f := &field.EndSeqNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndSeqNo reads a EndSeqNo from ResendRequest.
func (m ResendRequest) GetEndSeqNo(f *field.EndSeqNoField) errors.MessageRejectError {
	return m.Body.Get(f)
}
