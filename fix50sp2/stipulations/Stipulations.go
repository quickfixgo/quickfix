package stipulations

//NoStipulations is a repeating group in Stipulations
type NoStipulations struct {
	//StipulationType is a non-required field for NoStipulations.
	StipulationType *string `fix:"233"`
	//StipulationValue is a non-required field for NoStipulations.
	StipulationValue *string `fix:"234"`
}

//Component is a fix50sp2 Stipulations Component
type Component struct {
	//NoStipulations is a non-required field for Stipulations.
	NoStipulations []NoStipulations `fix:"232,omitempty"`
}

func New() *Component { return new(Component) }
