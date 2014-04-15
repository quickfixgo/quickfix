package fix50sp2

import (
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
	builder.Body.Set(userrequestid)
	builder.Body.Set(username)
	return builder
}

//UserRequestID is a required field for UserResponse.
func (m UserResponse) UserRequestID() (field.UserRequestID, error) {
	var f field.UserRequestID
	err := m.Body.Get(&f)
	return f, err
}

//Username is a required field for UserResponse.
func (m UserResponse) Username() (field.Username, error) {
	var f field.Username
	err := m.Body.Get(&f)
	return f, err
}

//UserStatus is a non-required field for UserResponse.
func (m UserResponse) UserStatus() (field.UserStatus, error) {
	var f field.UserStatus
	err := m.Body.Get(&f)
	return f, err
}

//UserStatusText is a non-required field for UserResponse.
func (m UserResponse) UserStatusText() (field.UserStatusText, error) {
	var f field.UserStatusText
	err := m.Body.Get(&f)
	return f, err
}
