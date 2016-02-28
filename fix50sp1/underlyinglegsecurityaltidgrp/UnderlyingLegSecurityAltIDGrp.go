package underlyinglegsecurityaltidgrp

//NoUnderlyingLegSecurityAltID is a repeating group in UnderlyingLegSecurityAltIDGrp
type NoUnderlyingLegSecurityAltID struct {
	//UnderlyingLegSecurityAltID is a non-required field for NoUnderlyingLegSecurityAltID.
	UnderlyingLegSecurityAltID *string `fix:"1335"`
	//UnderlyingLegSecurityAltIDSource is a non-required field for NoUnderlyingLegSecurityAltID.
	UnderlyingLegSecurityAltIDSource *string `fix:"1336"`
}

//Component is a fix50sp1 UnderlyingLegSecurityAltIDGrp Component
type Component struct {
	//NoUnderlyingLegSecurityAltID is a non-required field for UnderlyingLegSecurityAltIDGrp.
	NoUnderlyingLegSecurityAltID []NoUnderlyingLegSecurityAltID `fix:"1334,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoUnderlyingLegSecurityAltID(v []NoUnderlyingLegSecurityAltID) {
	m.NoUnderlyingLegSecurityAltID = v
}
