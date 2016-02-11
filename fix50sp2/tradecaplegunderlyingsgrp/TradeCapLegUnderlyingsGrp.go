package tradecaplegunderlyingsgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/underlyingleginstrument"
)

//NoOfLegUnderlyings is a repeating group in TradeCapLegUnderlyingsGrp
type NoOfLegUnderlyings struct {
	//UnderlyingLegInstrument Component
	UnderlyingLegInstrument underlyingleginstrument.Component
}

//Component is a fix50sp2 TradeCapLegUnderlyingsGrp Component
type Component struct {
	//NoOfLegUnderlyings is a non-required field for TradeCapLegUnderlyingsGrp.
	NoOfLegUnderlyings []NoOfLegUnderlyings `fix:"1342,omitempty"`
}

func New() *Component { return new(Component) }
