package cpctyconfgrp

//NoCapacities is a repeating group in CpctyConfGrp
type NoCapacities struct {
	//OrderCapacity is a required field for NoCapacities.
	OrderCapacity string `fix:"528"`
	//OrderRestrictions is a non-required field for NoCapacities.
	OrderRestrictions *string `fix:"529"`
	//OrderCapacityQty is a required field for NoCapacities.
	OrderCapacityQty float64 `fix:"863"`
}

//Component is a fix50sp2 CpctyConfGrp Component
type Component struct {
	//NoCapacities is a required field for CpctyConfGrp.
	NoCapacities []NoCapacities `fix:"862"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoCapacities(v []NoCapacities) { m.NoCapacities = v }
