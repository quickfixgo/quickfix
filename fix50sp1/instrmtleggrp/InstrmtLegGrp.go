package instrmtleggrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentleg"
)

//NoLegs is a repeating group in InstrmtLegGrp
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//Component is a fix50sp1 InstrmtLegGrp Component
type Component struct {
	//NoLegs is a non-required field for InstrmtLegGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func New() *Component { return new(Component) }
