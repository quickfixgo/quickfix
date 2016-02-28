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

//NoRelatedSym is a repeating group in SecLstUpdRelSymGrp
type NoRelatedSym struct {
	//Instrument Component
	Instrument instrument.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//SecLstUpdRelSymsLegGrp Component
	SecLstUpdRelSymsLegGrp seclstupdrelsymsleggrp.Component
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
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//Stipulations Component
	Stipulations stipulations.Component
	//ListUpdateAction is a non-required field for NoRelatedSym.
	ListUpdateAction *string `fix:"1324"`
	//SecurityTradingRules Component
	SecurityTradingRules securitytradingrules.Component
	//StrikeRules Component
	StrikeRules strikerules.Component
	//RelSymTransactTime is a non-required field for NoRelatedSym.
	RelSymTransactTime *time.Time `fix:"1504"`
}

//Component is a fix50sp2 SecLstUpdRelSymGrp Component
type Component struct {
	//NoRelatedSym is a non-required field for SecLstUpdRelSymGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
