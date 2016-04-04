package maturityrules

func New() *MaturityRules {
	var m MaturityRules
	return &m
}

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

func (m *NoMaturityRules) SetMaturityRuleID(v string)       { m.MaturityRuleID = &v }
func (m *NoMaturityRules) SetMaturityMonthYearFormat(v int) { m.MaturityMonthYearFormat = &v }
func (m *NoMaturityRules) SetMaturityMonthYearIncrementUnits(v int) {
	m.MaturityMonthYearIncrementUnits = &v
}
func (m *NoMaturityRules) SetStartMaturityMonthYear(v string)  { m.StartMaturityMonthYear = &v }
func (m *NoMaturityRules) SetEndMaturityMonthYear(v string)    { m.EndMaturityMonthYear = &v }
func (m *NoMaturityRules) SetMaturityMonthYearIncrement(v int) { m.MaturityMonthYearIncrement = &v }

//MaturityRules is a fix50sp2 Component
type MaturityRules struct {
	//NoMaturityRules is a non-required field for MaturityRules.
	NoMaturityRules []NoMaturityRules `fix:"1236,omitempty"`
}

func (m *MaturityRules) SetNoMaturityRules(v []NoMaturityRules) { m.NoMaturityRules = v }
