package rootparties

import (
	"github.com/quickfixgo/quickfix/fix50sp1/rootsubparties"
)

//NoRootPartyIDs is a repeating group in RootParties
type NoRootPartyIDs struct {
	//RootPartyID is a non-required field for NoRootPartyIDs.
	RootPartyID *string `fix:"1117"`
	//RootPartyIDSource is a non-required field for NoRootPartyIDs.
	RootPartyIDSource *string `fix:"1118"`
	//RootPartyRole is a non-required field for NoRootPartyIDs.
	RootPartyRole *int `fix:"1119"`
	//RootSubParties Component
	RootSubParties rootsubparties.Component
}

//Component is a fix50sp1 RootParties Component
type Component struct {
	//NoRootPartyIDs is a non-required field for RootParties.
	NoRootPartyIDs []NoRootPartyIDs `fix:"1116,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRootPartyIDs(v []NoRootPartyIDs) { m.NoRootPartyIDs = v }
