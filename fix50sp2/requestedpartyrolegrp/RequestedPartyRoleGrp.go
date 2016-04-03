package requestedpartyrolegrp

func New() *RequestedPartyRoleGrp {
	var m RequestedPartyRoleGrp
	return &m
}

//NoRequestedPartyRoles is a repeating group in RequestedPartyRoleGrp
type NoRequestedPartyRoles struct {
	//RequestedPartyRole is a non-required field for NoRequestedPartyRoles.
	RequestedPartyRole *int `fix:"1509"`
}

//RequestedPartyRoleGrp is a fix50sp2 Component
type RequestedPartyRoleGrp struct {
	//NoRequestedPartyRoles is a non-required field for RequestedPartyRoleGrp.
	NoRequestedPartyRoles []NoRequestedPartyRoles `fix:"1508,omitempty"`
}

func (m *RequestedPartyRoleGrp) SetNoRequestedPartyRoles(v []NoRequestedPartyRoles) {
	m.NoRequestedPartyRoles = v
}
