package partylistresponsetypegrp

//NoPartyListResponseTypes is a repeating group in PartyListResponseTypeGrp
type NoPartyListResponseTypes struct {
	//PartyListResponseType is a required field for NoPartyListResponseTypes.
	PartyListResponseType int `fix:"1507"`
}

//Component is a fix50sp2 PartyListResponseTypeGrp Component
type Component struct {
	//NoPartyListResponseTypes is a required field for PartyListResponseTypeGrp.
	NoPartyListResponseTypes []NoPartyListResponseTypes `fix:"1506"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoPartyListResponseTypes(v []NoPartyListResponseTypes) {
	m.NoPartyListResponseTypes = v
}
