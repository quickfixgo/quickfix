package nestedparties4

import (
	"github.com/quickfixgo/quickfix/fix50sp1/nstdptys4subgrp"
)

//NoNested4PartyIDs is a repeating group in NestedParties4
type NoNested4PartyIDs struct {
	//Nested4PartyID is a non-required field for NoNested4PartyIDs.
	Nested4PartyID *string `fix:"1415"`
	//Nested4PartyIDSource is a non-required field for NoNested4PartyIDs.
	Nested4PartyIDSource *string `fix:"1416"`
	//Nested4PartyRole is a non-required field for NoNested4PartyIDs.
	Nested4PartyRole *int `fix:"1417"`
	//NstdPtys4SubGrp Component
	nstdptys4subgrp.NstdPtys4SubGrp
}

//NestedParties4 is a fix50sp1 Component
type NestedParties4 struct {
	//NoNested4PartyIDs is a non-required field for NestedParties4.
	NoNested4PartyIDs []NoNested4PartyIDs `fix:"1414,omitempty"`
}

func (m *NestedParties4) SetNoNested4PartyIDs(v []NoNested4PartyIDs) { m.NoNested4PartyIDs = v }
