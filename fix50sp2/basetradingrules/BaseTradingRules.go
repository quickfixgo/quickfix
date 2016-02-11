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

//Component is a fix50sp2 BaseTradingRules Component
type Component struct {
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

func New() *Component { return new(Component) }
