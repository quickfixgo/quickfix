package securitytradingrules

import (
	"github.com/quickfixgo/quickfix/fix50sp2/tradingsessionrules"
)

//NoTickRules is a repeating group in SecurityTradingRules
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

//NoLotTypeRules is a repeating group in SecurityTradingRules
type NoLotTypeRules struct {
	//LotType is a non-required field for NoLotTypeRules.
	LotType *string `fix:"1093"`
	//MinLotSize is a non-required field for NoLotTypeRules.
	MinLotSize *float64 `fix:"1231"`
}

//NoTradingSessionRules is a repeating group in SecurityTradingRules
type NoTradingSessionRules struct {
	//TradingSessionID is a non-required field for NoTradingSessionRules.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessionRules.
	TradingSessionSubID *string `fix:"625"`
	//TradingSessionRules Component
	tradingsessionrules.TradingSessionRules
}

//NoNestedInstrAttrib is a repeating group in SecurityTradingRules
type NoNestedInstrAttrib struct {
	//NestedInstrAttribType is a non-required field for NoNestedInstrAttrib.
	NestedInstrAttribType *int `fix:"1210"`
	//NestedInstrAttribValue is a non-required field for NoNestedInstrAttrib.
	NestedInstrAttribValue *string `fix:"1211"`
}

//SecurityTradingRules is a fix50sp2 Component
type SecurityTradingRules struct {
	//NoTickRules is a non-required field for SecurityTradingRules.
	NoTickRules []NoTickRules `fix:"1205,omitempty"`
	//NoLotTypeRules is a non-required field for SecurityTradingRules.
	NoLotTypeRules []NoLotTypeRules `fix:"1234,omitempty"`
	//PriceLimitType is a non-required field for SecurityTradingRules.
	PriceLimitType *int `fix:"1306"`
	//LowLimitPrice is a non-required field for SecurityTradingRules.
	LowLimitPrice *float64 `fix:"1148"`
	//HighLimitPrice is a non-required field for SecurityTradingRules.
	HighLimitPrice *float64 `fix:"1149"`
	//TradingReferencePrice is a non-required field for SecurityTradingRules.
	TradingReferencePrice *float64 `fix:"1150"`
	//ExpirationCycle is a non-required field for SecurityTradingRules.
	ExpirationCycle *int `fix:"827"`
	//MinTradeVol is a non-required field for SecurityTradingRules.
	MinTradeVol *float64 `fix:"562"`
	//MaxTradeVol is a non-required field for SecurityTradingRules.
	MaxTradeVol *float64 `fix:"1140"`
	//MaxPriceVariation is a non-required field for SecurityTradingRules.
	MaxPriceVariation *float64 `fix:"1143"`
	//ImpliedMarketIndicator is a non-required field for SecurityTradingRules.
	ImpliedMarketIndicator *int `fix:"1144"`
	//TradingCurrency is a non-required field for SecurityTradingRules.
	TradingCurrency *string `fix:"1245"`
	//RoundLot is a non-required field for SecurityTradingRules.
	RoundLot *float64 `fix:"561"`
	//MultilegModel is a non-required field for SecurityTradingRules.
	MultilegModel *int `fix:"1377"`
	//MultilegPriceMethod is a non-required field for SecurityTradingRules.
	MultilegPriceMethod *int `fix:"1378"`
	//PriceType is a non-required field for SecurityTradingRules.
	PriceType *int `fix:"423"`
	//NoTradingSessionRules is a non-required field for SecurityTradingRules.
	NoTradingSessionRules []NoTradingSessionRules `fix:"1309,omitempty"`
	//NoNestedInstrAttrib is a non-required field for SecurityTradingRules.
	NoNestedInstrAttrib []NoNestedInstrAttrib `fix:"1312,omitempty"`
}

func (m *SecurityTradingRules) SetNoTickRules(v []NoTickRules)       { m.NoTickRules = v }
func (m *SecurityTradingRules) SetNoLotTypeRules(v []NoLotTypeRules) { m.NoLotTypeRules = v }
func (m *SecurityTradingRules) SetPriceLimitType(v int)              { m.PriceLimitType = &v }
func (m *SecurityTradingRules) SetLowLimitPrice(v float64)           { m.LowLimitPrice = &v }
func (m *SecurityTradingRules) SetHighLimitPrice(v float64)          { m.HighLimitPrice = &v }
func (m *SecurityTradingRules) SetTradingReferencePrice(v float64)   { m.TradingReferencePrice = &v }
func (m *SecurityTradingRules) SetExpirationCycle(v int)             { m.ExpirationCycle = &v }
func (m *SecurityTradingRules) SetMinTradeVol(v float64)             { m.MinTradeVol = &v }
func (m *SecurityTradingRules) SetMaxTradeVol(v float64)             { m.MaxTradeVol = &v }
func (m *SecurityTradingRules) SetMaxPriceVariation(v float64)       { m.MaxPriceVariation = &v }
func (m *SecurityTradingRules) SetImpliedMarketIndicator(v int)      { m.ImpliedMarketIndicator = &v }
func (m *SecurityTradingRules) SetTradingCurrency(v string)          { m.TradingCurrency = &v }
func (m *SecurityTradingRules) SetRoundLot(v float64)                { m.RoundLot = &v }
func (m *SecurityTradingRules) SetMultilegModel(v int)               { m.MultilegModel = &v }
func (m *SecurityTradingRules) SetMultilegPriceMethod(v int)         { m.MultilegPriceMethod = &v }
func (m *SecurityTradingRules) SetPriceType(v int)                   { m.PriceType = &v }
func (m *SecurityTradingRules) SetNoTradingSessionRules(v []NoTradingSessionRules) {
	m.NoTradingSessionRules = v
}
func (m *SecurityTradingRules) SetNoNestedInstrAttrib(v []NoNestedInstrAttrib) {
	m.NoNestedInstrAttrib = v
}
