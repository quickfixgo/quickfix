package instrmtstrkpxgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
)

//NoStrikes is a repeating group in InstrmtStrkPxGrp
type NoStrikes struct {
	//Instrument is a required component for NoStrikes.
	instrument.Instrument
	//UndInstrmtGrp is a non-required component for NoStrikes.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
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

//InstrmtStrkPxGrp is a fix50sp2 Component
type InstrmtStrkPxGrp struct {
	//NoStrikes is a required field for InstrmtStrkPxGrp.
	NoStrikes []NoStrikes `fix:"428"`
}

func (m *InstrmtStrkPxGrp) SetNoStrikes(v []NoStrikes) { m.NoStrikes = v }
