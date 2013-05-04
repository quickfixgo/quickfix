package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type OrderMassCancelReport struct {
	quickfixgo.Message
}

func (m *OrderMassCancelReport) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MassCancelRequestType() (*field.MassCancelRequestType, error) {
	f := new(field.MassCancelRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MassCancelResponse() (*field.MassCancelResponse, error) {
	f := new(field.MassCancelResponse)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) MassCancelRejectReason() (*field.MassCancelRejectReason, error) {
	f := new(field.MassCancelRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) TotalAffectedOrders() (*field.TotalAffectedOrders, error) {
	f := new(field.TotalAffectedOrders)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *OrderMassCancelReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
