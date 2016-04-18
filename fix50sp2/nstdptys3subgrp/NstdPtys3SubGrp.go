package nstdptys3subgrp

//New returns an initialized NstdPtys3SubGrp instance
func New() *NstdPtys3SubGrp {
	var m NstdPtys3SubGrp
	return &m
}

//NoNested3PartySubIDs is a repeating group in NstdPtys3SubGrp
type NoNested3PartySubIDs struct {
	//Nested3PartySubID is a non-required field for NoNested3PartySubIDs.
	Nested3PartySubID *string `fix:"953"`
	//Nested3PartySubIDType is a non-required field for NoNested3PartySubIDs.
	Nested3PartySubIDType *int `fix:"954"`
}

//NewNoNested3PartySubIDs returns an initialized NoNested3PartySubIDs instance
func NewNoNested3PartySubIDs() *NoNested3PartySubIDs {
	var m NoNested3PartySubIDs
	return &m
}

func (m *NoNested3PartySubIDs) SetNested3PartySubID(v string)  { m.Nested3PartySubID = &v }
func (m *NoNested3PartySubIDs) SetNested3PartySubIDType(v int) { m.Nested3PartySubIDType = &v }

//NstdPtys3SubGrp is a fix50sp2 Component
type NstdPtys3SubGrp struct {
	//NoNested3PartySubIDs is a non-required field for NstdPtys3SubGrp.
	NoNested3PartySubIDs []NoNested3PartySubIDs `fix:"952,omitempty"`
}

func (m *NstdPtys3SubGrp) SetNoNested3PartySubIDs(v []NoNested3PartySubIDs) {
	m.NoNested3PartySubIDs = v
}
