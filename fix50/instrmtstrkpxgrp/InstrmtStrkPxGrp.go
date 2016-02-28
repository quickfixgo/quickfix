package instrmtstrkpxgrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrument"
)

//NoStrikes is a repeating group in InstrmtStrkPxGrp
type NoStrikes struct {
	//Instrument Component
	Instrument instrument.Component
}

//Component is a fix50 InstrmtStrkPxGrp Component
type Component struct {
	//NoStrikes is a required field for InstrmtStrkPxGrp.
	NoStrikes []NoStrikes `fix:"428"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoStrikes(v []NoStrikes) { m.NoStrikes = v }
