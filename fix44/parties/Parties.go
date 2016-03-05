package parties

//NoPartyIDs is a repeating group in Parties
type NoPartyIDs struct {
	//PartyID is a non-required field for NoPartyIDs.
	PartyID *string `fix:"448"`
	//PartyIDSource is a non-required field for NoPartyIDs.
	PartyIDSource *string `fix:"447"`
	//PartyRole is a non-required field for NoPartyIDs.
	PartyRole *int `fix:"452"`
	//NoPartySubIDs is a non-required field for NoPartyIDs.
	NoPartySubIDs []NoPartySubIDs `fix:"802,omitempty"`
}

//NoPartySubIDs is a repeating group in NoPartyIDs
type NoPartySubIDs struct {
	//PartySubID is a non-required field for NoPartySubIDs.
	PartySubID *string `fix:"523"`
	//PartySubIDType is a non-required field for NoPartySubIDs.
	PartySubIDType *int `fix:"803"`
}

//Parties is a fix44 Component
type Parties struct {
	//NoPartyIDs is a non-required field for Parties.
	NoPartyIDs []NoPartyIDs `fix:"453,omitempty"`
}

func (m *Parties) SetNoPartyIDs(v []NoPartyIDs) { m.NoPartyIDs = v }
