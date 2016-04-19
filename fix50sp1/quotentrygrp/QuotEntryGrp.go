package quotentrygrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"time"
)

//New returns an initialized QuotEntryGrp instance
func New(noquoteentries []NoQuoteEntries) *QuotEntryGrp {
	var m QuotEntryGrp
	m.SetNoQuoteEntries(noquoteentries)
	return &m
}

//NoQuoteEntries is a repeating group in QuotEntryGrp
type NoQuoteEntries struct {
	//QuoteEntryID is a required field for NoQuoteEntries.
	QuoteEntryID string `fix:"299"`
	//Instrument is a non-required component for NoQuoteEntries.
	Instrument *instrument.Instrument
	//InstrmtLegGrp is a non-required component for NoQuoteEntries.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
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

func (m *NoQuoteEntries) SetQuoteEntryID(v string)                       { m.QuoteEntryID = v }
func (m *NoQuoteEntries) SetInstrument(v instrument.Instrument)          { m.Instrument = &v }
func (m *NoQuoteEntries) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp) { m.InstrmtLegGrp = &v }
func (m *NoQuoteEntries) SetBidPx(v float64)                             { m.BidPx = &v }
func (m *NoQuoteEntries) SetOfferPx(v float64)                           { m.OfferPx = &v }
func (m *NoQuoteEntries) SetBidSize(v float64)                           { m.BidSize = &v }
func (m *NoQuoteEntries) SetOfferSize(v float64)                         { m.OfferSize = &v }
func (m *NoQuoteEntries) SetValidUntilTime(v time.Time)                  { m.ValidUntilTime = &v }
func (m *NoQuoteEntries) SetBidSpotRate(v float64)                       { m.BidSpotRate = &v }
func (m *NoQuoteEntries) SetOfferSpotRate(v float64)                     { m.OfferSpotRate = &v }
func (m *NoQuoteEntries) SetBidForwardPoints(v float64)                  { m.BidForwardPoints = &v }
func (m *NoQuoteEntries) SetOfferForwardPoints(v float64)                { m.OfferForwardPoints = &v }
func (m *NoQuoteEntries) SetMidPx(v float64)                             { m.MidPx = &v }
func (m *NoQuoteEntries) SetBidYield(v float64)                          { m.BidYield = &v }
func (m *NoQuoteEntries) SetMidYield(v float64)                          { m.MidYield = &v }
func (m *NoQuoteEntries) SetOfferYield(v float64)                        { m.OfferYield = &v }
func (m *NoQuoteEntries) SetTransactTime(v time.Time)                    { m.TransactTime = &v }
func (m *NoQuoteEntries) SetTradingSessionID(v string)                   { m.TradingSessionID = &v }
func (m *NoQuoteEntries) SetTradingSessionSubID(v string)                { m.TradingSessionSubID = &v }
func (m *NoQuoteEntries) SetSettlDate(v string)                          { m.SettlDate = &v }
func (m *NoQuoteEntries) SetOrdType(v string)                            { m.OrdType = &v }
func (m *NoQuoteEntries) SetSettlDate2(v string)                         { m.SettlDate2 = &v }
func (m *NoQuoteEntries) SetOrderQty2(v float64)                         { m.OrderQty2 = &v }
func (m *NoQuoteEntries) SetBidForwardPoints2(v float64)                 { m.BidForwardPoints2 = &v }
func (m *NoQuoteEntries) SetOfferForwardPoints2(v float64)               { m.OfferForwardPoints2 = &v }
func (m *NoQuoteEntries) SetCurrency(v string)                           { m.Currency = &v }

//QuotEntryGrp is a fix50sp1 Component
type QuotEntryGrp struct {
	//NoQuoteEntries is a required field for QuotEntryGrp.
	NoQuoteEntries []NoQuoteEntries `fix:"295"`
}

func (m *QuotEntryGrp) SetNoQuoteEntries(v []NoQuoteEntries) { m.NoQuoteEntries = v }
