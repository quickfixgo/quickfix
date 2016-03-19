package instrmtlegseclistgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp2/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/legstipulations"
)

func New() *InstrmtLegSecListGrp {
	var m InstrmtLegSecListGrp
	return &m
}

//NoLegs is a repeating group in InstrmtLegSecListGrp
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

//InstrmtLegSecListGrp is a fix50sp2 Component
type InstrmtLegSecListGrp struct {
	//NoLegs is a non-required field for InstrmtLegSecListGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *InstrmtLegSecListGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
