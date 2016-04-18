package tickrules

//New returns an initialized TickRules instance
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

//NewNoTickRules returns an initialized NoTickRules instance
func NewNoTickRules() *NoTickRules {
	var m NoTickRules
	return &m
}

func (m *NoTickRules) SetStartTickPriceRange(v float64) { m.StartTickPriceRange = &v }
func (m *NoTickRules) SetEndTickPriceRange(v float64)   { m.EndTickPriceRange = &v }
func (m *NoTickRules) SetTickIncrement(v float64)       { m.TickIncrement = &v }
func (m *NoTickRules) SetTickRuleType(v int)            { m.TickRuleType = &v }

//TickRules is a fix50sp1 Component
type TickRules struct {
	//NoTickRules is a non-required field for TickRules.
	NoTickRules []NoTickRules `fix:"1205,omitempty"`
}

func (m *TickRules) SetNoTickRules(v []NoTickRules) { m.NoTickRules = v }
