package applidreportgrp

//New returns an initialized ApplIDReportGrp instance
func New() *ApplIDReportGrp {
	var m ApplIDReportGrp
	return &m
}

//NoApplIDs is a repeating group in ApplIDReportGrp
type NoApplIDs struct {
	//RefApplID is a non-required field for NoApplIDs.
	RefApplID *string `fix:"1355"`
	//ApplNewSeqNum is a non-required field for NoApplIDs.
	ApplNewSeqNum *int `fix:"1399"`
	//RefApplLastSeqNum is a non-required field for NoApplIDs.
	RefApplLastSeqNum *int `fix:"1357"`
}

//NewNoApplIDs returns an initialized NoApplIDs instance
func NewNoApplIDs() *NoApplIDs {
	var m NoApplIDs
	return &m
}

func (m *NoApplIDs) SetRefApplID(v string)      { m.RefApplID = &v }
func (m *NoApplIDs) SetApplNewSeqNum(v int)     { m.ApplNewSeqNum = &v }
func (m *NoApplIDs) SetRefApplLastSeqNum(v int) { m.RefApplLastSeqNum = &v }

//ApplIDReportGrp is a fix50sp2 Component
type ApplIDReportGrp struct {
	//NoApplIDs is a non-required field for ApplIDReportGrp.
	NoApplIDs []NoApplIDs `fix:"1351,omitempty"`
}

func (m *ApplIDReportGrp) SetNoApplIDs(v []NoApplIDs) { m.NoApplIDs = v }
