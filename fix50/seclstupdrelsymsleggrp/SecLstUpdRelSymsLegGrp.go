package seclstupdrelsymsleggrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/legstipulations"
)

//New returns an initialized SecLstUpdRelSymsLegGrp instance
func New() *SecLstUpdRelSymsLegGrp {
	var m SecLstUpdRelSymsLegGrp
	return &m
}

//NoLegs is a repeating group in SecLstUpdRelSymsLegGrp
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegStipulations is a non-required component for NoLegs.
	LegStipulations *legstipulations.LegStipulations
	//LegBenchmarkCurveData is a non-required component for NoLegs.
	LegBenchmarkCurveData *legbenchmarkcurvedata.LegBenchmarkCurveData
}

//NewNoLegs returns an initialized NoLegs instance
func NewNoLegs() *NoLegs {
	var m NoLegs
	return &m
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg)       { m.InstrumentLeg = &v }
func (m *NoLegs) SetLegSwapType(v int)                                 { m.LegSwapType = &v }
func (m *NoLegs) SetLegSettlType(v string)                             { m.LegSettlType = &v }
func (m *NoLegs) SetLegStipulations(v legstipulations.LegStipulations) { m.LegStipulations = &v }
func (m *NoLegs) SetLegBenchmarkCurveData(v legbenchmarkcurvedata.LegBenchmarkCurveData) {
	m.LegBenchmarkCurveData = &v
}

//SecLstUpdRelSymsLegGrp is a fix50 Component
type SecLstUpdRelSymsLegGrp struct {
	//NoLegs is a non-required field for SecLstUpdRelSymsLegGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *SecLstUpdRelSymsLegGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
