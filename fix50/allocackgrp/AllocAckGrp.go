package allocackgrp

import (
	"github.com/quickfixgo/quickfix/fix50/nestedparties"
)

func New() *AllocAckGrp {
	var m AllocAckGrp
	return &m
}

//NoAllocs is a repeating group in AllocAckGrp
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//AllocPrice is a non-required field for NoAllocs.
	AllocPrice *float64 `fix:"366"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//IndividualAllocRejCode is a non-required field for NoAllocs.
	IndividualAllocRejCode *int `fix:"776"`
	//AllocText is a non-required field for NoAllocs.
	AllocText *string `fix:"161"`
	//EncodedAllocTextLen is a non-required field for NoAllocs.
	EncodedAllocTextLen *int `fix:"360"`
	//EncodedAllocText is a non-required field for NoAllocs.
	EncodedAllocText *string `fix:"361"`
	//SecondaryIndividualAllocID is a non-required field for NoAllocs.
	SecondaryIndividualAllocID *string `fix:"989"`
	//AllocCustomerCapacity is a non-required field for NoAllocs.
	AllocCustomerCapacity *string `fix:"993"`
	//IndividualAllocType is a non-required field for NoAllocs.
	IndividualAllocType *int `fix:"992"`
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
	//NestedParties is a non-required component for NoAllocs.
	NestedParties *nestedparties.NestedParties
	//AllocPositionEffect is a non-required field for NoAllocs.
	AllocPositionEffect *string `fix:"1047"`
}

func (m *NoAllocs) SetAllocAccount(v string)                       { m.AllocAccount = &v }
func (m *NoAllocs) SetAllocAcctIDSource(v int)                     { m.AllocAcctIDSource = &v }
func (m *NoAllocs) SetAllocPrice(v float64)                        { m.AllocPrice = &v }
func (m *NoAllocs) SetIndividualAllocID(v string)                  { m.IndividualAllocID = &v }
func (m *NoAllocs) SetIndividualAllocRejCode(v int)                { m.IndividualAllocRejCode = &v }
func (m *NoAllocs) SetAllocText(v string)                          { m.AllocText = &v }
func (m *NoAllocs) SetEncodedAllocTextLen(v int)                   { m.EncodedAllocTextLen = &v }
func (m *NoAllocs) SetEncodedAllocText(v string)                   { m.EncodedAllocText = &v }
func (m *NoAllocs) SetSecondaryIndividualAllocID(v string)         { m.SecondaryIndividualAllocID = &v }
func (m *NoAllocs) SetAllocCustomerCapacity(v string)              { m.AllocCustomerCapacity = &v }
func (m *NoAllocs) SetIndividualAllocType(v int)                   { m.IndividualAllocType = &v }
func (m *NoAllocs) SetAllocQty(v float64)                          { m.AllocQty = &v }
func (m *NoAllocs) SetNestedParties(v nestedparties.NestedParties) { m.NestedParties = &v }
func (m *NoAllocs) SetAllocPositionEffect(v string)                { m.AllocPositionEffect = &v }

//AllocAckGrp is a fix50 Component
type AllocAckGrp struct {
	//NoAllocs is a non-required field for AllocAckGrp.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
}

func (m *AllocAckGrp) SetNoAllocs(v []NoAllocs) { m.NoAllocs = v }
