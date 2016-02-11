package parties

import (
	"github.com/quickfixgo/quickfix/fix50sp1/ptyssubgrp"
)

//NoPartyIDs is a repeating group in Parties
type NoPartyIDs struct {
	//PartyID is a non-required field for NoPartyIDs.
	PartyID *string `fix:"448"`
	//PartyIDSource is a non-required field for NoPartyIDs.
	PartyIDSource *string `fix:"447"`
	//PartyRole is a non-required field for NoPartyIDs.
	PartyRole *int `fix:"452"`
	//PtysSubGrp Component
	PtysSubGrp ptyssubgrp.Component
}

//Component is a fix50sp1 Parties Component
type Component struct {
	//NoPartyIDs is a non-required field for Parties.
	NoPartyIDs []NoPartyIDs `fix:"453,omitempty"`
}

func New() *Component { return new(Component) }
