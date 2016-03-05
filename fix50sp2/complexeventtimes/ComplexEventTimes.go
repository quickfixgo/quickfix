package complexeventtimes

//NoComplexEventTimes is a repeating group in ComplexEventTimes
type NoComplexEventTimes struct {
	//ComplexEventStartTime is a non-required field for NoComplexEventTimes.
	ComplexEventStartTime *string `fix:"1495"`
	//ComplexEventEndTime is a non-required field for NoComplexEventTimes.
	ComplexEventEndTime *string `fix:"1496"`
}

//ComplexEventTimes is a fix50sp2 Component
type ComplexEventTimes struct {
	//NoComplexEventTimes is a non-required field for ComplexEventTimes.
	NoComplexEventTimes []NoComplexEventTimes `fix:"1494,omitempty"`
}

func (m *ComplexEventTimes) SetNoComplexEventTimes(v []NoComplexEventTimes) { m.NoComplexEventTimes = v }
