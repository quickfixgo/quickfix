package posundinstrmtgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/underlyingamount"
	"github.com/quickfixgo/quickfix/fix50sp2/underlyinginstrument"
)

//New returns an initialized PosUndInstrmtGrp instance
func New() *PosUndInstrmtGrp {
	var m PosUndInstrmtGrp
	return &m
}

//NoUnderlyings is a repeating group in PosUndInstrmtGrp
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
	//UnderlyingSettlPrice is a non-required field for NoUnderlyings.
	UnderlyingSettlPrice *float64 `fix:"732"`
	//UnderlyingSettlPriceType is a non-required field for NoUnderlyings.
	UnderlyingSettlPriceType *int `fix:"733"`
	//UnderlyingAmount is a non-required component for NoUnderlyings.
	UnderlyingAmount *underlyingamount.UnderlyingAmount
	//UnderlyingDeliveryAmount is a non-required field for NoUnderlyings.
	UnderlyingDeliveryAmount *float64 `fix:"1037"`
}

//NewNoUnderlyings returns an initialized NoUnderlyings instance
func NewNoUnderlyings() *NoUnderlyings {
	var m NoUnderlyings
	return &m
}

func (m *NoUnderlyings) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}
func (m *NoUnderlyings) SetUnderlyingSettlPrice(v float64) { m.UnderlyingSettlPrice = &v }
func (m *NoUnderlyings) SetUnderlyingSettlPriceType(v int) { m.UnderlyingSettlPriceType = &v }
func (m *NoUnderlyings) SetUnderlyingAmount(v underlyingamount.UnderlyingAmount) {
	m.UnderlyingAmount = &v
}
func (m *NoUnderlyings) SetUnderlyingDeliveryAmount(v float64) { m.UnderlyingDeliveryAmount = &v }

//PosUndInstrmtGrp is a fix50sp2 Component
type PosUndInstrmtGrp struct {
	//NoUnderlyings is a non-required field for PosUndInstrmtGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func (m *PosUndInstrmtGrp) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
