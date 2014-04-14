package fix40

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ListStatusRequest msg type = M.
type ListStatusRequest struct {
	message.Message
}

//ListStatusRequestBuilder builds ListStatusRequest messages.
type ListStatusRequestBuilder struct {
	message.MessageBuilder
}

//NewListStatusRequestBuilder returns an initialized ListStatusRequestBuilder with specified required fields.
func NewListStatusRequestBuilder(
	listid field.ListID) *ListStatusRequestBuilder {
	builder := new(ListStatusRequestBuilder)
	builder.Body.Set(listid)
	return builder
}

//ListID is a required field for ListStatusRequest.
func (m *ListStatusRequest) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//WaveNo is a non-required field for ListStatusRequest.
func (m *ListStatusRequest) WaveNo() (*field.WaveNo, error) {
	f := new(field.WaveNo)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for ListStatusRequest.
func (m *ListStatusRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
