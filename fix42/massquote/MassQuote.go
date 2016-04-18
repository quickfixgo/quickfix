//Package massquote msg type = i.
package massquote

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//NoQuoteSets is a repeating group in MassQuote
type NoQuoteSets struct {
	//QuoteSetID is a required field for NoQuoteSets.
	QuoteSetID string `fix:"302"`
	//UnderlyingSymbol is a required field for NoQuoteSets.
	UnderlyingSymbol string `fix:"311"`
	//UnderlyingSymbolSfx is a non-required field for NoQuoteSets.
	UnderlyingSymbolSfx *string `fix:"312"`
	//UnderlyingSecurityID is a non-required field for NoQuoteSets.
	UnderlyingSecurityID *string `fix:"309"`
	//UnderlyingIDSource is a non-required field for NoQuoteSets.
	UnderlyingIDSource *string `fix:"305"`
	//UnderlyingSecurityType is a non-required field for NoQuoteSets.
	UnderlyingSecurityType *string `fix:"310"`
	//UnderlyingMaturityMonthYear is a non-required field for NoQuoteSets.
	UnderlyingMaturityMonthYear *string `fix:"313"`
	//UnderlyingMaturityDay is a non-required field for NoQuoteSets.
	UnderlyingMaturityDay *int `fix:"314"`
	//UnderlyingPutOrCall is a non-required field for NoQuoteSets.
	UnderlyingPutOrCall *int `fix:"315"`
	//UnderlyingStrikePrice is a non-required field for NoQuoteSets.
	UnderlyingStrikePrice *float64 `fix:"316"`
	//UnderlyingOptAttribute is a non-required field for NoQuoteSets.
	UnderlyingOptAttribute *string `fix:"317"`
	//UnderlyingContractMultiplier is a non-required field for NoQuoteSets.
	UnderlyingContractMultiplier *float64 `fix:"436"`
	//UnderlyingCouponRate is a non-required field for NoQuoteSets.
	UnderlyingCouponRate *float64 `fix:"435"`
	//UnderlyingSecurityExchange is a non-required field for NoQuoteSets.
	UnderlyingSecurityExchange *string `fix:"308"`
	//UnderlyingIssuer is a non-required field for NoQuoteSets.
	UnderlyingIssuer *string `fix:"306"`
	//EncodedUnderlyingIssuerLen is a non-required field for NoQuoteSets.
	EncodedUnderlyingIssuerLen *int `fix:"362"`
	//EncodedUnderlyingIssuer is a non-required field for NoQuoteSets.
	EncodedUnderlyingIssuer *string `fix:"363"`
	//UnderlyingSecurityDesc is a non-required field for NoQuoteSets.
	UnderlyingSecurityDesc *string `fix:"307"`
	//EncodedUnderlyingSecurityDescLen is a non-required field for NoQuoteSets.
	EncodedUnderlyingSecurityDescLen *int `fix:"364"`
	//EncodedUnderlyingSecurityDesc is a non-required field for NoQuoteSets.
	EncodedUnderlyingSecurityDesc *string `fix:"365"`
	//QuoteSetValidUntilTime is a non-required field for NoQuoteSets.
	QuoteSetValidUntilTime *time.Time `fix:"367"`
	//TotQuoteEntries is a required field for NoQuoteSets.
	TotQuoteEntries int `fix:"304"`
	//NoQuoteEntries is a required field for NoQuoteSets.
	NoQuoteEntries []NoQuoteEntries `fix:"295"`
}

//NewNoQuoteSets returns an initialized NoQuoteSets instance
func NewNoQuoteSets(quotesetid string, underlyingsymbol string, totquoteentries int, noquoteentries []NoQuoteEntries) *NoQuoteSets {
	var m NoQuoteSets
	m.SetQuoteSetID(quotesetid)
	m.SetUnderlyingSymbol(underlyingsymbol)
	m.SetTotQuoteEntries(totquoteentries)
	m.SetNoQuoteEntries(noquoteentries)
	return &m
}

