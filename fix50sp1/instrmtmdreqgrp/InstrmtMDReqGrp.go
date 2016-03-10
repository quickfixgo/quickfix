package instrmtmdreqgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
)

//NoRelatedSym is a repeating group in InstrmtMDReqGrp
type NoRelatedSym struct {
	//Instrument Component
	instrument.Instrument
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp Component
	instrmtleggrp.InstrmtLegGrp
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//QuoteType is a non-required field for NoRelatedSym.
	QuoteType *int `fix:"537"`
	//SettlType is a non-required field for NoRelatedSym.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NoRelatedSym.
	SettlDate *string `fix:"64"`
	//MDEntrySize is a non-required field for NoRelatedSym.
	MDEntrySize *float64 `fix:"271"`
}

//InstrmtMDReqGrp is a fix50sp1 Component
type InstrmtMDReqGrp struct {
	//NoRelatedSym is a required field for InstrmtMDReqGrp.
	NoRelatedSym []NoRelatedSym `fix:"146"`
}

func (m *InstrmtMDReqGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
