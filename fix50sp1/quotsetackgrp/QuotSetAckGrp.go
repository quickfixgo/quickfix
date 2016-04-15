package quotsetackgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/quotentryackgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/underlyinginstrument"
)

func New() *QuotSetAckGrp {
	var m QuotSetAckGrp
	return &m
}

//NoQuoteSets is a repeating group in QuotSetAckGrp
type NoQuoteSets struct {
	//QuoteSetID is a non-required field for NoQuoteSets.
	QuoteSetID *string `fix:"302"`
	//UnderlyingInstrument is a non-required component for NoQuoteSets.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
	//TotNoQuoteEntries is a non-required field for NoQuoteSets.
	TotNoQuoteEntries *int `fix:"304"`
	//LastFragment is a non-required field for NoQuoteSets.
	LastFragment *bool `fix:"893"`
	//QuotEntryAckGrp is a non-required component for NoQuoteSets.
	QuotEntryAckGrp *quotentryackgrp.QuotEntryAckGrp
	//TotNoCxldQuotes is a non-required field for NoQuoteSets.
	TotNoCxldQuotes *int `fix:"1168"`
	//TotNoAccQuotes is a non-required field for NoQuoteSets.
	TotNoAccQuotes *int `fix:"1169"`
	//TotNoRejQuotes is a non-required field for NoQuoteSets.
	TotNoRejQuotes *int `fix:"1170"`
}

func (m *NoQuoteSets) SetQuoteSetID(v string) { m.QuoteSetID = &v }
func (m *NoQuoteSets) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}
func (m *NoQuoteSets) SetTotNoQuoteEntries(v int)                           { m.TotNoQuoteEntries = &v }
func (m *NoQuoteSets) SetLastFragment(v bool)                               { m.LastFragment = &v }
func (m *NoQuoteSets) SetQuotEntryAckGrp(v quotentryackgrp.QuotEntryAckGrp) { m.QuotEntryAckGrp = &v }
func (m *NoQuoteSets) SetTotNoCxldQuotes(v int)                             { m.TotNoCxldQuotes = &v }
func (m *NoQuoteSets) SetTotNoAccQuotes(v int)                              { m.TotNoAccQuotes = &v }
func (m *NoQuoteSets) SetTotNoRejQuotes(v int)                              { m.TotNoRejQuotes = &v }

//QuotSetAckGrp is a fix50sp1 Component
type QuotSetAckGrp struct {
	//NoQuoteSets is a non-required field for QuotSetAckGrp.
	NoQuoteSets []NoQuoteSets `fix:"296,omitempty"`
}

func (m *QuotSetAckGrp) SetNoQuoteSets(v []NoQuoteSets) { m.NoQuoteSets = v }
