package quote

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//Quote is the fix40 Quote type, MsgType = S
type Quote struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a Quote from a quickfix.Message instance
func FromMessage(m quickfix.Message) Quote {
	return Quote{
		Header:      fix40.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix40.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Quote) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a Quote initialized with the required fields for Quote
func New(quoteid field.QuoteIDField, symbol field.SymbolField, bidpx field.BidPxField) (m Quote) {
	m.Header = fix40.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("S"))
	m.Header.Set(field.NewBeginString("FIX.4.0"))
	m.Set(quoteid)
	m.Set(symbol)
	m.Set(bidpx)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Quote, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "S", r
}

//SetIDSource sets IDSource, Tag 22
func (m Quote) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m Quote) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m Quote) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetValidUntilTime sets ValidUntilTime, Tag 62
func (m Quote) SetValidUntilTime(v time.Time) {
	m.Set(field.NewValidUntilTime(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m Quote) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m Quote) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m Quote) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetQuoteID sets QuoteID, Tag 117
func (m Quote) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

//SetQuoteReqID sets QuoteReqID, Tag 131
func (m Quote) SetQuoteReqID(v string) {
	m.Set(field.NewQuoteReqID(v))
}

//SetBidPx sets BidPx, Tag 132
func (m Quote) SetBidPx(v float64) {
	m.Set(field.NewBidPx(v))
}

//SetOfferPx sets OfferPx, Tag 133
func (m Quote) SetOfferPx(v float64) {
	m.Set(field.NewOfferPx(v))
}

//SetBidSize sets BidSize, Tag 134
func (m Quote) SetBidSize(v float64) {
	m.Set(field.NewBidSize(v))
}

//SetOfferSize sets OfferSize, Tag 135
func (m Quote) SetOfferSize(v float64) {
	m.Set(field.NewOfferSize(v))
}

//GetIDSource gets IDSource, Tag 22
func (m Quote) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m Quote) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m Quote) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetValidUntilTime gets ValidUntilTime, Tag 62
func (m Quote) GetValidUntilTime() (f field.ValidUntilTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m Quote) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m Quote) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m Quote) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteID gets QuoteID, Tag 117
func (m Quote) GetQuoteID() (f field.QuoteIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteReqID gets QuoteReqID, Tag 131
func (m Quote) GetQuoteReqID() (f field.QuoteReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidPx gets BidPx, Tag 132
func (m Quote) GetBidPx() (f field.BidPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOfferPx gets OfferPx, Tag 133
func (m Quote) GetOfferPx() (f field.OfferPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidSize gets BidSize, Tag 134
func (m Quote) GetBidSize() (f field.BidSizeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOfferSize gets OfferSize, Tag 135
func (m Quote) GetOfferSize() (f field.OfferSizeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m Quote) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m Quote) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m Quote) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasValidUntilTime returns true if ValidUntilTime is present, Tag 62
func (m Quote) HasValidUntilTime() bool {
	return m.Has(tag.ValidUntilTime)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m Quote) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m Quote) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m Quote) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasQuoteID returns true if QuoteID is present, Tag 117
func (m Quote) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

//HasQuoteReqID returns true if QuoteReqID is present, Tag 131
func (m Quote) HasQuoteReqID() bool {
	return m.Has(tag.QuoteReqID)
}

//HasBidPx returns true if BidPx is present, Tag 132
func (m Quote) HasBidPx() bool {
	return m.Has(tag.BidPx)
}

//HasOfferPx returns true if OfferPx is present, Tag 133
func (m Quote) HasOfferPx() bool {
	return m.Has(tag.OfferPx)
}

//HasBidSize returns true if BidSize is present, Tag 134
func (m Quote) HasBidSize() bool {
	return m.Has(tag.BidSize)
}

//HasOfferSize returns true if OfferSize is present, Tag 135
func (m Quote) HasOfferSize() bool {
	return m.Has(tag.OfferSize)
}
