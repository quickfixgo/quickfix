package legquotgrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/legstipulations"
	"github.com/quickfixgo/quickfix/fix50/nestedparties"
)

//New returns an initialized LegQuotGrp instance
func New() *LegQuotGrp {
	var m LegQuotGrp
	return &m
}

//NoLegs is a repeating group in LegQuotGrp
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
	//LegStipulations is a non-required component for NoLegs.
	LegStipulations *legstipulations.LegStipulations
	//NestedParties is a non-required component for NoLegs.
	NestedParties *nestedparties.NestedParties
	//LegPriceType is a non-required field for NoLegs.
	LegPriceType *int `fix:"686"`
	//LegBidPx is a non-required field for NoLegs.
	LegBidPx *float64 `fix:"681"`
	//LegOfferPx is a non-required field for NoLegs.
	LegOfferPx *float64 `fix:"684"`
	//LegBenchmarkCurveData is a non-required component for NoLegs.
	LegBenchmarkCurveData *legbenchmarkcurvedata.LegBenchmarkCurveData
	//LegOrderQty is a non-required field for NoLegs.
	LegOrderQty *float64 `fix:"685"`
	//LegRefID is a non-required field for NoLegs.
	LegRefID *string `fix:"654"`
	//LegBidForwardPoints is a non-required field for NoLegs.
	LegBidForwardPoints *float64 `fix:"1067"`
	//LegOfferForwardPoints is a non-required field for NoLegs.
	LegOfferForwardPoints *float64 `fix:"1068"`
}

//NewNoLegs returns an initialized NoLegs instance
func NewNoLegs() *NoLegs {
	var m NoLegs
	return &m
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg)       { m.InstrumentLeg = &v }
func (m *NoLegs) SetLegQty(v float64)                                  { m.LegQty = &v }
func (m *NoLegs) SetLegSwapType(v int)                                 { m.LegSwapType = &v }
func (m *NoLegs) SetLegSettlType(v string)                             { m.LegSettlType = &v }
func (m *NoLegs) SetLegSettlDate(v string)                             { m.LegSettlDate = &v }
func (m *NoLegs) SetLegStipulations(v legstipulations.LegStipulations) { m.LegStipulations = &v }
func (m *NoLegs) SetNestedParties(v nestedparties.NestedParties)       { m.NestedParties = &v }
func (m *NoLegs) SetLegPriceType(v int)                                { m.LegPriceType = &v }
func (m *NoLegs) SetLegBidPx(v float64)                                { m.LegBidPx = &v }
func (m *NoLegs) SetLegOfferPx(v float64)                              { m.LegOfferPx = &v }
func (m *NoLegs) SetLegBenchmarkCurveData(v legbenchmarkcurvedata.LegBenchmarkCurveData) {
	m.LegBenchmarkCurveData = &v
}
func (m *NoLegs) SetLegOrderQty(v float64)           { m.LegOrderQty = &v }
func (m *NoLegs) SetLegRefID(v string)               { m.LegRefID = &v }
func (m *NoLegs) SetLegBidForwardPoints(v float64)   { m.LegBidForwardPoints = &v }
func (m *NoLegs) SetLegOfferForwardPoints(v float64) { m.LegOfferForwardPoints = &v }

//LegQuotGrp is a fix50 Component
type LegQuotGrp struct {
	//NoLegs is a non-required field for LegQuotGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *LegQuotGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
