package undinstrmtstrkpxgrp

import (
	"github.com/quickfixgo/quickfix/fix50/underlyinginstrument"
)

func New() *UndInstrmtStrkPxGrp {
	var m UndInstrmtStrkPxGrp
	return &m
}

//NoUnderlyings is a repeating group in UndInstrmtStrkPxGrp
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
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

func (m *NoUnderlyings) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}
func (m *NoUnderlyings) SetPrevClosePx(v float64)     { m.PrevClosePx = &v }
func (m *NoUnderlyings) SetClOrdID(v string)          { m.ClOrdID = &v }
func (m *NoUnderlyings) SetSecondaryClOrdID(v string) { m.SecondaryClOrdID = &v }
func (m *NoUnderlyings) SetSide(v string)             { m.Side = &v }
func (m *NoUnderlyings) SetPrice(v float64)           { m.Price = v }
func (m *NoUnderlyings) SetCurrency(v string)         { m.Currency = &v }
func (m *NoUnderlyings) SetText(v string)             { m.Text = &v }
func (m *NoUnderlyings) SetEncodedTextLen(v int)      { m.EncodedTextLen = &v }
func (m *NoUnderlyings) SetEncodedText(v string)      { m.EncodedText = &v }

//UndInstrmtStrkPxGrp is a fix50 Component
type UndInstrmtStrkPxGrp struct {
	//NoUnderlyings is a non-required field for UndInstrmtStrkPxGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func (m *UndInstrmtStrkPxGrp) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
