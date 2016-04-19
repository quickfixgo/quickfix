package relationshiprisksecaltidgrp

//New returns an initialized RelationshipRiskSecAltIDGrp instance
func New() *RelationshipRiskSecAltIDGrp {
	var m RelationshipRiskSecAltIDGrp
	return &m
}

//NoRelationshipRiskSecurityAltID is a repeating group in RelationshipRiskSecAltIDGrp
type NoRelationshipRiskSecurityAltID struct {
	//RelationshipRiskSecurityAltID is a non-required field for NoRelationshipRiskSecurityAltID.
	RelationshipRiskSecurityAltID *string `fix:"1594"`
	//RelationshipRiskSecurityAltIDSource is a non-required field for NoRelationshipRiskSecurityAltID.
	RelationshipRiskSecurityAltIDSource *string `fix:"1595"`
}

//NewNoRelationshipRiskSecurityAltID returns an initialized NoRelationshipRiskSecurityAltID instance
func NewNoRelationshipRiskSecurityAltID() *NoRelationshipRiskSecurityAltID {
	var m NoRelationshipRiskSecurityAltID
	return &m
}

func (m *NoRelationshipRiskSecurityAltID) SetRelationshipRiskSecurityAltID(v string) {
	m.RelationshipRiskSecurityAltID = &v
}
func (m *NoRelationshipRiskSecurityAltID) SetRelationshipRiskSecurityAltIDSource(v string) {
	m.RelationshipRiskSecurityAltIDSource = &v
}

//RelationshipRiskSecAltIDGrp is a fix50sp2 Component
type RelationshipRiskSecAltIDGrp struct {
	//NoRelationshipRiskSecurityAltID is a non-required field for RelationshipRiskSecAltIDGrp.
	NoRelationshipRiskSecurityAltID []NoRelationshipRiskSecurityAltID `fix:"1593,omitempty"`
}

func (m *RelationshipRiskSecAltIDGrp) SetNoRelationshipRiskSecurityAltID(v []NoRelationshipRiskSecurityAltID) {
	m.NoRelationshipRiskSecurityAltID = v
}
