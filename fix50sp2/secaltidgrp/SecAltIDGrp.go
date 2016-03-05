package secaltidgrp

//NoSecurityAltID is a repeating group in SecAltIDGrp
type NoSecurityAltID struct {
	//SecurityAltID is a non-required field for NoSecurityAltID.
	SecurityAltID *string `fix:"455"`
	//SecurityAltIDSource is a non-required field for NoSecurityAltID.
	SecurityAltIDSource *string `fix:"456"`
}

//SecAltIDGrp is a fix50sp2 Component
type SecAltIDGrp struct {
	//NoSecurityAltID is a non-required field for SecAltIDGrp.
	NoSecurityAltID []NoSecurityAltID `fix:"454,omitempty"`
}

func (m *SecAltIDGrp) SetNoSecurityAltID(v []NoSecurityAltID) { m.NoSecurityAltID = v }
