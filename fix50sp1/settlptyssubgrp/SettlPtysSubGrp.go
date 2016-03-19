package settlptyssubgrp

func New() *SettlPtysSubGrp {
	var m SettlPtysSubGrp
	return &m
}

//NoSettlPartySubIDs is a repeating group in SettlPtysSubGrp
type NoSettlPartySubIDs struct {
	//SettlPartySubID is a non-required field for NoSettlPartySubIDs.
	SettlPartySubID *string `fix:"785"`
	//SettlPartySubIDType is a non-required field for NoSettlPartySubIDs.
	SettlPartySubIDType *int `fix:"786"`
}

//SettlPtysSubGrp is a fix50sp1 Component
type SettlPtysSubGrp struct {
	//NoSettlPartySubIDs is a non-required field for SettlPtysSubGrp.
	NoSettlPartySubIDs []NoSettlPartySubIDs `fix:"801,omitempty"`
}

func (m *SettlPtysSubGrp) SetNoSettlPartySubIDs(v []NoSettlPartySubIDs) { m.NoSettlPartySubIDs = v }
