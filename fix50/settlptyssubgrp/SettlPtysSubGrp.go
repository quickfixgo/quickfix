package settlptyssubgrp

//New returns an initialized SettlPtysSubGrp instance
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

//NewNoSettlPartySubIDs returns an initialized NoSettlPartySubIDs instance
func NewNoSettlPartySubIDs() *NoSettlPartySubIDs {
	var m NoSettlPartySubIDs
	return &m
}

func (m *NoSettlPartySubIDs) SetSettlPartySubID(v string)  { m.SettlPartySubID = &v }
func (m *NoSettlPartySubIDs) SetSettlPartySubIDType(v int) { m.SettlPartySubIDType = &v }

//SettlPtysSubGrp is a fix50 Component
type SettlPtysSubGrp struct {
	//NoSettlPartySubIDs is a non-required field for SettlPtysSubGrp.
	NoSettlPartySubIDs []NoSettlPartySubIDs `fix:"801,omitempty"`
}

func (m *SettlPtysSubGrp) SetNoSettlPartySubIDs(v []NoSettlPartySubIDs) { m.NoSettlPartySubIDs = v }
