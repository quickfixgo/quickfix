package fix50sp2

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

//NewSettlementObligationReportBuilder returns an initialized SettlementObligationReportBuilder with specified required fields.
func NewSettlementObligationReportBuilder(
	settlobligmsgid field.SettlObligMsgID,
	settlobligmode field.SettlObligMode) *SettlementObligationReportBuilder {
	builder := new(SettlementObligationReportBuilder)
	builder.Body.Set(settlobligmsgid)
	builder.Body.Set(settlobligmode)
	return builder
}

//ClearingBusinessDate is a non-required field for SettlementObligationReport.
func (m *SettlementObligationReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//SettlementCycleNo is a non-required field for SettlementObligationReport.
func (m *SettlementObligationReport) SettlementCycleNo() (*field.SettlementCycleNo, error) {
	f := new(field.SettlementCycleNo)
	err := m.Body.Get(f)
	return f, err
}

//SettlObligMsgID is a required field for SettlementObligationReport.
func (m *SettlementObligationReport) SettlObligMsgID() (*field.SettlObligMsgID, error) {
	f := new(field.SettlObligMsgID)
	err := m.Body.Get(f)
	return f, err
}

//SettlObligMode is a required field for SettlementObligationReport.
func (m *SettlementObligationReport) SettlObligMode() (*field.SettlObligMode, error) {
	f := new(field.SettlObligMode)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for SettlementObligationReport.
func (m *SettlementObligationReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for SettlementObligationReport.
func (m *SettlementObligationReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for SettlementObligationReport.
func (m *SettlementObligationReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for SettlementObligationReport.
func (m *SettlementObligationReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//NoSettlOblig is a non-required field for SettlementObligationReport.
func (m *SettlementObligationReport) NoSettlOblig() (*field.NoSettlOblig, error) {
	f := new(field.NoSettlOblig)
	err := m.Body.Get(f)
	return f, err
}

//ApplID is a non-required field for SettlementObligationReport.
func (m *SettlementObligationReport) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//ApplSeqNum is a non-required field for SettlementObligationReport.
func (m *SettlementObligationReport) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplLastSeqNum is a non-required field for SettlementObligationReport.
func (m *SettlementObligationReport) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplResendFlag is a non-required field for SettlementObligationReport.
func (m *SettlementObligationReport) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
