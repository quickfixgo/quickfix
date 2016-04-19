package legsecaltidgrp

//New returns an initialized LegSecAltIDGrp instance
func New() *LegSecAltIDGrp {
	var m LegSecAltIDGrp
	return &m
}

//NoLegSecurityAltID is a repeating group in LegSecAltIDGrp
type NoLegSecurityAltID struct {
	//LegSecurityAltID is a non-required field for NoLegSecurityAltID.
	LegSecurityAltID *string `fix:"605"`
	//LegSecurityAltIDSource is a non-required field for NoLegSecurityAltID.
	LegSecurityAltIDSource *string `fix:"606"`
}

//NewNoLegSecurityAltID returns an initialized NoLegSecurityAltID instance
func NewNoLegSecurityAltID() *NoLegSecurityAltID {
	var m NoLegSecurityAltID
	return &m
}

func (m *NoLegSecurityAltID) SetLegSecurityAltID(v string)       { m.LegSecurityAltID = &v }
func (m *NoLegSecurityAltID) SetLegSecurityAltIDSource(v string) { m.LegSecurityAltIDSource = &v }

//LegSecAltIDGrp is a fix50sp2 Component
type LegSecAltIDGrp struct {
	//NoLegSecurityAltID is a non-required field for LegSecAltIDGrp.
	NoLegSecurityAltID []NoLegSecurityAltID `fix:"604,omitempty"`
}

func (m *LegSecAltIDGrp) SetNoLegSecurityAltID(v []NoLegSecurityAltID) { m.NoLegSecurityAltID = v }
