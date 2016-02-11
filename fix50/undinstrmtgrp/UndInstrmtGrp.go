package undinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50/underlyinginstrument"
)

//NoUnderlyings is a repeating group in UndInstrmtGrp
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//Component is a fix50 UndInstrmtGrp Component
type Component struct {
	//NoUnderlyings is a non-required field for UndInstrmtGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func New() *Component { return new(Component) }
