package instrmtlegioigrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50/legstipulations"
)

//New returns an initialized InstrmtLegIOIGrp instance
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

//NewNoLegs returns an initialized NoLegs instance
func NewNoLegs() *NoLegs {
	var m NoLegs
	return &m
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg)       { m.InstrumentLeg = &v }
func (m *NoLegs) SetLegIOIQty(v string)                                { m.LegIOIQty = &v }
func (m *NoLegs) SetLegStipulations(v legstipulations.LegStipulations) { m.LegStipulations = &v }

//InstrmtLegIOIGrp is a fix50 Component
type InstrmtLegIOIGrp struct {
	//NoLegs is a non-required field for InstrmtLegIOIGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *InstrmtLegIOIGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
