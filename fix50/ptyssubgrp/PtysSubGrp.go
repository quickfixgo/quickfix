package ptyssubgrp

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

//PtysSubGrp is a fix50 Component
type PtysSubGrp struct {
	//NoPartySubIDs is a non-required field for PtysSubGrp.
	NoPartySubIDs []NoPartySubIDs `fix:"802,omitempty"`
}

func (m *PtysSubGrp) SetNoPartySubIDs(v []NoPartySubIDs) { m.NoPartySubIDs = v }
