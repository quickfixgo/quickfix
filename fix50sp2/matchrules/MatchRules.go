package matchrules

//NoMatchRules is a repeating group in MatchRules
type NoMatchRules struct {
	//MatchAlgorithm is a non-required field for NoMatchRules.
	MatchAlgorithm *string `fix:"1142"`
	//MatchType is a non-required field for NoMatchRules.
	MatchType *string `fix:"574"`
}

//Component is a fix50sp2 MatchRules Component
type Component struct {
	//NoMatchRules is a non-required field for MatchRules.
	NoMatchRules []NoMatchRules `fix:"1235,omitempty"`
}

func New() *Component { return new(Component) }
