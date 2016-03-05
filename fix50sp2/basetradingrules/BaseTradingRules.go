package basetradingrules

//NoTickRules is a repeating group in BaseTradingRules
type NoTickRules struct {
	//StartTickPriceRange is a non-required field for NoTickRules.
	StartTickPriceRange *float64 `fix:"1206"`
	//EndTickPriceRange is a non-required field for NoTickRules.
	EndTickPriceRange *float64 `fix:"1207"`
	//TickIncrement is a non-required field for NoTickRules.
	TickIncrement *float64 `fix:"1208"`
	//TickRuleType is a non-required field for NoTickRules.
	TickRuleType *int `fix:"1209"`
}

//NoLotTypeRules is a repeating group in BaseTradingRules
type NoLotTypeRules struct {
	//LotType is a non-required field for NoLotTypeRules.
	LotType *string `fix:"1093"`
	//MinLotSize is a non-required field for NoLotTypeRules.
	MinLotSize *float64 `fix:"1231"`
}

//BaseTradingRules is a fix50sp2 Component
type BaseTradingRules struct {
	//NoTickRules is a non-required field for BaseTradingRules.
	NoTickRules []NoTickRules `fix:"1205,omitempty"`
	//NoLotTypeRules is a non-required field for BaseTradingRules.
	NoLotTypeRules []NoLotTypeRules `fix:"1234,omitempty"`
	//PriceLimitType is a non-required field for BaseTradingRules.
	PriceLimitType *int `fix:"1306"`
	//LowLimitPrice is a non-required field for BaseTradingRules.
	LowLimitPrice *float64 `fix:"1148"`
	//HighLimitPrice is a non-required field for BaseTradingRules.
	HighLimitPrice *float64 `fix:"1149"`
	//TradingReferencePrice is a non-required field for BaseTradingRules.
	TradingReferencePrice *float64 `fix:"1150"`
	//ExpirationCycle is a non-required field for BaseTradingRules.
	ExpirationCycle *int `fix:"827"`
	//MinTradeVol is a non-required field for BaseTradingRules.
	MinTradeVol *float64 `fix:"562"`
	//MaxTradeVol is a non-required field for BaseTradingRules.
	MaxTradeVol *float64 `fix:"1140"`
	//MaxPriceVariation is a non-required field for BaseTradingRules.
	MaxPriceVariation *float64 `fix:"1143"`
	//ImpliedMarketIndicator is a non-required field for BaseTradingRules.
	ImpliedMarketIndicator *int `fix:"1144"`
	//TradingCurrency is a non-required field for BaseTradingRules.
	TradingCurrency *string `fix:"1245"`
	//RoundLot is a non-required field for BaseTradingRules.
	RoundLot *float64 `fix:"561"`
	//MultilegModel is a non-required field for BaseTradingRules.
	MultilegModel *int `fix:"1377"`
	//MultilegPriceMethod is a non-required field for BaseTradingRules.
	MultilegPriceMethod *int `fix:"1378"`
	//PriceType is a non-required field for BaseTradingRules.
	PriceType *int `fix:"423"`
}

func (m *BaseTradingRules) SetNoTickRules(v []NoTickRules)       { m.NoTickRules = v }
func (m *BaseTradingRules) SetNoLotTypeRules(v []NoLotTypeRules) { m.NoLotTypeRules = v }
func (m *BaseTradingRules) SetPriceLimitType(v int)              { m.PriceLimitType = &v }
func (m *BaseTradingRules) SetLowLimitPrice(v float64)           { m.LowLimitPrice = &v }
func (m *BaseTradingRules) SetHighLimitPrice(v float64)          { m.HighLimitPrice = &v }
func (m *BaseTradingRules) SetTradingReferencePrice(v float64)   { m.TradingReferencePrice = &v }
func (m *BaseTradingRules) SetExpirationCycle(v int)             { m.ExpirationCycle = &v }
func (m *BaseTradingRules) SetMinTradeVol(v float64)             { m.MinTradeVol = &v }
func (m *BaseTradingRules) SetMaxTradeVol(v float64)             { m.MaxTradeVol = &v }
func (m *BaseTradingRules) SetMaxPriceVariation(v float64)       { m.MaxPriceVariation = &v }
func (m *BaseTradingRules) SetImpliedMarketIndicator(v int)      { m.ImpliedMarketIndicator = &v }
func (m *BaseTradingRules) SetTradingCurrency(v string)          { m.TradingCurrency = &v }
func (m *BaseTradingRules) SetRoundLot(v float64)                { m.RoundLot = &v }
func (m *BaseTradingRules) SetMultilegModel(v int)               { m.MultilegModel = &v }
func (m *BaseTradingRules) SetMultilegPriceMethod(v int)         { m.MultilegPriceMethod = &v }
func (m *BaseTradingRules) SetPriceType(v int)                   { m.PriceType = &v }
