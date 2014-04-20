package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
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

//CreateLogoutBuilder returns an initialized LogoutBuilder with specified required fields.
func CreateLogoutBuilder() LogoutBuilder {
	var builder LogoutBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//Text is a non-required field for Logout.
func (m Logout) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Logout.
func (m Logout) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for Logout.
func (m Logout) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from Logout.
func (m Logout) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for Logout.
func (m Logout) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from Logout.
func (m Logout) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
