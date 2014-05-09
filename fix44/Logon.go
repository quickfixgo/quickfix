package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Logon msg type = A.
type Logon struct {
	message.Message
}

//LogonBuilder builds Logon messages.
type LogonBuilder struct {
	message.MessageBuilder
}

//CreateLogonBuilder returns an initialized LogonBuilder with specified required fields.
func CreateLogonBuilder(
	encryptmethod *field.EncryptMethodField,
	heartbtint *field.HeartBtIntField) LogonBuilder {
	var builder LogonBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("A"))
	builder.Body.Set(encryptmethod)
	builder.Body.Set(heartbtint)
	return builder
}

//EncryptMethod is a required field for Logon.
func (m Logon) EncryptMethod() (*field.EncryptMethodField, errors.MessageRejectError) {
	f := &field.EncryptMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptMethod reads a EncryptMethod from Logon.
func (m Logon) GetEncryptMethod(f *field.EncryptMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HeartBtInt is a required field for Logon.
func (m Logon) HeartBtInt() (*field.HeartBtIntField, errors.MessageRejectError) {
	f := &field.HeartBtIntField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHeartBtInt reads a HeartBtInt from Logon.
func (m Logon) GetHeartBtInt(f *field.HeartBtIntField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for Logon.
func (m Logon) RawDataLength() (*field.RawDataLengthField, errors.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from Logon.
func (m Logon) GetRawDataLength(f *field.RawDataLengthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for Logon.
func (m Logon) RawData() (*field.RawDataField, errors.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from Logon.
func (m Logon) GetRawData(f *field.RawDataField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResetSeqNumFlag is a non-required field for Logon.
func (m Logon) ResetSeqNumFlag() (*field.ResetSeqNumFlagField, errors.MessageRejectError) {
	f := &field.ResetSeqNumFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetResetSeqNumFlag reads a ResetSeqNumFlag from Logon.
func (m Logon) GetResetSeqNumFlag(f *field.ResetSeqNumFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NextExpectedMsgSeqNum is a non-required field for Logon.
func (m Logon) NextExpectedMsgSeqNum() (*field.NextExpectedMsgSeqNumField, errors.MessageRejectError) {
	f := &field.NextExpectedMsgSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNextExpectedMsgSeqNum reads a NextExpectedMsgSeqNum from Logon.
func (m Logon) GetNextExpectedMsgSeqNum(f *field.NextExpectedMsgSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxMessageSize is a non-required field for Logon.
func (m Logon) MaxMessageSize() (*field.MaxMessageSizeField, errors.MessageRejectError) {
	f := &field.MaxMessageSizeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxMessageSize reads a MaxMessageSize from Logon.
func (m Logon) GetMaxMessageSize(f *field.MaxMessageSizeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMsgTypes is a non-required field for Logon.
func (m Logon) NoMsgTypes() (*field.NoMsgTypesField, errors.MessageRejectError) {
	f := &field.NoMsgTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMsgTypes reads a NoMsgTypes from Logon.
func (m Logon) GetNoMsgTypes(f *field.NoMsgTypesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TestMessageIndicator is a non-required field for Logon.
func (m Logon) TestMessageIndicator() (*field.TestMessageIndicatorField, errors.MessageRejectError) {
	f := &field.TestMessageIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTestMessageIndicator reads a TestMessageIndicator from Logon.
func (m Logon) GetTestMessageIndicator(f *field.TestMessageIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Username is a non-required field for Logon.
func (m Logon) Username() (*field.UsernameField, errors.MessageRejectError) {
	f := &field.UsernameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUsername reads a Username from Logon.
func (m Logon) GetUsername(f *field.UsernameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Password is a non-required field for Logon.
func (m Logon) Password() (*field.PasswordField, errors.MessageRejectError) {
	f := &field.PasswordField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPassword reads a Password from Logon.
func (m Logon) GetPassword(f *field.PasswordField) errors.MessageRejectError {
	return m.Body.Get(f)
}
