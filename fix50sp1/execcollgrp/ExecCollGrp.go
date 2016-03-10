package execcollgrp

//NoExecs is a repeating group in ExecCollGrp
type NoExecs struct {
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
}

//ExecCollGrp is a fix50sp1 Component
type ExecCollGrp struct {
	//NoExecs is a non-required field for ExecCollGrp.
	NoExecs []NoExecs `fix:"124,omitempty"`
}

func (m *ExecCollGrp) SetNoExecs(v []NoExecs) { m.NoExecs = v }
