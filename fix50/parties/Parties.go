package parties

import (
	"github.com/quickfixgo/quickfix/fix50/ptyssubgrp"
)

func New() *Parties {
	var m Parties
	return &m
}

//NoPartyIDs is a repeating group in Parties
type NoPartyIDs struct {
	//PartyID is a non-required field for NoPartyIDs.
	PartyID *string `fix:"448"`
	//PartyIDSource is a non-required field for NoPartyIDs.
	PartyIDSource *string `fix:"447"`
	//PartyRole is a non-required field for NoPartyIDs.
	PartyRole *int `fix:"452"`
	//PtysSubGrp is a non-required component for NoPartyIDs.
	PtysSubGrp *ptyssubgrp.PtysSubGrp
}

func (m *NoPartyIDs) SetPartyID(v string)                   { m.PartyID = &v }
func (m *NoPartyIDs) SetPartyIDSource(v string)             { m.PartyIDSource = &v }
func (m *NoPartyIDs) SetPartyRole(v int)                    { m.PartyRole = &v }
func (m *NoPartyIDs) SetPtysSubGrp(v ptyssubgrp.PtysSubGrp) { m.PtysSubGrp = &v }

//Parties is a fix50 Component
type Parties struct {
	//NoPartyIDs is a non-required field for Parties.
	NoPartyIDs []NoPartyIDs `fix:"453,omitempty"`
}

func (m *Parties) SetNoPartyIDs(v []NoPartyIDs) { m.NoPartyIDs = v }
