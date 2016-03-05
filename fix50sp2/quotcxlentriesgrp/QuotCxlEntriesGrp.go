package quotcxlentriesgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
)

//NoQuoteEntries is a repeating group in QuotCxlEntriesGrp
type NoQuoteEntries struct {
	//Instrument Component
	instrument.Instrument
	//FinancingDetails Component
	financingdetails.FinancingDetails
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp Component
	instrmtleggrp.InstrmtLegGrp
}

//QuotCxlEntriesGrp is a fix50sp2 Component
type QuotCxlEntriesGrp struct {
	//NoQuoteEntries is a non-required field for QuotCxlEntriesGrp.
	NoQuoteEntries []NoQuoteEntries `fix:"295,omitempty"`
}

func (m *QuotCxlEntriesGrp) SetNoQuoteEntries(v []NoQuoteEntries) { m.NoQuoteEntries = v }
