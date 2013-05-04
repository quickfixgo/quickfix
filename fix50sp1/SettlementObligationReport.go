package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type SettlementObligationReport struct {
	quickfixgo.Message
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
func (m *SettlementObligationReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
