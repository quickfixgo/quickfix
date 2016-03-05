package relationshipriskwarninglevels

//NoRelationshipRiskWarningLevels is a repeating group in RelationshipRiskWarningLevels
type NoRelationshipRiskWarningLevels struct {
	//RelationshipRiskWarningLevelPercent is a non-required field for NoRelationshipRiskWarningLevels.
	RelationshipRiskWarningLevelPercent *float64 `fix:"1614"`
	//RelationshipRiskWarningLevelName is a non-required field for NoRelationshipRiskWarningLevels.
	RelationshipRiskWarningLevelName *string `fix:"1615"`
}

//RelationshipRiskWarningLevels is a fix50sp2 Component
type RelationshipRiskWarningLevels struct {
	//NoRelationshipRiskWarningLevels is a non-required field for RelationshipRiskWarningLevels.
	NoRelationshipRiskWarningLevels []NoRelationshipRiskWarningLevels `fix:"1613,omitempty"`
}

func (m *RelationshipRiskWarningLevels) SetNoRelationshipRiskWarningLevels(v []NoRelationshipRiskWarningLevels) {
	m.NoRelationshipRiskWarningLevels = v
}
