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

//TrdCapDtGrp is a fix50sp2 Component
type TrdCapDtGrp struct {
	//NoDates is a non-required field for TrdCapDtGrp.
	NoDates []NoDates `fix:"580,omitempty"`
}

func (m *TrdCapDtGrp) SetNoDates(v []NoDates) { m.NoDates = v }
