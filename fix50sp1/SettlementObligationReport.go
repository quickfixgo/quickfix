package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Body.Set(settlobligmsgid)
	builder.Body.Set(settlobligmode)
	return builder
}

//ClearingBusinessDate is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ClearingBusinessDate() (field.ClearingBusinessDate, errors.MessageRejectError) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlementCycleNo is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) SettlementCycleNo() (field.SettlementCycleNo, errors.MessageRejectError) {
	var f field.SettlementCycleNo
	err := m.Body.Get(&f)
	return f, err
}

//SettlObligMsgID is a required field for SettlementObligationReport.
func (m SettlementObligationReport) SettlObligMsgID() (field.SettlObligMsgID, errors.MessageRejectError) {
	var f field.SettlObligMsgID
	err := m.Body.Get(&f)
	return f, err
}

//SettlObligMode is a required field for SettlementObligationReport.
func (m SettlementObligationReport) SettlObligMode() (field.SettlObligMode, errors.MessageRejectError) {
	var f field.SettlObligMode
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoSettlOblig is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) NoSettlOblig() (field.NoSettlOblig, errors.MessageRejectError) {
	var f field.NoSettlOblig
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplID() (field.ApplID, errors.MessageRejectError) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplSeqNum() (field.ApplSeqNum, errors.MessageRejectError) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplLastSeqNum() (field.ApplLastSeqNum, errors.MessageRejectError) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplResendFlag() (field.ApplResendFlag, errors.MessageRejectError) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}
