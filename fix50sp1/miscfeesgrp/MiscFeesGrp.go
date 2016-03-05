package miscfeesgrp

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

//MiscFeesGrp is a fix50sp1 Component
type MiscFeesGrp struct {
	//NoMiscFees is a non-required field for MiscFeesGrp.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
}

func (m *MiscFeesGrp) SetNoMiscFees(v []NoMiscFees) { m.NoMiscFees = v }
