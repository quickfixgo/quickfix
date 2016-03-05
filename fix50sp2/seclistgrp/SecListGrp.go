package seclistgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtlegseclistgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp2/securitytradingrules"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/strikerules"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/yielddata"
	"time"
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
	//RelSymTransactTime is a non-required field for NoRelatedSym.
	RelSymTransactTime *time.Time `fix:"1504"`
}

//SecListGrp is a fix50sp2 Component
type SecListGrp struct {
	//NoRelatedSym is a non-required field for SecListGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *SecListGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
