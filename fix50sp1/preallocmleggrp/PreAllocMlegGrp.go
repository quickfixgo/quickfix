package preallocmleggrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/nestedparties3"
)

func New() *PreAllocMlegGrp {
	var m PreAllocMlegGrp
	return &m
}

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
	//NestedParties3 is a non-required component for NoAllocs.
	NestedParties3 *nestedparties3.NestedParties3
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

//PreAllocMlegGrp is a fix50sp1 Component
type PreAllocMlegGrp struct {
	//NoAllocs is a non-required field for PreAllocMlegGrp.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
}

func (m *PreAllocMlegGrp) SetNoAllocs(v []NoAllocs) { m.NoAllocs = v }
