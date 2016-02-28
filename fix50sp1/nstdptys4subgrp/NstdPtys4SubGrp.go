package nstdptys4subgrp

//NoNested4PartySubIDs is a repeating group in NstdPtys4SubGrp
type NoNested4PartySubIDs struct {
	//Nested4PartySubID is a non-required field for NoNested4PartySubIDs.
	Nested4PartySubID *string `fix:"1412"`
	//Nested4PartySubIDType is a non-required field for NoNested4PartySubIDs.
	Nested4PartySubIDType *int `fix:"1411"`
}

//Component is a fix50sp1 NstdPtys4SubGrp Component
type Component struct {
	//NoNested4PartySubIDs is a non-required field for NstdPtys4SubGrp.
	NoNested4PartySubIDs []NoNested4PartySubIDs `fix:"1413,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoNested4PartySubIDs(v []NoNested4PartySubIDs) { m.NoNested4PartySubIDs = v }
