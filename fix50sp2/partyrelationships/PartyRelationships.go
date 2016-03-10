package partyrelationships

//NoPartyRelationships is a repeating group in PartyRelationships
type NoPartyRelationships struct {
	//PartyRelationship is a non-required field for NoPartyRelationships.
	PartyRelationship *int `fix:"1515"`
}

//PartyRelationships is a fix50sp2 Component
type PartyRelationships struct {
	//NoPartyRelationships is a non-required field for PartyRelationships.
	NoPartyRelationships []NoPartyRelationships `fix:"1514,omitempty"`
}

func (m *PartyRelationships) SetNoPartyRelationships(v []NoPartyRelationships) {
	m.NoPartyRelationships = v
}
