package nestedparties4

import (
	"github.com/quickfixgo/quickfix/fix50sp1/nstdptys4subgrp"
)

func New() *NestedParties4 {
	var m NestedParties4
	return &m
}

//NoNested4PartyIDs is a repeating group in NestedParties4
type NoNested4PartyIDs struct {
	//Nested4PartyID is a non-required field for NoNested4PartyIDs.
	Nested4PartyID *string `fix:"1415"`
	//Nested4PartyIDSource is a non-required field for NoNested4PartyIDs.
	Nested4PartyIDSource *string `fix:"1416"`
	//Nested4PartyRole is a non-required field for NoNested4PartyIDs.
	Nested4PartyRole *int `fix:"1417"`
	//NstdPtys4SubGrp is a non-required component for NoNested4PartyIDs.
	NstdPtys4SubGrp *nstdptys4subgrp.NstdPtys4SubGrp
}

func (m *NoNested4PartyIDs) SetNested4PartyID(v string)       { m.Nested4PartyID = &v }
func (m *NoNested4PartyIDs) SetNested4PartyIDSource(v string) { m.Nested4PartyIDSource = &v }
func (m *NoNested4PartyIDs) SetNested4PartyRole(v int)        { m.Nested4PartyRole = &v }
func (m *NoNested4PartyIDs) SetNstdPtys4SubGrp(v nstdptys4subgrp.NstdPtys4SubGrp) {
	m.NstdPtys4SubGrp = &v
}

//NestedParties4 is a fix50sp1 Component
type NestedParties4 struct {
	//NoNested4PartyIDs is a non-required field for NestedParties4.
	NoNested4PartyIDs []NoNested4PartyIDs `fix:"1414,omitempty"`
}

func (m *NestedParties4) SetNoNested4PartyIDs(v []NoNested4PartyIDs) { m.NoNested4PartyIDs = v }
