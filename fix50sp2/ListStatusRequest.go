package fix50sp2

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

//CreateListStatusRequestBuilder returns an initialized ListStatusRequestBuilder with specified required fields.
func CreateListStatusRequestBuilder(
	listid field.ListID) ListStatusRequestBuilder {
	var builder ListStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(listid)
	return builder
}

//ListID is a required field for ListStatusRequest.
func (m ListStatusRequest) ListID() (field.ListID, error) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for ListStatusRequest.
func (m ListStatusRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for ListStatusRequest.
func (m ListStatusRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for ListStatusRequest.
func (m ListStatusRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
