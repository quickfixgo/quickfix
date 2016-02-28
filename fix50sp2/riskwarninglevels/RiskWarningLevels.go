package riskwarninglevels

//NoRiskWarningLevels is a repeating group in RiskWarningLevels
type NoRiskWarningLevels struct {
	//RiskWarningLevelPercent is a non-required field for NoRiskWarningLevels.
	RiskWarningLevelPercent *float64 `fix:"1560"`
	//RiskWarningLevelName is a non-required field for NoRiskWarningLevels.
	RiskWarningLevelName *string `fix:"1561"`
}

//Component is a fix50sp2 RiskWarningLevels Component
type Component struct {
	//NoRiskWarningLevels is a non-required field for RiskWarningLevels.
	NoRiskWarningLevels []NoRiskWarningLevels `fix:"1559,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRiskWarningLevels(v []NoRiskWarningLevels) { m.NoRiskWarningLevels = v }
