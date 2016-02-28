package nestedparties2

import (
	"github.com/quickfixgo/quickfix/fix50sp1/nstdptys2subgrp"
)

//NoNested2PartyIDs is a repeating group in NestedParties2
type NoNested2PartyIDs struct {
	//Nested2PartyID is a non-required field for NoNested2PartyIDs.
	Nested2PartyID *string `fix:"757"`
	//Nested2PartyIDSource is a non-required field for NoNested2PartyIDs.
	Nested2PartyIDSource *string `fix:"758"`
	//Nested2PartyRole is a non-required field for NoNested2PartyIDs.
	Nested2PartyRole *int `fix:"759"`
	//NstdPtys2SubGrp Component
	NstdPtys2SubGrp nstdptys2subgrp.Component
}

//Component is a fix50sp1 NestedParties2 Component
type Component struct {
	//NoNested2PartyIDs is a non-required field for NestedParties2.
	NoNested2PartyIDs []NoNested2PartyIDs `fix:"756,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoNested2PartyIDs(v []NoNested2PartyIDs) { m.NoNested2PartyIDs = v }
