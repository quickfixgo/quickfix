package fix44

import (
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

//NewResendRequestBuilder returns an initialized ResendRequestBuilder with specified required fields.
func NewResendRequestBuilder(
	beginseqno field.BeginSeqNo,
	endseqno field.EndSeqNo) *ResendRequestBuilder {
	builder := new(ResendRequestBuilder)
	builder.Body.Set(beginseqno)
	builder.Body.Set(endseqno)
	return builder
}

//BeginSeqNo is a required field for ResendRequest.
func (m *ResendRequest) BeginSeqNo() (*field.BeginSeqNo, error) {
	f := new(field.BeginSeqNo)
	err := m.Body.Get(f)
	return f, err
}

//EndSeqNo is a required field for ResendRequest.
func (m *ResendRequest) EndSeqNo() (*field.EndSeqNo, error) {
	f := new(field.EndSeqNo)
	err := m.Body.Get(f)
	return f, err
}
