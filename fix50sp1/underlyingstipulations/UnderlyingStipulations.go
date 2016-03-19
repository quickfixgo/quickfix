package underlyingstipulations

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

//UnderlyingStipulations is a fix50sp1 Component
type UnderlyingStipulations struct {
	//NoUnderlyingStips is a non-required field for UnderlyingStipulations.
	NoUnderlyingStips []NoUnderlyingStips `fix:"887,omitempty"`
}

func (m *UnderlyingStipulations) SetNoUnderlyingStips(v []NoUnderlyingStips) { m.NoUnderlyingStips = v }
