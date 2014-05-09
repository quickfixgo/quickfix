package fix50

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
	userrequestid *field.UserRequestIDField,
	userrequesttype *field.UserRequestTypeField,
	username *field.UsernameField) UserRequestBuilder {
	var builder UserRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.NewMsgType("BE"))
	builder.Body.Set(userrequestid)
	builder.Body.Set(userrequesttype)
	builder.Body.Set(username)
	return builder
}

//UserRequestID is a required field for UserRequest.
func (m UserRequest) UserRequestID() (*field.UserRequestIDField, errors.MessageRejectError) {
	f := &field.UserRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUserRequestID reads a UserRequestID from UserRequest.
func (m UserRequest) GetUserRequestID(f *field.UserRequestIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UserRequestType is a required field for UserRequest.
func (m UserRequest) UserRequestType() (*field.UserRequestTypeField, errors.MessageRejectError) {
	f := &field.UserRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUserRequestType reads a UserRequestType from UserRequest.
func (m UserRequest) GetUserRequestType(f *field.UserRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Username is a required field for UserRequest.
func (m UserRequest) Username() (*field.UsernameField, errors.MessageRejectError) {
	f := &field.UsernameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUsername reads a Username from UserRequest.
func (m UserRequest) GetUsername(f *field.UsernameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Password is a non-required field for UserRequest.
func (m UserRequest) Password() (*field.PasswordField, errors.MessageRejectError) {
	f := &field.PasswordField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPassword reads a Password from UserRequest.
func (m UserRequest) GetPassword(f *field.PasswordField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NewPassword is a non-required field for UserRequest.
func (m UserRequest) NewPassword() (*field.NewPasswordField, errors.MessageRejectError) {
	f := &field.NewPasswordField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNewPassword reads a NewPassword from UserRequest.
func (m UserRequest) GetNewPassword(f *field.NewPasswordField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for UserRequest.
func (m UserRequest) RawDataLength() (*field.RawDataLengthField, errors.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from UserRequest.
func (m UserRequest) GetRawDataLength(f *field.RawDataLengthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for UserRequest.
func (m UserRequest) RawData() (*field.RawDataField, errors.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from UserRequest.
func (m UserRequest) GetRawData(f *field.RawDataField) errors.MessageRejectError {
	return m.Body.Get(f)
}
