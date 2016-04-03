package legstipulations

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

//LegStipulations is a fix44 Component
type LegStipulations struct {
	//NoLegStipulations is a non-required field for LegStipulations.
	NoLegStipulations []NoLegStipulations `fix:"683,omitempty"`
}

func (m *LegStipulations) SetNoLegStipulations(v []NoLegStipulations) { m.NoLegStipulations = v }
