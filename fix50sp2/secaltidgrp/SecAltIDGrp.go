package secaltidgrp

//New returns an initialized SecAltIDGrp instance
func New() *SecAltIDGrp {
	var m SecAltIDGrp
	return &m
}

//NoSecurityAltID is a repeating group in SecAltIDGrp
type NoSecurityAltID struct {
	//SecurityAltID is a non-required field for NoSecurityAltID.
	SecurityAltID *string `fix:"455"`
	//SecurityAltIDSource is a non-required field for NoSecurityAltID.
	SecurityAltIDSource *string `fix:"456"`
}

//NewNoSecurityAltID returns an initialized NoSecurityAltID instance
func NewNoSecurityAltID() *NoSecurityAltID {
	var m NoSecurityAltID
	return &m
}

func (m *NoSecurityAltID) SetSecurityAltID(v string)       { m.SecurityAltID = &v }
func (m *NoSecurityAltID) SetSecurityAltIDSource(v string) { m.SecurityAltIDSource = &v }

//SecAltIDGrp is a fix50sp2 Component
type SecAltIDGrp struct {
	//NoSecurityAltID is a non-required field for SecAltIDGrp.
	NoSecurityAltID []NoSecurityAltID `fix:"454,omitempty"`
}

func (m *SecAltIDGrp) SetNoSecurityAltID(v []NoSecurityAltID) { m.NoSecurityAltID = v }
