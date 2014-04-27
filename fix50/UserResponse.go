package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	userrequestid field.UserRequestID,
	username field.Username) UserResponseBuilder {
	var builder UserResponseBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildMsgType("BF"))
	builder.Body.Set(userrequestid)
	builder.Body.Set(username)
	return builder
}

//UserRequestID is a required field for UserResponse.
func (m UserResponse) UserRequestID() (*field.UserRequestID, errors.MessageRejectError) {
	f := new(field.UserRequestID)
	err := m.Body.Get(f)
	return f, err
}

//GetUserRequestID reads a UserRequestID from UserResponse.
func (m UserResponse) GetUserRequestID(f *field.UserRequestID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Username is a required field for UserResponse.
func (m UserResponse) Username() (*field.Username, errors.MessageRejectError) {
	f := new(field.Username)
	err := m.Body.Get(f)
	return f, err
}

//GetUsername reads a Username from UserResponse.
func (m UserResponse) GetUsername(f *field.Username) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UserStatus is a non-required field for UserResponse.
func (m UserResponse) UserStatus() (*field.UserStatus, errors.MessageRejectError) {
	f := new(field.UserStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetUserStatus reads a UserStatus from UserResponse.
func (m UserResponse) GetUserStatus(f *field.UserStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UserStatusText is a non-required field for UserResponse.
func (m UserResponse) UserStatusText() (*field.UserStatusText, errors.MessageRejectError) {
	f := new(field.UserStatusText)
	err := m.Body.Get(f)
	return f, err
}

//GetUserStatusText reads a UserStatusText from UserResponse.
func (m UserResponse) GetUserStatusText(f *field.UserStatusText) errors.MessageRejectError {
	return m.Body.Get(f)
}
