package fix50

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
	builder.Body.Set(listid)
	builder.Body.Set(transacttime)
	return builder
}

//ListID is a required field for ListExecute.
func (m ListExecute) ListID() (field.ListID, errors.MessageRejectError) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//ClientBidID is a non-required field for ListExecute.
func (m ListExecute) ClientBidID() (field.ClientBidID, errors.MessageRejectError) {
	var f field.ClientBidID
	err := m.Body.Get(&f)
	return f, err
}

//BidID is a non-required field for ListExecute.
func (m ListExecute) BidID() (field.BidID, errors.MessageRejectError) {
	var f field.BidID
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for ListExecute.
func (m ListExecute) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for ListExecute.
func (m ListExecute) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for ListExecute.
func (m ListExecute) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for ListExecute.
func (m ListExecute) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
