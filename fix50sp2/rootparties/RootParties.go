package rootparties

import (
	"github.com/quickfixgo/quickfix/fix50sp2/rootsubparties"
)

//New returns an initialized RootParties instance
func New() *RootParties {
	var m RootParties
	return &m
}

//NoRootPartyIDs is a repeating group in RootParties
type NoRootPartyIDs struct {
	//RootPartyID is a non-required field for NoRootPartyIDs.
	RootPartyID *string `fix:"1117"`
	//RootPartyIDSource is a non-required field for NoRootPartyIDs.
	RootPartyIDSource *string `fix:"1118"`
	//RootPartyRole is a non-required field for NoRootPartyIDs.
	RootPartyRole *int `fix:"1119"`
	//RootSubParties is a non-required component for NoRootPartyIDs.
	RootSubParties *rootsubparties.RootSubParties
}

//NewNoRootPartyIDs returns an initialized NoRootPartyIDs instance
func NewNoRootPartyIDs() *NoRootPartyIDs {
	var m NoRootPartyIDs
	return &m
}

func (m *NoRootPartyIDs) SetRootPartyID(v string)                           { m.RootPartyID = &v }
func (m *NoRootPartyIDs) SetRootPartyIDSource(v string)                     { m.RootPartyIDSource = &v }
func (m *NoRootPartyIDs) SetRootPartyRole(v int)                            { m.RootPartyRole = &v }
func (m *NoRootPartyIDs) SetRootSubParties(v rootsubparties.RootSubParties) { m.RootSubParties = &v }

//RootParties is a fix50sp2 Component
type RootParties struct {
	//NoRootPartyIDs is a non-required field for RootParties.
	NoRootPartyIDs []NoRootPartyIDs `fix:"1116,omitempty"`
}

func (m *RootParties) SetNoRootPartyIDs(v []NoRootPartyIDs) { m.NoRootPartyIDs = v }
