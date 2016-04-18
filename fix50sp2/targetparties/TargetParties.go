package targetparties

//New returns an initialized TargetParties instance
func New() *TargetParties {
	var m TargetParties
	return &m
}

//NoTargetPartyIDs is a repeating group in TargetParties
type NoTargetPartyIDs struct {
	//TargetPartyID is a non-required field for NoTargetPartyIDs.
	TargetPartyID *string `fix:"1462"`
	//TargetPartyIDSource is a non-required field for NoTargetPartyIDs.
	TargetPartyIDSource *string `fix:"1463"`
	//TargetPartyRole is a non-required field for NoTargetPartyIDs.
	TargetPartyRole *int `fix:"1464"`
}

//NewNoTargetPartyIDs returns an initialized NoTargetPartyIDs instance
func NewNoTargetPartyIDs() *NoTargetPartyIDs {
	var m NoTargetPartyIDs
	return &m
}

func (m *NoTargetPartyIDs) SetTargetPartyID(v string)       { m.TargetPartyID = &v }
func (m *NoTargetPartyIDs) SetTargetPartyIDSource(v string) { m.TargetPartyIDSource = &v }
func (m *NoTargetPartyIDs) SetTargetPartyRole(v int)        { m.TargetPartyRole = &v }

//TargetParties is a fix50sp2 Component
type TargetParties struct {
	//NoTargetPartyIDs is a non-required field for TargetParties.
	NoTargetPartyIDs []NoTargetPartyIDs `fix:"1461,omitempty"`
}

func (m *TargetParties) SetNoTargetPartyIDs(v []NoTargetPartyIDs) { m.NoTargetPartyIDs = v }
