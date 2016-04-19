package altptyssubgrp

//New returns an initialized AltPtysSubGrp instance
func New() *AltPtysSubGrp {
	var m AltPtysSubGrp
	return &m
}

//NoPartyAltSubIDs is a repeating group in AltPtysSubGrp
type NoPartyAltSubIDs struct {
	//PartyAltSubID is a non-required field for NoPartyAltSubIDs.
	PartyAltSubID *string `fix:"1520"`
	//PartyAltSubIDType is a non-required field for NoPartyAltSubIDs.
	PartyAltSubIDType *int `fix:"1521"`
}

//NewNoPartyAltSubIDs returns an initialized NoPartyAltSubIDs instance
func NewNoPartyAltSubIDs() *NoPartyAltSubIDs {
	var m NoPartyAltSubIDs
	return &m
}

func (m *NoPartyAltSubIDs) SetPartyAltSubID(v string)  { m.PartyAltSubID = &v }
func (m *NoPartyAltSubIDs) SetPartyAltSubIDType(v int) { m.PartyAltSubIDType = &v }

//AltPtysSubGrp is a fix50sp2 Component
type AltPtysSubGrp struct {
	//NoPartyAltSubIDs is a non-required field for AltPtysSubGrp.
	NoPartyAltSubIDs []NoPartyAltSubIDs `fix:"1519,omitempty"`
}

func (m *AltPtysSubGrp) SetNoPartyAltSubIDs(v []NoPartyAltSubIDs) { m.NoPartyAltSubIDs = v }
