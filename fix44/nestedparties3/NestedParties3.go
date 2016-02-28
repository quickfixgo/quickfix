package nestedparties3

//NoNested3PartyIDs is a repeating group in NestedParties3
type NoNested3PartyIDs struct {
	//Nested3PartyID is a non-required field for NoNested3PartyIDs.
	Nested3PartyID *string `fix:"949"`
	//Nested3PartyIDSource is a non-required field for NoNested3PartyIDs.
	Nested3PartyIDSource *string `fix:"950"`
	//Nested3PartyRole is a non-required field for NoNested3PartyIDs.
	Nested3PartyRole *int `fix:"951"`
	//NoNested3PartySubIDs is a non-required field for NoNested3PartyIDs.
	NoNested3PartySubIDs []NoNested3PartySubIDs `fix:"952,omitempty"`
}

//NoNested3PartySubIDs is a repeating group in NoNested3PartyIDs
type NoNested3PartySubIDs struct {
	//Nested3PartySubID is a non-required field for NoNested3PartySubIDs.
	Nested3PartySubID *string `fix:"953"`
	//Nested3PartySubIDType is a non-required field for NoNested3PartySubIDs.
	Nested3PartySubIDType *int `fix:"954"`
}

//Component is a fix44 NestedParties3 Component
type Component struct {
	//NoNested3PartyIDs is a non-required field for NestedParties3.
	NoNested3PartyIDs []NoNested3PartyIDs `fix:"948,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoNested3PartyIDs(v []NoNested3PartyIDs) { m.NoNested3PartyIDs = v }
