package timeinforcerules

//NoTimeInForceRules is a repeating group in TimeInForceRules
type NoTimeInForceRules struct {
	//TimeInForce is a non-required field for NoTimeInForceRules.
	TimeInForce *string `fix:"59"`
}

//TimeInForceRules is a fix50sp1 Component
type TimeInForceRules struct {
	//NoTimeInForceRules is a non-required field for TimeInForceRules.
	NoTimeInForceRules []NoTimeInForceRules `fix:"1239,omitempty"`
}

func (m *TimeInForceRules) SetNoTimeInForceRules(v []NoTimeInForceRules) { m.NoTimeInForceRules = v }
