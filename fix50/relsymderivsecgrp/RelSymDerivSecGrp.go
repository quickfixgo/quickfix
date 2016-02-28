package relsymderivsecgrp

import (
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/instrumentextension"
)

//NoRelatedSym is a repeating group in RelSymDerivSecGrp
type NoRelatedSym struct {
	//Instrument Component
	Instrument instrument.Component
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//ExpirationCycle is a non-required field for NoRelatedSym.
	ExpirationCycle *int `fix:"827"`
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//TradingSessionID is a non-required field for NoRelatedSym.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoRelatedSym.
	TradingSessionSubID *string `fix:"625"`
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
}

//Component is a fix50 RelSymDerivSecGrp Component
type Component struct {
	//NoRelatedSym is a non-required field for RelSymDerivSecGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
