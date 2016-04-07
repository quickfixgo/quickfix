package instrmtstrkpxgrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrument"
)

func New(nostrikes []NoStrikes) *InstrmtStrkPxGrp {
	var m InstrmtStrkPxGrp
	m.SetNoStrikes(nostrikes)
	return &m
}

//NoStrikes is a repeating group in InstrmtStrkPxGrp
type NoStrikes struct {
	//Instrument is a required component for NoStrikes.
	instrument.Instrument
}

func (m *NoStrikes) SetInstrument(v instrument.Instrument) { m.Instrument = v }

//InstrmtStrkPxGrp is a fix50 Component
type InstrmtStrkPxGrp struct {
	//NoStrikes is a required field for InstrmtStrkPxGrp.
	NoStrikes []NoStrikes `fix:"428"`
}

func (m *InstrmtStrkPxGrp) SetNoStrikes(v []NoStrikes) { m.NoStrikes = v }
