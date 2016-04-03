package instrmtlegioigrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp1/legstipulations"
)

func New() *InstrmtLegIOIGrp {
	var m InstrmtLegIOIGrp
	return &m
}

//NoLegs is a repeating group in InstrmtLegIOIGrp
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
	//LegIOIQty is a non-required field for NoLegs.
	LegIOIQty *string `fix:"682"`
	//LegStipulations is a non-required component for NoLegs.
	LegStipulations *legstipulations.LegStipulations
}

//InstrmtLegIOIGrp is a fix50sp1 Component
type InstrmtLegIOIGrp struct {
	//NoLegs is a non-required field for InstrmtLegIOIGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *InstrmtLegIOIGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
