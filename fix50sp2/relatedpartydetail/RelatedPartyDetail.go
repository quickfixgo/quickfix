package relatedpartydetail

import (
	"github.com/quickfixgo/quickfix/fix50sp2/relatedcontextparties"
	"github.com/quickfixgo/quickfix/fix50sp2/relatedpartyaltids"
	"github.com/quickfixgo/quickfix/fix50sp2/relatedptyssubgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/relationshiprisklimits"
)

//RelatedPartyDetail is a fix50sp2 Component
type RelatedPartyDetail struct {
	//RelatedPartyID is a non-required field for RelatedPartyDetail.
	RelatedPartyID *string `fix:"1563"`
	//RelatedPartyIDSource is a non-required field for RelatedPartyDetail.
	RelatedPartyIDSource *string `fix:"1564"`
	//RelatedPartyRole is a non-required field for RelatedPartyDetail.
	RelatedPartyRole *int `fix:"1565"`
	//RelatedPtysSubGrp Component
	relatedptyssubgrp.RelatedPtysSubGrp
	//RelatedPartyAltIDs Component
	relatedpartyaltids.RelatedPartyAltIDs
	//RelatedContextParties Component
	relatedcontextparties.RelatedContextParties
	//RelationshipRiskLimits Component
	relationshiprisklimits.RelationshipRiskLimits
}

func (m *RelatedPartyDetail) SetRelatedPartyID(v string)       { m.RelatedPartyID = &v }
func (m *RelatedPartyDetail) SetRelatedPartyIDSource(v string) { m.RelatedPartyIDSource = &v }
func (m *RelatedPartyDetail) SetRelatedPartyRole(v int)        { m.RelatedPartyRole = &v }