func (m *NoQuoteSets) SetQuoteSetID(v string)                    { m.QuoteSetID = v }
func (m *NoQuoteSets) SetUnderlyingSymbol(v string)              { m.UnderlyingSymbol = v }
func (m *NoQuoteSets) SetUnderlyingSymbolSfx(v string)           { m.UnderlyingSymbolSfx = &v }
func (m *NoQuoteSets) SetUnderlyingSecurityID(v string)          { m.UnderlyingSecurityID = &v }
func (m *NoQuoteSets) SetUnderlyingIDSource(v string)            { m.UnderlyingIDSource = &v }
func (m *NoQuoteSets) SetUnderlyingSecurityType(v string)        { m.UnderlyingSecurityType = &v }
func (m *NoQuoteSets) SetUnderlyingMaturityMonthYear(v string)   { m.UnderlyingMaturityMonthYear = &v }
func (m *NoQuoteSets) SetUnderlyingMaturityDay(v int)            { m.UnderlyingMaturityDay = &v }
func (m *NoQuoteSets) SetUnderlyingPutOrCall(v int)              { m.UnderlyingPutOrCall = &v }
func (m *NoQuoteSets) SetUnderlyingStrikePrice(v float64)        { m.UnderlyingStrikePrice = &v }
func (m *NoQuoteSets) SetUnderlyingOptAttribute(v string)        { m.UnderlyingOptAttribute = &v }
func (m *NoQuoteSets) SetUnderlyingContractMultiplier(v float64) { m.UnderlyingContractMultiplier = &v }
func (m *NoQuoteSets) SetUnderlyingCouponRate(v float64)         { m.UnderlyingCouponRate = &v }
func (m *NoQuoteSets) SetUnderlyingSecurityExchange(v string)    { m.UnderlyingSecurityExchange = &v }
func (m *NoQuoteSets) SetUnderlyingIssuer(v string)              { m.UnderlyingIssuer = &v }
func (m *NoQuoteSets) SetEncodedUnderlyingIssuerLen(v int)       { m.EncodedUnderlyingIssuerLen = &v }
func (m *NoQuoteSets) SetEncodedUnderlyingIssuer(v string)       { m.EncodedUnderlyingIssuer = &v }
func (m *NoQuoteSets) SetUnderlyingSecurityDesc(v string)        { m.UnderlyingSecurityDesc = &v }
func (m *NoQuoteSets) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.EncodedUnderlyingSecurityDescLen = &v
}
func (m *NoQuoteSets) SetEncodedUnderlyingSecurityDesc(v string) { m.EncodedUnderlyingSecurityDesc = &v }
func (m *NoQuoteSets) SetQuoteSetValidUntilTime(v time.Time)     { m.QuoteSetValidUntilTime = &v }
func (m *NoQuoteSets) SetTotQuoteEntries(v int)                  { m.TotQuoteEntries = v }
func (m *NoQuoteSets) SetNoQuoteEntries(v []NoQuoteEntries)      { m.NoQuoteEntries = v }

//NoQuoteEntries is a repeating group in NoQuoteSets
type NoQuoteEntries struct {
	//QuoteEntryID is a required field for NoQuoteEntries.
	QuoteEntryID string `fix:"299"`
	//Symbol is a non-required field for NoQuoteEntries.
	Symbol *string `fix:"55"`
	//SymbolSfx is a non-required field for NoQuoteEntries.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for NoQuoteEntries.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for NoQuoteEntries.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for NoQuoteEntries.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for NoQuoteEntries.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for NoQuoteEntries.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for NoQuoteEntries.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for NoQuoteEntries.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for NoQuoteEntries.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for NoQuoteEntries.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for NoQuoteEntries.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for NoQuoteEntries.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for NoQuoteEntries.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for NoQuoteEntries.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for NoQuoteEntries.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for NoQuoteEntries.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for NoQuoteEntries.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for NoQuoteEntries.
	EncodedSecurityDesc *string `fix:"351"`
	//BidPx is a non-required field for NoQuoteEntries.
	BidPx *float64 `fix:"132"`
	//OfferPx is a non-required field for NoQuoteEntries.
	OfferPx *float64 `fix:"133"`
	//BidSize is a non-required field for NoQuoteEntries.
	BidSize *float64 `fix:"134"`
	//OfferSize is a non-required field for NoQuoteEntries.
	OfferSize *float64 `fix:"135"`
	//ValidUntilTime is a non-required field for NoQuoteEntries.
	ValidUntilTime *time.Time `fix:"62"`
	//BidSpotRate is a non-required field for NoQuoteEntries.
	BidSpotRate *float64 `fix:"188"`
	//OfferSpotRate is a non-required field for NoQuoteEntries.
	OfferSpotRate *float64 `fix:"190"`
	//BidForwardPoints is a non-required field for NoQuoteEntries.
	BidForwardPoints *float64 `fix:"189"`
	//OfferForwardPoints is a non-required field for NoQuoteEntries.
	OfferForwardPoints *float64 `fix:"191"`
	//TransactTime is a non-required field for NoQuoteEntries.
	TransactTime *time.Time `fix:"60"`
	//TradingSessionID is a non-required field for NoQuoteEntries.
	TradingSessionID *string `fix:"336"`
	//FutSettDate is a non-required field for NoQuoteEntries.
	FutSettDate *string `fix:"64"`
	//OrdType is a non-required field for NoQuoteEntries.
	OrdType *string `fix:"40"`
	//FutSettDate2 is a non-required field for NoQuoteEntries.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoQuoteEntries.
	OrderQty2 *float64 `fix:"192"`
	//Currency is a non-required field for NoQuoteEntries.
	Currency *string `fix:"15"`
}

//NewNoQuoteEntries returns an initialized NoQuoteEntries instance
func NewNoQuoteEntries(quoteentryid string) *NoQuoteEntries {
	var m NoQuoteEntries
	m.SetQuoteEntryID(quoteentryid)
	return &m
}

