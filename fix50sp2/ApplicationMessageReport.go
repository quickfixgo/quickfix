package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	applreportid *field.ApplReportIDField,
	applreporttype *field.ApplReportTypeField) ApplicationMessageReportBuilder {
	var builder ApplicationMessageReportBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("BY"))
	builder.Body().Set(applreportid)
	builder.Body().Set(applreporttype)
	return builder
}

//ApplReportID is a required field for ApplicationMessageReport.
func (m ApplicationMessageReport) ApplReportID() (*field.ApplReportIDField, errors.MessageRejectError) {
	f := &field.ApplReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReportID reads a ApplReportID from ApplicationMessageReport.
func (m ApplicationMessageReport) GetApplReportID(f *field.ApplReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReportType is a required field for ApplicationMessageReport.
func (m ApplicationMessageReport) ApplReportType() (*field.ApplReportTypeField, errors.MessageRejectError) {
	f := &field.ApplReportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReportType reads a ApplReportType from ApplicationMessageReport.
func (m ApplicationMessageReport) GetApplReportType(f *field.ApplReportTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoApplIDs is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) NoApplIDs() (*field.NoApplIDsField, errors.MessageRejectError) {
	f := &field.NoApplIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoApplIDs reads a NoApplIDs from ApplicationMessageReport.
func (m ApplicationMessageReport) GetNoApplIDs(f *field.NoApplIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ApplicationMessageReport.
func (m ApplicationMessageReport) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ApplicationMessageReport.
func (m ApplicationMessageReport) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ApplicationMessageReport.
func (m ApplicationMessageReport) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReqID is a non-required field for ApplicationMessageReport.
func (m ApplicationMessageReport) ApplReqID() (*field.ApplReqIDField, errors.MessageRejectError) {
	f := &field.ApplReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqID reads a ApplReqID from ApplicationMessageReport.
func (m ApplicationMessageReport) GetApplReqID(f *field.ApplReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}
