package undinstrmtcollgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/underlyinginstrument"
)

//NoUnderlyings is a repeating group in UndInstrmtCollGrp
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//CollAction is a non-required field for NoUnderlyings.
	CollAction *int `fix:"944"`
}

//Component is a fix50sp2 UndInstrmtCollGrp Component
type Component struct {
	//NoUnderlyings is a non-required field for UndInstrmtCollGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
