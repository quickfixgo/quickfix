package relsymderivsecgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp2/secondarypricelimits"
	"time"
)

func New() *RelSymDerivSecGrp {
	var m RelSymDerivSecGrp
	return &m
}

//NoRelatedSym is a repeating group in RelSymDerivSecGrp
type NoRelatedSym struct {
	//Instrument is a non-required component for NoRelatedSym.
	Instrument *instrument.Instrument
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//InstrumentExtension is a non-required component for NoRelatedSym.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//InstrmtLegGrp is a non-required component for NoRelatedSym.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
	//SecondaryPriceLimits is a non-required component for NoRelatedSym.
	SecondaryPriceLimits *secondarypricelimits.SecondaryPriceLimits
	//CorporateAction is a non-required field for NoRelatedSym.
	CorporateAction *string `fix:"292"`
	//RelSymTransactTime is a non-required field for NoRelatedSym.
	RelSymTransactTime *time.Time `fix:"1504"`
}

//RelSymDerivSecGrp is a fix50sp2 Component
type RelSymDerivSecGrp struct {
	//NoRelatedSym is a non-required field for RelSymDerivSecGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *RelSymDerivSecGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
