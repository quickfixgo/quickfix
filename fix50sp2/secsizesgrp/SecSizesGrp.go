package secsizesgrp

func New() *SecSizesGrp {
	var m SecSizesGrp
	return &m
}

//NoOfSecSizes is a repeating group in SecSizesGrp
type NoOfSecSizes struct {
	//MDSecSizeType is a non-required field for NoOfSecSizes.
	MDSecSizeType *int `fix:"1178"`
	//MDSecSize is a non-required field for NoOfSecSizes.
	MDSecSize *float64 `fix:"1179"`
}

func (m *NoOfSecSizes) SetMDSecSizeType(v int) { m.MDSecSizeType = &v }
func (m *NoOfSecSizes) SetMDSecSize(v float64) { m.MDSecSize = &v }

//SecSizesGrp is a fix50sp2 Component
type SecSizesGrp struct {
	//NoOfSecSizes is a non-required field for SecSizesGrp.
	NoOfSecSizes []NoOfSecSizes `fix:"1177,omitempty"`
}

func (m *SecSizesGrp) SetNoOfSecSizes(v []NoOfSecSizes) { m.NoOfSecSizes = v }
