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
	RiskInstrumentScope riskinstrumentscope.Component
	//RiskWarningLevels Component
	RiskWarningLevels riskwarninglevels.Component
}

//Component is a fix50sp2 RiskLimits Component
type Component struct {
	//NoRiskLimits is a non-required field for RiskLimits.
	NoRiskLimits []NoRiskLimits `fix:"1529,omitempty"`
}

func New() *Component { return new(Component) }
