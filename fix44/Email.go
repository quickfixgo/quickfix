package fix44

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type Email struct {
	quickfix.Message
}

func (m *Email) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
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
func (m *Email) NoRoutingIDs() (*field.NoRoutingIDs, error) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) NoLinesOfText() (*field.NoLinesOfText, error) {
	f := new(field.NoLinesOfText)
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
func (m *Email) EmailType() (*field.EmailType, error) {
	f := new(field.EmailType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) EncodedSubject() (*field.EncodedSubject, error) {
	f := new(field.EncodedSubject)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) EmailThreadID() (*field.EmailThreadID, error) {
	f := new(field.EmailThreadID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
