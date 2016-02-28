package instrmtstrkpxgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
)

//NoStrikes is a repeating group in InstrmtStrkPxGrp
type NoStrikes struct {
	//Instrument Component
	Instrument instrument.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//PrevClosePx is a non-required field for NoStrikes.
	PrevClosePx *float64 `fix:"140"`
	//ClOrdID is a non-required field for NoStrikes.
	ClOrdID *string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoStrikes.
	SecondaryClOrdID *string `fix:"526"`
	//Side is a non-required field for NoStrikes.
	Side *string `fix:"54"`
	//Price is a non-required field for NoStrikes.
	Price *float64 `fix:"44"`
	//Currency is a non-required field for NoStrikes.
	Currency *string `fix:"15"`
	//Text is a non-required field for NoStrikes.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoStrikes.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoStrikes.
	EncodedText *string `fix:"355"`
}

//Component is a fix50sp1 InstrmtStrkPxGrp Component
type Component struct {
	//NoStrikes is a required field for InstrmtStrkPxGrp.
	NoStrikes []NoStrikes `fix:"428"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoStrikes(v []NoStrikes) { m.NoStrikes = v }
