package fix41

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
	listid field.ListID) *ListExecuteBuilder {
	builder := new(ListExecuteBuilder)
	builder.Body.Set(listid)
	return builder
}

//ListID is a required field for ListExecute.
func (m *ListExecute) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//WaveNo is a non-required field for ListExecute.
func (m *ListExecute) WaveNo() (*field.WaveNo, error) {
	f := new(field.WaveNo)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for ListExecute.
func (m *ListExecute) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
