package yielddata

//Component is a fix43 YieldData Component
type Component struct {
	//YieldType is a non-required field for YieldData.
	YieldType *string `fix:"235"`
	//Yield is a non-required field for YieldData.
	Yield *float64 `fix:"236"`
}

func New() *Component { return new(Component) }

func (m *Component) SetYieldType(v string) { m.YieldType = &v }
func (m *Component) SetYield(v float64)    { m.Yield = &v }
