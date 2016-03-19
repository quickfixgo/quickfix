package cpctyconfgrp

func New(nocapacities []NoCapacities) *CpctyConfGrp {
	var m CpctyConfGrp
	m.SetNoCapacities(nocapacities)
	return &m
}

//NoCapacities is a repeating group in CpctyConfGrp
type NoCapacities struct {
	//OrderCapacity is a required field for NoCapacities.
	OrderCapacity string `fix:"528"`
	//OrderRestrictions is a non-required field for NoCapacities.
	OrderRestrictions *string `fix:"529"`
	//OrderCapacityQty is a required field for NoCapacities.
	OrderCapacityQty float64 `fix:"863"`
}

//CpctyConfGrp is a fix50 Component
type CpctyConfGrp struct {
	//NoCapacities is a required field for CpctyConfGrp.
	NoCapacities []NoCapacities `fix:"862"`
}

func (m *CpctyConfGrp) SetNoCapacities(v []NoCapacities) { m.NoCapacities = v }
