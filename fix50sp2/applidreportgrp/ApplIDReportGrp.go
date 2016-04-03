package applidreportgrp

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

//ApplIDReportGrp is a fix50sp2 Component
type ApplIDReportGrp struct {
	//NoApplIDs is a non-required field for ApplIDReportGrp.
	NoApplIDs []NoApplIDs `fix:"1351,omitempty"`
}

func (m *ApplIDReportGrp) SetNoApplIDs(v []NoApplIDs) { m.NoApplIDs = v }
