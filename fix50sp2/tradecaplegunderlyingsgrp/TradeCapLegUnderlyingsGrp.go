package tradecaplegunderlyingsgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/underlyingleginstrument"
)

//New returns an initialized TradeCapLegUnderlyingsGrp instance
func New() *TradeCapLegUnderlyingsGrp {
	var m TradeCapLegUnderlyingsGrp
	return &m
}

//NoOfLegUnderlyings is a repeating group in TradeCapLegUnderlyingsGrp
type NoOfLegUnderlyings struct {
	//UnderlyingLegInstrument is a non-required component for NoOfLegUnderlyings.
	UnderlyingLegInstrument *underlyingleginstrument.UnderlyingLegInstrument
}

//NewNoOfLegUnderlyings returns an initialized NoOfLegUnderlyings instance
func NewNoOfLegUnderlyings() *NoOfLegUnderlyings {
	var m NoOfLegUnderlyings
	return &m
}

func (m *NoOfLegUnderlyings) SetUnderlyingLegInstrument(v underlyingleginstrument.UnderlyingLegInstrument) {
	m.UnderlyingLegInstrument = &v
}

//TradeCapLegUnderlyingsGrp is a fix50sp2 Component
type TradeCapLegUnderlyingsGrp struct {
	//NoOfLegUnderlyings is a non-required field for TradeCapLegUnderlyingsGrp.
	NoOfLegUnderlyings []NoOfLegUnderlyings `fix:"1342,omitempty"`
}

func (m *TradeCapLegUnderlyingsGrp) SetNoOfLegUnderlyings(v []NoOfLegUnderlyings) {
	m.NoOfLegUnderlyings = v
}
