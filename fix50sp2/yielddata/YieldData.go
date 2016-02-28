package yielddata

//Component is a fix50sp2 YieldData Component
type Component struct {
	//YieldType is a non-required field for YieldData.
	YieldType *string `fix:"235"`
	//Yield is a non-required field for YieldData.
	Yield *float64 `fix:"236"`
	//YieldCalcDate is a non-required field for YieldData.
	YieldCalcDate *string `fix:"701"`
	//YieldRedemptionDate is a non-required field for YieldData.
	YieldRedemptionDate *string `fix:"696"`
	//YieldRedemptionPrice is a non-required field for YieldData.
	YieldRedemptionPrice *float64 `fix:"697"`
	//YieldRedemptionPriceType is a non-required field for YieldData.
	YieldRedemptionPriceType *int `fix:"698"`
}

func New() *Component { return new(Component) }

func (m *Component) SetYieldType(v string)             { m.YieldType = &v }
func (m *Component) SetYield(v float64)                { m.Yield = &v }
func (m *Component) SetYieldCalcDate(v string)         { m.YieldCalcDate = &v }
func (m *Component) SetYieldRedemptionDate(v string)   { m.YieldRedemptionDate = &v }
func (m *Component) SetYieldRedemptionPrice(v float64) { m.YieldRedemptionPrice = &v }
func (m *Component) SetYieldRedemptionPriceType(v int) { m.YieldRedemptionPriceType = &v }
