package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	userstatus *field.UserStatusField) UserNotificationBuilder {
	var builder UserNotificationBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("CB"))
	builder.Body().Set(userstatus)
	return builder
}

//Username is a non-required field for UserNotification.
func (m UserNotification) Username() (*field.UsernameField, errors.MessageRejectError) {
	f := &field.UsernameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUsername reads a Username from UserNotification.
func (m UserNotification) GetUsername(f *field.UsernameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UserStatus is a required field for UserNotification.
func (m UserNotification) UserStatus() (*field.UserStatusField, errors.MessageRejectError) {
	f := &field.UserStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUserStatus reads a UserStatus from UserNotification.
func (m UserNotification) GetUserStatus(f *field.UserStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for UserNotification.
func (m UserNotification) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from UserNotification.
func (m UserNotification) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for UserNotification.
func (m UserNotification) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from UserNotification.
func (m UserNotification) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for UserNotification.
func (m UserNotification) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from UserNotification.
func (m UserNotification) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
