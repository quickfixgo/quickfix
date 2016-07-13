package quoterequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//QuoteRequest is the fix40 QuoteRequest type, MsgType = R
type QuoteRequest struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a QuoteRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) QuoteRequest {
	return QuoteRequest{
		Header:      fix40.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix40.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m QuoteRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a QuoteRequest initialized with the required fields for QuoteRequest
func New(quotereqid field.QuoteReqIDField, symbol field.SymbolField) (m QuoteRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("R"))
	m.Header.Set(field.NewBeginString("FIX.4.0"))
	m.Set(quotereqid)
	m.Set(symbol)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg QuoteRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "R", r
}

//SetIDSource sets IDSource, Tag 22
func (m QuoteRequest) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m QuoteRequest) SetOrderQty(v float64) {
	m.Set(field.NewOrderQty(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m QuoteRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m QuoteRequest) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m QuoteRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m QuoteRequest) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m QuoteRequest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m QuoteRequest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetQuoteReqID sets QuoteReqID, Tag 131
func (m QuoteRequest) SetQuoteReqID(v string) {
	m.Set(field.NewQuoteReqID(v))
}

//SetPrevClosePx sets PrevClosePx, Tag 140
func (m QuoteRequest) SetPrevClosePx(v float64) {
	m.Set(field.NewPrevClosePx(v))
}

//GetIDSource gets IDSource, Tag 22
func (m QuoteRequest) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m QuoteRequest) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m QuoteRequest) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m QuoteRequest) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m QuoteRequest) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m QuoteRequest) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m QuoteRequest) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m QuoteRequest) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteReqID gets QuoteReqID, Tag 131
func (m QuoteRequest) GetQuoteReqID() (f field.QuoteReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrevClosePx gets PrevClosePx, Tag 140
func (m QuoteRequest) GetPrevClosePx() (f field.PrevClosePxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m QuoteRequest) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m QuoteRequest) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m QuoteRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m QuoteRequest) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m QuoteRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m QuoteRequest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m QuoteRequest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m QuoteRequest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasQuoteReqID returns true if QuoteReqID is present, Tag 131
func (m QuoteRequest) HasQuoteReqID() bool {
	return m.Has(tag.QuoteReqID)
}

//HasPrevClosePx returns true if PrevClosePx is present, Tag 140
func (m QuoteRequest) HasPrevClosePx() bool {
	return m.Has(tag.PrevClosePx)
}
