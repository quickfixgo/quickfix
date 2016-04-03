package quotcxlentriesgrp

import (
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
)

func New() *QuotCxlEntriesGrp {
	var m QuotCxlEntriesGrp
	return &m
}

//NoQuoteEntries is a repeating group in QuotCxlEntriesGrp
type NoQuoteEntries struct {
	//Instrument is a non-required component for NoQuoteEntries.
	Instrument *instrument.Instrument
	//FinancingDetails is a non-required component for NoQuoteEntries.
	FinancingDetails *financingdetails.FinancingDetails
	//UndInstrmtGrp is a non-required component for NoQuoteEntries.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp is a non-required component for NoQuoteEntries.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
}

//QuotCxlEntriesGrp is a fix50 Component
type QuotCxlEntriesGrp struct {
	//NoQuoteEntries is a non-required field for QuotCxlEntriesGrp.
	NoQuoteEntries []NoQuoteEntries `fix:"295,omitempty"`
}

func (m *QuotCxlEntriesGrp) SetNoQuoteEntries(v []NoQuoteEntries) { m.NoQuoteEntries = v }
