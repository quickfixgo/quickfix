package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
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

//EncodedTextLen is a non-required field for ListStatusRequest.
func (m ListStatusRequest) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ListStatusRequest.
func (m ListStatusRequest) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ListStatusRequest.
func (m ListStatusRequest) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ListStatusRequest.
func (m ListStatusRequest) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
