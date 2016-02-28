package relatedcontextparties

import (
	"github.com/quickfixgo/quickfix/fix50sp2/relatedcontextptyssubgrp"
)

//NoRelatedContextPartyIDs is a repeating group in RelatedContextParties
type NoRelatedContextPartyIDs struct {
	//RelatedContextPartyID is a non-required field for NoRelatedContextPartyIDs.
	RelatedContextPartyID *string `fix:"1576"`
	//RelatedContextPartyIDSource is a non-required field for NoRelatedContextPartyIDs.
	RelatedContextPartyIDSource *string `fix:"1577"`
	//RelatedContextPartyRole is a non-required field for NoRelatedContextPartyIDs.
	RelatedContextPartyRole *int `fix:"1578"`
	//RelatedContextPtysSubGrp Component
	RelatedContextPtysSubGrp relatedcontextptyssubgrp.Component
}

//Component is a fix50sp2 RelatedContextParties Component
type Component struct {
	//NoRelatedContextPartyIDs is a non-required field for RelatedContextParties.
	NoRelatedContextPartyIDs []NoRelatedContextPartyIDs `fix:"1575,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRelatedContextPartyIDs(v []NoRelatedContextPartyIDs) {
	m.NoRelatedContextPartyIDs = v
}
