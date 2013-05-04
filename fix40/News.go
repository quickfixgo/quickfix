package fix40

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type News struct {
	quickfixgo.Message
}

func (m *News) OrigTime() (*field.OrigTime, error) {
	f := new(field.OrigTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) Urgency() (*field.Urgency, error) {
	f := new(field.Urgency)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) RelatdSym() (*field.RelatdSym, error) {
	f := new(field.RelatdSym)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) LinesOfText() (*field.LinesOfText, error) {
	f := new(field.LinesOfText)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) RawData() (*field.RawData, error) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}
