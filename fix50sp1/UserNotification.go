package fix50sp1

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

//CreateUserNotificationBuilder returns an initialized UserNotificationBuilder with specified required fields.
func CreateUserNotificationBuilder(
	userstatus field.UserStatus) UserNotificationBuilder {
	var builder UserNotificationBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(userstatus)
	return builder
}

//Username is a non-required field for UserNotification.
func (m UserNotification) Username() (field.Username, error) {
	var f field.Username
	err := m.Body.Get(&f)
	return f, err
}

//UserStatus is a required field for UserNotification.
func (m UserNotification) UserStatus() (field.UserStatus, error) {
	var f field.UserStatus
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for UserNotification.
func (m UserNotification) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for UserNotification.
func (m UserNotification) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for UserNotification.
func (m UserNotification) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
