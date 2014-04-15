package fix50sp1

import (
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
func (m SettlementObligationReport) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlementCycleNo is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) SettlementCycleNo() (field.SettlementCycleNo, error) {
	var f field.SettlementCycleNo
	err := m.Body.Get(&f)
	return f, err
}

//SettlObligMsgID is a required field for SettlementObligationReport.
func (m SettlementObligationReport) SettlObligMsgID() (field.SettlObligMsgID, error) {
	var f field.SettlObligMsgID
	err := m.Body.Get(&f)
	return f, err
}

//SettlObligMode is a required field for SettlementObligationReport.
func (m SettlementObligationReport) SettlObligMode() (field.SettlObligMode, error) {
	var f field.SettlObligMode
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoSettlOblig is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) NoSettlOblig() (field.NoSettlOblig, error) {
	var f field.NoSettlOblig
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplID() (field.ApplID, error) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplSeqNum() (field.ApplSeqNum, error) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplLastSeqNum() (field.ApplLastSeqNum, error) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for SettlementObligationReport.
func (m SettlementObligationReport) ApplResendFlag() (field.ApplResendFlag, error) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}
