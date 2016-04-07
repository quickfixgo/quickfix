package trdrepindicatorsgrp

func New() *TrdRepIndicatorsGrp {
	var m TrdRepIndicatorsGrp
	return &m
}

//NoTrdRepIndicators is a repeating group in TrdRepIndicatorsGrp
type NoTrdRepIndicators struct {
	//TrdRepPartyRole is a non-required field for NoTrdRepIndicators.
	TrdRepPartyRole *int `fix:"1388"`
	//TrdRepIndicator is a non-required field for NoTrdRepIndicators.
	TrdRepIndicator *bool `fix:"1389"`
}

func (m *NoTrdRepIndicators) SetTrdRepPartyRole(v int)  { m.TrdRepPartyRole = &v }
func (m *NoTrdRepIndicators) SetTrdRepIndicator(v bool) { m.TrdRepIndicator = &v }

//TrdRepIndicatorsGrp is a fix50sp2 Component
type TrdRepIndicatorsGrp struct {
	//NoTrdRepIndicators is a non-required field for TrdRepIndicatorsGrp.
	NoTrdRepIndicators []NoTrdRepIndicators `fix:"1387,omitempty"`
}

func (m *TrdRepIndicatorsGrp) SetNoTrdRepIndicators(v []NoTrdRepIndicators) { m.NoTrdRepIndicators = v }
