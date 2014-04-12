package fix40

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type Email struct {
	message.Message
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
func (m *Email) LinesOfText() (*field.LinesOfText, error) {
	f := new(field.LinesOfText)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) RawData() (*field.RawData, error) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) RelatdSym() (*field.RelatdSym, error) {
	f := new(field.RelatdSym)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *Email) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}
