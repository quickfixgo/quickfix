package ordercancelreject

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//OrderCancelReject is the fix42 OrderCancelReject type, MsgType = 9
type OrderCancelReject struct {
	fix42.Header
	quickfix.Body
	fix42.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a OrderCancelReject from a quickfix.Message instance
func FromMessage(m quickfix.Message) OrderCancelReject {
	return OrderCancelReject{
		Header:      fix42.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix42.Trailer{Trailer: m.Trailer},
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
func New(orderid field.OrderIDField, clordid field.ClOrdIDField, origclordid field.OrigClOrdIDField, ordstatus field.OrdStatusField, cxlrejresponseto field.CxlRejResponseToField) (m OrderCancelReject) {
	m.Header = fix42.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("9"))
	m.Header.Set(field.NewBeginString("FIX.4.2"))
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
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
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
func (m OrderCancelReject) SetOrdStatus(v string) {
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
func (m OrderCancelReject) SetCxlRejReason(v int) {
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
func (m OrderCancelReject) SetCxlRejResponseTo(v string) {
	m.Set(field.NewCxlRejResponseTo(v))
}

//GetAccount gets Account, Tag 1
func (m OrderCancelReject) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
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

//GetOrdStatus gets OrdStatus, Tag 39
func (m OrderCancelReject) GetOrdStatus() (f field.OrdStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigClOrdID gets OrigClOrdID, Tag 41
func (m OrderCancelReject) GetOrigClOrdID() (f field.OrigClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m OrderCancelReject) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m OrderCancelReject) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
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

//GetSecondaryOrderID gets SecondaryOrderID, Tag 198
func (m OrderCancelReject) GetSecondaryOrderID() (f field.SecondaryOrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m OrderCancelReject) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m OrderCancelReject) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCxlRejResponseTo gets CxlRejResponseTo, Tag 434
func (m OrderCancelReject) GetCxlRejResponseTo() (f field.CxlRejResponseToField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
