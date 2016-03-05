package undinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/underlyinginstrument"
)

//NoUnderlyings is a repeating group in UndInstrmtGrp
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
}

//UndInstrmtGrp is a fix50sp2 Component
type UndInstrmtGrp struct {
	//NoUnderlyings is a non-required field for UndInstrmtGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func (m *UndInstrmtGrp) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
