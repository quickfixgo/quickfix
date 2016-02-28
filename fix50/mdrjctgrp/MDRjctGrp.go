package mdrjctgrp

//NoAltMDSource is a repeating group in MDRjctGrp
type NoAltMDSource struct {
	//AltMDSourceID is a non-required field for NoAltMDSource.
	AltMDSourceID *string `fix:"817"`
}

//Component is a fix50 MDRjctGrp Component
type Component struct {
	//NoAltMDSource is a non-required field for MDRjctGrp.
	NoAltMDSource []NoAltMDSource `fix:"816,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoAltMDSource(v []NoAltMDSource) { m.NoAltMDSource = v }
