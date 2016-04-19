package usernamegrp

//New returns an initialized UsernameGrp instance
func New() *UsernameGrp {
	var m UsernameGrp
	return &m
}

//UsernameGrp is a fix50sp1 Component
type UsernameGrp struct {
	//Username is a non-required field for UsernameGrp.
	Username *string `fix:"553"`
}

func (m *UsernameGrp) SetUsername(v string) { m.Username = &v }
