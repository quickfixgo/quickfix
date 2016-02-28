package ioiqualgrp

//NoIOIQualifiers is a repeating group in IOIQualGrp
type NoIOIQualifiers struct {
	//IOIQualifier is a non-required field for NoIOIQualifiers.
	IOIQualifier *string `fix:"104"`
}

//Component is a fix50 IOIQualGrp Component
type Component struct {
	//NoIOIQualifiers is a non-required field for IOIQualGrp.
	NoIOIQualifiers []NoIOIQualifiers `fix:"199,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoIOIQualifiers(v []NoIOIQualifiers) { m.NoIOIQualifiers = v }
