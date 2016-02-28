package relatedptyssubgrp

//NoRelatedPartySubIDs is a repeating group in RelatedPtysSubGrp
type NoRelatedPartySubIDs struct {
	//RelatedPartySubID is a non-required field for NoRelatedPartySubIDs.
	RelatedPartySubID *string `fix:"1567"`
	//RelatedPartySubIDType is a non-required field for NoRelatedPartySubIDs.
	RelatedPartySubIDType *int `fix:"1568"`
}

//Component is a fix50sp2 RelatedPtysSubGrp Component
type Component struct {
	//NoRelatedPartySubIDs is a non-required field for RelatedPtysSubGrp.
	NoRelatedPartySubIDs []NoRelatedPartySubIDs `fix:"1566,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRelatedPartySubIDs(v []NoRelatedPartySubIDs) { m.NoRelatedPartySubIDs = v }
