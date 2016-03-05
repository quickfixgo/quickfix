package relatedaltptyssubgrp

//NoRelatedPartyAltSubIDs is a repeating group in RelatedAltPtysSubGrp
type NoRelatedPartyAltSubIDs struct {
	//RelatedPartyAltSubID is a non-required field for NoRelatedPartyAltSubIDs.
	RelatedPartyAltSubID *string `fix:"1573"`
	//RelatedPartyAltSubIDType is a non-required field for NoRelatedPartyAltSubIDs.
	RelatedPartyAltSubIDType *int `fix:"1574"`
}

//RelatedAltPtysSubGrp is a fix50sp2 Component
type RelatedAltPtysSubGrp struct {
	//NoRelatedPartyAltSubIDs is a non-required field for RelatedAltPtysSubGrp.
	NoRelatedPartyAltSubIDs []NoRelatedPartyAltSubIDs `fix:"1572,omitempty"`
}

func (m *RelatedAltPtysSubGrp) SetNoRelatedPartyAltSubIDs(v []NoRelatedPartyAltSubIDs) {
	m.NoRelatedPartyAltSubIDs = v
}
