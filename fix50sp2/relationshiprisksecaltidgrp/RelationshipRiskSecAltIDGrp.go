package relationshiprisksecaltidgrp

//NoRelationshipRiskSecurityAltID is a repeating group in RelationshipRiskSecAltIDGrp
type NoRelationshipRiskSecurityAltID struct {
	//RelationshipRiskSecurityAltID is a non-required field for NoRelationshipRiskSecurityAltID.
	RelationshipRiskSecurityAltID *string `fix:"1594"`
	//RelationshipRiskSecurityAltIDSource is a non-required field for NoRelationshipRiskSecurityAltID.
	RelationshipRiskSecurityAltIDSource *string `fix:"1595"`
}

//RelationshipRiskSecAltIDGrp is a fix50sp2 Component
type RelationshipRiskSecAltIDGrp struct {
	//NoRelationshipRiskSecurityAltID is a non-required field for RelationshipRiskSecAltIDGrp.
	NoRelationshipRiskSecurityAltID []NoRelationshipRiskSecurityAltID `fix:"1593,omitempty"`
}

func (m *RelationshipRiskSecAltIDGrp) SetNoRelationshipRiskSecurityAltID(v []NoRelationshipRiskSecurityAltID) {
	m.NoRelationshipRiskSecurityAltID = v
}
