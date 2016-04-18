package secondarypricelimits

//New returns an initialized SecondaryPriceLimits instance
func New() *SecondaryPriceLimits {
	var m SecondaryPriceLimits
	return &m
}

//SecondaryPriceLimits is a fix50sp1 Component
type SecondaryPriceLimits struct {
	//SecondaryPriceLimitType is a non-required field for SecondaryPriceLimits.
	SecondaryPriceLimitType *int `fix:"1305"`
	//SecondaryLowLimitPrice is a non-required field for SecondaryPriceLimits.
	SecondaryLowLimitPrice *float64 `fix:"1221"`
	//SecondaryHighLimitPrice is a non-required field for SecondaryPriceLimits.
	SecondaryHighLimitPrice *float64 `fix:"1230"`
	//SecondaryTradingReferencePrice is a non-required field for SecondaryPriceLimits.
	SecondaryTradingReferencePrice *float64 `fix:"1240"`
}

func (m *SecondaryPriceLimits) SetSecondaryPriceLimitType(v int)     { m.SecondaryPriceLimitType = &v }
func (m *SecondaryPriceLimits) SetSecondaryLowLimitPrice(v float64)  { m.SecondaryLowLimitPrice = &v }
func (m *SecondaryPriceLimits) SetSecondaryHighLimitPrice(v float64) { m.SecondaryHighLimitPrice = &v }
func (m *SecondaryPriceLimits) SetSecondaryTradingReferencePrice(v float64) {
	m.SecondaryTradingReferencePrice = &v
}
