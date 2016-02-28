package trdallocgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/nestedparties2"
)

//NoAllocs is a repeating group in TrdAllocGrp
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties2 Component
	NestedParties2 nestedparties2.Component
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
	//AllocCustomerCapacity is a non-required field for NoAllocs.
	AllocCustomerCapacity *string `fix:"993"`
	//AllocMethod is a non-required field for NoAllocs.
	AllocMethod *int `fix:"1002"`
	//SecondaryIndividualAllocID is a non-required field for NoAllocs.
	SecondaryIndividualAllocID *string `fix:"989"`
	//AllocClearingFeeIndicator is a non-required field for NoAllocs.
	AllocClearingFeeIndicator *string `fix:"1136"`
}

//Component is a fix50sp1 TrdAllocGrp Component
type Component struct {
	//NoAllocs is a non-required field for TrdAllocGrp.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoAllocs(v []NoAllocs) { m.NoAllocs = v }
