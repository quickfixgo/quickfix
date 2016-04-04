package matchrules

func New() *MatchRules {
	var m MatchRules
	return &m
}

//NoMatchRules is a repeating group in MatchRules
type NoMatchRules struct {
	//MatchAlgorithm is a non-required field for NoMatchRules.
	MatchAlgorithm *string `fix:"1142"`
	//MatchType is a non-required field for NoMatchRules.
	MatchType *string `fix:"574"`
}

func (m *NoMatchRules) SetMatchAlgorithm(v string) { m.MatchAlgorithm = &v }
func (m *NoMatchRules) SetMatchType(v string)      { m.MatchType = &v }

//MatchRules is a fix50sp1 Component
type MatchRules struct {
	//NoMatchRules is a non-required field for MatchRules.
	NoMatchRules []NoMatchRules `fix:"1235,omitempty"`
}

func (m *MatchRules) SetNoMatchRules(v []NoMatchRules) { m.NoMatchRules = v }
