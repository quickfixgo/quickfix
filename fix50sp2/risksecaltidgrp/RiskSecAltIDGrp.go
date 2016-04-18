package risksecaltidgrp

//New returns an initialized RiskSecAltIDGrp instance
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

//NewNoRiskSecurityAltID returns an initialized NoRiskSecurityAltID instance
func NewNoRiskSecurityAltID() *NoRiskSecurityAltID {
	var m NoRiskSecurityAltID
	return &m
}

func (m *NoRiskSecurityAltID) SetRiskSecurityAltID(v string)       { m.RiskSecurityAltID = &v }
func (m *NoRiskSecurityAltID) SetRiskSecurityAltIDSource(v string) { m.RiskSecurityAltIDSource = &v }

//RiskSecAltIDGrp is a fix50sp2 Component
type RiskSecAltIDGrp struct {
	//NoRiskSecurityAltID is a non-required field for RiskSecAltIDGrp.
	NoRiskSecurityAltID []NoRiskSecurityAltID `fix:"1540,omitempty"`
}

func (m *RiskSecAltIDGrp) SetNoRiskSecurityAltID(v []NoRiskSecurityAltID) { m.NoRiskSecurityAltID = v }
