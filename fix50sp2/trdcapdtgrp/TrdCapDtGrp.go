package trdcapdtgrp

import (
	"time"
)

//NoDates is a repeating group in TrdCapDtGrp
type NoDates struct {
	//TradeDate is a non-required field for NoDates.
	TradeDate *string `fix:"75"`
	//TransactTime is a non-required field for NoDates.
	TransactTime *time.Time `fix:"60"`
	//LastUpdateTime is a non-required field for NoDates.
	LastUpdateTime *time.Time `fix:"779"`
}

//Component is a fix50sp2 TrdCapDtGrp Component
type Component struct {
	//NoDates is a non-required field for TrdCapDtGrp.
	NoDates []NoDates `fix:"580,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoDates(v []NoDates) { m.NoDates = v }
