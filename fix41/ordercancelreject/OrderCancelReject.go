//Package ordercancelreject msg type = 9.
package ordercancelreject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a OrderCancelReject wrapper for the generic Message type
type Message struct {
	message.Message
}

//OrderID is a required field for OrderCancelReject.
func (m Message) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderCancelReject.
func (m Message) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for OrderCancelReject.
func (m Message) SecondaryOrderID() (*field.SecondaryOrderIDField, errors.MessageRejectError) {
	f := &field.SecondaryOrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from OrderCancelReject.
func (m Message) GetSecondaryOrderID(f *field.SecondaryOrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for OrderCancelReject.
func (m Message) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderCancelReject.
func (m Message) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a required field for OrderCancelReject.
func (m Message) OrigClOrdID() (*field.OrigClOrdIDField, errors.MessageRejectError) {
	f := &field.OrigClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from OrderCancelReject.
func (m Message) GetOrigClOrdID(f *field.OrigClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatus is a required field for OrderCancelReject.
func (m Message) OrdStatus() (*field.OrdStatusField, errors.MessageRejectError) {
	f := &field.OrdStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatus reads a OrdStatus from OrderCancelReject.
func (m Message) GetOrdStatus(f *field.OrdStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for OrderCancelReject.
func (m Message) ClientID() (*field.ClientIDField, errors.MessageRejectError) {
	f := &field.ClientIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from OrderCancelReject.
func (m Message) GetClientID(f *field.ClientIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for OrderCancelReject.
func (m Message) ExecBroker() (*field.ExecBrokerField, errors.MessageRejectError) {
	f := &field.ExecBrokerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from OrderCancelReject.
func (m Message) GetExecBroker(f *field.ExecBrokerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for OrderCancelReject.
func (m Message) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from OrderCancelReject.
func (m Message) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CxlRejReason is a non-required field for OrderCancelReject.
func (m Message) CxlRejReason() (*field.CxlRejReasonField, errors.MessageRejectError) {
	f := &field.CxlRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCxlRejReason reads a CxlRejReason from OrderCancelReject.
func (m Message) GetCxlRejReason(f *field.CxlRejReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderCancelReject.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderCancelReject.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds OrderCancelReject messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for OrderCancelReject.
func Builder(
	orderid *field.OrderIDField,
	clordid *field.ClOrdIDField,
	origclordid *field.OrigClOrdIDField,
	ordstatus *field.OrdStatusField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX41))
	builder.Header().Set(field.NewMsgType("9"))
	builder.Body().Set(orderid)
	builder.Body().Set(clordid)
	builder.Body().Set(origclordid)
	builder.Body().Set(ordstatus)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX41, "9", r
}
