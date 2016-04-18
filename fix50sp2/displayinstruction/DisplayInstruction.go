package displayinstruction

//New returns an initialized DisplayInstruction instance
func New() *DisplayInstruction {
	var m DisplayInstruction
	return &m
}

//DisplayInstruction is a fix50sp2 Component
type DisplayInstruction struct {
	//SecondaryDisplayQty is a non-required field for DisplayInstruction.
	SecondaryDisplayQty *float64 `fix:"1082"`
	//DisplayWhen is a non-required field for DisplayInstruction.
	DisplayWhen *string `fix:"1083"`
	//DisplayMethod is a non-required field for DisplayInstruction.
	DisplayMethod *string `fix:"1084"`
	//DisplayLowQty is a non-required field for DisplayInstruction.
	DisplayLowQty *float64 `fix:"1085"`
	//DisplayHighQty is a non-required field for DisplayInstruction.
	DisplayHighQty *float64 `fix:"1086"`
	//DisplayMinIncr is a non-required field for DisplayInstruction.
	DisplayMinIncr *float64 `fix:"1087"`
	//RefreshQty is a non-required field for DisplayInstruction.
	RefreshQty *float64 `fix:"1088"`
	//DisplayQty is a non-required field for DisplayInstruction.
	DisplayQty *float64 `fix:"1138"`
}

func (m *DisplayInstruction) SetSecondaryDisplayQty(v float64) { m.SecondaryDisplayQty = &v }
func (m *DisplayInstruction) SetDisplayWhen(v string)          { m.DisplayWhen = &v }
func (m *DisplayInstruction) SetDisplayMethod(v string)        { m.DisplayMethod = &v }
func (m *DisplayInstruction) SetDisplayLowQty(v float64)       { m.DisplayLowQty = &v }
func (m *DisplayInstruction) SetDisplayHighQty(v float64)      { m.DisplayHighQty = &v }
func (m *DisplayInstruction) SetDisplayMinIncr(v float64)      { m.DisplayMinIncr = &v }
func (m *DisplayInstruction) SetRefreshQty(v float64)          { m.RefreshQty = &v }
func (m *DisplayInstruction) SetDisplayQty(v float64)          { m.DisplayQty = &v }
