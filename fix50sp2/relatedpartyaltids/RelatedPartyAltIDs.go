package relatedpartyaltids

import (
	"github.com/quickfixgo/quickfix/fix50sp2/relatedaltptyssubgrp"
)

//New returns an initialized RelatedPartyAltIDs instance
func New() *RelatedPartyAltIDs {
	var m RelatedPartyAltIDs
	return &m
}

//NoRelatedPartyAltIDs is a repeating group in RelatedPartyAltIDs
type NoRelatedPartyAltIDs struct {
	//RelatedPartyAltID is a non-required field for NoRelatedPartyAltIDs.
	RelatedPartyAltID *string `fix:"1570"`
	//RelatedPartyAltIDSource is a non-required field for NoRelatedPartyAltIDs.
	RelatedPartyAltIDSource *string `fix:"1571"`
	//RelatedAltPtysSubGrp is a non-required component for NoRelatedPartyAltIDs.
	RelatedAltPtysSubGrp *relatedaltptyssubgrp.RelatedAltPtysSubGrp
}

//NewNoRelatedPartyAltIDs returns an initialized NoRelatedPartyAltIDs instance
func NewNoRelatedPartyAltIDs() *NoRelatedPartyAltIDs {
	var m NoRelatedPartyAltIDs
	return &m
}

func (m *NoRelatedPartyAltIDs) SetRelatedPartyAltID(v string)       { m.RelatedPartyAltID = &v }
func (m *NoRelatedPartyAltIDs) SetRelatedPartyAltIDSource(v string) { m.RelatedPartyAltIDSource = &v }
func (m *NoRelatedPartyAltIDs) SetRelatedAltPtysSubGrp(v relatedaltptyssubgrp.RelatedAltPtysSubGrp) {
	m.RelatedAltPtysSubGrp = &v
}

//RelatedPartyAltIDs is a fix50sp2 Component
type RelatedPartyAltIDs struct {
	//NoRelatedPartyAltIDs is a non-required field for RelatedPartyAltIDs.
	NoRelatedPartyAltIDs []NoRelatedPartyAltIDs `fix:"1569,omitempty"`
}

func (m *RelatedPartyAltIDs) SetNoRelatedPartyAltIDs(v []NoRelatedPartyAltIDs) {
	m.NoRelatedPartyAltIDs = v
}
