package quotcxlentriesgrp

import (
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
)

//NoQuoteEntries is a repeating group in QuotCxlEntriesGrp
type NoQuoteEntries struct {
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
}

//Component is a fix50 QuotCxlEntriesGrp Component
type Component struct {
	//NoQuoteEntries is a non-required field for QuotCxlEntriesGrp.
	NoQuoteEntries []NoQuoteEntries `fix:"295,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoQuoteEntries(v []NoQuoteEntries) { m.NoQuoteEntries = v }
