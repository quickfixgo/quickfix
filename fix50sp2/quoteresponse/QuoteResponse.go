//Package quoteresponse msg type = AJ.
package quoteresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/legquotgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/quotqualgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a QuoteResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"AJ"`
	Header     fixt11.Header
	//QuoteRespID is a required field for QuoteResponse.
	QuoteRespID string `fix:"693"`
	//QuoteID is a non-required field for QuoteResponse.
	QuoteID *string `fix:"117"`
	//QuoteRespType is a required field for QuoteResponse.
	QuoteRespType int `fix:"694"`
	//ClOrdID is a non-required field for QuoteResponse.
	ClOrdID *string `fix:"11"`
	//OrderCapacity is a non-required field for QuoteResponse.
	OrderCapacity *string `fix:"528"`
	//IOIID is a non-required field for QuoteResponse.
	IOIID *string `fix:"23"`
	//QuoteType is a non-required field for QuoteResponse.
	QuoteType *int `fix:"537"`
	//QuotQualGrp Component
	QuotQualGrp quotqualgrp.Component
	//Parties Component
	Parties parties.Component
	//TradingSessionID is a non-required field for QuoteResponse.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for QuoteResponse.
	TradingSessionSubID *string `fix:"625"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//Side is a non-required field for QuoteResponse.
	Side *string `fix:"54"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//SettlType is a non-required field for QuoteResponse.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for QuoteResponse.
	SettlDate *string `fix:"64"`
	//SettlDate2 is a non-required field for QuoteResponse.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for QuoteResponse.
	OrderQty2 *float64 `fix:"192"`
	//Currency is a non-required field for QuoteResponse.
	Currency *string `fix:"15"`
	//Stipulations Component
	Stipulations stipulations.Component
	//Account is a non-required field for QuoteResponse.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for QuoteResponse.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for QuoteResponse.
	AccountType *int `fix:"581"`
	//LegQuotGrp Component
	LegQuotGrp legquotgrp.Component
	//BidPx is a non-required field for QuoteResponse.
	BidPx *float64 `fix:"132"`
	//OfferPx is a non-required field for QuoteResponse.
	OfferPx *float64 `fix:"133"`
	//MktBidPx is a non-required field for QuoteResponse.
	MktBidPx *float64 `fix:"645"`
	//MktOfferPx is a non-required field for QuoteResponse.
	MktOfferPx *float64 `fix:"646"`
	//MinBidSize is a non-required field for QuoteResponse.
	MinBidSize *float64 `fix:"647"`
	//BidSize is a non-required field for QuoteResponse.
	BidSize *float64 `fix:"134"`
	//MinOfferSize is a non-required field for QuoteResponse.
	MinOfferSize *float64 `fix:"648"`
	//OfferSize is a non-required field for QuoteResponse.
	OfferSize *float64 `fix:"135"`
	//ValidUntilTime is a non-required field for QuoteResponse.
	ValidUntilTime *time.Time `fix:"62"`
	//BidSpotRate is a non-required field for QuoteResponse.
	BidSpotRate *float64 `fix:"188"`
	//OfferSpotRate is a non-required field for QuoteResponse.
	OfferSpotRate *float64 `fix:"190"`
	//BidForwardPoints is a non-required field for QuoteResponse.
	BidForwardPoints *float64 `fix:"189"`
	//OfferForwardPoints is a non-required field for QuoteResponse.
	OfferForwardPoints *float64 `fix:"191"`
	//MidPx is a non-required field for QuoteResponse.
	MidPx *float64 `fix:"631"`
	//BidYield is a non-required field for QuoteResponse.
	BidYield *float64 `fix:"632"`
	//MidYield is a non-required field for QuoteResponse.
	MidYield *float64 `fix:"633"`
	//OfferYield is a non-required field for QuoteResponse.
	OfferYield *float64 `fix:"634"`
	//TransactTime is a non-required field for QuoteResponse.
	TransactTime *time.Time `fix:"60"`
	//OrdType is a non-required field for QuoteResponse.
	OrdType *string `fix:"40"`
	//BidForwardPoints2 is a non-required field for QuoteResponse.
	BidForwardPoints2 *float64 `fix:"642"`
	//OfferForwardPoints2 is a non-required field for QuoteResponse.
	OfferForwardPoints2 *float64 `fix:"643"`
	//SettlCurrBidFxRate is a non-required field for QuoteResponse.
	SettlCurrBidFxRate *float64 `fix:"656"`
	//SettlCurrOfferFxRate is a non-required field for QuoteResponse.
	SettlCurrOfferFxRate *float64 `fix:"657"`
	//SettlCurrFxRateCalc is a non-required field for QuoteResponse.
	SettlCurrFxRateCalc *string `fix:"156"`
	//Commission is a non-required field for QuoteResponse.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for QuoteResponse.
	CommType *string `fix:"13"`
	//CustOrderCapacity is a non-required field for QuoteResponse.
	CustOrderCapacity *int `fix:"582"`
	//ExDestination is a non-required field for QuoteResponse.
	ExDestination *string `fix:"100"`
	//Text is a non-required field for QuoteResponse.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for QuoteResponse.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for QuoteResponse.
	EncodedText *string `fix:"355"`
	//Price is a non-required field for QuoteResponse.
	Price *float64 `fix:"44"`
	//PriceType is a non-required field for QuoteResponse.
	PriceType *int `fix:"423"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//ExDestinationIDSource is a non-required field for QuoteResponse.
	ExDestinationIDSource *string `fix:"1133"`
	//QuoteMsgID is a non-required field for QuoteResponse.
	QuoteMsgID *string `fix:"1166"`
	//PreTradeAnonymity is a non-required field for QuoteResponse.
	PreTradeAnonymity *bool `fix:"1091"`
	//MinQty is a non-required field for QuoteResponse.
	MinQty *float64 `fix:"110"`
	//OrderRestrictions is a non-required field for QuoteResponse.
	OrderRestrictions *string `fix:"529"`
	Trailer           fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "AJ", r
}
