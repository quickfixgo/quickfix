package legquotstatgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp2/legstipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties"
)

//NoLegs is a repeating group in LegQuotStatGrp
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
	//LegOrderQty is a non-required field for NoLegs.
	LegOrderQty *float64 `fix:"685"`
}

//Component is a fix50sp2 LegQuotStatGrp Component
type Component struct {
	//NoLegs is a non-required field for LegQuotStatGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoLegs(v []NoLegs) { m.NoLegs = v }
