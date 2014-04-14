package fix44

import (
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

//NewLogonBuilder returns an initialized LogonBuilder with specified required fields.
func NewLogonBuilder(
	encryptmethod field.EncryptMethod,
	heartbtint field.HeartBtInt) *LogonBuilder {
	builder := new(LogonBuilder)
	builder.Body.Set(encryptmethod)
	builder.Body.Set(heartbtint)
	return builder
}

//EncryptMethod is a required field for Logon.
func (m *Logon) EncryptMethod() (*field.EncryptMethod, error) {
	f := new(field.EncryptMethod)
	err := m.Body.Get(f)
	return f, err
}

//HeartBtInt is a required field for Logon.
func (m *Logon) HeartBtInt() (*field.HeartBtInt, error) {
	f := new(field.HeartBtInt)
	err := m.Body.Get(f)
	return f, err
}

//RawDataLength is a non-required field for Logon.
func (m *Logon) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}

//RawData is a non-required field for Logon.
func (m *Logon) RawData() (*field.RawData, error) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}

//ResetSeqNumFlag is a non-required field for Logon.
func (m *Logon) ResetSeqNumFlag() (*field.ResetSeqNumFlag, error) {
	f := new(field.ResetSeqNumFlag)
	err := m.Body.Get(f)
	return f, err
}

//NextExpectedMsgSeqNum is a non-required field for Logon.
func (m *Logon) NextExpectedMsgSeqNum() (*field.NextExpectedMsgSeqNum, error) {
	f := new(field.NextExpectedMsgSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//MaxMessageSize is a non-required field for Logon.
func (m *Logon) MaxMessageSize() (*field.MaxMessageSize, error) {
	f := new(field.MaxMessageSize)
	err := m.Body.Get(f)
	return f, err
}

//NoMsgTypes is a non-required field for Logon.
func (m *Logon) NoMsgTypes() (*field.NoMsgTypes, error) {
	f := new(field.NoMsgTypes)
	err := m.Body.Get(f)
	return f, err
}

//TestMessageIndicator is a non-required field for Logon.
func (m *Logon) TestMessageIndicator() (*field.TestMessageIndicator, error) {
	f := new(field.TestMessageIndicator)
	err := m.Body.Get(f)
	return f, err
}

//Username is a non-required field for Logon.
func (m *Logon) Username() (*field.Username, error) {
	f := new(field.Username)
	err := m.Body.Get(f)
	return f, err
}

//Password is a non-required field for Logon.
func (m *Logon) Password() (*field.Password, error) {
	f := new(field.Password)
	err := m.Body.Get(f)
	return f, err
}
