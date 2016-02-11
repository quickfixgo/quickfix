package pricelimits

//Component is a fix50sp2 PriceLimits Component
type Component struct {
	//PriceLimitType is a non-required field for PriceLimits.
	PriceLimitType *int `fix:"1306"`
	//LowLimitPrice is a non-required field for PriceLimits.
	LowLimitPrice *float64 `fix:"1148"`
	//HighLimitPrice is a non-required field for PriceLimits.
	HighLimitPrice *float64 `fix:"1149"`
	//TradingReferencePrice is a non-required field for PriceLimits.
	TradingReferencePrice *float64 `fix:"1150"`
}

func New() *Component { return new(Component) }
