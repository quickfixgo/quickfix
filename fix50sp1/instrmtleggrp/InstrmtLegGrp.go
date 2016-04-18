package instrmtleggrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentleg"
)

//New returns an initialized InstrmtLegGrp instance
func New() *InstrmtLegGrp {
	var m InstrmtLegGrp
	return &m
}

//NoLegs is a repeating group in InstrmtLegGrp
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
}

//NewNoLegs returns an initialized NoLegs instance
func NewNoLegs() *NoLegs {
	var m NoLegs
	return &m
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg) { m.InstrumentLeg = &v }

//InstrmtLegGrp is a fix50sp1 Component
type InstrmtLegGrp struct {
	//NoLegs is a non-required field for InstrmtLegGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *InstrmtLegGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
