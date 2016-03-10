package nstdptys4subgrp

//NoNested4PartySubIDs is a repeating group in NstdPtys4SubGrp
type NoNested4PartySubIDs struct {
	//Nested4PartySubID is a non-required field for NoNested4PartySubIDs.
	Nested4PartySubID *string `fix:"1412"`
	//Nested4PartySubIDType is a non-required field for NoNested4PartySubIDs.
	Nested4PartySubIDType *int `fix:"1411"`
}

//NstdPtys4SubGrp is a fix50sp1 Component
type NstdPtys4SubGrp struct {
	//NoNested4PartySubIDs is a non-required field for NstdPtys4SubGrp.
	NoNested4PartySubIDs []NoNested4PartySubIDs `fix:"1413,omitempty"`
}

func (m *NstdPtys4SubGrp) SetNoNested4PartySubIDs(v []NoNested4PartySubIDs) {
	m.NoNested4PartySubIDs = v
}
