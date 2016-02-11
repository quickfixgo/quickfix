package partyrelationships

//NoPartyRelationships is a repeating group in PartyRelationships
type NoPartyRelationships struct {
	//PartyRelationship is a non-required field for NoPartyRelationships.
	PartyRelationship *int `fix:"1515"`
}

//Component is a fix50sp2 PartyRelationships Component
type Component struct {
	//NoPartyRelationships is a non-required field for PartyRelationships.
	NoPartyRelationships []NoPartyRelationships `fix:"1514,omitempty"`
}

func New() *Component { return new(Component) }
