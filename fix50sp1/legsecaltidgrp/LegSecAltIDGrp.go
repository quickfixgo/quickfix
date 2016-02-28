package legsecaltidgrp

//NoLegSecurityAltID is a repeating group in LegSecAltIDGrp
type NoLegSecurityAltID struct {
	//LegSecurityAltID is a non-required field for NoLegSecurityAltID.
	LegSecurityAltID *string `fix:"605"`
	//LegSecurityAltIDSource is a non-required field for NoLegSecurityAltID.
	LegSecurityAltIDSource *string `fix:"606"`
}

//Component is a fix50sp1 LegSecAltIDGrp Component
type Component struct {
	//NoLegSecurityAltID is a non-required field for LegSecAltIDGrp.
	NoLegSecurityAltID []NoLegSecurityAltID `fix:"604,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoLegSecurityAltID(v []NoLegSecurityAltID) { m.NoLegSecurityAltID = v }
