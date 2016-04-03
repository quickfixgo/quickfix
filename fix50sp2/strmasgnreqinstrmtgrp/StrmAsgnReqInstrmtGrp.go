package strmasgnreqinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
)

func New() *StrmAsgnReqInstrmtGrp {
	var m StrmAsgnReqInstrmtGrp
	return &m
}

//NoRelatedSym is a repeating group in StrmAsgnReqInstrmtGrp
type NoRelatedSym struct {
	//Instrument is a non-required component for NoRelatedSym.
	Instrument *instrument.Instrument
	//SettlType is a non-required field for NoRelatedSym.
	SettlType *string `fix:"63"`
	//MDEntrySize is a non-required field for NoRelatedSym.
	MDEntrySize *float64 `fix:"271"`
	//MDStreamID is a non-required field for NoRelatedSym.
	MDStreamID *string `fix:"1500"`
}

//StrmAsgnReqInstrmtGrp is a fix50sp2 Component
type StrmAsgnReqInstrmtGrp struct {
	//NoRelatedSym is a non-required field for StrmAsgnReqInstrmtGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *StrmAsgnReqInstrmtGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
