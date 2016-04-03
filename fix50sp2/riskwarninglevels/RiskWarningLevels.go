package riskwarninglevels

func New() *RiskWarningLevels {
	var m RiskWarningLevels
	return &m
}

//NoRiskWarningLevels is a repeating group in RiskWarningLevels
type NoRiskWarningLevels struct {
	//RiskWarningLevelPercent is a non-required field for NoRiskWarningLevels.
	RiskWarningLevelPercent *float64 `fix:"1560"`
	//RiskWarningLevelName is a non-required field for NoRiskWarningLevels.
	RiskWarningLevelName *string `fix:"1561"`
}

//RiskWarningLevels is a fix50sp2 Component
type RiskWarningLevels struct {
	//NoRiskWarningLevels is a non-required field for RiskWarningLevels.
	NoRiskWarningLevels []NoRiskWarningLevels `fix:"1559,omitempty"`
}

func (m *RiskWarningLevels) SetNoRiskWarningLevels(v []NoRiskWarningLevels) { m.NoRiskWarningLevels = v }
