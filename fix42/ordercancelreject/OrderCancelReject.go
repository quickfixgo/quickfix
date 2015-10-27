//Package ordercancelreject msg type = 9.
package ordercancelreject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a OrderCancelReject wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//OrderID is a required field for OrderCancelReject.
func (m Message) OrderID() (*field.OrderIDField, quickfix.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderCancelReject.
func (m Message) GetOrderID(f *field.OrderIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for OrderCancelReject.
func (m Message) SecondaryOrderID() (*field.SecondaryOrderIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryOrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from OrderCancelReject.
func (m Message) GetSecondaryOrderID(f *field.SecondaryOrderIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for OrderCancelReject.
func (m Message) ClOrdID() (*field.ClOrdIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderCancelReject.
func (m Message) GetClOrdID(f *field.ClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a required field for OrderCancelReject.
func (m Message) OrigClOrdID() (*field.OrigClOrdIDField, quickfix.MessageRejectError) {
	f := &field.OrigClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from OrderCancelReject.
func (m Message) GetOrigClOrdID(f *field.OrigClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatus is a required field for OrderCancelReject.
func (m Message) OrdStatus() (*field.OrdStatusField, quickfix.MessageRejectError) {
	f := &field.OrdStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatus reads a OrdStatus from OrderCancelReject.
func (m Message) GetOrdStatus(f *field.OrdStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for OrderCancelReject.
func (m Message) ClientID() (*field.ClientIDField, quickfix.MessageRejectError) {
	f := &field.ClientIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from OrderCancelReject.
func (m Message) GetClientID(f *field.ClientIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for OrderCancelReject.
func (m Message) ExecBroker() (*field.ExecBrokerField, quickfix.MessageRejectError) {
	f := &field.ExecBrokerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from OrderCancelReject.
func (m Message) GetExecBroker(f *field.ExecBrokerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for OrderCancelReject.
func (m Message) ListID() (*field.ListIDField, quickfix.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from OrderCancelReject.
func (m Message) GetListID(f *field.ListIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for OrderCancelReject.
func (m Message) Account() (*field.AccountField, quickfix.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from OrderCancelReject.
func (m Message) GetAccount(f *field.AccountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for OrderCancelReject.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from OrderCancelReject.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CxlRejResponseTo is a required field for OrderCancelReject.
func (m Message) CxlRejResponseTo() (*field.CxlRejResponseToField, quickfix.MessageRejectError) {
	f := &field.CxlRejResponseToField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCxlRejResponseTo reads a CxlRejResponseTo from OrderCancelReject.
func (m Message) GetCxlRejResponseTo(f *field.CxlRejResponseToField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CxlRejReason is a non-required field for OrderCancelReject.
func (m Message) CxlRejReason() (*field.CxlRejReasonField, quickfix.MessageRejectError) {
	f := &field.CxlRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCxlRejReason reads a CxlRejReason from OrderCancelReject.
func (m Message) GetCxlRejReason(f *field.CxlRejReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderCancelReject.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderCancelReject.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for OrderCancelReject.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from OrderCancelReject.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for OrderCancelReject.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from OrderCancelReject.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for OrderCancelReject.
func New(
	orderid *field.OrderIDField,
	clordid *field.ClOrdIDField,
	origclordid *field.OrigClOrdIDField,
	ordstatus *field.OrdStatusField,
	cxlrejresponseto *field.CxlRejResponseToField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX42))
	builder.Header.Set(field.NewMsgType("9"))
	builder.Body.Set(orderid)
	builder.Body.Set(clordid)
	builder.Body.Set(origclordid)
	builder.Body.Set(ordstatus)
	builder.Body.Set(cxlrejresponseto)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX42, "9", r
}
