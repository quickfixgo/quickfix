package nstdptyssubgrp

//New returns an initialized NstdPtysSubGrp instance
func New() *NstdPtysSubGrp {
	var m NstdPtysSubGrp
	return &m
}

//NoNestedPartySubIDs is a repeating group in NstdPtysSubGrp
type NoNestedPartySubIDs struct {
	//NestedPartySubID is a non-required field for NoNestedPartySubIDs.
	NestedPartySubID *string `fix:"545"`
	//NestedPartySubIDType is a non-required field for NoNestedPartySubIDs.
	NestedPartySubIDType *int `fix:"805"`
}

//NewNoNestedPartySubIDs returns an initialized NoNestedPartySubIDs instance
func NewNoNestedPartySubIDs() *NoNestedPartySubIDs {
	var m NoNestedPartySubIDs
	return &m
}

func (m *NoNestedPartySubIDs) SetNestedPartySubID(v string)  { m.NestedPartySubID = &v }
func (m *NoNestedPartySubIDs) SetNestedPartySubIDType(v int) { m.NestedPartySubIDType = &v }

//NstdPtysSubGrp is a fix50 Component
type NstdPtysSubGrp struct {
	//NoNestedPartySubIDs is a non-required field for NstdPtysSubGrp.
	NoNestedPartySubIDs []NoNestedPartySubIDs `fix:"804,omitempty"`
}

func (m *NstdPtysSubGrp) SetNoNestedPartySubIDs(v []NoNestedPartySubIDs) { m.NoNestedPartySubIDs = v }
