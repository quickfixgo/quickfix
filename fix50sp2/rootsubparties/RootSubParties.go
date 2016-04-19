package rootsubparties

//New returns an initialized RootSubParties instance
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

//NewNoRootPartySubIDs returns an initialized NoRootPartySubIDs instance
func NewNoRootPartySubIDs() *NoRootPartySubIDs {
	var m NoRootPartySubIDs
	return &m
}

func (m *NoRootPartySubIDs) SetRootPartySubID(v string)  { m.RootPartySubID = &v }
func (m *NoRootPartySubIDs) SetRootPartySubIDType(v int) { m.RootPartySubIDType = &v }

//RootSubParties is a fix50sp2 Component
type RootSubParties struct {
	//NoRootPartySubIDs is a non-required field for RootSubParties.
	NoRootPartySubIDs []NoRootPartySubIDs `fix:"1120,omitempty"`
}

func (m *RootSubParties) SetNoRootPartySubIDs(v []NoRootPartySubIDs) { m.NoRootPartySubIDs = v }
