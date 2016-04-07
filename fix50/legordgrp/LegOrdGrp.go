package legordgrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50/legpreallocgrp"
	"github.com/quickfixgo/quickfix/fix50/legstipulations"
	"github.com/quickfixgo/quickfix/fix50/nestedparties"
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

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg)       { m.InstrumentLeg = &v }
func (m *NoLegs) SetLegQty(v float64)                                  { m.LegQty = &v }
func (m *NoLegs) SetLegSwapType(v int)                                 { m.LegSwapType = &v }
func (m *NoLegs) SetLegStipulations(v legstipulations.LegStipulations) { m.LegStipulations = &v }
func (m *NoLegs) SetLegPreAllocGrp(v legpreallocgrp.LegPreAllocGrp)    { m.LegPreAllocGrp = &v }
func (m *NoLegs) SetLegPositionEffect(v string)                        { m.LegPositionEffect = &v }
func (m *NoLegs) SetLegCoveredOrUncovered(v int)                       { m.LegCoveredOrUncovered = &v }
func (m *NoLegs) SetNestedParties(v nestedparties.NestedParties)       { m.NestedParties = &v }
func (m *NoLegs) SetLegRefID(v string)                                 { m.LegRefID = &v }
func (m *NoLegs) SetLegPrice(v float64)                                { m.LegPrice = &v }
func (m *NoLegs) SetLegSettlType(v string)                             { m.LegSettlType = &v }
func (m *NoLegs) SetLegSettlDate(v string)                             { m.LegSettlDate = &v }
func (m *NoLegs) SetLegOptionRatio(v float64)                          { m.LegOptionRatio = &v }
func (m *NoLegs) SetLegOrderQty(v float64)                             { m.LegOrderQty = &v }

//LegOrdGrp is a fix50 Component
type LegOrdGrp struct {
	//NoLegs is a required field for LegOrdGrp.
	NoLegs []NoLegs `fix:"555"`
}

func (m *LegOrdGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
