package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type ContraryIntentionReport struct {
	quickfixgo.Message
}

func (m *ContraryIntentionReport) ContIntRptID() (*field.ContIntRptID, error) {
	f := new(field.ContIntRptID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ContraryIntentionReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *ContraryIntentionReport) LateIndicator() (*field.LateIndicator, error) {
	f := new(field.LateIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *ContraryIntentionReport) InputSource() (*field.InputSource, error) {
	f := new(field.InputSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *ContraryIntentionReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ContraryIntentionReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *ContraryIntentionReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *ContraryIntentionReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
