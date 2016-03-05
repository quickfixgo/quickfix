package instrmtlegseclistgrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/legstipulations"
)

//NoLegs is a repeating group in InstrmtLegSecListGrp
type NoLegs struct {
	//InstrumentLeg Component
	instrumentleg.InstrumentLeg
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegStipulations Component
	legstipulations.LegStipulations
	//LegBenchmarkCurveData Component
	legbenchmarkcurvedata.LegBenchmarkCurveData
}

//InstrmtLegSecListGrp is a fix50 Component
type InstrmtLegSecListGrp struct {
	//NoLegs is a non-required field for InstrmtLegSecListGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *InstrmtLegSecListGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
