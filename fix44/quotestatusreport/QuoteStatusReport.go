//Package quotestatusreport msg type = AI.
package quotestatusreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
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

//NoUnderlyings is a repeating group in QuoteStatusReport
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//NoLegs is a repeating group in QuoteStatusReport
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
}

//NoQuoteQualifiers is a repeating group in QuoteStatusReport
type NoQuoteQualifiers struct {
	//QuoteQualifier is a non-required field for NoQuoteQualifiers.
	QuoteQualifier *string `fix:"695"`
}

//Message is a QuoteStatusReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AI"`
	Header     fix44.Header
	//QuoteStatusReqID is a non-required field for QuoteStatusReport.
	QuoteStatusReqID *string `fix:"649"`
	//QuoteReqID is a non-required field for QuoteStatusReport.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for QuoteStatusReport.
	QuoteID string `fix:"117"`
	//QuoteRespID is a non-required field for QuoteStatusReport.
	QuoteRespID *string `fix:"693"`
	//QuoteType is a non-required field for QuoteStatusReport.
	QuoteType *int `fix:"537"`
	//Parties Component
	Parties parties.Component
	//TradingSessionID is a non-required field for QuoteStatusReport.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for QuoteStatusReport.
	TradingSessionSubID *string `fix:"625"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//NoUnderlyings is a non-required field for QuoteStatusReport.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//Side is a non-required field for QuoteStatusReport.
	Side *string `fix:"54"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//SettlType is a non-required field for QuoteStatusReport.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for QuoteStatusReport.
	SettlDate *string `fix:"64"`
	//SettlDate2 is a non-required field for QuoteStatusReport.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for QuoteStatusReport.
	OrderQty2 *float64 `fix:"192"`
	//Currency is a non-required field for QuoteStatusReport.
	Currency *string `fix:"15"`
	//Stipulations Component
	Stipulations stipulations.Component
	//Account is a non-required field for QuoteStatusReport.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for QuoteStatusReport.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for QuoteStatusReport.
	AccountType *int `fix:"581"`
	//NoLegs is a non-required field for QuoteStatusReport.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoQuoteQualifiers is a non-required field for QuoteStatusReport.
	NoQuoteQualifiers []NoQuoteQualifiers `fix:"735,omitempty"`
	//ExpireTime is a non-required field for QuoteStatusReport.
	ExpireTime *time.Time `fix:"126"`
	//Price is a non-required field for QuoteStatusReport.
	Price *float64 `fix:"44"`
	//PriceType is a non-required field for QuoteStatusReport.
	PriceType *int `fix:"423"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//BidPx is a non-required field for QuoteStatusReport.
	BidPx *float64 `fix:"132"`
	//OfferPx is a non-required field for QuoteStatusReport.
	OfferPx *float64 `fix:"133"`
	//MktBidPx is a non-required field for QuoteStatusReport.
	MktBidPx *float64 `fix:"645"`
	//MktOfferPx is a non-required field for QuoteStatusReport.
	MktOfferPx *float64 `fix:"646"`
	//MinBidSize is a non-required field for QuoteStatusReport.
	MinBidSize *float64 `fix:"647"`
	//BidSize is a non-required field for QuoteStatusReport.
	BidSize *float64 `fix:"134"`
	//MinOfferSize is a non-required field for QuoteStatusReport.
	MinOfferSize *float64 `fix:"648"`
	//OfferSize is a non-required field for QuoteStatusReport.
	OfferSize *float64 `fix:"135"`
	//ValidUntilTime is a non-required field for QuoteStatusReport.
	ValidUntilTime *time.Time `fix:"62"`
	//BidSpotRate is a non-required field for QuoteStatusReport.
	BidSpotRate *float64 `fix:"188"`
	//OfferSpotRate is a non-required field for QuoteStatusReport.
	OfferSpotRate *float64 `fix:"190"`
	//BidForwardPoints is a non-required field for QuoteStatusReport.
	BidForwardPoints *float64 `fix:"189"`
	//OfferForwardPoints is a non-required field for QuoteStatusReport.
	OfferForwardPoints *float64 `fix:"191"`
	//MidPx is a non-required field for QuoteStatusReport.
	MidPx *float64 `fix:"631"`
	//BidYield is a non-required field for QuoteStatusReport.
	BidYield *float64 `fix:"632"`
	//MidYield is a non-required field for QuoteStatusReport.
	MidYield *float64 `fix:"633"`
	//OfferYield is a non-required field for QuoteStatusReport.
	OfferYield *float64 `fix:"634"`
	//TransactTime is a non-required field for QuoteStatusReport.
	TransactTime *time.Time `fix:"60"`
	//OrdType is a non-required field for QuoteStatusReport.
	OrdType *string `fix:"40"`
	//BidForwardPoints2 is a non-required field for QuoteStatusReport.
	BidForwardPoints2 *float64 `fix:"642"`
	//OfferForwardPoints2 is a non-required field for QuoteStatusReport.
	OfferForwardPoints2 *float64 `fix:"643"`
	//SettlCurrBidFxRate is a non-required field for QuoteStatusReport.
	SettlCurrBidFxRate *float64 `fix:"656"`
	//SettlCurrOfferFxRate is a non-required field for QuoteStatusReport.
	SettlCurrOfferFxRate *float64 `fix:"657"`
	//SettlCurrFxRateCalc is a non-required field for QuoteStatusReport.
	SettlCurrFxRateCalc *string `fix:"156"`
	//CommType is a non-required field for QuoteStatusReport.
	CommType *string `fix:"13"`
	//Commission is a non-required field for QuoteStatusReport.
	Commission *float64 `fix:"12"`
	//CustOrderCapacity is a non-required field for QuoteStatusReport.
	CustOrderCapacity *int `fix:"582"`
	//ExDestination is a non-required field for QuoteStatusReport.
	ExDestination *string `fix:"100"`
	//QuoteStatus is a non-required field for QuoteStatusReport.
	QuoteStatus *int `fix:"297"`
	//Text is a non-required field for QuoteStatusReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for QuoteStatusReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for QuoteStatusReport.
	EncodedText *string `fix:"355"`
	Trailer     fix44.Trailer
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
	return enum.BeginStringFIX44, "AI", r
}
