package rfqreqgrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
)

func New(norelatedsym []NoRelatedSym) *RFQReqGrp {
	var m RFQReqGrp
	m.SetNoRelatedSym(norelatedsym)
	return &m
}

//NoRelatedSym is a repeating group in RFQReqGrp
type NoRelatedSym struct {
	//Instrument is a required component for NoRelatedSym.
	instrument.Instrument
	//UndInstrmtGrp is a non-required component for NoRelatedSym.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp is a non-required component for NoRelatedSym.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//PrevClosePx is a non-required field for NoRelatedSym.
	PrevClosePx *float64 `fix:"140"`
	//QuoteRequestType is a non-required field for NoRelatedSym.
	QuoteRequestType *int `fix:"303"`
	//QuoteType is a non-required field for NoRelatedSym.
	QuoteType *int `fix:"537"`
	//TradingSessionID is a non-required field for NoRelatedSym.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoRelatedSym.
	TradingSessionSubID *string `fix:"625"`
}

//RFQReqGrp is a fix50 Component
type RFQReqGrp struct {
	//NoRelatedSym is a required field for RFQReqGrp.
	NoRelatedSym []NoRelatedSym `fix:"146"`
}

func (m *RFQReqGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
