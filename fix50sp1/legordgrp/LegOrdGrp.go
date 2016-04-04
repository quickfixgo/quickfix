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

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg)       { m.InstrumentLeg = &v }
func (m *NoLegs) SetLegQty(v float64)                                  { m.LegQty = &v }
func (m *NoLegs) SetLegSwapType(v int)                                 { m.LegSwapType = &v }
func (m *NoLegs) SetLegStipulations(v legstipulations.LegStipulations) { m.LegStipulations = &v }
func (m *NoLegs) SetLegPreAllocGrp(v legpreallocgrp.LegPreAllocGrp)    { m.LegPreAllocGrp = &v }
func (m *NoLegs) SetLegPositionEffect(v string)                        { m.LegPositionEffect = &v }
func (m *NoLegs) SetLegCoveredOrUncovered(v int)                       { m.LegCoveredOrUncovered = &v }
func (m *NoLegs) SetNestedParties(v nestedparties.NestedParties)       { m.NestedParties = &v }
func (m *NoLegs) SetLegRefID(v string)                                 { m.LegRefID = &v }
func (m *NoLegs) SetLegSettlType(v string)                             { m.LegSettlType = &v }
func (m *NoLegs) SetLegSettlDate(v string)                             { m.LegSettlDate = &v }
func (m *NoLegs) SetLegOrderQty(v float64)                             { m.LegOrderQty = &v }
func (m *NoLegs) SetLegAllocID(v string)                               { m.LegAllocID = &v }
func (m *NoLegs) SetLegVolatility(v float64)                           { m.LegVolatility = &v }
func (m *NoLegs) SetLegDividendYield(v float64)                        { m.LegDividendYield = &v }
func (m *NoLegs) SetLegCurrencyRatio(v float64)                        { m.LegCurrencyRatio = &v }
func (m *NoLegs) SetLegExecInst(v string)                              { m.LegExecInst = &v }
func (m *NoLegs) SetLegSettlCurrency(v string)                         { m.LegSettlCurrency = &v }

//LegOrdGrp is a fix50sp1 Component
type LegOrdGrp struct {
	//NoLegs is a required field for LegOrdGrp.
	NoLegs []NoLegs `fix:"555"`
}

func (m *LegOrdGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
