package quotsetackgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/quotentryackgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/underlyinginstrument"
	"time"
)

//NoQuoteSets is a repeating group in QuotSetAckGrp
type NoQuoteSets struct {
	//QuoteSetID is a non-required field for NoQuoteSets.
	QuoteSetID *string `fix:"302"`
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//TotNoQuoteEntries is a non-required field for NoQuoteSets.
	TotNoQuoteEntries *int `fix:"304"`
	//LastFragment is a non-required field for NoQuoteSets.
	LastFragment *bool `fix:"893"`
	//QuotEntryAckGrp Component
	QuotEntryAckGrp quotentryackgrp.Component
	//TotNoCxldQuotes is a non-required field for NoQuoteSets.
	TotNoCxldQuotes *int `fix:"1168"`
	//TotNoAccQuotes is a non-required field for NoQuoteSets.
	TotNoAccQuotes *int `fix:"1169"`
	//TotNoRejQuotes is a non-required field for NoQuoteSets.
	TotNoRejQuotes *int `fix:"1170"`
	//QuoteSetValidUntilTime is a non-required field for NoQuoteSets.
	QuoteSetValidUntilTime *time.Time `fix:"367"`
}

//Component is a fix50sp2 QuotSetAckGrp Component
type Component struct {
	//NoQuoteSets is a non-required field for QuotSetAckGrp.
	NoQuoteSets []NoQuoteSets `fix:"296,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoQuoteSets(v []NoQuoteSets) { m.NoQuoteSets = v }
