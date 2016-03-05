package underlyinglegsecurityaltidgrp

//NoUnderlyingLegSecurityAltID is a repeating group in UnderlyingLegSecurityAltIDGrp
type NoUnderlyingLegSecurityAltID struct {
	//UnderlyingLegSecurityAltID is a non-required field for NoUnderlyingLegSecurityAltID.
	UnderlyingLegSecurityAltID *string `fix:"1335"`
	//UnderlyingLegSecurityAltIDSource is a non-required field for NoUnderlyingLegSecurityAltID.
	UnderlyingLegSecurityAltIDSource *string `fix:"1336"`
}

//UnderlyingLegSecurityAltIDGrp is a fix50sp1 Component
type UnderlyingLegSecurityAltIDGrp struct {
	//NoUnderlyingLegSecurityAltID is a non-required field for UnderlyingLegSecurityAltIDGrp.
	NoUnderlyingLegSecurityAltID []NoUnderlyingLegSecurityAltID `fix:"1334,omitempty"`
}

func (m *UnderlyingLegSecurityAltIDGrp) SetNoUnderlyingLegSecurityAltID(v []NoUnderlyingLegSecurityAltID) {
	m.NoUnderlyingLegSecurityAltID = v
}
