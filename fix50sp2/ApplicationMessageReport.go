package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
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
func (m ApplicationMessageReport) ApplReportID() (*field.ApplReportID, errors.MessageRejectError) {
	f := new(field.ApplReportID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplReportID reads a ApplReportID from ApplicationMessageReport.
func (m ApplicationMessageReport) GetApplReportID(f *field.ApplReportID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReportType is a required field for ApplicationMessageReport.
func (m ApplicationMessageReport) ApplReportType() (*field.ApplReportType, errors.MessageRejectError) {
	f := new(field.ApplReportType)
	err := m.Body.Get(f)
	return f, err
}

//GetApplReportType reads a ApplReportType from ApplicationMessageReport.
func (m ApplicationMessageReport) GetApplReportType(f *field.ApplReportType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoApplIDs is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) NoApplIDs() (*field.NoApplIDs, errors.MessageRejectError) {
	f := new(field.NoApplIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoApplIDs reads a NoApplIDs from ApplicationMessageReport.
func (m ApplicationMessageReport) GetNoApplIDs(f *field.NoApplIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ApplicationMessageReport.
func (m ApplicationMessageReport) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ApplicationMessageReport.
func (m ApplicationMessageReport) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ApplicationMessageReport.
func (m ApplicationMessageReport) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReqID is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) ApplReqID() (*field.ApplReqID, errors.MessageRejectError) {
	f := new(field.ApplReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqID reads a ApplReqID from ApplicationMessageReport.
func (m ApplicationMessageReport) GetApplReqID(f *field.ApplReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}