func (m *NoQuoteEntries) SetQuoteEntryID(v string)        { m.QuoteEntryID = v }
func (m *NoQuoteEntries) SetSymbol(v string)              { m.Symbol = &v }
func (m *NoQuoteEntries) SetSymbolSfx(v string)           { m.SymbolSfx = &v }
func (m *NoQuoteEntries) SetSecurityID(v string)          { m.SecurityID = &v }
func (m *NoQuoteEntries) SetIDSource(v string)            { m.IDSource = &v }
func (m *NoQuoteEntries) SetSecurityType(v string)        { m.SecurityType = &v }
func (m *NoQuoteEntries) SetMaturityMonthYear(v string)   { m.MaturityMonthYear = &v }
func (m *NoQuoteEntries) SetMaturityDay(v int)            { m.MaturityDay = &v }
func (m *NoQuoteEntries) SetPutOrCall(v int)              { m.PutOrCall = &v }
func (m *NoQuoteEntries) SetStrikePrice(v float64)        { m.StrikePrice = &v }
func (m *NoQuoteEntries) SetOptAttribute(v string)        { m.OptAttribute = &v }
func (m *NoQuoteEntries) SetContractMultiplier(v float64) { m.ContractMultiplier = &v }
func (m *NoQuoteEntries) SetCouponRate(v float64)         { m.CouponRate = &v }
func (m *NoQuoteEntries) SetSecurityExchange(v string)    { m.SecurityExchange = &v }
func (m *NoQuoteEntries) SetIssuer(v string)              { m.Issuer = &v }
func (m *NoQuoteEntries) SetEncodedIssuerLen(v int)       { m.EncodedIssuerLen = &v }
func (m *NoQuoteEntries) SetEncodedIssuer(v string)       { m.EncodedIssuer = &v }
func (m *NoQuoteEntries) SetSecurityDesc(v string)        { m.SecurityDesc = &v }
func (m *NoQuoteEntries) SetEncodedSecurityDescLen(v int) { m.EncodedSecurityDescLen = &v }
func (m *NoQuoteEntries) SetEncodedSecurityDesc(v string) { m.EncodedSecurityDesc = &v }
func (m *NoQuoteEntries) SetBidPx(v float64)              { m.BidPx = &v }
func (m *NoQuoteEntries) SetOfferPx(v float64)            { m.OfferPx = &v }
func (m *NoQuoteEntries) SetBidSize(v float64)            { m.BidSize = &v }
func (m *NoQuoteEntries) SetOfferSize(v float64)          { m.OfferSize = &v }
func (m *NoQuoteEntries) SetValidUntilTime(v time.Time)   { m.ValidUntilTime = &v }
func (m *NoQuoteEntries) SetBidSpotRate(v float64)        { m.BidSpotRate = &v }
func (m *NoQuoteEntries) SetOfferSpotRate(v float64)      { m.OfferSpotRate = &v }
func (m *NoQuoteEntries) SetBidForwardPoints(v float64)   { m.BidForwardPoints = &v }
func (m *NoQuoteEntries) SetOfferForwardPoints(v float64) { m.OfferForwardPoints = &v }
func (m *NoQuoteEntries) SetTransactTime(v time.Time)     { m.TransactTime = &v }
func (m *NoQuoteEntries) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoQuoteEntries) SetFutSettDate(v string)         { m.FutSettDate = &v }
func (m *NoQuoteEntries) SetOrdType(v string)             { m.OrdType = &v }
func (m *NoQuoteEntries) SetFutSettDate2(v string)        { m.FutSettDate2 = &v }
func (m *NoQuoteEntries) SetOrderQty2(v float64)          { m.OrderQty2 = &v }
func (m *NoQuoteEntries) SetCurrency(v string)            { m.Currency = &v }

//Message is a MassQuote FIX Message
type Message struct {
	FIXMsgType string `fix:"i"`
	fix42.Header
	//QuoteReqID is a non-required field for MassQuote.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for MassQuote.
	QuoteID string `fix:"117"`
	//QuoteResponseLevel is a non-required field for MassQuote.
	QuoteResponseLevel *int `fix:"301"`
	//DefBidSize is a non-required field for MassQuote.
	DefBidSize *float64 `fix:"293"`
	//DefOfferSize is a non-required field for MassQuote.
	DefOfferSize *float64 `fix:"294"`
	//NoQuoteSets is a required field for MassQuote.
	NoQuoteSets []NoQuoteSets `fix:"296"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized MassQuote instance
func New(quoteid string, noquotesets []NoQuoteSets) *Message {
	var m Message
	m.SetQuoteID(quoteid)
	m.SetNoQuoteSets(noquotesets)
	return &m
}

func (m *Message) SetQuoteReqID(v string)         { m.QuoteReqID = &v }
func (m *Message) SetQuoteID(v string)            { m.QuoteID = v }
func (m *Message) SetQuoteResponseLevel(v int)    { m.QuoteResponseLevel = &v }
func (m *Message) SetDefBidSize(v float64)        { m.DefBidSize = &v }
func (m *Message) SetDefOfferSize(v float64)      { m.DefOfferSize = &v }
func (m *Message) SetNoQuoteSets(v []NoQuoteSets) { m.NoQuoteSets = v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX42, "i", r
}
