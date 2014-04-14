package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderList msg type = E.
type NewOrderList struct {
	message.Message
}

//NewOrderListBuilder builds NewOrderList messages.
type NewOrderListBuilder struct {
	message.MessageBuilder
}

//NewNewOrderListBuilder returns an initialized NewOrderListBuilder with specified required fields.
func NewNewOrderListBuilder(
	listid field.ListID,
	bidtype field.BidType,
	totnoorders field.TotNoOrders,
	noorders field.NoOrders) *NewOrderListBuilder {
	builder := new(NewOrderListBuilder)
	builder.Body.Set(listid)
	builder.Body.Set(bidtype)
	builder.Body.Set(totnoorders)
	builder.Body.Set(noorders)
	return builder
}

//ListID is a required field for NewOrderList.
func (m *NewOrderList) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//BidID is a non-required field for NewOrderList.
func (m *NewOrderList) BidID() (*field.BidID, error) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}

//ClientBidID is a non-required field for NewOrderList.
func (m *NewOrderList) ClientBidID() (*field.ClientBidID, error) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}

//ProgRptReqs is a non-required field for NewOrderList.
func (m *NewOrderList) ProgRptReqs() (*field.ProgRptReqs, error) {
	f := new(field.ProgRptReqs)
	err := m.Body.Get(f)
	return f, err
}

//BidType is a required field for NewOrderList.
func (m *NewOrderList) BidType() (*field.BidType, error) {
	f := new(field.BidType)
	err := m.Body.Get(f)
	return f, err
}

//ProgPeriodInterval is a non-required field for NewOrderList.
func (m *NewOrderList) ProgPeriodInterval() (*field.ProgPeriodInterval, error) {
	f := new(field.ProgPeriodInterval)
	err := m.Body.Get(f)
	return f, err
}

//CancellationRights is a non-required field for NewOrderList.
func (m *NewOrderList) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for NewOrderList.
func (m *NewOrderList) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//RegistID is a non-required field for NewOrderList.
func (m *NewOrderList) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//ListExecInstType is a non-required field for NewOrderList.
func (m *NewOrderList) ListExecInstType() (*field.ListExecInstType, error) {
	f := new(field.ListExecInstType)
	err := m.Body.Get(f)
	return f, err
}

//ListExecInst is a non-required field for NewOrderList.
func (m *NewOrderList) ListExecInst() (*field.ListExecInst, error) {
	f := new(field.ListExecInst)
	err := m.Body.Get(f)
	return f, err
}

//EncodedListExecInstLen is a non-required field for NewOrderList.
func (m *NewOrderList) EncodedListExecInstLen() (*field.EncodedListExecInstLen, error) {
	f := new(field.EncodedListExecInstLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedListExecInst is a non-required field for NewOrderList.
func (m *NewOrderList) EncodedListExecInst() (*field.EncodedListExecInst, error) {
	f := new(field.EncodedListExecInst)
	err := m.Body.Get(f)
	return f, err
}

//AllowableOneSidednessPct is a non-required field for NewOrderList.
func (m *NewOrderList) AllowableOneSidednessPct() (*field.AllowableOneSidednessPct, error) {
	f := new(field.AllowableOneSidednessPct)
	err := m.Body.Get(f)
	return f, err
}

//AllowableOneSidednessValue is a non-required field for NewOrderList.
func (m *NewOrderList) AllowableOneSidednessValue() (*field.AllowableOneSidednessValue, error) {
	f := new(field.AllowableOneSidednessValue)
	err := m.Body.Get(f)
	return f, err
}

//AllowableOneSidednessCurr is a non-required field for NewOrderList.
func (m *NewOrderList) AllowableOneSidednessCurr() (*field.AllowableOneSidednessCurr, error) {
	f := new(field.AllowableOneSidednessCurr)
	err := m.Body.Get(f)
	return f, err
}

//TotNoOrders is a required field for NewOrderList.
func (m *NewOrderList) TotNoOrders() (*field.TotNoOrders, error) {
	f := new(field.TotNoOrders)
	err := m.Body.Get(f)
	return f, err
}

//LastFragment is a non-required field for NewOrderList.
func (m *NewOrderList) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//NoOrders is a required field for NewOrderList.
func (m *NewOrderList) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//NoRootPartyIDs is a non-required field for NewOrderList.
func (m *NewOrderList) NoRootPartyIDs() (*field.NoRootPartyIDs, error) {
	f := new(field.NoRootPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
