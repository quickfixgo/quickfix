package timeinforcerules

//NoTimeInForceRules is a repeating group in TimeInForceRules
type NoTimeInForceRules struct {
	//TimeInForce is a non-required field for NoTimeInForceRules.
	TimeInForce *string `fix:"59"`
}

//Component is a fix50sp1 TimeInForceRules Component
type Component struct {
	//NoTimeInForceRules is a non-required field for TimeInForceRules.
	NoTimeInForceRules []NoTimeInForceRules `fix:"1239,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoTimeInForceRules(v []NoTimeInForceRules) { m.NoTimeInForceRules = v }
