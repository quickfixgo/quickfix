//Package quote msg type = S.
package quote

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/legstipulations"
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
	"github.com/quickfixgo/quickfix/fix44/orderqtydata"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/stipulations"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fix44/yielddata"
	"time"
)

//NoQuoteQualifiers is a repeating group in Quote
type NoQuoteQualifiers struct {
	//QuoteQualifier is a non-required field for NoQuoteQualifiers.
	QuoteQualifier *string `fix:"695"`
}

//NoUnderlyings is a repeating group in Quote
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//NoLegs is a repeating group in Quote
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
	//LegStipulations Component
	LegStipulations legstipulations.Component
	//NestedParties Component
	NestedParties nestedparties.Component
	//LegPriceType is a non-required field for NoLegs.
	LegPriceType *int `fix:"686"`
	//LegBidPx is a non-required field for NoLegs.
	LegBidPx *float64 `fix:"681"`
	//LegOfferPx is a non-required field for NoLegs.
	LegOfferPx *float64 `fix:"684"`
	//LegBenchmarkCurveData Component
	LegBenchmarkCurveData legbenchmarkcurvedata.Component
}

//Message is a Quote FIX Message
type Message struct {
	FIXMsgType string `fix:"S"`
	Header     fix44.Header
	//QuoteReqID is a non-required field for Quote.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for Quote.
	QuoteID string `fix:"117"`
	//QuoteRespID is a non-required field for Quote.
	QuoteRespID *string `fix:"693"`
	//QuoteType is a non-required field for Quote.
	QuoteType *int `fix:"537"`
	//NoQuoteQualifiers is a non-required field for Quote.
	NoQuoteQualifiers []NoQuoteQualifiers `fix:"735,omitempty"`
	//QuoteResponseLevel is a non-required field for Quote.
	QuoteResponseLevel *int `fix:"301"`
	//Parties Component
	Parties parties.Component
	//TradingSessionID is a non-required field for Quote.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for Quote.
	TradingSessionSubID *string `fix:"625"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//NoUnderlyings is a non-required field for Quote.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//Side is a non-required field for Quote.
	Side *string `fix:"54"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//SettlType is a non-required field for Quote.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for Quote.
	SettlDate *string `fix:"64"`
	//SettlDate2 is a non-required field for Quote.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for Quote.
	OrderQty2 *float64 `fix:"192"`
	//Currency is a non-required field for Quote.
	Currency *string `fix:"15"`
	//Stipulations Component
	Stipulations stipulations.Component
	//Account is a non-required field for Quote.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for Quote.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for Quote.
	AccountType *int `fix:"581"`
	//NoLegs is a non-required field for Quote.
	NoLegs []NoLegs `fix:"555,omitempty"`
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
	//OrdType is a non-required field for Quote.
	OrdType *string `fix:"40"`
	//BidForwardPoints2 is a non-required field for Quote.
	BidForwardPoints2 *float64 `fix:"642"`
	//OfferForwardPoints2 is a non-required field for Quote.
	OfferForwardPoints2 *float64 `fix:"643"`
	//SettlCurrBidFxRate is a non-required field for Quote.
	SettlCurrBidFxRate *float64 `fix:"656"`
	//SettlCurrOfferFxRate is a non-required field for Quote.
	SettlCurrOfferFxRate *float64 `fix:"657"`
	//SettlCurrFxRateCalc is a non-required field for Quote.
	SettlCurrFxRateCalc *string `fix:"156"`
	//CommType is a non-required field for Quote.
	CommType *string `fix:"13"`
	//Commission is a non-required field for Quote.
	Commission *float64 `fix:"12"`
	//CustOrderCapacity is a non-required field for Quote.
	CustOrderCapacity *int `fix:"582"`
	//ExDestination is a non-required field for Quote.
	ExDestination *string `fix:"100"`
	//OrderCapacity is a non-required field for Quote.
	OrderCapacity *string `fix:"528"`
	//PriceType is a non-required field for Quote.
	PriceType *int `fix:"423"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//Text is a non-required field for Quote.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for Quote.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for Quote.
	EncodedText *string `fix:"355"`
	Trailer     fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.BeginStringFIX44, "S", r
}
