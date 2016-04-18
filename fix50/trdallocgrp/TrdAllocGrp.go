package trdallocgrp

import (
	"github.com/quickfixgo/quickfix/fix50/nestedparties2"
)

//New returns an initialized TrdAllocGrp instance
func New() *TrdAllocGrp {
	var m TrdAllocGrp
	return &m
}

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
	//NestedParties2 is a non-required component for NoAllocs.
	NestedParties2 *nestedparties2.NestedParties2
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

//NewNoAllocs returns an initialized NoAllocs instance
func NewNoAllocs() *NoAllocs {
	var m NoAllocs
	return &m
}

func (m *NoAllocs) SetAllocAccount(v string)                          { m.AllocAccount = &v }
func (m *NoAllocs) SetAllocAcctIDSource(v int)                        { m.AllocAcctIDSource = &v }
func (m *NoAllocs) SetAllocSettlCurrency(v string)                    { m.AllocSettlCurrency = &v }
func (m *NoAllocs) SetIndividualAllocID(v string)                     { m.IndividualAllocID = &v }
func (m *NoAllocs) SetNestedParties2(v nestedparties2.NestedParties2) { m.NestedParties2 = &v }
func (m *NoAllocs) SetAllocQty(v float64)                             { m.AllocQty = &v }
func (m *NoAllocs) SetAllocCustomerCapacity(v string)                 { m.AllocCustomerCapacity = &v }
func (m *NoAllocs) SetAllocMethod(v int)                              { m.AllocMethod = &v }
func (m *NoAllocs) SetSecondaryIndividualAllocID(v string)            { m.SecondaryIndividualAllocID = &v }
func (m *NoAllocs) SetAllocClearingFeeIndicator(v string)             { m.AllocClearingFeeIndicator = &v }

//TrdAllocGrp is a fix50 Component
type TrdAllocGrp struct {
	//NoAllocs is a non-required field for TrdAllocGrp.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
}

func (m *TrdAllocGrp) SetNoAllocs(v []NoAllocs) { m.NoAllocs = v }
