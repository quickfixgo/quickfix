package targetparties

//NoTargetPartyIDs is a repeating group in TargetParties
type NoTargetPartyIDs struct {
	//TargetPartyID is a non-required field for NoTargetPartyIDs.
	TargetPartyID *string `fix:"1462"`
	//TargetPartyIDSource is a non-required field for NoTargetPartyIDs.
	TargetPartyIDSource *string `fix:"1463"`
	//TargetPartyRole is a non-required field for NoTargetPartyIDs.
	TargetPartyRole *int `fix:"1464"`
}

//TargetParties is a fix50sp2 Component
type TargetParties struct {
	//NoTargetPartyIDs is a non-required field for TargetParties.
	NoTargetPartyIDs []NoTargetPartyIDs `fix:"1461,omitempty"`
}

func (m *TargetParties) SetNoTargetPartyIDs(v []NoTargetPartyIDs) { m.NoTargetPartyIDs = v }
