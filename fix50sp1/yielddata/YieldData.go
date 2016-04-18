package yielddata

//New returns an initialized YieldData instance
func New() *YieldData {
	var m YieldData
	return &m
}

//YieldData is a fix50sp1 Component
type YieldData struct {
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

func (m *YieldData) SetYieldType(v string)             { m.YieldType = &v }
func (m *YieldData) SetYield(v float64)                { m.Yield = &v }
func (m *YieldData) SetYieldCalcDate(v string)         { m.YieldCalcDate = &v }
func (m *YieldData) SetYieldRedemptionDate(v string)   { m.YieldRedemptionDate = &v }
func (m *YieldData) SetYieldRedemptionPrice(v float64) { m.YieldRedemptionPrice = &v }
func (m *YieldData) SetYieldRedemptionPriceType(v int) { m.YieldRedemptionPriceType = &v }
