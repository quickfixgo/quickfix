package tickrules

//NoTickRules is a repeating group in TickRules
type NoTickRules struct {
	//StartTickPriceRange is a non-required field for NoTickRules.
	StartTickPriceRange *float64 `fix:"1206"`
	//EndTickPriceRange is a non-required field for NoTickRules.
	EndTickPriceRange *float64 `fix:"1207"`
	//TickIncrement is a non-required field for NoTickRules.
	TickIncrement *float64 `fix:"1208"`
	//TickRuleType is a non-required field for NoTickRules.
	TickRuleType *int `fix:"1209"`
}

//Component is a fix50sp1 TickRules Component
type Component struct {
	//NoTickRules is a non-required field for TickRules.
	NoTickRules []NoTickRules `fix:"1205,omitempty"`
}

func New() *Component { return new(Component) }
