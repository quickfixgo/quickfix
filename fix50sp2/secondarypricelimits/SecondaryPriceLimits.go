package secondarypricelimits

//Component is a fix50sp2 SecondaryPriceLimits Component
type Component struct {
	//SecondaryPriceLimitType is a non-required field for SecondaryPriceLimits.
	SecondaryPriceLimitType *int `fix:"1305"`
	//SecondaryLowLimitPrice is a non-required field for SecondaryPriceLimits.
	SecondaryLowLimitPrice *float64 `fix:"1221"`
	//SecondaryHighLimitPrice is a non-required field for SecondaryPriceLimits.
	SecondaryHighLimitPrice *float64 `fix:"1230"`
	//SecondaryTradingReferencePrice is a non-required field for SecondaryPriceLimits.
	SecondaryTradingReferencePrice *float64 `fix:"1240"`
}

func New() *Component { return new(Component) }

func (m *Component) SetSecondaryPriceLimitType(v int)     { m.SecondaryPriceLimitType = &v }
func (m *Component) SetSecondaryLowLimitPrice(v float64)  { m.SecondaryLowLimitPrice = &v }
func (m *Component) SetSecondaryHighLimitPrice(v float64) { m.SecondaryHighLimitPrice = &v }
func (m *Component) SetSecondaryTradingReferencePrice(v float64) {
	m.SecondaryTradingReferencePrice = &v
}
