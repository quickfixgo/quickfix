package nstdptys3subgrp

//NoNested3PartySubIDs is a repeating group in NstdPtys3SubGrp
type NoNested3PartySubIDs struct {
	//Nested3PartySubID is a non-required field for NoNested3PartySubIDs.
	Nested3PartySubID *string `fix:"953"`
	//Nested3PartySubIDType is a non-required field for NoNested3PartySubIDs.
	Nested3PartySubIDType *int `fix:"954"`
}

//Component is a fix50sp1 NstdPtys3SubGrp Component
type Component struct {
	//NoNested3PartySubIDs is a non-required field for NstdPtys3SubGrp.
	NoNested3PartySubIDs []NoNested3PartySubIDs `fix:"952,omitempty"`
}

func New() *Component { return new(Component) }
