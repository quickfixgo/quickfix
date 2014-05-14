//Package neworderlist msg type = E.
package neworderlist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a NewOrderList wrapper for the generic Message type
type Message struct {
	message.Message
}

//ListID is a required field for NewOrderList.
func (m Message) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from NewOrderList.
func (m Message) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidID is a non-required field for NewOrderList.
func (m Message) BidID() (*field.BidIDField, errors.MessageRejectError) {
	f := &field.BidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidID reads a BidID from NewOrderList.
func (m Message) GetBidID(f *field.BidIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientBidID is a non-required field for NewOrderList.
func (m Message) ClientBidID() (*field.ClientBidIDField, errors.MessageRejectError) {
	f := &field.ClientBidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientBidID reads a ClientBidID from NewOrderList.
func (m Message) GetClientBidID(f *field.ClientBidIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgRptReqs is a non-required field for NewOrderList.
func (m Message) ProgRptReqs() (*field.ProgRptReqsField, errors.MessageRejectError) {
	f := &field.ProgRptReqsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProgRptReqs reads a ProgRptReqs from NewOrderList.
func (m Message) GetProgRptReqs(f *field.ProgRptReqsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidType is a required field for NewOrderList.
func (m Message) BidType() (*field.BidTypeField, errors.MessageRejectError) {
	f := &field.BidTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidType reads a BidType from NewOrderList.
func (m Message) GetBidType(f *field.BidTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgPeriodInterval is a non-required field for NewOrderList.
func (m Message) ProgPeriodInterval() (*field.ProgPeriodIntervalField, errors.MessageRejectError) {
	f := &field.ProgPeriodIntervalField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProgPeriodInterval reads a ProgPeriodInterval from NewOrderList.
func (m Message) GetProgPeriodInterval(f *field.ProgPeriodIntervalField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListExecInstType is a non-required field for NewOrderList.
func (m Message) ListExecInstType() (*field.ListExecInstTypeField, errors.MessageRejectError) {
	f := &field.ListExecInstTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListExecInstType reads a ListExecInstType from NewOrderList.
func (m Message) GetListExecInstType(f *field.ListExecInstTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListExecInst is a non-required field for NewOrderList.
func (m Message) ListExecInst() (*field.ListExecInstField, errors.MessageRejectError) {
	f := &field.ListExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListExecInst reads a ListExecInst from NewOrderList.
func (m Message) GetListExecInst(f *field.ListExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListExecInstLen is a non-required field for NewOrderList.
func (m Message) EncodedListExecInstLen() (*field.EncodedListExecInstLenField, errors.MessageRejectError) {
	f := &field.EncodedListExecInstLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListExecInstLen reads a EncodedListExecInstLen from NewOrderList.
func (m Message) GetEncodedListExecInstLen(f *field.EncodedListExecInstLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListExecInst is a non-required field for NewOrderList.
func (m Message) EncodedListExecInst() (*field.EncodedListExecInstField, errors.MessageRejectError) {
	f := &field.EncodedListExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListExecInst reads a EncodedListExecInst from NewOrderList.
func (m Message) GetEncodedListExecInst(f *field.EncodedListExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoOrders is a required field for NewOrderList.
func (m Message) TotNoOrders() (*field.TotNoOrdersField, errors.MessageRejectError) {
	f := &field.TotNoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoOrders reads a TotNoOrders from NewOrderList.
func (m Message) GetTotNoOrders(f *field.TotNoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a required field for NewOrderList.
func (m Message) NoOrders() (*field.NoOrdersField, errors.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from NewOrderList.
func (m Message) GetNoOrders(f *field.NoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds NewOrderList messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for NewOrderList.
func Builder(
	listid *field.ListIDField,
	bidtype *field.BidTypeField,
	totnoorders *field.TotNoOrdersField,
	noorders *field.NoOrdersField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header().Set(field.NewMsgType("E"))
	builder.Body().Set(listid)
	builder.Body().Set(bidtype)
	builder.Body().Set(totnoorders)
	builder.Body().Set(noorders)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX42, "E", r
}
