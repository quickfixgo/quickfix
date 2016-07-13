package dontknowtrade

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//DontKnowTrade is the fix40 DontKnowTrade type, MsgType = Q
type DontKnowTrade struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a DontKnowTrade from a quickfix.Message instance
func FromMessage(m quickfix.Message) DontKnowTrade {
	return DontKnowTrade{
		Header:      fix40.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix40.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m DontKnowTrade) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a DontKnowTrade initialized with the required fields for DontKnowTrade
func New(dkreason field.DKReasonField, symbol field.SymbolField, side field.SideField, orderqty field.OrderQtyField, lastshares field.LastSharesField, lastpx field.LastPxField) (m DontKnowTrade) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("Q"))
	m.Header.Set(field.NewBeginString("FIX.4.0"))
	m.Set(dkreason)
	m.Set(symbol)
	m.Set(side)
	m.Set(orderqty)
	m.Set(lastshares)
	m.Set(lastpx)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg DontKnowTrade, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "Q", r
}

//SetExecID sets ExecID, Tag 17
func (m DontKnowTrade) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetLastPx sets LastPx, Tag 31
func (m DontKnowTrade) SetLastPx(v float64) {
	m.Set(field.NewLastPx(v))
}

//SetLastShares sets LastShares, Tag 32
func (m DontKnowTrade) SetLastShares(v float64) {
	m.Set(field.NewLastShares(v))
}

//SetOrderID sets OrderID, Tag 37
func (m DontKnowTrade) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m DontKnowTrade) SetOrderQty(v float64) {
	m.Set(field.NewOrderQty(v))
}

//SetSide sets Side, Tag 54
func (m DontKnowTrade) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m DontKnowTrade) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m DontKnowTrade) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetDKReason sets DKReason, Tag 127
func (m DontKnowTrade) SetDKReason(v string) {
	m.Set(field.NewDKReason(v))
}

//GetExecID gets ExecID, Tag 17
func (m DontKnowTrade) GetExecID() (f field.ExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastPx gets LastPx, Tag 31
func (m DontKnowTrade) GetLastPx() (f field.LastPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastShares gets LastShares, Tag 32
func (m DontKnowTrade) GetLastShares() (f field.LastSharesField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m DontKnowTrade) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m DontKnowTrade) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m DontKnowTrade) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m DontKnowTrade) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m DontKnowTrade) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDKReason gets DKReason, Tag 127
func (m DontKnowTrade) GetDKReason() (f field.DKReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasExecID returns true if ExecID is present, Tag 17
func (m DontKnowTrade) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//HasLastPx returns true if LastPx is present, Tag 31
func (m DontKnowTrade) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

//HasLastShares returns true if LastShares is present, Tag 32
func (m DontKnowTrade) HasLastShares() bool {
	return m.Has(tag.LastShares)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m DontKnowTrade) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m DontKnowTrade) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasSide returns true if Side is present, Tag 54
func (m DontKnowTrade) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m DontKnowTrade) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m DontKnowTrade) HasText() bool {
	return m.Has(tag.Text)
}

//HasDKReason returns true if DKReason is present, Tag 127
func (m DontKnowTrade) HasDKReason() bool {
	return m.Has(tag.DKReason)
}
