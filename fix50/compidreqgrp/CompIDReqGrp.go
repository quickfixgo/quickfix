package compidreqgrp

//New returns an initialized CompIDReqGrp instance
func New() *CompIDReqGrp {
	var m CompIDReqGrp
	return &m
}

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

//NewNoCompIDs returns an initialized NoCompIDs instance
func NewNoCompIDs() *NoCompIDs {
	var m NoCompIDs
	return &m
}

func (m *NoCompIDs) SetRefCompID(v string)  { m.RefCompID = &v }
func (m *NoCompIDs) SetRefSubID(v string)   { m.RefSubID = &v }
func (m *NoCompIDs) SetLocationID(v string) { m.LocationID = &v }
func (m *NoCompIDs) SetDeskID(v string)     { m.DeskID = &v }

//CompIDReqGrp is a fix50 Component
type CompIDReqGrp struct {
	//NoCompIDs is a non-required field for CompIDReqGrp.
	NoCompIDs []NoCompIDs `fix:"936,omitempty"`
}

func (m *CompIDReqGrp) SetNoCompIDs(v []NoCompIDs) { m.NoCompIDs = v }
