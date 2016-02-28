package clrinstgrp

//NoClearingInstructions is a repeating group in ClrInstGrp
type NoClearingInstructions struct {
	//ClearingInstruction is a non-required field for NoClearingInstructions.
	ClearingInstruction *int `fix:"577"`
}

//Component is a fix50 ClrInstGrp Component
type Component struct {
	//NoClearingInstructions is a non-required field for ClrInstGrp.
	NoClearingInstructions []NoClearingInstructions `fix:"576,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoClearingInstructions(v []NoClearingInstructions) {
	m.NoClearingInstructions = v
}
