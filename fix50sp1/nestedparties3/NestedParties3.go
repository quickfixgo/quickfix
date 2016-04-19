package nestedparties3

import (
	"github.com/quickfixgo/quickfix/fix50sp1/nstdptys3subgrp"
)

//New returns an initialized NestedParties3 instance
func New() *NestedParties3 {
	var m NestedParties3
	return &m
}

//NoNested3PartyIDs is a repeating group in NestedParties3
type NoNested3PartyIDs struct {
	//Nested3PartyID is a non-required field for NoNested3PartyIDs.
	Nested3PartyID *string `fix:"949"`
	//Nested3PartyIDSource is a non-required field for NoNested3PartyIDs.
	Nested3PartyIDSource *string `fix:"950"`
	//Nested3PartyRole is a non-required field for NoNested3PartyIDs.
	Nested3PartyRole *int `fix:"951"`
	//NstdPtys3SubGrp is a non-required component for NoNested3PartyIDs.
	NstdPtys3SubGrp *nstdptys3subgrp.NstdPtys3SubGrp
}

//NewNoNested3PartyIDs returns an initialized NoNested3PartyIDs instance
func NewNoNested3PartyIDs() *NoNested3PartyIDs {
	var m NoNested3PartyIDs
	return &m
}

func (m *NoNested3PartyIDs) SetNested3PartyID(v string)       { m.Nested3PartyID = &v }
func (m *NoNested3PartyIDs) SetNested3PartyIDSource(v string) { m.Nested3PartyIDSource = &v }
func (m *NoNested3PartyIDs) SetNested3PartyRole(v int)        { m.Nested3PartyRole = &v }
func (m *NoNested3PartyIDs) SetNstdPtys3SubGrp(v nstdptys3subgrp.NstdPtys3SubGrp) {
	m.NstdPtys3SubGrp = &v
}

//NestedParties3 is a fix50sp1 Component
type NestedParties3 struct {
	//NoNested3PartyIDs is a non-required field for NestedParties3.
	NoNested3PartyIDs []NoNested3PartyIDs `fix:"948,omitempty"`
}

func (m *NestedParties3) SetNoNested3PartyIDs(v []NoNested3PartyIDs) { m.NoNested3PartyIDs = v }
