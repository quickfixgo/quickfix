package strmasgnrptinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
)

func New() *StrmAsgnRptInstrmtGrp {
	var m StrmAsgnRptInstrmtGrp
	return &m
}

//NoRelatedSym is a repeating group in StrmAsgnRptInstrmtGrp
type NoRelatedSym struct {
	//Instrument is a non-required component for NoRelatedSym.
	Instrument *instrument.Instrument
	//SettlType is a non-required field for NoRelatedSym.
	SettlType *string `fix:"63"`
	//StreamAsgnType is a non-required field for NoRelatedSym.
	StreamAsgnType *int `fix:"1617"`
	//MDStreamID is a non-required field for NoRelatedSym.
	MDStreamID *string `fix:"1500"`
	//StreamAsgnRejReason is a non-required field for NoRelatedSym.
	StreamAsgnRejReason *int `fix:"1502"`
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
}

//StrmAsgnRptInstrmtGrp is a fix50sp2 Component
type StrmAsgnRptInstrmtGrp struct {
	//NoRelatedSym is a non-required field for StrmAsgnRptInstrmtGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *StrmAsgnRptInstrmtGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
