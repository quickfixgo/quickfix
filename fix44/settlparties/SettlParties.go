package settlparties

//NoSettlPartyIDs is a repeating group in SettlParties
type NoSettlPartyIDs struct {
	//SettlPartyID is a non-required field for NoSettlPartyIDs.
	SettlPartyID *string `fix:"782"`
	//SettlPartyIDSource is a non-required field for NoSettlPartyIDs.
	SettlPartyIDSource *string `fix:"783"`
	//SettlPartyRole is a non-required field for NoSettlPartyIDs.
	SettlPartyRole *int `fix:"784"`
	//NoSettlPartySubIDs is a non-required field for NoSettlPartyIDs.
	NoSettlPartySubIDs []NoSettlPartySubIDs `fix:"801,omitempty"`
}

//NoSettlPartySubIDs is a repeating group in NoSettlPartyIDs
type NoSettlPartySubIDs struct {
	//SettlPartySubID is a non-required field for NoSettlPartySubIDs.
	SettlPartySubID *string `fix:"785"`
	//SettlPartySubIDType is a non-required field for NoSettlPartySubIDs.
	SettlPartySubIDType *int `fix:"786"`
}

//Component is a fix44 SettlParties Component
type Component struct {
	//NoSettlPartyIDs is a non-required field for SettlParties.
	NoSettlPartyIDs []NoSettlPartyIDs `fix:"781,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoSettlPartyIDs(v []NoSettlPartyIDs) { m.NoSettlPartyIDs = v }
