package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type UserNotification struct {
	quickfixgo.Message
}

func (m *UserNotification) UserStatus() (*field.UserStatus, error) {
	f := new(field.UserStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserNotification) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserNotification) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserNotification) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
