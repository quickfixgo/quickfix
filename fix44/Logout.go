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
func (m Logout) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for Logout.
func (m Logout) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for Logout.
func (m Logout) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
