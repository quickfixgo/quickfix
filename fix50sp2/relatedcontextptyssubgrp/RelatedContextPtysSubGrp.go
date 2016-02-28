package relatedcontextptyssubgrp

//NoRelatedContextPartySubIDs is a repeating group in RelatedContextPtysSubGrp
type NoRelatedContextPartySubIDs struct {
	//RelatedContextPartySubID is a non-required field for NoRelatedContextPartySubIDs.
	RelatedContextPartySubID *string `fix:"1580"`
	//RelatedContextPartySubIDType is a non-required field for NoRelatedContextPartySubIDs.
	RelatedContextPartySubIDType *int `fix:"1581"`
}

//Component is a fix50sp2 RelatedContextPtysSubGrp Component
type Component struct {
	//NoRelatedContextPartySubIDs is a non-required field for RelatedContextPtysSubGrp.
	NoRelatedContextPartySubIDs []NoRelatedContextPartySubIDs `fix:"1579,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRelatedContextPartySubIDs(v []NoRelatedContextPartySubIDs) {
	m.NoRelatedContextPartySubIDs = v
}
