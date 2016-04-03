package quotsetackgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/quotentryackgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/underlyinginstrument"
	"time"
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
	//QuoteSetValidUntilTime is a non-required field for NoQuoteSets.
	QuoteSetValidUntilTime *time.Time `fix:"367"`
}

//QuotSetAckGrp is a fix50sp2 Component
type QuotSetAckGrp struct {
	//NoQuoteSets is a non-required field for QuotSetAckGrp.
	NoQuoteSets []NoQuoteSets `fix:"296,omitempty"`
}

func (m *QuotSetAckGrp) SetNoQuoteSets(v []NoQuoteSets) { m.NoQuoteSets = v }
