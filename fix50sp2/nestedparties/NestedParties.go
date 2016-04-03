package nestedparties

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nstdptyssubgrp"
)

func New() *NestedParties {
	var m NestedParties
	return &m
}

//NoNestedPartyIDs is a repeating group in NestedParties
type NoNestedPartyIDs struct {
	//NestedPartyID is a non-required field for NoNestedPartyIDs.
	NestedPartyID *string `fix:"524"`
	//NestedPartyIDSource is a non-required field for NoNestedPartyIDs.
	NestedPartyIDSource *string `fix:"525"`
	//NestedPartyRole is a non-required field for NoNestedPartyIDs.
	NestedPartyRole *int `fix:"538"`
	//NstdPtysSubGrp is a non-required component for NoNestedPartyIDs.
	NstdPtysSubGrp *nstdptyssubgrp.NstdPtysSubGrp
}

//NestedParties is a fix50sp2 Component
type NestedParties struct {
	//NoNestedPartyIDs is a non-required field for NestedParties.
	NoNestedPartyIDs []NoNestedPartyIDs `fix:"539,omitempty"`
}

func (m *NestedParties) SetNoNestedPartyIDs(v []NoNestedPartyIDs) { m.NoNestedPartyIDs = v }
