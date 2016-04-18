package nstdptys4subgrp

//New returns an initialized NstdPtys4SubGrp instance
func New() *NstdPtys4SubGrp {
	var m NstdPtys4SubGrp
	return &m
}

//NoNested4PartySubIDs is a repeating group in NstdPtys4SubGrp
type NoNested4PartySubIDs struct {
	//Nested4PartySubID is a non-required field for NoNested4PartySubIDs.
	Nested4PartySubID *string `fix:"1412"`
	//Nested4PartySubIDType is a non-required field for NoNested4PartySubIDs.
	Nested4PartySubIDType *int `fix:"1411"`
}

//NewNoNested4PartySubIDs returns an initialized NoNested4PartySubIDs instance
func NewNoNested4PartySubIDs() *NoNested4PartySubIDs {
	var m NoNested4PartySubIDs
	return &m
}

func (m *NoNested4PartySubIDs) SetNested4PartySubID(v string)  { m.Nested4PartySubID = &v }
func (m *NoNested4PartySubIDs) SetNested4PartySubIDType(v int) { m.Nested4PartySubIDType = &v }

//NstdPtys4SubGrp is a fix50sp1 Component
type NstdPtys4SubGrp struct {
	//NoNested4PartySubIDs is a non-required field for NstdPtys4SubGrp.
	NoNested4PartySubIDs []NoNested4PartySubIDs `fix:"1413,omitempty"`
}

func (m *NstdPtys4SubGrp) SetNoNested4PartySubIDs(v []NoNested4PartySubIDs) {
	m.NoNested4PartySubIDs = v
}
