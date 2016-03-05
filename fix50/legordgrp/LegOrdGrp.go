package legordgrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50/legpreallocgrp"
	"github.com/quickfixgo/quickfix/fix50/legstipulations"
	"github.com/quickfixgo/quickfix/fix50/nestedparties"
)

//NoLegs is a repeating group in LegOrdGrp
type NoLegs struct {
	//InstrumentLeg Component
	instrumentleg.InstrumentLeg
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegStipulations Component
	legstipulations.LegStipulations
	//LegPreAllocGrp Component
	legpreallocgrp.LegPreAllocGrp
	//LegPositionEffect is a non-required field for NoLegs.
	LegPositionEffect *string `fix:"564"`
	//LegCoveredOrUncovered is a non-required field for NoLegs.
	LegCoveredOrUncovered *int `fix:"565"`
	//NestedParties Component
	nestedparties.NestedParties
	//LegRefID is a non-required field for NoLegs.
	LegRefID *string `fix:"654"`
	//LegPrice is a non-required field for NoLegs.
	LegPrice *float64 `fix:"566"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
	//LegOptionRatio is a non-required field for NoLegs.
	LegOptionRatio *float64 `fix:"1017"`
	//LegOrderQty is a non-required field for NoLegs.
	LegOrderQty *float64 `fix:"685"`
}

//LegOrdGrp is a fix50 Component
type LegOrdGrp struct {
	//NoLegs is a required field for LegOrdGrp.
	NoLegs []NoLegs `fix:"555"`
}

func (m *LegOrdGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
