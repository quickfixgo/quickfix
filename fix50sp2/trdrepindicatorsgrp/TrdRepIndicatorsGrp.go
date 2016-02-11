package trdrepindicatorsgrp

//NoTrdRepIndicators is a repeating group in TrdRepIndicatorsGrp
type NoTrdRepIndicators struct {
	//TrdRepPartyRole is a non-required field for NoTrdRepIndicators.
	TrdRepPartyRole *int `fix:"1388"`
	//TrdRepIndicator is a non-required field for NoTrdRepIndicators.
	TrdRepIndicator *bool `fix:"1389"`
}

//Component is a fix50sp2 TrdRepIndicatorsGrp Component
type Component struct {
	//NoTrdRepIndicators is a non-required field for TrdRepIndicatorsGrp.
	NoTrdRepIndicators []NoTrdRepIndicators `fix:"1387,omitempty"`
}

func New() *Component { return new(Component) }
