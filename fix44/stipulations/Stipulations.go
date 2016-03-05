package stipulations

//NoStipulations is a repeating group in Stipulations
type NoStipulations struct {
	//StipulationType is a non-required field for NoStipulations.
	StipulationType *string `fix:"233"`
	//StipulationValue is a non-required field for NoStipulations.
	StipulationValue *string `fix:"234"`
}

//Stipulations is a fix44 Component
type Stipulations struct {
	//NoStipulations is a non-required field for Stipulations.
	NoStipulations []NoStipulations `fix:"232,omitempty"`
}

func (m *Stipulations) SetNoStipulations(v []NoStipulations) { m.NoStipulations = v }
