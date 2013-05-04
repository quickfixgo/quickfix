package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type PositionMaintenanceReport struct {
	quickfixgo.Message
}

func (m *PositionMaintenanceReport) PosMaintRptID() (*field.PosMaintRptID, error) {
	f := new(field.PosMaintRptID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) PosTransType() (*field.PosTransType, error) {
	f := new(field.PosTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) PosReqID() (*field.PosReqID, error) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) PosMaintAction() (*field.PosMaintAction, error) {
	f := new(field.PosMaintAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) OrigPosReqRefID() (*field.OrigPosReqRefID, error) {
	f := new(field.OrigPosReqRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) PosMaintStatus() (*field.PosMaintStatus, error) {
	f := new(field.PosMaintStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) PosMaintResult() (*field.PosMaintResult, error) {
	f := new(field.PosMaintResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) AdjustmentType() (*field.AdjustmentType, error) {
	f := new(field.AdjustmentType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) ThresholdAmount() (*field.ThresholdAmount, error) {
	f := new(field.ThresholdAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
