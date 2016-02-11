package attrbgrp

//NoInstrAttrib is a repeating group in AttrbGrp
type NoInstrAttrib struct {
	//InstrAttribType is a non-required field for NoInstrAttrib.
	InstrAttribType *int `fix:"871"`
	//InstrAttribValue is a non-required field for NoInstrAttrib.
	InstrAttribValue *string `fix:"872"`
}

//Component is a fix50sp2 AttrbGrp Component
type Component struct {
	//NoInstrAttrib is a non-required field for AttrbGrp.
	NoInstrAttrib []NoInstrAttrib `fix:"870,omitempty"`
}

func New() *Component { return new(Component) }
