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

//CreateOrderCancelRejectBuilder returns an initialized OrderCancelRejectBuilder with specified required fields.
func CreateOrderCancelRejectBuilder(
	orderid field.OrderID,
	clordid field.ClOrdID,
	origclordid field.OrigClOrdID,
	ordstatus field.OrdStatus,
	cxlrejresponseto field.CxlRejResponseTo) OrderCancelRejectBuilder {
	var builder OrderCancelRejectBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(orderid)
	builder.Body.Set(clordid)
	builder.Body.Set(origclordid)
	builder.Body.Set(ordstatus)
	builder.Body.Set(cxlrejresponseto)
	return builder
}

//OrderID is a required field for OrderCancelReject.
func (m OrderCancelReject) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) SecondaryOrderID() (field.SecondaryOrderID, error) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a required field for OrderCancelReject.
func (m OrderCancelReject) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrigClOrdID is a required field for OrderCancelReject.
func (m OrderCancelReject) OrigClOrdID() (field.OrigClOrdID, error) {
	var f field.OrigClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrdStatus is a required field for OrderCancelReject.
func (m OrderCancelReject) OrdStatus() (field.OrdStatus, error) {
	var f field.OrdStatus
	err := m.Body.Get(&f)
	return f, err
}

//ClientID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ClientID() (field.ClientID, error) {
	var f field.ClientID
	err := m.Body.Get(&f)
	return f, err
}

//ExecBroker is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ExecBroker() (field.ExecBroker, error) {
	var f field.ExecBroker
	err := m.Body.Get(&f)
	return f, err
}

//ListID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ListID() (field.ListID, error) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for OrderCancelReject.
func (m OrderCancelReject) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for OrderCancelReject.
func (m OrderCancelReject) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//CxlRejResponseTo is a required field for OrderCancelReject.
func (m OrderCancelReject) CxlRejResponseTo() (field.CxlRejResponseTo, error) {
	var f field.CxlRejResponseTo
	err := m.Body.Get(&f)
	return f, err
}

//CxlRejReason is a non-required field for OrderCancelReject.
func (m OrderCancelReject) CxlRejReason() (field.CxlRejReason, error) {
	var f field.CxlRejReason
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for OrderCancelReject.
func (m OrderCancelReject) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for OrderCancelReject.
func (m OrderCancelReject) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for OrderCancelReject.
func (m OrderCancelReject) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
