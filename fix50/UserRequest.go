package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type UserRequest struct {
	message.Message
}

func (m *UserRequest) UserRequestID() (*field.UserRequestID, error) {
	f := new(field.UserRequestID)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserRequest) UserRequestType() (*field.UserRequestType, error) {
	f := new(field.UserRequestType)
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
