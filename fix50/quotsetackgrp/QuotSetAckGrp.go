package quotsetackgrp

import (
	"github.com/quickfixgo/quickfix/fix50/quotentryackgrp"
	"github.com/quickfixgo/quickfix/fix50/underlyinginstrument"
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
}

//QuotSetAckGrp is a fix50 Component
type QuotSetAckGrp struct {
	//NoQuoteSets is a non-required field for QuotSetAckGrp.
	NoQuoteSets []NoQuoteSets `fix:"296,omitempty"`
}

func (m *QuotSetAckGrp) SetNoQuoteSets(v []NoQuoteSets) { m.NoQuoteSets = v }
