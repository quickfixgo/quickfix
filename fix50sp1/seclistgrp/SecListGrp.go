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
	instrument.Instrument
	//InstrumentExtension Component
	instrumentextension.InstrumentExtension
	//FinancingDetails Component
	financingdetails.FinancingDetails
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//Stipulations Component
	stipulations.Stipulations
	//InstrmtLegSecListGrp Component
	instrmtlegseclistgrp.InstrmtLegSecListGrp
	//SpreadOrBenchmarkCurveData Component
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData Component
	yielddata.YieldData
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
	//SecurityTradingRules Component
	securitytradingrules.SecurityTradingRules
	//StrikeRules Component
	strikerules.StrikeRules
}

//SecListGrp is a fix50sp1 Component
type SecListGrp struct {
	//NoRelatedSym is a non-required field for SecListGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *SecListGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
