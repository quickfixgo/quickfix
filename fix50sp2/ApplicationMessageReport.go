package fix50sp2

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

//CreateApplicationMessageReportBuilder returns an initialized ApplicationMessageReportBuilder with specified required fields.
func CreateApplicationMessageReportBuilder(
	applreportid field.ApplReportID,
	applreporttype field.ApplReportType) ApplicationMessageReportBuilder {
	var builder ApplicationMessageReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(applreportid)
	builder.Body.Set(applreporttype)
	return builder
}

//ApplReportID is a required field for ApplicationMessageReport.
func (m ApplicationMessageReport) ApplReportID() (field.ApplReportID, error) {
	var f field.ApplReportID
	err := m.Body.Get(&f)
	return f, err
}

//ApplReportType is a required field for ApplicationMessageReport.
func (m ApplicationMessageReport) ApplReportType() (field.ApplReportType, error) {
	var f field.ApplReportType
	err := m.Body.Get(&f)
	return f, err
}

//NoApplIDs is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) NoApplIDs() (field.NoApplIDs, error) {
	var f field.NoApplIDs
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//ApplReqID is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) ApplReqID() (field.ApplReqID, error) {
	var f field.ApplReqID
	err := m.Body.Get(&f)
	return f, err
}
