package instrmtmdreqgrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
)

//NoRelatedSym is a repeating group in InstrmtMDReqGrp
type NoRelatedSym struct {
	//Instrument Component
	Instrument instrument.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
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

//Component is a fix50 InstrmtMDReqGrp Component
type Component struct {
	//NoRelatedSym is a required field for InstrmtMDReqGrp.
	NoRelatedSym []NoRelatedSym `fix:"146"`
}

func New() *Component { return new(Component) }
