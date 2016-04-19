package execcollgrp

//New returns an initialized ExecCollGrp instance
func New() *ExecCollGrp {
	var m ExecCollGrp
	return &m
}

//NoExecs is a repeating group in ExecCollGrp
type NoExecs struct {
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
}

//NewNoExecs returns an initialized NoExecs instance
func NewNoExecs() *NoExecs {
	var m NoExecs
	return &m
}

func (m *NoExecs) SetExecID(v string) { m.ExecID = &v }

//ExecCollGrp is a fix50sp1 Component
type ExecCollGrp struct {
	//NoExecs is a non-required field for ExecCollGrp.
	NoExecs []NoExecs `fix:"124,omitempty"`
}

func (m *ExecCollGrp) SetNoExecs(v []NoExecs) { m.NoExecs = v }
