package instrmtlegioigrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp1/legstipulations"
)

//NoLegs is a repeating group in InstrmtLegIOIGrp
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
	//LegIOIQty is a non-required field for NoLegs.
	LegIOIQty *string `fix:"682"`
	//LegStipulations Component
	LegStipulations legstipulations.Component
}

//Component is a fix50sp1 InstrmtLegIOIGrp Component
type Component struct {
	//NoLegs is a non-required field for InstrmtLegIOIGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoLegs(v []NoLegs) { m.NoLegs = v }
