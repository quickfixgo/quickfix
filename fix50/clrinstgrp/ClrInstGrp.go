package clrinstgrp

//NoClearingInstructions is a repeating group in ClrInstGrp
type NoClearingInstructions struct {
	//ClearingInstruction is a non-required field for NoClearingInstructions.
	ClearingInstruction *int `fix:"577"`
}

//ClrInstGrp is a fix50 Component
type ClrInstGrp struct {
	//NoClearingInstructions is a non-required field for ClrInstGrp.
	NoClearingInstructions []NoClearingInstructions `fix:"576,omitempty"`
}

func (m *ClrInstGrp) SetNoClearingInstructions(v []NoClearingInstructions) {
	m.NoClearingInstructions = v
}
