package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type AssignmentReport struct {
	quickfixgo.Message
}

func (m *AssignmentReport) AsgnRptID() (*field.AsgnRptID, error) {
	f := new(field.AsgnRptID)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) TotNumAssignmentReports() (*field.TotNumAssignmentReports, error) {
	f := new(field.TotNumAssignmentReports)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) LastRptRequested() (*field.LastRptRequested, error) {
	f := new(field.LastRptRequested)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) ThresholdAmount() (*field.ThresholdAmount, error) {
	f := new(field.ThresholdAmount)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) SettlPrice() (*field.SettlPrice, error) {
	f := new(field.SettlPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) SettlPriceType() (*field.SettlPriceType, error) {
	f := new(field.SettlPriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) UnderlyingSettlPrice() (*field.UnderlyingSettlPrice, error) {
	f := new(field.UnderlyingSettlPrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) AssignmentMethod() (*field.AssignmentMethod, error) {
	f := new(field.AssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) AssignmentUnit() (*field.AssignmentUnit, error) {
	f := new(field.AssignmentUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) OpenInterest() (*field.OpenInterest, error) {
	f := new(field.OpenInterest)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) ExerciseMethod() (*field.ExerciseMethod, error) {
	f := new(field.ExerciseMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *AssignmentReport) PriorSettlPrice() (*field.PriorSettlPrice, error) {
	f := new(field.PriorSettlPrice)
	err := m.Body.Get(f)
	return f, err
}
