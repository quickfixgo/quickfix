package execcollgrp

//NoExecs is a repeating group in ExecCollGrp
type NoExecs struct {
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
}

//Component is a fix50sp2 ExecCollGrp Component
type Component struct {
	//NoExecs is a non-required field for ExecCollGrp.
	NoExecs []NoExecs `fix:"124,omitempty"`
}

func New() *Component { return new(Component) }
