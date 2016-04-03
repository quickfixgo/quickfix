package undsecaltidgrp

func New() *UndSecAltIDGrp {
	var m UndSecAltIDGrp
	return &m
}

//NoUnderlyingSecurityAltID is a repeating group in UndSecAltIDGrp
type NoUnderlyingSecurityAltID struct {
	//UnderlyingSecurityAltID is a non-required field for NoUnderlyingSecurityAltID.
	UnderlyingSecurityAltID *string `fix:"458"`
	//UnderlyingSecurityAltIDSource is a non-required field for NoUnderlyingSecurityAltID.
	UnderlyingSecurityAltIDSource *string `fix:"459"`
}

//UndSecAltIDGrp is a fix50 Component
type UndSecAltIDGrp struct {
	//NoUnderlyingSecurityAltID is a non-required field for UndSecAltIDGrp.
	NoUnderlyingSecurityAltID []NoUnderlyingSecurityAltID `fix:"457,omitempty"`
}

func (m *UndSecAltIDGrp) SetNoUnderlyingSecurityAltID(v []NoUnderlyingSecurityAltID) {
	m.NoUnderlyingSecurityAltID = v
}
