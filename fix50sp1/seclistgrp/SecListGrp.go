package seclistgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtlegseclistgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp1/securitytradingrules"
	"github.com/quickfixgo/quickfix/fix50sp1/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp1/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp1/strikerules"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/yielddata"
)

//NoRelatedSym is a repeating group in SecListGrp
type NoRelatedSym struct {
	//Instrument Component
	Instrument instrument.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//Stipulations Component
	Stipulations stipulations.Component
	//InstrmtLegSecListGrp Component
	InstrmtLegSecListGrp instrmtlegseclistgrp.Component
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
	//SecurityTradingRules Component
	SecurityTradingRules securitytradingrules.Component
	//StrikeRules Component
	StrikeRules strikerules.Component
}

//Component is a fix50sp1 SecListGrp Component
type Component struct {
	//NoRelatedSym is a non-required field for SecListGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
