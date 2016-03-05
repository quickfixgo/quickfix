package nestedparties3

import (
	"github.com/quickfixgo/quickfix/fix50/nstdptys3subgrp"
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
	nstdptys3subgrp.NstdPtys3SubGrp
}

//NestedParties3 is a fix50 Component
type NestedParties3 struct {
	//NoNested3PartyIDs is a non-required field for NestedParties3.
	NoNested3PartyIDs []NoNested3PartyIDs `fix:"948,omitempty"`
}

func (m *NestedParties3) SetNoNested3PartyIDs(v []NoNested3PartyIDs) { m.NoNested3PartyIDs = v }
