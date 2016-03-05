package quotsetgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/quotentrygrp"
	"github.com/quickfixgo/quickfix/fix50sp1/underlyinginstrument"
	"time"
)

//NoQuoteSets is a repeating group in QuotSetGrp
type NoQuoteSets struct {
	//QuoteSetID is a required field for NoQuoteSets.
	QuoteSetID string `fix:"302"`
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
	//QuoteSetValidUntilTime is a non-required field for NoQuoteSets.
	QuoteSetValidUntilTime *time.Time `fix:"367"`
	//TotNoQuoteEntries is a required field for NoQuoteSets.
	TotNoQuoteEntries int `fix:"304"`
	//LastFragment is a non-required field for NoQuoteSets.
	LastFragment *bool `fix:"893"`
	//QuotEntryGrp Component
	quotentrygrp.QuotEntryGrp
}

//QuotSetGrp is a fix50sp1 Component
type QuotSetGrp struct {
	//NoQuoteSets is a required field for QuotSetGrp.
	NoQuoteSets []NoQuoteSets `fix:"296"`
}

func (m *QuotSetGrp) SetNoQuoteSets(v []NoQuoteSets) { m.NoQuoteSets = v }
