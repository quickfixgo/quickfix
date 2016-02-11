package settlptyssubgrp

//NoSettlPartySubIDs is a repeating group in SettlPtysSubGrp
type NoSettlPartySubIDs struct {
	//SettlPartySubID is a non-required field for NoSettlPartySubIDs.
	SettlPartySubID *string `fix:"785"`
	//SettlPartySubIDType is a non-required field for NoSettlPartySubIDs.
	SettlPartySubIDType *int `fix:"786"`
}

//Component is a fix50 SettlPtysSubGrp Component
type Component struct {
	//NoSettlPartySubIDs is a non-required field for SettlPtysSubGrp.
	NoSettlPartySubIDs []NoSettlPartySubIDs `fix:"801,omitempty"`
}

func New() *Component { return new(Component) }
