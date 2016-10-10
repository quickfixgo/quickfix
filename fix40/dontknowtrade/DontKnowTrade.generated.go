package dontknowtrade

import (
	"github.com/shopspring/decimal"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//DontKnowTrade is the fix40 DontKnowTrade type, MsgType = Q
type DontKnowTrade struct {
	fix40.Header
	*quickfix.Body
	fix40.Trailer
	Message *quickfix.Message
}

//FromMessage creates a DontKnowTrade from a quickfix.Message instance
func FromMessage(m *quickfix.Message) DontKnowTrade {
	return DontKnowTrade{
		Header:  fix40.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix40.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m DontKnowTrade) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a DontKnowTrade initialized with the required fields for DontKnowTrade
func New(dkreason field.DKReasonField, symbol field.SymbolField, side field.SideField, orderqty field.OrderQtyField, lastshares field.LastSharesField, lastpx field.LastPxField) (m DontKnowTrade) {
	m.Message = quickfix.NewMessage()
	m.Header = fix40.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("Q"))
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
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "Q", r
}

//SetExecID sets ExecID, Tag 17
func (m DontKnowTrade) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetLastPx sets LastPx, Tag 31
func (m DontKnowTrade) SetLastPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastPx(value, scale))
}

//SetLastShares sets LastShares, Tag 32
func (m DontKnowTrade) SetLastShares(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastShares(value, scale))
}

//SetOrderID sets OrderID, Tag 37
func (m DontKnowTrade) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m DontKnowTrade) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

//SetSide sets Side, Tag 54
func (m DontKnowTrade) SetSide(v enum.Side) {
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
func (m DontKnowTrade) SetDKReason(v enum.DKReason) {
	m.Set(field.NewDKReason(v))
}

//GetExecID gets ExecID, Tag 17
func (m DontKnowTrade) GetExecID() (v string, err quickfix.MessageRejectError) {
	var f field.ExecIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastPx gets LastPx, Tag 31
func (m DontKnowTrade) GetLastPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastShares gets LastShares, Tag 32
func (m DontKnowTrade) GetLastShares() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastSharesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderID gets OrderID, Tag 37
func (m DontKnowTrade) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m DontKnowTrade) GetOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSide gets Side, Tag 54
func (m DontKnowTrade) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m DontKnowTrade) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m DontKnowTrade) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDKReason gets DKReason, Tag 127
func (m DontKnowTrade) GetDKReason() (v enum.DKReason, err quickfix.MessageRejectError) {
	var f field.DKReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
