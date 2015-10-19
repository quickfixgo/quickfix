//Package userrequest msg type = BE.
package userrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a UserRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//UserRequestID is a required field for UserRequest.
func (m Message) UserRequestID() (*field.UserRequestIDField, quickfix.MessageRejectError) {
	f := &field.UserRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUserRequestID reads a UserRequestID from UserRequest.
func (m Message) GetUserRequestID(f *field.UserRequestIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UserRequestType is a required field for UserRequest.
func (m Message) UserRequestType() (*field.UserRequestTypeField, quickfix.MessageRejectError) {
	f := &field.UserRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUserRequestType reads a UserRequestType from UserRequest.
func (m Message) GetUserRequestType(f *field.UserRequestTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Username is a required field for UserRequest.
func (m Message) Username() (*field.UsernameField, quickfix.MessageRejectError) {
	f := &field.UsernameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUsername reads a Username from UserRequest.
func (m Message) GetUsername(f *field.UsernameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Password is a non-required field for UserRequest.
func (m Message) Password() (*field.PasswordField, quickfix.MessageRejectError) {
	f := &field.PasswordField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPassword reads a Password from UserRequest.
func (m Message) GetPassword(f *field.PasswordField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NewPassword is a non-required field for UserRequest.
func (m Message) NewPassword() (*field.NewPasswordField, quickfix.MessageRejectError) {
	f := &field.NewPasswordField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNewPassword reads a NewPassword from UserRequest.
func (m Message) GetNewPassword(f *field.NewPasswordField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for UserRequest.
func (m Message) RawDataLength() (*field.RawDataLengthField, quickfix.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from UserRequest.
func (m Message) GetRawDataLength(f *field.RawDataLengthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for UserRequest.
func (m Message) RawData() (*field.RawDataField, quickfix.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from UserRequest.
func (m Message) GetRawData(f *field.RawDataField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncryptedPasswordMethod is a non-required field for UserRequest.
func (m Message) EncryptedPasswordMethod() (*field.EncryptedPasswordMethodField, quickfix.MessageRejectError) {
	f := &field.EncryptedPasswordMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptedPasswordMethod reads a EncryptedPasswordMethod from UserRequest.
func (m Message) GetEncryptedPasswordMethod(f *field.EncryptedPasswordMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncryptedPasswordLen is a non-required field for UserRequest.
func (m Message) EncryptedPasswordLen() (*field.EncryptedPasswordLenField, quickfix.MessageRejectError) {
	f := &field.EncryptedPasswordLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptedPasswordLen reads a EncryptedPasswordLen from UserRequest.
func (m Message) GetEncryptedPasswordLen(f *field.EncryptedPasswordLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncryptedPassword is a non-required field for UserRequest.
func (m Message) EncryptedPassword() (*field.EncryptedPasswordField, quickfix.MessageRejectError) {
	f := &field.EncryptedPasswordField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptedPassword reads a EncryptedPassword from UserRequest.
func (m Message) GetEncryptedPassword(f *field.EncryptedPasswordField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncryptedNewPasswordLen is a non-required field for UserRequest.
func (m Message) EncryptedNewPasswordLen() (*field.EncryptedNewPasswordLenField, quickfix.MessageRejectError) {
	f := &field.EncryptedNewPasswordLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptedNewPasswordLen reads a EncryptedNewPasswordLen from UserRequest.
func (m Message) GetEncryptedNewPasswordLen(f *field.EncryptedNewPasswordLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncryptedNewPassword is a non-required field for UserRequest.
func (m Message) EncryptedNewPassword() (*field.EncryptedNewPasswordField, quickfix.MessageRejectError) {
	f := &field.EncryptedNewPasswordField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptedNewPassword reads a EncryptedNewPassword from UserRequest.
func (m Message) GetEncryptedNewPassword(f *field.EncryptedNewPasswordField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for UserRequest.
func New(
	userrequestid *field.UserRequestIDField,
	userrequesttype *field.UserRequestTypeField,
	username *field.UsernameField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("BE"))
	builder.Body.Set(userrequestid)
	builder.Body.Set(userrequesttype)
	builder.Body.Set(username)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "BE", r
}
