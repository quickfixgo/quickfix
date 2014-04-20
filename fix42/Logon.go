package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
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
	encryptmethod field.EncryptMethod,
	heartbtint field.HeartBtInt) LogonBuilder {
	var builder LogonBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(encryptmethod)
	builder.Body.Set(heartbtint)
	return builder
}

//EncryptMethod is a required field for Logon.
func (m Logon) EncryptMethod() (*field.EncryptMethod, errors.MessageRejectError) {
	f := new(field.EncryptMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetEncryptMethod reads a EncryptMethod from Logon.
func (m Logon) GetEncryptMethod(f *field.EncryptMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HeartBtInt is a required field for Logon.
func (m Logon) HeartBtInt() (*field.HeartBtInt, errors.MessageRejectError) {
	f := new(field.HeartBtInt)
	err := m.Body.Get(f)
	return f, err
}

//GetHeartBtInt reads a HeartBtInt from Logon.
func (m Logon) GetHeartBtInt(f *field.HeartBtInt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for Logon.
func (m Logon) RawDataLength() (*field.RawDataLength, errors.MessageRejectError) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from Logon.
func (m Logon) GetRawDataLength(f *field.RawDataLength) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for Logon.
func (m Logon) RawData() (*field.RawData, errors.MessageRejectError) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from Logon.
func (m Logon) GetRawData(f *field.RawData) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResetSeqNumFlag is a non-required field for Logon.
func (m Logon) ResetSeqNumFlag() (*field.ResetSeqNumFlag, errors.MessageRejectError) {
	f := new(field.ResetSeqNumFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetResetSeqNumFlag reads a ResetSeqNumFlag from Logon.
func (m Logon) GetResetSeqNumFlag(f *field.ResetSeqNumFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxMessageSize is a non-required field for Logon.
func (m Logon) MaxMessageSize() (*field.MaxMessageSize, errors.MessageRejectError) {
	f := new(field.MaxMessageSize)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxMessageSize reads a MaxMessageSize from Logon.
func (m Logon) GetMaxMessageSize(f *field.MaxMessageSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMsgTypes is a non-required field for Logon.
func (m Logon) NoMsgTypes() (*field.NoMsgTypes, errors.MessageRejectError) {
	f := new(field.NoMsgTypes)
	err := m.Body.Get(f)
	return f, err
}

//GetNoMsgTypes reads a NoMsgTypes from Logon.
func (m Logon) GetNoMsgTypes(f *field.NoMsgTypes) errors.MessageRejectError {
	return m.Body.Get(f)
}
