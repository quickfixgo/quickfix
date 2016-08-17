package dontknowtrade

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//DontKnowTrade is the fix41 DontKnowTrade type, MsgType = Q
type DontKnowTrade struct {
	fix41.Header
	quickfix.Body
	fix41.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a DontKnowTrade from a quickfix.Message instance
func FromMessage(m quickfix.Message) DontKnowTrade {
	return DontKnowTrade{
		Header:      fix41.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix41.Trailer{Trailer: m.Trailer},
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
func New(dkreason field.DKReasonField, symbol field.SymbolField, side field.SideField) (m DontKnowTrade) {
	m.Header = fix41.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("Q"))
	m.Set(dkreason)
	m.Set(symbol)
	m.Set(side)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg DontKnowTrade, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "Q", r
}

//SetExecID sets ExecID, Tag 17
func (m DontKnowTrade) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetIDSource sets IDSource, Tag 22
func (m DontKnowTrade) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
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

//SetSecurityID sets SecurityID, Tag 48
func (m DontKnowTrade) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
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

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m DontKnowTrade) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m DontKnowTrade) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m DontKnowTrade) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetDKReason sets DKReason, Tag 127
func (m DontKnowTrade) SetDKReason(v string) {
	m.Set(field.NewDKReason(v))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m DontKnowTrade) SetCashOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOrderQty(value, scale))
}

//SetSecurityType sets SecurityType, Tag 167
func (m DontKnowTrade) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m DontKnowTrade) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m DontKnowTrade) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m DontKnowTrade) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m DontKnowTrade) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m DontKnowTrade) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m DontKnowTrade) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//GetExecID gets ExecID, Tag 17
func (m DontKnowTrade) GetExecID() (f field.ExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m DontKnowTrade) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
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

//GetSecurityID gets SecurityID, Tag 48
func (m DontKnowTrade) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
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

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m DontKnowTrade) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m DontKnowTrade) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m DontKnowTrade) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDKReason gets DKReason, Tag 127
func (m DontKnowTrade) GetDKReason() (f field.DKReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m DontKnowTrade) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m DontKnowTrade) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m DontKnowTrade) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m DontKnowTrade) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m DontKnowTrade) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m DontKnowTrade) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m DontKnowTrade) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m DontKnowTrade) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasExecID returns true if ExecID is present, Tag 17
func (m DontKnowTrade) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m DontKnowTrade) HasIDSource() bool {
	return m.Has(tag.IDSource)
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

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m DontKnowTrade) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
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

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m DontKnowTrade) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m DontKnowTrade) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m DontKnowTrade) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasDKReason returns true if DKReason is present, Tag 127
func (m DontKnowTrade) HasDKReason() bool {
	return m.Has(tag.DKReason)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m DontKnowTrade) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m DontKnowTrade) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m DontKnowTrade) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m DontKnowTrade) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m DontKnowTrade) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m DontKnowTrade) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m DontKnowTrade) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m DontKnowTrade) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}
