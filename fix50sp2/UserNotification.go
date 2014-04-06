package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type UserNotification struct {
	quickfix.Message
}

func (m *UserNotification) Username() (*field.Username, error) {
	f := new(field.Username)
	err := m.Body.Get(f)
	return f, err
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
