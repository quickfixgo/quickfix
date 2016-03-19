package parties

func New() *Parties {
	var m Parties
	return &m
}

//NoPartyIDs is a repeating group in Parties
type NoPartyIDs struct {
	//PartyID is a non-required field for NoPartyIDs.
	PartyID *string `fix:"448"`
	//PartyIDSource is a non-required field for NoPartyIDs.
	PartyIDSource *string `fix:"447"`
	//PartyRole is a non-required field for NoPartyIDs.
	PartyRole *int `fix:"452"`
	//PartySubID is a non-required field for NoPartyIDs.
	PartySubID *string `fix:"523"`
}

//Parties is a fix43 Component
type Parties struct {
	//NoPartyIDs is a non-required field for Parties.
	NoPartyIDs []NoPartyIDs `fix:"453,omitempty"`
}

func (m *Parties) SetNoPartyIDs(v []NoPartyIDs) { m.NoPartyIDs = v }
