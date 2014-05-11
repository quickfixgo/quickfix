package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	listid *field.ListIDField) ListStatusRequestBuilder {
	var builder ListStatusRequestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX41))
	builder.Header().Set(field.NewMsgType("M"))
	builder.Body().Set(listid)
	return builder
}

//ListID is a required field for ListStatusRequest.
func (m ListStatusRequest) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListStatusRequest.
func (m ListStatusRequest) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WaveNo is a non-required field for ListStatusRequest.
func (m ListStatusRequest) WaveNo() (*field.WaveNoField, errors.MessageRejectError) {
	f := &field.WaveNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWaveNo reads a WaveNo from ListStatusRequest.
func (m ListStatusRequest) GetWaveNo(f *field.WaveNoField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ListStatusRequest.
func (m ListStatusRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ListStatusRequest.
func (m ListStatusRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
