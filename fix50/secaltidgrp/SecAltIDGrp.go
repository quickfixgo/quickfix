package secaltidgrp

//NoSecurityAltID is a repeating group in SecAltIDGrp
type NoSecurityAltID struct {
	//SecurityAltID is a non-required field for NoSecurityAltID.
	SecurityAltID *string `fix:"455"`
	//SecurityAltIDSource is a non-required field for NoSecurityAltID.
	SecurityAltIDSource *string `fix:"456"`
}

//Component is a fix50 SecAltIDGrp Component
type Component struct {
	//NoSecurityAltID is a non-required field for SecAltIDGrp.
	NoSecurityAltID []NoSecurityAltID `fix:"454,omitempty"`
}

func New() *Component { return new(Component) }
