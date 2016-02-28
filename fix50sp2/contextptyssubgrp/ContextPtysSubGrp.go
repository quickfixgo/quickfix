package contextptyssubgrp

//NoContextPartySubIDs is a repeating group in ContextPtysSubGrp
type NoContextPartySubIDs struct {
	//ContextPartySubID is a non-required field for NoContextPartySubIDs.
	ContextPartySubID *string `fix:"1527"`
	//ContextPartySubIDType is a non-required field for NoContextPartySubIDs.
	ContextPartySubIDType *int `fix:"1528"`
}

//Component is a fix50sp2 ContextPtysSubGrp Component
type Component struct {
	//NoContextPartySubIDs is a non-required field for ContextPtysSubGrp.
	NoContextPartySubIDs []NoContextPartySubIDs `fix:"1526,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoContextPartySubIDs(v []NoContextPartySubIDs) { m.NoContextPartySubIDs = v }
