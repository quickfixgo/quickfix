package complexeventtimes

//New returns an initialized ComplexEventTimes instance
func New() *ComplexEventTimes {
	var m ComplexEventTimes
	return &m
}

//NoComplexEventTimes is a repeating group in ComplexEventTimes
type NoComplexEventTimes struct {
	//ComplexEventStartTime is a non-required field for NoComplexEventTimes.
	ComplexEventStartTime *string `fix:"1495"`
	//ComplexEventEndTime is a non-required field for NoComplexEventTimes.
	ComplexEventEndTime *string `fix:"1496"`
}

//NewNoComplexEventTimes returns an initialized NoComplexEventTimes instance
func NewNoComplexEventTimes() *NoComplexEventTimes {
	var m NoComplexEventTimes
	return &m
}

func (m *NoComplexEventTimes) SetComplexEventStartTime(v string) { m.ComplexEventStartTime = &v }
func (m *NoComplexEventTimes) SetComplexEventEndTime(v string)   { m.ComplexEventEndTime = &v }

//ComplexEventTimes is a fix50sp2 Component
type ComplexEventTimes struct {
	//NoComplexEventTimes is a non-required field for ComplexEventTimes.
	NoComplexEventTimes []NoComplexEventTimes `fix:"1494,omitempty"`
}

func (m *ComplexEventTimes) SetNoComplexEventTimes(v []NoComplexEventTimes) { m.NoComplexEventTimes = v }
