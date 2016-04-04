package applidrequestgrp

func New() *ApplIDRequestGrp {
	var m ApplIDRequestGrp
	return &m
}

//NoApplIDs is a repeating group in ApplIDRequestGrp
type NoApplIDs struct {
	//RefApplID is a non-required field for NoApplIDs.
	RefApplID *string `fix:"1355"`
	//ApplBegSeqNum is a non-required field for NoApplIDs.
	ApplBegSeqNum *int `fix:"1182"`
	//ApplEndSeqNum is a non-required field for NoApplIDs.
	ApplEndSeqNum *int `fix:"1183"`
}

func (m *NoApplIDs) SetRefApplID(v string)  { m.RefApplID = &v }
func (m *NoApplIDs) SetApplBegSeqNum(v int) { m.ApplBegSeqNum = &v }
func (m *NoApplIDs) SetApplEndSeqNum(v int) { m.ApplEndSeqNum = &v }

//ApplIDRequestGrp is a fix50sp1 Component
type ApplIDRequestGrp struct {
	//NoApplIDs is a non-required field for ApplIDRequestGrp.
	NoApplIDs []NoApplIDs `fix:"1351,omitempty"`
}

func (m *ApplIDRequestGrp) SetNoApplIDs(v []NoApplIDs) { m.NoApplIDs = v }
