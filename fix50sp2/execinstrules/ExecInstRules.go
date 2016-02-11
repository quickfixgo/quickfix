package execinstrules

//NoExecInstRules is a repeating group in ExecInstRules
type NoExecInstRules struct {
	//ExecInstValue is a non-required field for NoExecInstRules.
	ExecInstValue *string `fix:"1308"`
}

//Component is a fix50sp2 ExecInstRules Component
type Component struct {
	//NoExecInstRules is a non-required field for ExecInstRules.
	NoExecInstRules []NoExecInstRules `fix:"1232,omitempty"`
}

func New() *Component { return new(Component) }
