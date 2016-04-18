package relatedaltptyssubgrp

//New returns an initialized RelatedAltPtysSubGrp instance
func New() *RelatedAltPtysSubGrp {
	var m RelatedAltPtysSubGrp
	return &m
}

//NoRelatedPartyAltSubIDs is a repeating group in RelatedAltPtysSubGrp
type NoRelatedPartyAltSubIDs struct {
	//RelatedPartyAltSubID is a non-required field for NoRelatedPartyAltSubIDs.
	RelatedPartyAltSubID *string `fix:"1573"`
	//RelatedPartyAltSubIDType is a non-required field for NoRelatedPartyAltSubIDs.
	RelatedPartyAltSubIDType *int `fix:"1574"`
}

//NewNoRelatedPartyAltSubIDs returns an initialized NoRelatedPartyAltSubIDs instance
func NewNoRelatedPartyAltSubIDs() *NoRelatedPartyAltSubIDs {
	var m NoRelatedPartyAltSubIDs
	return &m
}

func (m *NoRelatedPartyAltSubIDs) SetRelatedPartyAltSubID(v string)  { m.RelatedPartyAltSubID = &v }
func (m *NoRelatedPartyAltSubIDs) SetRelatedPartyAltSubIDType(v int) { m.RelatedPartyAltSubIDType = &v }

//RelatedAltPtysSubGrp is a fix50sp2 Component
type RelatedAltPtysSubGrp struct {
	//NoRelatedPartyAltSubIDs is a non-required field for RelatedAltPtysSubGrp.
	NoRelatedPartyAltSubIDs []NoRelatedPartyAltSubIDs `fix:"1572,omitempty"`
}

func (m *RelatedAltPtysSubGrp) SetNoRelatedPartyAltSubIDs(v []NoRelatedPartyAltSubIDs) {
	m.NoRelatedPartyAltSubIDs = v
}
