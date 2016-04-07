package ioiqualgrp

func New() *IOIQualGrp {
	var m IOIQualGrp
	return &m
}

//NoIOIQualifiers is a repeating group in IOIQualGrp
type NoIOIQualifiers struct {
	//IOIQualifier is a non-required field for NoIOIQualifiers.
	IOIQualifier *string `fix:"104"`
}

func (m *NoIOIQualifiers) SetIOIQualifier(v string) { m.IOIQualifier = &v }

//IOIQualGrp is a fix50sp2 Component
type IOIQualGrp struct {
	//NoIOIQualifiers is a non-required field for IOIQualGrp.
	NoIOIQualifiers []NoIOIQualifiers `fix:"199,omitempty"`
}

func (m *IOIQualGrp) SetNoIOIQualifiers(v []NoIOIQualifiers) { m.NoIOIQualifiers = v }
