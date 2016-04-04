package nestedparties

func New() *NestedParties {
	var m NestedParties
	return &m
}

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

func (m *NoNestedPartyIDs) SetNestedPartyID(v string)                      { m.NestedPartyID = &v }
func (m *NoNestedPartyIDs) SetNestedPartyIDSource(v string)                { m.NestedPartyIDSource = &v }
func (m *NoNestedPartyIDs) SetNestedPartyRole(v int)                       { m.NestedPartyRole = &v }
func (m *NoNestedPartyIDs) SetNoNestedPartySubIDs(v []NoNestedPartySubIDs) { m.NoNestedPartySubIDs = v }

//NoNestedPartySubIDs is a repeating group in NoNestedPartyIDs
type NoNestedPartySubIDs struct {
	//NestedPartySubID is a non-required field for NoNestedPartySubIDs.
	NestedPartySubID *string `fix:"545"`
	//NestedPartySubIDType is a non-required field for NoNestedPartySubIDs.
	NestedPartySubIDType *int `fix:"805"`
}

func (m *NoNestedPartySubIDs) SetNestedPartySubID(v string)  { m.NestedPartySubID = &v }
func (m *NoNestedPartySubIDs) SetNestedPartySubIDType(v int) { m.NestedPartySubIDType = &v }

//NestedParties is a fix44 Component
type NestedParties struct {
	//NoNestedPartyIDs is a non-required field for NestedParties.
	NoNestedPartyIDs []NoNestedPartyIDs `fix:"539,omitempty"`
}

func (m *NestedParties) SetNoNestedPartyIDs(v []NoNestedPartyIDs) { m.NoNestedPartyIDs = v }
