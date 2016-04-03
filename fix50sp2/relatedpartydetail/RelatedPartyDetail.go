package relatedpartydetail

import (
	"github.com/quickfixgo/quickfix/fix50sp2/relatedcontextparties"
	"github.com/quickfixgo/quickfix/fix50sp2/relatedpartyaltids"
	"github.com/quickfixgo/quickfix/fix50sp2/relatedptyssubgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/relationshiprisklimits"
)

func New() *RelatedPartyDetail {
	var m RelatedPartyDetail
	return &m
}

//RelatedPartyDetail is a fix50sp2 Component
type RelatedPartyDetail struct {
	//RelatedPartyID is a non-required field for RelatedPartyDetail.
	RelatedPartyID *string `fix:"1563"`
	//RelatedPartyIDSource is a non-required field for RelatedPartyDetail.
	RelatedPartyIDSource *string `fix:"1564"`
	//RelatedPartyRole is a non-required field for RelatedPartyDetail.
	RelatedPartyRole *int `fix:"1565"`
	//RelatedPtysSubGrp is a non-required component for RelatedPartyDetail.
	RelatedPtysSubGrp *relatedptyssubgrp.RelatedPtysSubGrp
	//RelatedPartyAltIDs is a non-required component for RelatedPartyDetail.
	RelatedPartyAltIDs *relatedpartyaltids.RelatedPartyAltIDs
	//RelatedContextParties is a non-required component for RelatedPartyDetail.
	RelatedContextParties *relatedcontextparties.RelatedContextParties
	//RelationshipRiskLimits is a non-required component for RelatedPartyDetail.
	RelationshipRiskLimits *relationshiprisklimits.RelationshipRiskLimits
}

func (m *RelatedPartyDetail) SetRelatedPartyID(v string)       { m.RelatedPartyID = &v }
func (m *RelatedPartyDetail) SetRelatedPartyIDSource(v string) { m.RelatedPartyIDSource = &v }
func (m *RelatedPartyDetail) SetRelatedPartyRole(v int)        { m.RelatedPartyRole = &v }
func (m *RelatedPartyDetail) SetRelatedPtysSubGrp(v relatedptyssubgrp.RelatedPtysSubGrp) {
	m.RelatedPtysSubGrp = &v
}
func (m *RelatedPartyDetail) SetRelatedPartyAltIDs(v relatedpartyaltids.RelatedPartyAltIDs) {
	m.RelatedPartyAltIDs = &v
}
func (m *RelatedPartyDetail) SetRelatedContextParties(v relatedcontextparties.RelatedContextParties) {
	m.RelatedContextParties = &v
}
func (m *RelatedPartyDetail) SetRelationshipRiskLimits(v relationshiprisklimits.RelationshipRiskLimits) {
	m.RelationshipRiskLimits = &v
}
