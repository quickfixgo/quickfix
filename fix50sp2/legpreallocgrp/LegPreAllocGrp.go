package legpreallocgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties2"
)

func New() *LegPreAllocGrp {
	var m LegPreAllocGrp
	return &m
}

//NoLegAllocs is a repeating group in LegPreAllocGrp
type NoLegAllocs struct {
	//LegAllocAccount is a non-required field for NoLegAllocs.
	LegAllocAccount *string `fix:"671"`
	//LegIndividualAllocID is a non-required field for NoLegAllocs.
	LegIndividualAllocID *string `fix:"672"`
	//LegAllocQty is a non-required field for NoLegAllocs.
	LegAllocQty *float64 `fix:"673"`
	//LegAllocAcctIDSource is a non-required field for NoLegAllocs.
	LegAllocAcctIDSource *string `fix:"674"`
	//LegAllocSettlCurrency is a non-required field for NoLegAllocs.
	LegAllocSettlCurrency *string `fix:"1367"`
	//NestedParties2 is a non-required component for NoLegAllocs.
	NestedParties2 *nestedparties2.NestedParties2
}

//LegPreAllocGrp is a fix50sp2 Component
type LegPreAllocGrp struct {
	//NoLegAllocs is a non-required field for LegPreAllocGrp.
	NoLegAllocs []NoLegAllocs `fix:"670,omitempty"`
}

func (m *LegPreAllocGrp) SetNoLegAllocs(v []NoLegAllocs) { m.NoLegAllocs = v }
