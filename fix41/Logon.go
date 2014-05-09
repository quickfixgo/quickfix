package fix41

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
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX41))
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
