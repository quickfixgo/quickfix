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
	relatedaltptyssubgrp.RelatedAltPtysSubGrp
}

//RelatedPartyAltIDs is a fix50sp2 Component
type RelatedPartyAltIDs struct {
	//NoRelatedPartyAltIDs is a non-required field for RelatedPartyAltIDs.
	NoRelatedPartyAltIDs []NoRelatedPartyAltIDs `fix:"1569,omitempty"`
}

func (m *RelatedPartyAltIDs) SetNoRelatedPartyAltIDs(v []NoRelatedPartyAltIDs) {
	m.NoRelatedPartyAltIDs = v
}
