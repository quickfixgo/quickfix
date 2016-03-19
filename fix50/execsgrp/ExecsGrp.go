package execsgrp

func New() *ExecsGrp {
	var m ExecsGrp
	return &m
}

//NoExecs is a repeating group in ExecsGrp
type NoExecs struct {
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
}

//ExecsGrp is a fix50 Component
type ExecsGrp struct {
	//NoExecs is a non-required field for ExecsGrp.
	NoExecs []NoExecs `fix:"124,omitempty"`
}

func (m *ExecsGrp) SetNoExecs(v []NoExecs) { m.NoExecs = v }
