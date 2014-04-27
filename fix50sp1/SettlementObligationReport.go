package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	settlobligmsgid field.SettlObligMsgID,
	settlobligmode field.SettlObligMode) SettlementObligationReportBuilder {
	var builder SettlementObligationReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildMsgType("BQ"))
	builder.Body.Set(settlobligmsgid)
	builder.Body.Set(settlobligmode)
	return builder
}

//ClearingBusinessDate is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from SettlementObligationReport.
func (m SettlementObligationReport) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlementCycleNo is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) SettlementCycleNo() (*field.SettlementCycleNo, errors.MessageRejectError) {
	f := new(field.SettlementCycleNo)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlementCycleNo reads a SettlementCycleNo from SettlementObligationReport.
func (m SettlementObligationReport) GetSettlementCycleNo(f *field.SettlementCycleNo) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlObligMsgID is a required field for SettlementObligationReport.
func (m SettlementObligationReport) SettlObligMsgID() (*field.SettlObligMsgID, errors.MessageRejectError) {
	f := new(field.SettlObligMsgID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlObligMsgID reads a SettlObligMsgID from SettlementObligationReport.
func (m SettlementObligationReport) GetSettlObligMsgID(f *field.SettlObligMsgID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlObligMode is a required field for SettlementObligationReport.
func (m SettlementObligationReport) SettlObligMode() (*field.SettlObligMode, errors.MessageRejectError) {
	f := new(field.SettlObligMode)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlObligMode reads a SettlObligMode from SettlementObligationReport.
func (m SettlementObligationReport) GetSettlObligMode(f *field.SettlObligMode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SettlementObligationReport.
func (m SettlementObligationReport) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SettlementObligationReport.
func (m SettlementObligationReport) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SettlementObligationReport.
func (m SettlementObligationReport) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementObligationReport.
func (m SettlementObligationReport) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSettlOblig is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) NoSettlOblig() (*field.NoSettlOblig, errors.MessageRejectError) {
	f := new(field.NoSettlOblig)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSettlOblig reads a NoSettlOblig from SettlementObligationReport.
func (m SettlementObligationReport) GetNoSettlOblig(f *field.NoSettlOblig) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from SettlementObligationReport.
func (m SettlementObligationReport) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from SettlementObligationReport.
func (m SettlementObligationReport) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from SettlementObligationReport.
func (m SettlementObligationReport) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from SettlementObligationReport.
func (m SettlementObligationReport) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}
