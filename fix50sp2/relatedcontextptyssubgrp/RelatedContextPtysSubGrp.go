package relatedcontextptyssubgrp

//New returns an initialized RelatedContextPtysSubGrp instance
func New() *RelatedContextPtysSubGrp {
	var m RelatedContextPtysSubGrp
	return &m
}

//NoRelatedContextPartySubIDs is a repeating group in RelatedContextPtysSubGrp
type NoRelatedContextPartySubIDs struct {
	//RelatedContextPartySubID is a non-required field for NoRelatedContextPartySubIDs.
	RelatedContextPartySubID *string `fix:"1580"`
	//RelatedContextPartySubIDType is a non-required field for NoRelatedContextPartySubIDs.
	RelatedContextPartySubIDType *int `fix:"1581"`
}

//NewNoRelatedContextPartySubIDs returns an initialized NoRelatedContextPartySubIDs instance
func NewNoRelatedContextPartySubIDs() *NoRelatedContextPartySubIDs {
	var m NoRelatedContextPartySubIDs
	return &m
}

func (m *NoRelatedContextPartySubIDs) SetRelatedContextPartySubID(v string) {
	m.RelatedContextPartySubID = &v
}
func (m *NoRelatedContextPartySubIDs) SetRelatedContextPartySubIDType(v int) {
	m.RelatedContextPartySubIDType = &v
}

//RelatedContextPtysSubGrp is a fix50sp2 Component
type RelatedContextPtysSubGrp struct {
	//NoRelatedContextPartySubIDs is a non-required field for RelatedContextPtysSubGrp.
	NoRelatedContextPartySubIDs []NoRelatedContextPartySubIDs `fix:"1579,omitempty"`
}

func (m *RelatedContextPtysSubGrp) SetNoRelatedContextPartySubIDs(v []NoRelatedContextPartySubIDs) {
	m.NoRelatedContextPartySubIDs = v
}
