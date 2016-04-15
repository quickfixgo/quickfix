package relationshiprisklimits

import (
	"github.com/quickfixgo/quickfix/fix50sp2/relationshipriskinstrumentscope"
	"github.com/quickfixgo/quickfix/fix50sp2/relationshipriskwarninglevels"
)

func New() *RelationshipRiskLimits {
	var m RelationshipRiskLimits
	return &m
}

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
	//RelationshipRiskInstrumentScope is a non-required component for NoRelationshipRiskLimits.
	RelationshipRiskInstrumentScope *relationshipriskinstrumentscope.RelationshipRiskInstrumentScope
	//RelationshipRiskWarningLevels is a non-required component for NoRelationshipRiskLimits.
	RelationshipRiskWarningLevels *relationshipriskwarninglevels.RelationshipRiskWarningLevels
}

func (m *NoRelationshipRiskLimits) SetRelationshipRiskLimitType(v int) {
	m.RelationshipRiskLimitType = &v
}
func (m *NoRelationshipRiskLimits) SetRelationshipRiskLimitAmount(v float64) {
	m.RelationshipRiskLimitAmount = &v
}
func (m *NoRelationshipRiskLimits) SetRelationshipRiskLimitCurrency(v string) {
	m.RelationshipRiskLimitCurrency = &v
}
func (m *NoRelationshipRiskLimits) SetRelationshipRiskLimitPlatform(v string) {
	m.RelationshipRiskLimitPlatform = &v
}
func (m *NoRelationshipRiskLimits) SetRelationshipRiskInstrumentScope(v relationshipriskinstrumentscope.RelationshipRiskInstrumentScope) {
	m.RelationshipRiskInstrumentScope = &v
}
func (m *NoRelationshipRiskLimits) SetRelationshipRiskWarningLevels(v relationshipriskwarninglevels.RelationshipRiskWarningLevels) {
	m.RelationshipRiskWarningLevels = &v
}

//RelationshipRiskLimits is a fix50sp2 Component
type RelationshipRiskLimits struct {
	//NoRelationshipRiskLimits is a non-required field for RelationshipRiskLimits.
	NoRelationshipRiskLimits []NoRelationshipRiskLimits `fix:"1582,omitempty"`
}

func (m *RelationshipRiskLimits) SetNoRelationshipRiskLimits(v []NoRelationshipRiskLimits) {
	m.NoRelationshipRiskLimits = v
}
