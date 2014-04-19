package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
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
	clordid field.ClOrdID) OrderCancelRejectBuilder {
	var builder OrderCancelRejectBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(orderid)
	builder.Body.Set(clordid)
	return builder
}

//OrderID is a required field for OrderCancelReject.
func (m OrderCancelReject) OrderID() (field.OrderID, errors.MessageRejectError) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a required field for OrderCancelReject.
func (m OrderCancelReject) ClOrdID() (field.ClOrdID, errors.MessageRejectError) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClientID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ClientID() (field.ClientID, errors.MessageRejectError) {
	var f field.ClientID
	err := m.Body.Get(&f)
	return f, err
}

//ExecBroker is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ExecBroker() (field.ExecBroker, errors.MessageRejectError) {
	var f field.ExecBroker
	err := m.Body.Get(&f)
	return f, err
}

//ListID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ListID() (field.ListID, errors.MessageRejectError) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//CxlRejReason is a non-required field for OrderCancelReject.
func (m OrderCancelReject) CxlRejReason() (field.CxlRejReason, errors.MessageRejectError) {
	var f field.CxlRejReason
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for OrderCancelReject.
func (m OrderCancelReject) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}
