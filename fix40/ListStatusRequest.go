package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Header.Set(field.BuildMsgType("M"))
	builder.Body.Set(listid)
	return builder
}

//ListID is a required field for ListStatusRequest.
func (m ListStatusRequest) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListStatusRequest.
func (m ListStatusRequest) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WaveNo is a non-required field for ListStatusRequest.
func (m ListStatusRequest) WaveNo() (*field.WaveNo, errors.MessageRejectError) {
	f := new(field.WaveNo)
	err := m.Body.Get(f)
	return f, err
}

//GetWaveNo reads a WaveNo from ListStatusRequest.
func (m ListStatusRequest) GetWaveNo(f *field.WaveNo) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ListStatusRequest.
func (m ListStatusRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ListStatusRequest.
func (m ListStatusRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}
