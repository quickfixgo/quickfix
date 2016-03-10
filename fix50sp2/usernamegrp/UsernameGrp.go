package usernamegrp

//UsernameGrp is a fix50sp2 Component
type UsernameGrp struct {
	//Username is a non-required field for UsernameGrp.
	Username *string `fix:"553"`
}

func (m *UsernameGrp) SetUsername(v string) { m.Username = &v }
