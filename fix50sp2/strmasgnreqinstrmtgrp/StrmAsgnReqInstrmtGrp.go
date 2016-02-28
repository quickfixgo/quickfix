package strmasgnreqinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
)

//NoRelatedSym is a repeating group in StrmAsgnReqInstrmtGrp
type NoRelatedSym struct {
	//Instrument Component
	Instrument instrument.Component
	//SettlType is a non-required field for NoRelatedSym.
	SettlType *string `fix:"63"`
	//MDEntrySize is a non-required field for NoRelatedSym.
	MDEntrySize *float64 `fix:"271"`
	//MDStreamID is a non-required field for NoRelatedSym.
	MDStreamID *string `fix:"1500"`
}

//Component is a fix50sp2 StrmAsgnReqInstrmtGrp Component
type Component struct {
	//NoRelatedSym is a non-required field for StrmAsgnReqInstrmtGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
