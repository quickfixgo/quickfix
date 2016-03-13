package tradecaplegunderlyingsgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/underlyingleginstrument"
)

//NoOfLegUnderlyings is a repeating group in TradeCapLegUnderlyingsGrp
type NoOfLegUnderlyings struct {
	//UnderlyingLegInstrument is a non-required component for NoOfLegUnderlyings.
	UnderlyingLegInstrument *underlyingleginstrument.UnderlyingLegInstrument
}

//TradeCapLegUnderlyingsGrp is a fix50sp1 Component
type TradeCapLegUnderlyingsGrp struct {
	//NoOfLegUnderlyings is a non-required field for TradeCapLegUnderlyingsGrp.
	NoOfLegUnderlyings []NoOfLegUnderlyings `fix:"1342,omitempty"`
}

func (m *TradeCapLegUnderlyingsGrp) SetNoOfLegUnderlyings(v []NoOfLegUnderlyings) {
	m.NoOfLegUnderlyings = v
}
