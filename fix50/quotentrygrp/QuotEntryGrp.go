package quotentrygrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"time"
)

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

//QuotEntryGrp is a fix50 Component
type QuotEntryGrp struct {
	//NoQuoteEntries is a required field for QuotEntryGrp.
	NoQuoteEntries []NoQuoteEntries `fix:"295"`
}

func (m *QuotEntryGrp) SetNoQuoteEntries(v []NoQuoteEntries) { m.NoQuoteEntries = v }
