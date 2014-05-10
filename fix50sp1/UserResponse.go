package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//UserResponse msg type = BF.
type UserResponse struct {
	message.Message
}

//UserResponseBuilder builds UserResponse messages.
type UserResponseBuilder struct {
	message.MessageBuilder
}

//CreateUserResponseBuilder returns an initialized UserResponseBuilder with specified required fields.
func CreateUserResponseBuilder(
	userrequestid *field.UserRequestIDField,
	username *field.UsernameField) UserResponseBuilder {
	var builder UserResponseBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("BF"))
	builder.Body.Set(userrequestid)
	builder.Body.Set(username)
	return builder
}

//UserRequestID is a required field for UserResponse.
func (m UserResponse) UserRequestID() (*field.UserRequestIDField, errors.MessageRejectError) {
	f := &field.UserRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUserRequestID reads a UserRequestID from UserResponse.
func (m UserResponse) GetUserRequestID(f *field.UserRequestIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Username is a required field for UserResponse.
func (m UserResponse) Username() (*field.UsernameField, errors.MessageRejectError) {
	f := &field.UsernameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUsername reads a Username from UserResponse.
func (m UserResponse) GetUsername(f *field.UsernameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UserStatus is a non-required field for UserResponse.
func (m UserResponse) UserStatus() (*field.UserStatusField, errors.MessageRejectError) {
	f := &field.UserStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUserStatus reads a UserStatus from UserResponse.
func (m UserResponse) GetUserStatus(f *field.UserStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UserStatusText is a non-required field for UserResponse.
func (m UserResponse) UserStatusText() (*field.UserStatusTextField, errors.MessageRejectError) {
	f := &field.UserStatusTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUserStatusText reads a UserStatusText from UserResponse.
func (m UserResponse) GetUserStatusText(f *field.UserStatusTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
