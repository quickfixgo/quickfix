//Package massquoteacknowledgement msg type = b.
package massquoteacknowledgement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"github.com/quickfixgo/quickfix/fix43/underlyinginstrument"
	"time"
)

//NoQuoteSets is a repeating group in MassQuoteAcknowledgement
type NoQuoteSets struct {
	//QuoteSetID is a non-required field for NoQuoteSets.
	QuoteSetID *string `fix:"302"`
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
	//TotQuoteEntries is a non-required field for NoQuoteSets.
	TotQuoteEntries *int `fix:"304"`
	//NoQuoteEntries is a non-required field for NoQuoteSets.
	NoQuoteEntries []NoQuoteEntries `fix:"295,omitempty"`
}

func (m *NoQuoteSets) SetQuoteSetID(v string)               { m.QuoteSetID = &v }
func (m *NoQuoteSets) SetTotQuoteEntries(v int)             { m.TotQuoteEntries = &v }
func (m *NoQuoteSets) SetNoQuoteEntries(v []NoQuoteEntries) { m.NoQuoteEntries = v }

//NoQuoteEntries is a repeating group in NoQuoteSets
type NoQuoteEntries struct {
	//QuoteEntryID is a non-required field for NoQuoteEntries.
	QuoteEntryID *string `fix:"299"`
	//Instrument Component
	instrument.Instrument
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
	//FutSettDate is a non-required field for NoQuoteEntries.
	FutSettDate *string `fix:"64"`
	//OrdType is a non-required field for NoQuoteEntries.
	OrdType *string `fix:"40"`
	//FutSettDate2 is a non-required field for NoQuoteEntries.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoQuoteEntries.
	OrderQty2 *float64 `fix:"192"`
	//BidForwardPoints2 is a non-required field for NoQuoteEntries.
	BidForwardPoints2 *float64 `fix:"642"`
	//OfferForwardPoints2 is a non-required field for NoQuoteEntries.
	OfferForwardPoints2 *float64 `fix:"643"`
	//Currency is a non-required field for NoQuoteEntries.
	Currency *string `fix:"15"`
	//QuoteEntryRejectReason is a non-required field for NoQuoteEntries.
	QuoteEntryRejectReason *int `fix:"368"`
}

func (m *NoQuoteEntries) SetQuoteEntryID(v string)         { m.QuoteEntryID = &v }
func (m *NoQuoteEntries) SetBidPx(v float64)               { m.BidPx = &v }
func (m *NoQuoteEntries) SetOfferPx(v float64)             { m.OfferPx = &v }
func (m *NoQuoteEntries) SetBidSize(v float64)             { m.BidSize = &v }
func (m *NoQuoteEntries) SetOfferSize(v float64)           { m.OfferSize = &v }
func (m *NoQuoteEntries) SetValidUntilTime(v time.Time)    { m.ValidUntilTime = &v }
func (m *NoQuoteEntries) SetBidSpotRate(v float64)         { m.BidSpotRate = &v }
func (m *NoQuoteEntries) SetOfferSpotRate(v float64)       { m.OfferSpotRate = &v }
func (m *NoQuoteEntries) SetBidForwardPoints(v float64)    { m.BidForwardPoints = &v }
func (m *NoQuoteEntries) SetOfferForwardPoints(v float64)  { m.OfferForwardPoints = &v }
func (m *NoQuoteEntries) SetMidPx(v float64)               { m.MidPx = &v }
func (m *NoQuoteEntries) SetBidYield(v float64)            { m.BidYield = &v }
func (m *NoQuoteEntries) SetMidYield(v float64)            { m.MidYield = &v }
func (m *NoQuoteEntries) SetOfferYield(v float64)          { m.OfferYield = &v }
func (m *NoQuoteEntries) SetTransactTime(v time.Time)      { m.TransactTime = &v }
func (m *NoQuoteEntries) SetTradingSessionID(v string)     { m.TradingSessionID = &v }
func (m *NoQuoteEntries) SetTradingSessionSubID(v string)  { m.TradingSessionSubID = &v }
func (m *NoQuoteEntries) SetFutSettDate(v string)          { m.FutSettDate = &v }
func (m *NoQuoteEntries) SetOrdType(v string)              { m.OrdType = &v }
func (m *NoQuoteEntries) SetFutSettDate2(v string)         { m.FutSettDate2 = &v }
func (m *NoQuoteEntries) SetOrderQty2(v float64)           { m.OrderQty2 = &v }
func (m *NoQuoteEntries) SetBidForwardPoints2(v float64)   { m.BidForwardPoints2 = &v }
func (m *NoQuoteEntries) SetOfferForwardPoints2(v float64) { m.OfferForwardPoints2 = &v }
func (m *NoQuoteEntries) SetCurrency(v string)             { m.Currency = &v }
func (m *NoQuoteEntries) SetQuoteEntryRejectReason(v int)  { m.QuoteEntryRejectReason = &v }

//Message is a MassQuoteAcknowledgement FIX Message
type Message struct {
	FIXMsgType string `fix:"b"`
	fix43.Header
	//QuoteReqID is a non-required field for MassQuoteAcknowledgement.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a non-required field for MassQuoteAcknowledgement.
	QuoteID *string `fix:"117"`
	//QuoteStatus is a required field for MassQuoteAcknowledgement.
	QuoteStatus int `fix:"297"`
	//QuoteRejectReason is a non-required field for MassQuoteAcknowledgement.
	QuoteRejectReason *int `fix:"300"`
	//QuoteResponseLevel is a non-required field for MassQuoteAcknowledgement.
	QuoteResponseLevel *int `fix:"301"`
	//QuoteType is a non-required field for MassQuoteAcknowledgement.
	QuoteType *int `fix:"537"`
	//Parties Component
	parties.Parties
	//Account is a non-required field for MassQuoteAcknowledgement.
	Account *string `fix:"1"`
	//AccountType is a non-required field for MassQuoteAcknowledgement.
	AccountType *int `fix:"581"`
	//Text is a non-required field for MassQuoteAcknowledgement.
	Text *string `fix:"58"`
	//NoQuoteSets is a non-required field for MassQuoteAcknowledgement.
	NoQuoteSets []NoQuoteSets `fix:"296,omitempty"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteReqID(v string)         { m.QuoteReqID = &v }
func (m *Message) SetQuoteID(v string)            { m.QuoteID = &v }
func (m *Message) SetQuoteStatus(v int)           { m.QuoteStatus = v }
func (m *Message) SetQuoteRejectReason(v int)     { m.QuoteRejectReason = &v }
func (m *Message) SetQuoteResponseLevel(v int)    { m.QuoteResponseLevel = &v }
func (m *Message) SetQuoteType(v int)             { m.QuoteType = &v }
func (m *Message) SetAccount(v string)            { m.Account = &v }
func (m *Message) SetAccountType(v int)           { m.AccountType = &v }
func (m *Message) SetText(v string)               { m.Text = &v }
func (m *Message) SetNoQuoteSets(v []NoQuoteSets) { m.NoQuoteSets = v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX43, "b", r
}
