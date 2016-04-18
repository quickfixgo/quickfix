package partyaltids

import (
	"github.com/quickfixgo/quickfix/fix50sp2/altptyssubgrp"
)

//New returns an initialized PartyAltIDs instance
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

//NewNoPartyAltIDs returns an initialized NoPartyAltIDs instance
func NewNoPartyAltIDs() *NoPartyAltIDs {
	var m NoPartyAltIDs
	return &m
}

func (m *NoPartyAltIDs) SetPartyAltID(v string)                         { m.PartyAltID = &v }
func (m *NoPartyAltIDs) SetPartyAltIDSource(v string)                   { m.PartyAltIDSource = &v }
func (m *NoPartyAltIDs) SetAltPtysSubGrp(v altptyssubgrp.AltPtysSubGrp) { m.AltPtysSubGrp = &v }

//PartyAltIDs is a fix50sp2 Component
type PartyAltIDs struct {
	//NoPartyAltIDs is a non-required field for PartyAltIDs.
	NoPartyAltIDs []NoPartyAltIDs `fix:"1516,omitempty"`
}

func (m *PartyAltIDs) SetNoPartyAltIDs(v []NoPartyAltIDs) { m.NoPartyAltIDs = v }
