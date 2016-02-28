package quotsetackgrp

import (
	"github.com/quickfixgo/quickfix/fix50/quotentryackgrp"
	"github.com/quickfixgo/quickfix/fix50/underlyinginstrument"
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
}

//Component is a fix50 QuotSetAckGrp Component
type Component struct {
	//NoQuoteSets is a non-required field for QuotSetAckGrp.
	NoQuoteSets []NoQuoteSets `fix:"296,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoQuoteSets(v []NoQuoteSets) { m.NoQuoteSets = v }
