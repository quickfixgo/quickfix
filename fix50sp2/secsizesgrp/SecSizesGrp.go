package secsizesgrp

//NoOfSecSizes is a repeating group in SecSizesGrp
type NoOfSecSizes struct {
	//MDSecSizeType is a non-required field for NoOfSecSizes.
	MDSecSizeType *int `fix:"1178"`
	//MDSecSize is a non-required field for NoOfSecSizes.
	MDSecSize *float64 `fix:"1179"`
}

//Component is a fix50sp2 SecSizesGrp Component
type Component struct {
	//NoOfSecSizes is a non-required field for SecSizesGrp.
	NoOfSecSizes []NoOfSecSizes `fix:"1177,omitempty"`
}

func New() *Component { return new(Component) }
