package legquotgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50sp1/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp1/legstipulations"
	"github.com/quickfixgo/quickfix/fix50sp1/nestedparties"
)

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

//LegQuotGrp is a fix50sp1 Component
type LegQuotGrp struct {
	//NoLegs is a non-required field for LegQuotGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *LegQuotGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
