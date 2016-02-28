package risksecaltidgrp

//NoRiskSecurityAltID is a repeating group in RiskSecAltIDGrp
type NoRiskSecurityAltID struct {
	//RiskSecurityAltID is a non-required field for NoRiskSecurityAltID.
	RiskSecurityAltID *string `fix:"1541"`
	//RiskSecurityAltIDSource is a non-required field for NoRiskSecurityAltID.
	RiskSecurityAltIDSource *string `fix:"1542"`
}

//Component is a fix50sp2 RiskSecAltIDGrp Component
type Component struct {
	//NoRiskSecurityAltID is a non-required field for RiskSecAltIDGrp.
	NoRiskSecurityAltID []NoRiskSecurityAltID `fix:"1540,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRiskSecurityAltID(v []NoRiskSecurityAltID) { m.NoRiskSecurityAltID = v }
