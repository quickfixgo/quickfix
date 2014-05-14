//Package userresponse msg type = BF.
package userresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a UserResponse wrapper for the generic Message type
type Message struct {
	message.Message
}

//UserRequestID is a required field for UserResponse.
func (m Message) UserRequestID() (*field.UserRequestIDField, errors.MessageRejectError) {
	f := &field.UserRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUserRequestID reads a UserRequestID from UserResponse.
func (m Message) GetUserRequestID(f *field.UserRequestIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Username is a required field for UserResponse.
func (m Message) Username() (*field.UsernameField, errors.MessageRejectError) {
	f := &field.UsernameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUsername reads a Username from UserResponse.
func (m Message) GetUsername(f *field.UsernameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UserStatus is a non-required field for UserResponse.
func (m Message) UserStatus() (*field.UserStatusField, errors.MessageRejectError) {
	f := &field.UserStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUserStatus reads a UserStatus from UserResponse.
func (m Message) GetUserStatus(f *field.UserStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UserStatusText is a non-required field for UserResponse.
func (m Message) UserStatusText() (*field.UserStatusTextField, errors.MessageRejectError) {
	f := &field.UserStatusTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUserStatusText reads a UserStatusText from UserResponse.
func (m Message) GetUserStatusText(f *field.UserStatusTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds UserResponse messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for UserResponse.
func Builder(
	userrequestid *field.UserRequestIDField,
	username *field.UsernameField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("BF"))
	builder.Body().Set(userrequestid)
	builder.Body().Set(username)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "BF", r
}
