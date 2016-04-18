package rfqreqgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
)

//New returns an initialized RFQReqGrp instance
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

//NewNoRelatedSym returns an initialized NoRelatedSym instance
func NewNoRelatedSym(instrument instrument.Instrument) *NoRelatedSym {
	var m NoRelatedSym
	m.SetInstrument(instrument)
	return &m
}

func (m *NoRelatedSym) SetInstrument(v instrument.Instrument)          { m.Instrument = v }
func (m *NoRelatedSym) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
func (m *NoRelatedSym) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp) { m.InstrmtLegGrp = &v }
func (m *NoRelatedSym) SetPrevClosePx(v float64)                       { m.PrevClosePx = &v }
func (m *NoRelatedSym) SetQuoteRequestType(v int)                      { m.QuoteRequestType = &v }
func (m *NoRelatedSym) SetQuoteType(v int)                             { m.QuoteType = &v }
func (m *NoRelatedSym) SetTradingSessionID(v string)                   { m.TradingSessionID = &v }
func (m *NoRelatedSym) SetTradingSessionSubID(v string)                { m.TradingSessionSubID = &v }

//RFQReqGrp is a fix50sp1 Component
type RFQReqGrp struct {
	//NoRelatedSym is a required field for RFQReqGrp.
	NoRelatedSym []NoRelatedSym `fix:"146"`
}

func (m *RFQReqGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
