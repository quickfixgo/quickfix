package instrmtstrkpxgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
)

func New(nostrikes []NoStrikes) *InstrmtStrkPxGrp {
	var m InstrmtStrkPxGrp
	m.SetNoStrikes(nostrikes)
	return &m
}

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

func (m *NoStrikes) SetInstrument(v instrument.Instrument)          { m.Instrument = v }
func (m *NoStrikes) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
func (m *NoStrikes) SetPrevClosePx(v float64)                       { m.PrevClosePx = &v }
func (m *NoStrikes) SetClOrdID(v string)                            { m.ClOrdID = &v }
func (m *NoStrikes) SetSecondaryClOrdID(v string)                   { m.SecondaryClOrdID = &v }
func (m *NoStrikes) SetSide(v string)                               { m.Side = &v }
func (m *NoStrikes) SetPrice(v float64)                             { m.Price = &v }
func (m *NoStrikes) SetCurrency(v string)                           { m.Currency = &v }
func (m *NoStrikes) SetText(v string)                               { m.Text = &v }
func (m *NoStrikes) SetEncodedTextLen(v int)                        { m.EncodedTextLen = &v }
func (m *NoStrikes) SetEncodedText(v string)                        { m.EncodedText = &v }

//InstrmtStrkPxGrp is a fix50sp2 Component
type InstrmtStrkPxGrp struct {
	//NoStrikes is a required field for InstrmtStrkPxGrp.
	NoStrikes []NoStrikes `fix:"428"`
}

func (m *InstrmtStrkPxGrp) SetNoStrikes(v []NoStrikes) { m.NoStrikes = v }
