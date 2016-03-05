package risklimits

import (
	"github.com/quickfixgo/quickfix/fix50sp2/riskinstrumentscope"
	"github.com/quickfixgo/quickfix/fix50sp2/riskwarninglevels"
)

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
	//RiskInstrumentScope Component
	riskinstrumentscope.RiskInstrumentScope
	//RiskWarningLevels Component
	riskwarninglevels.RiskWarningLevels
}

//RiskLimits is a fix50sp2 Component
type RiskLimits struct {
	//NoRiskLimits is a non-required field for RiskLimits.
	NoRiskLimits []NoRiskLimits `fix:"1529,omitempty"`
}

func (m *RiskLimits) SetNoRiskLimits(v []NoRiskLimits) { m.NoRiskLimits = v }
