package nestedparties

//NoNestedPartyIDs is a repeating group in NestedParties
type NoNestedPartyIDs struct {
	//NestedPartyID is a non-required field for NoNestedPartyIDs.
	NestedPartyID *string `fix:"524"`
	//NestedPartyIDSource is a non-required field for NoNestedPartyIDs.
	NestedPartyIDSource *string `fix:"525"`
	//NestedPartyRole is a non-required field for NoNestedPartyIDs.
	NestedPartyRole *int `fix:"538"`
	//NoNestedPartySubIDs is a non-required field for NoNestedPartyIDs.
	NoNestedPartySubIDs []NoNestedPartySubIDs `fix:"804,omitempty"`
}

//NoNestedPartySubIDs is a repeating group in NoNestedPartyIDs
type NoNestedPartySubIDs struct {
	//NestedPartySubID is a non-required field for NoNestedPartySubIDs.
	NestedPartySubID *string `fix:"545"`
	//NestedPartySubIDType is a non-required field for NoNestedPartySubIDs.
	NestedPartySubIDType *int `fix:"805"`
}

//Component is a fix44 NestedParties Component
type Component struct {
	//NoNestedPartyIDs is a non-required field for NestedParties.
	NoNestedPartyIDs []NoNestedPartyIDs `fix:"539,omitempty"`
}

func New() *Component { return new(Component) }
