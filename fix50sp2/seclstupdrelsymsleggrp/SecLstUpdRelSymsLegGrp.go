package seclstupdrelsymsleggrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp2/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/legstipulations"
)

//NoLegs is a repeating group in SecLstUpdRelSymsLegGrp
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegStipulations Component
	LegStipulations legstipulations.Component
	//LegBenchmarkCurveData Component
	LegBenchmarkCurveData legbenchmarkcurvedata.Component
}

//Component is a fix50sp2 SecLstUpdRelSymsLegGrp Component
type Component struct {
	//NoLegs is a non-required field for SecLstUpdRelSymsLegGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func New() *Component { return new(Component) }
