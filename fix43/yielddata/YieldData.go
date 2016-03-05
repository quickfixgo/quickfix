package yielddata

//YieldData is a fix43 Component
type YieldData struct {
	//YieldType is a non-required field for YieldData.
	YieldType *string `fix:"235"`
	//Yield is a non-required field for YieldData.
	Yield *float64 `fix:"236"`
}

func (m *YieldData) SetYieldType(v string) { m.YieldType = &v }
func (m *YieldData) SetYield(v float64)    { m.Yield = &v }
