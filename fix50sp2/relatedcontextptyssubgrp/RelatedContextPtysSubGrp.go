package relatedcontextptyssubgrp

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
