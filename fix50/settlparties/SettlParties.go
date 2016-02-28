package settlparties

import (
	"github.com/quickfixgo/quickfix/fix50/settlptyssubgrp"
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
	SettlPtysSubGrp settlptyssubgrp.Component
}

//Component is a fix50 SettlParties Component
type Component struct {
	//NoSettlPartyIDs is a non-required field for SettlParties.
	NoSettlPartyIDs []NoSettlPartyIDs `fix:"781,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoSettlPartyIDs(v []NoSettlPartyIDs) { m.NoSettlPartyIDs = v }
