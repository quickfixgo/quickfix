package undinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/underlyinginstrument"
)

func New() *UndInstrmtGrp {
	var m UndInstrmtGrp
	return &m
}

//NoUnderlyings is a repeating group in UndInstrmtGrp
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
}

//UndInstrmtGrp is a fix50sp1 Component
type UndInstrmtGrp struct {
	//NoUnderlyings is a non-required field for UndInstrmtGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func (m *UndInstrmtGrp) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
