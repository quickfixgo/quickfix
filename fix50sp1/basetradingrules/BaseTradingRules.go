package basetradingrules

import (
	"github.com/quickfixgo/quickfix/fix50sp1/lottyperules"
	"github.com/quickfixgo/quickfix/fix50sp1/pricelimits"
	"github.com/quickfixgo/quickfix/fix50sp1/tickrules"
)

func New() *BaseTradingRules {
	var m BaseTradingRules
	return &m
}

//BaseTradingRules is a fix50sp1 Component
type BaseTradingRules struct {
	//TickRules is a non-required component for BaseTradingRules.
	TickRules *tickrules.TickRules
	//LotTypeRules is a non-required component for BaseTradingRules.
	LotTypeRules *lottyperules.LotTypeRules
	//PriceLimits is a non-required component for BaseTradingRules.
	PriceLimits *pricelimits.PriceLimits
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

func (m *BaseTradingRules) SetTickRules(v tickrules.TickRules)          { m.TickRules = &v }
func (m *BaseTradingRules) SetLotTypeRules(v lottyperules.LotTypeRules) { m.LotTypeRules = &v }
func (m *BaseTradingRules) SetPriceLimits(v pricelimits.PriceLimits)    { m.PriceLimits = &v }
func (m *BaseTradingRules) SetExpirationCycle(v int)                    { m.ExpirationCycle = &v }
func (m *BaseTradingRules) SetMinTradeVol(v float64)                    { m.MinTradeVol = &v }
func (m *BaseTradingRules) SetMaxTradeVol(v float64)                    { m.MaxTradeVol = &v }
func (m *BaseTradingRules) SetMaxPriceVariation(v float64)              { m.MaxPriceVariation = &v }
func (m *BaseTradingRules) SetImpliedMarketIndicator(v int)             { m.ImpliedMarketIndicator = &v }
func (m *BaseTradingRules) SetTradingCurrency(v string)                 { m.TradingCurrency = &v }
func (m *BaseTradingRules) SetRoundLot(v float64)                       { m.RoundLot = &v }
func (m *BaseTradingRules) SetMultilegModel(v int)                      { m.MultilegModel = &v }
func (m *BaseTradingRules) SetMultilegPriceMethod(v int)                { m.MultilegPriceMethod = &v }
func (m *BaseTradingRules) SetPriceType(v int)                          { m.PriceType = &v }
