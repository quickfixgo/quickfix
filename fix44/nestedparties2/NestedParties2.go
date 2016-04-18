package nestedparties2

//New returns an initialized NestedParties2 instance
func New() *NestedParties2 {
	var m NestedParties2
	return &m
}

//NoNested2PartyIDs is a repeating group in NestedParties2
type NoNested2PartyIDs struct {
	//Nested2PartyID is a non-required field for NoNested2PartyIDs.
	Nested2PartyID *string `fix:"757"`
	//Nested2PartyIDSource is a non-required field for NoNested2PartyIDs.
	Nested2PartyIDSource *string `fix:"758"`
	//Nested2PartyRole is a non-required field for NoNested2PartyIDs.
	Nested2PartyRole *int `fix:"759"`
	//NoNested2PartySubIDs is a non-required field for NoNested2PartyIDs.
	NoNested2PartySubIDs []NoNested2PartySubIDs `fix:"806,omitempty"`
}

//NewNoNested2PartyIDs returns an initialized NoNested2PartyIDs instance
func NewNoNested2PartyIDs() *NoNested2PartyIDs {
	var m NoNested2PartyIDs
	return &m
}

func (m *NoNested2PartyIDs) SetNested2PartyID(v string)       { m.Nested2PartyID = &v }
func (m *NoNested2PartyIDs) SetNested2PartyIDSource(v string) { m.Nested2PartyIDSource = &v }
func (m *NoNested2PartyIDs) SetNested2PartyRole(v int)        { m.Nested2PartyRole = &v }
func (m *NoNested2PartyIDs) SetNoNested2PartySubIDs(v []NoNested2PartySubIDs) {
	m.NoNested2PartySubIDs = v
}

//NoNested2PartySubIDs is a repeating group in NoNested2PartyIDs
type NoNested2PartySubIDs struct {
	//Nested2PartySubID is a non-required field for NoNested2PartySubIDs.
	Nested2PartySubID *string `fix:"760"`
	//Nested2PartySubIDType is a non-required field for NoNested2PartySubIDs.
	Nested2PartySubIDType *int `fix:"807"`
}

//NewNoNested2PartySubIDs returns an initialized NoNested2PartySubIDs instance
func NewNoNested2PartySubIDs() *NoNested2PartySubIDs {
	var m NoNested2PartySubIDs
	return &m
}

func (m *NoNested2PartySubIDs) SetNested2PartySubID(v string)  { m.Nested2PartySubID = &v }
func (m *NoNested2PartySubIDs) SetNested2PartySubIDType(v int) { m.Nested2PartySubIDType = &v }

//NestedParties2 is a fix44 Component
type NestedParties2 struct {
	//NoNested2PartyIDs is a non-required field for NestedParties2.
	NoNested2PartyIDs []NoNested2PartyIDs `fix:"756,omitempty"`
}

func (m *NestedParties2) SetNoNested2PartyIDs(v []NoNested2PartyIDs) { m.NoNested2PartyIDs = v }
