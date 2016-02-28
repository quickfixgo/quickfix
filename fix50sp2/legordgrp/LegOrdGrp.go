package legordgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp2/legpreallocgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/legstipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties"
)

//NoLegs is a repeating group in LegOrdGrp
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegStipulations Component
	LegStipulations legstipulations.Component
	//LegPreAllocGrp Component
	LegPreAllocGrp legpreallocgrp.Component
	//LegPositionEffect is a non-required field for NoLegs.
	LegPositionEffect *string `fix:"564"`
	//LegCoveredOrUncovered is a non-required field for NoLegs.
	LegCoveredOrUncovered *int `fix:"565"`
	//NestedParties Component
	NestedParties nestedparties.Component
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

//Component is a fix50sp2 LegOrdGrp Component
type Component struct {
	//NoLegs is a required field for LegOrdGrp.
	NoLegs []NoLegs `fix:"555"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoLegs(v []NoLegs) { m.NoLegs = v }
