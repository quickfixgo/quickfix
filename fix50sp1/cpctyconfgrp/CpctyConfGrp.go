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

func (m *NoCapacities) SetOrderCapacity(v string)     { m.OrderCapacity = v }
func (m *NoCapacities) SetOrderRestrictions(v string) { m.OrderRestrictions = &v }
func (m *NoCapacities) SetOrderCapacityQty(v float64) { m.OrderCapacityQty = v }

//CpctyConfGrp is a fix50sp1 Component
type CpctyConfGrp struct {
	//NoCapacities is a required field for CpctyConfGrp.
	NoCapacities []NoCapacities `fix:"862"`
}

func (m *CpctyConfGrp) SetNoCapacities(v []NoCapacities) { m.NoCapacities = v }
