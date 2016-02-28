package requestedpartyrolegrp

//NoRequestedPartyRoles is a repeating group in RequestedPartyRoleGrp
type NoRequestedPartyRoles struct {
	//RequestedPartyRole is a non-required field for NoRequestedPartyRoles.
	RequestedPartyRole *int `fix:"1509"`
}

//Component is a fix50sp2 RequestedPartyRoleGrp Component
type Component struct {
	//NoRequestedPartyRoles is a non-required field for RequestedPartyRoleGrp.
	NoRequestedPartyRoles []NoRequestedPartyRoles `fix:"1508,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoRequestedPartyRoles(v []NoRequestedPartyRoles) { m.NoRequestedPartyRoles = v }
