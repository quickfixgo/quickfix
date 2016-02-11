package legquotgrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrumentleg"
	"github.com/quickfixgo/quickfix/fix50/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/legstipulations"
	"github.com/quickfixgo/quickfix/fix50/nestedparties"
)

//NoLegs is a repeating group in LegQuotGrp
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
	//LegPriceType is a non-required field for NoLegs.
	LegPriceType *int `fix:"686"`
	//LegBidPx is a non-required field for NoLegs.
	LegBidPx *float64 `fix:"681"`
	//LegOfferPx is a non-required field for NoLegs.
	LegOfferPx *float64 `fix:"684"`
	//LegBenchmarkCurveData Component
	LegBenchmarkCurveData legbenchmarkcurvedata.Component
	//LegOrderQty is a non-required field for NoLegs.
	LegOrderQty *float64 `fix:"685"`
	//LegRefID is a non-required field for NoLegs.
	LegRefID *string `fix:"654"`
	//LegBidForwardPoints is a non-required field for NoLegs.
	LegBidForwardPoints *float64 `fix:"1067"`
	//LegOfferForwardPoints is a non-required field for NoLegs.
	LegOfferForwardPoints *float64 `fix:"1068"`
}

//Component is a fix50 LegQuotGrp Component
type Component struct {
	//NoLegs is a non-required field for LegQuotGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func New() *Component { return new(Component) }
