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

func New() *SecListGrp {
	var m SecListGrp
	return &m
}

//NoRelatedSym is a repeating group in SecListGrp
type NoRelatedSym struct {
	//Instrument is a non-required component for NoRelatedSym.
	Instrument *instrument.Instrument
	//InstrumentExtension is a non-required component for NoRelatedSym.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//FinancingDetails is a non-required component for NoRelatedSym.
	FinancingDetails *financingdetails.FinancingDetails
	//UndInstrmtGrp is a non-required component for NoRelatedSym.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//Stipulations is a non-required component for NoRelatedSym.
	Stipulations *stipulations.Stipulations
	//InstrmtLegSecListGrp is a non-required component for NoRelatedSym.
	InstrmtLegSecListGrp *instrmtlegseclistgrp.InstrmtLegSecListGrp
	//SpreadOrBenchmarkCurveData is a non-required component for NoRelatedSym.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData is a non-required component for NoRelatedSym.
	YieldData *yielddata.YieldData
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
	//SecurityTradingRules is a non-required component for NoRelatedSym.
	SecurityTradingRules *securitytradingrules.SecurityTradingRules
	//StrikeRules is a non-required component for NoRelatedSym.
	StrikeRules *strikerules.StrikeRules
}

func (m *NoRelatedSym) SetInstrument(v instrument.Instrument) { m.Instrument = &v }
func (m *NoRelatedSym) SetInstrumentExtension(v instrumentextension.InstrumentExtension) {
	m.InstrumentExtension = &v
}
func (m *NoRelatedSym) SetFinancingDetails(v financingdetails.FinancingDetails) {
	m.FinancingDetails = &v
}
func (m *NoRelatedSym) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
func (m *NoRelatedSym) SetCurrency(v string)                           { m.Currency = &v }
func (m *NoRelatedSym) SetStipulations(v stipulations.Stipulations)    { m.Stipulations = &v }
func (m *NoRelatedSym) SetInstrmtLegSecListGrp(v instrmtlegseclistgrp.InstrmtLegSecListGrp) {
	m.InstrmtLegSecListGrp = &v
}
func (m *NoRelatedSym) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *NoRelatedSym) SetYieldData(v yielddata.YieldData) { m.YieldData = &v }
func (m *NoRelatedSym) SetText(v string)                   { m.Text = &v }
func (m *NoRelatedSym) SetEncodedTextLen(v int)            { m.EncodedTextLen = &v }
func (m *NoRelatedSym) SetEncodedText(v string)            { m.EncodedText = &v }
func (m *NoRelatedSym) SetSecurityTradingRules(v securitytradingrules.SecurityTradingRules) {
	m.SecurityTradingRules = &v
}
func (m *NoRelatedSym) SetStrikeRules(v strikerules.StrikeRules) { m.StrikeRules = &v }

//SecListGrp is a fix50sp1 Component
type SecListGrp struct {
	//NoRelatedSym is a non-required field for SecListGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *SecListGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
