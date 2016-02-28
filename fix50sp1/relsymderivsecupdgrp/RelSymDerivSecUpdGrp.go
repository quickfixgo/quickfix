package relsymderivsecupdgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp1/secondarypricelimits"
)

//NoRelatedSym is a repeating group in RelSymDerivSecUpdGrp
type NoRelatedSym struct {
	//ListUpdateAction is a non-required field for NoRelatedSym.
	ListUpdateAction *string `fix:"1324"`
	//Instrument Component
	Instrument instrument.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//SecondaryPriceLimits Component
	SecondaryPriceLimits secondarypricelimits.Component
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
	//CorporateAction is a non-required field for NoRelatedSym.
	CorporateAction *string `fix:"292"`
}

//Component is a fix50sp1 RelSymDerivSecUpdGrp Component
type Component struct {
	//NoRelatedSym is a non-required field for RelSymDerivSecUpdGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
