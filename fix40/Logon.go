package fix40

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type Logon struct {
	quickfix.Message
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
