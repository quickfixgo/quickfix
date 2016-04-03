package risksecaltidgrp

func New() *RiskSecAltIDGrp {
	var m RiskSecAltIDGrp
	return &m
}

//NoRiskSecurityAltID is a repeating group in RiskSecAltIDGrp
type NoRiskSecurityAltID struct {
	//RiskSecurityAltID is a non-required field for NoRiskSecurityAltID.
	RiskSecurityAltID *string `fix:"1541"`
	//RiskSecurityAltIDSource is a non-required field for NoRiskSecurityAltID.
	RiskSecurityAltIDSource *string `fix:"1542"`
}

//RiskSecAltIDGrp is a fix50sp2 Component
type RiskSecAltIDGrp struct {
	//NoRiskSecurityAltID is a non-required field for RiskSecAltIDGrp.
	NoRiskSecurityAltID []NoRiskSecurityAltID `fix:"1540,omitempty"`
}

func (m *RiskSecAltIDGrp) SetNoRiskSecurityAltID(v []NoRiskSecurityAltID) { m.NoRiskSecurityAltID = v }
