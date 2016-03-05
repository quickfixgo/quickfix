package instrmtstrkpxgrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrument"
)

//NoStrikes is a repeating group in InstrmtStrkPxGrp
type NoStrikes struct {
	//Instrument Component
	instrument.Instrument
}

//InstrmtStrkPxGrp is a fix50 Component
type InstrmtStrkPxGrp struct {
	//NoStrikes is a required field for InstrmtStrkPxGrp.
	NoStrikes []NoStrikes `fix:"428"`
}

func (m *InstrmtStrkPxGrp) SetNoStrikes(v []NoStrikes) { m.NoStrikes = v }
