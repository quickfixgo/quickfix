package instrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
)

func New() *InstrmtGrp {
	var m InstrmtGrp
	return &m
}

//NoRelatedSym is a repeating group in InstrmtGrp
type NoRelatedSym struct {
	//Instrument is a non-required component for NoRelatedSym.
	Instrument *instrument.Instrument
}

func (m *NoRelatedSym) SetInstrument(v instrument.Instrument) { m.Instrument = &v }

//InstrmtGrp is a fix50sp1 Component
type InstrmtGrp struct {
	//NoRelatedSym is a non-required field for InstrmtGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *InstrmtGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
