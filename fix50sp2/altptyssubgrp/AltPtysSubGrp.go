package altptyssubgrp

//NoPartyAltSubIDs is a repeating group in AltPtysSubGrp
type NoPartyAltSubIDs struct {
	//PartyAltSubID is a non-required field for NoPartyAltSubIDs.
	PartyAltSubID *string `fix:"1520"`
	//PartyAltSubIDType is a non-required field for NoPartyAltSubIDs.
	PartyAltSubIDType *int `fix:"1521"`
}

//Component is a fix50sp2 AltPtysSubGrp Component
type Component struct {
	//NoPartyAltSubIDs is a non-required field for AltPtysSubGrp.
	NoPartyAltSubIDs []NoPartyAltSubIDs `fix:"1519,omitempty"`
}

func New() *Component { return new(Component) }
