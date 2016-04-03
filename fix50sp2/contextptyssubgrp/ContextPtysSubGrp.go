package contextptyssubgrp

func New() *ContextPtysSubGrp {
	var m ContextPtysSubGrp
	return &m
}

//NoContextPartySubIDs is a repeating group in ContextPtysSubGrp
type NoContextPartySubIDs struct {
	//ContextPartySubID is a non-required field for NoContextPartySubIDs.
	ContextPartySubID *string `fix:"1527"`
	//ContextPartySubIDType is a non-required field for NoContextPartySubIDs.
	ContextPartySubIDType *int `fix:"1528"`
}

//ContextPtysSubGrp is a fix50sp2 Component
type ContextPtysSubGrp struct {
	//NoContextPartySubIDs is a non-required field for ContextPtysSubGrp.
	NoContextPartySubIDs []NoContextPartySubIDs `fix:"1526,omitempty"`
}

func (m *ContextPtysSubGrp) SetNoContextPartySubIDs(v []NoContextPartySubIDs) {
	m.NoContextPartySubIDs = v
}
