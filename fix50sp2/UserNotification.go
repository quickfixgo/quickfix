package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//UserNotification msg type = CB.
type UserNotification struct {
	message.Message
}

//UserNotificationBuilder builds UserNotification messages.
type UserNotificationBuilder struct {
	message.MessageBuilder
}

//NewUserNotificationBuilder returns an initialized UserNotificationBuilder with specified required fields.
func NewUserNotificationBuilder(
	userstatus field.UserStatus) *UserNotificationBuilder {
	builder := new(UserNotificationBuilder)
	builder.Body.Set(userstatus)
	return builder
}

//Username is a non-required field for UserNotification.
func (m *UserNotification) Username() (*field.Username, error) {
	f := new(field.Username)
	err := m.Body.Get(f)
	return f, err
}

//UserStatus is a required field for UserNotification.
func (m *UserNotification) UserStatus() (*field.UserStatus, error) {
	f := new(field.UserStatus)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for UserNotification.
func (m *UserNotification) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for UserNotification.
func (m *UserNotification) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for UserNotification.
func (m *UserNotification) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
