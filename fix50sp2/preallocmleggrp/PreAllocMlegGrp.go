package preallocmleggrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties3"
)

//New returns an initialized PreAllocMlegGrp instance
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

//NewNoAllocs returns an initialized NoAllocs instance
func NewNoAllocs() *NoAllocs {
	var m NoAllocs
	return &m
}

func (m *NoAllocs) SetAllocAccount(v string)                          { m.AllocAccount = &v }
func (m *NoAllocs) SetAllocAcctIDSource(v int)                        { m.AllocAcctIDSource = &v }
func (m *NoAllocs) SetAllocSettlCurrency(v string)                    { m.AllocSettlCurrency = &v }
func (m *NoAllocs) SetIndividualAllocID(v string)                     { m.IndividualAllocID = &v }
func (m *NoAllocs) SetNestedParties3(v nestedparties3.NestedParties3) { m.NestedParties3 = &v }
func (m *NoAllocs) SetAllocQty(v float64)                             { m.AllocQty = &v }

//PreAllocMlegGrp is a fix50sp2 Component
type PreAllocMlegGrp struct {
	//NoAllocs is a non-required field for PreAllocMlegGrp.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
}

func (m *PreAllocMlegGrp) SetNoAllocs(v []NoAllocs) { m.NoAllocs = v }
