package preallocmleggrp

import (
	"github.com/quickfixgo/quickfix/fix50/nestedparties3"
)

//NoAllocs is a repeating group in PreAllocMlegGrp
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties3 Component
	NestedParties3 nestedparties3.Component
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

//Component is a fix50 PreAllocMlegGrp Component
type Component struct {
	//NoAllocs is a non-required field for PreAllocMlegGrp.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoAllocs(v []NoAllocs) { m.NoAllocs = v }
