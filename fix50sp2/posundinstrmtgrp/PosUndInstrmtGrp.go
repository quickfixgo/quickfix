package posundinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/underlyingamount"
	"github.com/quickfixgo/quickfix/fix50sp2/underlyinginstrument"
)

//NoUnderlyings is a repeating group in PosUndInstrmtGrp
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
	//UnderlyingSettlPrice is a non-required field for NoUnderlyings.
	UnderlyingSettlPrice *float64 `fix:"732"`
	//UnderlyingSettlPriceType is a non-required field for NoUnderlyings.
	UnderlyingSettlPriceType *int `fix:"733"`
	//UnderlyingAmount Component
	underlyingamount.UnderlyingAmount
	//UnderlyingDeliveryAmount is a non-required field for NoUnderlyings.
	UnderlyingDeliveryAmount *float64 `fix:"1037"`
}

//PosUndInstrmtGrp is a fix50sp2 Component
type PosUndInstrmtGrp struct {
	//NoUnderlyings is a non-required field for PosUndInstrmtGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func (m *PosUndInstrmtGrp) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
