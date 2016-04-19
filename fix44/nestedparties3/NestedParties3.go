package nestedparties3

//New returns an initialized NestedParties3 instance
func New() *NestedParties3 {
	var m NestedParties3
	return &m
}

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

//NewNoNested3PartyIDs returns an initialized NoNested3PartyIDs instance
func NewNoNested3PartyIDs() *NoNested3PartyIDs {
	var m NoNested3PartyIDs
	return &m
}

func (m *NoNested3PartyIDs) SetNested3PartyID(v string)       { m.Nested3PartyID = &v }
func (m *NoNested3PartyIDs) SetNested3PartyIDSource(v string) { m.Nested3PartyIDSource = &v }
func (m *NoNested3PartyIDs) SetNested3PartyRole(v int)        { m.Nested3PartyRole = &v }
func (m *NoNested3PartyIDs) SetNoNested3PartySubIDs(v []NoNested3PartySubIDs) {
	m.NoNested3PartySubIDs = v
}

//NoNested3PartySubIDs is a repeating group in NoNested3PartyIDs
type NoNested3PartySubIDs struct {
	//Nested3PartySubID is a non-required field for NoNested3PartySubIDs.
	Nested3PartySubID *string `fix:"953"`
	//Nested3PartySubIDType is a non-required field for NoNested3PartySubIDs.
	Nested3PartySubIDType *int `fix:"954"`
}

//NewNoNested3PartySubIDs returns an initialized NoNested3PartySubIDs instance
func NewNoNested3PartySubIDs() *NoNested3PartySubIDs {
	var m NoNested3PartySubIDs
	return &m
}

func (m *NoNested3PartySubIDs) SetNested3PartySubID(v string)  { m.Nested3PartySubID = &v }
func (m *NoNested3PartySubIDs) SetNested3PartySubIDType(v int) { m.Nested3PartySubIDType = &v }

//NestedParties3 is a fix44 Component
type NestedParties3 struct {
	//NoNested3PartyIDs is a non-required field for NestedParties3.
	NoNested3PartyIDs []NoNested3PartyIDs `fix:"948,omitempty"`
}

func (m *NestedParties3) SetNoNested3PartyIDs(v []NoNested3PartyIDs) { m.NoNested3PartyIDs = v }
