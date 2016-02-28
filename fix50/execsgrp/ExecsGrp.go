package execsgrp

//NoExecs is a repeating group in ExecsGrp
type NoExecs struct {
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
}

//Component is a fix50 ExecsGrp Component
type Component struct {
	//NoExecs is a non-required field for ExecsGrp.
	NoExecs []NoExecs `fix:"124,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoExecs(v []NoExecs) { m.NoExecs = v }
