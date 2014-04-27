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
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.BuildMsgType("BE"))
	builder.Body.Set(userrequestid)
	builder.Body.Set(userrequesttype)
	builder.Body.Set(username)
	return builder
}

//UserRequestID is a required field for UserRequest.
func (m UserRequest) UserRequestID() (*field.UserRequestID, errors.MessageRejectError) {
	f := new(field.UserRequestID)
	err := m.Body.Get(f)
	return f, err
}

//GetUserRequestID reads a UserRequestID from UserRequest.
func (m UserRequest) GetUserRequestID(f *field.UserRequestID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UserRequestType is a required field for UserRequest.
func (m UserRequest) UserRequestType() (*field.UserRequestType, errors.MessageRejectError) {
	f := new(field.UserRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetUserRequestType reads a UserRequestType from UserRequest.
func (m UserRequest) GetUserRequestType(f *field.UserRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Username is a required field for UserRequest.
func (m UserRequest) Username() (*field.Username, errors.MessageRejectError) {
	f := new(field.Username)
	err := m.Body.Get(f)
	return f, err
}

//GetUsername reads a Username from UserRequest.
func (m UserRequest) GetUsername(f *field.Username) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Password is a non-required field for UserRequest.
func (m UserRequest) Password() (*field.Password, errors.MessageRejectError) {
	f := new(field.Password)
	err := m.Body.Get(f)
	return f, err
}

//GetPassword reads a Password from UserRequest.
func (m UserRequest) GetPassword(f *field.Password) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NewPassword is a non-required field for UserRequest.
func (m UserRequest) NewPassword() (*field.NewPassword, errors.MessageRejectError) {
	f := new(field.NewPassword)
	err := m.Body.Get(f)
	return f, err
}

//GetNewPassword reads a NewPassword from UserRequest.
func (m UserRequest) GetNewPassword(f *field.NewPassword) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for UserRequest.
func (m UserRequest) RawDataLength() (*field.RawDataLength, errors.MessageRejectError) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from UserRequest.
func (m UserRequest) GetRawDataLength(f *field.RawDataLength) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for UserRequest.
func (m UserRequest) RawData() (*field.RawData, errors.MessageRejectError) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from UserRequest.
func (m UserRequest) GetRawData(f *field.RawData) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncryptedPasswordMethod is a non-required field for UserRequest.
func (m UserRequest) EncryptedPasswordMethod() (*field.EncryptedPasswordMethod, errors.MessageRejectError) {
	f := new(field.EncryptedPasswordMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptedPasswordMethod reads a EncryptedPasswordMethod from UserRequest.
func (m UserRequest) GetEncryptedPasswordMethod(f *field.EncryptedPasswordMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncryptedPasswordLen is a non-required field for UserRequest.
func (m UserRequest) EncryptedPasswordLen() (*field.EncryptedPasswordLen, errors.MessageRejectError) {
	f := new(field.EncryptedPasswordLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptedPasswordLen reads a EncryptedPasswordLen from UserRequest.
func (m UserRequest) GetEncryptedPasswordLen(f *field.EncryptedPasswordLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncryptedPassword is a non-required field for UserRequest.
func (m UserRequest) EncryptedPassword() (*field.EncryptedPassword, errors.MessageRejectError) {
	f := new(field.EncryptedPassword)
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptedPassword reads a EncryptedPassword from UserRequest.
func (m UserRequest) GetEncryptedPassword(f *field.EncryptedPassword) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncryptedNewPasswordLen is a non-required field for UserRequest.
func (m UserRequest) EncryptedNewPasswordLen() (*field.EncryptedNewPasswordLen, errors.MessageRejectError) {
	f := new(field.EncryptedNewPasswordLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptedNewPasswordLen reads a EncryptedNewPasswordLen from UserRequest.
func (m UserRequest) GetEncryptedNewPasswordLen(f *field.EncryptedNewPasswordLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncryptedNewPassword is a non-required field for UserRequest.
func (m UserRequest) EncryptedNewPassword() (*field.EncryptedNewPassword, errors.MessageRejectError) {
	f := new(field.EncryptedNewPassword)
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptedNewPassword reads a EncryptedNewPassword from UserRequest.
func (m UserRequest) GetEncryptedNewPassword(f *field.EncryptedNewPassword) errors.MessageRejectError {
	return m.Body.Get(f)
}
