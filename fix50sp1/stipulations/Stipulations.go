package stipulations

//New returns an initialized Stipulations instance
func New() *Stipulations {
	var m Stipulations
	return &m
}

//NoStipulations is a repeating group in Stipulations
type NoStipulations struct {
	//StipulationType is a non-required field for NoStipulations.
	StipulationType *string `fix:"233"`
	//StipulationValue is a non-required field for NoStipulations.
	StipulationValue *string `fix:"234"`
}

//NewNoStipulations returns an initialized NoStipulations instance
func NewNoStipulations() *NoStipulations {
	var m NoStipulations
	return &m
}

func (m *NoStipulations) SetStipulationType(v string)  { m.StipulationType = &v }
func (m *NoStipulations) SetStipulationValue(v string) { m.StipulationValue = &v }

//Stipulations is a fix50sp1 Component
type Stipulations struct {
	//NoStipulations is a non-required field for Stipulations.
	NoStipulations []NoStipulations `fix:"232,omitempty"`
}

func (m *Stipulations) SetNoStipulations(v []NoStipulations) { m.NoStipulations = v }
