//Package liststatus msg type = N.
package liststatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a ListStatus wrapper for the generic Message type
type Message struct {
	message.Message
}

//ListID is a required field for ListStatus.
func (m Message) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListStatus.
func (m Message) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListStatusType is a required field for ListStatus.
func (m Message) ListStatusType() (*field.ListStatusTypeField, errors.MessageRejectError) {
	f := &field.ListStatusTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListStatusType reads a ListStatusType from ListStatus.
func (m Message) GetListStatusType(f *field.ListStatusTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRpts is a required field for ListStatus.
func (m Message) NoRpts() (*field.NoRptsField, errors.MessageRejectError) {
	f := &field.NoRptsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRpts reads a NoRpts from ListStatus.
func (m Message) GetNoRpts(f *field.NoRptsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListOrderStatus is a required field for ListStatus.
func (m Message) ListOrderStatus() (*field.ListOrderStatusField, errors.MessageRejectError) {
	f := &field.ListOrderStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListOrderStatus reads a ListOrderStatus from ListStatus.
func (m Message) GetListOrderStatus(f *field.ListOrderStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RptSeq is a required field for ListStatus.
func (m Message) RptSeq() (*field.RptSeqField, errors.MessageRejectError) {
	f := &field.RptSeqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRptSeq reads a RptSeq from ListStatus.
func (m Message) GetRptSeq(f *field.RptSeqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListStatusText is a non-required field for ListStatus.
func (m Message) ListStatusText() (*field.ListStatusTextField, errors.MessageRejectError) {
	f := &field.ListStatusTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListStatusText reads a ListStatusText from ListStatus.
func (m Message) GetListStatusText(f *field.ListStatusTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListStatusTextLen is a non-required field for ListStatus.
func (m Message) EncodedListStatusTextLen() (*field.EncodedListStatusTextLenField, errors.MessageRejectError) {
	f := &field.EncodedListStatusTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListStatusTextLen reads a EncodedListStatusTextLen from ListStatus.
func (m Message) GetEncodedListStatusTextLen(f *field.EncodedListStatusTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListStatusText is a non-required field for ListStatus.
func (m Message) EncodedListStatusText() (*field.EncodedListStatusTextField, errors.MessageRejectError) {
	f := &field.EncodedListStatusTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListStatusText reads a EncodedListStatusText from ListStatus.
func (m Message) GetEncodedListStatusText(f *field.EncodedListStatusTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for ListStatus.
func (m Message) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ListStatus.
func (m Message) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoOrders is a required field for ListStatus.
func (m Message) TotNoOrders() (*field.TotNoOrdersField, errors.MessageRejectError) {
	f := &field.TotNoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoOrders reads a TotNoOrders from ListStatus.
func (m Message) GetTotNoOrders(f *field.TotNoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for ListStatus.
func (m Message) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from ListStatus.
func (m Message) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a required field for ListStatus.
func (m Message) NoOrders() (*field.NoOrdersField, errors.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from ListStatus.
func (m Message) GetNoOrders(f *field.NoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds ListStatus messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for ListStatus.
func Builder(
	listid *field.ListIDField,
	liststatustype *field.ListStatusTypeField,
	norpts *field.NoRptsField,
	listorderstatus *field.ListOrderStatusField,
	rptseq *field.RptSeqField,
	totnoorders *field.TotNoOrdersField,
	noorders *field.NoOrdersField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header().Set(field.NewMsgType("N"))
	builder.Body().Set(listid)
	builder.Body().Set(liststatustype)
	builder.Body().Set(norpts)
	builder.Body().Set(listorderstatus)
	builder.Body().Set(rptseq)
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
	return fix.BeginString_FIX50, "N", r
}
