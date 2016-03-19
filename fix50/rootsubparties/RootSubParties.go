package rootsubparties

func New() *RootSubParties {
	var m RootSubParties
	return &m
}

//NoRootPartySubIDs is a repeating group in RootSubParties
type NoRootPartySubIDs struct {
	//RootPartySubID is a non-required field for NoRootPartySubIDs.
	RootPartySubID *string `fix:"1121"`
	//RootPartySubIDType is a non-required field for NoRootPartySubIDs.
	RootPartySubIDType *int `fix:"1122"`
}

//RootSubParties is a fix50 Component
type RootSubParties struct {
	//NoRootPartySubIDs is a non-required field for RootSubParties.
	NoRootPartySubIDs []NoRootPartySubIDs `fix:"1120,omitempty"`
}

func (m *RootSubParties) SetNoRootPartySubIDs(v []NoRootPartySubIDs) { m.NoRootPartySubIDs = v }
