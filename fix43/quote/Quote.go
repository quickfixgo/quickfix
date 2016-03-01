//Package quote msg type = S.
package quote

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"time"
)

//Message is a Quote FIX Message
type Message struct {
	FIXMsgType string `fix:"S"`
	Header     fix43.Header
	//QuoteReqID is a non-required field for Quote.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for Quote.
	QuoteID string `fix:"117"`
	//QuoteType is a non-required field for Quote.
	QuoteType *int `fix:"537"`
	//QuoteResponseLevel is a non-required field for Quote.
	QuoteResponseLevel *int `fix:"301"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for Quote.
	Account *string `fix:"1"`
	//AccountType is a non-required field for Quote.
	AccountType *int `fix:"581"`
	//TradingSessionID is a non-required field for Quote.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for Quote.
	TradingSessionSubID *string `fix:"625"`
	//Instrument Component
	Instrument instrument.Component
	//BidPx is a non-required field for Quote.
	BidPx *float64 `fix:"132"`
	//OfferPx is a non-required field for Quote.
	OfferPx *float64 `fix:"133"`
	//MktBidPx is a non-required field for Quote.
	MktBidPx *float64 `fix:"645"`
	//MktOfferPx is a non-required field for Quote.
	MktOfferPx *float64 `fix:"646"`
	//MinBidSize is a non-required field for Quote.
	MinBidSize *float64 `fix:"647"`
	//BidSize is a non-required field for Quote.
	BidSize *float64 `fix:"134"`
	//MinOfferSize is a non-required field for Quote.
	MinOfferSize *float64 `fix:"648"`
	//OfferSize is a non-required field for Quote.
	OfferSize *float64 `fix:"135"`
	//ValidUntilTime is a non-required field for Quote.
	ValidUntilTime *time.Time `fix:"62"`
	//BidSpotRate is a non-required field for Quote.
	BidSpotRate *float64 `fix:"188"`
	//OfferSpotRate is a non-required field for Quote.
	OfferSpotRate *float64 `fix:"190"`
	//BidForwardPoints is a non-required field for Quote.
	BidForwardPoints *float64 `fix:"189"`
	//OfferForwardPoints is a non-required field for Quote.
	OfferForwardPoints *float64 `fix:"191"`
	//MidPx is a non-required field for Quote.
	MidPx *float64 `fix:"631"`
	//BidYield is a non-required field for Quote.
	BidYield *float64 `fix:"632"`
	//MidYield is a non-required field for Quote.
	MidYield *float64 `fix:"633"`
	//OfferYield is a non-required field for Quote.
	OfferYield *float64 `fix:"634"`
	//TransactTime is a non-required field for Quote.
	TransactTime *time.Time `fix:"60"`
	//SettlmntTyp is a non-required field for Quote.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for Quote.
	FutSettDate *string `fix:"64"`
	//OrdType is a non-required field for Quote.
	OrdType *string `fix:"40"`
	//FutSettDate2 is a non-required field for Quote.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for Quote.
	OrderQty2 *float64 `fix:"192"`
	//BidForwardPoints2 is a non-required field for Quote.
	BidForwardPoints2 *float64 `fix:"642"`
	//OfferForwardPoints2 is a non-required field for Quote.
	OfferForwardPoints2 *float64 `fix:"643"`
	//Currency is a non-required field for Quote.
	Currency *string `fix:"15"`
	//SettlCurrBidFxRate is a non-required field for Quote.
	SettlCurrBidFxRate *float64 `fix:"656"`
	//SettlCurrOfferFxRate is a non-required field for Quote.
	SettlCurrOfferFxRate *float64 `fix:"657"`
	//SettlCurrFxRateCalc is a non-required field for Quote.
	SettlCurrFxRateCalc *string `fix:"156"`
	//Commission is a non-required field for Quote.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for Quote.
	CommType *string `fix:"13"`
	//CustOrderCapacity is a non-required field for Quote.
	CustOrderCapacity *int `fix:"582"`
	//ExDestination is a non-required field for Quote.
	ExDestination *string `fix:"100"`
	//Text is a non-required field for Quote.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for Quote.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for Quote.
	EncodedText *string `fix:"355"`
	Trailer     fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.BeginStringFIX43, "S", r
}
