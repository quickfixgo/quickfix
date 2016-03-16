package seclstupdrelsymsleggrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp1/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp1/legstipulations"
)

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

//SecLstUpdRelSymsLegGrp is a fix50sp1 Component
type SecLstUpdRelSymsLegGrp struct {
	//NoLegs is a non-required field for SecLstUpdRelSymsLegGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *SecLstUpdRelSymsLegGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
