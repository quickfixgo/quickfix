package attrbgrp

//New returns an initialized AttrbGrp instance
func New() *AttrbGrp {
	var m AttrbGrp
	return &m
}

//NoInstrAttrib is a repeating group in AttrbGrp
type NoInstrAttrib struct {
	//InstrAttribType is a non-required field for NoInstrAttrib.
	InstrAttribType *int `fix:"871"`
	//InstrAttribValue is a non-required field for NoInstrAttrib.
	InstrAttribValue *string `fix:"872"`
}

//NewNoInstrAttrib returns an initialized NoInstrAttrib instance
func NewNoInstrAttrib() *NoInstrAttrib {
	var m NoInstrAttrib
	return &m
}

func (m *NoInstrAttrib) SetInstrAttribType(v int)     { m.InstrAttribType = &v }
func (m *NoInstrAttrib) SetInstrAttribValue(v string) { m.InstrAttribValue = &v }

//AttrbGrp is a fix50 Component
type AttrbGrp struct {
	//NoInstrAttrib is a non-required field for AttrbGrp.
	NoInstrAttrib []NoInstrAttrib `fix:"870,omitempty"`
}

func (m *AttrbGrp) SetNoInstrAttrib(v []NoInstrAttrib) { m.NoInstrAttrib = v }
