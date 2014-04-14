package fix50sp1

import (
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

//NewUserRequestBuilder returns an initialized UserRequestBuilder with specified required fields.
func NewUserRequestBuilder(
	userrequestid field.UserRequestID,
	userrequesttype field.UserRequestType,
	username field.Username) *UserRequestBuilder {
	builder := new(UserRequestBuilder)
	builder.Body.Set(userrequestid)
	builder.Body.Set(userrequesttype)
	builder.Body.Set(username)
	return builder
}

//UserRequestID is a required field for UserRequest.
func (m *UserRequest) UserRequestID() (*field.UserRequestID, error) {
	f := new(field.UserRequestID)
	err := m.Body.Get(f)
	return f, err
}

//UserRequestType is a required field for UserRequest.
func (m *UserRequest) UserRequestType() (*field.UserRequestType, error) {
	f := new(field.UserRequestType)
	err := m.Body.Get(f)
	return f, err
}

//Username is a required field for UserRequest.
func (m *UserRequest) Username() (*field.Username, error) {
	f := new(field.Username)
	err := m.Body.Get(f)
	return f, err
}

//Password is a non-required field for UserRequest.
func (m *UserRequest) Password() (*field.Password, error) {
	f := new(field.Password)
	err := m.Body.Get(f)
	return f, err
}

//NewPassword is a non-required field for UserRequest.
func (m *UserRequest) NewPassword() (*field.NewPassword, error) {
	f := new(field.NewPassword)
	err := m.Body.Get(f)
	return f, err
}

//RawDataLength is a non-required field for UserRequest.
func (m *UserRequest) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}

//RawData is a non-required field for UserRequest.
func (m *UserRequest) RawData() (*field.RawData, error) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}

//EncryptedPasswordMethod is a non-required field for UserRequest.
func (m *UserRequest) EncryptedPasswordMethod() (*field.EncryptedPasswordMethod, error) {
	f := new(field.EncryptedPasswordMethod)
	err := m.Body.Get(f)
	return f, err
}

//EncryptedPasswordLen is a non-required field for UserRequest.
func (m *UserRequest) EncryptedPasswordLen() (*field.EncryptedPasswordLen, error) {
	f := new(field.EncryptedPasswordLen)
	err := m.Body.Get(f)
	return f, err
}

//EncryptedPassword is a non-required field for UserRequest.
func (m *UserRequest) EncryptedPassword() (*field.EncryptedPassword, error) {
	f := new(field.EncryptedPassword)
	err := m.Body.Get(f)
	return f, err
}

//EncryptedNewPasswordLen is a non-required field for UserRequest.
func (m *UserRequest) EncryptedNewPasswordLen() (*field.EncryptedNewPasswordLen, error) {
	f := new(field.EncryptedNewPasswordLen)
	err := m.Body.Get(f)
	return f, err
}

//EncryptedNewPassword is a non-required field for UserRequest.
func (m *UserRequest) EncryptedNewPassword() (*field.EncryptedNewPassword, error) {
	f := new(field.EncryptedNewPassword)
	err := m.Body.Get(f)
	return f, err
}
