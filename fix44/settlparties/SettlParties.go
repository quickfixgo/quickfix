package settlparties

func New() *SettlParties {
	var m SettlParties
	return &m
}

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

func (m *NoSettlPartyIDs) SetSettlPartyID(v string)                     { m.SettlPartyID = &v }
func (m *NoSettlPartyIDs) SetSettlPartyIDSource(v string)               { m.SettlPartyIDSource = &v }
func (m *NoSettlPartyIDs) SetSettlPartyRole(v int)                      { m.SettlPartyRole = &v }
func (m *NoSettlPartyIDs) SetNoSettlPartySubIDs(v []NoSettlPartySubIDs) { m.NoSettlPartySubIDs = v }

//NoSettlPartySubIDs is a repeating group in NoSettlPartyIDs
type NoSettlPartySubIDs struct {
	//SettlPartySubID is a non-required field for NoSettlPartySubIDs.
	SettlPartySubID *string `fix:"785"`
	//SettlPartySubIDType is a non-required field for NoSettlPartySubIDs.
	SettlPartySubIDType *int `fix:"786"`
}

func (m *NoSettlPartySubIDs) SetSettlPartySubID(v string)  { m.SettlPartySubID = &v }
func (m *NoSettlPartySubIDs) SetSettlPartySubIDType(v int) { m.SettlPartySubIDType = &v }

//SettlParties is a fix44 Component
type SettlParties struct {
	//NoSettlPartyIDs is a non-required field for SettlParties.
	NoSettlPartyIDs []NoSettlPartyIDs `fix:"781,omitempty"`
}

func (m *SettlParties) SetNoSettlPartyIDs(v []NoSettlPartyIDs) { m.NoSettlPartyIDs = v }
