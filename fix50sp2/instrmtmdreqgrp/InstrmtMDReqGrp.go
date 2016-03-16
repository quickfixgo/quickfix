package instrmtmdreqgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
)

//NoRelatedSym is a repeating group in InstrmtMDReqGrp
type NoRelatedSym struct {
	//Instrument is a required component for NoRelatedSym.
	instrument.Instrument
	//UndInstrmtGrp is a non-required component for NoRelatedSym.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp is a non-required component for NoRelatedSym.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
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
	//MDStreamID is a non-required field for NoRelatedSym.
	MDStreamID *string `fix:"1500"`
}

//InstrmtMDReqGrp is a fix50sp2 Component
type InstrmtMDReqGrp struct {
	//NoRelatedSym is a required field for InstrmtMDReqGrp.
	NoRelatedSym []NoRelatedSym `fix:"146"`
}

func (m *InstrmtMDReqGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
