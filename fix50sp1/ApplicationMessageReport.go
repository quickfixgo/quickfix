package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ApplicationMessageReport msg type = BY.
type ApplicationMessageReport struct {
	message.Message
}

//ApplicationMessageReportBuilder builds ApplicationMessageReport messages.
type ApplicationMessageReportBuilder struct {
	message.MessageBuilder
}

//NewApplicationMessageReportBuilder returns an initialized ApplicationMessageReportBuilder with specified required fields.
func NewApplicationMessageReportBuilder(
	applreportid field.ApplReportID,
	applreporttype field.ApplReportType) *ApplicationMessageReportBuilder {
	builder := new(ApplicationMessageReportBuilder)
	builder.Body.Set(applreportid)
	builder.Body.Set(applreporttype)
	return builder
}

//ApplReportID is a required field for ApplicationMessageReport.
func (m *ApplicationMessageReport) ApplReportID() (*field.ApplReportID, error) {
	f := new(field.ApplReportID)
	err := m.Body.Get(f)
	return f, err
}

//ApplReportType is a required field for ApplicationMessageReport.
func (m *ApplicationMessageReport) ApplReportType() (*field.ApplReportType, error) {
	f := new(field.ApplReportType)
	err := m.Body.Get(f)
	return f, err
}

//NoApplIDs is a non-required field for ApplicationMessageReport.
func (m *ApplicationMessageReport) NoApplIDs() (*field.NoApplIDs, error) {
	f := new(field.NoApplIDs)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for ApplicationMessageReport.
func (m *ApplicationMessageReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for ApplicationMessageReport.
func (m *ApplicationMessageReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for ApplicationMessageReport.
func (m *ApplicationMessageReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
