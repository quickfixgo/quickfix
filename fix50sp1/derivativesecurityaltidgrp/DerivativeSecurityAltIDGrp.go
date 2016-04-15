package derivativesecurityaltidgrp

func New() *DerivativeSecurityAltIDGrp {
	var m DerivativeSecurityAltIDGrp
	return &m
}

//NoDerivativeSecurityAltID is a repeating group in DerivativeSecurityAltIDGrp
type NoDerivativeSecurityAltID struct {
	//DerivativeSecurityAltID is a non-required field for NoDerivativeSecurityAltID.
	DerivativeSecurityAltID *string `fix:"1219"`
	//DerivativeSecurityAltIDSource is a non-required field for NoDerivativeSecurityAltID.
	DerivativeSecurityAltIDSource *string `fix:"1220"`
}

func (m *NoDerivativeSecurityAltID) SetDerivativeSecurityAltID(v string) {
	m.DerivativeSecurityAltID = &v
}
func (m *NoDerivativeSecurityAltID) SetDerivativeSecurityAltIDSource(v string) {
	m.DerivativeSecurityAltIDSource = &v
}

//DerivativeSecurityAltIDGrp is a fix50sp1 Component
type DerivativeSecurityAltIDGrp struct {
	//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityAltIDGrp.
	NoDerivativeSecurityAltID []NoDerivativeSecurityAltID `fix:"1218,omitempty"`
}

func (m *DerivativeSecurityAltIDGrp) SetNoDerivativeSecurityAltID(v []NoDerivativeSecurityAltID) {
	m.NoDerivativeSecurityAltID = v
}
