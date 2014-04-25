package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
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

//CreateListCancelRequestBuilder returns an initialized ListCancelRequestBuilder with specified required fields.
func CreateListCancelRequestBuilder(
	listid field.ListID,
	transacttime field.TransactTime) ListCancelRequestBuilder {
	var builder ListCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("K"))
	builder.Body.Set(listid)
	builder.Body.Set(transacttime)
	return builder
}

//ListID is a required field for ListCancelRequest.
func (m ListCancelRequest) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListCancelRequest.
func (m ListCancelRequest) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for ListCancelRequest.
func (m ListCancelRequest) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ListCancelRequest.
func (m ListCancelRequest) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ListCancelRequest.
func (m ListCancelRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ListCancelRequest.
func (m ListCancelRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ListCancelRequest.
func (m ListCancelRequest) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ListCancelRequest.
func (m ListCancelRequest) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ListCancelRequest.
func (m ListCancelRequest) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ListCancelRequest.
func (m ListCancelRequest) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
