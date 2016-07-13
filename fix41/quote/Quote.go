package quote

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//Quote is the fix41 Quote type, MsgType = S
type Quote struct {
	fix41.Header
	quickfix.Body
	fix41.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a Quote from a quickfix.Message instance
func FromMessage(m quickfix.Message) Quote {
	return Quote{
		Header:      fix41.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix41.Trailer{Trailer: m.Trailer},
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
func New(quoteid field.QuoteIDField, symbol field.SymbolField) (m Quote) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("S"))
	m.Header.Set(field.NewBeginString("FIX.4.1"))
	m.Set(quoteid)
	m.Set(symbol)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Quote, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "S", r
}

//SetIDSource sets IDSource, Tag 22
func (m Quote) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetOrdType sets OrdType, Tag 40
func (m Quote) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m Quote) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m Quote) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m Quote) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetValidUntilTime sets ValidUntilTime, Tag 62
func (m Quote) SetValidUntilTime(v time.Time) {
	m.Set(field.NewValidUntilTime(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m Quote) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
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

//SetSecurityType sets SecurityType, Tag 167
func (m Quote) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetBidSpotRate sets BidSpotRate, Tag 188
func (m Quote) SetBidSpotRate(v float64) {
	m.Set(field.NewBidSpotRate(v))
}

//SetBidForwardPoints sets BidForwardPoints, Tag 189
func (m Quote) SetBidForwardPoints(v float64) {
	m.Set(field.NewBidForwardPoints(v))
}

//SetOfferSpotRate sets OfferSpotRate, Tag 190
func (m Quote) SetOfferSpotRate(v float64) {
	m.Set(field.NewOfferSpotRate(v))
}

//SetOfferForwardPoints sets OfferForwardPoints, Tag 191
func (m Quote) SetOfferForwardPoints(v float64) {
	m.Set(field.NewOfferForwardPoints(v))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m Quote) SetOrderQty2(v float64) {
	m.Set(field.NewOrderQty2(v))
}

//SetFutSettDate2 sets FutSettDate2, Tag 193
func (m Quote) SetFutSettDate2(v string) {
	m.Set(field.NewFutSettDate2(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m Quote) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m Quote) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m Quote) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m Quote) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m Quote) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m Quote) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//GetIDSource gets IDSource, Tag 22
func (m Quote) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m Quote) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
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

//GetTransactTime gets TransactTime, Tag 60
func (m Quote) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetValidUntilTime gets ValidUntilTime, Tag 62
func (m Quote) GetValidUntilTime() (f field.ValidUntilTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m Quote) GetFutSettDate() (f field.FutSettDateField, err quickfix.MessageRejectError) {
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

//GetSecurityType gets SecurityType, Tag 167
func (m Quote) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidSpotRate gets BidSpotRate, Tag 188
func (m Quote) GetBidSpotRate() (f field.BidSpotRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidForwardPoints gets BidForwardPoints, Tag 189
func (m Quote) GetBidForwardPoints() (f field.BidForwardPointsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOfferSpotRate gets OfferSpotRate, Tag 190
func (m Quote) GetOfferSpotRate() (f field.OfferSpotRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOfferForwardPoints gets OfferForwardPoints, Tag 191
func (m Quote) GetOfferForwardPoints() (f field.OfferForwardPointsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m Quote) GetOrderQty2() (f field.OrderQty2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate2 gets FutSettDate2, Tag 193
func (m Quote) GetFutSettDate2() (f field.FutSettDate2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m Quote) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m Quote) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m Quote) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m Quote) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m Quote) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m Quote) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m Quote) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m Quote) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m Quote) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m Quote) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m Quote) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasValidUntilTime returns true if ValidUntilTime is present, Tag 62
func (m Quote) HasValidUntilTime() bool {
	return m.Has(tag.ValidUntilTime)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m Quote) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
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

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m Quote) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasBidSpotRate returns true if BidSpotRate is present, Tag 188
func (m Quote) HasBidSpotRate() bool {
	return m.Has(tag.BidSpotRate)
}

//HasBidForwardPoints returns true if BidForwardPoints is present, Tag 189
func (m Quote) HasBidForwardPoints() bool {
	return m.Has(tag.BidForwardPoints)
}

//HasOfferSpotRate returns true if OfferSpotRate is present, Tag 190
func (m Quote) HasOfferSpotRate() bool {
	return m.Has(tag.OfferSpotRate)
}

//HasOfferForwardPoints returns true if OfferForwardPoints is present, Tag 191
func (m Quote) HasOfferForwardPoints() bool {
	return m.Has(tag.OfferForwardPoints)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m Quote) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasFutSettDate2 returns true if FutSettDate2 is present, Tag 193
func (m Quote) HasFutSettDate2() bool {
	return m.Has(tag.FutSettDate2)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m Quote) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m Quote) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m Quote) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m Quote) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m Quote) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m Quote) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}
