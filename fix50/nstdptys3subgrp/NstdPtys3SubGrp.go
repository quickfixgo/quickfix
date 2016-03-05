package nstdptys3subgrp

//NoNested3PartySubIDs is a repeating group in NstdPtys3SubGrp
type NoNested3PartySubIDs struct {
	//Nested3PartySubID is a non-required field for NoNested3PartySubIDs.
	Nested3PartySubID *string `fix:"953"`
	//Nested3PartySubIDType is a non-required field for NoNested3PartySubIDs.
	Nested3PartySubIDType *int `fix:"954"`
}

//NstdPtys3SubGrp is a fix50 Component
type NstdPtys3SubGrp struct {
	//NoNested3PartySubIDs is a non-required field for NstdPtys3SubGrp.
	NoNested3PartySubIDs []NoNested3PartySubIDs `fix:"952,omitempty"`
}

func (m *NstdPtys3SubGrp) SetNoNested3PartySubIDs(v []NoNested3PartySubIDs) {
	m.NoNested3PartySubIDs = v
}
