package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header().Set(field.NewMsgType("5"))
	return builder
}

//Text is a non-required field for Logout.
func (m Logout) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Logout.
func (m Logout) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
