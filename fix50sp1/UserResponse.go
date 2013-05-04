package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type UserResponse struct {
	quickfixgo.Message
}

func (m *UserResponse) UserRequestID() (*field.UserRequestID, error) {
	f := new(field.UserRequestID)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserResponse) Username() (*field.Username, error) {
	f := new(field.Username)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserResponse) UserStatus() (*field.UserStatus, error) {
	f := new(field.UserStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *UserResponse) UserStatusText() (*field.UserStatusText, error) {
	f := new(field.UserStatusText)
	err := m.Body.Get(f)
	return f, err
}
