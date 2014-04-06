package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type ApplicationMessageReport struct {
	quickfix.Message
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
func (m *ApplicationMessageReport) NoApplIDs() (*field.NoApplIDs, error) {
	f := new(field.NoApplIDs)
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
