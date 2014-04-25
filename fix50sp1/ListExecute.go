package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ListExecute msg type = L.
type ListExecute struct {
	message.Message
}

//ListExecuteBuilder builds ListExecute messages.
type ListExecuteBuilder struct {
	message.MessageBuilder
}

//CreateListExecuteBuilder returns an initialized ListExecuteBuilder with specified required fields.
func CreateListExecuteBuilder(
	listid field.ListID,
	transacttime field.TransactTime) ListExecuteBuilder {
	var builder ListExecuteBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("L"))
	builder.Body.Set(listid)
	builder.Body.Set(transacttime)
	return builder
}

//ListID is a required field for ListExecute.
func (m ListExecute) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListExecute.
func (m ListExecute) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientBidID is a non-required field for ListExecute.
func (m ListExecute) ClientBidID() (*field.ClientBidID, errors.MessageRejectError) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}

//GetClientBidID reads a ClientBidID from ListExecute.
func (m ListExecute) GetClientBidID(f *field.ClientBidID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidID is a non-required field for ListExecute.
func (m ListExecute) BidID() (*field.BidID, errors.MessageRejectError) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}

//GetBidID reads a BidID from ListExecute.
func (m ListExecute) GetBidID(f *field.BidID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for ListExecute.
func (m ListExecute) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ListExecute.
func (m ListExecute) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ListExecute.
func (m ListExecute) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ListExecute.
func (m ListExecute) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ListExecute.
func (m ListExecute) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ListExecute.
func (m ListExecute) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ListExecute.
func (m ListExecute) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ListExecute.
func (m ListExecute) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
