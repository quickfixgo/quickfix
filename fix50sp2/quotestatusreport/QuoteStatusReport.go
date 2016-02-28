//Package quotestatusreport msg type = AI.
package quotestatusreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/legquotstatgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/quotqualgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/targetparties"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a QuoteStatusReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AI"`
	Header     fixt11.Header
	//QuoteStatusReqID is a non-required field for QuoteStatusReport.
	QuoteStatusReqID *string `fix:"649"`
	//QuoteReqID is a non-required field for QuoteStatusReport.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a non-required field for QuoteStatusReport.
	QuoteID *string `fix:"117"`
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
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
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
	//LegQuotStatGrp Component
	LegQuotStatGrp legquotstatgrp.Component
	//QuotQualGrp Component
	QuotQualGrp quotqualgrp.Component
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
	//ExDestinationIDSource is a non-required field for QuoteStatusReport.
	ExDestinationIDSource *string `fix:"1133"`
	//QuoteCancelType is a non-required field for QuoteStatusReport.
	QuoteCancelType *int `fix:"298"`
	//QuoteMsgID is a non-required field for QuoteStatusReport.
	QuoteMsgID *string `fix:"1166"`
	//QuoteRejectReason is a non-required field for QuoteStatusReport.
	QuoteRejectReason *int `fix:"300"`
	//MinQty is a non-required field for QuoteStatusReport.
	MinQty *float64 `fix:"110"`
	//BookingType is a non-required field for QuoteStatusReport.
	BookingType *int `fix:"775"`
	//OrderCapacity is a non-required field for QuoteStatusReport.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for QuoteStatusReport.
	OrderRestrictions *string `fix:"529"`
	//TargetParties Component
	TargetParties targetparties.Component
	Trailer       fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteStatusReqID(v string)      { m.QuoteStatusReqID = &v }
func (m *Message) SetQuoteReqID(v string)            { m.QuoteReqID = &v }
func (m *Message) SetQuoteID(v string)               { m.QuoteID = &v }
func (m *Message) SetQuoteRespID(v string)           { m.QuoteRespID = &v }
func (m *Message) SetQuoteType(v int)                { m.QuoteType = &v }
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
func (m *Message) SetExpireTime(v time.Time)         { m.ExpireTime = &v }
func (m *Message) SetPrice(v float64)                { m.Price = &v }
func (m *Message) SetPriceType(v int)                { m.PriceType = &v }
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
func (m *Message) SetQuoteStatus(v int)              { m.QuoteStatus = &v }
func (m *Message) SetText(v string)                  { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)           { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)           { m.EncodedText = &v }
func (m *Message) SetExDestinationIDSource(v string) { m.ExDestinationIDSource = &v }
func (m *Message) SetQuoteCancelType(v int)          { m.QuoteCancelType = &v }
func (m *Message) SetQuoteMsgID(v string)            { m.QuoteMsgID = &v }
func (m *Message) SetQuoteRejectReason(v int)        { m.QuoteRejectReason = &v }
func (m *Message) SetMinQty(v float64)               { m.MinQty = &v }
func (m *Message) SetBookingType(v int)              { m.BookingType = &v }
func (m *Message) SetOrderCapacity(v string)         { m.OrderCapacity = &v }
func (m *Message) SetOrderRestrictions(v string)     { m.OrderRestrictions = &v }

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
	return enum.ApplVerID_FIX50SP2, "AI", r
}
