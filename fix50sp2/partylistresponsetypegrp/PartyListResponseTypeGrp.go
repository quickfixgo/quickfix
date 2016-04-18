package partylistresponsetypegrp

//New returns an initialized PartyListResponseTypeGrp instance
func New(nopartylistresponsetypes []NoPartyListResponseTypes) *PartyListResponseTypeGrp {
	var m PartyListResponseTypeGrp
	m.SetNoPartyListResponseTypes(nopartylistresponsetypes)
	return &m
}

//NoPartyListResponseTypes is a repeating group in PartyListResponseTypeGrp
type NoPartyListResponseTypes struct {
	//PartyListResponseType is a required field for NoPartyListResponseTypes.
	PartyListResponseType int `fix:"1507"`
}

//NewNoPartyListResponseTypes returns an initialized NoPartyListResponseTypes instance
func NewNoPartyListResponseTypes(partylistresponsetype int) *NoPartyListResponseTypes {
	var m NoPartyListResponseTypes
	m.SetPartyListResponseType(partylistresponsetype)
	return &m
}

func (m *NoPartyListResponseTypes) SetPartyListResponseType(v int) { m.PartyListResponseType = v }

//PartyListResponseTypeGrp is a fix50sp2 Component
type PartyListResponseTypeGrp struct {
	//NoPartyListResponseTypes is a required field for PartyListResponseTypeGrp.
	NoPartyListResponseTypes []NoPartyListResponseTypes `fix:"1506"`
}

func (m *PartyListResponseTypeGrp) SetNoPartyListResponseTypes(v []NoPartyListResponseTypes) {
	m.NoPartyListResponseTypes = v
}
