package fix41

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
	listid field.ListID) ListCancelRequestBuilder {
	var builder ListCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("K"))
	builder.Body.Set(listid)
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

//WaveNo is a non-required field for ListCancelRequest.
func (m ListCancelRequest) WaveNo() (*field.WaveNo, errors.MessageRejectError) {
	f := new(field.WaveNo)
	err := m.Body.Get(f)
	return f, err
}

//GetWaveNo reads a WaveNo from ListCancelRequest.
func (m ListCancelRequest) GetWaveNo(f *field.WaveNo) errors.MessageRejectError {
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
