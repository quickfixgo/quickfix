package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type NewOrderList struct {
	message.Message
}

func (m *NewOrderList) BidID() (*field.BidID, error) {
	f := new(field.BidID)
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
func (m *NewOrderList) TotNoOrders() (*field.TotNoOrders, error) {
	f := new(field.TotNoOrders)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) ListID() (*field.ListID, error) {
	f := new(field.ListID)
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
func (m *NewOrderList) BidType() (*field.BidType, error) {
	f := new(field.BidType)
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
func (m *NewOrderList) EncodedListExecInst() (*field.EncodedListExecInst, error) {
	f := new(field.EncodedListExecInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderList) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}
