package tickrules

func New() *TickRules {
	var m TickRules
	return &m
}

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

//TickRules is a fix50sp2 Component
type TickRules struct {
	//NoTickRules is a non-required field for TickRules.
	NoTickRules []NoTickRules `fix:"1205,omitempty"`
}

func (m *TickRules) SetNoTickRules(v []NoTickRules) { m.NoTickRules = v }
