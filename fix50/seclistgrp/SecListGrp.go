package seclistgrp

import (
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrmtlegseclistgrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/stipulations"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50/yielddata"
)

//New returns an initialized SecListGrp instance
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
	//RoundLot is a non-required field for NoRelatedSym.
	RoundLot *float64 `fix:"561"`
	//MinTradeVol is a non-required field for NoRelatedSym.
	MinTradeVol *float64 `fix:"562"`
	//TradingSessionID is a non-required field for NoRelatedSym.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoRelatedSym.
	TradingSessionSubID *string `fix:"625"`
	//ExpirationCycle is a non-required field for NoRelatedSym.
	ExpirationCycle *int `fix:"827"`
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
}

//NewNoRelatedSym returns an initialized NoRelatedSym instance
func NewNoRelatedSym() *NoRelatedSym {
	var m NoRelatedSym
	return &m
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
func (m *NoRelatedSym) SetRoundLot(v float64)              { m.RoundLot = &v }
func (m *NoRelatedSym) SetMinTradeVol(v float64)           { m.MinTradeVol = &v }
func (m *NoRelatedSym) SetTradingSessionID(v string)       { m.TradingSessionID = &v }
func (m *NoRelatedSym) SetTradingSessionSubID(v string)    { m.TradingSessionSubID = &v }
func (m *NoRelatedSym) SetExpirationCycle(v int)           { m.ExpirationCycle = &v }
func (m *NoRelatedSym) SetText(v string)                   { m.Text = &v }
func (m *NoRelatedSym) SetEncodedTextLen(v int)            { m.EncodedTextLen = &v }
func (m *NoRelatedSym) SetEncodedText(v string)            { m.EncodedText = &v }

//SecListGrp is a fix50 Component
type SecListGrp struct {
	//NoRelatedSym is a non-required field for SecListGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *SecListGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
