package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//UserRequest msg type = BE.
type UserRequest struct {
	message.Message
}

//UserRequestBuilder builds UserRequest messages.
type UserRequestBuilder struct {
	message.MessageBuilder
}

//CreateUserRequestBuilder returns an initialized UserRequestBuilder with specified required fields.
func CreateUserRequestBuilder(
	userrequestid field.UserRequestID,
	userrequesttype field.UserRequestType,
	username field.Username) UserRequestBuilder {
	var builder UserRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(userrequestid)
	builder.Body.Set(userrequesttype)
	builder.Body.Set(username)
	return builder
}

//UserRequestID is a required field for UserRequest.
func (m UserRequest) UserRequestID() (field.UserRequestID, errors.MessageRejectError) {
	var f field.UserRequestID
	err := m.Body.Get(&f)
	return f, err
}

//UserRequestType is a required field for UserRequest.
func (m UserRequest) UserRequestType() (field.UserRequestType, errors.MessageRejectError) {
	var f field.UserRequestType
	err := m.Body.Get(&f)
	return f, err
}

//Username is a required field for UserRequest.
func (m UserRequest) Username() (field.Username, errors.MessageRejectError) {
	var f field.Username
	err := m.Body.Get(&f)
	return f, err
}

//Password is a non-required field for UserRequest.
func (m UserRequest) Password() (field.Password, errors.MessageRejectError) {
	var f field.Password
	err := m.Body.Get(&f)
	return f, err
}

//NewPassword is a non-required field for UserRequest.
func (m UserRequest) NewPassword() (field.NewPassword, errors.MessageRejectError) {
	var f field.NewPassword
	err := m.Body.Get(&f)
	return f, err
}

//RawDataLength is a non-required field for UserRequest.
func (m UserRequest) RawDataLength() (field.RawDataLength, errors.MessageRejectError) {
	var f field.RawDataLength
	err := m.Body.Get(&f)
	return f, err
}

//RawData is a non-required field for UserRequest.
func (m UserRequest) RawData() (field.RawData, errors.MessageRejectError) {
	var f field.RawData
	err := m.Body.Get(&f)
	return f, err
}
