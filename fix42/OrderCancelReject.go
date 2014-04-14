package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderCancelReject msg type = 9.
type OrderCancelReject struct {
	message.Message
}

//OrderCancelRejectBuilder builds OrderCancelReject messages.
type OrderCancelRejectBuilder struct {
	message.MessageBuilder
}

//NewOrderCancelRejectBuilder returns an initialized OrderCancelRejectBuilder with specified required fields.
func NewOrderCancelRejectBuilder(
	orderid field.OrderID,
	clordid field.ClOrdID,
	origclordid field.OrigClOrdID,
	ordstatus field.OrdStatus,
	cxlrejresponseto field.CxlRejResponseTo) *OrderCancelRejectBuilder {
	builder := new(OrderCancelRejectBuilder)
	builder.Body.Set(orderid)
	builder.Body.Set(clordid)
	builder.Body.Set(origclordid)
	builder.Body.Set(ordstatus)
	builder.Body.Set(cxlrejresponseto)
	return builder
}

//OrderID is a required field for OrderCancelReject.
func (m *OrderCancelReject) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryOrderID is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a required field for OrderCancelReject.
func (m *OrderCancelReject) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//OrigClOrdID is a required field for OrderCancelReject.
func (m *OrderCancelReject) OrigClOrdID() (*field.OrigClOrdID, error) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//OrdStatus is a required field for OrderCancelReject.
func (m *OrderCancelReject) OrdStatus() (*field.OrdStatus, error) {
	f := new(field.OrdStatus)
	err := m.Body.Get(f)
	return f, err
}

//ClientID is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) ClientID() (*field.ClientID, error) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}

//ExecBroker is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) ExecBroker() (*field.ExecBroker, error) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}

//ListID is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//CxlRejResponseTo is a required field for OrderCancelReject.
func (m *OrderCancelReject) CxlRejResponseTo() (*field.CxlRejResponseTo, error) {
	f := new(field.CxlRejResponseTo)
	err := m.Body.Get(f)
	return f, err
}

//CxlRejReason is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) CxlRejReason() (*field.CxlRejReason, error) {
	f := new(field.CxlRejReason)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
