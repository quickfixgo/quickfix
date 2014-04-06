package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type UserRequest struct {
	quickfix.Message
}

func (m *UserRequest) UserRequestType() (*field.UserRequestType, error) {
	f := new(field.UserRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserRequest) NewPassword() (*field.NewPassword, error) {
	f := new(field.NewPassword)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserRequest) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserRequest) RawData() (*field.RawData, error) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserRequest) EncryptedPasswordMethod() (*field.EncryptedPasswordMethod, error) {
	f := new(field.EncryptedPasswordMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserRequest) EncryptedPasswordLen() (*field.EncryptedPasswordLen, error) {
	f := new(field.EncryptedPasswordLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserRequest) EncryptedNewPassword() (*field.EncryptedNewPassword, error) {
	f := new(field.EncryptedNewPassword)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserRequest) UserRequestID() (*field.UserRequestID, error) {
	f := new(field.UserRequestID)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserRequest) Username() (*field.Username, error) {
	f := new(field.Username)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserRequest) Password() (*field.Password, error) {
	f := new(field.Password)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserRequest) EncryptedPassword() (*field.EncryptedPassword, error) {
	f := new(field.EncryptedPassword)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserRequest) EncryptedNewPasswordLen() (*field.EncryptedNewPasswordLen, error) {
	f := new(field.EncryptedNewPasswordLen)
	err := m.Body.Get(f)
	return f, err
}
