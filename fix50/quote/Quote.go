//Package quote msg type = S.
package quote

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/legquotgrp"
	"github.com/quickfixgo/quickfix/fix50/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/quotqualgrp"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/stipulations"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a Quote FIX Message
type Message struct {
	FIXMsgType string `fix:"S"`
	Header     fixt11.Header
	//QuoteReqID is a non-required field for Quote.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for Quote.
	QuoteID string `fix:"117"`
	//QuoteRespID is a non-required field for Quote.
	QuoteRespID *string `fix:"693"`
	//QuoteType is a non-required field for Quote.
	QuoteType *int `fix:"537"`
	//QuotQualGrp Component
	QuotQualGrp quotqualgrp.Component
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
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
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
	//LegQuotGrp Component
	LegQuotGrp legquotgrp.Component
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
	//BidSwapPoints is a non-required field for Quote.
	BidSwapPoints *float64 `fix:"1065"`
	//OfferSwapPoints is a non-required field for Quote.
	OfferSwapPoints *float64 `fix:"1066"`
	//ExDestinationIDSource is a non-required field for Quote.
	ExDestinationIDSource *string `fix:"1133"`
	Trailer               fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteReqID(v string)            { m.QuoteReqID = &v }
func (m *Message) SetQuoteID(v string)               { m.QuoteID = v }
func (m *Message) SetQuoteRespID(v string)           { m.QuoteRespID = &v }
func (m *Message) SetQuoteType(v int)                { m.QuoteType = &v }
func (m *Message) SetQuoteResponseLevel(v int)       { m.QuoteResponseLevel = &v }
func (m *Message) SetTradingSessionID(v string)      { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)   { m.TradingSessionSubID = &v }
func (m *Message) SetSide(v string)                  { m.Side = &v }
func (m *Message) SetSettlType(v string)             { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)             { m.SettlDate = &v }
func (m *Message) SetSettlDate2(v string)            { m.SettlDate2 = &v }
func (m *Message) SetOrderQty2(v float64)            { m.OrderQty2 = &v }
func (m *Message) SetCurrency(v string)              { m.Currency = &v }
func (m *Message) SetAccount(v string)               { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)             { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)              { m.AccountType = &v }
func (m *Message) SetBidPx(v float64)                { m.BidPx = &v }
func (m *Message) SetOfferPx(v float64)              { m.OfferPx = &v }
func (m *Message) SetMktBidPx(v float64)             { m.MktBidPx = &v }
func (m *Message) SetMktOfferPx(v float64)           { m.MktOfferPx = &v }
func (m *Message) SetMinBidSize(v float64)           { m.MinBidSize = &v }
func (m *Message) SetBidSize(v float64)              { m.BidSize = &v }
func (m *Message) SetMinOfferSize(v float64)         { m.MinOfferSize = &v }
func (m *Message) SetOfferSize(v float64)            { m.OfferSize = &v }
func (m *Message) SetValidUntilTime(v time.Time)     { m.ValidUntilTime = &v }
func (m *Message) SetBidSpotRate(v float64)          { m.BidSpotRate = &v }
func (m *Message) SetOfferSpotRate(v float64)        { m.OfferSpotRate = &v }
func (m *Message) SetBidForwardPoints(v float64)     { m.BidForwardPoints = &v }
func (m *Message) SetOfferForwardPoints(v float64)   { m.OfferForwardPoints = &v }
func (m *Message) SetMidPx(v float64)                { m.MidPx = &v }
func (m *Message) SetBidYield(v float64)             { m.BidYield = &v }
func (m *Message) SetMidYield(v float64)             { m.MidYield = &v }
func (m *Message) SetOfferYield(v float64)           { m.OfferYield = &v }
func (m *Message) SetTransactTime(v time.Time)       { m.TransactTime = &v }
func (m *Message) SetOrdType(v string)               { m.OrdType = &v }
func (m *Message) SetBidForwardPoints2(v float64)    { m.BidForwardPoints2 = &v }
func (m *Message) SetOfferForwardPoints2(v float64)  { m.OfferForwardPoints2 = &v }
func (m *Message) SetSettlCurrBidFxRate(v float64)   { m.SettlCurrBidFxRate = &v }
func (m *Message) SetSettlCurrOfferFxRate(v float64) { m.SettlCurrOfferFxRate = &v }
func (m *Message) SetSettlCurrFxRateCalc(v string)   { m.SettlCurrFxRateCalc = &v }
func (m *Message) SetCommType(v string)              { m.CommType = &v }
func (m *Message) SetCommission(v float64)           { m.Commission = &v }
func (m *Message) SetCustOrderCapacity(v int)        { m.CustOrderCapacity = &v }
func (m *Message) SetExDestination(v string)         { m.ExDestination = &v }
func (m *Message) SetOrderCapacity(v string)         { m.OrderCapacity = &v }
func (m *Message) SetPriceType(v int)                { m.PriceType = &v }
func (m *Message) SetText(v string)                  { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)           { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)           { m.EncodedText = &v }
func (m *Message) SetBidSwapPoints(v float64)        { m.BidSwapPoints = &v }
func (m *Message) SetOfferSwapPoints(v float64)      { m.OfferSwapPoints = &v }
func (m *Message) SetExDestinationIDSource(v string) { m.ExDestinationIDSource = &v }

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
	return enum.BeginStringFIX50, "S", r
}
