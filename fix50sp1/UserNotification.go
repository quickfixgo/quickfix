package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Header.Set(field.BuildMsgType("CB"))
	builder.Body.Set(userstatus)
	return builder
}

//Username is a non-required field for UserNotification.
func (m UserNotification) Username() (*field.Username, errors.MessageRejectError) {
	f := new(field.Username)
	err := m.Body.Get(f)
	return f, err
}

//GetUsername reads a Username from UserNotification.
func (m UserNotification) GetUsername(f *field.Username) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UserStatus is a required field for UserNotification.
func (m UserNotification) UserStatus() (*field.UserStatus, errors.MessageRejectError) {
	f := new(field.UserStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetUserStatus reads a UserStatus from UserNotification.
func (m UserNotification) GetUserStatus(f *field.UserStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for UserNotification.
func (m UserNotification) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from UserNotification.
func (m UserNotification) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for UserNotification.
func (m UserNotification) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from UserNotification.
func (m UserNotification) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for UserNotification.
func (m UserNotification) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from UserNotification.
func (m UserNotification) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
