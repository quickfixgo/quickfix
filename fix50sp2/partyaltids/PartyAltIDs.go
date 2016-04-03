package partyaltids

import (
	"github.com/quickfixgo/quickfix/fix50sp2/altptyssubgrp"
)

func New() *PartyAltIDs {
	var m PartyAltIDs
	return &m
}

//NoPartyAltIDs is a repeating group in PartyAltIDs
type NoPartyAltIDs struct {
	//PartyAltID is a non-required field for NoPartyAltIDs.
	PartyAltID *string `fix:"1517"`
	//PartyAltIDSource is a non-required field for NoPartyAltIDs.
	PartyAltIDSource *string `fix:"1518"`
	//AltPtysSubGrp is a non-required component for NoPartyAltIDs.
	AltPtysSubGrp *altptyssubgrp.AltPtysSubGrp
}

//PartyAltIDs is a fix50sp2 Component
type PartyAltIDs struct {
	//NoPartyAltIDs is a non-required field for PartyAltIDs.
	NoPartyAltIDs []NoPartyAltIDs `fix:"1516,omitempty"`
}

func (m *PartyAltIDs) SetNoPartyAltIDs(v []NoPartyAltIDs) { m.NoPartyAltIDs = v }
