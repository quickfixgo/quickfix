package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//SettlementObligationReport msg type = BQ.
type SettlementObligationReport struct {
	message.Message
}

//SettlementObligationReportBuilder builds SettlementObligationReport messages.
type SettlementObligationReportBuilder struct {
	message.MessageBuilder
}

//CreateSettlementObligationReportBuilder returns an initialized SettlementObligationReportBuilder with specified required fields.
func CreateSettlementObligationReportBuilder(
	settlobligmsgid *field.SettlObligMsgIDField,
	settlobligmode *field.SettlObligModeField) SettlementObligationReportBuilder {
	var builder SettlementObligationReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("BQ"))
	builder.Body.Set(settlobligmsgid)
	builder.Body.Set(settlobligmode)
	return builder
}

//ClearingBusinessDate is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from SettlementObligationReport.
func (m SettlementObligationReport) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlementCycleNo is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) SettlementCycleNo() (*field.SettlementCycleNoField, errors.MessageRejectError) {
	f := &field.SettlementCycleNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlementCycleNo reads a SettlementCycleNo from SettlementObligationReport.
func (m SettlementObligationReport) GetSettlementCycleNo(f *field.SettlementCycleNoField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlObligMsgID is a required field for SettlementObligationReport.
func (m SettlementObligationReport) SettlObligMsgID() (*field.SettlObligMsgIDField, errors.MessageRejectError) {
	f := &field.SettlObligMsgIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlObligMsgID reads a SettlObligMsgID from SettlementObligationReport.
func (m SettlementObligationReport) GetSettlObligMsgID(f *field.SettlObligMsgIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlObligMode is a required field for SettlementObligationReport.
func (m SettlementObligationReport) SettlObligMode() (*field.SettlObligModeField, errors.MessageRejectError) {
	f := &field.SettlObligModeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlObligMode reads a SettlObligMode from SettlementObligationReport.
func (m SettlementObligationReport) GetSettlObligMode(f *field.SettlObligModeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SettlementObligationReport.
func (m SettlementObligationReport) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SettlementObligationReport.
func (m SettlementObligationReport) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SettlementObligationReport.
func (m SettlementObligationReport) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementObligationReport.
func (m SettlementObligationReport) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSettlOblig is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) NoSettlOblig() (*field.NoSettlObligField, errors.MessageRejectError) {
	f := &field.NoSettlObligField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSettlOblig reads a NoSettlOblig from SettlementObligationReport.
func (m SettlementObligationReport) GetNoSettlOblig(f *field.NoSettlObligField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from SettlementObligationReport.
func (m SettlementObligationReport) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from SettlementObligationReport.
func (m SettlementObligationReport) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from SettlementObligationReport.
func (m SettlementObligationReport) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from SettlementObligationReport.
func (m SettlementObligationReport) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}
