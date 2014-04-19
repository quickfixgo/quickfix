package fix41

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
	listid field.ListID) ListExecuteBuilder {
	var builder ListExecuteBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(listid)
	return builder
}

//ListID is a required field for ListExecute.
func (m ListExecute) ListID() (field.ListID, errors.MessageRejectError) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//WaveNo is a non-required field for ListExecute.
func (m ListExecute) WaveNo() (field.WaveNo, errors.MessageRejectError) {
	var f field.WaveNo
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for ListExecute.
func (m ListExecute) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}
