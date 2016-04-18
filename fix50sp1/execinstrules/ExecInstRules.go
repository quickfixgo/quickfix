package execinstrules

//New returns an initialized ExecInstRules instance
func New() *ExecInstRules {
	var m ExecInstRules
	return &m
}

//NoExecInstRules is a repeating group in ExecInstRules
type NoExecInstRules struct {
	//ExecInstValue is a non-required field for NoExecInstRules.
	ExecInstValue *string `fix:"1308"`
}

//NewNoExecInstRules returns an initialized NoExecInstRules instance
func NewNoExecInstRules() *NoExecInstRules {
	var m NoExecInstRules
	return &m
}

func (m *NoExecInstRules) SetExecInstValue(v string) { m.ExecInstValue = &v }

//ExecInstRules is a fix50sp1 Component
type ExecInstRules struct {
	//NoExecInstRules is a non-required field for ExecInstRules.
	NoExecInstRules []NoExecInstRules `fix:"1232,omitempty"`
}

func (m *ExecInstRules) SetNoExecInstRules(v []NoExecInstRules) { m.NoExecInstRules = v }
