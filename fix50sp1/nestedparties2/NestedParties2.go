package nestedparties2

import (
	"github.com/quickfixgo/quickfix/fix50sp1/nstdptys2subgrp"
)

//New returns an initialized NestedParties2 instance
func New() *NestedParties2 {
	var m NestedParties2
	return &m
}

//NoNested2PartyIDs is a repeating group in NestedParties2
type NoNested2PartyIDs struct {
	//Nested2PartyID is a non-required field for NoNested2PartyIDs.
	Nested2PartyID *string `fix:"757"`
	//Nested2PartyIDSource is a non-required field for NoNested2PartyIDs.
	Nested2PartyIDSource *string `fix:"758"`
	//Nested2PartyRole is a non-required field for NoNested2PartyIDs.
	Nested2PartyRole *int `fix:"759"`
	//NstdPtys2SubGrp is a non-required component for NoNested2PartyIDs.
	NstdPtys2SubGrp *nstdptys2subgrp.NstdPtys2SubGrp
}

//NewNoNested2PartyIDs returns an initialized NoNested2PartyIDs instance
func NewNoNested2PartyIDs() *NoNested2PartyIDs {
	var m NoNested2PartyIDs
	return &m
}

func (m *NoNested2PartyIDs) SetNested2PartyID(v string)       { m.Nested2PartyID = &v }
func (m *NoNested2PartyIDs) SetNested2PartyIDSource(v string) { m.Nested2PartyIDSource = &v }
func (m *NoNested2PartyIDs) SetNested2PartyRole(v int)        { m.Nested2PartyRole = &v }
func (m *NoNested2PartyIDs) SetNstdPtys2SubGrp(v nstdptys2subgrp.NstdPtys2SubGrp) {
	m.NstdPtys2SubGrp = &v
}

//NestedParties2 is a fix50sp1 Component
type NestedParties2 struct {
	//NoNested2PartyIDs is a non-required field for NestedParties2.
	NoNested2PartyIDs []NoNested2PartyIDs `fix:"756,omitempty"`
}

func (m *NestedParties2) SetNoNested2PartyIDs(v []NoNested2PartyIDs) { m.NoNested2PartyIDs = v }
