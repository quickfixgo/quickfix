package ordercancelreject

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//OrderCancelReject is the fix40 OrderCancelReject type, MsgType = 9
type OrderCancelReject struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a OrderCancelReject from a quickfix.Message instance
func FromMessage(m quickfix.Message) OrderCancelReject {
	return OrderCancelReject{
		Header:      fix40.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix40.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m OrderCancelReject) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a OrderCancelReject initialized with the required fields for OrderCancelReject
func New(orderid field.OrderIDField, clordid field.ClOrdIDField) (m OrderCancelReject) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("9"))
	m.Set(orderid)
	m.Set(clordid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg OrderCancelReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "9", r
}

//SetClOrdID sets ClOrdID, Tag 11
func (m OrderCancelReject) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetOrderID sets OrderID, Tag 37
func (m OrderCancelReject) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetText sets Text, Tag 58
func (m OrderCancelReject) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetListID sets ListID, Tag 66
func (m OrderCancelReject) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetExecBroker sets ExecBroker, Tag 76
func (m OrderCancelReject) SetExecBroker(v string) {
	m.Set(field.NewExecBroker(v))
}

//SetCxlRejReason sets CxlRejReason, Tag 102
func (m OrderCancelReject) SetCxlRejReason(v int) {
	m.Set(field.NewCxlRejReason(v))
}

//SetClientID sets ClientID, Tag 109
func (m OrderCancelReject) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

//GetClOrdID gets ClOrdID, Tag 11
func (m OrderCancelReject) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m OrderCancelReject) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m OrderCancelReject) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m OrderCancelReject) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m OrderCancelReject) GetExecBroker() (f field.ExecBrokerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCxlRejReason gets CxlRejReason, Tag 102
func (m OrderCancelReject) GetCxlRejReason() (f field.CxlRejReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientID gets ClientID, Tag 109
func (m OrderCancelReject) GetClientID() (f field.ClientIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m OrderCancelReject) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m OrderCancelReject) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasText returns true if Text is present, Tag 58
func (m OrderCancelReject) HasText() bool {
	return m.Has(tag.Text)
}

//HasListID returns true if ListID is present, Tag 66
func (m OrderCancelReject) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasExecBroker returns true if ExecBroker is present, Tag 76
func (m OrderCancelReject) HasExecBroker() bool {
	return m.Has(tag.ExecBroker)
}

//HasCxlRejReason returns true if CxlRejReason is present, Tag 102
func (m OrderCancelReject) HasCxlRejReason() bool {
	return m.Has(tag.CxlRejReason)
}

//HasClientID returns true if ClientID is present, Tag 109
func (m OrderCancelReject) HasClientID() bool {
	return m.Has(tag.ClientID)
}
