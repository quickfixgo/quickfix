package nstdptyssubgrp

//NoNestedPartySubIDs is a repeating group in NstdPtysSubGrp
type NoNestedPartySubIDs struct {
	//NestedPartySubID is a non-required field for NoNestedPartySubIDs.
	NestedPartySubID *string `fix:"545"`
	//NestedPartySubIDType is a non-required field for NoNestedPartySubIDs.
	NestedPartySubIDType *int `fix:"805"`
}

//Component is a fix50sp2 NstdPtysSubGrp Component
type Component struct {
	//NoNestedPartySubIDs is a non-required field for NstdPtysSubGrp.
	NoNestedPartySubIDs []NoNestedPartySubIDs `fix:"804,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoNestedPartySubIDs(v []NoNestedPartySubIDs) { m.NoNestedPartySubIDs = v }
