package quotcxlentriesgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
)

//New returns an initialized QuotCxlEntriesGrp instance
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

//NewNoQuoteEntries returns an initialized NoQuoteEntries instance
func NewNoQuoteEntries() *NoQuoteEntries {
	var m NoQuoteEntries
	return &m
}

func (m *NoQuoteEntries) SetInstrument(v instrument.Instrument) { m.Instrument = &v }
func (m *NoQuoteEntries) SetFinancingDetails(v financingdetails.FinancingDetails) {
	m.FinancingDetails = &v
}
func (m *NoQuoteEntries) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
func (m *NoQuoteEntries) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp) { m.InstrmtLegGrp = &v }

//QuotCxlEntriesGrp is a fix50sp2 Component
type QuotCxlEntriesGrp struct {
	//NoQuoteEntries is a non-required field for QuotCxlEntriesGrp.
	NoQuoteEntries []NoQuoteEntries `fix:"295,omitempty"`
}

func (m *QuotCxlEntriesGrp) SetNoQuoteEntries(v []NoQuoteEntries) { m.NoQuoteEntries = v }
