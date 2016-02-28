package displayinstruction

//Component is a fix50 DisplayInstruction Component
type Component struct {
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

func New() *Component { return new(Component) }

func (m *Component) SetSecondaryDisplayQty(v float64) { m.SecondaryDisplayQty = &v }
func (m *Component) SetDisplayWhen(v string)          { m.DisplayWhen = &v }
func (m *Component) SetDisplayMethod(v string)        { m.DisplayMethod = &v }
func (m *Component) SetDisplayLowQty(v float64)       { m.DisplayLowQty = &v }
func (m *Component) SetDisplayHighQty(v float64)      { m.DisplayHighQty = &v }
func (m *Component) SetDisplayMinIncr(v float64)      { m.DisplayMinIncr = &v }
func (m *Component) SetRefreshQty(v float64)          { m.RefreshQty = &v }
func (m *Component) SetDisplayQty(v float64)          { m.DisplayQty = &v }
