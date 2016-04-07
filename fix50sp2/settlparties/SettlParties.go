package settlparties

import (
	"github.com/quickfixgo/quickfix/fix50sp2/settlptyssubgrp"
)

func New() *SettlParties {
	var m SettlParties
	return &m
}

//NoSettlPartyIDs is a repeating group in SettlParties
type NoSettlPartyIDs struct {
	//SettlPartyID is a non-required field for NoSettlPartyIDs.
	SettlPartyID *string `fix:"782"`
	//SettlPartyIDSource is a non-required field for NoSettlPartyIDs.
	SettlPartyIDSource *string `fix:"783"`
	//SettlPartyRole is a non-required field for NoSettlPartyIDs.
	SettlPartyRole *int `fix:"784"`
	//SettlPtysSubGrp is a non-required component for NoSettlPartyIDs.
	SettlPtysSubGrp *settlptyssubgrp.SettlPtysSubGrp
}

func (m *NoSettlPartyIDs) SetSettlPartyID(v string)       { m.SettlPartyID = &v }
func (m *NoSettlPartyIDs) SetSettlPartyIDSource(v string) { m.SettlPartyIDSource = &v }
func (m *NoSettlPartyIDs) SetSettlPartyRole(v int)        { m.SettlPartyRole = &v }
func (m *NoSettlPartyIDs) SetSettlPtysSubGrp(v settlptyssubgrp.SettlPtysSubGrp) {
	m.SettlPtysSubGrp = &v
}

//SettlParties is a fix50sp2 Component
type SettlParties struct {
	//NoSettlPartyIDs is a non-required field for SettlParties.
	NoSettlPartyIDs []NoSettlPartyIDs `fix:"781,omitempty"`
}

func (m *SettlParties) SetNoSettlPartyIDs(v []NoSettlPartyIDs) { m.NoSettlPartyIDs = v }
