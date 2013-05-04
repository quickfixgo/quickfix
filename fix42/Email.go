package fix42

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type Email struct {
	quickfixgo.Message
}

func (m *Email) EmailThreadID() (*field.EmailThreadID, error) {
	f := new(field.EmailThreadID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) EmailType() (*field.EmailType, error) {
	f := new(field.EmailType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) OrigTime() (*field.OrigTime, error) {
	f := new(field.OrigTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) Subject() (*field.Subject, error) {
	f := new(field.Subject)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) EncodedSubjectLen() (*field.EncodedSubjectLen, error) {
	f := new(field.EncodedSubjectLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) EncodedSubject() (*field.EncodedSubject, error) {
	f := new(field.EncodedSubject)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) RawData() (*field.RawData, error) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}
