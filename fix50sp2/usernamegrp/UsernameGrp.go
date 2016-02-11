package usernamegrp

//Component is a fix50sp2 UsernameGrp Component
type Component struct {
	//Username is a non-required field for UsernameGrp.
	Username *string `fix:"553"`
}

func New() *Component { return new(Component) }
