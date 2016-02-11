package yielddata

//Component is a fix50 YieldData Component
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
