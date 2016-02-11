package rootsubparties

//NoRootPartySubIDs is a repeating group in RootSubParties
type NoRootPartySubIDs struct {
	//RootPartySubID is a non-required field for NoRootPartySubIDs.
	RootPartySubID *string `fix:"1121"`
	//RootPartySubIDType is a non-required field for NoRootPartySubIDs.
	RootPartySubIDType *int `fix:"1122"`
}

//Component is a fix50sp1 RootSubParties Component
type Component struct {
	//NoRootPartySubIDs is a non-required field for RootSubParties.
	NoRootPartySubIDs []NoRootPartySubIDs `fix:"1120,omitempty"`
}

func New() *Component { return new(Component) }
