package fixt11

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

//CreateLogonBuilder returns an initialized LogonBuilder with specified required fields.
func CreateLogonBuilder(
	encryptmethod field.EncryptMethod,
	heartbtint field.HeartBtInt,
	defaultapplverid field.DefaultApplVerID) LogonBuilder {
	var builder LogonBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(encryptmethod)
	builder.Body.Set(heartbtint)
	builder.Body.Set(defaultapplverid)
	return builder
}

//EncryptMethod is a required field for Logon.
func (m Logon) EncryptMethod() (field.EncryptMethod, error) {
	var f field.EncryptMethod
	err := m.Body.Get(&f)
	return f, err
}

//HeartBtInt is a required field for Logon.
func (m Logon) HeartBtInt() (field.HeartBtInt, error) {
	var f field.HeartBtInt
	err := m.Body.Get(&f)
	return f, err
}

//RawDataLength is a non-required field for Logon.
func (m Logon) RawDataLength() (field.RawDataLength, error) {
	var f field.RawDataLength
	err := m.Body.Get(&f)
	return f, err
}

//RawData is a non-required field for Logon.
func (m Logon) RawData() (field.RawData, error) {
	var f field.RawData
	err := m.Body.Get(&f)
	return f, err
}

//ResetSeqNumFlag is a non-required field for Logon.
func (m Logon) ResetSeqNumFlag() (field.ResetSeqNumFlag, error) {
	var f field.ResetSeqNumFlag
	err := m.Body.Get(&f)
	return f, err
}

//NextExpectedMsgSeqNum is a non-required field for Logon.
func (m Logon) NextExpectedMsgSeqNum() (field.NextExpectedMsgSeqNum, error) {
	var f field.NextExpectedMsgSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//MaxMessageSize is a non-required field for Logon.
func (m Logon) MaxMessageSize() (field.MaxMessageSize, error) {
	var f field.MaxMessageSize
	err := m.Body.Get(&f)
	return f, err
}

//TestMessageIndicator is a non-required field for Logon.
func (m Logon) TestMessageIndicator() (field.TestMessageIndicator, error) {
	var f field.TestMessageIndicator
	err := m.Body.Get(&f)
	return f, err
}

//Username is a non-required field for Logon.
func (m Logon) Username() (field.Username, error) {
	var f field.Username
	err := m.Body.Get(&f)
	return f, err
}

//Password is a non-required field for Logon.
func (m Logon) Password() (field.Password, error) {
	var f field.Password
	err := m.Body.Get(&f)
	return f, err
}

//DefaultApplVerID is a required field for Logon.
func (m Logon) DefaultApplVerID() (field.DefaultApplVerID, error) {
	var f field.DefaultApplVerID
	err := m.Body.Get(&f)
	return f, err
}

//NoMsgTypes is a non-required field for Logon.
func (m Logon) NoMsgTypes() (field.NoMsgTypes, error) {
	var f field.NoMsgTypes
	err := m.Body.Get(&f)
	return f, err
}
