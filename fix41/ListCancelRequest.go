package fix41

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ListCancelRequest msg type = K.
type ListCancelRequest struct {
	message.Message
}

//ListCancelRequestBuilder builds ListCancelRequest messages.
type ListCancelRequestBuilder struct {
	message.MessageBuilder
}

//NewListCancelRequestBuilder returns an initialized ListCancelRequestBuilder with specified required fields.
func NewListCancelRequestBuilder(
	listid field.ListID) *ListCancelRequestBuilder {
	builder := new(ListCancelRequestBuilder)
	builder.Body.Set(listid)
	return builder
}

//ListID is a required field for ListCancelRequest.
func (m *ListCancelRequest) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//WaveNo is a non-required field for ListCancelRequest.
func (m *ListCancelRequest) WaveNo() (*field.WaveNo, error) {
	f := new(field.WaveNo)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for ListCancelRequest.
func (m *ListCancelRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
