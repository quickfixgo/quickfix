//Package logon msg type = A.
package logon

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a Logon wrapper for the generic Message type
type Message struct {
	message.Message
}

//EncryptMethod is a required field for Logon.
func (m Message) EncryptMethod() (*field.EncryptMethodField, errors.MessageRejectError) {
	f := &field.EncryptMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptMethod reads a EncryptMethod from Logon.
func (m Message) GetEncryptMethod(f *field.EncryptMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HeartBtInt is a required field for Logon.
func (m Message) HeartBtInt() (*field.HeartBtIntField, errors.MessageRejectError) {
	f := &field.HeartBtIntField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHeartBtInt reads a HeartBtInt from Logon.
func (m Message) GetHeartBtInt(f *field.HeartBtIntField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for Logon.
func (m Message) RawDataLength() (*field.RawDataLengthField, errors.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from Logon.
func (m Message) GetRawDataLength(f *field.RawDataLengthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for Logon.
func (m Message) RawData() (*field.RawDataField, errors.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from Logon.
func (m Message) GetRawData(f *field.RawDataField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResetSeqNumFlag is a non-required field for Logon.
func (m Message) ResetSeqNumFlag() (*field.ResetSeqNumFlagField, errors.MessageRejectError) {
	f := &field.ResetSeqNumFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetResetSeqNumFlag reads a ResetSeqNumFlag from Logon.
func (m Message) GetResetSeqNumFlag(f *field.ResetSeqNumFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxMessageSize is a non-required field for Logon.
func (m Message) MaxMessageSize() (*field.MaxMessageSizeField, errors.MessageRejectError) {
	f := &field.MaxMessageSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxMessageSize reads a MaxMessageSize from Logon.
func (m Message) GetMaxMessageSize(f *field.MaxMessageSizeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMsgTypes is a non-required field for Logon.
func (m Message) NoMsgTypes() (*field.NoMsgTypesField, errors.MessageRejectError) {
	f := &field.NoMsgTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMsgTypes reads a NoMsgTypes from Logon.
func (m Message) GetNoMsgTypes(f *field.NoMsgTypesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TestMessageIndicator is a non-required field for Logon.
func (m Message) TestMessageIndicator() (*field.TestMessageIndicatorField, errors.MessageRejectError) {
	f := &field.TestMessageIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTestMessageIndicator reads a TestMessageIndicator from Logon.
func (m Message) GetTestMessageIndicator(f *field.TestMessageIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Username is a non-required field for Logon.
func (m Message) Username() (*field.UsernameField, errors.MessageRejectError) {
	f := &field.UsernameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUsername reads a Username from Logon.
func (m Message) GetUsername(f *field.UsernameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Password is a non-required field for Logon.
func (m Message) Password() (*field.PasswordField, errors.MessageRejectError) {
	f := &field.PasswordField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPassword reads a Password from Logon.
func (m Message) GetPassword(f *field.PasswordField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds Logon messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for Logon.
func Builder(
	encryptmethod *field.EncryptMethodField,
	heartbtint *field.HeartBtIntField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header().Set(field.NewMsgType("A"))
	builder.Body().Set(encryptmethod)
	builder.Body().Set(heartbtint)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX43, "A", r
}
