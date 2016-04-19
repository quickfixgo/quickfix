package legstipulations

//New returns an initialized LegStipulations instance
func New() *LegStipulations {
	var m LegStipulations
	return &m
}

//NoLegStipulations is a repeating group in LegStipulations
type NoLegStipulations struct {
	//LegStipulationType is a non-required field for NoLegStipulations.
	LegStipulationType *string `fix:"688"`
	//LegStipulationValue is a non-required field for NoLegStipulations.
	LegStipulationValue *string `fix:"689"`
}

//NewNoLegStipulations returns an initialized NoLegStipulations instance
func NewNoLegStipulations() *NoLegStipulations {
	var m NoLegStipulations
	return &m
}

func (m *NoLegStipulations) SetLegStipulationType(v string)  { m.LegStipulationType = &v }
func (m *NoLegStipulations) SetLegStipulationValue(v string) { m.LegStipulationValue = &v }

//LegStipulations is a fix44 Component
type LegStipulations struct {
	//NoLegStipulations is a non-required field for LegStipulations.
	NoLegStipulations []NoLegStipulations `fix:"683,omitempty"`
}

func (m *LegStipulations) SetNoLegStipulations(v []NoLegStipulations) { m.NoLegStipulations = v }
