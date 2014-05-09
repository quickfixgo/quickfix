package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	listid *field.ListIDField) ListCancelRequestBuilder {
	var builder ListCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header.Set(field.NewMsgType("K"))
	builder.Body.Set(listid)
	return builder
}

//ListID is a required field for ListCancelRequest.
func (m ListCancelRequest) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListCancelRequest.
func (m ListCancelRequest) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WaveNo is a non-required field for ListCancelRequest.
func (m ListCancelRequest) WaveNo() (*field.WaveNoField, errors.MessageRejectError) {
	f := &field.WaveNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWaveNo reads a WaveNo from ListCancelRequest.
func (m ListCancelRequest) GetWaveNo(f *field.WaveNoField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ListCancelRequest.
func (m ListCancelRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ListCancelRequest.
func (m ListCancelRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
