//Package massquote msg type = i.
package massquote

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"time"
)

//NoQuoteSets is a repeating group in MassQuote
type NoQuoteSets struct {
	//QuoteSetID is a required field for NoQuoteSets.
	QuoteSetID string `fix:"302"`
	//UnderlyingInstrument is a non-required component for NoQuoteSets.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
	//QuoteSetValidUntilTime is a non-required field for NoQuoteSets.
	QuoteSetValidUntilTime *time.Time `fix:"367"`
	//TotNoQuoteEntries is a required field for NoQuoteSets.
	TotNoQuoteEntries int `fix:"304"`
	//LastFragment is a non-required field for NoQuoteSets.
	LastFragment *bool `fix:"893"`
	//NoQuoteEntries is a required field for NoQuoteSets.
	NoQuoteEntries []NoQuoteEntries `fix:"295"`
}

//NewNoQuoteSets returns an initialized NoQuoteSets instance
func NewNoQuoteSets(quotesetid string, totnoquoteentries int, noquoteentries []NoQuoteEntries) *NoQuoteSets {
	var m NoQuoteSets
	m.SetQuoteSetID(quotesetid)
	m.SetTotNoQuoteEntries(totnoquoteentries)
	m.SetNoQuoteEntries(noquoteentries)
	return &m
}

func (m *NoQuoteSets) SetQuoteSetID(v string) { m.QuoteSetID = v }
func (m *NoQuoteSets) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}
func (m *NoQuoteSets) SetQuoteSetValidUntilTime(v time.Time) { m.QuoteSetValidUntilTime = &v }
func (m *NoQuoteSets) SetTotNoQuoteEntries(v int)            { m.TotNoQuoteEntries = v }
func (m *NoQuoteSets) SetLastFragment(v bool)                { m.LastFragment = &v }
func (m *NoQuoteSets) SetNoQuoteEntries(v []NoQuoteEntries)  { m.NoQuoteEntries = v }

//NoQuoteEntries is a repeating group in NoQuoteSets
type NoQuoteEntries struct {
	//QuoteEntryID is a required field for NoQuoteEntries.
	QuoteEntryID string `fix:"299"`
	//Instrument is a non-required component for NoQuoteEntries.
	Instrument *instrument.Instrument
	//NoLegs is a non-required field for NoQuoteEntries.
	NoLegs []NoLegs `fix:"555,omitempty"`
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
	//MidPx is a non-required field for NoQuoteEntries.
	MidPx *float64 `fix:"631"`
	//BidYield is a non-required field for NoQuoteEntries.
	BidYield *float64 `fix:"632"`
	//MidYield is a non-required field for NoQuoteEntries.
	MidYield *float64 `fix:"633"`
	//OfferYield is a non-required field for NoQuoteEntries.
	OfferYield *float64 `fix:"634"`
	//TransactTime is a non-required field for NoQuoteEntries.
	TransactTime *time.Time `fix:"60"`
	//TradingSessionID is a non-required field for NoQuoteEntries.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoQuoteEntries.
	TradingSessionSubID *string `fix:"625"`
	//SettlDate is a non-required field for NoQuoteEntries.
	SettlDate *string `fix:"64"`
	//OrdType is a non-required field for NoQuoteEntries.
	OrdType *string `fix:"40"`
	//SettlDate2 is a non-required field for NoQuoteEntries.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoQuoteEntries.
	OrderQty2 *float64 `fix:"192"`
	//BidForwardPoints2 is a non-required field for NoQuoteEntries.
	BidForwardPoints2 *float64 `fix:"642"`
	//OfferForwardPoints2 is a non-required field for NoQuoteEntries.
	OfferForwardPoints2 *float64 `fix:"643"`
	//Currency is a non-required field for NoQuoteEntries.
	Currency *string `fix:"15"`
}

//NewNoQuoteEntries returns an initialized NoQuoteEntries instance
func NewNoQuoteEntries(quoteentryid string) *NoQuoteEntries {
	var m NoQuoteEntries
	m.SetQuoteEntryID(quoteentryid)
	return &m
}

