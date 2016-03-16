package preallocgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties"
)

//NoAllocs is a repeating group in PreAllocGrp
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties is a non-required component for NoAllocs.
	NestedParties *nestedparties.NestedParties
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

//PreAllocGrp is a fix50sp2 Component
type PreAllocGrp struct {
	//NoAllocs is a non-required field for PreAllocGrp.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
}

func (m *PreAllocGrp) SetNoAllocs(v []NoAllocs) { m.NoAllocs = v }
