package rfqreqgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
)

//NoRelatedSym is a repeating group in RFQReqGrp
type NoRelatedSym struct {
	//Instrument Component
	instrument.Instrument
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp Component
	instrmtleggrp.InstrmtLegGrp
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

//RFQReqGrp is a fix50sp2 Component
type RFQReqGrp struct {
	//NoRelatedSym is a required field for RFQReqGrp.
	NoRelatedSym []NoRelatedSym `fix:"146"`
}

func (m *RFQReqGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
