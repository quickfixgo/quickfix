package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Logout msg type = 5.
type Logout struct {
	message.Message
}

//LogoutBuilder builds Logout messages.
type LogoutBuilder struct {
	message.MessageBuilder
}

//NewLogoutBuilder returns an initialized LogoutBuilder with specified required fields.
func NewLogoutBuilder() *LogoutBuilder {
	builder := new(LogoutBuilder)
	return builder
}

//Text is a non-required field for Logout.
func (m *Logout) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for Logout.
func (m *Logout) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for Logout.
func (m *Logout) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
