package rootparties

import (
	"github.com/quickfixgo/quickfix/fix50sp2/rootsubparties"
)

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

//RootParties is a fix50sp2 Component
type RootParties struct {
	//NoRootPartyIDs is a non-required field for RootParties.
	NoRootPartyIDs []NoRootPartyIDs `fix:"1116,omitempty"`
}

func (m *RootParties) SetNoRootPartyIDs(v []NoRootPartyIDs) { m.NoRootPartyIDs = v }
