package undsecaltidgrp

//NoUnderlyingSecurityAltID is a repeating group in UndSecAltIDGrp
type NoUnderlyingSecurityAltID struct {
	//UnderlyingSecurityAltID is a non-required field for NoUnderlyingSecurityAltID.
	UnderlyingSecurityAltID *string `fix:"458"`
	//UnderlyingSecurityAltIDSource is a non-required field for NoUnderlyingSecurityAltID.
	UnderlyingSecurityAltIDSource *string `fix:"459"`
}

//Component is a fix50 UndSecAltIDGrp Component
type Component struct {
	//NoUnderlyingSecurityAltID is a non-required field for UndSecAltIDGrp.
	NoUnderlyingSecurityAltID []NoUnderlyingSecurityAltID `fix:"457,omitempty"`
}

func New() *Component { return new(Component) }
