//Package logon msg type = A.
package logon

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a Logon wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//EncryptMethod is a required field for Logon.
func (m Message) EncryptMethod() (*field.EncryptMethodField, quickfix.MessageRejectError) {
	f := &field.EncryptMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptMethod reads a EncryptMethod from Logon.
func (m Message) GetEncryptMethod(f *field.EncryptMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//HeartBtInt is a required field for Logon.
func (m Message) HeartBtInt() (*field.HeartBtIntField, quickfix.MessageRejectError) {
	f := &field.HeartBtIntField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHeartBtInt reads a HeartBtInt from Logon.
func (m Message) GetHeartBtInt(f *field.HeartBtIntField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for Logon.
func (m Message) RawDataLength() (*field.RawDataLengthField, quickfix.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from Logon.
func (m Message) GetRawDataLength(f *field.RawDataLengthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for Logon.
func (m Message) RawData() (*field.RawDataField, quickfix.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from Logon.
func (m Message) GetRawData(f *field.RawDataField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for Logon.
func New(
	encryptmethod *field.EncryptMethodField,
	heartbtint *field.HeartBtIntField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX40))
	builder.Header.Set(field.NewMsgType("A"))
	builder.Body.Set(encryptmethod)
	builder.Body.Set(heartbtint)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX40, "A", r
}
