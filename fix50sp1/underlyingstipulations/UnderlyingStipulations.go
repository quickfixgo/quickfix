package underlyingstipulations

//New returns an initialized UnderlyingStipulations instance
func New() *UnderlyingStipulations {
	var m UnderlyingStipulations
	return &m
}

//NoUnderlyingStips is a repeating group in UnderlyingStipulations
type NoUnderlyingStips struct {
	//UnderlyingStipType is a non-required field for NoUnderlyingStips.
	UnderlyingStipType *string `fix:"888"`
	//UnderlyingStipValue is a non-required field for NoUnderlyingStips.
	UnderlyingStipValue *string `fix:"889"`
}

//NewNoUnderlyingStips returns an initialized NoUnderlyingStips instance
func NewNoUnderlyingStips() *NoUnderlyingStips {
	var m NoUnderlyingStips
	return &m
}

func (m *NoUnderlyingStips) SetUnderlyingStipType(v string)  { m.UnderlyingStipType = &v }
func (m *NoUnderlyingStips) SetUnderlyingStipValue(v string) { m.UnderlyingStipValue = &v }

//UnderlyingStipulations is a fix50sp1 Component
type UnderlyingStipulations struct {
	//NoUnderlyingStips is a non-required field for UnderlyingStipulations.
	NoUnderlyingStips []NoUnderlyingStips `fix:"887,omitempty"`
}

func (m *UnderlyingStipulations) SetNoUnderlyingStips(v []NoUnderlyingStips) { m.NoUnderlyingStips = v }
