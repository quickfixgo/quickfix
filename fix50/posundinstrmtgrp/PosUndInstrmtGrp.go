package posundinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50/underlyingamount"
	"github.com/quickfixgo/quickfix/fix50/underlyinginstrument"
)

//NoUnderlyings is a repeating group in PosUndInstrmtGrp
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//UnderlyingSettlPrice is a non-required field for NoUnderlyings.
	UnderlyingSettlPrice *float64 `fix:"732"`
	//UnderlyingSettlPriceType is a non-required field for NoUnderlyings.
	UnderlyingSettlPriceType *int `fix:"733"`
	//UnderlyingAmount Component
	UnderlyingAmount underlyingamount.Component
	//UnderlyingDeliveryAmount is a non-required field for NoUnderlyings.
	UnderlyingDeliveryAmount *float64 `fix:"1037"`
}

//Component is a fix50 PosUndInstrmtGrp Component
type Component struct {
	//NoUnderlyings is a non-required field for PosUndInstrmtGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
