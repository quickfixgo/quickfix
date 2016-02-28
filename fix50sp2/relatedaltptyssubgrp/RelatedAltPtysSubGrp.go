package relatedaltptyssubgrp

//NoRelatedPartyAltSubIDs is a repeating group in RelatedAltPtysSubGrp
type NoRelatedPartyAltSubIDs struct {
	//RelatedPartyAltSubID is a non-required field for NoRelatedPartyAltSubIDs.
	RelatedPartyAltSubID *string `fix:"1573"`
	//RelatedPartyAltSubIDType is a non-required field for NoRelatedPartyAltSubIDs.
	RelatedPartyAltSubIDType *int `fix:"1574"`
}

//Component is a fix50sp2 RelatedAltPtysSubGrp Component
type Component struct {
	//NoRelatedPartyAltSubIDs is a non-required field for RelatedAltPtysSubGrp.
	NoRelatedPartyAltSubIDs []NoRelatedPartyAltSubIDs `fix:"1572,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRelatedPartyAltSubIDs(v []NoRelatedPartyAltSubIDs) {
	m.NoRelatedPartyAltSubIDs = v
}
