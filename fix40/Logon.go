package fix40

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
	heartbtint field.HeartBtInt) LogonBuilder {
	var builder LogonBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(encryptmethod)
	builder.Body.Set(heartbtint)
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
