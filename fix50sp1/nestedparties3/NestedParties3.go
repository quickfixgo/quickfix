package nestedparties3

import (
	"github.com/quickfixgo/quickfix/fix50sp1/nstdptys3subgrp"
)

//NoNested3PartyIDs is a repeating group in NestedParties3
type NoNested3PartyIDs struct {
	//Nested3PartyID is a non-required field for NoNested3PartyIDs.
	Nested3PartyID *string `fix:"949"`
	//Nested3PartyIDSource is a non-required field for NoNested3PartyIDs.
	Nested3PartyIDSource *string `fix:"950"`
	//Nested3PartyRole is a non-required field for NoNested3PartyIDs.
	Nested3PartyRole *int `fix:"951"`
	//NstdPtys3SubGrp Component
	NstdPtys3SubGrp nstdptys3subgrp.Component
}

//Component is a fix50sp1 NestedParties3 Component
type Component struct {
	//NoNested3PartyIDs is a non-required field for NestedParties3.
	NoNested3PartyIDs []NoNested3PartyIDs `fix:"948,omitempty"`
}

func New() *Component { return new(Component) }
