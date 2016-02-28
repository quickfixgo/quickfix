package allocackgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties"
)

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
	//NestedParties Component
	NestedParties nestedparties.Component
	//AllocPositionEffect is a non-required field for NoAllocs.
	AllocPositionEffect *string `fix:"1047"`
}

//Component is a fix50sp2 AllocAckGrp Component
type Component struct {
	//NoAllocs is a non-required field for AllocAckGrp.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoAllocs(v []NoAllocs) { m.NoAllocs = v }
