package timeinforcerules

//New returns an initialized TimeInForceRules instance
func New() *TimeInForceRules {
	var m TimeInForceRules
	return &m
}

//NoTimeInForceRules is a repeating group in TimeInForceRules
type NoTimeInForceRules struct {
	//TimeInForce is a non-required field for NoTimeInForceRules.
	TimeInForce *string `fix:"59"`
}

//NewNoTimeInForceRules returns an initialized NoTimeInForceRules instance
func NewNoTimeInForceRules() *NoTimeInForceRules {
	var m NoTimeInForceRules
	return &m
}

func (m *NoTimeInForceRules) SetTimeInForce(v string) { m.TimeInForce = &v }

//TimeInForceRules is a fix50sp1 Component
type TimeInForceRules struct {
	//NoTimeInForceRules is a non-required field for TimeInForceRules.
	NoTimeInForceRules []NoTimeInForceRules `fix:"1239,omitempty"`
}

func (m *TimeInForceRules) SetNoTimeInForceRules(v []NoTimeInForceRules) { m.NoTimeInForceRules = v }