func (m *NoQuoteEntries) SetQuoteEntryID(v string)              { m.QuoteEntryID = v }
func (m *NoQuoteEntries) SetInstrument(v instrument.Instrument) { m.Instrument = &v }
func (m *NoQuoteEntries) SetNoLegs(v []NoLegs)                  { m.NoLegs = v }
func (m *NoQuoteEntries) SetBidPx(v float64)                    { m.BidPx = &v }
func (m *NoQuoteEntries) SetOfferPx(v float64)                  { m.OfferPx = &v }
func (m *NoQuoteEntries) SetBidSize(v float64)                  { m.BidSize = &v }
func (m *NoQuoteEntries) SetOfferSize(v float64)                { m.OfferSize = &v }
func (m *NoQuoteEntries) SetValidUntilTime(v time.Time)         { m.ValidUntilTime = &v }
func (m *NoQuoteEntries) SetBidSpotRate(v float64)              { m.BidSpotRate = &v }
func (m *NoQuoteEntries) SetOfferSpotRate(v float64)            { m.OfferSpotRate = &v }
func (m *NoQuoteEntries) SetBidForwardPoints(v float64)         { m.BidForwardPoints = &v }
func (m *NoQuoteEntries) SetOfferForwardPoints(v float64)       { m.OfferForwardPoints = &v }
func (m *NoQuoteEntries) SetMidPx(v float64)                    { m.MidPx = &v }
func (m *NoQuoteEntries) SetBidYield(v float64)                 { m.BidYield = &v }
func (m *NoQuoteEntries) SetMidYield(v float64)                 { m.MidYield = &v }
func (m *NoQuoteEntries) SetOfferYield(v float64)               { m.OfferYield = &v }
func (m *NoQuoteEntries) SetTransactTime(v time.Time)           { m.TransactTime = &v }
func (m *NoQuoteEntries) SetTradingSessionID(v string)          { m.TradingSessionID = &v }
func (m *NoQuoteEntries) SetTradingSessionSubID(v string)       { m.TradingSessionSubID = &v }
func (m *NoQuoteEntries) SetSettlDate(v string)                 { m.SettlDate = &v }
func (m *NoQuoteEntries) SetOrdType(v string)                   { m.OrdType = &v }
func (m *NoQuoteEntries) SetSettlDate2(v string)                { m.SettlDate2 = &v }
func (m *NoQuoteEntries) SetOrderQty2(v float64)                { m.OrderQty2 = &v }
func (m *NoQuoteEntries) SetBidForwardPoints2(v float64)        { m.BidForwardPoints2 = &v }
func (m *NoQuoteEntries) SetOfferForwardPoints2(v float64)      { m.OfferForwardPoints2 = &v }
func (m *NoQuoteEntries) SetCurrency(v string)                  { m.Currency = &v }

//NoLegs is a repeating group in NoQuoteEntries
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
}

//NewNoLegs returns an initialized NoLegs instance
func NewNoLegs() *NoLegs {
	var m NoLegs
	return &m
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg) { m.InstrumentLeg = &v }

//Message is a MassQuote FIX Message
type Message struct {
	FIXMsgType string `fix:"i"`
	fix44.Header
	//QuoteReqID is a non-required field for MassQuote.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for MassQuote.
	QuoteID string `fix:"117"`
	//QuoteType is a non-required field for MassQuote.
	QuoteType *int `fix:"537"`
	//QuoteResponseLevel is a non-required field for MassQuote.
	QuoteResponseLevel *int `fix:"301"`
	//Parties is a non-required component for MassQuote.
	Parties *parties.Parties
	//Account is a non-required field for MassQuote.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for MassQuote.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for MassQuote.
	AccountType *int `fix:"581"`
	//DefBidSize is a non-required field for MassQuote.
	DefBidSize *float64 `fix:"293"`
	//DefOfferSize is a non-required field for MassQuote.
	DefOfferSize *float64 `fix:"294"`
	//NoQuoteSets is a required field for MassQuote.
	NoQuoteSets []NoQuoteSets `fix:"296"`
	fix44.Trailer
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
func (m *Message) SetQuoteType(v int)             { m.QuoteType = &v }
func (m *Message) SetQuoteResponseLevel(v int)    { m.QuoteResponseLevel = &v }
func (m *Message) SetParties(v parties.Parties)   { m.Parties = &v }
func (m *Message) SetAccount(v string)            { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)          { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)           { m.AccountType = &v }
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
	return enum.BeginStringFIX44, "i", r
}
