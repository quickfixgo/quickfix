package pricelimits

//New returns an initialized PriceLimits instance
func New() *PriceLimits {
	var m PriceLimits
	return &m
}

//PriceLimits is a fix50sp2 Component
type PriceLimits struct {
	//PriceLimitType is a non-required field for PriceLimits.
	PriceLimitType *int `fix:"1306"`
	//LowLimitPrice is a non-required field for PriceLimits.
	LowLimitPrice *float64 `fix:"1148"`
	//HighLimitPrice is a non-required field for PriceLimits.
	HighLimitPrice *float64 `fix:"1149"`
	//TradingReferencePrice is a non-required field for PriceLimits.
	TradingReferencePrice *float64 `fix:"1150"`
}

func (m *PriceLimits) SetPriceLimitType(v int)            { m.PriceLimitType = &v }
func (m *PriceLimits) SetLowLimitPrice(v float64)         { m.LowLimitPrice = &v }
func (m *PriceLimits) SetHighLimitPrice(v float64)        { m.HighLimitPrice = &v }
func (m *PriceLimits) SetTradingReferencePrice(v float64) { m.TradingReferencePrice = &v }
