package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type OrderStatusRequest struct {
	quickfixgo.Message
}

func (m *OrderStatusRequest) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderStatusRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderStatusRequest) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderStatusRequest) ClOrdLinkID() (*field.ClOrdLinkID, error) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderStatusRequest) OrdStatusReqID() (*field.OrdStatusReqID, error) {
	f := new(field.OrdStatusReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderStatusRequest) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderStatusRequest) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderStatusRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
