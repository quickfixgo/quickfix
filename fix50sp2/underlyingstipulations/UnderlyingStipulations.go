package underlyingstipulations

//NoUnderlyingStips is a repeating group in UnderlyingStipulations
type NoUnderlyingStips struct {
	//UnderlyingStipType is a non-required field for NoUnderlyingStips.
	UnderlyingStipType *string `fix:"888"`
	//UnderlyingStipValue is a non-required field for NoUnderlyingStips.
	UnderlyingStipValue *string `fix:"889"`
}

//Component is a fix50sp2 UnderlyingStipulations Component
type Component struct {
	//NoUnderlyingStips is a non-required field for UnderlyingStipulations.
	NoUnderlyingStips []NoUnderlyingStips `fix:"887,omitempty"`
}

func New() *Component { return new(Component) }
