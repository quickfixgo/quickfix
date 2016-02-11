package compidreqgrp

//NoCompIDs is a repeating group in CompIDReqGrp
type NoCompIDs struct {
	//RefCompID is a non-required field for NoCompIDs.
	RefCompID *string `fix:"930"`
	//RefSubID is a non-required field for NoCompIDs.
	RefSubID *string `fix:"931"`
	//LocationID is a non-required field for NoCompIDs.
	LocationID *string `fix:"283"`
	//DeskID is a non-required field for NoCompIDs.
	DeskID *string `fix:"284"`
}

//Component is a fix50sp1 CompIDReqGrp Component
type Component struct {
	//NoCompIDs is a non-required field for CompIDReqGrp.
	NoCompIDs []NoCompIDs `fix:"936,omitempty"`
}

func New() *Component { return new(Component) }
