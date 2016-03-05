package settlparties

import (
	"github.com/quickfixgo/quickfix/fix50sp2/settlptyssubgrp"
)

//NoSettlPartyIDs is a repeating group in SettlParties
type NoSettlPartyIDs struct {
	//SettlPartyID is a non-required field for NoSettlPartyIDs.
	SettlPartyID *string `fix:"782"`
	//SettlPartyIDSource is a non-required field for NoSettlPartyIDs.
	SettlPartyIDSource *string `fix:"783"`
	//SettlPartyRole is a non-required field for NoSettlPartyIDs.
	SettlPartyRole *int `fix:"784"`
	//SettlPtysSubGrp Component
	settlptyssubgrp.SettlPtysSubGrp
}

//SettlParties is a fix50sp2 Component
type SettlParties struct {
	//NoSettlPartyIDs is a non-required field for SettlParties.
	NoSettlPartyIDs []NoSettlPartyIDs `fix:"781,omitempty"`
}

func (m *SettlParties) SetNoSettlPartyIDs(v []NoSettlPartyIDs) { m.NoSettlPartyIDs = v }
