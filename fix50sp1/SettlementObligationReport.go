package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type SettlementObligationReport struct {
	quickfixgo.Message
}

func (m *SettlementObligationReport) SettlObligMsgID() (*field.SettlObligMsgID, error) {
	f := new(field.SettlObligMsgID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementObligationReport) SettlObligMode() (*field.SettlObligMode, error) {
	f := new(field.SettlObligMode)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementObligationReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementObligationReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementObligationReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementObligationReport) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementObligationReport) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementObligationReport) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementObligationReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementObligationReport) SettlementCycleNo() (*field.SettlementCycleNo, error) {
	f := new(field.SettlementCycleNo)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementObligationReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementObligationReport) NoSettlOblig() (*field.NoSettlOblig, error) {
	f := new(field.NoSettlOblig)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementObligationReport) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}
