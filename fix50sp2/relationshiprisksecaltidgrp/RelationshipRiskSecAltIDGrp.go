package relationshiprisksecaltidgrp

//NoRelationshipRiskSecurityAltID is a repeating group in RelationshipRiskSecAltIDGrp
type NoRelationshipRiskSecurityAltID struct {
	//RelationshipRiskSecurityAltID is a non-required field for NoRelationshipRiskSecurityAltID.
	RelationshipRiskSecurityAltID *string `fix:"1594"`
	//RelationshipRiskSecurityAltIDSource is a non-required field for NoRelationshipRiskSecurityAltID.
	RelationshipRiskSecurityAltIDSource *string `fix:"1595"`
}

//Component is a fix50sp2 RelationshipRiskSecAltIDGrp Component
type Component struct {
	//NoRelationshipRiskSecurityAltID is a non-required field for RelationshipRiskSecAltIDGrp.
	NoRelationshipRiskSecurityAltID []NoRelationshipRiskSecurityAltID `fix:"1593,omitempty"`
}

func New() *Component { return new(Component) }
