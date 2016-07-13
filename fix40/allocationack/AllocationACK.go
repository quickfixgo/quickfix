package allocationack

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//AllocationACK is the fix40 AllocationACK type, MsgType = P
type AllocationACK struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a AllocationACK from a quickfix.Message instance
func FromMessage(m quickfix.Message) AllocationACK {
	return AllocationACK{
		Header:      fix40.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix40.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m AllocationACK) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a AllocationACK initialized with the required fields for AllocationACK
func New(allocid field.AllocIDField, tradedate field.TradeDateField, allocstatus field.AllocStatusField) (m AllocationACK) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("P"))
	m.Header.Set(field.NewBeginString("FIX.4.0"))
	m.Set(allocid)
	m.Set(tradedate)
	m.Set(allocstatus)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg AllocationACK, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "P", r
}

//SetText sets Text, Tag 58
func (m AllocationACK) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m AllocationACK) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetAllocID sets AllocID, Tag 70
func (m AllocationACK) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m AllocationACK) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetExecBroker sets ExecBroker, Tag 76
func (m AllocationACK) SetExecBroker(v string) {
	m.Set(field.NewExecBroker(v))
}

//SetAllocStatus sets AllocStatus, Tag 87
func (m AllocationACK) SetAllocStatus(v int) {
	m.Set(field.NewAllocStatus(v))
}

//SetAllocRejCode sets AllocRejCode, Tag 88
func (m AllocationACK) SetAllocRejCode(v int) {
	m.Set(field.NewAllocRejCode(v))
}

//SetClientID sets ClientID, Tag 109
func (m AllocationACK) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

//GetText gets Text, Tag 58
func (m AllocationACK) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m AllocationACK) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocID gets AllocID, Tag 70
func (m AllocationACK) GetAllocID() (f field.AllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m AllocationACK) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m AllocationACK) GetExecBroker() (f field.ExecBrokerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocStatus gets AllocStatus, Tag 87
func (m AllocationACK) GetAllocStatus() (f field.AllocStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocRejCode gets AllocRejCode, Tag 88
func (m AllocationACK) GetAllocRejCode() (f field.AllocRejCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientID gets ClientID, Tag 109
func (m AllocationACK) GetClientID() (f field.ClientIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m AllocationACK) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m AllocationACK) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m AllocationACK) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m AllocationACK) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasExecBroker returns true if ExecBroker is present, Tag 76
func (m AllocationACK) HasExecBroker() bool {
	return m.Has(tag.ExecBroker)
}

//HasAllocStatus returns true if AllocStatus is present, Tag 87
func (m AllocationACK) HasAllocStatus() bool {
	return m.Has(tag.AllocStatus)
}

//HasAllocRejCode returns true if AllocRejCode is present, Tag 88
func (m AllocationACK) HasAllocRejCode() bool {
	return m.Has(tag.AllocRejCode)
}

//HasClientID returns true if ClientID is present, Tag 109
func (m AllocationACK) HasClientID() bool {
	return m.Has(tag.ClientID)
}
