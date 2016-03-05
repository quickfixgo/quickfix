package undinstrmtstrkpxgrp

import (
	"github.com/quickfixgo/quickfix/fix50/underlyinginstrument"
)

//NoUnderlyings is a repeating group in UndInstrmtStrkPxGrp
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
	//PrevClosePx is a non-required field for NoUnderlyings.
	PrevClosePx *float64 `fix:"140"`
	//ClOrdID is a non-required field for NoUnderlyings.
	ClOrdID *string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoUnderlyings.
	SecondaryClOrdID *string `fix:"526"`
	//Side is a non-required field for NoUnderlyings.
	Side *string `fix:"54"`
	//Price is a required field for NoUnderlyings.
	Price float64 `fix:"44"`
	//Currency is a non-required field for NoUnderlyings.
	Currency *string `fix:"15"`
	//Text is a non-required field for NoUnderlyings.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoUnderlyings.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoUnderlyings.
	EncodedText *string `fix:"355"`
}

//UndInstrmtStrkPxGrp is a fix50 Component
type UndInstrmtStrkPxGrp struct {
	//NoUnderlyings is a non-required field for UndInstrmtStrkPxGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func (m *UndInstrmtStrkPxGrp) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
