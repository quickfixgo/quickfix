package execinstrules

//NoExecInstRules is a repeating group in ExecInstRules
type NoExecInstRules struct {
	//ExecInstValue is a non-required field for NoExecInstRules.
	ExecInstValue *string `fix:"1308"`
}

//ExecInstRules is a fix50sp1 Component
type ExecInstRules struct {
	//NoExecInstRules is a non-required field for ExecInstRules.
	NoExecInstRules []NoExecInstRules `fix:"1232,omitempty"`
}

func (m *ExecInstRules) SetNoExecInstRules(v []NoExecInstRules) { m.NoExecInstRules = v }
