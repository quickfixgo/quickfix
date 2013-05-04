package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type PositionMaintenanceRequest struct {
	quickfixgo.Message
}

func (m *PositionMaintenanceRequest) PosReqID() (*field.PosReqID, error) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) PosTransType() (*field.PosTransType, error) {
	f := new(field.PosTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) PosMaintAction() (*field.PosMaintAction, error) {
	f := new(field.PosMaintAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) OrigPosReqRefID() (*field.OrigPosReqRefID, error) {
	f := new(field.OrigPosReqRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) PosMaintRptRefID() (*field.PosMaintRptRefID, error) {
	f := new(field.PosMaintRptRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) AdjustmentType() (*field.AdjustmentType, error) {
	f := new(field.AdjustmentType)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) ContraryInstructionIndicator() (*field.ContraryInstructionIndicator, error) {
	f := new(field.ContraryInstructionIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) PriorSpreadIndicator() (*field.PriorSpreadIndicator, error) {
	f := new(field.PriorSpreadIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) ThresholdAmount() (*field.ThresholdAmount, error) {
	f := new(field.ThresholdAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *PositionMaintenanceRequest) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
