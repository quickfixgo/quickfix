package quotreqlegsgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp1/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp1/legstipulations"
	"github.com/quickfixgo/quickfix/fix50sp1/nestedparties"
)

//NoLegs is a repeating group in QuotReqLegsGrp
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
	//LegStipulations Component
	LegStipulations legstipulations.Component
	//NestedParties Component
	NestedParties nestedparties.Component
	//LegBenchmarkCurveData Component
	LegBenchmarkCurveData legbenchmarkcurvedata.Component
	//LegOrderQty is a non-required field for NoLegs.
	LegOrderQty *float64 `fix:"685"`
	//LegRefID is a non-required field for NoLegs.
	LegRefID *string `fix:"654"`
}

//Component is a fix50sp1 QuotReqLegsGrp Component
type Component struct {
	//NoLegs is a non-required field for QuotReqLegsGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoLegs(v []NoLegs) { m.NoLegs = v }
