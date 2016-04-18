package strmasgnreqinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
)

//New returns an initialized StrmAsgnReqInstrmtGrp instance
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

//NewNoRelatedSym returns an initialized NoRelatedSym instance
func NewNoRelatedSym() *NoRelatedSym {
	var m NoRelatedSym
	return &m
}

func (m *NoRelatedSym) SetInstrument(v instrument.Instrument) { m.Instrument = &v }
func (m *NoRelatedSym) SetSettlType(v string)                 { m.SettlType = &v }
func (m *NoRelatedSym) SetMDEntrySize(v float64)              { m.MDEntrySize = &v }
func (m *NoRelatedSym) SetMDStreamID(v string)                { m.MDStreamID = &v }

//StrmAsgnReqInstrmtGrp is a fix50sp2 Component
type StrmAsgnReqInstrmtGrp struct {
	//NoRelatedSym is a non-required field for StrmAsgnReqInstrmtGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *StrmAsgnReqInstrmtGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
