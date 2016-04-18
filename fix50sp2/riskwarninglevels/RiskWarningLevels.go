package riskwarninglevels

//New returns an initialized RiskWarningLevels instance
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

//NewNoRiskWarningLevels returns an initialized NoRiskWarningLevels instance
func NewNoRiskWarningLevels() *NoRiskWarningLevels {
	var m NoRiskWarningLevels
	return &m
}

func (m *NoRiskWarningLevels) SetRiskWarningLevelPercent(v float64) { m.RiskWarningLevelPercent = &v }
func (m *NoRiskWarningLevels) SetRiskWarningLevelName(v string)     { m.RiskWarningLevelName = &v }

//RiskWarningLevels is a fix50sp2 Component
type RiskWarningLevels struct {
	//NoRiskWarningLevels is a non-required field for RiskWarningLevels.
	NoRiskWarningLevels []NoRiskWarningLevels `fix:"1559,omitempty"`
}

func (m *RiskWarningLevels) SetNoRiskWarningLevels(v []NoRiskWarningLevels) { m.NoRiskWarningLevels = v }
