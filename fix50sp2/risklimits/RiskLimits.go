package risklimits

import (
	"github.com/quickfixgo/quickfix/fix50sp2/riskinstrumentscope"
	"github.com/quickfixgo/quickfix/fix50sp2/riskwarninglevels"
)

//New returns an initialized RiskLimits instance
func New() *RiskLimits {
	var m RiskLimits
	return &m
}

//NoRiskLimits is a repeating group in RiskLimits
type NoRiskLimits struct {
	//RiskLimitType is a non-required field for NoRiskLimits.
	RiskLimitType *int `fix:"1530"`
	//RiskLimitAmount is a non-required field for NoRiskLimits.
	RiskLimitAmount *float64 `fix:"1531"`
	//RiskLimitCurrency is a non-required field for NoRiskLimits.
	RiskLimitCurrency *string `fix:"1532"`
	//RiskLimitPlatform is a non-required field for NoRiskLimits.
	RiskLimitPlatform *string `fix:"1533"`
	//RiskInstrumentScope is a non-required component for NoRiskLimits.
	RiskInstrumentScope *riskinstrumentscope.RiskInstrumentScope
	//RiskWarningLevels is a non-required component for NoRiskLimits.
	RiskWarningLevels *riskwarninglevels.RiskWarningLevels
}

//NewNoRiskLimits returns an initialized NoRiskLimits instance
func NewNoRiskLimits() *NoRiskLimits {
	var m NoRiskLimits
	return &m
}

func (m *NoRiskLimits) SetRiskLimitType(v int)        { m.RiskLimitType = &v }
func (m *NoRiskLimits) SetRiskLimitAmount(v float64)  { m.RiskLimitAmount = &v }
func (m *NoRiskLimits) SetRiskLimitCurrency(v string) { m.RiskLimitCurrency = &v }
func (m *NoRiskLimits) SetRiskLimitPlatform(v string) { m.RiskLimitPlatform = &v }
func (m *NoRiskLimits) SetRiskInstrumentScope(v riskinstrumentscope.RiskInstrumentScope) {
	m.RiskInstrumentScope = &v
}
func (m *NoRiskLimits) SetRiskWarningLevels(v riskwarninglevels.RiskWarningLevels) {
	m.RiskWarningLevels = &v
}

//RiskLimits is a fix50sp2 Component
type RiskLimits struct {
	//NoRiskLimits is a non-required field for RiskLimits.
	NoRiskLimits []NoRiskLimits `fix:"1529,omitempty"`
}

func (m *RiskLimits) SetNoRiskLimits(v []NoRiskLimits) { m.NoRiskLimits = v }
