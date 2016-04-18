package miscfeesgrp

//New returns an initialized MiscFeesGrp instance
func New() *MiscFeesGrp {
	var m MiscFeesGrp
	return &m
}

//NoMiscFees is a repeating group in MiscFeesGrp
type NoMiscFees struct {
	//MiscFeeAmt is a non-required field for NoMiscFees.
	MiscFeeAmt *float64 `fix:"137"`
	//MiscFeeCurr is a non-required field for NoMiscFees.
	MiscFeeCurr *string `fix:"138"`
	//MiscFeeType is a non-required field for NoMiscFees.
	MiscFeeType *string `fix:"139"`
	//MiscFeeBasis is a non-required field for NoMiscFees.
	MiscFeeBasis *int `fix:"891"`
}

//NewNoMiscFees returns an initialized NoMiscFees instance
func NewNoMiscFees() *NoMiscFees {
	var m NoMiscFees
	return &m
}

func (m *NoMiscFees) SetMiscFeeAmt(v float64) { m.MiscFeeAmt = &v }
func (m *NoMiscFees) SetMiscFeeCurr(v string) { m.MiscFeeCurr = &v }
func (m *NoMiscFees) SetMiscFeeType(v string) { m.MiscFeeType = &v }
func (m *NoMiscFees) SetMiscFeeBasis(v int)   { m.MiscFeeBasis = &v }

//MiscFeesGrp is a fix50 Component
type MiscFeesGrp struct {
	//NoMiscFees is a non-required field for MiscFeesGrp.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
}

func (m *MiscFeesGrp) SetNoMiscFees(v []NoMiscFees) { m.NoMiscFees = v }
