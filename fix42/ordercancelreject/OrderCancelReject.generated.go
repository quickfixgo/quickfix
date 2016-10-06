package ordercancelreject

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//OrderCancelReject is the fix42 OrderCancelReject type, MsgType = 9
type OrderCancelReject struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

//FromMessage creates a OrderCancelReject from a quickfix.Message instance
func FromMessage(m *quickfix.Message) OrderCancelReject {
	return OrderCancelReject{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m OrderCancelReject) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a OrderCancelReject initialized with the required fields for OrderCancelReject
func New(orderid field.OrderIDField, clordid field.ClOrdIDField, origclordid field.OrigClOrdIDField, ordstatus field.OrdStatusField, cxlrejresponseto field.CxlRejResponseToField) (m OrderCancelReject) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("9"))
	m.Set(orderid)
	m.Set(clordid)
	m.Set(origclordid)
	m.Set(ordstatus)
	m.Set(cxlrejresponseto)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg OrderCancelReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "9", r
}

//SetAccount sets Account, Tag 1
func (m OrderCancelReject) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m OrderCancelReject) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetOrderID sets OrderID, Tag 37
func (m OrderCancelReject) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrdStatus sets OrdStatus, Tag 39
func (m OrderCancelReject) SetOrdStatus(v enum.OrdStatus) {
	m.Set(field.NewOrdStatus(v))
}

//SetOrigClOrdID sets OrigClOrdID, Tag 41
func (m OrderCancelReject) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

//SetText sets Text, Tag 58
func (m OrderCancelReject) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m OrderCancelReject) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
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
func (m OrderCancelReject) SetCxlRejReason(v enum.CxlRejReason) {
	m.Set(field.NewCxlRejReason(v))
}

//SetClientID sets ClientID, Tag 109
func (m OrderCancelReject) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

//SetSecondaryOrderID sets SecondaryOrderID, Tag 198
func (m OrderCancelReject) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m OrderCancelReject) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m OrderCancelReject) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetCxlRejResponseTo sets CxlRejResponseTo, Tag 434
func (m OrderCancelReject) SetCxlRejResponseTo(v enum.CxlRejResponseTo) {
	m.Set(field.NewCxlRejResponseTo(v))
}

//GetAccount gets Account, Tag 1
func (m OrderCancelReject) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m OrderCancelReject) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderID gets OrderID, Tag 37
func (m OrderCancelReject) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrdStatus gets OrdStatus, Tag 39
func (m OrderCancelReject) GetOrdStatus() (v enum.OrdStatus, err quickfix.MessageRejectError) {
	var f field.OrdStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrigClOrdID gets OrigClOrdID, Tag 41
func (m OrderCancelReject) GetOrigClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.OrigClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m OrderCancelReject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m OrderCancelReject) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetListID gets ListID, Tag 66
func (m OrderCancelReject) GetListID() (v string, err quickfix.MessageRejectError) {
	var f field.ListIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m OrderCancelReject) GetExecBroker() (v string, err quickfix.MessageRejectError) {
	var f field.ExecBrokerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCxlRejReason gets CxlRejReason, Tag 102
func (m OrderCancelReject) GetCxlRejReason() (v enum.CxlRejReason, err quickfix.MessageRejectError) {
	var f field.CxlRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClientID gets ClientID, Tag 109
func (m OrderCancelReject) GetClientID() (v string, err quickfix.MessageRejectError) {
	var f field.ClientIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryOrderID gets SecondaryOrderID, Tag 198
func (m OrderCancelReject) GetSecondaryOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryOrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m OrderCancelReject) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m OrderCancelReject) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCxlRejResponseTo gets CxlRejResponseTo, Tag 434
func (m OrderCancelReject) GetCxlRejResponseTo() (v enum.CxlRejResponseTo, err quickfix.MessageRejectError) {
	var f field.CxlRejResponseToField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m OrderCancelReject) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m OrderCancelReject) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m OrderCancelReject) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasOrdStatus returns true if OrdStatus is present, Tag 39
func (m OrderCancelReject) HasOrdStatus() bool {
	return m.Has(tag.OrdStatus)
}

//HasOrigClOrdID returns true if OrigClOrdID is present, Tag 41
func (m OrderCancelReject) HasOrigClOrdID() bool {
	return m.Has(tag.OrigClOrdID)
}

//HasText returns true if Text is present, Tag 58
func (m OrderCancelReject) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m OrderCancelReject) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
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

//HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198
func (m OrderCancelReject) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m OrderCancelReject) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m OrderCancelReject) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasCxlRejResponseTo returns true if CxlRejResponseTo is present, Tag 434
func (m OrderCancelReject) HasCxlRejResponseTo() bool {
	return m.Has(tag.CxlRejResponseTo)
}
