package tradingsessionrules

import (
	"github.com/quickfixgo/quickfix/fix50sp1/execinstrules"
	"github.com/quickfixgo/quickfix/fix50sp1/marketdatafeedtypes"
	"github.com/quickfixgo/quickfix/fix50sp1/matchrules"
	"github.com/quickfixgo/quickfix/fix50sp1/ordtyperules"
	"github.com/quickfixgo/quickfix/fix50sp1/timeinforcerules"
)

//New returns an initialized TradingSessionRules instance
func New() *TradingSessionRules {
	var m TradingSessionRules
	return &m
}

//TradingSessionRules is a fix50sp1 Component
type TradingSessionRules struct {
	//OrdTypeRules is a non-required component for TradingSessionRules.
	OrdTypeRules *ordtyperules.OrdTypeRules
	//TimeInForceRules is a non-required component for TradingSessionRules.
	TimeInForceRules *timeinforcerules.TimeInForceRules
	//ExecInstRules is a non-required component for TradingSessionRules.
	ExecInstRules *execinstrules.ExecInstRules
	//MatchRules is a non-required component for TradingSessionRules.
	MatchRules *matchrules.MatchRules
	//MarketDataFeedTypes is a non-required component for TradingSessionRules.
	MarketDataFeedTypes *marketdatafeedtypes.MarketDataFeedTypes
}

func (m *TradingSessionRules) SetOrdTypeRules(v ordtyperules.OrdTypeRules) { m.OrdTypeRules = &v }
func (m *TradingSessionRules) SetTimeInForceRules(v timeinforcerules.TimeInForceRules) {
	m.TimeInForceRules = &v
}
func (m *TradingSessionRules) SetExecInstRules(v execinstrules.ExecInstRules) { m.ExecInstRules = &v }
func (m *TradingSessionRules) SetMatchRules(v matchrules.MatchRules)          { m.MatchRules = &v }
func (m *TradingSessionRules) SetMarketDataFeedTypes(v marketdatafeedtypes.MarketDataFeedTypes) {
	m.MarketDataFeedTypes = &v
}
