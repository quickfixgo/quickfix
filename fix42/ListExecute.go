package fix42

import (
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

//NewListExecuteBuilder returns an initialized ListExecuteBuilder with specified required fields.
func NewListExecuteBuilder(
	listid field.ListID,
	transacttime field.TransactTime) *ListExecuteBuilder {
	builder := new(ListExecuteBuilder)
	builder.Body.Set(listid)
	builder.Body.Set(transacttime)
	return builder
}

//ListID is a required field for ListExecute.
func (m *ListExecute) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//ClientBidID is a non-required field for ListExecute.
func (m *ListExecute) ClientBidID() (*field.ClientBidID, error) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}

//BidID is a non-required field for ListExecute.
func (m *ListExecute) BidID() (*field.BidID, error) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for ListExecute.
func (m *ListExecute) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for ListExecute.
func (m *ListExecute) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for ListExecute.
func (m *ListExecute) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for ListExecute.
func (m *ListExecute) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
