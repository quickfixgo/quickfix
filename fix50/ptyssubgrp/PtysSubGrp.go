package ptyssubgrp

//NoPartySubIDs is a repeating group in PtysSubGrp
type NoPartySubIDs struct {
	//PartySubID is a non-required field for NoPartySubIDs.
	PartySubID *string `fix:"523"`
	//PartySubIDType is a non-required field for NoPartySubIDs.
	PartySubIDType *int `fix:"803"`
}

//Component is a fix50 PtysSubGrp Component
type Component struct {
	//NoPartySubIDs is a non-required field for PtysSubGrp.
	NoPartySubIDs []NoPartySubIDs `fix:"802,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoPartySubIDs(v []NoPartySubIDs) { m.NoPartySubIDs = v }
