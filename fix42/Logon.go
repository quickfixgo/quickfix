package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type Logon struct {
	message.Message
}

func (m *Logon) EncryptMethod() (*field.EncryptMethod, error) {
	f := new(field.EncryptMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) HeartBtInt() (*field.HeartBtInt, error) {
	f := new(field.HeartBtInt)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) RawData() (*field.RawData, error) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) ResetSeqNumFlag() (*field.ResetSeqNumFlag, error) {
	f := new(field.ResetSeqNumFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) MaxMessageSize() (*field.MaxMessageSize, error) {
	f := new(field.MaxMessageSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) NoMsgTypes() (*field.NoMsgTypes, error) {
	f := new(field.NoMsgTypes)
	err := m.Body.Get(f)
	return f, err
}
