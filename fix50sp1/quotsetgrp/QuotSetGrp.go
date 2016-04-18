package quotsetgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/quotentrygrp"
	"github.com/quickfixgo/quickfix/fix50sp1/underlyinginstrument"
	"time"
)

//New returns an initialized QuotSetGrp instance
func New(noquotesets []NoQuoteSets) *QuotSetGrp {
	var m QuotSetGrp
	m.SetNoQuoteSets(noquotesets)
	return &m
}

//NoQuoteSets is a repeating group in QuotSetGrp
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
	//QuotEntryGrp is a required component for NoQuoteSets.
	quotentrygrp.QuotEntryGrp
}

//NewNoQuoteSets returns an initialized NoQuoteSets instance
func NewNoQuoteSets(quotesetid string, totnoquoteentries int, quotentrygrp quotentrygrp.QuotEntryGrp) *NoQuoteSets {
	var m NoQuoteSets
	m.SetQuoteSetID(quotesetid)
	m.SetTotNoQuoteEntries(totnoquoteentries)
	m.SetQuotEntryGrp(quotentrygrp)
	return &m
}

func (m *NoQuoteSets) SetQuoteSetID(v string) { m.QuoteSetID = v }
func (m *NoQuoteSets) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}
func (m *NoQuoteSets) SetQuoteSetValidUntilTime(v time.Time)       { m.QuoteSetValidUntilTime = &v }
func (m *NoQuoteSets) SetTotNoQuoteEntries(v int)                  { m.TotNoQuoteEntries = v }
func (m *NoQuoteSets) SetLastFragment(v bool)                      { m.LastFragment = &v }
func (m *NoQuoteSets) SetQuotEntryGrp(v quotentrygrp.QuotEntryGrp) { m.QuotEntryGrp = v }

//QuotSetGrp is a fix50sp1 Component
type QuotSetGrp struct {
	//NoQuoteSets is a required field for QuotSetGrp.
	NoQuoteSets []NoQuoteSets `fix:"296"`
}

func (m *QuotSetGrp) SetNoQuoteSets(v []NoQuoteSets) { m.NoQuoteSets = v }
