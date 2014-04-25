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
	builder.Header.Set(field.BuildMsgType("9"))
	builder.Body.Set(orderid)
	builder.Body.Set(clordid)
	return builder
}

//OrderID is a required field for OrderCancelReject.
func (m OrderCancelReject) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderCancelReject.
func (m OrderCancelReject) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for OrderCancelReject.
func (m OrderCancelReject) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderCancelReject.
func (m OrderCancelReject) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ClientID() (*field.ClientID, errors.MessageRejectError) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from OrderCancelReject.
func (m OrderCancelReject) GetClientID(f *field.ClientID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ExecBroker() (*field.ExecBroker, errors.MessageRejectError) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from OrderCancelReject.
func (m OrderCancelReject) GetExecBroker(f *field.ExecBroker) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from OrderCancelReject.
func (m OrderCancelReject) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CxlRejReason is a non-required field for OrderCancelReject.
func (m OrderCancelReject) CxlRejReason() (*field.CxlRejReason, errors.MessageRejectError) {
	f := new(field.CxlRejReason)
	err := m.Body.Get(f)
	return f, err
}

//GetCxlRejReason reads a CxlRejReason from OrderCancelReject.
func (m OrderCancelReject) GetCxlRejReason(f *field.CxlRejReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderCancelReject.
func (m OrderCancelReject) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderCancelReject.
func (m OrderCancelReject) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}
