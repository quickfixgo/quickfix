package legordgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp1/legpreallocgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/legstipulations"
	"github.com/quickfixgo/quickfix/fix50sp1/nestedparties"
)

func New(nolegs []NoLegs) *LegOrdGrp {
	var m LegOrdGrp
	m.SetNoLegs(nolegs)
	return &m
}

//NoLegs is a repeating group in LegOrdGrp
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegStipulations is a non-required component for NoLegs.
	LegStipulations *legstipulations.LegStipulations
	//LegPreAllocGrp is a non-required component for NoLegs.
	LegPreAllocGrp *legpreallocgrp.LegPreAllocGrp
	//LegPositionEffect is a non-required field for NoLegs.
	LegPositionEffect *string `fix:"564"`
	//LegCoveredOrUncovered is a non-required field for NoLegs.
	LegCoveredOrUncovered *int `fix:"565"`
	//NestedParties is a non-required component for NoLegs.
	NestedParties *nestedparties.NestedParties
	//LegRefID is a non-required field for NoLegs.
	LegRefID *string `fix:"654"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
	//LegOrderQty is a non-required field for NoLegs.
	LegOrderQty *float64 `fix:"685"`
	//LegAllocID is a non-required field for NoLegs.
	LegAllocID *string `fix:"1366"`
	//LegVolatility is a non-required field for NoLegs.
	LegVolatility *float64 `fix:"1379"`
	//LegDividendYield is a non-required field for NoLegs.
	LegDividendYield *float64 `fix:"1381"`
	//LegCurrencyRatio is a non-required field for NoLegs.
	LegCurrencyRatio *float64 `fix:"1383"`
	//LegExecInst is a non-required field for NoLegs.
	LegExecInst *string `fix:"1384"`
	//LegSettlCurrency is a non-required field for NoLegs.
	LegSettlCurrency *string `fix:"675"`
}

//LegOrdGrp is a fix50sp1 Component
type LegOrdGrp struct {
	//NoLegs is a required field for LegOrdGrp.
	NoLegs []NoLegs `fix:"555"`
}

func (m *LegOrdGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
