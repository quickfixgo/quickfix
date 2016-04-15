package strikerules

import (
	"github.com/quickfixgo/quickfix/fix50sp1/maturityrules"
)

func New() *StrikeRules {
	var m StrikeRules
	return &m
}

//NoStrikeRules is a repeating group in StrikeRules
type NoStrikeRules struct {
	//StrikeRuleID is a non-required field for NoStrikeRules.
	StrikeRuleID *string `fix:"1223"`
	//StartStrikePxRange is a non-required field for NoStrikeRules.
	StartStrikePxRange *float64 `fix:"1202"`
	//EndStrikePxRange is a non-required field for NoStrikeRules.
	EndStrikePxRange *float64 `fix:"1203"`
	//StrikeIncrement is a non-required field for NoStrikeRules.
	StrikeIncrement *float64 `fix:"1204"`
	//StrikeExerciseStyle is a non-required field for NoStrikeRules.
	StrikeExerciseStyle *int `fix:"1304"`
	//MaturityRules is a non-required component for NoStrikeRules.
	MaturityRules *maturityrules.MaturityRules
}

func (m *NoStrikeRules) SetStrikeRuleID(v string)                       { m.StrikeRuleID = &v }
func (m *NoStrikeRules) SetStartStrikePxRange(v float64)                { m.StartStrikePxRange = &v }
func (m *NoStrikeRules) SetEndStrikePxRange(v float64)                  { m.EndStrikePxRange = &v }
func (m *NoStrikeRules) SetStrikeIncrement(v float64)                   { m.StrikeIncrement = &v }
func (m *NoStrikeRules) SetStrikeExerciseStyle(v int)                   { m.StrikeExerciseStyle = &v }
func (m *NoStrikeRules) SetMaturityRules(v maturityrules.MaturityRules) { m.MaturityRules = &v }

//StrikeRules is a fix50sp1 Component
type StrikeRules struct {
	//NoStrikeRules is a non-required field for StrikeRules.
	NoStrikeRules []NoStrikeRules `fix:"1201,omitempty"`
}

func (m *StrikeRules) SetNoStrikeRules(v []NoStrikeRules) { m.NoStrikeRules = v }
