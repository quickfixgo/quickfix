package legstipulations

//NoLegStipulations is a repeating group in LegStipulations
type NoLegStipulations struct {
	//LegStipulationType is a non-required field for NoLegStipulations.
	LegStipulationType *string `fix:"688"`
	//LegStipulationValue is a non-required field for NoLegStipulations.
	LegStipulationValue *string `fix:"689"`
}

//Component is a fix44 LegStipulations Component
type Component struct {
	//NoLegStipulations is a non-required field for LegStipulations.
	NoLegStipulations []NoLegStipulations `fix:"683,omitempty"`
}

func New() *Component { return new(Component) }
