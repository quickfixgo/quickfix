package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type NewOrderList struct {
	message.Message
}

func (m *NewOrderList) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) EncodedListExecInst() (*field.EncodedListExecInst, error) {
	f := new(field.EncodedListExecInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) AllowableOneSidednessCurr() (*field.AllowableOneSidednessCurr, error) {
	f := new(field.AllowableOneSidednessCurr)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) ContingencyType() (*field.ContingencyType, error) {
	f := new(field.ContingencyType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) BidType() (*field.BidType, error) {
	f := new(field.BidType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) ListExecInstType() (*field.ListExecInstType, error) {
	f := new(field.ListExecInstType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) EncodedListExecInstLen() (*field.EncodedListExecInstLen, error) {
	f := new(field.EncodedListExecInstLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) AllowableOneSidednessPct() (*field.AllowableOneSidednessPct, error) {
	f := new(field.AllowableOneSidednessPct)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) ClientBidID() (*field.ClientBidID, error) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) ProgRptReqs() (*field.ProgRptReqs, error) {
	f := new(field.ProgRptReqs)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) TotNoOrders() (*field.TotNoOrders, error) {
	f := new(field.TotNoOrders)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) NoRootPartyIDs() (*field.NoRootPartyIDs, error) {
	f := new(field.NoRootPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) BidID() (*field.BidID, error) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) ProgPeriodInterval() (*field.ProgPeriodInterval, error) {
	f := new(field.ProgPeriodInterval)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) ListExecInst() (*field.ListExecInst, error) {
	f := new(field.ListExecInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) AllowableOneSidednessValue() (*field.AllowableOneSidednessValue, error) {
	f := new(field.AllowableOneSidednessValue)
	err := m.Body.Get(f)
	return f, err
}
