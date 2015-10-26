//Package usernotification msg type = CB.
package usernotification

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a UserNotification wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//Username is a non-required field for UserNotification.
func (m Message) Username() (*field.UsernameField, quickfix.MessageRejectError) {
	f := &field.UsernameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUsername reads a Username from UserNotification.
func (m Message) GetUsername(f *field.UsernameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UserStatus is a required field for UserNotification.
func (m Message) UserStatus() (*field.UserStatusField, quickfix.MessageRejectError) {
	f := &field.UserStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUserStatus reads a UserStatus from UserNotification.
func (m Message) GetUserStatus(f *field.UserStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for UserNotification.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from UserNotification.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for UserNotification.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from UserNotification.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for UserNotification.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from UserNotification.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for UserNotification.
func New(
	userstatus *field.UserStatusField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("CB"))
	builder.Body.Set(userstatus)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "CB", r
}
