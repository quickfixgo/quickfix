package undinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/underlyinginstrument"
)

//New returns an initialized UndInstrmtGrp instance
func New() *UndInstrmtGrp {
	var m UndInstrmtGrp
	return &m
}

//NoUnderlyings is a repeating group in UndInstrmtGrp
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
}

//NewNoUnderlyings returns an initialized NoUnderlyings instance
func NewNoUnderlyings() *NoUnderlyings {
	var m NoUnderlyings
	return &m
}

func (m *NoUnderlyings) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}

//UndInstrmtGrp is a fix50sp2 Component
type UndInstrmtGrp struct {
	//NoUnderlyings is a non-required field for UndInstrmtGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func (m *UndInstrmtGrp) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
