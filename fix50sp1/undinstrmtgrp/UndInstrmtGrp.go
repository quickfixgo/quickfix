package undinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/underlyinginstrument"
)

//NoUnderlyings is a repeating group in UndInstrmtGrp
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//Component is a fix50sp1 UndInstrmtGrp Component
type Component struct {
	//NoUnderlyings is a non-required field for UndInstrmtGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
