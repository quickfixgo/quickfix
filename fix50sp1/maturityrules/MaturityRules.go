package maturityrules

//NoMaturityRules is a repeating group in MaturityRules
type NoMaturityRules struct {
	//MaturityRuleID is a non-required field for NoMaturityRules.
	MaturityRuleID *string `fix:"1222"`
	//MaturityMonthYearFormat is a non-required field for NoMaturityRules.
	MaturityMonthYearFormat *int `fix:"1303"`
	//MaturityMonthYearIncrementUnits is a non-required field for NoMaturityRules.
	MaturityMonthYearIncrementUnits *int `fix:"1302"`
	//StartMaturityMonthYear is a non-required field for NoMaturityRules.
	StartMaturityMonthYear *string `fix:"1241"`
	//EndMaturityMonthYear is a non-required field for NoMaturityRules.
	EndMaturityMonthYear *string `fix:"1226"`
	//MaturityMonthYearIncrement is a non-required field for NoMaturityRules.
	MaturityMonthYearIncrement *int `fix:"1229"`
}

//Component is a fix50sp1 MaturityRules Component
type Component struct {
	//NoMaturityRules is a non-required field for MaturityRules.
	NoMaturityRules []NoMaturityRules `fix:"1236,omitempty"`
}

func New() *Component { return new(Component) }
