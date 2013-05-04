package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type ApplicationMessageReport struct {
	quickfixgo.Message
}

func (m *ApplicationMessageReport) ApplReportID() (*field.ApplReportID, error) {
	f := new(field.ApplReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageReport) ApplReportType() (*field.ApplReportType, error) {
	f := new(field.ApplReportType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageReport) ApplReqID() (*field.ApplReqID, error) {
	f := new(field.ApplReqID)
	err := m.Body.Get(f)
	return f, err
}
