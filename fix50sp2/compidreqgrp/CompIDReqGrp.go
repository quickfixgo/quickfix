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

//CompIDReqGrp is a fix50sp2 Component
type CompIDReqGrp struct {
	//NoCompIDs is a non-required field for CompIDReqGrp.
	NoCompIDs []NoCompIDs `fix:"936,omitempty"`
}

func (m *CompIDReqGrp) SetNoCompIDs(v []NoCompIDs) { m.NoCompIDs = v }
