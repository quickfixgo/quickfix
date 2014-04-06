package fixt11

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type Logon struct {
	quickfix.Message
}

func (m *Logon) HeartBtInt() (*field.HeartBtInt, error) {
	f := new(field.HeartBtInt)
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
func (m *Logon) NextExpectedMsgSeqNum() (*field.NextExpectedMsgSeqNum, error) {
	f := new(field.NextExpectedMsgSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) MaxMessageSize() (*field.MaxMessageSize, error) {
	f := new(field.MaxMessageSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) Password() (*field.Password, error) {
	f := new(field.Password)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) DefaultApplVerID() (*field.DefaultApplVerID, error) {
	f := new(field.DefaultApplVerID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) EncryptMethod() (*field.EncryptMethod, error) {
	f := new(field.EncryptMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) TestMessageIndicator() (*field.TestMessageIndicator, error) {
	f := new(field.TestMessageIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) Username() (*field.Username, error) {
	f := new(field.Username)
	err := m.Body.Get(f)
	return f, err
}
func (m *Logon) NoMsgTypes() (*field.NoMsgTypes, error) {
	f := new(field.NoMsgTypes)
	err := m.Body.Get(f)
	return f, err
}
