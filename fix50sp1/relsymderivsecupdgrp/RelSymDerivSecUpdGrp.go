package relsymderivsecupdgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp1/secondarypricelimits"
)

func New() *RelSymDerivSecUpdGrp {
	var m RelSymDerivSecUpdGrp
	return &m
}

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
}

func (m *NoRelatedSym) SetListUpdateAction(v string)          { m.ListUpdateAction = &v }
func (m *NoRelatedSym) SetInstrument(v instrument.Instrument) { m.Instrument = &v }
func (m *NoRelatedSym) SetInstrumentExtension(v instrumentextension.InstrumentExtension) {
	m.InstrumentExtension = &v
}
func (m *NoRelatedSym) SetSecondaryPriceLimits(v secondarypricelimits.SecondaryPriceLimits) {
	m.SecondaryPriceLimits = &v
}
func (m *NoRelatedSym) SetCurrency(v string)                           { m.Currency = &v }
func (m *NoRelatedSym) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp) { m.InstrmtLegGrp = &v }
func (m *NoRelatedSym) SetText(v string)                               { m.Text = &v }
func (m *NoRelatedSym) SetEncodedTextLen(v int)                        { m.EncodedTextLen = &v }
func (m *NoRelatedSym) SetEncodedText(v string)                        { m.EncodedText = &v }
func (m *NoRelatedSym) SetCorporateAction(v string)                    { m.CorporateAction = &v }

//RelSymDerivSecUpdGrp is a fix50sp1 Component
type RelSymDerivSecUpdGrp struct {
	//NoRelatedSym is a non-required field for RelSymDerivSecUpdGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func (m *RelSymDerivSecUpdGrp) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
