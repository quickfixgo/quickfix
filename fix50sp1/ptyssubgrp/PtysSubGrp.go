package ptyssubgrp

//New returns an initialized PtysSubGrp instance
func New() *PtysSubGrp {
	var m PtysSubGrp
	return &m
}

//NoPartySubIDs is a repeating group in PtysSubGrp
type NoPartySubIDs struct {
	//PartySubID is a non-required field for NoPartySubIDs.
	PartySubID *string `fix:"523"`
	//PartySubIDType is a non-required field for NoPartySubIDs.
	PartySubIDType *int `fix:"803"`
}

//NewNoPartySubIDs returns an initialized NoPartySubIDs instance
func NewNoPartySubIDs() *NoPartySubIDs {
	var m NoPartySubIDs
	return &m
}

func (m *NoPartySubIDs) SetPartySubID(v string)  { m.PartySubID = &v }
func (m *NoPartySubIDs) SetPartySubIDType(v int) { m.PartySubIDType = &v }

//PtysSubGrp is a fix50sp1 Component
type PtysSubGrp struct {
	//NoPartySubIDs is a non-required field for PtysSubGrp.
	NoPartySubIDs []NoPartySubIDs `fix:"802,omitempty"`
}

func (m *PtysSubGrp) SetNoPartySubIDs(v []NoPartySubIDs) { m.NoPartySubIDs = v }
