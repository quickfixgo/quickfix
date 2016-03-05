package nstdptyssubgrp

//NoNestedPartySubIDs is a repeating group in NstdPtysSubGrp
type NoNestedPartySubIDs struct {
	//NestedPartySubID is a non-required field for NoNestedPartySubIDs.
	NestedPartySubID *string `fix:"545"`
	//NestedPartySubIDType is a non-required field for NoNestedPartySubIDs.
	NestedPartySubIDType *int `fix:"805"`
}

//NstdPtysSubGrp is a fix50 Component
type NstdPtysSubGrp struct {
	//NoNestedPartySubIDs is a non-required field for NstdPtysSubGrp.
	NoNestedPartySubIDs []NoNestedPartySubIDs `fix:"804,omitempty"`
}

func (m *NstdPtysSubGrp) SetNoNestedPartySubIDs(v []NoNestedPartySubIDs) { m.NoNestedPartySubIDs = v }
