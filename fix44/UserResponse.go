package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type UserResponse struct {
	message.Message
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
