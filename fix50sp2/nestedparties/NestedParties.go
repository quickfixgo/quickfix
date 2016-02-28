package nestedparties

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nstdptyssubgrp"
)

//NoNestedPartyIDs is a repeating group in NestedParties
type NoNestedPartyIDs struct {
	//NestedPartyID is a non-required field for NoNestedPartyIDs.
	NestedPartyID *string `fix:"524"`
	//NestedPartyIDSource is a non-required field for NoNestedPartyIDs.
	NestedPartyIDSource *string `fix:"525"`
	//NestedPartyRole is a non-required field for NoNestedPartyIDs.
	NestedPartyRole *int `fix:"538"`
	//NstdPtysSubGrp Component
	NstdPtysSubGrp nstdptyssubgrp.Component
}

//Component is a fix50sp2 NestedParties Component
type Component struct {
	//NoNestedPartyIDs is a non-required field for NestedParties.
	NoNestedPartyIDs []NoNestedPartyIDs `fix:"539,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoNestedPartyIDs(v []NoNestedPartyIDs) { m.NoNestedPartyIDs = v }
