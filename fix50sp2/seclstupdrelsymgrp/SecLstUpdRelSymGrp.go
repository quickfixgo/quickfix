package seclstupdrelsymgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp2/seclstupdrelsymsleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/securitytradingrules"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/strikerules"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/yielddata"
	"time"
)

//New returns an initialized SecLstUpdRelSymGrp instance
func New() *SecLstUpdRelSymGrp {
	var m SecLstUpdRelSymGrp
	return &m
}

//NoRelatedSym is a repeating group in SecLstUpdRelSymGrp
type NoRelatedSym struct {
	//Instrument is a non-required component for NoRelatedSym.
	Instrument *instrument.Instrument
	//InstrumentExtension is a non-required component for NoRelatedSym.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//FinancingDetails is a non-required component for NoRelatedSym.
	FinancingDetails *financingdetails.FinancingDetails
	//SecLstUpdRelSymsLegGrp is a non-required component for NoRelatedSym.
	SecLstUpdRelSymsLegGrp *seclstupdrelsymsleggrp.SecLstUpdRelSymsLegGrp
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
	//UndInstrmtGrp is a non-required component for NoRelatedSym.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//Stipulations is a non-required component for NoRelatedSym.
	Stipulations *stipulations.Stipulations
	//ListUpdateAction is a non-required field for NoRelatedSym.
	ListUpdateAction *string `fix:"1324"`
	//SecurityTradingRules is a non-required component for NoRelatedSym.
	SecurityTradingRules *securitytradingrules.SecurityTradingRules
	//StrikeRules is a non-required component for NoRelatedSym.
	StrikeRules *strikerules.StrikeRules
	//RelSymTransactTime is a non-required field for NoRelatedSym.
	RelSymTransactTime *time.Time `fix:"1504"`
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
func (m *NoRelatedSym) SetSecLstUpdRelSymsLegGrp(v seclstupdrelsymsleggrp.SecLstUpdRelSymsLegGrp) {
	m.SecLstUpdRelSymsLegGrp = &v
}
func (m *NoRelatedSym) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *NoRelatedSym) SetYieldData(v yielddata.YieldData)             { m.YieldData = &v }
func (m *NoRelatedSym) SetText(v string)                               { m.Text = &v }
func (m *NoRelatedSym) SetEncodedTextLen(v int)                        { m.EncodedTextLen = &v }
func (m *NoRelatedSym) SetEncodedText(v string)                        { m.EncodedText = &v }
func (m *NoRelatedSym) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
func (m *NoRelatedSym) SetCurrency(v string)                           { m.Currency = &v }
func (m *NoRelatedSym) SetStipulations(v stipulations.Stipulations)    { m.Stipulations = &v }
func (m *NoRelatedSym) SetListUpdateAction(v string)                   { m.ListUpdateAction = &v }
func (m *NoRelatedSym) SetSecurityTradingRules(v securitytradingrules.SecurityTradingRules) {
	m.SecurityTradingRules = &v
}
func (m *NoRelatedSym) SetStrikeRules(v strikerules.StrikeRules) { m.StrikeRules = &v }
func (m *NoRelatedSym) SetRelSymTransactTime(v time.Time)        { m.RelSymTransactTime = &v }

//SecLstUpdRelSymGrp is a fix50sp2 Component
type SecLstUpdRelSymGrp struct {
	//NoRelatedSym is a non-required field for SecLstUpdRelSymGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *SecLstUpdRelSymGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
