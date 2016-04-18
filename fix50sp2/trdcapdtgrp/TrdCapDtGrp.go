package trdcapdtgrp

import (
	"time"
)

//New returns an initialized TrdCapDtGrp instance
func New() *TrdCapDtGrp {
	var m TrdCapDtGrp
	return &m
}

//NoDates is a repeating group in TrdCapDtGrp
type NoDates struct {
	//TradeDate is a non-required field for NoDates.
	TradeDate *string `fix:"75"`
	//TransactTime is a non-required field for NoDates.
	TransactTime *time.Time `fix:"60"`
	//LastUpdateTime is a non-required field for NoDates.
	LastUpdateTime *time.Time `fix:"779"`
}

//NewNoDates returns an initialized NoDates instance
func NewNoDates() *NoDates {
	var m NoDates
	return &m
}

func (m *NoDates) SetTradeDate(v string)         { m.TradeDate = &v }
func (m *NoDates) SetTransactTime(v time.Time)   { m.TransactTime = &v }
func (m *NoDates) SetLastUpdateTime(v time.Time) { m.LastUpdateTime = &v }

//TrdCapDtGrp is a fix50sp2 Component
type TrdCapDtGrp struct {
	//NoDates is a non-required field for TrdCapDtGrp.
	NoDates []NoDates `fix:"580,omitempty"`
}

func (m *TrdCapDtGrp) SetNoDates(v []NoDates) { m.NoDates = v }
