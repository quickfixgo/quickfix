package quotreqlegsgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp2/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/legstipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties"
)

//NoLegs is a repeating group in QuotReqLegsGrp
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
	//LegStipulations is a non-required component for NoLegs.
	LegStipulations *legstipulations.LegStipulations
	//NestedParties is a non-required component for NoLegs.
	NestedParties *nestedparties.NestedParties
	//LegBenchmarkCurveData is a non-required component for NoLegs.
	LegBenchmarkCurveData *legbenchmarkcurvedata.LegBenchmarkCurveData
	//LegOrderQty is a non-required field for NoLegs.
	LegOrderQty *float64 `fix:"685"`
	//LegRefID is a non-required field for NoLegs.
	LegRefID *string `fix:"654"`
}

//QuotReqLegsGrp is a fix50sp2 Component
type QuotReqLegsGrp struct {
	//NoLegs is a non-required field for QuotReqLegsGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *QuotReqLegsGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
