package mdrjctgrp

func New() *MDRjctGrp {
	var m MDRjctGrp
	return &m
}

//NoAltMDSource is a repeating group in MDRjctGrp
type NoAltMDSource struct {
	//AltMDSourceID is a non-required field for NoAltMDSource.
	AltMDSourceID *string `fix:"817"`
}

func (m *NoAltMDSource) SetAltMDSourceID(v string) { m.AltMDSourceID = &v }

//MDRjctGrp is a fix50sp1 Component
type MDRjctGrp struct {
	//NoAltMDSource is a non-required field for MDRjctGrp.
	NoAltMDSource []NoAltMDSource `fix:"816,omitempty"`
}

func (m *MDRjctGrp) SetNoAltMDSource(v []NoAltMDSource) { m.NoAltMDSource = v }
