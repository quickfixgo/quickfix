package relsymderivsecupdgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp2/secondarypricelimits"
	"time"
)

//NoRelatedSym is a repeating group in RelSymDerivSecUpdGrp
type NoRelatedSym struct {
	//ListUpdateAction is a non-required field for NoRelatedSym.
	ListUpdateAction *string `fix:"1324"`
	//Instrument is a non-required component for NoRelatedSym.
	Instrument *instrument.Instrument
	//InstrumentExtension is a non-required component for NoRelatedSym.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//SecondaryPriceLimits is a non-required component for NoRelatedSym.
	SecondaryPriceLimits *secondarypricelimits.SecondaryPriceLimits
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//InstrmtLegGrp is a non-required component for NoRelatedSym.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
	//CorporateAction is a non-required field for NoRelatedSym.
	CorporateAction *string `fix:"292"`
	//RelSymTransactTime is a non-required field for NoRelatedSym.
	RelSymTransactTime *time.Time `fix:"1504"`
}

//RelSymDerivSecUpdGrp is a fix50sp2 Component
type RelSymDerivSecUpdGrp struct {
	//NoRelatedSym is a non-required field for RelSymDerivSecUpdGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *RelSymDerivSecUpdGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
