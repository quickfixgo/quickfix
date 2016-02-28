package relatedpartyaltids

import (
	"github.com/quickfixgo/quickfix/fix50sp2/relatedaltptyssubgrp"
)

//NoRelatedPartyAltIDs is a repeating group in RelatedPartyAltIDs
type NoRelatedPartyAltIDs struct {
	//RelatedPartyAltID is a non-required field for NoRelatedPartyAltIDs.
	RelatedPartyAltID *string `fix:"1570"`
	//RelatedPartyAltIDSource is a non-required field for NoRelatedPartyAltIDs.
	RelatedPartyAltIDSource *string `fix:"1571"`
	//RelatedAltPtysSubGrp Component
	RelatedAltPtysSubGrp relatedaltptyssubgrp.Component
}

//Component is a fix50sp2 RelatedPartyAltIDs Component
type Component struct {
	//NoRelatedPartyAltIDs is a non-required field for RelatedPartyAltIDs.
	NoRelatedPartyAltIDs []NoRelatedPartyAltIDs `fix:"1569,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRelatedPartyAltIDs(v []NoRelatedPartyAltIDs) { m.NoRelatedPartyAltIDs = v }
