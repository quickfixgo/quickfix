package legpreallocgrp

import (
	"github.com/quickfixgo/quickfix/fix50/nestedparties2"
)

//NoLegAllocs is a repeating group in LegPreAllocGrp
type NoLegAllocs struct {
	//LegAllocAccount is a non-required field for NoLegAllocs.
	LegAllocAccount *string `fix:"671"`
	//LegIndividualAllocID is a non-required field for NoLegAllocs.
	LegIndividualAllocID *string `fix:"672"`
	//NestedParties2 Component
	NestedParties2 nestedparties2.Component
	//LegAllocQty is a non-required field for NoLegAllocs.
	LegAllocQty *float64 `fix:"673"`
	//LegAllocAcctIDSource is a non-required field for NoLegAllocs.
	LegAllocAcctIDSource *string `fix:"674"`
	//LegSettlCurrency is a non-required field for NoLegAllocs.
	LegSettlCurrency *string `fix:"675"`
}

//Component is a fix50 LegPreAllocGrp Component
type Component struct {
	//NoLegAllocs is a non-required field for LegPreAllocGrp.
	NoLegAllocs []NoLegAllocs `fix:"670,omitempty"`
}

func New() *Component { return new(Component) }
