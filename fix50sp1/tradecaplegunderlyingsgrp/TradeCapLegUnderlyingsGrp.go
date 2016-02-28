package tradecaplegunderlyingsgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/underlyingleginstrument"
)

//NoOfLegUnderlyings is a repeating group in TradeCapLegUnderlyingsGrp
type NoOfLegUnderlyings struct {
	//UnderlyingLegInstrument Component
	UnderlyingLegInstrument underlyingleginstrument.Component
}

//Component is a fix50sp1 TradeCapLegUnderlyingsGrp Component
type Component struct {
	//NoOfLegUnderlyings is a non-required field for TradeCapLegUnderlyingsGrp.
	NoOfLegUnderlyings []NoOfLegUnderlyings `fix:"1342,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoOfLegUnderlyings(v []NoOfLegUnderlyings) { m.NoOfLegUnderlyings = v }
