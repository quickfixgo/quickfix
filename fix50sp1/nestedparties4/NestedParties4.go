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
	NstdPtys4SubGrp nstdptys4subgrp.Component
}

//Component is a fix50sp1 NestedParties4 Component
type Component struct {
	//NoNested4PartyIDs is a non-required field for NestedParties4.
	NoNested4PartyIDs []NoNested4PartyIDs `fix:"1414,omitempty"`
}

func New() *Component { return new(Component) }
