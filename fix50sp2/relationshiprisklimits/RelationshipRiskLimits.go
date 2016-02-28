package relationshiprisklimits

import (
	"github.com/quickfixgo/quickfix/fix50sp2/relationshipriskinstrumentscope"
	"github.com/quickfixgo/quickfix/fix50sp2/relationshipriskwarninglevels"
)

//NoRelationshipRiskLimits is a repeating group in RelationshipRiskLimits
type NoRelationshipRiskLimits struct {
	//RelationshipRiskLimitType is a non-required field for NoRelationshipRiskLimits.
	RelationshipRiskLimitType *int `fix:"1583"`
	//RelationshipRiskLimitAmount is a non-required field for NoRelationshipRiskLimits.
	RelationshipRiskLimitAmount *float64 `fix:"1584"`
	//RelationshipRiskLimitCurrency is a non-required field for NoRelationshipRiskLimits.
	RelationshipRiskLimitCurrency *string `fix:"1585"`
	//RelationshipRiskLimitPlatform is a non-required field for NoRelationshipRiskLimits.
	RelationshipRiskLimitPlatform *string `fix:"1586"`
	//RelationshipRiskInstrumentScope Component
	RelationshipRiskInstrumentScope relationshipriskinstrumentscope.Component
	//RelationshipRiskWarningLevels Component
	RelationshipRiskWarningLevels relationshipriskwarninglevels.Component
}

//Component is a fix50sp2 RelationshipRiskLimits Component
type Component struct {
	//NoRelationshipRiskLimits is a non-required field for RelationshipRiskLimits.
	NoRelationshipRiskLimits []NoRelationshipRiskLimits `fix:"1582,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRelationshipRiskLimits(v []NoRelationshipRiskLimits) {
	m.NoRelationshipRiskLimits = v
}
